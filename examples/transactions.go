package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/core/types"
	"github.com/bcl-chain/web3.go/ethclient"
	"github.com/bcl-chain/web3.go/gos/context"
	"github.com/bcl-chain/web3.go/gos/math/big"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(104)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
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

		if msg, err := tx.AsMessage(types.NewHomesteadSigner()); err != nil {
			fmt.Println(msg.From().Hex()) //
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status()) //
	}

	blockHash := common.HexToHash("0xfb0f445dd139f98095241d83bc1aaf16aaeb68d03a02863bc90383a7318967ee")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := int64(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex()) //
	}

	txHash := common.HexToHash("0x9852fc34f26b791e0877a81d094bb6b5ccea6d5a78302e381690ec393d071124")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) //
	fmt.Println(isPending)       //
}
