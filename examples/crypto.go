package main

import (
	"log"

	"github.com/bcl-chain/web3.go/crypto"
)

func main() {
	// 1. get private key
	priv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privBytes := crypto.FromECDSA(priv)
	log.Println(privBytes)

	// 2. get public key from private key
	pub := priv.Public()
	pubBytes := crypto.FromECDSAPub(pub)
	log.Println(pubBytes)

	// 3. get address from public key
	address := crypto.PubkeyToAddress(pub)
	log.Println(address)
}
