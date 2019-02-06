package main

import (
	"fmt"

	"github.com/bcl-chain/web3.go/api"
)

func main() {
	// 1. set bcl address
	address := api.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
