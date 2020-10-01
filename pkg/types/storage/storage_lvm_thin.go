package storage

type StorageLVMThin interface {
	Storage

	VolumeGroup() string
	ThinPool() string
}

type StorageLVMThinProperties struct {
	VolumeGroup string
	ThinPool    string
}

func NewStorageLVMThinProperties(props ExtraProperties) (*StorageLVMThinProperties, error) {
	obj := new(StorageLVMThinProperties)

	if v, ok := props["vgname"].(string); ok {
		obj.VolumeGroup = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "vgname")
		return nil, err
	}

	if v, ok := props["thinpool"].(string); ok {
		obj.ThinPool = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "thinpool")
		return nil, err
	}

	return obj, nil
}

const (
	StorageLVMThinContents    = ContentQEMUData & ContentContainerData
	StorageLVMThinImageFormat = ImageFormatRaw
	StorageLVMThinShared      = AllowShareNever
	StorageLVMThinSnapshots   = AllowSnapshotAll
	StorageLVMThinClones      = AllowCloneAll
)
