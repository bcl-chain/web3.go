package main

import (
	"bytes"
	"fmt"
	"github.com/bcl-chain/web3.go/api"
	"log"
)

func main() {
	// 1. prepare private key
	privateKey, err := api.HexToECDSA("3062101dc0f00fd929bfb817d2ea21bf499106abbd7108f6bd33450ab7ec6ece")
	if err != nil {
		log.Fatal(err)
	}

	// 2. get public key form private key
	publicKey := privateKey.Public()
	publicKeyBytes := api.FromECDSAPub(publicKey)

	// 3. prepaer data to be signed
	data := []byte("hello")
	hash := api.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	// 4. sign data with private key
	signature, err := api.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(api.Encode(signature))

	// 5. recover public key in []byte type from signature
	sigPublicKey, err := api.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	// 6. check if recovered public key matches original public key
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches)

	// 7. recover public key in ECDSA type from signature
	sigPublicKeyECDSA, err := api.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	// 8. check if recovered public key matches original public key
	sigPublicKeyBytes := api.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	// 9. check if signature is valid
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := api.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified)
}
