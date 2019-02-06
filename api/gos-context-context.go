package api

import (
	"context"
)

type Context struct {
	context context.Context
}

func fromContext(ctx context.Context) *Context {
	if ctx != nil {
		return &Context{
			context: ctx,
		}
	} else {
		return nil
	}
}

func toContext(wctx *Context) context.Context {
	if wctx != nil {
		return wctx.context
	} else {
		return nil
	}
}

func Background() *Context {
	ctx := context.Background()
	return fromContext(ctx)
}
