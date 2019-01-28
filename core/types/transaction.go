package types

import (
  "github.com/ethereum/go-ethereum/core/types"
)

type Transactions struct {
  Transactions interface{}
}

func (ws Transactions) Len() int {
  s, _ := ws.Transactions.(types.Transactions)
  return s.Len()
}
