package web3go

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/bcl-chain/web3.go/contract"
)

type ERC20 struct {
	erc20 *contract.ERC20
}

func NewERC20(address common.Address, client *EthereumClient) (*ERC20, error) {
	erc20, err := contract.NewERC20(address, client.client)
	if err != nil {
		return nil, err
	}
	return &ERC20{erc20}, nil
}

func (erc20 *ERC20) BalanceOf(who common.Address) (*BigInt, error) {
	balance, err := erc20.erc20.BalanceOf(nil, who)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}
