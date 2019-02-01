package big

import (
	"fmt"

	"math/big"
)

type Int struct {
	Int interface{}
}

func InitInt() *Int {
	return &Int{
		Int: new(big.Int),
	}
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

func (wz *Int) SetString(s string, base int) (*Int, bool) {
	z, _ := wz.Int.(*big.Int)
	i, b := z.SetString(s, base)
	fmt.Println(i,b)
	return &Int{
		Int: i,
	}, b
}

func (wz *Int) Bytes() []byte {
	z, _ := wz.Int.(*big.Int)
	return z.Bytes()
}
