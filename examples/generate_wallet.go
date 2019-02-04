package main

import (
	"log"
	"fmt"

	"github.com/bcl-chain/web3.go/api"
//	"github.com/bcl-chain/web3.go/common/hexutil"
//	"github.com/bcl-chain/web3.go/crypto"
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
