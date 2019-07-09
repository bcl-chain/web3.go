package web3go

import (
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func BatchTxsUserToPlatform(pk string, to string, contract string, amount *BigInt, value *BigInt, gasLimit int64, gasPrice *BigInt)(*Transaction, *Transaction, error){
	client, err := NewEthereumClient("http://119.3.43.136:8203")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := HexToECDSA(pk)
	publicKey := privateKey.Public()
	fromAddress := PubkeyToAddress(publicKey)
	nonce, err := client.GetPendingNonceAt(NewContext(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	toAddress, err := NewAddressFromHex(to)
	var data []byte

	//token交易
	address, _ := NewAddressFromHex(contract)
	erc20, err := NewERC20(address, client)
	pk_, err := crypto.HexToECDSA(pk)
	tokentxopt := &TransactOpts{bind.NewKeyedTransactor(pk_)}
	tokentxopt.SetNonce(nonce)
	tokentxopt.SetGasLimit(gasLimit)
	tokentxopt.SetGasPrice(gasPrice)
	signedTokenTx, err :=  erc20.Transfer(tokentxopt, address, amount)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	//手续费交易
	nonce = nonce + 1
	data = []byte(signedTokenTx.GetHash().GetHex())
	gastxopt := NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedGasTx, err := SignTx(gastxopt, NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	err = client.SendTransaction(NewContext(), signedGasTx)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	return signedTokenTx, signedGasTx, nil
}

func BuildBatchTxsUserToPlatform(pk string, to string, contract string, amount *BigInt, value *BigInt, gasLimit int64, gasPrice *BigInt)(*Transaction, *Transaction, error){
	client, err := NewEthereumClient("http://119.3.43.136:8203")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := HexToECDSA(pk)
	publicKey := privateKey.Public()
	fromAddress := PubkeyToAddress(publicKey)
	nonce, err := client.GetPendingNonceAt(NewContext(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	toAddress, err := NewAddressFromHex(to)
	var data []byte

	//token交易
	address, _ := NewAddressFromHex(contract)
	erc20, err := NewERC20(address, client)
	pk_, err := crypto.HexToECDSA(pk)
	tokentxopt := &TransactOpts{bind.NewKeyedTransactor(pk_)}
	tokentxopt.SetNonce(nonce)
	tokentxopt.SetGasLimit(gasLimit)
	tokentxopt.SetGasPrice(gasPrice)
	signedTokenTx, err :=  erc20.BuildTransfer(tokentxopt, address, amount)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	//手续费交易
	nonce = nonce + 1
	data = []byte(signedTokenTx.GetHash().GetHex())
	gastxopt := NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedGasTx, err := SignTx(gastxopt, NewHomesteadSigner(), privateKey)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	err = client.SendTransaction(NewContext(), signedGasTx)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	return signedTokenTx, signedGasTx, nil
}