package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
)

type StorageZFS interface {
	Storage

	PoolName() string

	BlockSize() string
	UseSparse() bool

	LocalPath() string
}

const (
	StorageZFSContents    = ContentQEMUData & ContentContainerData
	StorageZFSImageFormat = ImageFormatRaw & ImageFormatSubVolume
	StorageZFSShared      = AllowShareNever
	StorageZFSSnapshots   = AllowSnapshotAll
	StorageZFSClones      = AllowCloneAll
)

type StorageZFSProperties struct {
	PoolName  string
	BlockSize string
	UseSparse bool
	LocalPath string
}

const (
	mkZFSPoolName  = "pool"
	mkZFSBlockSize = "blocksize"
	mkZFSUseSparse = "sparse"
	mkZFSLocalPath = "mountpoint"
)

const (
	DefaultStorageZFSBlockSize  = "8192"
	DefaultStorageZFSUseSparse  = false
	DefaultStorageZFSMountPoint = ""
)

func NewStorageZFSProperties(
	props types.Properties,
) (*StorageZFSProperties, error) {
	obj := new(StorageZFSProperties)

	if err := props.SetRequiredString(mkZFSPoolName, &obj.PoolName, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkZFSBlockSize, &obj.BlockSize, DefaultStorageZFSBlockSize, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkZFSUseSparse, &obj.UseSparse, DefaultStorageZFSUseSparse, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkZFSLocalPath, &obj.LocalPath, DefaultStorageZFSMountPoint, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
