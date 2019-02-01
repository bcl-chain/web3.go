package main

import (
	"log"
	"fmt"

	"github.com/bcl-chain/web3.go/common/hexutil"
	"github.com/bcl-chain/web3.go/crypto"
)

func main() {
	// 1. get private key
	priv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privBytes := crypto.FromECDSA(priv)
	fmt.Println(hexutil.Encode(privBytes)[2:])

	// 2. get public key from private key
	pub := priv.Public()
	pubBytes := crypto.FromECDSAPub(pub)
	fmt.Println(hexutil.Encode(pubBytes)[4:])

	// 3. get address from public key
	address := crypto.PubkeyToAddress(pub)
//	fmt.Println(address)
	fmt.Println(address.Hex())
}
