package big

import (
	"math/big"
)

type Int struct {
	Int interface{}
}

func NewInt(x int64) *Int {
	return &Int{
		Int: big.NewInt(x),
	}
}

func (wx *Int) Int64() int64 {
	x, _ := wx.Int.(*big.Int)
	return x.Int64()
}
