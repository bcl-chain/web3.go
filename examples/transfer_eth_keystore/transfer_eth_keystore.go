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

	// 2. create new keyStore and load 0th account
	keyStore := web3go.NewKeyStore("/home/k-katsushima/Desktop/ethereum/go-etheruem/keystore", web3go.StandardScryptN, web3go.StandardScryptP)
	accounts := keyStore.GetAccounts()
	if accounts.Size() <= 0 {
		log.Fatal("there is no account")
	}
	account0, _ := accounts.Get(0)
	address0 := account0.GetAddress()
	fmt.Println(address0.GetHex())

	// 3. unlock 0th account
	keyStore.Unlock(account0, "12345678")

	// 4. get nonce of sender
	nonce, err := client.GetPendingNonceAt(web3go.NewContext(), address0)
	if err != nil {
		log.Fatal(err)
	}

	// 5. set transaction parameters
	value := web3go.NewBigInt(100000000000000000)
	gasLimit := int64(21000)
	gasPrice := web3go.NewBigInt(30000000000)
	toAddress, _ := web3go.NewAddressFromHex("0xFBcDe7c69381f5e3e0C02f983da48A65d88DD02a")

	// 6. sign transaction
	var data []byte
	tx := web3go.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID := web3go.NewBigInt(15)
	signedTx, err := keyStore.SignTx(account0, tx, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// 7. send transaction
	err = client.SendTransaction(web3go.NewContext(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.GetHash().GetHex())
}
