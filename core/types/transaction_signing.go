package types

import (
	"github.com/ethereum/go-ethereum/core/types"
)

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
