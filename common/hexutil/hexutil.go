package hexutil

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Encode(b []byte) string {
	return hexutil.Encode(b)
}
