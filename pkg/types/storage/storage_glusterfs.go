package storage

import (
	"encoding/json"

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

const (
	StorageGlusterFSContents    = ContentQEMUData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageGlusterFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageGlusterFSShared      = AllowShareForced
	StorageGlusterFSSnapshots   = AllowSnapshotQcow2
	StorageGlusterFSClones      = AllowSnapshotQcow2
)

type StorageGlusterFSProperties struct {
	MainServer   string
	BackupServer string
	Transport    GlusterFSTransport

	Volume string
}

const (
	mkGlusterFSMainServer   = "server"
	mkGlusterFSBackupServer = "server2"
	mkGlusterFSTransport    = "transport"
	mkGlusterFSVolume       = "volume"
)

const (
	DefaultStorageGlusterFSBackupServer = ""
	DefaultStorageGlusterFSTransport    = GlusterFSTransportNone
)

func NewStorageGlusterFSProperties(
	props types.Properties,
) (*StorageGlusterFSProperties, error) {
	obj := new(StorageGlusterFSProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkGlusterFSMainServer,
				&obj.MainServer,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkGlusterFSBackupServer,
				&obj.BackupServer,
				DefaultStorageGlusterFSBackupServer,
				nil,
			)
		},
		func() error {
			return props.SetFixedValue(
				mkGlusterFSTransport,
				&obj.Transport,
				DefaultStorageGlusterFSTransport,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkGlusterFSVolume,
				&obj.Volume,
				nil,
			)
		},
	)
}

type GlusterFSTransport string

const (
	GlusterFSTransportNone GlusterFSTransport = ""
	GlusterFSTransportTCP  GlusterFSTransport = "tcp"
	GlusterFSTransportUNIX GlusterFSTransport = "unix"
	GlusterFSTransportRDMA GlusterFSTransport = "rdma"
)

func (obj GlusterFSTransport) IsValid() bool {
	switch obj {
	case GlusterFSTransportNone,
		GlusterFSTransportTCP,
		GlusterFSTransportUNIX,
		GlusterFSTransportRDMA:
		return true
	default:
		return false
	}
}

func (obj GlusterFSTransport) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj GlusterFSTransport) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *GlusterFSTransport) Unmarshal(s string) error {
	*obj = GlusterFSTransport(s)
	return nil
}

func (obj *GlusterFSTransport) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
