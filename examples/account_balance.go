package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/api"
//	"github.com/bcl-chain/web3.go/common"
//	"github.com/bcl-chain/web3.go/ethclient"
//	"github.com/bcl-chain/web3.go/gos/context"
//	"github.com/bcl-chain/web3.go/gos/math/big"
)

func main() {
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	account := api.HexToAddress("0x0abb28e0270074d5552a66d5dd172fbcb9db4fd7")
	balance, err := client.BalanceAt(api.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	blockNumber := api.NewInt(3)
	balanceAt, err := client.BalanceAt(api.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)
}
