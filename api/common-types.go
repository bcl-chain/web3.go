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

func BytesToHash(b []byte) *Hash {
	h := common.BytesToHash(b)
	return fromHash(h)
}

func BigToHash(wx *Int) *Hash {
	x := toBigInt(wx)
	h := common.BigToHash(x)
	return fromHash(h)
}

func HexToHash(s string) *Hash {
	h := common.HexToHash(s)
	return fromHash(h)
}

func (wh *Hash) Bytes() []byte {
	h := toHash(wh)
	return h.Bytes()
}

func (wh *Hash) Big() *Int {
	h := toHash(wh)
	x := h.Big()
	return fromBigInt(x)
}

func (wh *Hash) Hex() string {
	h := toHash(wh)
	return h.Hex()
}

func (wh *Hash) TerminalString() string {
	h := toHash(wh)
	return h.TerminalString()
}

func (wh *Hash) String() string {
	h := toHash(wh)
	return h.String()
}

// TODO
//func Format {
//}

func (wh *Hash) UnmarshalText(input []byte) error {
	h := toHash(wh)
	return h.UnmarshalText(input)
}

func (wh *Hash) UnmarshalJSON(input []byte) error {
	h := toHash(wh)
	return h.UnmarshalJSON(input)
}

func (wh *Hash) MarshalText() ([]byte, error) {
	h := toHash(wh)
	return h.MarshalText()
}

func (wh *Hash) SetBytes(b []byte) {
	h := toHash(wh)
	h.SetBytes(b)
}

//TODO
//func Generate {
//}

//TODO
//func Scan {
//}

//TODO
//func Value {
//}

//TODO
//func UnmarshalText {
//}

//TODO
//func MarshalText {
//}

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

func BigToAddress(wx *Int) *Address {
	x := toBigInt(wx)
	a := common.BigToAddress(x)
	return fromAddress(a)
}

func HexToAddress(s string) *Address {
	a := common.HexToAddress(s)
	return fromAddress(a)
}

func IsHexAddress(s string) bool {
	return common.IsHexAddress(s)
}

func (wa *Address) Bytes() []byte {
	a := toAddress(wa)
	return a.Bytes()
}

func (wa *Address) Big() *Int {
	a := toAddress(wa)
	x := a.Big()
	return fromBigInt(x)
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

func (wa *Address) String() string {
	a := toAddress(wa)
	return a.String()
}

// TODO
//func Format {
//}

func (wa *Address) SetBytes(b []byte) {
	a := toAddress(wa)
	a.SetBytes(b)
}

func (wa *Address) MarshalText() ([]byte, error) {
	a := toAddress(wa)
	return a.MarshalText()
}

func (wa *Address) UnmarshalText(input []byte) error {
	a := toAddress(wa)
	return a.UnmarshalText(input)
}

func (wa *Address) UnmarshalJSON(input []byte) error {
	a := toAddress(wa)
	return a.UnmarshalJSON(input)
}

//TODO
//func Scan {
//}

//TODO
//func Value {
//}

//TODO
//func UnmarshalText {
//}

//TODO
//func MarshalText {
//}

//TODO
//func NewMixedcaseAddress {
//}

//TODO
//func NewMixedcaseAddressFromString {
//}

//TODO
//func UnmarshalJSON {
//}

//TODO
//func MarshalJSON {
//}

//TODO
//func Address {
//}

//TODO
//func String {
//}

//TODO
//func ValidateCheckSum {
//}

//TODO
//func Original {
//}
