package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. load private key
	privateKey, err := geth.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	// 2. prepare data to be signed
	data := []byte("hello")
	hash := geth.Keccak256Hash(data)
	fmt.Println(hash.GetHex())

	// 3. sign data with private key
	signature, err := geth.Sign(hash.GetBytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(geth.Encode(signature))
}
