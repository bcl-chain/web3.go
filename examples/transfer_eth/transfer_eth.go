package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://119.3.43.136:8203")
	if err != nil {
		log.Fatal(err)
	}

	// // 2. load private key
	// privateKey, err := web3go.HexToECDSA("4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 3. get public key from private key
	// publicKey := privateKey.Public()

	// // 4. get address from public key
	// fromAddress := web3go.PubkeyToAddress(publicKey)

	// // 5. get nonce of sender
	// nonce, err := client.GetPendingNonceAt(web3go.NewContext(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 6. set transaction parameters
	// value := web3go.NewBigInt(10)
	// gasLimit := int64(21000000)
	// gasPrice := web3go.NewBigInt(3000)
	// toAddress, _ := web3go.NewAddressFromHex("0x5BF987a39C38479A8057264aC904D64b20410427")

	// // 7. sign transaction
	// var data []byte
	// tx := web3go.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	// signedTx, err := web3go.SignTx(tx, web3go.NewHomesteadSigner(), privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 8. send transaction
	// err = client.SendTransaction(web3go.NewContext(), signedTx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("tx sent: %s\n", signedTx.GetHash().GetHex())

	// 9. send batch transaction
	var data []byte
	amount := web3go.NewBigInt(150)
    value := web3go.NewBigInt(10)
	pk := "4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41"
	to := "0x5BF987a39C38479A8057264aC904D64b20410427"
	toaddr, _ := web3go.NewAddressFromHex(to)
	
	//tokentx
	tcontract := "0xc4297a75cded5d8e33ec571c79351895c9b362df"
	txopt := web3go.NewTransactOpts(pk)
	nonce, err := client.GetPendingNonceAt(web3go.NewContext(), txopt.GetFrom())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := int64(50000000)
	gasPrice, err := client.SuggestGasPrice(web3go.NewContext())
	if err != nil {
		log.Fatal(err)
	}
	txopt.SetNonce(nonce)
	txopt.SetGasLimit(gasLimit)
	txopt.SetGasPrice(gasPrice)

	//gitx
	txopt_ := web3go.NewTransaction(nonce+1, toaddr, value, gasLimit, gasPrice, data)
 
	TokenTx, err := web3go.TokenTX(txopt, toaddr, amount, tcontract, client)
	GITx, err := web3go.GITx(txopt_ , pk)
	fmt.Println(TokenTx.GetHash().GetHex())
	fmt.Println(GITx.GetHash().GetHex())
	if err!= nil {
		log.Fatal(err)
	}

	// 9. send batch
	txres, err := web3go.SendTxs(txopt, toaddr, amount, tcontract, client, GITx) 
	if err!=nil {
		log.Fatal(err)
	}
	TokenTxres, err := txres.Get(0)
	GITxres, err := txres.Get(1)
    fmt.Println(TokenTxres.GetHash().GetHex())
	fmt.Println(GITxres.GetHash().GetHex())
}
