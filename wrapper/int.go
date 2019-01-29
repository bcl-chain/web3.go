package wrapper

import (
	"math/big"

	wbig "github.com/bcl-chain/web3.go/goos/math/big"
)

func FromBigInt(bigInt *big.Int) *wbig.Int {
	if bigInt != nil {
		return &wbig.Int{
			Int: bigInt,
		}
	} else {
		return nil
	}
}

func ToBigInt(wbigInt *wbig.Int) *big.Int {
	if wbigInt != nil {
		if bigInt, ok := wbigInt.Int.(*big.Int); ok {
			return bigInt
		} else {
			return nil
		}
	} else {
		return nil
	}
}

