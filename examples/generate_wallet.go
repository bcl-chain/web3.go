package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/api"
)

func main() {
	// 1. get private key
	priv, err := api.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privBytes := api.FromECDSA(priv)
	fmt.Println(api.Encode(privBytes)[2:])

	// 2. get public key from private key
	pub := priv.Public()
	pubBytes := api.FromECDSAPub(pub)
	fmt.Println(api.Encode(pubBytes)[4:])

	// 3. get address from public key
	address := api.PubkeyToAddress(pub)
	//	fmt.Println(address)
	fmt.Println(address.Hex())
}
