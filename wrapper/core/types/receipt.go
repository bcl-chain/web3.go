package wrapper

import (
	"github.com/ethereum/go-ethereum/core/types"

	wtypes "github.com/bcl-chain/web3.go/core/types"
)

func FromReceipt(receipt *types.Receipt) *wtypes.Receipt {
	if receipt != nil {
		return &wtypes.Receipt{
			Receipt: receipt,
		}
	} else {
		return nil
	}
}

func ToReceipt(wReceipt *wtypes.Receipt) *types.Receipt {
	if wReceipt != nil {
		if receipt, ok := wReceipt.Receipt.(*types.Receipt); ok {
			return receipt
		} else {
			return nil
		}
	} else {
		return nil
	}
}

