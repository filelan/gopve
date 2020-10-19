package storage

import (
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

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkZFSPoolName,
				&obj.PoolName,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkZFSBlockSize,
				&obj.BlockSize,
				DefaultStorageZFSBlockSize,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkZFSUseSparse,
				&obj.UseSparse,
				DefaultStorageZFSUseSparse,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkZFSLocalPath,
				&obj.LocalPath,
				DefaultStorageZFSMountPoint,
				nil,
			)
		},
	)
}
