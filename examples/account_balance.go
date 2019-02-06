package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/api"
)

func main() {
	// 1. connect to client
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get ether balance in latest block
	account := api.HexToAddress("0xFBcDe7c69381f5e3e0C02f983da48A65d88DD02a")
	balance, err := client.BalanceAt(api.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 3. get ether balance in specific block
	blockNumber := api.NewInt(2574)
	balanceAt, err := client.BalanceAt(api.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)
}
