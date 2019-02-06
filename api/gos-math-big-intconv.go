package api

import (
//	"math/big"
)

func (wx *Int) String() string {
	x := toBigInt(wx)
	return x.String()
}
