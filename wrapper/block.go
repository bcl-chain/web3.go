package wrapper

import(
  "github.com/ethereum/go-ethereum/core/types"

  wtypes "github.com/bcl-chain/web3.go/core/types"
)

func FromHeader(header *types.Header) *wtypes.Header {
  if header != nil {
    return &wtypes.Header{
      Header: header,
    }
  } else {
    return nil
  }
}

func ToHeader(wheader *wtypes.Header) *types.Header {
  if wheader != nil {
    if header, ok := wheader.Header.(*types.Header); ok {
      return header
    } else {
      return nil
    }
  } else {
    return nil
  }
}

func FromBlock(block *types.Block) *wtypes.Block {
  if block != nil {
    return &wtypes.Block{
      Block: block,
    }
  } else {
    return nil
  }
}

func ToBlock(wblock *wtypes.Block) *types.Block {
  if wblock != nil {
    if block, ok := wblock.Block.(*types.Block); ok {
      return block
    } else {
      return nil
    }
  } else {
    return nil
  }
}

