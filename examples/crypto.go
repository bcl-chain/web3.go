package main

import (
	"log"

	"github.com/bcl-chain/web3.go/crypto"
	"github.com/bcl-chain/web3.go/common"
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
	//log.Println(address.Hex())

	// 4. get address from Hex private key
	priv_key, err :=crypto.HexToECDSA("e09ae607ff4fb3320133e73a76d4fc8e5b784663b2f34662fb910f3ff5d8d5ef")
	if err != nil {
		log.Fatal(err)
	}
	privBytes = crypto.FromECDSA(priv)
	log.Println(privBytes)

	pub_key := priv_key.Public()
	pubBytes = crypto.FromECDSAPub(pub)
	log.Println(pubBytes)

	address_bytes := crypto.PubkeyToAddress(pub_key)
	addr_new := common.BytesToAddress(address_bytes)
	log.Println(addr_new.Hex(), "0x15d48078AB8532b8857e0568311fc3792a5562ab")

}

