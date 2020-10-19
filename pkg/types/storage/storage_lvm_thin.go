package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageLVMThin interface {
	Storage

	VolumeGroup() string
	ThinPool() string
}

const (
	StorageLVMThinContents    = ContentQEMUData & ContentContainerData
	StorageLVMThinImageFormat = ImageFormatRaw
	StorageLVMThinShared      = AllowShareNever
	StorageLVMThinSnapshots   = AllowSnapshotAll
	StorageLVMThinClones      = AllowCloneAll
)

type StorageLVMThinProperties struct {
	VolumeGroup string
	ThinPool    string
}

const (
	mkLVMThinVolumeGroup = "vgname"
	mkLVMThinThinPool    = "thinpool"
)

func NewStorageLVMThinProperties(
	props types.Properties,
) (*StorageLVMThinProperties, error) {
	obj := new(StorageLVMThinProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkLVMThinVolumeGroup,
				&obj.VolumeGroup,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkLVMThinThinPool,
				&obj.ThinPool,
				nil,
			)
		},
	)
}
