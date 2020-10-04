package storage

import (
	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageZFS interface {
	Storage

	PoolName() string

	BlockSize() string
	UseSparse() bool

	LocalPath() string
}

type StorageZFSProperties struct {
	PoolName  string
	BlockSize string
	UseSparse bool
	LocalPath string
}

func NewStorageZFSProperties(
	props types.Properties,
) (*StorageZFSProperties, error) {
	obj := new(StorageZFSProperties)

	if v, ok := props["pool"].(string); ok {
		obj.PoolName = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "pool")
		return nil, err
	}

	if v, ok := props["blocksize"].(string); ok {
		obj.BlockSize = v
	} else {
		obj.BlockSize = DefaultStorageZFSBlockSize
	}

	if v, ok := props["sparse"].(float64); ok {
		obj.UseSparse = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.UseSparse = DefaultStorageZFSUseSparse
	}

	if v, ok := props["mountpoint"].(string); ok {
		obj.LocalPath = v
	} else {
		obj.LocalPath = DefaultStorageZFSMountPoint
	}

	return obj, nil
}

const (
	StorageZFSContents    = ContentQEMUData & ContentContainerData
	StorageZFSImageFormat = ImageFormatRaw & ImageFormatSubVolume
	StorageZFSShared      = AllowShareNever
	StorageZFSSnapshots   = AllowSnapshotAll
	StorageZFSClones      = AllowCloneAll
)

const (
	DefaultStorageZFSBlockSize  = "8192"
	DefaultStorageZFSUseSparse  = false
	DefaultStorageZFSMountPoint = ""
)
