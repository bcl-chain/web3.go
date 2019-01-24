package common

import (
  "github.com/ethereum/go-ethereum/common"
)

type Address struct {
  address common.Address
}

type Hash struct {
  hash common.Hash
}

func (a Address) Hex() string {
  return a.address.Hex()
}
