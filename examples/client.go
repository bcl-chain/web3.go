package main

import (
	"fmt"
	"github.com/bcl-chain/web3.go/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client // we'll use this in the next section
}
