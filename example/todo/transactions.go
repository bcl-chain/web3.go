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

	// 2. get block using block number
	blockNumber := int64(344)
	block, err := client.GetBlockByNumber(geth.NewContext(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 3. get each transaction in the block
	// TODO: cannnot range over wrapped Transactions
	for i := 0; i < block.GetTransactions().Size(); i++ {
		// 3-1. get each transaction
		tx, _ := block.GetTransactions().Get(i)
		fmt.Println("hash:", tx.GetHash().GetHex())
		fmt.Println("value:", tx.GetValue())
		fmt.Println("gas:", tx.GetGas())
		fmt.Println("gasPrice:", tx.GetGasPrice())
		fmt.Println("nonce:", tx.GetNonce())
		fmt.Println("data:", tx.GetData())
		fmt.Println("to:", tx.GetTo().GetHex())

		// 3-2. get sender address
		// if msg, err := tx.AsMessage(api.NewHomesteadSigner()); err != nil {
		// fmt.Println(msg.From().Hex())
		// }

		// 3-2. get sender address
		// if sender, err := client.TransactionSender(geth.NewContext(), tx, block.GetHash(), i); err == nil { //TODO: not implemented in mobile
		// fmt.Println("sender:", sender.Hex())
		// }

		// 3-3. get receipt of the transaction
		receipt, err := client.GetTransactionReceipt(geth.NewContext(), tx.GetHash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("status:", receipt.GetStatus())
	}

	// 4. get block using block hash
	blockHash, _ := geth.NewHashFromHex("0x4198cab5087eb41c3f1ec3db69e3dd828564050c223eaf80678b3155baf178d3")
	count, err := client.GetTransactionCount(geth.NewContext(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// 5. get each transaction
	for idx := int(0); idx < count; idx++ {
		// 5-1. get each transaction
		tx, err := client.GetTransactionInBlock(geth.NewContext(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hash:", tx.GetHash().GetHex())
	}

	// 6. get transaction using transaction hash
	txHash, _ := geth.NewHashFromHex("0xb003c71c7b18a06d01af8494d7aaeb65a34d2db332a06cfe0f353a8c7d920138")
	tx, err := client.GetTransactionByHash(geth.NewContext(), txHash)
	// isPending, err := client.GetTransactionByHashIsPending(geth.NewContext(), txHash) // TODO: not implemented in mobile
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hash:", tx.GetHash().GetHex())
	// fmt.Println("isPending:", isPending)
}
