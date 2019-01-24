// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package crypto implements a wrapper library for
// github.com/bcl-chain/go-ethereum/crypto.
//

package crypto

import (
	"crypto/ecdsa"

	"github.com/astra-x/go-ethereum/crypto"
)

type PublicKey struct {
	publicKey interface{}
}

type PrivateKey struct {
	privateKey interface{}
}

func (priv *PrivateKey) Public() *PublicKey {
	privateKey, _ := priv.privateKey.(*ecdsa.PrivateKey)
	return &PublicKey{
		publicKey: &privateKey.PublicKey,
	}
}

func GenerateKey() (*PrivateKey, error) {
	if privateKey, err := crypto.GenerateKey(); err == nil {
		return &PrivateKey{
			privateKey: privateKey,
		}, nil
	} else {
		return nil, err
	}
}

func FromECDSA(priv *PrivateKey) []byte {
	privateKey, _ := priv.privateKey.(*ecdsa.PrivateKey)
	return crypto.FromECDSA(privateKey)
}

func FromECDSAPub(pub *PublicKey) []byte {
	publicKey, _ := pub.publicKey.(*ecdsa.PublicKey)
	return crypto.FromECDSAPub(publicKey)
}

func PubkeyToAddress(pub *PublicKey) []byte {
	pubBytes := FromECDSAPub(pub)
	return crypto.Keccak256(pubBytes[1:])[12:]
}
