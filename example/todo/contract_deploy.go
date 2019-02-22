package main

// TODO: this code is not completed

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"

	// store "./contracts" // for demo
	// store "./contracts_deploy" // for demo
)

func main() {
	// 1. connect to cilent
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 2. prepare private key
	privateKey, err := api.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
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

	// 6. set gas price
	gasPrice, err := client.SuggestGasPrice(api.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 7. set parameters for contract deploy
	auth := api.NewKeyedTransactor(privateKey)
	auth.SetNonce(api.NewInt(int64(nonce)))
	auth.SetValue(api.NewInt(0))    // in wei
	auth.SetGasLimit(int64(300000)) // in units
	auth.SetGasPrice(gasPrice)

	// 8. deploy contract
	// TODO: need to wrap DeployStore, which is generated by abigen
	//input := "1.0"
	//address, err := store.DeployStore(auth, client, input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(address.Hex())
	//	fmt.Println(tx.Hash().Hex())

	//	_ = instance
}
