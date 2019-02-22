package geth

import(
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func NewKeccak256() *GosHash {
	return &GosHash{sha3.NewKeccak256()	}
}
