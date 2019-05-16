package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://192.168.2.89:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	erc20, err := web3go.NewERC20(common.HexToAddress("0xea26b3c7a6eb0a9b37cc145cb96f1409ef0023f2"), client)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := erc20.BalanceOf(common.HexToAddress("0x0e4ae5082f4307022c598121b81ad39f4554170e"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())

	//	// 2. get networkID
	//	id, err := web3go.NetworkID(api.NewContext())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("networkId:", id)
	//
	//	// 3. close connection to client
	//	geth.Close()
	//	fmt.Println("connection is closed")
}
