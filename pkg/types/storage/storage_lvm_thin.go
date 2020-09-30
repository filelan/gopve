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

func (obj *StorageLVMThinProperties) Unmarshal(props ExtraProperties) error {
	if v, ok := props["vgname"].(string); ok {
		obj.VolumeGroup = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "vgname")
		return err
	}

	if v, ok := props["thinpool"].(string); ok {
		obj.ThinPool = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "thinpool")
		return err
	}

	return nil
}

const (
	StorageLVMThinContents    = ContentQEMUData & ContentContainerData
	StorageLVMThinImageFormat = ImageFormatRaw
	StorageLVMThinShared      = AllowShareNever
	StorageLVMThinSnapshots   = AllowSnapshotAll
	StorageLVMThinClones      = AllowCloneAll
)
