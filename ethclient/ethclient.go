package ethclient

import (
  "github.com/ethereum/go-ethereum/ethclient"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/goos/context"
	"github.com/bcl-chain/web3.go/goos/math/big"
	"github.com/bcl-chain/web3.go/wrapper"
  "github.com/bcl-chain/web3.go/core/types"
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

func (ec *Client) BlockByNumber(wctx *context.Context, wnumber *big.Int) (*types.Block, error) {
  c, _ := ec.c.(*ethclient.Client)
  context := wrapper.ToContext(wctx)
  blockNumber := wrapper.ToBigInt(wnumber)

  if block, err := c.BlockByNumber(context, blockNumber); err == nil {
    return wrapper.FromBlock(block), nil
  } else {
    return nil, err
  }
}

func (ec *Client) HeaderByNumber(wctx *context.Context, wnumber *big.Int) (*types.Header, error) {
  c, _ := ec.c.(*ethclient.Client)
  context := wrapper.ToContext(wctx)
  blockNumber := wrapper.ToBigInt(wnumber)

  if header, err := c.HeaderByNumber(context, blockNumber); err == nil {
    return wrapper.FromHeader(header), nil
  } else {
    return nil, err
  }
}

// TODO: wrap uint type
func (ec *Client) TransactionCount(wctx *context.Context, wblockHash *common.Hash) (uint, error) {
  c, _ := ec.c.(*ethclient.Client)
  context := wrapper.ToContext(wctx)
  blockHash := wrapper.ToHash(wblockHash)

  count, err := c.TransactionCount(context, blockHash)
  return count, err
}

