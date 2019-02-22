package geth

import (
	"github.com/ethereum/go-ethereum/common"
)

func LeftPadBytes(slice []byte, l int) []byte {
	return common.LeftPadBytes(slice, l)
}
