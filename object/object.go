// Copyright 2019 The bcl-chain Authors. All rights reserved.
// Package object implements ancestor for all structs in web3.go.
//

package object

import (
	"reflect"
)

var objs map[uintptr]interface{}

func init() {
	objs = make(map[uintptr]interface{})
}

func key2Ptr(key interface{}) uintptr {
	val := reflect.ValueOf(key)
	return val.Pointer()
}

func newObject(ptr uintptr, value interface{}) (err error) {
	objs[ptr] = value
	return
}

func NewObject(key, value interface{}) interface{} {
	if err := newObject(key2Ptr(key), value); err != nil {
		return nil
	}
	return key
}

func getObject(ptr uintptr) interface{} {
	return objs[ptr]
}

func GetObject(key interface{}) interface{} {
	return getObject(key2Ptr(key))
}

func freeObject(ptr uintptr) {
	delete(objs, ptr)
}

func FreeObject(key interface{}) {
	freeObject(key2Ptr(key))
}
