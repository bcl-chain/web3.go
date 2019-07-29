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
	address, _ := web3go.NewAddressFromHex("0xC2870160634934c79ac93757E37b72AE17141715")
	erc20, err := web3go.NewERC20V2(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//打印用户token余额
	address, _ = web3go.NewAddressFromHex("0x941c69B23CeF5f5021b5966f4ba85fE6Bf9A58E1")
	balance, err := erc20.BalanceOfV2(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())

	//签名某交易
	opts := web3go.NewTransactOpts("72149D7D1FAD0A66AF39AB1693F28E66E0600A2032D874FF321C8ABEBDC41670")
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
	address, _ = web3go.NewAddressFromHex("0x71ce870012d1B6cd6e94E764e03643048617cA0F")
	tx, _ := erc20.BuildTransferV2(opts, address, "1500", 18, 8888)
	fmt.Println(tx.GetHash().GetHex())
	err = erc20.SendTransferV2(tx)
	if err != nil {
		log.Fatal(err)
	}

	//打印账户token余额
	address, _ = web3go.NewAddressFromHex("0xfe04cb1d7d6715169edc07c8e3c2fdba3a0854af")
	balance, err = erc20.BalanceOfV2(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())
}
