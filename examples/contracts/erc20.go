package main

import (
	"fmt"
	"log"

	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	// 1. connect to client
	client, err := web3go.NewEthereumClient("http://192.168.1.15:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	address, _ := web3go.NewAddressFromHex("0xc558ed48e80b27cf8c57ebe5267997ceac5ea2c1")
	erc20, err := web3go.NewERC20(address, client)
	if err != nil {
		log.Fatal(err)
	}
	address, _ = web3go.NewAddressFromHex("0xfe04cb1d7d6715169edc07c8e3c2fdba3a0854af")
	balance, err := erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())

	opts := web3go.NewTransactOpts("ac5c70f5a1c9f5b121edbbabb054d8763ac1f5066aeacbb14a5e4c10216fb605")
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

	address, _ = web3go.NewAddressFromHex("0xf80af2b63ee5e02a53e28c3ed470d8dd54a068d6")
	tx, _ := erc20.BuildTransfer(opts, address, web3go.NewBigInt(10))
	fmt.Println(tx.GetHash().GetHex())
	tx, _ = erc20.Transfer(opts, address, web3go.NewBigInt(10))
	fmt.Println(tx.GetHash().GetHex())

	address, _ = web3go.NewAddressFromHex("0xfe04cb1d7d6715169edc07c8e3c2fdba3a0854af")
	balance, err = erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())
}
