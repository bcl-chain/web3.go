package api

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewKeyedTransactor(wprv *PrivateKey) *TransactOpts {
	prv := toPrivateKey(wprv)

	return fromTransactOpts(bind.NewKeyedTransactor(prv))
}
