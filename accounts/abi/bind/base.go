package bind

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/bcl-chain/web3.go/goos/math/big"

	wbig "github.com/bcl-chain/web3.go/wrapper/goos/math/big"
)

type TransactOpts struct {
	TransactOpts interface{}
}

func (wto *TransactOpts) SetNonce(wn *big.Int) {
	to, _ := wto.TransactOpts.(bind.TransactOpts)
	n := wbig.ToBigInt(wn)

	to.Nonce = n
}

func (wto *TransactOpts) SetValue(wv *big.Int) {
  to, _ := wto.TransactOpts.(bind.TransactOpts)
  v := wbig.ToBigInt(wv)

  to.Value = v
}

func (wto *TransactOpts) SetGasPrice(wgp *big.Int) {
  to, _ := wto.TransactOpts.(bind.TransactOpts)
  gp := wbig.ToBigInt(wgp)

  to.GasPrice = gp
}

func (wto *TransactOpts) SetGasLimit(wgl int64) {
  to, _ := wto.TransactOpts.(bind.TransactOpts)
  gl := uint64(wgl)

  to.GasLimit = gl
}

