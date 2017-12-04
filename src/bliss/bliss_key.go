package bliss

import (
	"fmt"
	"params"
	"poly"
	"sampler"
)

type BlissPrivateKey struct {
	s1 *poly.PolyArray
	s2 *poly.PolyArray
	a  *poly.PolyArray
}

type BlissPublicKey struct {
	a *poly.PolyArray
}

func GeneratePrivateKey(version int, entropy *sampler.Entropy) (*BlissPrivateKey, error) {
	// Generate g
	s2 := poly.UniformPoly(version, entropy)
	if s2 == nil {
		return nil, fmt.Errorf("Failed to generate uniform polynomial g")
	}
	// s2 = 2g-1
	s2.ScalarMul(2)
	s2.ScalarInc(-1)

	t, err := s2.NTT()
	if err != nil {
		return nil, err
	}

	for j := 0; j < 4; j++ {
		s1 := poly.UniformPoly(version, entropy)
		if s1 == nil {
			return nil, fmt.Errorf("Failed to generate uniform polynomial f")
		}
		u, err := s1.NTT()
		if err != nil {
			return nil, err
		}
		u, err = u.InvertAsNTT()
		if err != nil {
			continue
		}
		t.MulModQ(u)
		t, err = t.INTT()
		if err != nil {
			return nil, err
		}
		t.ScalarMulModQ(-1)
		a, err := t.NTT()
		if err != nil {
			return nil, err
		}
		key := BlissPrivateKey{s1, s2, a}
		return &key, nil
	}
	return nil, fmt.Errorf("Failed to generate invertible polynomial")
}

func (privateKey *BlissPrivateKey) PublicKey() *BlissPublicKey {
	return &BlissPublicKey{privateKey.a}
}

func (privateKey *BlissPrivateKey) Param() *params.BlissBParam {
	return privateKey.s1.Param()
}

func (publicKey *BlissPublicKey) Param() *params.BlissBParam {
	return publicKey.a.Param()
}

func (privateKey *BlissPrivateKey) String() string {
	return fmt.Sprintf("{s1:%s,s2:%s,a:%s}",
		privateKey.s1.String(), privateKey.s2.String(), privateKey.a.String())
}

func (publicKey *BlissPublicKey) String() string {
	return fmt.Sprintf("{a:%s}", publicKey.a.String())
}

func (publicKey *BlissPublicKey) Encode() []byte {
	n := publicKey.Param().N
	data := publicKey.a.GetData()
	ret := make([]byte, n*2+1)
	ret[0] = byte(publicKey.Param().Version)
	for i := 0; i < int(n); i++ {
		ret[i*2+1] = byte(uint16(data[i]) >> 8)
		ret[i*2+2] = byte(uint16(data[i]) & 0xff)
	}
	return ret[:]
}

func DecodeBlissPublicKey(data []byte) (*BlissPublicKey, error) {
	a, err := poly.New(int(data[0]))
	if err != nil {
		return nil, fmt.Errorf("Error in generating new polyarray: %s", err.Error())
	}
	ret := &BlissPublicKey{a}
	n := a.Param().N
	retdata := make([]int32, n)
	for i := 0; i < int(n); i++ {
		retdata[i] = (int32(data[i*2+1]) << 8) | (int32(data[i*2+2]))
	}
	a.SetData(retdata)
	return ret, nil
}
