package types

import (
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/goos/math/big"
)

type Transaction struct {
  Transaction interface{}
}

func (wtx *Transaction) Data() []byte {
  tx, _ := wtx.Transaction.(*types.Transaction)
  return tx.Data()
}

func (wtx *Transaction) Gas() int64 {
  tx, _ := wtx.Transaction.(*types.Transaction)
	return int64(tx.Gas())
}

func (wtx *Transaction) GasPrice() *big.Int {
  tx, _ := wtx.Transaction.(*types.Transaction)
	return &big.Int{
		Int: tx.GasPrice(),
	}
}

func (wtx *Transaction) Value() *big.Int {
  tx, _ := wtx.Transaction.(*types.Transaction)
  return &big.Int{
    Int: tx.Value(),
  }
}

func (wtx *Transaction) Nonce() int64 {
  tx, _ := wtx.Transaction.(*types.Transaction)
  return int64(tx.Nonce())
}

func (wtx *Transaction) To() *common.Address {
  tx, _ := wtx.Transaction.(*types.Transaction)
  return &common.Address{
    Address: *tx.To(),
  }
}

func (wtx *Transaction) Hash() *common.Hash {
  tx, _ := wtx.Transaction.(*types.Transaction)
  return &common.Hash{
    Hash: tx.Hash(),
  }
}

func (wtx *Transaction) AsMessage(ws *Signer) (*Message, error) {
  tx, _ := wtx.Transaction.(*types.Transaction)
	s, _ := ws.Signer.(types.Signer)
	msg, err := tx.AsMessage(s)
	return &Message{
		Message: msg,
	}, err
}

type Transactions struct {
	Transactions interface{}
}

func (ws Transactions) Len() int {
	s, _ := ws.Transactions.(types.Transactions)
	return s.Len()
}

// function not in geth, added by K 
func (ws Transactions) Get(i int)  *Transaction {
	s, _ := ws.Transactions.(types.Transactions)
	return &Transaction{
		Transaction: s[i],
	}
}

type Message struct {
	Message interface{}
}

func (wm Message) From() *common.Address {
  m, _ := wm.Message.(types.Message)
  return &common.Address{
    Address: m.From(),
  }
}

func (wm Message) To() *common.Address {
  m, _ := wm.Message.(types.Message)
  return &common.Address{
    Address: *m.To(),
  }
}

