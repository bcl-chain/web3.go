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
	fmt.Println("we have a connection")

	_ = client

	//	// 2. get networkID
	//	id, err := web3go.NetworkID(api.NewContext())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("networkId:", id)
	//
	//	// 3. close connection to client
	//	geth.Close()
	//	fmt.Println("connection is closed")
}
