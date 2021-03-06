package web3go

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bcl-chain/web3.go/contract/erc20"
)

type ERC20 struct {
	abi     abi.ABI
	address common.Address
	erc20   *contract.ERC20
}

//封装erc20
func NewERC20(address *Address, client *EthereumClient) (*ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(contract.ERC20ABI))
	if err != nil {
		return nil, err
	}
	erc20, err := contract.NewERC20(address.address, client.client)
	if err != nil {
		return nil, err
	}
	return &ERC20{
		abi:     parsed,
		address: address.address,
		erc20:   erc20,
	}, nil
}

func (erc20 *ERC20) BalanceOf(who *Address) (*BigInt, error) {
	balance, err := erc20.erc20.BalanceOf(nil, who.address)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}


func (erc20 *ERC20) BuildTransfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	input, err := erc20.abi.Pack("transfer", to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), erc20.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

func (erc20 *ERC20) Transfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	tx, err := erc20.erc20.Transfer(opts.opts, to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx}, nil
}
