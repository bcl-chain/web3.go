package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/ethclient"
	"github.com/bcl-chain/web3.go/gos/context"
	"github.com/bcl-chain/web3.go/gos/math/big"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number()
	fmt.Println(blockNumber.String()) //

	blockNumber = big.NewInt(28)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Int64())     //
	fmt.Println(block.Time().Int64())       //
	fmt.Println(block.Difficulty().Int64()) //
	fmt.Println(block.Hash().Hex())         //
	fmt.Println(block.Transactions().Len()) //

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) //
}
