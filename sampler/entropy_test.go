package sampler

import (
	"fmt"
	"testing"
	"strings"
	"io/ioutil"
)

func TestEntropy(t *testing.T) {
	testfile,err := ioutil.ReadFile("test_data/entropy_test")
	if err != nil {
		t.Errorf("Failed to open file: %s",err.Error())
	}
	filecontent := strings.TrimSpace(string(testfile))
	vs := strings.Split(filecontent,"\n")

	seed := make([]uint8,SHA_512_DIGEST_LENGTH)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(i % 8)
	}
	entropy,err := NewEntropy(seed)
	if err != nil {
		t.Errorf("Error in initializing entropy: %s",err.Error())
	}

	for i := 0; i < 128; i++ {
		ns := strings.Split(strings.TrimSpace(vs[i])," ")
		bit := entropy.Bit()
		vchar := entropy.Char()
		vint16 := entropy.Uint16()
		vint64 := entropy.Uint64()
		bits := entropy.Bits(i%32)
		if (bit && ns[0] == "0") || (!bit && ns[0] == "1") {
			t.Errorf("Error in bit in line %d: expect %s",i,ns[0])
		}
		if ns[1] != fmt.Sprintf("%d",vchar) {
			t.Errorf("Error in char in line %d: expect %s, got %d",i,ns[1],vchar)
		}
		if ns[2] != fmt.Sprintf("%d",vint16) {
			t.Errorf("Error in int16 in line %d: expect %s, got %d",i,ns[2],vint16)
		}
		if ns[3] != fmt.Sprintf("%d",vint64) {
			t.Errorf("Error in int64 in line %d: expect %s, got %d",i,ns[3],vint64)
		}
		if ns[4] != fmt.Sprintf("%d",bits) {
			t.Errorf("Error in bits in line %d: expect %s, got %d",i,ns[4],bits)
		}
	}
}
