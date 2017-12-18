package poly

import (
	"testing"
)

func TestPolyArrayInc(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{2, 5, 8, 4, 7, 10, 8, 6, 6, 6}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	f.Inc(g)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.Inc(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.Inc(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayAdd(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{2, 5, 8, 4, 7, 10, 8, 6, 6, 6}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	h := f.Add(g)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.Add(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.Add(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayDec(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{-2, -3, -4, 2, 1, 0, 4, 4, 2, 0}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	f.Dec(g)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.Dec(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.Dec(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArraySub(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{-2, -3, -4, 2, 1, 0, 4, 4, 2, 0}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	h := f.Sub(g)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.Sub(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.Sub(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayMul(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{0, 4, 12, 3, 12, 25, 12, 5, 8, 9}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	f.Mul(g)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.Mul(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.Mul(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayScalarMul(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	hdata := []int32{0, 4, 8, 12, 16, 20, 24, 20, 16, 12}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	f.ScalarMul(4)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.ScalarMul(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.ScalarMul(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayScalarTimes(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	hdata := []int32{0, 4, 8, 12, 16, 20, 24, 20, 16, 12}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	h := f.ScalarTimes(4)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.ScalarTimes(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.ScalarTimes(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayTimes(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{0, 4, 12, 3, 12, 25, 12, 5, 8, 9}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	h := f.Times(g)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.Times(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.Times(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayNorm2(t *testing.T) {
	fdata := []int32{0, -1, 2, -3, -4, 5, -6, 5, 4, -3}
	norm := int32(141)
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	res := f.Norm2()
	if res != norm {
		t.Errorf("Error in computing f.Norm2(): expect %d, got %d", norm, res)
	}
}
