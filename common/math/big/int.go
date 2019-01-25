package big

import(
  "math/big"
)

type Int struct {
  Int interface{}
}

func NewInt(x int64) *Int {
  return &Int{
    Int: big.NewInt(x),
  }
}
