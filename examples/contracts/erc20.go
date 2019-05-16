package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://192.168.2.89:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	address, _ := web3go.NewAddressFromHex("0x501b63b36d109427e9e5e2c59c83b55e4275133c")
	erc20, err := web3go.NewERC20(address, client)
	if err != nil {
		log.Fatal(err)
	}
	address, _ = web3go.NewAddressFromHex("0x3776faee65fde11eda3acdf213da539805f83aa4")
	balance, err := erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())

	opts := web3go.NewTransactOpts("943c44a3f2bbc4230ef40556cd9a6ad46f2e4ce5f2e04f17573a168b8356f3f0")
	nonce, err := client.GetPendingNonceAt(web3go.NewContext(), opts.GetFrom())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(web3go.NewContext())
	if err != nil {
		log.Fatal(err)
	}
	opts.SetNonce(nonce)
	opts.SetGasLimit(300000)
	opts.SetGasPrice(gasPrice)

	address, _ = web3go.NewAddressFromHex("0xc9fa4b1a25397dba5be5db658fcf2d191b253030")
	erc20.Transfer(opts, address, web3go.NewBigInt(10))

	address, _ = web3go.NewAddressFromHex("0x3776faee65fde11eda3acdf213da539805f83aa4")
	balance, err = erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())
}
