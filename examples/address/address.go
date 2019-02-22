package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	address, err := geth.NewAddressFromHex("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.GetHex())
	fmt.Println(address.GetHash().GetHex())
	fmt.Println(address.GetBytes())
}
