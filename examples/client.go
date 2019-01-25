package main

import (
	"fmt"
  "log"

  "github.com/bcl-chain/web3.go/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client // we'll use this in the next section
}
