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
}
