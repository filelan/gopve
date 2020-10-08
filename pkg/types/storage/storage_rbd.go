package storage

import (
	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageRBD interface {
	Storage

	MonitorHosts() []string
	Username() string

	// Always access rbd through krbd kernel module
	UseKRBD() bool

	PoolName() string
}

const (
	StorageRBDContents    = ContentQEMUData & ContentContainerData
	StorageRBDImageFormat = ImageFormatRaw
	StorageRBDShared      = AllowShareForced
	StorageRBDSnapshots   = AllowSnapshotAll
	StorageRBDClones      = AllowCloneAll
)

type StorageRBDProperties struct {
	MonitorHosts []string
	Username     string

	UseKRBD bool

	PoolName string
}

const (
	mkRBDMonitorHosts = "monhost"
	mkRBDUsername     = "username"
	mkRBDUseKRBD      = "krbd"
	mkRBDPoolName     = "pool"
)

var DefaultStorageRBDMonitorHosts = []string{}

const (
	DefaultStorageRBDUsername = ""
	DefaultStorageRBDUseKRBD  = false
	DefaultStorageRBDPoolName = "rbd"
)

func NewStorageRBDProperties(
	props types.Properties,
) (*StorageRBDProperties, error) {
	obj := new(StorageRBDProperties)

	if v, ok := props[mkRBDMonitorHosts].(string); ok {
		monitorHosts := internal_types.PVEList{Separator: " "}
		if err := (&monitorHosts).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkRBDMonitorHosts)
			err.AddKey("value", v)
			return nil, err
		}

		obj.MonitorHosts = monitorHosts.List()
	} else {
		obj.MonitorHosts = DefaultStorageRBDMonitorHosts
	}

	if err := props.SetString(mkRBDUsername, &obj.Username, DefaultStorageRBDUsername, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkRBDUseKRBD, &obj.UseKRBD, DefaultStorageRBDUseKRBD, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkRBDPoolName, &obj.PoolName, DefaultStorageRBDPoolName, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
