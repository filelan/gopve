package storage

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

const (
	DefaultStorageLVMSafeRemove           = false
	DefaultStorageLVMSaveRemoveThroughput = -10485760
	DefaultStorageLVMTaggedOnly           = false
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
