package web3go

import (
	"log"
	// "context"
	// "fmt"
	// "math/rand"
	// "net"
	// "net/http"
	// "net/http/httptest"
	// "os"
	// "reflect"
	// "runtime"
	// "sync"
	// "testing"
	// "time"

	"github.com/ethereum/go-ethereum/core/types"
)

// func BatchTxs(client *EthereumClient, pk string, to string, contract string, amount *BigInt, value *BigInt, gasLimit int64, gasPrice *BigInt)(*Transactions, error){
// 	txs := make(types.Transactions, 0, 2)

// 	privateKey, err := HexToECDSA(pk)
// 	publicKey := privateKey.Public()
// 	fromAddress := PubkeyToAddress(publicKey)
// 	nonce, err := client.GetPendingNonceAt(NewContext(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	toAddress, err := NewAddressFromHex(to)
// 	var data []byte

// 	//token交易
//     //token txs
// 	address, _ := NewAddressFromHex(contract)
// 	erc20, err := NewERC20(address, client)
// 	pk_, err := crypto.HexToECDSA(pk)
// 	tokentxopt := &TransactOpts{bind.NewKeyedTransactor(pk_)}
// 	tokentxopt.SetNonce(nonce)
// 	tokentxopt.SetGasLimit(gasLimit)
// 	tokentxopt.SetGasPrice(gasPrice)
// 	signedTokenTx, err :=  erc20.Transfer(tokentxopt, address, amount)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	txs = append(txs, signedTokenTx)

// 	//手续费交易
// 	//GI txs
// 	nonce = nonce + 1
// 	data = []byte(signedTokenTx.GetHash().GetHex())
// 	gastxopt := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
// 	signedGasTx, err := types.SignTx(gastxopt, NewHomesteadSigner(), privateKey)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	err = client.SendTransaction(NewContext(), signedGasTx)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	txs = append(txs, gastxopt)

// 	return &Transactions{txs: txs}, nil
// }

// func BuildBatchTxsUserToPlatform(*Transactions)(error){
// 	client, err := NewEthereumClient("http://119.3.43.136:8203")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	privateKey, err := HexToECDSA(pk)
// 	publicKey := privateKey.Public()
// 	fromAddress := PubkeyToAddress(publicKey)
// 	nonce, err := client.GetPendingNonceAt(NewContext(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	toAddress, err := NewAddressFromHex(to)
// 	var data []byte

// 	//token交易
// 	//token txs
// 	address, _ := NewAddressFromHex(contract)
// 	erc20, err := NewERC20(address, client)
// 	pk_, err := crypto.HexToECDSA(pk)
// 	tokentxopt := &TransactOpts{bind.NewKeyedTransactor(pk_)}
// 	tokentxopt.SetNonce(nonce)
// 	tokentxopt.SetGasLimit(gasLimit)
// 	tokentxopt.SetGasPrice(gasPrice)
// 	signedTokenTx, err :=  erc20.BuildTransfer(tokentxopt, address, amount)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, nil, err
// 	}

// 	//手续费交易
// 	//GI txs
// 	nonce = nonce + 1
// 	data = []byte(signedTokenTx.GetHash().GetHex())
// 	gastxopt := NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
// 	signedGasTx, err := SignTx(gastxopt, NewHomesteadSigner(), privateKey)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, nil, err
// 	}

// 	return signedTokenTx, signedGasTx, nil
// }

func TokenTX(tokentxopt *TransactOpts, toaddr *Address, iamount string, decimals int, contract string, client *EthereumClient, chainId int64) (*Transaction, error) {
	address, _ := NewAddressFromHex(contract)
	erc20, _ := NewERC20V2(address, client)
	signedTokenTx, err := erc20.BuildTransferV2(tokentxopt, toaddr, iamount, decimals, chainId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return signedTokenTx, nil
}

func GITx(gitxopt *Transaction, privateKey string) (*Transaction, error) {
	privateKey_, err := HexToECDSA(privateKey)
	signedGITx, err := SignTx(gitxopt, NewHomesteadSigner(), privateKey_)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return signedGITx, nil
}

//step one: invoke rpc function to send the two transaction
//step two: invoke batch rpc function to send the packed two function
func SendTxs(tokentxopt *TransactOpts, toaddr *Address, iamount string, decimals int, chainId int64, contract string, client *EthereumClient, gitxopt *Transaction) (*Transactions, error) {
	txs := make(types.Transactions, 0, 2)

	address, _ := NewAddressFromHex(contract)
	erc20, _ := NewERC20V2(address, client)
	signedTokenTx, err := erc20.BuildTransferV2(tokentxopt, toaddr, iamount, decimals, chainId)
	if err != nil {
		log.Fatal(err)
		return &Transactions{txs: txs}, err
	}
	err = erc20.SendTransferV2(signedTokenTx)
	if err != nil {
		log.Fatal(err)
		return &Transactions{txs: txs}, err
	}

	//txs.txs[0]=signedTokenTx
	txs = append(txs, signedTokenTx.tx)

	err = erc20.SendTransferV2(gitxopt)
	if err != nil {
		log.Fatal(err)
		return &Transactions{txs: txs}, err
	}

	//txs.txs[1]=gitxopt
	txs = append(txs, gitxopt.tx)

	return &Transactions{txs: txs}, nil
}
