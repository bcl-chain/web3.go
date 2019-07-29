package web3go

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	contract "github.com/bcl-chain/web3.go/contract/erc20"
)

type ERC20V2 struct {
	abi     abi.ABI
	address common.Address
	erc20   *contract.ERC20
	client  *EthereumClient
}

//封装erc20
func NewERC20V2(address *Address, client *EthereumClient) (*ERC20V2, error) {
	parsed, err := abi.JSON(strings.NewReader(contract.ERC20ABI))
	if err != nil {
		return nil, err
	}
	erc20, err := contract.NewERC20(address.address, client.client)
	if err != nil {
		return nil, err
	}
	return &ERC20V2{
		abi:     parsed,
		address: address.address,
		erc20:   erc20,
		client:  client,
	}, nil
}

func (erc20 *ERC20V2) BalanceOfV2(who *Address) (*BigInt, error) {
	balance, err := erc20.erc20.BalanceOf(nil, who.address)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}

func (erc20 *ERC20V2) BuildTransferV2(opts *TransactOpts, to *Address, iamount string, decimals int, chainId int64) (*Transaction, error) {
	value := ToWei(iamount, decimals)
	input, err := erc20.abi.Pack("transfer", to.address, value)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), erc20.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

func (erc20 *ERC20V2) SendTransferV2(tx *Transaction) error {
	return erc20.client.SendTransaction(NewContext(), tx)
}
