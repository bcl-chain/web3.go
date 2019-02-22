package geth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Decode ...
func Decode(input string) ([]byte, error) {
	return hexutil.Decode(input)
}

// MustDecode ...
func MustDecode(s string) []byte {
	return hexutil.MustDecode(s)
}

// Encode ...
func Encode(b []byte) string {
	return hexutil.Encode(b)
}

// TODO
//func DecodeUint64 {
//}

// TODO
//func MustDecodeUint64 {
//}

// TODO
//func EncodeUint64 {
//}

// DecodeBig ...
func DecodeBig(s string) (*BigInt, error) {
	bigint, err := hexutil.DecodeBig(s)
	if err == nil {
		return &BigInt{bigint: bigint}, nil
	}
	return nil, err
}

// MustDecodeBig ...
func MustDecodeBig(s string) *BigInt {
	x := hexutil.MustDecodeBig(s)
	return &BigInt{bigint: x}
}

// EncodeBig ...
func EncodeBig(wbigint *BigInt) string {
	bigint := wbigint.bigint
	return hexutil.EncodeBig(bigint)
}
