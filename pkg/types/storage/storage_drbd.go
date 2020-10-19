package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageDRBD interface {
	Storage

	Redundancy() uint
}

const (
	StorageDRBDContents    = ContentQEMUData & ContentContainerData
	StorageDRBDImageFormat = ImageFormatRaw
	StorageDRBDShared      = AllowShareForced
	StorageDRBDSnapshots   = AllowSnapshotNever
	StorageDRBDClones      = AllowCloneNever
)

type StorageDRBDProperties struct {
	Redundancy uint
}

const (
	mkDRBDRedundancy = "redundancy"
)

const (
	DefaultStorageDRBDRedundancy uint = 2
)

func NewStorageDRBDProperties(
	props types.Properties,
) (*StorageDRBDProperties, error) {
	obj := new(StorageDRBDProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetUint(
				mkDRBDRedundancy,
				&obj.Redundancy,
				DefaultStorageDRBDRedundancy,
				nil,
			)
		},
	)
}
