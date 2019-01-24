package ethclient

import (
  "context"
  "math/big"

  "github.com/bcl-chain/web3.go/common"
  "github.com/bcl-chain/web3.go/wrapper"

  "github.com/ethereum/go-ethereum/ethclient"
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

func (ec *Client) BalanceAt(ctx context.Context, account *common.Address, blockNumber *big.Int) (*big.Int, error) {
  c, _ := ec.c.(*ethclient.Client)
  address := wrapper.ToAddress(account)
  return c.BalanceAt(ctx, address, blockNumber)
}

