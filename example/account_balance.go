package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := geth.NewEthereumClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get ether balance from specific block
	account, _ := geth.NewAddressFromHex("0x0abb28e0270074d5552a66d5dd172fbcb9db4fd7")
	balance, err := client.GetBalanceAt(geth.NewContext(), account, 200)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 3. get ether balance from the latest block
	balance2, err := client.GetBalanceAt(geth.NewContext(), account, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance2)
}
