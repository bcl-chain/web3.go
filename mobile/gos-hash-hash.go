package geth

import (
	"hash"
)

// GosHash ...
type GosHash struct {
	hash hash.Hash
}

// Sum ...
func (wh *GosHash) Sum(b []byte) []byte {
	h := wh.hash
	return h.Sum(b)
}

// Size ...
func (wh *GosHash) Size() int {
	h := wh.hash
	return h.Size()
}

// Write ...
func (wh *GosHash) Write(p []byte) (int, error) {
	h := wh.hash
	return h.Write(p)
}
