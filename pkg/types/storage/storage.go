package storage

type Storage interface {
	Name() string
	Kind() Kind

	Content() Content
	Shared() bool
	Disabled() bool

	ImageFormat() ImageFormat
	MaxBackupsPerVM() uint

	Nodes() []string

	Digest() string
}

type Properties struct {
	Content  Content
	Shared   bool
	Disabled bool

	ImageFormat     ImageFormat
	MaxBackupsPerVM uint

	Nodes []string

	Digest string
}

type AllowShare int

const (
	AllowShareNever AllowShare = iota
	AllowSharePossible
	AllowShareForced
)

type AllowSnapshot int

const (
	AllowSnapshotNever AllowSnapshot = iota
	AllowSnapshotQcow2
	AllowSnapshotAll
)

type AllowClone int

const (
	AllowCloneNever AllowClone = iota
	AllowCloneQcow2
	AllowCloneAll
)
