package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get ether balance from specific block
	account, _ := web3go.NewAddressFromHex("0xa3bcb71efafe7fdeb13333baaf8f0a06cafc2a11")
	balance, err := client.GetBalanceAt(web3go.NewContext(), account, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 3. get ether balance from the latest block
	balance2, err := client.GetBalanceAt(web3go.NewContext(), account, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance2)
}
