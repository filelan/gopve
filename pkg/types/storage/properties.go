package storage

type Properties struct {
	Content  Content
	Shared   bool
	Disabled bool

	ImageFormat     ImageFormat
	MaxBackupsPerVM uint

	Nodes []string

	ExtraProperties ExtraProperties
}
