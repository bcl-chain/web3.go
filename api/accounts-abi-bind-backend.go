package api

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ContractBackend struct {
	contractBackend bind.ContractBackend
}

func fromContractBackend(cb bind.ContractBackend) *ContractBackend {
	if cb != nil {
		return &ContractBackend{
			contractBackend: cb,
		}
	} else {
		return nil
	}
}

func toContractBackend(wcb *ContractBackend) bind.ContractBackend {
	if wcb != nil {
		return wcb.contractBackend
	} else {
		return nil
	}
}
