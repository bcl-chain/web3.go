package sha3

import(
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/bcl-chain/web3.go/goos/hash"
)

func NewKeccak256() *hash.Hash{
	return &hash.Hash{
		Hash: sha3.NewKeccak256(),
	}
}
