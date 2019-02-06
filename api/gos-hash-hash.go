package api

import (
	"hash"
)

type GosHash struct {
	hash hash.Hash
}

func fromGosHash(h hash.Hash) *GosHash {
	if h != nil {
		return &GosHash{
			hash: h,
		}
	} else {
		return nil
	}
}

func toGosHash(wh *GosHash) hash.Hash {
	if wh != nil {
		return wh.hash
	} else {
		return nil
	}
}

func (wh *GosHash) Sum(b []byte) []byte {
	h := toGosHash(wh)
	return h.Sum(b)
}

func (wh *GosHash) Size() int {
	h := toGosHash(wh)
	return h.Size()
}

func (wh *GosHash) Write(p []byte) (int, error) {
	h := toGosHash(wh)
	return h.Write(p)
}
