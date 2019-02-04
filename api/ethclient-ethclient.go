package api

import (
	"github.com/ethereum/go-ethereum/ethclient"
	//	"github.com/bcl-chain/web3.go/common"
	//	"github.com/bcl-chain/web3.go/core/types"
	//	"github.com/bcl-chain/web3.go/gos/context"
	//	"github.com/bcl-chain/web3.go/gos/math/big"
	//	wcommon "github.com/bcl-chain/web3.go/wrapper/common"
	//	wtypes "github.com/bcl-chain/web3.go/wrapper/core/types"
	//	wcontext "github.com/bcl-chain/web3.go/wrapper/gos/context"
	//	wbig "github.com/bcl-chain/web3.go/wrapper/gos/math/big"
)

type Client struct {
	c *ethclient.Client
}

func fromClient(c *ethclient.Client) *Client {
	if c != nil {
		return &Client{
			c: c,
		}
	} else {
		return nil
	}
}

func toClient(wc *Client) *ethclient.Client {
	if wc != nil {
		c := wc.c
		return c
	} else {
		return nil
	}
}

func Dial(rawurl string) (*Client, error) {
	if c, err := ethclient.Dial(rawurl); err == nil {
		return fromClient(c), nil
	} else {
		return nil, err
	}
}

func (ec *Client) BalanceAt(wctx *Context, waccount *Address, wblockNumber *Int) (*Int, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	address := toAddress(waccount)
	blockNumber := toBigInt(wblockNumber)

	if bigInt, err := c.BalanceAt(ctx, address, blockNumber); err == nil {
		return fromBigInt(bigInt), nil
	} else {
		return nil, err
	}
}

func (ec *Client) BlockByNumber(wctx *Context, wnumber *Int) (*Block, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockNumber := toBigInt(wnumber)

	if block, err := c.BlockByNumber(ctx, blockNumber); err == nil {
		return fromBlock(block), nil
	} else {
		return nil, err
	}
}

func (ec *Client) HeaderByNumber(wctx *Context, wnumber *Int) (*Header, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockNumber := toBigInt(wnumber)

	if header, err := c.HeaderByNumber(ctx, blockNumber); err == nil {
		return fromHeader(header), nil
	} else {
		return nil, err
	}
}

// divied into two functions; TransactionByHash and TransactionByHashIsPending
func (ec *Client) TransactionByHash(wctx *Context, wblockHash *Hash) (*Transaction, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	if tx, _, err := c.TransactionByHash(ctx, blockHash); err == nil {
		return fromTransaction(tx), err
	} else {
		return nil, err
	}
}

func (ec *Client) TransactionByHashIsPending(wctx *Context, wblockHash *Hash) (bool, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	if _, isPending, err := c.TransactionByHash(ctx, blockHash); err == nil {
		return isPending, err
	} else {
		return false, err
	}
}

func (ec *Client) TransactionCount(wctx *Context, wblockHash *Hash) (int64, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	count, err := c.TransactionCount(ctx, blockHash)
	return int64(count), err
}

func (ec *Client) TransactionInBlock(wctx *Context, wblockHash *Hash, index int64) (*Transaction, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)
	uindex := uint(index)

	if tx, err := c.TransactionInBlock(ctx, blockHash, uindex); err == nil {
		return fromTransaction(tx), err
	} else {
		return nil, err
	}
}

func (ec *Client) TransactionReceipt(wctx *Context, wblockHash *Hash) (*Receipt, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	if receipt, err := c.TransactionReceipt(ctx, blockHash); err == nil {
		return fromReceipt(receipt), nil
	} else {
		return nil, err
	}
}

func (ec *Client) PendingNonceAt(wctx *Context, waccount *Address) (int64, error) {
	c := toClient(ec)
	ctx := toContext(wctx)
	account := toAddress(waccount)

	result, err := c.PendingNonceAt(ctx, account)
	return int64(result), err
}

func (ec *Client) SuggestGasPrice(wctx *Context) (*Int, error) {
	c := toClient(ec)
	ctx := toContext(wctx)

	if gasPrice, err := c.SuggestGasPrice(ctx); err == nil {
		return fromBigInt(gasPrice), nil
	} else {
		return nil, err
	}
}

func (ec *Client) SendTransaction(wctx *Context, wtx *Transaction) error {
	c := toClient(ec)
	ctx := toContext(wctx)
	tx := toTransaction(wtx)

	return c.SendTransaction(ctx, tx)
}
