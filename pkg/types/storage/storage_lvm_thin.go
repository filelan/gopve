package storage

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
