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
	fmt.Println("we have a connection")

	// 2. get networkID
	id, err := client.NetworkID(api.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("networkId:", id)

	// 3. close connection to client
	client.Close()
	fmt.Println("connection is closed")
}
