package api

import (
	"github.com/ethereum/go-ethereum/core/types"
	//	"github.com/bcl-chain/web3.go/common"
	//	"github.com/bcl-chain/web3.go/gos/math/big"
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

func (wh *Header) Number() *Int {
	h := toHeader(wh)
	return fromBigInt(h.Number)
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

func (wb *Block) Number() *Int {
	b := toBlock(wb)
	n := b.Number()
	return fromBigInt(n)
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
