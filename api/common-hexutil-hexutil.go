package api

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

func DecodeBig(s string) (*Int, error) {
	if bigint, err := hexutil.DecodeBig(s); err == nil {
		return fromBigInt(bigint), nil
	} else {
		return nil, err
	}
}

func MustDecodeBig(s string) *Int {
	x := hexutil.MustDecodeBig(s)
	return fromBigInt(x)
}

func EncodeBig(wbigint *Int) string {
	bigint := toBigInt(wbigint)
	return hexutil.EncodeBig(bigint)
}
