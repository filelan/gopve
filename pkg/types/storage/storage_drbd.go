package storage

import "github.com/xabinapal/gopve/pkg/types"

type StorageDRBD interface {
	Storage

	Redundancy() uint
}

type StorageDRBDProperties struct {
	Redundancy uint
}

func NewStorageDRBDProperties(
	props types.Properties,
) (*StorageDRBDProperties, error) {
	obj := new(StorageDRBDProperties)

	if v, ok := props["redundancy"].(float64); ok {
		obj.Redundancy = uint(v)
	} else {
		obj.Redundancy = DefaultStorageDRBDRedundancy
	}

	return obj, nil
}

const (
	StorageDRBDContents    = ContentQEMUData & ContentContainerData
	StorageDRBDImageFormat = ImageFormatRaw
	StorageDRBDShared      = AllowShareForced
	StorageDRBDSnapshots   = AllowSnapshotNever
	StorageDRBDClones      = AllowCloneNever
)

const (
	DefaultStorageDRBDRedundancy uint = 2
)
