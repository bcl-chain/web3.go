package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://119.3.43.136.8203")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get block of given block number
	block, err := client.GetBlockByNumber(web3go.NewContext(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.GetNumber())
	fmt.Println(block.GetGasLimit())
	fmt.Println(block.GetGasUsed())
	fmt.Println(block.GetDifficulty().GetInt64())
	fmt.Println(block.GetTime())
	fmt.Println(block.GetMixDigest().GetHex())
	fmt.Println(block.GetNonce())
	fmt.Println(block.GetCoinbase().GetHex())
	fmt.Println(block.GetRoot().GetHex())
	fmt.Println(block.GetHash().GetHex())
	fmt.Println(block.GetTransactions().Size())

	// 3. get transaction count of the block
	count, err := client.GetTransactionCount(web3go.NewContext(), block.GetHash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
