package ethclient

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/goos/context"
	"github.com/bcl-chain/web3.go/goos/math/big"
	"github.com/bcl-chain/web3.go/core/types"
	wcontext "github.com/bcl-chain/web3.go/wrapper/goos/context"
	wbig "github.com/bcl-chain/web3.go/wrapper/goos/math/big"
	wcommon "github.com/bcl-chain/web3.go/wrapper/common"
  wtypes "github.com/bcl-chain/web3.go/wrapper/core/types"
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
	ctx := wcontext.ToContext(wctx)
	address := wcommon.ToAddress(waccount)
	blockNumber := wbig.ToBigInt(wblockNumber)

	if bigInt, err := c.BalanceAt(ctx, address, blockNumber); err == nil {
		return wbig.FromBigInt(bigInt), nil
	} else {
		return nil, err
	}
}

func (ec *Client) BlockByNumber(wctx *context.Context, wnumber *big.Int) (*types.Block, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	blockNumber := wbig.ToBigInt(wnumber)

	if block, err := c.BlockByNumber(ctx, blockNumber); err == nil {
		return wtypes.FromBlock(block), nil
	} else {
		return nil, err
	}
}

func (ec *Client) HeaderByNumber(wctx *context.Context, wnumber *big.Int) (*types.Header, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	blockNumber := wbig.ToBigInt(wnumber)

	if header, err := c.HeaderByNumber(ctx, blockNumber); err == nil {
		return wtypes.FromHeader(header), nil
	} else {
		return nil, err
	}
}

func (ec *Client) TransactionByHash(wctx *context.Context, wblockHash *common.Hash) (*types.Transaction, bool, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	blockHash := wcommon.ToHash(wblockHash)

	if tx, isPending, err := c.TransactionByHash(ctx, blockHash); err == nil {
		return wtypes.FromTransaction(tx), isPending, err
	} else {
		return nil, false, err
	}
}

func (ec *Client) TransactionCount(wctx *context.Context, wblockHash *common.Hash) (int64, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	blockHash := wcommon.ToHash(wblockHash)

	count, err := c.TransactionCount(ctx, blockHash)
	return int64(count), err
}

func (ec *Client) TransactionInBlock(wctx *context.Context, wblockHash *common.Hash, index int64) (*types.Transaction, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	blockHash := wcommon.ToHash(wblockHash)
	uindex := uint(index)

	if tx, err := c.TransactionInBlock(ctx, blockHash, uindex); err == nil {
		return wtypes.FromTransaction(tx), err
	} else {
		return nil, err
	}
}

func (ec *Client) TransactionReceipt(wctx *context.Context, wblockHash *common.Hash) (*types.Receipt, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	blockHash := wcommon.ToHash(wblockHash)

	if receipt, err := c.TransactionReceipt(ctx, blockHash); err == nil {
		return wtypes.FromReceipt(receipt), nil
	} else {
		return nil, err
	}
}

func (ec *Client) PendingNonceAt(wctx *context.Context, waccount *common.Address) (int64, error) {
	c, _ := ec.c.(*ethclient.Client)
	ctx := wcontext.ToContext(wctx)
	account := wcommon.ToAddress(waccount)

  result, err := c.PendingNonceAt(ctx, account)
  return int64(result), err
}

func (ec *Client) SendTransaction(wctx *context.Context, wtx *types.Transaction) error {
  c, _ := ec.c.(*ethclient.Client)
  ctx := wcontext.ToContext(wctx)
  tx := wtypes.ToTransaction(wtx)

  return c.SendTransaction(ctx, tx)
}
