package bliss

import (
	"fmt"
	"github.com/LoCCS/bliss/params"
	"github.com/LoCCS/bliss/sampler"
	_ "io/ioutil"
	"reflect"
	_ "strconv"
	_ "strings"
	"testing"
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
		_, err = pub.Verify(msg, sig)
		if err != nil {
			t.Errorf("Failed to verify signature for version %d: %s", i, err.Error())
		}
	}
}

func TestSignVerifyAgainstChannel(t *testing.T) {
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
		sig, err := key.SignAgainstSideChannel(msg, entropy)
		if err != nil {
			t.Errorf("Failed to generate signature for version %d: %s", i, err.Error())
		}
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

func benchSign(b *testing.B, version int) {
	seed := make([]uint8, sampler.SHA_512_DIGEST_LENGTH)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(i % 8)
	}
	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		b.Errorf("Error in initializing entropy: %s", err.Error())
	}

	key, err := GeneratePrivateKey(version, entropy)
	if err != nil {
		b.Errorf("Error in generating private key: %s", err.Error())
	}

	msg := []byte("Hello world")
	for i := 0; i < b.N; i++ {
		key.Sign(msg, entropy)
	}
}

func benchSignAgainstSideChannel(b *testing.B, version int) {
	seed := make([]uint8, sampler.SHA_512_DIGEST_LENGTH)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(i % 8)
	}
	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		b.Errorf("Error in initializing entropy: %s", err.Error())
	}

	key, err := GeneratePrivateKey(version, entropy)
	if err != nil {
		b.Errorf("Error in generating private key: %s", err.Error())
	}

	msg := []byte("Hello world")
	for i := 0; i < b.N; i++ {
		key.SignAgainstSideChannel(msg, entropy)
	}
}

func BenchmarkSignBliss0(b *testing.B) {
	benchSign(b, params.BLISS_B_0)
}

func BenchmarkSignBliss1(b *testing.B) {
	benchSign(b, params.BLISS_B_1)
}

func BenchmarkSignBliss2(b *testing.B) {
	benchSign(b, params.BLISS_B_2)
}

func BenchmarkSignBliss3(b *testing.B) {
	benchSign(b, params.BLISS_B_3)
}

func BenchmarkSignBliss4(b *testing.B) {
	benchSign(b, params.BLISS_B_4)
}

func BenchmarkSignBliss0AgainstSideChannel(b *testing.B) {
	benchSignAgainstSideChannel(b, params.BLISS_B_0)
}

func BenchmarkSignBliss1AgainstSideChannel(b *testing.B) {
	benchSignAgainstSideChannel(b, params.BLISS_B_1)
}

func BenchmarkSignBliss2AgainstSideChannel(b *testing.B) {
	benchSignAgainstSideChannel(b, params.BLISS_B_2)
}

func BenchmarkSignBliss3AgainstSideChannel(b *testing.B) {
	benchSignAgainstSideChannel(b, params.BLISS_B_3)
}

func BenchmarkSignBliss4AgainstSideChannel(b *testing.B) {
	benchSignAgainstSideChannel(b, params.BLISS_B_4)
}

func benchVerify(b *testing.B, version int) {
	seed := make([]uint8, sampler.SHA_512_DIGEST_LENGTH)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(i % 8)
	}
	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		b.Errorf("Error in initializing entropy: %s", err.Error())
	}

	key, err := GeneratePrivateKey(version, entropy)
	if err != nil {
		b.Errorf("Error in generating private key: %s", err.Error())
	}

	pub := key.PublicKey()
	msg := []byte("Hello world")
	sig, err := key.Sign(msg, entropy)
	for i := 0; i < b.N; i++ {
		pub.Verify(msg, sig)
	}
}

func BenchmarkVerifyBliss0(b *testing.B) {
	benchVerify(b, params.BLISS_B_0)
}

func BenchmarkVerifyBliss1(b *testing.B) {
	benchVerify(b, params.BLISS_B_1)
}

func BenchmarkVerifyBliss2(b *testing.B) {
	benchVerify(b, params.BLISS_B_2)
}

func BenchmarkVerifyBliss3(b *testing.B) {
	benchVerify(b, params.BLISS_B_3)
}

func BenchmarkVerifyBliss4(b *testing.B) {
	benchVerify(b, params.BLISS_B_4)
}

func TestSignatureSerializeDeserialize(t *testing.T) {
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

		enc := sig.Serialize()
		fmt.Printf("Size of signature for BLISS-%d: %d bytes (%d bits)\n", i, len(enc), len(enc)*8)
		if len(enc) == 0 {
			t.Errorf("Failed to encode signature for version %d", i)
			continue
		}
		tmp, err := DeserializeBlissSignature(enc)
		if err != nil {
			t.Errorf("Error in decoding signature : %s", err.Error())
		}
		if !reflect.DeepEqual(sig, tmp) {
			t.Errorf("Different signature decoded for version %d!\nOriginal:\n%s\ngot:\n%s\n",
				i, sig.String(), tmp.String())
		}
	}
}
