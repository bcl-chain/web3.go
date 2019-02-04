package main

import (
	"bytes"
	"fmt"
	"github.com/bcl-chain/web3.go/api"
	"log"
)

func main() {
	privateKey, err := api.HexToECDSA("3062101dc0f00fd929bfb817d2ea21bf499106abbd7108f6bd33450ab7ec6ece")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyBytes := api.FromECDSAPub(publicKey)

	data := []byte("hello")
	hash := api.Keccak256Hash(data)
	fmt.Println(hash.Hex()) //

	signature, err := api.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(api.Encode(signature)) //
	sigPublicKey, err := api.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := api.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := api.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := api.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}
