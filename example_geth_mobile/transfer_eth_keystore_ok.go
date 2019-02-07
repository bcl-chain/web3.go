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

	// 2. create new keyStore and load 0th account
	keyStore := geth.NewKeyStore("/home/k-katsushima/Desktop/ethereum/go-etheruem/keystore", geth.StandardScryptN, geth.StandardScryptP)
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
	nonce, err := client.GetPendingNonceAt(geth.NewContext(), address0)
	if err != nil {
		log.Fatal(err)
	}

	// 5. set transaction parameters
	value := geth.NewBigInt(100000000000000000) // in wei (1 eth)
	gasLimit := int64(21000)                    //
	gasPrice := geth.NewBigInt(30000000000)     // in wei (30 gwei)
	toAddress, _ := geth.NewAddressFromHex("0xFBcDe7c69381f5e3e0C02f983da48A65d88DD02a")

	// 6. sign transaction
	var data []byte
	tx := geth.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainId := geth.NewBigInt(15)
	signedTx, err := keyStore.SignTx(account0, tx, chainId)
	if err != nil {
		log.Fatal(err)
	}

	// 7. send transaction
	err = client.SendTransaction(geth.NewContext(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.GetHash().GetHex())
}
