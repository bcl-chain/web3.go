package api

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Receipt struct {
	receipt *types.Receipt
}

func fromReceipt(r *types.Receipt) *Receipt {
	if r != nil {
		return &Receipt{
			receipt: r,
		}
	} else {
		return nil
	}
}

func toReceipt(wr *Receipt) *types.Receipt {
	if wr != nil {
		return wr.receipt
	} else {
		return nil
	}
}

func (wr *Receipt) Status() int64 {
	r := toReceipt(wr)
	return int64(r.Status)
}
