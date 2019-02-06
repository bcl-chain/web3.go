// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package crypto implements a wrapper library for
// github.com/bcl-chain/go-ethereum/crypto.
//

package api

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	secp256k1N = InitInt().SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
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

func (priv *PrivateKey) Public() *PublicKey {
	privateKey := priv.privateKey
	return fromPublicKey(&privateKey.PublicKey)
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

// TODO: input type of original Keccak256 is ...[]byte.
//       K changed to []byte.
//       have to check if this modification is fine or not.
func Keccak256(data []byte) []byte {
	return crypto.Keccak256(data)
}

// TODO: input type of original Keccak256Hash is ...[]byte.
//       K changed to []byte.
//       have to check if this modification is fine or not.
func Keccak256Hash(data []byte) *Hash {
	h := crypto.Keccak256Hash(data)
	return fromHash(h)
}

// TODO: input type of original Keccak512 is ...[]byte.
//       K changed to []byte.
//       have to check if this modification is fine or not.
func Keccak512(data []byte) []byte {
	return crypto.Keccak512(data)
}

func CreateAddress(wb *Address, wnonce int64) *Address {
	b := toAddress(wb)
	nonce := uint64(wnonce)
	a := crypto.CreateAddress(b, nonce)
	return fromAddress(a)
}

// TODO
//func CreateAddress2 {
//}

func ToECDSA(d []byte) (*PrivateKey, error) {
	if priv, err := crypto.ToECDSA(d); err == nil {
		return fromPrivateKey(priv), nil
	} else {
		return nil, err
	}
}

func ToECDSAUnsafe(d []byte) *PrivateKey {
	priv := crypto.ToECDSAUnsafe(d)
	return fromPrivateKey(priv)
}

func FromECDSA(priv *PrivateKey) []byte {
	privateKey := toPrivateKey(priv)
	return crypto.FromECDSA(privateKey)
}

func UnmarshalPubkey(pub []byte) (*PublicKey, error) {
	if publicKey, err := crypto.UnmarshalPubkey(pub); err == nil {
		return fromPublicKey(publicKey), nil
	} else {
		return nil, err
	}
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

func LoadECDSA(file string) (*PrivateKey, error) {
	if privateKey, err := crypto.LoadECDSA(file); err == nil {
		return fromPrivateKey(privateKey), nil
	} else {
		return nil, err
	}
}

func SaveECDSA(file string, wpriv *PrivateKey) error {
	priv := toPrivateKey(wpriv)
	return crypto.SaveECDSA(file, priv)
}

func GenerateKey() (*PrivateKey, error) {
	if privateKey, err := crypto.GenerateKey(); err == nil {
		return fromPrivateKey(privateKey), nil
	} else {
		return nil, err
	}
}

func ValidateSignatureValues(v byte, wr, ws *Int, homestead bool) bool {
	r := toBigInt(wr)
	s := toBigInt(ws)
	return crypto.ValidateSignatureValues(v, r, s, homestead)
}

func PubkeyToAddress(pub *PublicKey) *Address {
	pubBytes := FromECDSAPub(pub)
	return BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
}
