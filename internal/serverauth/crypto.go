package serverauth

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

const (
	rsaKeyType   string = "RSA PRIVATE KEY"
	ecdsaKeyType string = "EC PRIVATE KEY"
	certType     string = "CERTIFICATE"
)

func pubKey[T crypto.PrivateKey](privKey T) (pk crypto.PublicKey) {
	switch p := (any)(privKey).(type) {
	case *rsa.PrivateKey:
		pk = &p.PublicKey
	case *ecdsa.PrivateKey:
		pk = &p.PublicKey
	}

	return
}

func pemBlock[T crypto.PrivateKey](privKey T) (block *pem.Block) {
	block = new(pem.Block)

	switch key := (any)(privKey).(type) {
	case *rsa.PrivateKey:
		block.Type = rsaKeyType
		block.Bytes = x509.MarshalPKCS1PrivateKey(key)
	case *ecdsa.PrivateKey:
		var err error

		block.Type = ecdsaKeyType
		block.Bytes, err = x509.MarshalECPrivateKey(key)
		if err != nil {
			panic(errors.WithMessage(err, "error marshalling ecdsa private key"))
		}
	}

	return block
}
