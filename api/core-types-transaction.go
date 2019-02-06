package api

import (
	"github.com/ethereum/go-ethereum/core/types"
	//	"github.com/bcl-chain/web3.go/common"
	//	"github.com/bcl-chain/web3.go/gos/math/big"
	//	wcommon "github.com/bcl-chain/web3.go/wrapper/common"
	//	wbig "github.com/bcl-chain/web3.go/wrapper/gos/math/big"
)

type Transaction struct {
	transaction *types.Transaction
}

func fromTransaction(tx *types.Transaction) *Transaction {
	if tx != nil {
		return &Transaction{
			transaction: tx,
		}
	} else {
		return nil
	}
}

func toTransaction(wtx *Transaction) *types.Transaction {
	if wtx != nil {
		return wtx.transaction
	} else {
		return nil
	}
}

func NewTransaction(nonce int64, wto *Address, wamount *Int, gasLimit int64, wgasPrice *Int, data []byte) *Transaction {
	to := toAddress(wto)
	amount := toBigInt(wamount)
	gasPrice := toBigInt(wgasPrice)

	tx := types.NewTransaction(uint64(nonce), to, amount, uint64(gasLimit), gasPrice, data)
	return fromTransaction(tx)
}

func (wtx *Transaction) Data() []byte {
	tx := toTransaction(wtx)
	return tx.Data()
}

func (wtx *Transaction) Gas() int64 {
	tx := toTransaction(wtx)
	return int64(tx.Gas())
}

func (wtx *Transaction) GasPrice() *Int {
	tx := toTransaction(wtx)
	gp := tx.GasPrice()
	return fromBigInt(gp)
}

func (wtx *Transaction) Value() *Int {
	tx := toTransaction(wtx)
	v := tx.Value()
	return fromBigInt(v)
}

func (wtx *Transaction) Nonce() int64 {
	tx := toTransaction(wtx)
	return int64(tx.Nonce())
}

func (wtx *Transaction) To() *Address {
	tx := toTransaction(wtx)
	return fromAddress(*tx.To())
	//	return &common.Address{
	//		Address: *tx.To(),
	//	}
}

func (wtx *Transaction) Hash() *Hash {
	tx := toTransaction(wtx)
	h := tx.Hash()
	return fromHash(h)
}

func (wtx *Transaction) AsMessage(ws *Signer) (*Message, error) {
	tx := toTransaction(wtx)
	s := toSigner(ws)
	msg, err := tx.AsMessage(s)
	return fromMessage(msg), err
}

func (wtx *Transaction) WithSignature(ws *Signer, sig []byte) (*Transaction, error) {
	tx := toTransaction(wtx)
	s := toSigner(ws)
	tx2, err := tx.WithSignature(s, sig)
	return fromTransaction(tx2), err
}

type Transactions struct {
	transactions types.Transactions
}

func fromTransactions(txs types.Transactions) *Transactions {
	return &Transactions{
		transactions: txs,
	}
}

func toTransactions(wtxs *Transactions) types.Transactions {
	if wtxs != nil {
		return wtxs.transactions
	} else {
		return nil
	}
}

func (ws *Transactions) Len() int {
	s := toTransactions(ws)
	return s.Len()
}

// function not in geth, added by K
func (ws *Transactions) Get(i int) *Transaction {
	s := toTransactions(ws)
	return fromTransaction(s[i])
}

type Message struct {
	message types.Message
}

func fromMessage(m types.Message) *Message {
	return &Message{
		message: m,
	}
}

func toMessage(wm *Message) types.Message {
	return wm.message
}

func (wm *Message) From() *Address {
	m := toMessage(wm)
	f := m.From()
	return fromAddress(f)
	//	return &common.Address {
	//		Address: m.From(),
	//	}
}

func (wm *Message) To() *Address {
	m := toMessage(wm)
	return fromAddress(*m.To())
	//	return &common.Address{
	//		Address: *m.To(),
	//	}
}
