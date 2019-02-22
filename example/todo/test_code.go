package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// connect to client
	client, err := geth.NewEthereumClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// get latest block
	block, err := client.GetBlockByNumber(geth.NewContext(), -1)
	if err != nil {
		log.Fatal(err)
	}

	// get block of given block hash
	block2, err := client.GetBlockByHash(geth.NewContext(), block.GetHash())
	if err != nil {
		log.Fatal(err)
	}

	// check if block matches block2
	if blocksAreEqual(block, block2) {
		log.Fatal("wrong result, blockes do not match")
	}

	// get header of given block number
	header, err := client.GetHeaderByNumber(geth.NewContext(), block.GetNumber())
	if err != nil {
		log.Fatal(err)
	}

	// get header of given block hash
	header2, err := client.GetHeaderByHash(geth.NewContext(), block.GetHash())
	if err != nil {
		log.Fatal(err)
	}

	// check if header matches header2
	if headersAreEqual(header, header2) {
		fmt.Println(header.GetNumber(), header2.GetNumber())
		fmt.Printf("%#v %#v\n", header, header2)
		log.Fatal("wrong result, headers do not match")
	}

	// 4. get transaction count of the block
	count, err := client.GetTransactionCount(geth.NewContext(), block.GetHash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transaction count", count)

	fmt.Println("test passed")
}

func blocksAreEqual(b1, b2 *geth.Block) bool {
	if !headersAreEqual(b1.GetHeader(), b2.GetHeader()) {
		return false
	}

	if !transactionsAreEqual(b1.GetTransactions(), b2.GetTransactions()) {
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

func transactionsAreEqual(txs1, txs2 *geth.Transactions) bool {
	if txs1.Size() != txs2.Size() {
		return false
	}

	for i := 0; i < txs1.Size(); i++ {
		tx1, _ := txs1.Get(i)
		tx2, _ := txs2.Get(i)
		if tx1.GetHash().GetHex() != tx2.GetHash().GetHex() {
			return false
		}
	}

	return true
}

func headersAreEqual(h1, h2 *geth.Header) bool {
	if h1.GetNumber() != h2.GetNumber() {
		return false
	}

	if h1.GetGasLimit() != h2.GetGasLimit() {
		return false
	}

	if h1.GetGasUsed() != h2.GetGasUsed() {
		return false
	}

	if h1.GetDifficulty().GetInt64() != h2.GetDifficulty().GetInt64() {
		return false
	}

	if h1.GetTime() != h2.GetTime() {
		return false
	}

	if h1.GetMixDigest().GetHex() != h2.GetMixDigest().GetHex() {
		return false
	}

	if h1.GetNonce() != h2.GetNonce() {
		return false
	}

	if h1.GetCoinbase().GetHex() != h2.GetCoinbase().GetHex() {
		return false
	}

	if h1.GetRoot().GetHex() != h2.GetRoot().GetHex() {
		return false
	}

	if h1.GetParentHash().GetHex() != h2.GetParentHash().GetHex() {
		return false
	}

	if h1.GetTxHash().GetHex() != h2.GetTxHash().GetHex() {
		return false
	}

	if h1.GetReceiptHash().GetHex() != h2.GetReceiptHash().GetHex() {
		return false
	}

	if h1.GetUncleHash().GetHex() != h2.GetUncleHash().GetHex() {
		return false
	}

	if !reflect.DeepEqual(h1.GetExtra(), h2.GetExtra()) {
		return false
	}

	return true
}
