package api

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type TransactOpts struct {
	transactOpts *bind.TransactOpts
}

func fromTransactOpts(to *bind.TransactOpts) *TransactOpts {
	if to != nil {
		return &TransactOpts{
			transactOpts: to,
		}
	} else {
		return nil
	}
}

func toTransactOpts(wto *TransactOpts) *bind.TransactOpts {
	if wto != nil {
		return wto.transactOpts
	} else {
		return nil
	}
}

func (wto *TransactOpts) SetNonce(wn *Int) {
	to := toTransactOpts(wto)
	n := toBigInt(wn)

	to.Nonce = n
}

func (wto *TransactOpts) SetValue(wv *Int) {
	to := toTransactOpts(wto)
	v := toBigInt(wv)

	to.Value = v
}

func (wto *TransactOpts) SetGasPrice(wgp *Int) {
	to := toTransactOpts(wto)
	gp := toBigInt(wgp)

	to.GasPrice = gp
}

func (wto *TransactOpts) SetGasLimit(wgl int64) {
	to := toTransactOpts(wto)
	gl := uint64(wgl)

	to.GasLimit = gl
}
