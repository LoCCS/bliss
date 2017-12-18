package sampler

import (
	"fmt"
	"github.com/LoCCS/bliss/params"
)


type Sampler struct {
	sigma      uint32
	ell        uint32
	prec       uint32
	columns    uint32
	kSigma     uint16
	kSigmaBits uint16

	ctable   []uint8
	cdttable []uint8

	random *Entropy
}

func invalidSampler() *Sampler {
	return &Sampler{0,0,0,0,0,0,[]uint8{},[]uint8{},nil}
}

func NewSampler(sigma,ell,prec uint32, entropy *Entropy) (*Sampler, error) {
	columns := prec/8
	ctable,err := getTable(sigma,ell,prec)
	if err != nil {
		return invalidSampler(),err
	}
	ksigma := getKSigma(sigma,prec)
	if ksigma == 0 {
		return invalidSampler(),fmt.Errorf("Failed to get kSigma")
	}
	ksigmabits := getKSigmaBits(sigma,prec)
	if ksigmabits == 0 {
		return invalidSampler(),fmt.Errorf("Failed to get kSigmaBits")
	}
	return &Sampler{sigma,ell,prec,columns,ksigma,ksigmabits,ctable,[]uint8{},entropy},nil
}

func New(version int, entropy *Entropy) (*Sampler, error) {
	param := params.GetParam(version)
	if param == nil {
		return nil,fmt.Errorf("Failed to get parameter")
	}
	return NewSampler(param.Sigma,param.Ell,param.Prec,entropy)
}

// Sample Bernoulli distribution with probability p
// p is stored as a large big-endian integer in an array
// the real probability is p/2^d, where d is the number of
// bits of p
func (sampler *Sampler) sampleBer(p []uint8) bool {
	for _,pi := range p {
		uc := sampler.random.Char()
		if uc < pi {
			return true
		}
		if uc > pi {
			return false
		}
	}
	return true
}

// Sample Bernoulli distribution with probability p = exp(-x/(2*sigma^2))
func (sampler *Sampler) SampleBerExp(x uint32) bool {
	ri := sampler.ell - 1
	mask := uint32(1) << ri
	start := ri * sampler.columns
	for mask > 0 {
		if x & mask != 0 {
			if !sampler.sampleBer(sampler.ctable[start:start+sampler.columns]) {
				return false
			}
		}
		mask >>= 1
		start -= sampler.columns
	}
	return true
}

// Sample Bernoulli distribution with probability p = 1/cosh(-x/(2*sigma^2))
func (sampler *Sampler) SampleBerCosh(x int32) bool {
	if x < 0 {
		x = -x
	}
	x <<= 1
	for {
		bit := sampler.SampleBerExp(uint32(x))
		if bit {
			return true
		}
		bit = sampler.random.Bit()
		if !bit {
			bit = sampler.SampleBerExp(uint32(x))
			if !bit {
				return false
			}
		}
	}
}

// Discrete Binary Gauss distribution is Discrete Gauss Distribution with
// a specific variance sigma = sqrt(1/(2 ln 2)) = 0.849...
// This is used as foundation of SampleGauss.
func (sampler *Sampler) SampleBinaryGauss() uint32 {
restart:
	if sampler.random.Bit() {
		return 0
	}
	for i := 1; i <= 16; i++ {
		u := sampler.random.Bits(2*i - 1)
		if u == 0 {
			return uint32(i)
		}
		if u != 1 {
			goto restart
		}
	}
	return 0
}

// Sample according to Discrete Gauss Distribution
// exp(-x^2/(2*sigma*sigma))
func (sampler *Sampler) SampleGauss() int32 {
	var x,y uint32
	var u bool
	for {
		x = sampler.SampleBinaryGauss()
		for {
			y = sampler.random.Bits(int(sampler.kSigmaBits))
			if y < uint32(sampler.kSigma) {
				break
			}
		}
		e := y * (y + 2 * uint32(sampler.kSigma) * x)
		u = sampler.random.Bit()
		if (x | y) != 0 || u {
			if sampler.SampleBerExp(e) {
				break
			}
		}
	}

	valPos := int32(uint32(sampler.kSigma) * x + y)
	if u {
		return valPos
	} else {
		return -valPos
	}
}
