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

	// 2. get latest header
	header, err := client.HeaderByNumber(api.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := header.Number()
	fmt.Println(blockNumber.String())

	// 3. get header of given block number
	blockNumber = api.NewInt(15)
	block, err := client.BlockByNumber(api.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number().Int64())
	fmt.Println(block.GasLimit())
	fmt.Println(block.GasUsed())
	fmt.Println(block.Difficulty().Int64())
	fmt.Println(block.Time().Int64())
	fmt.Println(block.MixDigest().Hex())
	fmt.Println(block.Nonce())
	fmt.Println(block.Coinbase().Hex())
	fmt.Println(block.Root().Hex())
	fmt.Println(block.Hash().Hex())
	fmt.Println(block.Transactions().Len())

	// 4. get transaction count of the block
	count, err := client.TransactionCount(api.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
