package wrapper

import (
	"github.com/ethereum/go-ethereum/common"

	wcommon "github.com/bcl-chain/web3.go/common"
)

func FromHash(hash common.Hash) *wcommon.Hash {
	return &wcommon.Hash{
		Hash: hash,
	}
}

func ToHash(whash *wcommon.Hash) common.Hash {
	// TODO : handle case where whash is nil
	//        need to find out how to output nil
	//        nil is not accepted as common.Hash
	hash, _ := whash.Hash.(common.Hash)
	return hash
}

func FromAddress(address common.Address) *wcommon.Address {
	return &wcommon.Address{
		Address: address,
	}
}

func ToAddress(waddress *wcommon.Address) common.Address {
	// TODO : handle case where waddress is nil
	//        need to find out how to output nil
	//        nil is not accepted as common.Address
	address, _ := waddress.Address.(common.Address)
	return address
}
