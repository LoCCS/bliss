package main

import (
	"bliss"
	"fmt"
	"params"
	"sampler"
)

func main() {
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
		fmt.Printf("Error in creating entropy: %s\n", err.Error())
		return
	}

	key, err := bliss.GeneratePrivateKey(version, entropy)
	if err != nil {
		fmt.Printf("Error in generating private key: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Private Key: %s\n", key.String())
	}

	pub := key.PublicKey()
	fmt.Printf("Public Key: %s\n", pub.String())

	sig, err := key.Sign([]byte(msg), entropy)
	if err != nil {
		fmt.Printf("Error in signing: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Signature: %s\n", sig.String())
	}

	res, err := pub.Verify([]byte(msg), sig)
	if res {
		fmt.Printf("Verified!\n")
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
