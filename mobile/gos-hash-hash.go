package geth

import (
	"hash"
)

type GosHash struct {
	hash hash.Hash
}

func (wh *GosHash) Sum(b []byte) []byte {
	h := wh.hash
	return h.Sum(b)
}

func (wh *GosHash) Size() int {
	h := wh.hash
	return h.Size()
}

func (wh *GosHash) Write(p []byte) (int, error) {
	h := wh.hash
	return h.Write(p)
}
