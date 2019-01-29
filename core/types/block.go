package types

import (
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bcl-chain/web3.go/common"
	"github.com/bcl-chain/web3.go/goos/math/big"
)

type Header struct {
	Header interface{}
}

func (h *Header) Number() *big.Int {
	header, _ := h.Header.(*types.Header)
	return &big.Int{
		Int: header.Number,
	}
}

type Block struct {
	Block interface{}
}

func (b *Block) Number() *big.Int {
	block, _ := b.Block.(*types.Block)
	return &big.Int{
		Int: block.Number(),
	}
}

func (b *Block) Difficulty() *big.Int {
	block, _ := b.Block.(*types.Block)
	return &big.Int{
		Int: block.Difficulty(),
	}
}

func (b *Block) Time() *big.Int {
	block, _ := b.Block.(*types.Block)
	return &big.Int{
		Int: block.Time(),
	}
}

func (b *Block) Hash() *common.Hash {
	block, _ := b.Block.(*types.Block)
	return &common.Hash{
		Hash: block.Hash(),
	}
}

func (b *Block) Transactions() *Transactions {
	block, _ := b.Block.(*types.Block)
	return &Transactions{
		Transactions: block.Transactions(),
	}
}
