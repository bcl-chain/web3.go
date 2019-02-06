package api

import (
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func NewKeccak256() *GosHash {
	h := sha3.NewKeccak256()
	return fromGosHash(h)
}
