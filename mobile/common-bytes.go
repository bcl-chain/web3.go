package geth

import (
	"github.com/ethereum/go-ethereum/common"
)

//LeftPadBytes ...
func LeftPadBytes(slice []byte, l int) []byte {
	return common.LeftPadBytes(slice, l)
}
