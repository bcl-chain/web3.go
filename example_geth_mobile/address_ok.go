package main

import (
	"fmt"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. set bcl address
	address, _ := geth.NewAddressFromHex("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Println(address.GetHex())
	// fmt.Println(address.GetHash().GetHex())
	fmt.Println(address.GetBytes())
}
