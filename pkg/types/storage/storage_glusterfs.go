package storage

import (
	"encoding/json"
	"fmt"
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

const (
	DefaultStorageGlusterFSBackupService = ""
	DefaultStorageGlusterFSTransport     = GlusterFSTransportNone
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
