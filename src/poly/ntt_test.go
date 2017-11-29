package poly

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestFFT(t *testing.T) {
	for i := 0; i <= 4; i++ {
		for k := 0; k < 2; k++ {
			testfile, err := ioutil.ReadFile(fmt.Sprintf("test_data/fft_test_%d%d", k, i))
			if err != nil {
				t.Errorf("Failed to open file: %s", err.Error())
			}
			filecontent := strings.TrimSpace(string(testfile))
			vs := strings.Split(filecontent, "\n")
			if len(vs) != 2 {
				t.Errorf("Error in data read from test_data: len(vs) = %d", len(vs))
			}
			v1 := strings.Split(strings.TrimSpace(vs[0]), " ")
			v2 := strings.Split(strings.TrimSpace(vs[1]), " ")
			poly, err := New(i)
			if err != nil {
				t.Errorf("Failed to create polynomial: %s", err.Error())
			}
			if int(poly.n) != len(v1) || int(poly.n) != len(v2) {
				t.Errorf("Data size invalid: n = %d, but len(v1) = %d, len(v2) = %d",
					len(v1), len(v2))
			}
			for j := 0; j < int(poly.n); j++ {
				tmp, err := strconv.Atoi(v1[j])
				if err != nil {
					t.Errorf("Invalid integer: ", v1[j])
				}
				poly.data[j] = int32(tmp)
			}
			array, err := poly.FFT()
			if err != nil {
				t.Errorf("Error in FFT(): %s", err.Error())
			}
			for j := 0; j < int(poly.n); j++ {
				tmp, err := strconv.Atoi(v2[j])
				if err != nil {
					t.Errorf("Invalid integer: ", v2[j])
				}
				if tmp != int(array.data[j]) {
					t.Errorf("Wrong result: expect %d, got %d", tmp, array.data[j])
				}
			}
		}
	}
}

func TestInvertAsNTT(t *testing.T) {
	for i := 0; i <= 4; i++ {
		testfile, err := ioutil.ReadFile(fmt.Sprintf("test_data/ntt_test_%d", i))
		if err != nil {
			t.Errorf("Failed to open file: %s", err.Error())
		}
		filecontent := strings.TrimSpace(string(testfile))
		vs := strings.Split(filecontent, "\n")
		if len(vs) != 2 {
			t.Errorf("Error in data read from test_data: len(vs) = %d", len(vs))
		}
		v1 := strings.Split(strings.TrimSpace(vs[0]), " ")
		v2 := strings.Split(strings.TrimSpace(vs[1]), " ")
		poly, err := New(i)
		if err != nil {
			t.Errorf("Failed to create polynomial: %s", err.Error())
		}
		if int(poly.n) != len(v1) || int(poly.n) != len(v2) {
			t.Errorf("Data size invalid: n = %d, but len(v1) = %d, len(v2) = %d",
				len(v1), len(v2))
		}
		for j := 0; j < int(poly.n); j++ {
			tmp, err := strconv.Atoi(v1[j])
			if err != nil {
				t.Errorf("Invalid integer: ", v1[j])
			}
			poly.data[j] = int32(tmp)
		}
		ntt, err := poly.NTT()
		if err != nil {
			t.Errorf("Error in FFT(): %s", err.Error())
		}
		for j := 0; j < int(poly.n); j++ {
			tmp, err := strconv.Atoi(v2[j])
			if err != nil {
				t.Errorf("Invalid integer: %s", v2[j])
			}
			if tmp != int(ntt.data[j]) {
				t.Errorf("Wrong result of FFT(): expect %d, got %d", tmp, ntt.data[j])
			}
		}
		inv, err := ntt.InvertAsNTT()
		if err == nil {
			test := inv.TimesModQ(ntt)
			for j := 0; j < int(test.n); j++ {
				if test.data[j] != 1 {
					t.Errorf("Wrong result of Invert(): expect 1, got %d", test.data[j])
				}
			}
		} else {
			fmt.Printf("Test polynomial test_data/ntt_test_%d not invertible.\n", i)
		}
	}
}

func TestNTT(t *testing.T) {
	for i := 0; i <= 4; i++ {
		testfile, err := ioutil.ReadFile(fmt.Sprintf("test_data/ntt_test_%d", i))
		if err != nil {
			t.Errorf("Failed to open file: %s", err.Error())
		}
		filecontent := strings.TrimSpace(string(testfile))
		vs := strings.Split(filecontent, "\n")
		if len(vs) != 2 {
			t.Errorf("Error in data read from test_data: len(vs) = %d", len(vs))
		}
		v1 := strings.Split(strings.TrimSpace(vs[0]), " ")
		v2 := strings.Split(strings.TrimSpace(vs[1]), " ")
		poly, err := New(i)
		if err != nil {
			t.Errorf("Failed to create polynomial: %s", err.Error())
		}
		if int(poly.n) != len(v1) || int(poly.n) != len(v2) {
			t.Errorf("Data size invalid: n = %d, but len(v1) = %d, len(v2) = %d",
				len(v1), len(v2))
		}
		for j := 0; j < int(poly.n); j++ {
			tmp, err := strconv.Atoi(v1[j])
			if err != nil {
				t.Errorf("Invalid integer: ", v1[j])
			}
			poly.data[j] = int32(tmp)
		}
		ntt, err := poly.NTT()
		if err != nil {
			t.Errorf("Error in FFT(): %s", err.Error())
		}
		for j := 0; j < int(poly.n); j++ {
			tmp, err := strconv.Atoi(v2[j])
			if err != nil {
				t.Errorf("Invalid integer: ", v2[j])
			}
			if tmp != int(ntt.data[j]) {
				t.Errorf("Wrong result of FFT(): expect %d, got %d", tmp, ntt.data[j])
			}
		}
		npoly, err := ntt.INTT()
		if err != nil {
			t.Errorf("Error in Poly(): %s", err.Error())
		}
		for j := 0; j < int(poly.n); j++ {
			if npoly.data[j] != poly.NumModQ(poly.data[j]) {
				t.Errorf("Wrong result of Poly(): expect %d, got %d", poly.data[j], npoly.data[j])
			}
		}
	}
}
