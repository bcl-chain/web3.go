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

func NewContractCreation(nonce int64, wamount *Int, gasLimit int64, wgasPrice *Int, data []byte) *Transaction {
	amount := toBigInt(wamount)
	gasPrice := toBigInt(wgasPrice)

	tx := types.NewContractCreation(uint64(nonce), amount, uint64(gasLimit), gasPrice, data)
	return fromTransaction(tx)
}

func (wtx *Transaction) ChainId() *Int {
  tx := toTransaction(wtx)
	x := tx.ChainId()
	return fromBigInt(x)
}

func (wtx *Transaction) Protected() bool {
  tx := toTransaction(wtx)
	return tx.Protected()
}

// TODO
//func EncodeRLP {
//}

// TODO
//func DecodeRLP {
//}

func (wtx *Transaction) MarshalJSON() ([]byte, error) {
	tx := toTransaction(wtx)
	return tx.MarshalJSON()
}

func (wtx *Transaction) UnmarshalJSON(input []byte) error {
	tx := toTransaction(wtx)
	return tx.UnmarshalJSON(input)
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

func (wtx *Transaction) CheckNonce() bool {
	tx := toTransaction(wtx)
	return tx.CheckNonce()
}

func (wtx *Transaction) To() *Address {
	tx := toTransaction(wtx)
	return fromAddress(*tx.To())
}

func (wtx *Transaction) Hash() *Hash {
	tx := toTransaction(wtx)
	h := tx.Hash()
	return fromHash(h)
}

func (wtx *Transaction) Size() *StorageSize {
	tx := toTransaction(wtx)
	s := tx.Size()
	return fromStorageSize(s)
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

func (wtx *Transaction) Cost() *Int {
	tx := toTransaction(wtx)
	c := tx.Cost()
	return fromBigInt(c)
}

// TODO
//func RawSignatureValues {
//}

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

// function not in geth, added by K
func (wtxs *Transactions) Get(i int) *Transaction {
	txs := toTransactions(wtxs)
	return fromTransaction(txs[i])
}

func (wtxs *Transactions) Len() int {
	txs := toTransactions(wtxs)
	return txs.Len()
}

func (wtxs *Transactions) Swap(i, j int) {
	txs := toTransactions(wtxs)
	txs.Swap(i, j)
}

func (wtxs *Transactions) GetRlp(i int) []byte {
	txs := toTransactions(wtxs)
	return txs.GetRlp(i)
}

func TxDifference(wa, wb *Transactions) *Transactions {
	a := toTransactions(wa)
	b := toTransactions(wb)
	res := types.TxDifference(a, b)
	return fromTransactions(res)
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
}

func (wm *Message) To() *Address {
	m := toMessage(wm)
	to := *m.To()
	return fromAddress(to)
}
