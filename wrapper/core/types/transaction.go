package wrapper

import (
	"github.com/ethereum/go-ethereum/core/types"

	wtypes "github.com/bcl-chain/web3.go/core/types"
)

func FromTransaction(transaction *types.Transaction) *wtypes.Transaction {
  if transaction != nil {
    return &wtypes.Transaction{
      Transaction: transaction,
    }
  } else {
    return nil
  }
}

func ToTransaction(wTransaction *wtypes.Transaction) *types.Transaction {
  if wTransaction != nil {
    if transaction, ok := wTransaction.Transaction.(*types.Transaction); ok {
      return transaction
    } else {
      return nil
    }
  } else {
    return nil
  }
}

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

