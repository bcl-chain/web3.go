package geth

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func SignTx(wtx *Transaction, ws *Signer2, wprv *PrivateKey) (*Transaction, error) {
	tx := wtx.tx
	s := toSigner2(ws)
	prv := toPrivateKey(wprv)

	if tx, err := types.SignTx(tx, s, prv); err == nil {
	 return &Transaction{tx:tx}, nil
	} else {
		return nil, err
	}
}

func Sender(ws *Signer2, wtx *Transaction) (*Address, error) {
	s := toSigner2(ws)
	tx := wtx.tx

	if a, err := types.Sender(s, tx); err == nil {
	 return &Address{address:a}, nil
	} else {
	 return nil, err
	}
}

type Signer2 struct {
	signer types.Signer
}

func fromSigner2(s types.Signer) *Signer2 {
	if s != nil {
		return &Signer2{
			signer: s,
		}
	} else {
		return nil
	}
}

func toSigner2(ws *Signer2) types.Signer {
	if ws != nil {
		return ws.signer
	} else {
		return nil
	}
}

type HomesteadSigner struct {
	homesteadSigner types.HomesteadSigner
}

func NewHomesteadSigner() *Signer2 {
	return fromSigner2(types.HomesteadSigner{})
}
