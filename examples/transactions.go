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

	// 2. get block using block number
	blockNumber := api.NewInt(27)
	block, err := client.BlockByNumber(api.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 3. get each transaction in the block
	// TODO: cannnot range over wrapped Transactions
	for i := 0; i < block.Transactions().Len(); i++ {
		// 3-1. get each transaction
		tx := block.Transactions().Get(i)
		fmt.Println("hash:", tx.Hash().Hex())
		fmt.Println("value:", tx.Value().String())
		fmt.Println("gas:", tx.Gas())
		fmt.Println("gasPrice:", tx.GasPrice().Int64())
		fmt.Println("nonce:", tx.Nonce())
		fmt.Println("data:", tx.Data())
		fmt.Println("to:", tx.To().Hex())

		// 3-2. get sender address
		// if msg, err := tx.AsMessage(api.NewHomesteadSigner()); err != nil {
		// fmt.Println(msg.From().Hex())
		// }

		// 3-2. get sender address
		if sender, err := client.TransactionSender(api.Background(), tx, block.Hash(), i); err == nil {
			fmt.Println("sender:", sender.Hex())
		}

		// 3-3. get receipt of the transaction
		receipt, err := client.TransactionReceipt(api.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("status:", receipt.Status())
	}

	// 4. get block usng block hash
	blockHash := api.HexToHash("0xf6f66a835fb2abca4a59a4b81f613a5c3824459ea23068b8f6c6777922af78d6")
	count, err := client.TransactionCount(api.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// 5. get each transaction
	for idx := int64(0); idx < count; idx++ {
		// 5-1. get each transaction
		tx, err := client.TransactionInBlock(api.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hash:", tx.Hash().Hex())
	}

	// 6. get transaction using transaction hash
	txHash := api.HexToHash("0x11d24404fe01937fcbbe61e0d0f5807e317439e13ef10cb429d29290fb2f8733")
	tx, err := client.TransactionByHash(api.Background(), txHash)
	isPending, err := client.TransactionByHashIsPending(api.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hash:", tx.Hash().Hex())
	fmt.Println("isPending:", isPending)
}
