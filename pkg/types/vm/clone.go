package vm

type CloneOptions struct {
	VMID         uint
	Name         string
	Description  string
	PoolName     string
	SnapshotName string

	BandwidthLimit    uint
	TemplateFullClone bool
	ImageFormat       QEMUImageFormat

	TargetNode    string
	TargetStorage string
}
