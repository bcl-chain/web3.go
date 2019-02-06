package api

import (
	//	"bytes"
	//	"testing"

	checker "gopkg.in/check.v1"
)

type BytesSuite struct{}

var _ = checker.Suite(&BytesSuite{})

func (s *BytesSuite) TestLeftPadBytes(c *checker.C) {
	val1 := []byte{1, 2, 3, 4}
	exp1 := []byte{0, 0, 0, 0, 1, 2, 3, 4}

	res1 := LeftPadBytes(val1, 8)
	res2 := LeftPadBytes(val1, 2)

	c.Assert(res1, checker.DeepEquals, exp1)
	c.Assert(res2, checker.DeepEquals, val1)
}
