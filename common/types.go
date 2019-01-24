package common

import (
//  "github.com/bcl-chain/web3.go/wrapper"

  "github.com/ethereum/go-ethereum/common"
)

type Address struct {
	Address interface{}
}

type Hash struct {
	Hash interface{}
}

func BytesToAddress(b []byte) *Address {
  address := common.BytesToAddress(b)
  return &Address{
    Address: address,
  }
}

func (a *Address) Hex() string {
  address, _ := a.Address.(common.Address)
  return address.Hex()
}

func HexToAddress(s string) *Address {
  address := common.HexToAddress(s)
  return &Address{
    Address: address,
  }
}
