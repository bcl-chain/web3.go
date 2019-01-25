package main

import (
//  "context"
  "fmt"
  "log"

  "github.com/bcl-chain/web3.go/common"
  "github.com/bcl-chain/web3.go/common/math/big"
  "github.com/bcl-chain/web3.go/common/context"
  "github.com/bcl-chain/web3.go/ethclient"
)

func main() {
  client, err := ethclient.Dial("http://127.0.0.1:8545")
  if err != nil {
    log.Fatal(err)
  }

  account := common.HexToAddress("0x0abb28e0270074d5552a66d5dd172fbcb9db4fd7")
  balance, err := client.BalanceAt(context.Background(), account, nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(balance)

  blockNumber := big.NewInt(3)
  balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(balanceAt)
}
