package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/core/types"
	"github.com/bcl-chain/web3.go/crypto"
	"github.com/bcl-chain/web3.go/ethclient"
	"github.com/bcl-chain/web3.go/gos/context"
	"github.com/bcl-chain/web3.go/gos/math/big"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("6d5284302baab1cf041e3b8baa724bd681163e82688574a757799a16ac1bbdc1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", privateKey)

	publicKey := privateKey.Public()
	//	log.Fatal("error casting public key to ECDSA")

	fromAddress := crypto.PubkeyToAddress(publicKey)
	fmt.Println(fromAddress.Hex())
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := int64(21000)                 //
	gasPrice := big.NewInt(30000000000)      // in wei (30 gwei)

	toAddress := common.HexToAddress("0x0abb28e0270074d5552a66d5dd172fbcb9db4fd7")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
