package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	address, err := web3go.NewAddressFromHex("0xa3bcb71efafe7fdeb13333baaf8f0a06cafc2a11")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.GetHex())
	fmt.Println(address.GetHash().GetHex())
	fmt.Println(address.GetBytes())

	privKey, _ := web3go.HexToECDSA("943c44a3f2bbc4230ef40556cd9a6ad46f2e4ce5f2e04f17573a168b8356f3f0")
	pubKey := privKey.Public()
	address = web3go.PubkeyToAddress(pubKey)
	fmt.Println(address.GetHex())
}
