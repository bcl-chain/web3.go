package main

import (
	"fmt"
	"github.com/bcl-chain/web3.go/api"
	"log"
)

func main() {
	// 1. prepare private key
	privateKey, err := api.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	// 2. prepaer data to be signed
	data := []byte("hello")
	hash := api.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	// 3. sign data with private key
	signature, err := api.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(api.Encode(signature))
}