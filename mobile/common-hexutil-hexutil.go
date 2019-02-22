package geth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Decode(input string) ([]byte, error) {
	return hexutil.Decode(input)
}

func MustDecode(s string) []byte {
	return hexutil.MustDecode(s)
}

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

func DecodeBig(s string) (*BigInt, error) {
	if bigint, err := hexutil.DecodeBig(s); err == nil {
		return &BigInt{bigint: bigint}, nil
	} else {
		return nil, err
	}
}

func MustDecodeBig(s string) *BigInt {
	x := hexutil.MustDecodeBig(s)
	return &BigInt{bigint: x}
}

func EncodeBig(wbigint *BigInt) string {
	bigint := wbigint.bigint
	return hexutil.EncodeBig(bigint)
}
