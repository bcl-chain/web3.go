package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. prepare private key
	privateKey, err := web3go.HexToECDSA("3062101dc0f00fd929bfb817d2ea21bf499106abbd7108f6bd33450ab7ec6ece")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get public key form private key
	publicKey := privateKey.Public()
	publicKeyBytes := web3go.FromECDSAPub(publicKey)

	// 3. prepaer data to be signed
	data := []byte("hello")
	hash := web3go.Keccak256Hash(data)
	fmt.Println(hash.GetHex())

	// 4. sign data with private key
	signature, err := web3go.Sign(hash.GetBytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(web3go.Encode(signature))

	// 5. recover public key in []byte type from signature
	sigPublicKey, err := web3go.Ecrecover(hash.GetBytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	// 6. check if recovered public key matches original public key
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches)

	// 7. recover public key in ECDSA type from signature
	sigPublicKeyECDSA, err := web3go.SigToPub(hash.GetBytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	// 8. check if recovered public key matches original public key
	sigPublicKeyBytes := web3go.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	// 9. check if signature is valid
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := web3go.VerifySignature(publicKeyBytes, hash.GetBytes(), signatureNoRecoverID)
	fmt.Println(verified)
}
