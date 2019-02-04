// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package crypto implements a wrapper library for
// github.com/bcl-chain/go-ethereum/crypto.
//

package api

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type PublicKey struct {
	publicKey *ecdsa.PublicKey
}

func fromPublicKey(pub *ecdsa.PublicKey) *PublicKey {
	if pub != nil {
		return &PublicKey{
			publicKey: pub,
		}
	} else {
		return nil
	}
}

func toPublicKey(wpub *PublicKey) *ecdsa.PublicKey {
	if wpub != nil {
		return wpub.publicKey
	} else {
		return nil
	}
}

type PrivateKey struct {
	privateKey *ecdsa.PrivateKey
}

func fromPrivateKey(prv *ecdsa.PrivateKey) *PrivateKey {
	if prv != nil {
		return &PrivateKey{
			privateKey: prv,
		}
	} else {
		return nil
	}
}

func toPrivateKey(wprv *PrivateKey) *ecdsa.PrivateKey {
	if wprv != nil {
		return wprv.privateKey
	} else {
		return nil
	}
}

func Keccak256Hash(data []byte) *Hash {
	h := crypto.Keccak256Hash(data)
	return fromHash(h)
}

func (priv *PrivateKey) Public() *PublicKey {
	privateKey := priv.privateKey
	return fromPublicKey(&privateKey.PublicKey)
}

func GenerateKey() (*PrivateKey, error) {
	if privateKey, err := crypto.GenerateKey(); err == nil {
		return fromPrivateKey(privateKey), nil
	} else {
		return nil, err
	}
}

func FromECDSA(priv *PrivateKey) []byte {
	privateKey := toPrivateKey(priv)
	return crypto.FromECDSA(privateKey)
}

func FromECDSAPub(pub *PublicKey) []byte {
	publicKey := toPublicKey(pub)
	return crypto.FromECDSAPub(publicKey)
}

func HexToECDSA(hexkey string) (*PrivateKey, error) {
	if privateKey, err := crypto.HexToECDSA(hexkey); err == nil {
		return fromPrivateKey(privateKey), nil
	} else {
		return nil, err
	}
}

func PubkeyToAddress(pub *PublicKey) *Address {
	pubBytes := FromECDSAPub(pub)
	return BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
}
