package web3go

import (
	"golang.org/x/crypto/sha3"
)

// New256 ...
func New256() *GosHash {
	return &GosHash{sha3.New256()}
}
