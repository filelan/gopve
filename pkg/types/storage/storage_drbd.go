package storage

import "github.com/xabinapal/gopve/pkg/types"

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

	if err := props.SetUint(mkDRBDRedundancy, &obj.Redundancy, DefaultStorageDRBDRedundancy, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
