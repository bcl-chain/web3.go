package web3go

import (
  "math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bcl-chain/web3.go/contract/dytoken"
)

type DyToken_ struct {
	abi     abi.ABI
	address common.Address
	DyToken   *DyTokenContract.DyToken
}

func NewDyToken(address *Address, client *EthereumClient) (*DyToken_, error) {

  parsed, err := abi.JSON(strings.NewReader(DyTokenContract.DyTokenABI))
	if err != nil {
		return nil, err
	}
	DyToken, err := DyTokenContract.NewDyToken(address.address, client.client)
	if err != nil {
		return nil, err
	}

	return &DyToken_{
		abi:     parsed,
		address: address.address,
		DyToken:   DyToken,
	}, nil

}

func (DyToken *DyToken_) BalanceOf(who *Address) (*BigInt, error) {
	balance, err := DyToken.DyToken.BalanceOf(nil, who.address)
	if err != nil {
		return nil, err
	}
	return &BigInt{balance}, nil
}

//把交易发送到网络签名
func (DyToken *DyToken_) Transfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	tx, err := DyToken.DyToken.Transfer(opts.opts, to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx}, nil
}

func (DyToken *DyToken_) BuildTransfer(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	input, err := DyToken.abi.Pack("transfer", to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//增发功能
func (DyToken *DyToken_) Mint(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error){
    tx, err := DyToken.DyToken.Mint(opts.opts, to.address, value.bigint)
    if err != nil {
      return nil, err
    }
    return &Transaction{tx},err
}

func (DyToken *DyToken_) BuildMint(opts *TransactOpts, to *Address, value *BigInt) (*Transaction, error) {
	input, err := DyToken.abi.Pack("mint", to.address, value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//销毁功能
func (DyToken *DyToken_) Burn(opts *TransactOpts, value *BigInt) (*Transaction, error){
    tx, err := DyToken.DyToken.Burn(opts.opts, value.bigint)
    if err != nil {
      return nil, err
    }
    return &Transaction{tx},err
}

func (DyToken *DyToken_) BuildBurn(opts *TransactOpts, value *BigInt) (*Transaction, error) {
	input, err := DyToken.abi.Pack("burn", value.bigint)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//转移权限功能
func (DyToken *DyToken_) TransferGRCOwnership (opts *TransactOpts, to *Address)(*Transaction, error){
  tx, err := DyToken.DyToken.TransferGRCOwnership(opts.opts, to.address)
  if err != nil {
    return nil, err
  }
  return &Transaction{tx},err
}

func (DyToken *DyToken_) BuildTransferGRCOwnership(opts *TransactOpts, to *Address) (*Transaction, error) {
	input, err := DyToken.abi.Pack("transferGRCOwnership", to.address)
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//禁用功能
func (DyToken *DyToken_) Pause(opts *TransactOpts) (*Transaction, error){
    tx, err := DyToken.DyToken.Pause(opts.opts)
    if err != nil {
      return nil, err
    }
    return &Transaction{tx},err
}

func (DyToken *DyToken_) BuildPause(opts *TransactOpts) (*Transaction, error) {
	input, err := DyToken.abi.Pack("pause")
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//启用功能
func (DyToken *DyToken_) UnPause(opts *TransactOpts) (*Transaction, error){
    tx, err := DyToken.DyToken.Unpause(opts.opts)
    if err != nil {
      return nil, err
    }
    return &Transaction{tx},err
}

func (DyToken *DyToken_) BuildUnPause(opts *TransactOpts) (*Transaction, error) {
	input, err := DyToken.abi.Pack("unpause")
	if err != nil {
		return nil, err
	}
	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	rawTx := types.NewTransaction(opts.opts.Nonce.Uint64(), DyToken.address, amount, opts.opts.GasLimit, opts.opts.GasPrice, input)
	signedTx, err := opts.opts.Signer(types.HomesteadSigner{}, opts.opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return &Transaction{signedTx}, nil
}

//是否被禁用
func (DyToken *DyToken_) Paused() (bool, error){
    flag, err := DyToken.DyToken.Paused(nil)
    if err != nil {
      return false, err
    }
    return flag,err
}

//判断是否有禁用权限
func (DyToken *DyToken_) IsPauser(to *Address) (bool,error){
    ispauser, err := DyToken.DyToken.IsPauser(nil, to.address)
    if err !=nil {
      return false, err
    }
    return ispauser, nil
}

//判断是否有增发权限
func (DyToken *DyToken_) IsMinter(to *Address) (bool,error){
    isminter, err := DyToken.DyToken.IsMinter(nil, to.address)
    if err !=nil {
      return false, err
    }
    return isminter, nil
}
