package geth

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// Ecrecover ...
func Ecrecover(hash, sig []byte) ([]byte, error) {
	return crypto.Ecrecover(hash, sig)
}

// SigToPub ...
func SigToPub(hash, sig []byte) (*PublicKey, error) {
	pub, err := crypto.SigToPub(hash, sig)
	if err == nil {
		return &PublicKey{pub}, nil
	}
	return nil, err
}

// Sign ...
func Sign(hash []byte, wprv *PrivateKey) ([]byte, error) {
	prv := wprv.privateKey
	return crypto.Sign(hash, prv)
}

// VerifySignature ...
func VerifySignature(pubkey, hash, sig []byte) bool {
	return crypto.VerifySignature(pubkey, hash, sig)
}

// DecompressPubkey ...
func DecompressPubkey(pubkey []byte) (*PublicKey, error) {
	pub, err := crypto.DecompressPubkey(pubkey)
	if err == nil {
		return &PublicKey{pub}, nil
	}
	return nil, err
}

// CompressPubkey ...
func CompressPubkey(wpubkey *PublicKey) []byte {
	pubkey := wpubkey.publicKey
	return crypto.CompressPubkey(pubkey)
}

// TODO
//func S256 {
//}
