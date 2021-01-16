package main

import (
	"fmt"
	"crypto"
	"crypto/ed25519"

	"github.com/stellar/go/keypair"
)

func main() {
	pubKey, privKey, err := ed25519.GenerateKey(nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Public Key: %+v (%d)\n", pubKey, len(pubKey))
	fmt.Printf("Public Key: %+v\n", privKey.Public())
	fmt.Printf("Private Key: %+v (%d)\n", privKey, len(privKey))
	fmt.Printf("Private Key Seed: %+v (%d)\n", privKey.Seed(), len(privKey.Seed()))

	message := []byte("The quick brown fox jumps over the lazy dog.")
	signature, err := privKey.Sign(nil, message, crypto.Hash(0))

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Signature: %+v\n", signature)

	fmt.Printf("Verified: %t\n", ed25519.Verify(pubKey, message, signature))

	var seed [32]byte
	copy(seed[:], privKey.Seed()[:])
	keyPair, err := keypair.FromRawSeed(seed)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", keyPair)

	keyPair2, err := keypair.Parse(keyPair.Seed())

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", keyPair2)

	fmt.Println(keyPair.Address())
	fmt.Println(keyPair2.Address())
}
