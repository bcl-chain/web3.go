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

	blockNumber := api.NewInt(15)
	block, err := client.BlockByNumber(api.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: cannnot range over wrapped Transactions
	for i := 0; i < block.Transactions().Len(); i++ {
		tx := block.Transactions().Get(i)
		fmt.Println(tx.Hash().Hex())       //
		fmt.Println(tx.Value().String())   //
		fmt.Println(tx.Gas())              //
		fmt.Println(tx.GasPrice().Int64()) //
		fmt.Println(tx.Nonce())            //
		fmt.Println(tx.Data())             //
		fmt.Println(tx.To().Hex())         //

		if msg, err := tx.AsMessage(api.NewHomesteadSigner()); err != nil {
			fmt.Println(msg.From().Hex()) //
		}

		receipt, err := client.TransactionReceipt(api.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status()) //
	}

	blockHash := api.HexToHash("0xe1bcee950d7acd1de033e4d143e8683d15ed3aeb6d0814d579024424f4eac009")
	count, err := client.TransactionCount(api.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := int64(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(api.Background(), blockHash, idx)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex()) //
	}

	txHash := api.HexToHash("0x3195e6a4cd745254fafda021fcfde00384d1b379e523311b0750dd4135a2e516")
	tx, err := client.TransactionByHash(api.Background(), txHash)
	isPending, err := client.TransactionByHashIsPending(api.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) //
	fmt.Println(isPending)       //
}
