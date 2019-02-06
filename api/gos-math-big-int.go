package api

import (
	"math/big"
)

type Int struct {
	bigInt *big.Int
}

func fromBigInt(bigInt *big.Int) *Int {
	if bigInt != nil {
		return &Int{
			bigInt: bigInt,
		}
	} else {
		return nil
	}
}

func toBigInt(wbigInt *Int) *big.Int {
	if wbigInt != nil {
		return wbigInt.bigInt
	} else {
		return nil
	}
}

func InitInt() *Int {
	return fromBigInt(new(big.Int))
}

func NewInt(x int64) *Int {
	i := big.NewInt(x)
	return fromBigInt(i)
}

func (wx *Int) Int64() int64 {
	x := toBigInt(wx)
	return x.Int64()
}

// modified SetString so that one value is returned
func (wz *Int) SetString(s string, base int) *Int {
	z := toBigInt(wz)
	i, _ := z.SetString(s, base)
	return fromBigInt(i)
}

func (wz *Int) Bytes() []byte {
	z := toBigInt(wz)
	return z.Bytes()
}
