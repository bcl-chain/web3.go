// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package crypto implements a wrapper library for
// github.com/bcl-chain/go-ethereum/crypto.
//

package crypto

import (
	"crypto/ecdsa"

	"github.com/bcl-chain/web3.go/common"
	"github.com/ethereum/go-ethereum/crypto"
	wcommon "github.com/bcl-chain/web3.go/wrapper/common"
)

type PublicKey struct {
	PublicKey interface{}
}

type PrivateKey struct {
	PrivateKey interface{}
}

func Keccak256Hash(data []byte) *common.Hash {
	return wcommon.FromHash(crypto.Keccak256Hash(data))
}

func (priv *PrivateKey) Public() *PublicKey {
	privateKey, _ := priv.PrivateKey.(*ecdsa.PrivateKey)
	return &PublicKey{
		PublicKey: &privateKey.PublicKey,
	}
}

func GenerateKey() (*PrivateKey, error) {
	if privateKey, err := crypto.GenerateKey(); err == nil {
		return &PrivateKey{
			PrivateKey: privateKey,
		}, nil
	} else {
		return nil, err
	}
}

func FromECDSA(priv *PrivateKey) []byte {
	privateKey, _ := priv.PrivateKey.(*ecdsa.PrivateKey)
	return crypto.FromECDSA(privateKey)
}

func FromECDSAPub(pub *PublicKey) []byte {
	publicKey, _ := pub.PublicKey.(*ecdsa.PublicKey)
	return crypto.FromECDSAPub(publicKey)
}

func HexToECDSA(hexkey string) (*PrivateKey, error) {
	if privateKey, err := crypto.HexToECDSA(hexkey); err == nil {
		return &PrivateKey{
			PrivateKey: privateKey,
		}, nil
	} else {
		return nil, err
	}
}

func PubkeyToAddress(pub *PublicKey) *common.Address {
	pubBytes := FromECDSAPub(pub)
	return common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
}
