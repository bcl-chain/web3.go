package web3go

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	DyTokenContract "github.com/bcl-chain/web3.go/contract/dytoken"
)

type DyTokenV2_ struct {
	abi     abi.ABI
	address common.Address
	DyToken *DyTokenContract.DyToken
	client  *EthereumClient
}

func NewDyTokenV2(address *Address, client *EthereumClient) (*DyTokenV2_, error) {
	parsed, err := abi.JSON(strings.NewReader(DyTokenContract.DyTokenABI))
	if err != nil {
		return nil, err
	}
	DyToken, err := DyTokenContract.NewDyToken(address.address, client.client)
	if err != nil {
		return nil, err
	}

	return &DyTokenV2_{
		abi:     parsed,
		address: address.address,
		DyToken: DyToken,
		client:  client,
	}, nil

}

func (DyToken *DyTokenV2_) BalanceOfV2(who *Address) (*BigInt, error) {
	balance, err := DyToken.DyToken.BalanceOf(nil, who.address)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}

func (DyToken *DyTokenV2_) SendTransferV2(tx *Transaction) error {
	return DyToken.client.SendTransaction(NewContext(), tx)
}

//把交易发送到网络签名
func (DyToken *DyTokenV2_) BuildTransferV2(opts *TransactOpts, to *Address, iamount string, decimals int, chainId int64) (*Transaction, error) {
	value := ToWei(iamount, decimals)
	input, err := DyToken.abi.Pack("transfer", to.address, value)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//增发功能

func (DyToken *DyTokenV2_) BuildMintV2(opts *TransactOpts, to *Address, iamount string, decimals int, chainId int64) (*Transaction, error) {
	value := ToWei(iamount, decimals)
	input, err := DyToken.abi.Pack("mint", to.address, value)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//销毁功能

func (DyToken *DyTokenV2_) BuildBurnV2(opts *TransactOpts, iamount string, decimals int, chainId int64) (*Transaction, error) {
	value := ToWei(iamount, decimals)
	input, err := DyToken.abi.Pack("burn", value)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//转移权限功能

func (DyToken *DyTokenV2_) BuildTransferGRCOwnershipV2(opts *TransactOpts, to *Address, chainId int64) (*Transaction, error) {
	input, err := DyToken.abi.Pack("transferGRCOwnership", to.address)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//禁用功能
func (DyToken *DyTokenV2_) BuildPauseV2(opts *TransactOpts, chainId int64) (*Transaction, error) {
	input, err := DyToken.abi.Pack("pause")
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//启用功能
func (DyToken *DyTokenV2_) BuildUnPauseV2(opts *TransactOpts, chainId int64) (*Transaction, error) {
	input, err := DyToken.abi.Pack("unpause")
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.NewEIP155Signer(big.NewInt(chainId)), opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//是否被禁用
// if it is paused
func (DyToken *DyTokenV2_) PausedV2() (bool, error) {
	flag, err := DyToken.DyToken.Paused(nil)
	if err != nil {
		return false, err
	}
	return flag, err
}

//判断是否有禁用权限
// if someone has the right to pause
func (DyToken *DyTokenV2_) IsPauserV2(to *Address) (bool, error) {
	ispauser, err := DyToken.DyToken.IsPauser(nil, to.address)
	if err != nil {
		return false, err
	}
	return ispauser, nil
}

//判断是否有增发权限
// if someone has the right to mint
func (DyToken *DyTokenV2_) IsMinterV2(to *Address) (bool, error) {
	isminter, err := DyToken.DyToken.IsMinter(nil, to.address)
	if err != nil {
		return false, err
	}
	return isminter, nil
}
