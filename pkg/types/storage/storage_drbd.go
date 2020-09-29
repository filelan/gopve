package storage

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

const (
	DefaultStorageDRBDRedundancy uint = 2
)
