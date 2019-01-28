package big

import(
  "math/big"
)

func (wx *Int) String() string {
  x, _ := wx.Int.(*big.Int)
  return x.String()
}
