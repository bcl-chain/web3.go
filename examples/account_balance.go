package main

import (
	//  "context"
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/ethclient"
	"github.com/bcl-chain/web3.go/goos/context"
	"github.com/bcl-chain/web3.go/goos/math/big"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xfcbcc6047e8f7b12ba5f422666e4d8d7290a3459")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	blockNumber := big.NewInt(0)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)
}
