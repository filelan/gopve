package storage

type StorageISCSIUser interface {
	Storage

	// iSCSI target
	Portal() string

	// iSCSI portal (IP or DNS name with optional port).
	Target() string
}

const (
	StorageISCSIUserContents    = ContentQEMUData
	StorageISCSIUserImageFormat = ImageFormatRaw
	StorageISCSIUserShared      = AllowShareForced
	StorageISCSIUserSnapshots   = AllowSnapshotNever
	StorageISCSIUserClones      = AllowCloneNever
)
