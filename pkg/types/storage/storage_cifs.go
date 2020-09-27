package storage

import (
	"encoding/json"
	"fmt"
)

type StorageCIFS interface {
	Storage

	Server() string
	SMBVersion() SMBVersion

	Domain() string
	Username() string
	Password() string

	ServerShare() string
	LocalPath() string
	CreateLocalPath() bool
}

const (
	StorageCIFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCIFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageCIFSShared      = AllowShareForced
	StorageCIFSSnapshots   = AllowSnapshotQcow2
	StorageCIFSClones      = AllowCloneQcow2
)

const (
	DefaultStorageCIFSSMBVersion      = SMBVersion30
	DefaultStorageCIFSDomain          = ""
	DefaultStorageCIFSUsername        = ""
	DefaultStorageCIFSPassword        = ""
	DefaultStorageCIFSCreateLocalPath = false
)

type SMBVersion uint

const (
	SMBVersion20 SMBVersion = iota
	SMBVersion21
	SMBVersion30
)

func (obj SMBVersion) Marshal() (string, error) {
	switch obj {
	case SMBVersion20:
		return "2.0", nil
	case SMBVersion21:
		return "2.1", nil
	case SMBVersion30:
		return "3.0", nil
	default:
		return "", fmt.Errorf("unknown smb version")
	}
}

func (obj *SMBVersion) Unmarshal(s string) error {
	switch s {
	case "2.0":
		*obj = SMBVersion20
	case "2.1":
		*obj = SMBVersion21
	case "3.0":
		*obj = SMBVersion30
	default:
		return fmt.Errorf("can't unmarshal smb version %s", s)
	}

	return nil
}

func (obj *SMBVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
