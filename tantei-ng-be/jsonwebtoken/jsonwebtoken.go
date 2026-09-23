package jsonwebtoken

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

var PriKey *ecdsa.PrivateKey

func GenerateKeyPairs() {
	var err error

	PriKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// priDER, err := x509.MarshalECPrivateKey(PriKey)

	if err != nil {
		panic(err)
	}

	pubDER, err := x509.MarshalPKIXPublicKey(&PriKey.PublicKey)

	if err != nil {
		panic(err)
	}

	// priPEM := pem.EncodeToMemory(&pem.Block{
	// 	Type:  "PRIVATE KEY",
	// 	Bytes: priDER,
	// })

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	})

	fmt.Printf("%s\n", pubPEM)
}
