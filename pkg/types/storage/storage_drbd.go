package storage

type StorageDRBD interface {
	Storage

	Redundancy() uint
}

type StorageDRBDProperties struct {
	Redundancy uint
}

func NewStorageDRBDProperties(props ExtraProperties) (*StorageDRBDProperties, error) {
	obj := new(StorageDRBDProperties)

	if v, ok := props["redundancy"].(int); ok {
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
