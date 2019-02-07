package api

import (
	"golang.org/x/crypto/sha3"
)

func NewKeccak256() *GosHash {
	h := sha3.NewLegacyKeccak256()
	return fromGosHash(h)
}
