package wrapper

import(
  "context"

   wcontext "github.com/bcl-chain/web3.go/common/context"
)

func FromContext(context context.Context) *wcontext.Context {
  if context != nil {
    return &wcontext.Context{
      Context: context,
    }
  } else {
    return nil
  }
}

func ToContext(wcontext *wcontext.Context) context.Context {
  if wcontext != nil {
    if context, ok := wcontext.Context.(context.Context); ok {
      return context
    } else {
      return nil
    }
  } else {
    return nil
  }
}

