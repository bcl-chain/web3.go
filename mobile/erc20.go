package web3go

import (
	"github.com/bcl-chain/web3.go/contract"
)

type ERC20 struct {
	erc20 *contract.ERC20
}

func NewERC20(address *Address, client *EthereumClient) (*ERC20, error) {
	erc20, err := contract.NewERC20(address.address, client.client)
	if err != nil {
		return nil, err
	}
	return &ERC20{erc20}, nil
}

func (erc20 *ERC20) BalanceOf(who *Address) (*BigInt, error) {
	balance, err := erc20.erc20.BalanceOf(nil, who.address)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}

func (erc20 *ERC20) Transfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	tx, err := erc20.erc20.Transfer(opts.opts, to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx}, nil
}
