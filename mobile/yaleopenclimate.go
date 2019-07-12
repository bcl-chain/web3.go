package web3go

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/bcl-chain/web3.go/contract/yaleopenclimate"
)

type YToken_ struct {
	abi     abi.ABI
	address common.Address
	YToken   *YTokenContract.YToken
}

func NewYToken(address *Address, client *EthereumClient) (*YToken_, error) {

	parsed, err := abi.JSON(strings.NewReader(YTokenContract.YTokenABI))
	if err != nil {
		return nil, err
	}
	YToken, err := YTokenContract.NewYToken(address.address, client.client)
	if err != nil {
		return nil, err
	}

	return &YToken_{
		abi:     parsed,
		address: address.address,
		YToken:   YToken,
	}, nil

}

func (YToken *YToken_) BalanceOf(who *Address) (*BigInt, error) {
	balance, err := YToken.YToken.BalanceOf(nil, who.address)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}

//send txs to net to sign
func (YToken *YToken_) Transfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	tx, err := YToken.YToken.Transfer(opts.opts, to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx}, nil
}

func (YToken *YToken_) BuildTransfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	input, err := YToken.abi.Pack("transfer", to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), YToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//mint function
func (YToken *YToken_) Mint(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error){
	tx, err := YToken.YToken.Mint(opts.opts, to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx},err
}

func (YToken *YToken_) BuildMint(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	input, err := YToken.abi.Pack("mint", to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), YToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//burn function
func (YToken *YToken_) Burn(opts *TransactOpts, value *BigInt) (*Transaction, error){
	tx, err := YToken.YToken.Burn(opts.opts, value.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx},err
}

func (YToken *YToken_) BuildBurn(opts *TransactOpts, value *BigInt) (*Transaction, error) {
	input, err := YToken.abi.Pack("burn", value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), YToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//transfer token's ownership
func (YToken *YToken_) TransferGRCOwnership (opts *TransactOpts, to *Address)(*Transaction, error){
	tx, err := YToken.YToken.TransferGRCOwnership(opts.opts, to.address)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx},err
}

func (YToken *YToken_) BuildTransferGRCOwnership(opts *TransactOpts, to *Address) (*Transaction, error) {
	input, err := YToken.abi.Pack("transferGRCOwnership", to.address)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), YToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//pause function
func (YToken *YToken_) Pause(opts *TransactOpts) (*Transaction, error){
	tx, err := YToken.YToken.Pause(opts.opts)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx},err
}

func (YToken *YToken_) BuildPause(opts *TransactOpts) (*Transaction, error) {
	input, err := YToken.abi.Pack("pause")
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), YToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//unpause function
func (YToken *YToken_) UnPause(opts *TransactOpts) (*Transaction, error){
	tx, err := YToken.YToken.Unpause(opts.opts)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx},err
}

func (YToken *YToken_) BuildUnPause(opts *TransactOpts) (*Transaction, error) {
	input, err := YToken.abi.Pack("unpause")
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), YToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}


// if it is paused
func (YToken *YToken_) Paused() (bool, error){
	flag, err := YToken.YToken.Paused(nil)
	if err != nil {
		return false, err
	}
	return flag,err
}

// if someone has the right to pause
func (YToken *YToken_) IsPauser(to *Address) (bool,error){
	ispauser, err := YToken.YToken.IsPauser(nil, to.address)
	if err !=nil {
		return false, err
	}
	return ispauser, nil
}

// if someone has the right to mint
func (YToken *YToken_) IsMinter(to *Address) (bool,error){
	isminter, err := YToken.YToken.IsMinter(nil, to.address)
	if err !=nil {
		return false, err
	}
	return isminter, nil
}
