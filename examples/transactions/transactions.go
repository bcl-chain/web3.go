package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get block using block number
	blockNumber := int64(10)
	block, err := client.GetBlockByNumber(web3go.NewContext(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 3. get each transaction in the block
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
		if sender, err := client.GetTransactionSender(web3go.NewContext(), tx, block.GetHash(), i); err == nil {
			fmt.Println("sender:", sender.GetHex())
		}

		// 3-3. get receipt of the transaction
		receipt, err := client.GetTransactionReceipt(web3go.NewContext(), tx.GetHash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("status:", receipt.GetStatus())
	}

	// 4. get block using block hash
	blockHash, _ := web3go.NewHashFromHex("0x88fb82c126d2b0d47980faa7b633fbe6a8880eb49ca61eb0767f713808522fc8")
	count, err := client.GetTransactionCount(web3go.NewContext(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// 5. get each transaction
	for idx := int(0); idx < count; idx++ {
		// 5-1. get each transaction
		tx, err := client.GetTransactionInBlock(web3go.NewContext(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hash:", tx.GetHash().GetHex())
	}

	// 6. get transaction using transaction hash
	txHash, _ := web3go.NewHashFromHex("0x9b082a8ac665b0987ab683c0385f5cf2c6d448c2f97c318146232d2e3bbe4a21")
	tx, err := client.GetTransactionByHash(web3go.NewContext(), txHash)
	isPending, err := client.GetTransactionByHashIsPending(web3go.NewContext(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hash:", tx.GetHash().GetHex())
	fmt.Println("isPending:", isPending)
}
