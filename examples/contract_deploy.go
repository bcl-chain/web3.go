package main

import (
//	store "./contracts" // for demo
//	"fmt"
	"github.com/bcl-chain/web3.go/accounts/abi/bind"
	"github.com/bcl-chain/web3.go/crypto"
	"github.com/bcl-chain/web3.go/ethclient"
	"github.com/bcl-chain/web3.go/goos/context"
	"github.com/bcl-chain/web3.go/goos/math/big"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	//	log.Fatal("error casting public key to ECDSA")

	fromAddress := crypto.PubkeyToAddress(publicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.SetNonce(big.NewInt(int64(nonce)))
	auth.SetValue(big.NewInt(0))    // in wei
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
