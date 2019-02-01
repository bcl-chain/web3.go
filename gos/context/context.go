package context

import (
	"context"
)

type Context struct {
	Context interface{}
}

func Background() *Context {
	return &Context{
		Context: context.Background(),
	}
}
