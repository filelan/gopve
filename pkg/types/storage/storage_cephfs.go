package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageCephFS interface {
	Storage

	MonitorHosts() []string
	Username() string

	// Mount CephFS through FUSE.
	UseFUSE() bool

	ServerPath() string
	LocalPath() string
}

const (
	StorageCephFSContents    = ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCephFSImageFormat = ContentNone
	StorageCephFSShared      = AllowShareNever
	StorageCephFSSnapshots   = AllowSnapshotAll
	StorageCephFSClones      = AllowCloneNever
)

type StorageCephFSProperties struct {
	MonitorHosts []string
	Username     string

	UseFUSE bool

	ServerPath string
	LocalPath  string
}

const (
	mkCephFSMonitorHosts = "monhost"
	mkCephFSUsername     = "username"
	mkCephFSUseFUSE      = "fuse"
	mkCephFSServerPath   = "subdir"
	mkCephFSLocalPath    = "path"
)

var DefaultStorageCephFSMonitorHosts = []string{}

const (
	DefaultStorageCephFSUsername   = ""
	DefaultStorageCephFSUseFUSE    = false
	DefaultStorageCephFSServerPath = "/"
)

func NewStorageCephFSProperties(
	props types.Properties,
) (*StorageCephFSProperties, error) {
	obj := new(StorageCephFSProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			if v, err := props.GetAsList(mkCephFSMonitorHosts, " "); err == nil {
				obj.MonitorHosts = v.List()
			} else if errors.ErrMissingProperty.IsBase(err) {
				obj.MonitorHosts = DefaultStorageCephFSMonitorHosts
			} else {
				return err
			}

			return nil
		},
		func() error {
			return props.SetString(
				mkCephFSUsername,
				&obj.Username,
				DefaultStorageCephFSUsername,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkCephFSUseFUSE,
				&obj.UseFUSE,
				DefaultStorageCephFSUseFUSE,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkCephFSServerPath,
				&obj.ServerPath,
				DefaultStorageCephFSServerPath,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkCephFSLocalPath,
				&obj.LocalPath,
				nil,
			)
		},
	)
}
