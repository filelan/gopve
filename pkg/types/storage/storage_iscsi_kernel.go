package storage

type StorageISCSIKernel interface {
	Storage

	Portal() string
	Target() string
}

const (
	StorageISCSIKernelContents    = ContentQEMUData
	StorageISCSIKernelImageFormat = ImageFormatRaw
	StorageISCSIKernelShared      = AllowShareForced
	StorageISCSIKernelSnapshots   = AllowSnapshotNever
	StorageISCSIKernelClones      = AllowCloneNever
)
