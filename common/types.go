package common

import (
	"github.com/ethereum/go-ethereum/common"
)

/////////// Hash

type Hash struct {
  Hash interface{}
}

func HexToHash(s string) *Hash {
	return &Hash{
		Hash: common.HexToHash(s),
	}
}

func (wh Hash) Hex() string {
  h, _ := wh.Hash.(common.Hash)
  return h.Hex()
}

/////////// Address

type Address struct {
  Address interface{}
}

func BytesToAddress(b []byte) *Address {
  a := common.BytesToAddress(b)
  return &Address{
    Address: a,
  }
}

func (wa *Address) Hex() string {
  a, _ := wa.Address.(common.Address)
	return a.Hex()
}

func HexToAddress(s string) *Address {
  a := common.HexToAddress(s)
  return &Address{
    Address: a,
  }
}

