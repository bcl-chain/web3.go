package wrapper

import (
  "github.com/ethereum/go-ethereum/core/types"

  wtypes "github.com/bcl-chain/web3.go/core/types"
)

func FromTransactions(transactions *types.Transactions) *wtypes.Transactions {
	if transactions != nil {
		return &wtypes.Transactions{
			Transactions: transactions,
		}
	} else {
		return nil
	}
}

func ToTransactions(wTransactions *wtypes.Transactions) *types.Transactions {
	if wTransactions != nil {
		if transactions, ok := wTransactions.Transactions.(*types.Transactions); ok {
			return transactions
		} else {
			return nil
		}
	} else {
		return nil
	}
}

