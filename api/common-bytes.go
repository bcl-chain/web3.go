package api

import (
	"github.com/ethereum/go-ethereum/common"
)

func ToHex(b []byte) string {
	return common.ToHex(b)
}

// TODO
//func ToHexArray {
//}

func FromHex(s string) []byte {
	return common.FromHex(s)
}

func CopyBytes(b []byte) []byte {
	return common.CopyBytes(b)
}

func Bytes2Hex(d []byte) string {
	return common.Bytes2Hex(d)
}

func Hex2Bytes(str string) []byte {
	return common.Hex2Bytes(str)
}

func Hex2BytesFixed(str string, flen int) []byte {
	return common.Hex2BytesFixed(str, flen)
}

func RightPadBytes(slice []byte, l int) []byte {
	return common.RightPadBytes(slice, l)
}

func LeftPadBytes(slice []byte, l int) []byte {
	return common.LeftPadBytes(slice, l)
}
