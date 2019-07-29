package main

import (
	"fmt"
	"log"

	web3go "github.com/bcl-chain/web3.go/mobile"
)

func main() {
	//连接节点
	client, err := web3go.NewEthereumClient("http://119.3.43.136:8203")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	//实例化合约
	address, _ := web3go.NewAddressFromHex("0xc4297a75cded5d8e33ec571c79351895c9b362df")
	erc20, err := web3go.NewERC20(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//打印用户token余额
	address, _ = web3go.NewAddressFromHex("0xfe04cb1d7d6715169edc07c8e3c2fdba3a0854af")
	balance, err := erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())

	//签名某交易
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

	//从某账户地址发起交易,打印交易签名
	address, _ = web3go.NewAddressFromHex("0xf80af2b63ee5e02a53e28c3ed470d8dd54a068d6")
	tx, _ := erc20.BuildTransfer(opts, address, web3go.NewBigInt(10))
	fmt.Println(tx.GetHash().GetHex())
	tx, _ = erc20.Transfer(opts, address, web3go.NewBigInt(10))
	fmt.Println(tx.GetHash().GetHex())

	//打印账户token余额
	address, _ = web3go.NewAddressFromHex("0xfe04cb1d7d6715169edc07c8e3c2fdba3a0854af")
	balance, err = erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())
}
