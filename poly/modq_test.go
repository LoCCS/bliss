package poly

import (
	"testing"
)

func TestPolyArrayModQ(t *testing.T) {
	fdata := []int32{0, -1, -2, 3, -4, 5, -6, 5, -4, 3}
	hdata := []int32{0, 6, 5, 3, 3, 5, 1, 5, 3, 3}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	f.ModQ()
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

func TestPolyArrayIncModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{2, 5, 1, 4, 0, 3, 1, 6, 6, 6}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	f.IncModQ(g)
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

func TestPolyArrayAddModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{2, 5, 1, 4, 0, 3, 1, 6, 6, 6}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	h := f.AddModQ(g)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.AddModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.AddModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayDecModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{5, 4, 3, 2, 1, 0, 4, 4, 2, 0}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	f.DecModQ(g)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.DecModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.DecModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArraySubModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{5, 4, 3, 2, 1, 0, 4, 4, 2, 0}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	h := f.SubModQ(g)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.SubModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.SubModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayMulModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{0, 4, 5, 3, 5, 4, 5, 5, 1, 2}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	f.MulModQ(g)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.MulModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.MulModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayScalarMulModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	hdata := []int32{0, 4, 1, 5, 2, 6, 3, 6, 2, 5}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	f.ScalarMulModQ(4)
	res := f.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.ScalarMulModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.ScalarMulModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayScalarTimesModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	hdata := []int32{0, 4, 1, 5, 2, 6, 3, 6, 2, 5}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	h := f.ScalarTimesModQ(4)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.ScalarTimesModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.ScalarTimesModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayTimesModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	gdata := []int32{2, 4, 6, 1, 3, 5, 2, 1, 2, 3}
	hdata := []int32{0, 4, 5, 3, 5, 4, 5, 5, 1, 2}
	f, _ := newPolyArray(10, 7)
	g, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	g.SetData(gdata)
	h := f.TimesModQ(g)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.TimesModQ(g): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.TimesModQ(g): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}

func TestPolyArrayExpModQ(t *testing.T) {
	fdata := []int32{0, 1, 2, 3, 4, 5, 6, 5, 4, 3}
	hdata := []int32{0, 1, 4, 5, 2, 3, 6, 3, 2, 5}
	f, _ := newPolyArray(10, 7)
	f.SetData(fdata)
	h := f.ExpModQ(5)
	res := h.GetData()
	if len(hdata) != len(res) {
		t.Errorf("Error in computing f.ExpModQ(5): incorrect length %d", len(res))
	}
	for i := 0; i < len(hdata); i++ {
		if hdata[i] != res[i] {
			t.Errorf("Error in computing f.ExpModQ(5): %d != %d at pos %d",
				hdata[i], res[i], i)
		}
	}
}
