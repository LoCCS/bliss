package bliss

import (
	_ "fmt"
	_ "io/ioutil"
	"github.com/LoCCS/bliss/sampler"
	_ "strconv"
	_ "strings"
	"testing"
	"reflect"
)

func TestSignVerify(t *testing.T) {
	for i := 0; i <= 4; i++ {
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

		pub := key.PublicKey()
		msg := []byte("Hello world")
		sig, err := key.Sign(msg, entropy)
		if err != nil {
			t.Errorf("Failed to generate signature for version %d: %s", i, err.Error())
		}
		/*
			z1data := sig.z1.GetData()
			z2data := sig.z2.GetData()
			fmt.Printf("z1: ")
			for j := 0; j < len(z1data); j++ {
				fmt.Printf("%d ", z1data[j])
			}
			fmt.Printf("\n")
			fmt.Printf("z2: ")
			for j := 0; j < len(z2data); j++ {
				fmt.Printf("%d ", z2data[j])
			}
			fmt.Printf("\n")
			fmt.Printf("c: ")
			for j := 0; j < len(sig.c); j++ {
				fmt.Printf("%d ", sig.c[j])
			}
			fmt.Printf("\n")
		*/
		_, err = pub.Verify(msg, sig)
		if err != nil {
			t.Errorf("Failed to verify signature for version %d: %s", i, err.Error())
		}
	}
}


func TestSignatureEncodeDecode(t *testing.T) {
	for i := 0; i <= 4; i++ {
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

		msg := []byte("Hello world")
		sig, err := key.Sign(msg, entropy)
		if err != nil {
			t.Errorf("Failed to generate signature for version %d: %s", i, err.Error())
		}

		enc := sig.Encode()
		tmp, err := DecodeSignature(enc)
		if err != nil {
			t.Errorf("Error in decoding signature : %s", err.Error())
		}
		if !reflect.DeepEqual(sig, tmp) {
			t.Errorf("Different signature decoded for version %d!\nOriginal:\n%s\ngot:\n%s\n",
				i, sig.String(), tmp.String())
		}
	}
}