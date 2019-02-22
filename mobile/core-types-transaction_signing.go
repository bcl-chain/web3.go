package geth

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// SignTx ...
func SignTx(wtx *Transaction, ws *Signer2, wprv *PrivateKey) (*Transaction, error) {
	tx := wtx.tx
	s := ws.signer
	prv := wprv.privateKey

	tx, err := types.SignTx(tx, s, prv)
	if err == nil {
		return &Transaction{tx: tx}, nil
	}
	return nil, err
}

// Sender ...
func Sender(ws *Signer2, wtx *Transaction) (*Address, error) {
	s := ws.signer
	tx := wtx.tx

	a, err := types.Sender(s, tx)
	if err == nil {
		return &Address{address: a}, nil
	}
	return nil, err
}

// Signer2 ...
type Signer2 struct {
	signer types.Signer
}

// HomesteadSigner ...
type HomesteadSigner struct {
	homesteadSigner types.HomesteadSigner
}

// NewHomesteadSigner ...
func NewHomesteadSigner() *Signer2 {
	return &Signer2{types.HomesteadSigner{}}
}
