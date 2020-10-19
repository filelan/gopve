package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageLVM interface {
	Storage

	// Base volume. This volume is automatically activated.
	BaseStorage() string

	// Volume group name.
	VolumeGroup() string

	// Zero-out data when removing LVs.
	SafeRemove() bool

	// Limit the thoughput of the data stream, in bytes per second. If the value is positive, tries to keep the overall rate at the specified value for the whole session. If the value is negative, it is an upper limit for each read/write system call pair. In other words, the negative number will never exceed that limit, the positive number will exceed it to make good for previous underutilization
	SafeRemoveThroughput() int

	// Only use logical volumes tagged with 'pve-vm-ID'.
	TaggedOnly() bool
}

const (
	StorageLVMContent       = ContentQEMUData & ContentContainerData
	StorageLVMImageFormat   = ImageFormatRaw
	StorageLVMAllowShare    = AllowSharePossible
	StorageLVMAllowSnapshot = AllowSnapshotNever
	StorageLVMAllowClone    = AllowCloneNever
)

type StorageLVMProperties struct {
	BaseStorage          string
	VolumeGroup          string
	SafeRemove           bool
	SafeRemoveThroughput int
	TaggedOnly           bool
}

const (
	mkLVMBaseStorage          = "base"
	mkLVMVolumeGroup          = "vgname"
	mkLVMSafeRemove           = "saferemove"
	mkLVMSafeRemoveThroughput = "saferemove_throughput"
	mkLVMTaggedOnly           = "tagged_only"
)

const (
	DefaultStorageLVMBaseStorage          = ""
	DefaultStorageLVMSafeRemove           = false
	DefaultStorageLVMSafeRemoveThroughput = -10485760
	DefaultStorageLVMTaggedOnly           = false
)

func NewStorageLVMProperties(
	props types.Properties,
) (*StorageLVMProperties, error) {
	obj := new(StorageLVMProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetString(
				mkLVMBaseStorage,
				&obj.BaseStorage,
				DefaultStorageLVMBaseStorage,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkLVMVolumeGroup,
				&obj.VolumeGroup,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkLVMSafeRemove,
				&obj.SafeRemove,
				DefaultStorageLVMSafeRemove,
				nil,
			)
		},
		func() error {
			return props.SetInt(
				mkLVMSafeRemoveThroughput,
				&obj.SafeRemoveThroughput,
				DefaultStorageLVMSafeRemoveThroughput,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkLVMTaggedOnly,
				&obj.TaggedOnly,
				DefaultStorageLVMTaggedOnly,
				nil,
			)
		},
	)
}
