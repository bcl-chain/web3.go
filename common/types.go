package common

import (
  "github.com/ethereum/go-ethereum/common"
)

/////////// Hash

type Hash struct {
	Hash interface{}
}

func (h Hash) Hex() string{
  hash, _ := h.Hash.(common.Hash)
  return hash.Hex()
}

/////////// Address

type Address struct {
  Address interface{}
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
