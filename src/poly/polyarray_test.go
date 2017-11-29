package poly

import (
	"params"
	"sampler"
	"testing"
)

func TestPolyArrayFlip(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	hdata := []int32{0, 3, 4, 5, 6, 5, 4, 3, 2, 1}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	f.flip()
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.flip(): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.flip(): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestNewPolyArray(t *testing.T) {
	var polyarray *PolyArray
	var err error
	var n, q uint32
	n = 10
	q = 10
	polyarray, err = newPolyArray(n, q)
	if polyarray == nil || err != nil || polyarray.n != n || polyarray.q != q {
		t.Errorf("Failed to create modular polyarray for n = %d, q = %d\n", n, q)
	}

	n = 0
	q = 10
	polyarray, err = newPolyArray(n, q)
	if polyarray != nil || err == nil {
		t.Errorf("Created modular polyarray for n = %d, q = %d\n", n, q)
	}

	n = 10
	q = 0
	polyarray, err = newPolyArray(n, q)
	if polyarray != nil || err == nil {
		t.Errorf("Created modular polyarray for n = %d, q = %d\n", n, q)
	}

	n = 0
	q = 0
	polyarray, err = newPolyArray(n, q)
	if polyarray != nil || err == nil {
		t.Errorf("Created modular polyarray for n = %d, q = %d\n", n, q)
	}
}

func TestNew(t *testing.T) {
	var polyarray *PolyArray
	var err error
	polyarray, err = New(params.BLISS_B_0)
	if polyarray == nil || err != nil {
		t.Errorf("Failed to create modular polyarray for BLISS_B_0: %s\n", err.Error())
	}
	polyarray, err = New(params.BLISS_B_1)
	if polyarray == nil || err != nil {
		t.Errorf("Failed to create modular polyarray for BLISS_B_1: %s\n", err.Error())
	}
	polyarray, err = New(params.BLISS_B_2)
	if polyarray == nil || err != nil {
		t.Errorf("Failed to create modular polyarray for BLISS_B_2: %s\n", err.Error())
	}
	polyarray, err = New(params.BLISS_B_3)
	if polyarray == nil || err != nil {
		t.Errorf("Failed to create modular polyarray for BLISS_B_3: %s\n", err.Error())
	}
	polyarray, err = New(params.BLISS_B_4)
	if polyarray == nil || err != nil {
		t.Errorf("Failed to create modular polyarray for BLISS_B_4: %s\n", err.Error())
	}
}

func TestUniformPoly(t *testing.T) {
	seed := make([]uint8, sampler.SHA_512_DIGEST_LENGTH)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(i % 8)
	}
	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		t.Errorf("Failed to create entropy")
	}
	p := UniformPoly(params.BLISS_B_4, entropy)
	if p == nil {
		t.Errorf("Failed to generate uniforma polynomial")
	}
	count0 := 0
	count1 := 0
	count2 := 0
	for i := 0; i < int(p.n); i++ {
		if p.data[i] == 0 {
			count0++
		} else if p.data[i] == 1 || p.data[i] == -1 {
			count1++
		} else if p.data[i] == 2 || p.data[i] == -2 {
			count2++
		}
	}
	if count1 != int(p.param.Nz1) {
		t.Errorf("Number of +-1 does not match: expect %d, got %d",
			p.param.Nz1, count1)
	}
	if count2 != int(p.param.Nz2) {
		t.Errorf("Number of +-2 does not match: expect %d, got %d",
			p.param.Nz2, count2)
	}
	if count0 != int(p.n-p.param.Nz1-p.param.Nz2) {
		t.Errorf("Number of 0 does not match: expect %d, got %d",
			p.n-p.param.Nz1-p.param.Nz2, count0)
	}
}
