package crypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return crypto.Ecrecover(hash, sig)
}

func SigToPub(hash, sig []byte) (*PublicKey, error) {
  if pub, err := crypto.SigToPub(hash, sig); err == nil {
		return &PublicKey{
			PublicKey: pub,
		}, nil
	} else {
		return nil, err
	}
}

func Sign(hash []byte, wprv *PrivateKey) ([]byte, error) {
	prv := wprv.PrivateKey.(*ecdsa.PrivateKey)
	return crypto.Sign(hash, prv)
}

func VerifySignature(pubkey, hash, sig []byte) bool {
	return crypto.VerifySignature(pubkey, hash, sig)
}
