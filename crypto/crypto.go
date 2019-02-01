// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package crypto implements a wrapper library for
// github.com/bcl-chain/go-ethereum/crypto.
//

package crypto

import (
	"crypto/ecdsa"

	"github.com/bcl-chain/web3.go/common"
	obj "github.com/bcl-chain/web3.go/object"
	"github.com/ethereum/go-ethereum/crypto"
)

type PublicKey struct{}

type PrivateKey struct{}

func (priv *PrivateKey) Public() *PublicKey {
	privateKey, _ := obj.GetObject(priv).(*ecdsa.PrivateKey)
	return obj.NewObject(&PublicKey{}, &privateKey.PublicKey).(*PublicKey)
}

func GenerateKey() (*PrivateKey, error) {
	if privateKey, err := crypto.GenerateKey(); err == nil {
		return obj.NewObject(&PrivateKey{}, privateKey).(*PrivateKey), nil
	} else {
		return nil, err
	}
}

func FromECDSA(priv *PrivateKey) []byte {
	privateKey, _ := obj.GetObject(priv).(*ecdsa.PrivateKey)
	return crypto.FromECDSA(privateKey)
}

func FromECDSAPub(pub *PublicKey) []byte {
	publicKey, _ := obj.GetObject(pub).(*ecdsa.PublicKey)
	return crypto.FromECDSAPub(publicKey)
}

func HexToECDSA(hexkey string) (*PrivateKey, error) {
	if privateKey, err := crypto.HexToECDSA(hexkey); err == nil {
		return obj.NewObject(&PrivateKey{}, privateKey).(*PrivateKey), nil
	} else {
		return nil, err
	}
}

func PubkeyToAddress(pub *PublicKey) *common.Address {
	pubBytes := FromECDSAPub(pub)
	return common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
}
