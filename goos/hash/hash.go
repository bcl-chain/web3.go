package hash

import(
	"hash"
)

type Hash struct {
	Hash interface{}
}

func (wh *Hash) Sum(b []byte) []byte {
  h, _ := wh.Hash.(hash.Hash)
  return h.Sum(b)
}

func (wh *Hash) Size() int {
	h, _ := wh.Hash.(hash.Hash)
	return h.Size()
}

func (wh *Hash) Write(p []byte) (int, error) {
  h, _ := wh.Hash.(hash.Hash)
  return h.Write(p)
}


