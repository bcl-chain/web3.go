package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/api"
)

func main() {
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(api.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number()
	fmt.Println(blockNumber.String()) //

	blockNumber = api.NewInt(15)
	block, err := client.BlockByNumber(api.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Int64())     //
	fmt.Println(block.Time().Int64())       //
	fmt.Println(block.Difficulty().Int64()) //
	fmt.Println(block.Hash().Hex())         //
	fmt.Println(block.Transactions().Len()) //

	count, err := client.TransactionCount(api.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) //
}
