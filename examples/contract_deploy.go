package main

import (
//	store "./contracts" // for demo
//	"fmt"
	"github.com/bcl-chain/web3.go/api"
	"log"
)

func main() {
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := api.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	fromAddress := api.PubkeyToAddress(publicKey)
	nonce, err := client.PendingNonceAt(api.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(api.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := api.NewKeyedTransactor(privateKey)
	auth.SetNonce(api.NewInt(int64(nonce)))
	auth.SetValue(api.NewInt(0))    // in wei
	auth.SetGasLimit(int64(300000)) // in units
	auth.SetGasPrice(gasPrice)

	// TODO: need to wrap DeployStore, which is generated by abigen
//	input := "1.0"
//	address, tx, instance, err := store.DeployStore(auth, client, input)
//	if err != nil {
//		log.Fatal(err)
//	}

//	fmt.Println(address.Hex())   //
//	fmt.Println(tx.Hash().Hex()) //

//	_ = instance
}
