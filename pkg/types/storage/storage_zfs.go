package storage

import "github.com/xabinapal/gopve/internal/types"

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

func (obj *StorageZFSProperties) Unmarshal(props ExtraProperties) error {
	if v, ok := props["pool"].(string); ok {
		obj.PoolName = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "pool")
		return err
	}

	if v, ok := props["blocksize"].(string); ok {
		obj.BlockSize = v
	} else {
		obj.BlockSize = DefaultStorageZFSBlockSize
	}

	if v, ok := props["sparse"].(int); ok {
		obj.UseSparse = types.NewPVEBoolFromInt(v).Bool()
	} else {
		obj.UseSparse = DefaultStorageZFSUseSparse
	}

	if v, ok := props["mountpoint"].(string); ok {
		obj.LocalPath = v
	} else {
		obj.LocalPath = DefaultStorageZFSMountPoint
	}

	return nil
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
