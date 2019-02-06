package api

import (
	"github.com/ethereum/go-ethereum/common"
)

type StorageSize struct {
	storageSize common.StorageSize
}

func fromStorageSize(ss common.StorageSize) *StorageSize {
	return &StorageSize{
		storageSize: ss,
	}
}

func toStorageSize(wss *StorageSize) common.StorageSize {
	ss := wss.storageSize
	return ss
}

func StorageSizeInt(i int) *StorageSize {
	ss := common.StorageSize(i)
	return &StorageSize{
		storageSize: ss,
	}
}
