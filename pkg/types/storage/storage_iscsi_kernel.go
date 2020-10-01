package storage

type StorageISCSIKernel interface {
	Storage

	Portal() string
	Target() string
}

type StorageISCSIKernelProperties struct {
	Portal string
	Target string
}

func NewStorageISCSIKernelProperties(props ExtraProperties) (*StorageISCSIKernelProperties, error) {
	obj := new(StorageISCSIKernelProperties)

	if v, ok := props["portal"].(string); ok {
		obj.Portal = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "portal")
		return nil, err
	}

	if v, ok := props["target"].(string); ok {
		obj.Target = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "target")
		return nil, err
	}

	return obj, nil
}

const (
	StorageISCSIKernelContents    = ContentQEMUData
	StorageISCSIKernelImageFormat = ImageFormatRaw
	StorageISCSIKernelShared      = AllowShareForced
	StorageISCSIKernelSnapshots   = AllowSnapshotNever
	StorageISCSIKernelClones      = AllowCloneNever
)
