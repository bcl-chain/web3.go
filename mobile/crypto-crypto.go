// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package crypto implements a wrapper library for
// github.com/bcl-chain/go-ethereum/crypto.
//

package web3go

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
// secp256k1N = NewBigInt(0).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
)

// PublicKey ...
type PublicKey struct {
	publicKey *ecdsa.PublicKey
}

// Public ...
func (priv *PrivateKey) Public() *PublicKey {
	privateKey := priv.privateKey
	return &PublicKey{&privateKey.PublicKey}
}

// PrivateKey ...
type PrivateKey struct {
	privateKey *ecdsa.PrivateKey
}

// Keccak256 ...
// TODO: input type of original Keccak256 is ...[]byte.
//       K changed to []byte.
//       have to check if this modification is fine or not.
func Keccak256(data []byte) []byte {
	return crypto.Keccak256(data)
}

// Keccak256Hash ...
// TODO: input type of original Keccak256Hash is ...[]byte.
//       K changed to []byte.
//       have to check if this modification is fine or not.
func Keccak256Hash(data []byte) *Hash {
	h := crypto.Keccak256Hash(data)
	return &Hash{hash: h}
}

// Keccak512 ...
// TODO: input type of original Keccak512 is ...[]byte.
//       K changed to []byte.
//       have to check if this modification is fine or not.
func Keccak512(data []byte) []byte {
	return crypto.Keccak512(data)
}

// CreateAddress ...
func CreateAddress(wb *Address, wnonce int64) *Address {
	b := wb.address
	nonce := uint64(wnonce)
	a := crypto.CreateAddress(b, nonce)
	return &Address{address: a}
}

// TODO
//func CreateAddress2 {
//}

// ToECDSA ...
func ToECDSA(d []byte) (*PrivateKey, error) {
	priv, err := crypto.ToECDSA(d)
	if err == nil {
		return &PrivateKey{priv}, nil
	}
	return nil, err
}

// ToECDSAUnsafe ...
func ToECDSAUnsafe(d []byte) *PrivateKey {
	priv := crypto.ToECDSAUnsafe(d)
	return &PrivateKey{priv}
}

// FromECDSA ...
func FromECDSA(priv *PrivateKey) []byte {
	privateKey := priv.privateKey
	return crypto.FromECDSA(privateKey)
}

// UnmarshalPubkey ...
func UnmarshalPubkey(pub []byte) (*PublicKey, error) {
	publicKey, err := crypto.UnmarshalPubkey(pub)
	if err == nil {
		return &PublicKey{publicKey}, nil
	}
	return nil, err

}

// FromECDSAPub ...
func FromECDSAPub(pub *PublicKey) []byte {
	publicKey := pub.publicKey
	return crypto.FromECDSAPub(publicKey)
}

// HexToECDSA ...
func HexToECDSA(hexkey string) (*PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(hexkey)
	if err == nil {
		return &PrivateKey{privateKey}, nil
	}
	return nil, err

}

// LoadECDSA ...
func LoadECDSA(file string) (*PrivateKey, error) {
	privateKey, err := crypto.LoadECDSA(file)
	if err == nil {
		return &PrivateKey{privateKey}, nil
	}
	return nil, err
}

// SaveECDSA ...
func SaveECDSA(file string, wpriv *PrivateKey) error {
	priv := wpriv.privateKey
	return crypto.SaveECDSA(file, priv)
}

// GenerateKey ...
func GenerateKey() (*PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err == nil {
		return &PrivateKey{privateKey}, nil
	}
	return nil, err
}

// TODO; this function does not work in ios binding, byte is not suppotred!!
//func ValidateSignatureValues(v byte, wr, ws *BigInt, homestead bool) bool {
//	r := wr.bigint
//	s := ws.bigint
//	return crypto.ValidateSignatureValues(v, r, s, homestead)
//}

// PubkeyToAddress ...
func PubkeyToAddress(pub *PublicKey) *Address {
	pubBytes := FromECDSAPub(pub)
	a, err := NewAddressFromBytes(crypto.Keccak256(pubBytes[1:])[12:])
	if err == nil {
		return a
	}
	return nil
}
