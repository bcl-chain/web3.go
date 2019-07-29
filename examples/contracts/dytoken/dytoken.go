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

	//实例化合约
	address, _ := web3go.NewAddressFromHex("0x864eD3ACe69Fd8a5ca11819048e9b2F6f3a7310B")
	dytoken, err := web3go.NewDyToken(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//增发权限功能测试,发布token的地址拥有增发权利
	address, _ = web3go.NewAddressFromHex("0x659f7Cef90f2Cd2d7ea186789338d53E5e2cCc86")
	isminter, _ := dytoken.IsMinter(address)
	if isminter {
		fmt.Println("有增发权限")
	} else {
		fmt.Println("无增发权限")
	}

	//增发功能测试
	fmt.Println("增发功能测试")
	txopt := web3go.NewTransactOpts("4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41")
	nonce, err := client.GetPendingNonceAt(web3go.NewContext(), txopt.GetFrom())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(web3go.NewContext())
	if err != nil {
		log.Fatal(err)
	}
	txopt.SetNonce(nonce)
	txopt.SetGasLimit(500000)
	txopt.SetGasPrice(gasPrice)

	receiverAddr, _ := web3go.NewAddressFromHex("0xf80af2b63ee5e02a53e28c3ed470d8dd54a068d6")
	minttx, _ := dytoken.BuildMint(txopt, receiverAddr, "5", 18, 8888)
	fmt.Println(minttx.GetHash().GetHex())
	minttx, _ = dytoken.Mint(txopt, receiverAddr, "5", 18)
	fmt.Println(minttx.GetHash().GetHex())

	address, _ = web3go.NewAddressFromHex("0xf80af2b63ee5e02a53e28c3ed470d8dd54a068d6")
	balance, err := dytoken.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("增发余额", balance.String())

	// //销毁功能测试
	// //burn test
	// fmt.Println("销毁功能测试")
	// btxopt := web3go.NewTransactOpts("4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41")
	// nonce, err = client.GetPendingNonceAt(web3go.NewContext(), btxopt.GetFrom())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// gasPrice, err = client.SuggestGasPrice(web3go.NewContext())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// btxopt.SetNonce(nonce)
	// btxopt.SetGasLimit(500000)
	// btxopt.SetGasPrice(gasPrice)
	//
	// burntx, _ := dytoken.BuildBurn(btxopt,web3go.NewBigInt(10))
	// fmt.Println(burntx.GetHash().GetHex())
	// burntx, _ = dytoken.Burn(btxopt,web3go.NewBigInt(10))
	// fmt.Println(burntx.GetHash().GetHex())
	//
	// address, _ = web3go.NewAddressFromHex("0x659f7Cef90f2Cd2d7ea186789338d53E5e2cCc86")
	// balance, err = dytoken.BalanceOf(address)
	// if err != nil {
	//   log.Fatal(err)
	// }
	// fmt.Println("销毁余额",balance.String())

	// //转移功能测试
	// //transferOwnship test
	// fmt.Println("转移权限功能测试")
	// ownershipReceiverAddr, _ := web3go.NewAddressFromHex("0x5BF987a39C38479A8057264aC904D64b20410427")
	// isminter, _ := dytoken.IsMinter(ownershipReceiverAddr)
	// if isminter {
	//   fmt.Println("有增发权限")
	// }else{
	//   fmt.Println("无增发权限")
	// }
	//
	// ttxopt := web3go.NewTransactOpts("4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41")
	// nonce, err = client.GetPendingNonceAt(web3go.NewContext(), ttxopt.GetFrom())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// gasPrice, err = client.SuggestGasPrice(web3go.NewContext())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// ttxopt.SetNonce(nonce)
	// ttxopt.SetGasLimit(500000)
	// ttxopt.SetGasPrice(gasPrice)
	//
	// transfertx, _ := dytoken.BuildTransferGRCOwnership(ttxopt, ownershipReceiverAddr)
	// fmt.Println(transfertx.GetHash().GetHex())
	// transfertx, _ = dytoken.TransferGRCOwnership(ttxopt, ownershipReceiverAddr)
	// fmt.Println(transfertx.GetHash().GetHex())
	// isminter, _ = dytoken.IsMinter(ownershipReceiverAddr)
	// if isminter {
	//   fmt.Println("有增发权限")
	// }else{
	//   fmt.Println("无增发权限")
	// }

	// //禁用功能测试
	// //pause test
	// fmt.Println("禁用功能测试")
	// ptxopt := web3go.NewTransactOpts("4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41")
	// nonce, err = client.GetPendingNonceAt(web3go.NewContext(), ptxopt.GetFrom())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// gasPrice, err = client.SuggestGasPrice(web3go.NewContext())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// ptxopt.SetNonce(nonce)
	// ptxopt.SetGasLimit(500000)
	// ptxopt.SetGasPrice(gasPrice)
	//
	// ptx, err := dytoken.BuildPause(ptxopt)
	// fmt.Println(ptx.GetHash().GetHex())
	// ptx, err = dytoken.Pause(ptxopt)
	// fmt.Println(ptx.GetHash().GetHex())
	//
	// ispaused,err := dytoken.Paused()
	// if err != nil{
	//   log.Fatal(err)
	// }
	// if ispaused {
	//     fmt.Println("禁用状态")
	// }else{
	//     fmt.Println("启用状态")
	// }
	//
	// addr, _ := web3go.NewAddressFromHex("0x659f7Cef90f2Cd2d7ea186789338d53E5e2cCc86")
	// ispaused,err = dytoken.IsPauser(addr)
	// if err != nil{
	//   log.Fatal(err)
	// }
	// if ispaused {
	//     fmt.Println("有禁用权限")
	// }else{
	//     fmt.Println("无禁用权限")
	// }

	// //启用功能测试
	// //unpause test
	// fmt.Println("启用功能测试")
	// unptxopt := web3go.NewTransactOpts("4B0D53040702FE1A72DD3B370B58903BBA2476C6920FA2D4A8197369B0B1FC41")
	// nonce, err = client.GetPendingNonceAt(web3go.NewContext(), unptxopt.GetFrom())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// gasPrice, err = client.SuggestGasPrice(web3go.NewContext())
	// if err != nil {
	//   log.Fatal(err)
	// }
	// unptxopt.SetNonce(nonce)
	// unptxopt.SetGasLimit(500000)
	// unptxopt.SetGasPrice(gasPrice)
	// unptx, _ := dytoken.BuildUnPause(unptxopt)
	// fmt.Println(unptx.GetHash().GetHex())
	// unptx, _ = dytoken.UnPause(unptxopt)
	// fmt.Println(unptx.GetHash().GetHex())
	//
	// ispaused,err := dytoken.Paused()
	// if err != nil{
	//   log.Fatal(err)
	// }
	// if ispaused {
	//     fmt.Println("禁用状态")
	// }else{
	//     fmt.Println("启用状态")
	// }

}
