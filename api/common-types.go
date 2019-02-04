package api

import (
	"github.com/ethereum/go-ethereum/common"
)

/////////// Hash

type Hash struct {
	hash common.Hash
}

func fromHash(hash common.Hash) *Hash {
	return &Hash{
		hash: hash,
	}
}

func toHash(whash *Hash) common.Hash {
	// TODO : handle case where whash is nil
	//        need to find out how to output nil
	//        nil is not accepted as common.Hash
	hash := whash.hash
	return hash
}

func HexToHash(s string) *Hash {
	h := common.HexToHash(s)
	return fromHash(h)
}

func (wh *Hash) Bytes() []byte {
	h := toHash(wh)
	return h.Bytes()
}

func (wh *Hash) Hex() string {
	h := toHash(wh)
	return h.Hex()
}

/////////// Address

type Address struct {
	address common.Address
}

func fromAddress(address common.Address) *Address {
	return &Address{
		address: address,
	}
}

func toAddress(waddress *Address) common.Address {
	// TODO : handle case where waddress is nil
	//        need to find out how to output nil
	//        nil is not accepted as common.Address
	address := waddress.address
	return address
}

func BytesToAddress(b []byte) *Address {
	a := common.BytesToAddress(b)
	return fromAddress(a)
}

func (wa *Address) Bytes() []byte {
	a := toAddress(wa)
	return a.Bytes()
}

func (wa *Address) Hash() *Hash {
	a := toAddress(wa)
	h := a.Hash()
	return fromHash(h)
}

func (wa *Address) Hex() string {
	a := toAddress(wa)
	return a.Hex()
}

func HexToAddress(s string) *Address {
	a := common.HexToAddress(s)
	return fromAddress(a)
}
