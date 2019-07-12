package main

import (
	"fmt"
	"log"
	"github.com/bcl-chain/web3.go/mobile"
)

func main() {
	//New Ethereum Client
	client, err := web3go.NewEthereumClient("https://kovan.infura.io/v3/def7370cf49d49d791b9df949986b9a0")
	if err != nil {
		log.Fatal(err)
	}

	address, _ := web3go.NewAddressFromHex("0xfe1827f2f1c366c04d458b3c07b8bd207d42eab4")
	// 3. get ether balance from the latest block
	balance, err := client.GetBalanceAt(web3go.NewContext(), address, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	//Access the Token
	tokenAddress, _ := web3go.NewAddressFromHex("0xb19ac159b87a4491b3b8bef4554b59da2bf42555")
	ytoken, err := web3go.NewYToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//Check if the address is a Token owner
	isminter, _ := ytoken.IsMinter(address)
	if isminter {
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}

	tokenBalance, err := ytoken.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("YOCL Token Balance",tokenBalance.String())
}
