package storage

import (
	"encoding/json"
	"fmt"
)

type StorageNFS interface {
	Storage

	Server() string
	NFSVersion() NFSVersion

	ServerPath() string
	LocalPath() string
	CreateLocalPath() bool
}

const (
	StorageNFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageNFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageNFSShared      = AllowShareForced
	StorageNFSSnapshots   = AllowSnapshotQcow2
	StorageNFSClones      = AllowCloneQcow2
)

const (
	DefaultStorageNFSVersion         = NFSVersionNone
	DefaultStorageNFSCreateLocalPath = false
)

type NFSVersion uint

const (
	NFSVersionNone NFSVersion = iota
	NFSVersion30
	NFSVersion40
	NFSVersion41
	NFSVersion42
)

func (obj NFSVersion) Marshal() (string, error) {
	switch obj {
	case NFSVersion30:
		return "3", nil
	case NFSVersion40:
		return "4", nil
	case NFSVersion41:
		return "4.1", nil
	case NFSVersion42:
		return "4.2", nil
	default:
		return "", fmt.Errorf("unknown nfs version")
	}
}

func (obj *NFSVersion) Unmarshal(s string) error {
	switch s {
	case "3":
		*obj = NFSVersion30
	case "4":
		*obj = NFSVersion40
	case "4.1":
		*obj = NFSVersion41
	case "4.2":
		*obj = NFSVersion42
	default:
		return fmt.Errorf("can't unmarshal nfs version %s", s)
	}

	return nil
}

func (obj *NFSVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
