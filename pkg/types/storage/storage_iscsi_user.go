package storage

type StorageISCSIUser interface {
	Storage

	// iSCSI target
	Portal() string

	// iSCSI portal (IP or DNS name with optional port).
	Target() string
}

type StorageISCSIUserProperties struct {
	Portal string
	Target string
}

func NewStorageISCSIUserProperties(props ExtraProperties) (*StorageISCSIUserProperties, error) {
	obj := new(StorageISCSIUserProperties)

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
	StorageISCSIUserContents    = ContentQEMUData
	StorageISCSIUserImageFormat = ImageFormatRaw
	StorageISCSIUserShared      = AllowShareForced
	StorageISCSIUserSnapshots   = AllowSnapshotNever
	StorageISCSIUserClones      = AllowCloneNever
)
