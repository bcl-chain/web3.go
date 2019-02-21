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

	// 6. set parameters of transaction
	value := api.NewInt(0) // in wei (0 eth)
	gasLimit := int64(2000000)
	gasPrice, err := client.SuggestGasPrice(api.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := api.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

	// 7. set contract address
	tokenAddress := api.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	// 8. set method to call
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := api.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(api.Encode(methodID)) // 0xa9059cbb

	// 9. set parameters of method to call
	paddedAddress := api.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(api.Encode(paddedAddress))
	amount := api.InitInt()
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := api.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(api.Encode(paddedAmount))

	// 10. prepare data
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 11. sign transaction
	tx := api.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := api.SignTx(tx, api.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 12. send transaction
	err = client.SendTransaction(api.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
