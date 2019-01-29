package types

import (
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bcl-chain/web3.go/crypto"
	wcrypto "github.com/bcl-chain/web3.go/wrapper/crypto"
)

func SignTx(wtx *Transaction, ws *Signer, wprv *crypto.PrivateKey) (*Transaction, error) {
	tx, _ := wtx.Transaction.(*types.Transaction)
	s, _ := ws.Signer.(types.Signer)
	prv := wcrypto.ToPrivateKey(wprv)

	if tx, err := types.SignTx(tx, s, prv); err == nil {
		return &Transaction{
			Transaction: tx,
		}, nil
	} else {
		return nil, err
	}
}

type Signer struct {
	Signer interface{}
}

type HomesteadSigner struct {
	HomesteadSigner interface{}
}

func NewHomesteadSigner() *Signer {
	return &Signer{
		Signer: types.HomesteadSigner{},
	}
}
