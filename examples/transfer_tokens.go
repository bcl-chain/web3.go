package main

import (
	"github.com/bcl-chain/web3.go/goos/context"
	"fmt"
	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/common/hexutil"
	"github.com/bcl-chain/web3.go/core/types"
	"github.com/bcl-chain/web3.go/crypto"
	"github.com/bcl-chain/web3.go/crypto/sha3"
	"github.com/bcl-chain/web3.go/ethclient"
	"log"
	"github.com/bcl-chain/web3.go/goos/math/big"
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

	value := big.NewInt(0)      // in wei (0 eth)
	gasLimit := int64(2000000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	fmt.Println(nonce, hash.Size())
	fmt.Printf("%#v\n", hash)
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) //
	amount := big.InitInt()
	fmt.Printf("%#v\n", amount)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) //

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent:
}
