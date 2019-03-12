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
	privateKey, err := web3go.HexToECDSA("3062101dc0f00fd929bfb817d2ea21bf499106abbd7108f6bd33450ab7ec6ece")
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

	// 6. set parameters of transaction
	value := web3go.NewBigInt(0)
	gasLimit := int64(2000000)
	gasPrice, err := client.SuggestGasPrice(web3go.NewContext())
	if err != nil {
		log.Fatal(err)
	}
	toAddress, _ := web3go.NewAddressFromHex("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

	// 7. set contract address
	tokenAddress, _ := web3go.NewAddressFromHex("0x28b149020d2152179873ec60bed6bf7cd705775d")

	// 8. set method to call
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := web3go.New256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(web3go.Encode(methodID)) // 0xa9059cbb

	// 9. set parameters of method to call
	paddedAddress := web3go.LeftPadBytes(toAddress.GetBytes(), 32)
	fmt.Println(web3go.Encode(paddedAddress))
	amount := web3go.NewBigInt(0)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := web3go.LeftPadBytes(amount.GetBytes(), 32)
	fmt.Println(web3go.Encode(paddedAmount))

	// 10. prepare data
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 11. sign transaction
	tx := web3go.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := web3go.SignTx(tx, web3go.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 12. send transaction
	err = client.SendTransaction(web3go.NewContext(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.GetHash().GetHex())
}
