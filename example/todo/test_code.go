package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// connect to client
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// get latest block
	block, err := client.BlockByNumber(api.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// get block of given block hash
	block2, err := client.BlockByHash(api.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// check if block matches block2
	if blocksAreEqual(block, block2) {
		log.Fatal("wrong result, blockes do not match")
	}

	// get header of given block number
	header, err := client.HeaderByNumber(api.Background(), block.Number())
	if err != nil {
		log.Fatal(err)
	}

	// get header of given block hash
	header2, err := client.HeaderByHash(api.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// check if header matches header2
	if headersAreEqual(header, header2) {
		fmt.Println(header.Number(), header2.Number())
		fmt.Printf("%#v %#v\n", header, header2)
		log.Fatal("wrong result, headers do not match")
	}

	// 4. get transaction count of the block
	count, err := client.TransactionCount(api.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func blocksAreEqual(b1, b2 *api.Block) bool {
	if !headersAreEqual(b1.Header(), b2.Header()) {
		return false
	}

	if !transactionsAreEqual(b1.Transactions(), b2.Transactions()) {
		return false
	}

	//	if b1.Number() != b2.Number() {
	//		return false
	//	}
	//
	//	if b1.GasLimit() != b2.GasLimit() {
	//		return false
	//	}
	//
	//	if b1.GasUsed() != b2.GasUsed() {
	//		return false
	//	}
	//
	//	if b1.Difficulty().Int64() != b2.Difficulty().Int64() {
	//		return false
	//	}
	//
	//	if b1.Time().Int64() != b2.Time().Int64() {
	//		return false
	//	}
	//
	//	if b1.MixDigest().Hex() != b2.MixDigest().Hex() {
	//		return false
	//	}
	//
	//	if b1.Nonce() != b2.Nonce() {
	//		return false
	//	}
	//
	//	if b1.Coinbase().Hex() != b2.Coinbase().Hex() {
	//		return false
	//	}
	//
	//	if b1.Root().Hex() != b2.Root().Hex() {
	//		return false
	//	}
	//
	//	if b1.ParentHash().Hex() != b2.ParentHash().Hex() {
	//		return false
	//	}
	//
	//	if b1.TxHash().Hex() != b2.TxHash().Hex() {
	//		return false
	//	}
	//
	//	if b1.ReceiptHash().Hex() != b2.ReceiptHash().Hex() {
	//		return false
	//	}
	//
	//	if b1.UncleHash().Hex() != b2.UncleHash().Hex() {
	//		return false
	//	}
	//
	//	if !reflect.DeepEqual(b1.Extra(), b2.Extra()) {
	//		return false
	//	}

	return true
}

func transactionsAreEqual(txs1, txs2 *api.Transactions) bool {
	if txs1.Len() != txs2.Len() {
		return false
	}

	for i := 0; i < txs1.Len(); i++ {
		tx1 := txs1.Get(i)
		tx2 := txs2.Get(i)
		if tx1.Hash().Hex() != tx2.Hash().Hex() {
			return false
		}
	}

	return true
}

func headersAreEqual(h1, h2 *api.Header) bool {
	if h1.Number() != h2.Number() {
		return false
	}

	if h1.GasLimit() != h2.GasLimit() {
		return false
	}

	if h1.GasUsed() != h2.GasUsed() {
		return false
	}

	if h1.Difficulty().Int64() != h2.Difficulty().Int64() {
		return false
	}

	if h1.Time().Int64() != h2.Time().Int64() {
		return false
	}

	if h1.MixDigest().Hex() != h2.MixDigest().Hex() {
		return false
	}

	if h1.Nonce() != h2.Nonce() {
		return false
	}

	if h1.Coinbase().Hex() != h2.Coinbase().Hex() {
		return false
	}

	if h1.Root().Hex() != h2.Root().Hex() {
		return false
	}

	if h1.ParentHash().Hex() != h2.ParentHash().Hex() {
		return false
	}

	if h1.TxHash().Hex() != h2.TxHash().Hex() {
		return false
	}

	if h1.ReceiptHash().Hex() != h2.ReceiptHash().Hex() {
		return false
	}

	if h1.UncleHash().Hex() != h2.UncleHash().Hex() {
		return false
	}

	if !reflect.DeepEqual(h1.Extra(), h2.Extra()) {
		return false
	}

	return true
}
