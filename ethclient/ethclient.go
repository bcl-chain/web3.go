package ethclient

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/goos/context"
	"github.com/bcl-chain/web3.go/goos/math/big"
	"github.com/bcl-chain/web3.go/wrapper"
)

type Client struct {
	c interface{}
}

func Dial(rawurl string) (*Client, error) {
	if c, err := ethclient.Dial(rawurl); err == nil {
		return &Client{
			c: c,
		}, nil
	} else {
		return nil, err
	}
}

func (ec *Client) BalanceAt(wctx *context.Context, waccount *common.Address, wblockNumber *big.Int) (*big.Int, error) {
	c, _ := ec.c.(*ethclient.Client)
	context := wrapper.ToContext(wctx)
	address := wrapper.ToAddress(waccount)
	blockNumber := wrapper.ToBigInt(wblockNumber)

	if bigInt, err := c.BalanceAt(context, address, blockNumber); err == nil {
		return wrapper.FromBigInt(bigInt), nil
	} else {
		return nil, err
	}
}
