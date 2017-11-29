package bliss

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"poly"
	"sampler"
)

type BlissSignature struct {
	z1 *poly.PolyArray
	z2 *poly.PolyArray
	c  []uint32
}

func (sig *BlissSignature) String() string {
	return fmt.Sprintf("{z1:%s,z2:%s,c:%d}",
		sig.z1.String(), sig.z2.String(), sig.c)
}

func computeC(kappa uint32, u *poly.PolyArray, hash []byte) []uint32 {
	indices := make([]uint32, kappa)
	data := u.GetData()
	n := len(data)
	for i := 0; i < n; i++ {
		hash = append(hash, byte(data[i]&0xff))
		hash = append(hash, byte((data[i]>>8)&0xff))
	}
	for tries := 0; tries < 256; tries++ {
		hash[len(hash)-1]++
		whash := sha3.Sum512(hash)
		array := make([]bool, n)
		if n == 256 {
			// BLISS_B_0: we need kappa indices of 8 bits
			i := 0
			for j := 0; j < int(sampler.SHA_512_DIGEST_LENGTH); j++ {
				index := whash[j]
				if !array[index] {
					indices[i] = uint32(index)
					array[index] = true
					i++
					if i >= int(kappa) {
						return indices
					}
				}
			}
		} else {
			// BLIS_B_1234: we need kappa indices of 9 bits
			// n should be 512 now
			extraBits := byte(0)
			i := 0
			j := 0
			for j < int(sampler.SHA_512_DIGEST_LENGTH) {
				if j&7 == 0 {
					extraBits = whash[j]
					j++
				}
				index := (uint32(whash[j]) << 1) | uint32(extraBits&1)
				extraBits >>= 1
				j++
				if !array[index] {
					indices[i] = index
					array[index] = true
					i++
					if i >= int(kappa) {
						return indices
					}
				}
			}
		}
	}
	return []uint32{}
}

func greedySc(indices []uint32, s1, s2 *poly.PolyArray) (v1, v2 *poly.PolyArray) {
	n := s1.Param().N
	v1, _ = poly.NewPolyArray(s1.Param())
	v2, _ = poly.NewPolyArray(s2.Param())
	s1data := s1.GetData()
	s2data := s2.GetData()
	v1data := v1.GetData()
	v2data := v2.GetData()
	for k := 0; k < len(indices); k++ {
		index := indices[k]
		sign := int32(0)
		for i := uint32(0); i < n-index; i++ {
			sign += s1data[i]*v1data[index+i] + s2data[i]*v2data[index+i]
		}
		for i := n - index; i < n; i++ {
			sign -= s1data[i]*v1data[index+i-n] + s2data[i]*v2data[index+i-n]
		}
		if sign > 0 {
			for i := uint32(0); i < n-index; i++ {
				v1data[index+i] -= s1data[i]
				v2data[index+i] -= s2data[i]
			}
			for i := n - index; i < n; i++ {
				v1data[index+i-n] += s1data[i]
				v2data[index+i-n] += s2data[i]
			}
		} else {
			for i := uint32(0); i < n-index; i++ {
				v1data[index+i] += s1data[i]
				v2data[index+i] += s2data[i]
			}
			for i := n - index; i < n; i++ {
				v1data[index+i-n] -= s1data[i]
				v2data[index+i-n] -= s2data[i]
			}
		}
	}
	return
}

func (key *BlissPrivateKey) Sign(msg []byte, entropy *sampler.Entropy) (*BlissSignature, error) {
	kappa := key.Param().Kappa
	version := key.Param().Version
	Binf := key.Param().Binf
	Bl2 := key.Param().Bl2
	M := key.Param().M
	sampler, err := sampler.New(version, entropy)
	if err != nil {
		return nil, err
	}
	hash := sha3.Sum512(msg)
restart:
	y1 := poly.GaussPoly(version, sampler)
	y2 := poly.GaussPoly(version, sampler)
	v, err := y1.MultiplyNTT(key.a)
	if err != nil {
		return nil, err
	}
	v.ScalarMul(2)
	v.ScalarMul(int32(key.Param().OneQ2))
	v.Inc(y2)
	v = v.Mod2Q()
	dv := v.DropBits().ModP()
	indices := computeC(kappa, dv, hash[:])
	v1, v2 := greedySc(indices, key.s1, key.s2)
	normV := v1.Norm2() + v2.Norm2()
	if M <= uint32(normV) {
		return nil, fmt.Errorf("|v|^2 is larger than M")
	}
	if !sampler.SampleBerExp(M - uint32(normV)) {
		goto restart
	}
	var z1, z2 *poly.PolyArray
	b := entropy.Bit()
	if b {
		z1 = y1.Sub(v1)
		z2 = y2.Sub(v2)
	} else {
		z1 = y1.Add(v1)
		z2 = y2.Add(v2)
	}
	prodZV := z1.InnerProduct(v1) + z2.InnerProduct(v2)
	if !sampler.SampleBerCosh(prodZV) {
		goto restart
	}
	y1 = v.Sub(z2).Mod2Q().DropBits()
	v = v.DropBits()
	z2 = v.Sub(y1).BoundByP()
	if z1.MaxNorm() > int32(Binf) {
		goto restart
	}
	y2 = z2.Mul2d()
	if y2.MaxNorm() > int32(Binf) {
		goto restart
	}
	if z1.Norm2()+y2.Norm2() > int32(Bl2) {
		goto restart
	}
	return &BlissSignature{z1, z2, indices}, nil
}

func (key *BlissPublicKey) Verify(msg []byte, sig *BlissSignature) (bool, error) {
	if key.a.Param().Version != sig.z1.Param().Version {
		return false, fmt.Errorf("Mismatched signature version")
	}
	z1, z2, indices := sig.z1, sig.z2, sig.c
	param := key.a.Param()
	if z1.MaxNorm() > int32(param.Binf) {
		return false, fmt.Errorf("z1 max norm too large")
	}
	tz2 := z2.Mul2d()
	if tz2.MaxNorm() > int32(param.Binf) {
		return false, fmt.Errorf("z2 max norm too large")
	}
	if z1.Norm2()+tz2.Norm2() > int32(param.Bl2) {
		return false, fmt.Errorf("t1,z2 L2 norm too large")
	}
	hash := sha3.Sum512(msg)
	v, err := z1.MultiplyNTT(key.a)
	if err != nil {
		return false, err
	}
	v.ScalarMul(2)
	v.ScalarMul(int32(key.Param().OneQ2))
	v = v.Mod2Q()
	vdata := v.GetData()
	for i := 0; i < len(indices); i++ {
		qq := param.Q * param.OneQ2
		vdata[indices[i]] = v.NumMod2Q(vdata[indices[i]] + int32(qq))
	}
	v = v.DropBits().Add(z2).ModP()
	indicesp := computeC(param.Kappa, v, hash[:])
	for i := 0; i < len(indices); i++ {
		if indices[i] != indicesp[i] {
			return false, fmt.Errorf("Indices mismatch!")
		}
	}
	return true, nil
}
