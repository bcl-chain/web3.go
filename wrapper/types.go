package wrapper

import(
  "github.com/ethereum/go-ethereum/common"

  wcommon "github.com/bcl-chain/web3.go/common"
)

//type Address struct {
//  Address interface{}
//}

//type Hash struct {
//  Hash interface{}
//}

func FromAddress(address common.Address) *wcommon.Address {
  return &wcommon.Address{
    Address: address,
  }
}

func ToAddress(waddress *wcommon.Address) common.Address {
  address, _ := waddress.Address.(common.Address)
  return address
}

