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

	// 2. load private key
	privateKey, err := web3go.HexToECDSA("6b18145b7b1fe2fc46cb261776ead9024e348a253653f9086209333b68d8e83b")
	if err != nil {
		log.Fatal(err)
	}

	// 3. get public key from private key
	publicKey := privateKey.Public()

	// 4. get address from public key
	fromAddress := web3go.PubkeyToAddress(publicKey)

	// 5. get nonce of sender
	nonce, err := client.GetPendingNonceAt(web3go.NewContext(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 6. set transaction parameters
	value := web3go.NewBigInt(10000000000000000)
	gasLimit := int64(21000)
	gasPrice := web3go.NewBigInt(30000000000)
	toAddress, _ := web3go.NewAddressFromHex("0xf27d3cb712c4594915bc5ab0e2d72960af41a89d")

	// 7. sign transaction
	var data []byte
	tx := web3go.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := web3go.SignTx(tx, web3go.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 8. send transaction
	err = client.SendTransaction(web3go.NewContext(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.GetHash().GetHex())
}
