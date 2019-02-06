package api

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/core/types"
)

type Header struct {
	header *types.Header
}

func fromHeader(h *types.Header) *Header {
	if h != nil {
		return &Header{
			header: h,
		}
	} else {
		return nil
	}
}

func toHeader(wh *Header) *types.Header {
	if wh != nil {
		return wh.header
	} else {
		return nil
	}
}

func (wh *Header) ParentHash() *Hash {
	h := toHeader(wh)
	ph := h.ParentHash
	return fromHash(ph)
}

func (wh *Header) UncleHash() *Hash {
	h := toHeader(wh)
	uh := h.UncleHash
	return fromHash(uh)
}

func (wh *Header) Coinbase() *Address {
	h := toHeader(wh)
	c := h.Coinbase
	return fromAddress(c)
}

func (wh *Header) Root() *Hash {
	h := toHeader(wh)
	r := h.Root
	return fromHash(r)
}

func (wh *Header) TxHash() *Hash {
	h := toHeader(wh)
	th := h.TxHash
	return fromHash(th)
}

func (wh *Header) ReceiptHash() *Hash {
	h := toHeader(wh)
	rh := h.ReceiptHash
	return fromHash(rh)
}

// TODO
//func Bloom {
//}

func (wh *Header) Difficulty() *Int {
	h := toHeader(wh)
	d := h.Difficulty
	return fromBigInt(d)
}

func (wh *Header) Number() *Int {
	h := toHeader(wh)
	n := h.Number
	return fromBigInt(n)
}

func (wh *Header) GasLimit() int64 {
	h := toHeader(wh)
	gl := h.GasLimit
	return int64(gl)
}

func (wh *Header) GasUsed() int64 {
	h := toHeader(wh)
	gu := h.GasUsed
	return int64(gu)
}

func (wh *Header) Time() *Int {
	h := toHeader(wh)
	t := h.Time
	return fromBigInt(t)
}

func (wh *Header) Extra() []byte {
	h := toHeader(wh)
	return h.Extra
}

func (wh *Header) MixDigest() *Hash {
	h := toHeader(wh)
	md := h.MixDigest
	return fromHash(md)
}

func (wh *Header) Nonce() int64 {
	h := toHeader(wh)
	n := h.Nonce[:]
	un := binary.BigEndian.Uint64(n)
	return int64(un)
}

type Block struct {
	block *types.Block
}

func fromBlock(b *types.Block) *Block {
	if b != nil {
		return &Block{
			block: b,
		}
	} else {
		return nil
	}
}

func toBlock(wb *Block) *types.Block {
	if wb != nil {
		return wb.block
	} else {
		return nil
	}
}

func NewBlock() *Block {
	return fromBlock(new(types.Block))
}

func (wb *Block) Number() *Int {
	b := toBlock(wb)
	n := b.Number()
	return fromBigInt(n)
}

func (wb *Block) GasLimit() int64 {
	b := toBlock(wb)
	gl := b.GasLimit()
	return int64(gl)
}

func (wb *Block) GasUsed() int64 {
	b := toBlock(wb)
	gu := b.GasUsed()
	return int64(gu)
}

func (wb *Block) Difficulty() *Int {
	b := toBlock(wb)
	d := b.Difficulty()
	return fromBigInt(d)
}

func (wb *Block) Time() *Int {
	b := toBlock(wb)
	t := b.Time()
	return fromBigInt(t)
}

func (wb *Block) MixDigest() *Hash {
	b := toBlock(wb)
	md := b.MixDigest()
	return fromHash(md)
}

func (wb *Block) Nonce() int64 {
	b := toBlock(wb)
	n := b.Nonce()
	return int64(n)
}

// TODO
//func Bloom {
//}

func (wb *Block) Coinbase() *Address {
	b := toBlock(wb)
	c := b.Coinbase()
	return fromAddress(c)
}

func (wb *Block) Root() *Hash {
	b := toBlock(wb)
	r := b.Root()
	return fromHash(r)
}

func (wb *Block) ParentHash() *Hash {
	b := toBlock(wb)
	ph := b.ParentHash()
	return fromHash(ph)
}

func (wb *Block) TxHash() *Hash {
	b := toBlock(wb)
	th := b.TxHash()
	return fromHash(th)
}

func (wb *Block) ReceiptHash() *Hash {
	b := toBlock(wb)
	rh := b.ReceiptHash()
	return fromHash(rh)
}

func (wb *Block) UncleHash() *Hash {
	b := toBlock(wb)
	uh := b.UncleHash()
	return fromHash(uh)
}

func (wb *Block) Extra() []byte {
	b := toBlock(wb)
	return b.Extra()
}

func (wb *Block) Header() *Header {
	b := toBlock(wb)
	h := b.Header()
	return fromHeader(h)
}

func (wb *Block) Size() *StorageSize {
	b := toBlock(wb)
	ss := b.Size()
	return fromStorageSize(ss)
}

func (wb *Block) Hash() *Hash {
	b := toBlock(wb)
	h := b.Hash()
	return fromHash(h)
}

func (wb *Block) Transactions() *Transactions {
	b := toBlock(wb)
	txs := b.Transactions()
	return fromTransactions(txs)
}
