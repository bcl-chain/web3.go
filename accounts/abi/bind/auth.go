package bind

import (
	"github.com/bcl-chain/web3.go/crypto"

	wcrypto "github.com/bcl-chain/web3.go/wrapper/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewKeyedTransactor(wprv *crypto.PrivateKey) *TransactOpts {
	prv := wcrypto.ToPrivateKey(wprv)

	return &TransactOpts{
		TransactOpts: bind.NewKeyedTransactor(prv),
	}
}
