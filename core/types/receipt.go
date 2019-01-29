package types

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Receipt struct {
	Receipt interface{}
}

func (wr *Receipt) Status() int64 {
	r, _ := wr.Receipt.(*types.Receipt)
	return int64(r.Status)
}
