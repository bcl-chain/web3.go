package api

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func SignTx(wtx *Transaction, ws *Signer, wprv *PrivateKey) (*Transaction, error) {
	tx := toTransaction(wtx)
	s := toSigner(ws)
	prv := toPrivateKey(wprv)

	if tx, err := types.SignTx(tx, s, prv); err == nil {
		return fromTransaction(tx), nil
	} else {
		return nil, err
	}
}

func Sender(ws *Signer, wtx *Transaction) (*Address, error) {
	s := toSigner(ws)
	tx := toTransaction(wtx)

	if a, err := types.Sender(s, tx); err == nil {
		return fromAddress(a), nil
	} else {
	 return nil, err
	}
}

type Signer struct {
	signer types.Signer
}

func fromSigner(s types.Signer) *Signer {
	if s != nil {
		return &Signer{
			signer: s,
		}
	} else {
		return nil
	}
}

func toSigner(ws *Signer) types.Signer {
	if ws != nil {
		return ws.signer
	} else {
		return nil
	}
}

type HomesteadSigner struct {
	homesteadSigner types.HomesteadSigner
}

func NewHomesteadSigner() *Signer {
	return fromSigner(types.HomesteadSigner{})
}
