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

	// 2. get latest header
//	header, err := client.GetHeaderByNumber(geth.NewContext(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	blockNumber := header.Number()
//	fmt.Println(blockNumber.String())

	// 3. get header of given block number
	block, err := client.GetBlockByNumber(geth.NewContext(), 15)
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

	// 4. get transaction count of the block
	count, err := client.GetTransactionCount(geth.NewContext(), block.GetHash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
