package api

import (
	"github.com/ethereum/go-ethereum/ethclient"
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

func DialContext(wctx *Context, rawurl string) (*Client, error) {
	ctx := toContext(wctx)
	if c, err := ethclient.DialContext(ctx, rawurl); err == nil {
		return fromClient(c), nil
	} else {
		return nil, err
	}
}

// TODO
//func NewClient() *Client {
//}

func (wec *Client) Close() {
	ec := toClient(wec)
	ec.Close()
}

func (wec *Client) BlockByHash(wctx *Context, whash *Hash) (*Block, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	hash := toHash(whash)

	if block, err := ec.BlockByHash(ctx, hash); err == nil {
		return fromBlock(block), nil
	} else {
		return nil, err
	}
}

func (wec *Client) BlockByNumber(wctx *Context, wnumber *Int) (*Block, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockNumber := toBigInt(wnumber)

	if block, err := ec.BlockByNumber(ctx, blockNumber); err == nil {
		return fromBlock(block), nil
	} else {
		return nil, err
	}
}

func (wec *Client) BalanceAt(wctx *Context, waccount *Address, wblockNumber *Int) (*Int, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	address := toAddress(waccount)
	blockNumber := toBigInt(wblockNumber)

	if bigInt, err := ec.BalanceAt(ctx, address, blockNumber); err == nil {
		return fromBigInt(bigInt), nil
	} else {
		return nil, err
	}
}

func (wec *Client) HeaderByHash(wctx *Context, whash *Hash) (*Header, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	hash := toHash(whash)

	if header, err := ec.HeaderByHash(ctx, hash); err == nil {
		return fromHeader(header), nil
	} else {
		return nil, err
	}
}

func (wec *Client) HeaderByNumber(wctx *Context, wnumber *Int) (*Header, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockNumber := toBigInt(wnumber)

	if header, err := ec.HeaderByNumber(ctx, blockNumber); err == nil {
		return fromHeader(header), nil
	} else {
		return nil, err
	}
}

// divied into two functions; TransactionByHash and TransactionByHashIsPending
func (wec *Client) TransactionByHash(wctx *Context, wblockHash *Hash) (*Transaction, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	if tx, _, err := ec.TransactionByHash(ctx, blockHash); err == nil {
		return fromTransaction(tx), err
	} else {
		return nil, err
	}
}

func (wec *Client) TransactionByHashIsPending(wctx *Context, wblockHash *Hash) (bool, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	if _, isPending, err := ec.TransactionByHash(ctx, blockHash); err == nil {
		return isPending, err
	} else {
		return false, err
	}
}

func (wec *Client) TransactionSender(wctx *Context, wtx *Transaction, wblock *Hash, windex int) (*Address, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	tx := toTransaction(wtx)
	block := toHash(wblock)
	index := uint(windex)

	if address, err := ec.TransactionSender(ctx, tx, block, index); err == nil {
		return fromAddress(address), nil
	} else {
		return nil, err
	}
}

func (wec *Client) TransactionCount(wctx *Context, wblockHash *Hash) (int64, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	count, err := ec.TransactionCount(ctx, blockHash)
	return int64(count), err
}

func (wec *Client) TransactionInBlock(wctx *Context, wblockHash *Hash, index int64) (*Transaction, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)
	uindex := uint(index)

	if tx, err := ec.TransactionInBlock(ctx, blockHash, uindex); err == nil {
		return fromTransaction(tx), err
	} else {
		return nil, err
	}
}

func (wec *Client) TransactionReceipt(wctx *Context, wblockHash *Hash) (*Receipt, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	blockHash := toHash(wblockHash)

	if receipt, err := ec.TransactionReceipt(ctx, blockHash); err == nil {
		return fromReceipt(receipt), nil
	} else {
		return nil, err
	}
}

// TODO
//func SyncProgress {
//}

// TODO
//func SubscribeNewHead {
//}

func (wec *Client) NetworkID(wctx *Context) (*Int, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)

	if i, err := ec.NetworkID(ctx); err == nil {
		return fromBigInt(i), nil
	} else {
		return nil, err
	}
}

func (wec *Client) StorageAt(wctx *Context, waccount *Address, wkey *Hash, wblockNumber *Int) ([]byte, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)
	key := toHash(wkey)
	blockNumber := toBigInt(wblockNumber)

	return ec.StorageAt(ctx, account, key, blockNumber)
}

func (wec *Client) CodeAt(wctx *Context, waccount *Address, wblockNumber *Int) ([]byte, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)
	blockNumber := toBigInt(wblockNumber)

	return ec.CodeAt(ctx, account, blockNumber)
}

// changed return from uint64 to int64
func (wec *Client) NonceAt(wctx *Context, waccount *Address, wblockNumber *Int) (int64, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)
	blockNumber := toBigInt(wblockNumber)

	i, err := ec.NonceAt(ctx, account, blockNumber)
	return int64(i), err
}

// TODO
//func FilterLogs {
//}

// TODO
//func SubscribeFilterLogs {
//}

func (wec *Client) PendingBalanceAt(wctx *Context, waccount *Address) (*Int, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)

	result, err := ec.PendingBalanceAt(ctx, account)
	return fromBigInt(result), err
}

func (wec *Client) PendingStorageAt(wctx *Context, waccount *Address, wkey *Hash) ([]byte, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)
	key := toHash(wkey)

	return ec.PendingStorageAt(ctx, account, key)
}

func (wec *Client) PendingCodeAt(wctx *Context, waccount *Address) ([]byte, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)

	return ec.PendingCodeAt(ctx, account)
}

// changed return from uint64 to int64
func (wec *Client) PendingNonceAt(wctx *Context, waccount *Address) (int64, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)
	account := toAddress(waccount)

	result, err := ec.PendingNonceAt(ctx, account)
	return int64(result), err
}

// changed return from uint64 to int64
func (wec *Client) PendingTransactionCount(wctx *Context) (int64, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)

	result, err := ec.PendingTransactionCount(ctx)
	return int64(result), err
}

// TODO
//func CallContract {
//}

// TODO
//func PendingCallContract {
//}

func (wec *Client) SuggestGasPrice(wctx *Context) (*Int, error) {
	ec := toClient(wec)
	ctx := toContext(wctx)

	if gasPrice, err := ec.SuggestGasPrice(ctx); err == nil {
		return fromBigInt(gasPrice), nil
	} else {
		return nil, err
	}
}

// TODO
//func EstimateGas {
//}

func (wec *Client) SendTransaction(wctx *Context, wtx *Transaction) error {
	ec := toClient(wec)
	ctx := toContext(wctx)
	tx := toTransaction(wtx)

	return ec.SendTransaction(ctx, tx)
}
