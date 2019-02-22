package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. get private key
	priv, err := geth.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privBytes := geth.FromECDSA(priv)
	fmt.Println(geth.Encode(privBytes)[2:])

	// 2. get public key from private key
	pub := priv.Public()
	pubBytes := geth.FromECDSAPub(pub)
	fmt.Println(geth.Encode(pubBytes)[4:])

	// 3. get address from public key
	address := geth.PubkeyToAddress(pub)
	fmt.Println(address.GetHex())

	// 4. get address from Hex private key
	privKey, err := geth.HexToECDSA("e09ae607ff4fb3320133e73a76d4fc8e5b784663b2f34662fb910f3ff5d8d5ef")
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privKey.Public()
	address = geth.PubkeyToAddress(pubKey)
	if address.GetHex() == "0x15d48078AB8532b8857e0568311fc3792a5562ab" {
		fmt.Println(address.GetHex())
	} else {
		log.Fatal(err)
	}
}
