package storage

type Storage interface {
	Name() string
	Kind() (Kind, error)
	Content() (Content, error)

	Shared() (bool, error)
	Disabled() (bool, error)

	ImageFormat() (ImageFormat, error)
	MaxBackupsPerVM() (uint, error)

	Nodes() ([]string, error)
}

type ExtraProperties map[string]interface{}

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
