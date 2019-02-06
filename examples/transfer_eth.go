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

	// 2. prepare private key
	privateKey, err := api.HexToECDSA("3062101dc0f00fd929bfb817d2ea21bf499106abbd7108f6bd33450ab7ec6ece")
	if err != nil {
		log.Fatal(err)
	}

	// 3. get public key from private key
	publicKey := privateKey.Public()

	// 4. get address from public key
	fromAddress := api.PubkeyToAddress(publicKey)

	// 5. get nonce of sender
	nonce, err := client.PendingNonceAt(api.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 6. set transaction parameters
	value := api.NewInt(10000000000000000) // in wei (1 eth)
	gasLimit := int64(21000)               //
	gasPrice := api.NewInt(30000000000)    // in wei (30 gwei)
	toAddress := api.HexToAddress("0x0abb28e0270074d5552a66d5dd172fbcb9db4fd7")

	// 7. sign transaction
	var data []byte
	tx := api.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := api.SignTx(tx, api.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 8. send transaction
	err = client.SendTransaction(api.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
