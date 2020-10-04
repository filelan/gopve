package storage

import (
	"encoding/json"
	"fmt"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageGlusterFS interface {
	Storage

	MainServer() string
	BackupServer() string
	Transport() GlusterFSTransport

	Volume() string
}

type StorageGlusterFSProperties struct {
	MainServer   string
	BackupServer string
	Transport    GlusterFSTransport

	Volume string
}

func NewStorageGlusterFSProperties(
	props types.Properties,
) (*StorageGlusterFSProperties, error) {
	obj := new(StorageGlusterFSProperties)

	if v, ok := props["server"].(string); ok {
		obj.MainServer = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	if v, ok := props["server2"].(string); ok {
		obj.BackupServer = v
	} else {
		obj.BackupServer = DefaultStorageGlusterFSBackupServer
	}

	if v, ok := props["transport"].(string); ok {
		if err := (&obj.Transport).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", "transport")
			err.AddKey("value", v)
			return nil, err
		}
	} else {
		obj.Transport = DefaultStorageGlusterFSTransport
	}

	if v, ok := props["volume"].(string); ok {
		obj.Volume = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "volume")
		return nil, err
	}

	return obj, nil
}

const (
	StorageGlusterFSContents    = ContentQEMUData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageGlusterFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageGlusterFSShared      = AllowShareForced
	StorageGlusterFSSnapshots   = AllowSnapshotQcow2
	StorageGlusterFSClones      = AllowSnapshotQcow2
)

const (
	DefaultStorageGlusterFSBackupServer = ""
	DefaultStorageGlusterFSTransport    = GlusterFSTransportNone
)

type GlusterFSTransport uint

const (
	GlusterFSTransportNone GlusterFSTransport = iota
	GlusterFSTransportTCP
	GlusterFSTransportUNIX
	GlusterFSTransportRDMA
)

func (obj GlusterFSTransport) Marshal() (string, error) {
	switch obj {
	case GlusterFSTransportNone:
		return "", nil
	case GlusterFSTransportTCP:
		return "tcp", nil
	case GlusterFSTransportUNIX:
		return "unix", nil
	case GlusterFSTransportRDMA:
		return "rdma", nil
	default:
		return "", fmt.Errorf("unknown glusterfs transport")
	}
}

func (obj *GlusterFSTransport) Unmarshal(s string) error {
	switch s {
	case "tcp":
		*obj = GlusterFSTransportTCP
	case "unix":
		*obj = GlusterFSTransportUNIX
	case "rdma":
		*obj = GlusterFSTransportRDMA
	default:
		return fmt.Errorf("can't unmarshal glusterfs transport %s", s)
	}

	return nil
}

func (obj *GlusterFSTransport) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
