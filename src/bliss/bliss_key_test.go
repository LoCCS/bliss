package bliss

import (
	"fmt"
	"io/ioutil"
	"sampler"
	"strconv"
	"strings"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	for i := 0; i <= 4; i++ {
		testfile, err := ioutil.ReadFile(fmt.Sprintf("test_data/key_test_%d", i))
		if err != nil {
			t.Errorf("Failed to open file: %s", err.Error())
		}
		filecontent := strings.TrimSpace(string(testfile))
		vs := strings.Split(filecontent, "\n")
		if len(vs) != 3 {
			t.Errorf("Error in data read from test_data: len(vs) = %d", len(vs))
		}
		v1 := strings.Split(strings.TrimSpace(vs[0]), " ")
		v2 := strings.Split(strings.TrimSpace(vs[1]), " ")
		v3 := strings.Split(strings.TrimSpace(vs[2]), " ")

		seed := make([]uint8, sampler.SHA_512_DIGEST_LENGTH)
		for i := 0; i < len(seed); i++ {
			seed[i] = uint8(i % 8)
		}
		entropy, err := sampler.NewEntropy(seed)
		if err != nil {
			t.Errorf("Error in initializing entropy: %s", err.Error())
		}

		key, err := GeneratePrivateKey(i, entropy)
		if err != nil {
			t.Errorf("Error in generating private key: %s", err.Error())
		}

		for j := 0; j < int(key.s1.Size()); j++ {
			tmp, err := strconv.Atoi(v1[j])
			s1 := key.s1.GetData()
			if err != nil {
				t.Errorf("Invalid integer: %s", v1[j])
			}
			if int32(tmp) != s1[j] {
				t.Errorf("Wrong s1 at %d: expect %d, got %d", j, tmp, s1[j])
			}
			tmp, err = strconv.Atoi(v2[j])
			s2 := key.s2.GetData()
			if err != nil {
				t.Errorf("Invalid integer: %s", v2[j])
			}
			if int32(tmp) != s2[j] {
				t.Errorf("Wrong s2 at %d: expect %d, got %d", j, tmp, s2[j])
			}
			tmp, err = strconv.Atoi(v3[j])
			s3 := key.a.GetData()
			if err != nil {
				t.Errorf("Invalid integer: %s", v3[j])
			}
			if int32(tmp) != s3[j] {
				t.Errorf("Wrong a at %d: expect %d, got %d", j, tmp, s3[j])
			}
		}
	}
}
