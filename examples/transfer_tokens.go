package main

import (
	"fmt"
	"github.com/bcl-chain/web3.go/api"
	"log"
)

func main() {
	client, err := api.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := api.HexToECDSA("3062101dc0f00fd929bfb817d2ea21bf499106abbd7108f6bd33450ab7ec6ece")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	fromAddress := api.PubkeyToAddress(publicKey)
	nonce, err := client.PendingNonceAt(api.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := api.NewInt(0)     // in wei (0 eth)
	gasLimit := int64(2000000) // in units
	gasPrice, err := client.SuggestGasPrice(api.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := api.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tokenAddress := api.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := api.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(api.Encode(methodID)) // 0xa9059cbb

	paddedAddress := api.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(api.Encode(paddedAddress)) //
	amount := api.InitInt()
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := api.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(api.Encode(paddedAmount)) //

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := api.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := api.SignTx(tx, api.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(api.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent:
}
