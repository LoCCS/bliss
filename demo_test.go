package bliss

import (
	"github.com/LoCCS/bliss/params"
	"github.com/LoCCS/bliss/sampler"
	"testing"
	"fmt"
)

func TestDemo(t *testing.T) {
	version := params.BLISS_B_0
	seed := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
	}
	msg := "Hello world"

	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		t.Errorf("Error in creating entropy: %s\n", err.Error())
		return
	}

	key, err := GeneratePrivateKey(version, entropy)
	if err != nil {
		t.Errorf("Error in generating private key: %s\n", err.Error())
		return
	} else {
		fmt.Errorf("Private Key: %s\n", key.String())
	}

	pub := key.PublicKey()
	fmt.Printf("Public Key: %s\n", pub.String())

	sig, err := key.Sign([]byte(msg), entropy)
	if err != nil {
		t.Errorf("Error in signing: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Signature: %s\n", sig.String())
	}

	res, err := pub.Verify([]byte(msg), sig)
	if res {
		fmt.Printf("Verified!\n")
	} else {
		t.Errorf("Error: %s\n", err.Error())
	}
}
