package crypto

import (
	"crypto/ecdsa"

	"github.com/bcl-chain/web3.go/crypto"
)

func FromPrivateKey(prv *ecdsa.PrivateKey) *crypto.PrivateKey {
	if prv != nil {
		return &crypto.PrivateKey{
			PrivateKey: prv,
		}
	} else {
		return nil
	}
}

func ToPrivateKey(wprv *crypto.PrivateKey) *ecdsa.PrivateKey {
	if wprv != nil {
		if prv, ok := wprv.PrivateKey.(*ecdsa.PrivateKey); ok {
			return prv
		} else {
			return nil
		}
	} else {
		return nil
	}
}
