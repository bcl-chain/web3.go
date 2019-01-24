package common

import (
	"github.com/ethereum/go-ethereum/common"
)

type Address struct {
	address interface{}
}

type Hash struct {
	hash interface{}
}

func BytesToAddress(b []byte) *Address {
	return &Address{
		address: common.BytesToAddress(b),
	}
}

func (a *Address) Hex() string {
	address, _ := a.address.(common.Address)
	return address.Hex()
}
