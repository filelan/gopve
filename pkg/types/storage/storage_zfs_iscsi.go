package storage

import (
	"encoding/json"
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageZFSOverISCSI interface {
	Storage

	Portal() string
	Target() string

	PoolName() string

	BlockSize() string
	UseSparse() bool
	WriteCache() bool

	ISCSIProvider() ISCSIProvider

	ComstarHostGroup() string
	ComstarTargetGroup() string
	LIOTargetPortalGroup() string
}

type StorageZFSOverISCSIProperties struct {
	Portal string
	Target string

	PoolName string

	BlockSize  string
	UseSparse  bool
	WriteCache bool

	ISCSIProvider ISCSIProvider

	ComstarHostGroup     string
	ComstarTargetGroup   string
	LIOTargetPortalGroup string
}

func NewStorageZFSOverISCSIProperties(
	props types.Properties,
) (*StorageZFSOverISCSIProperties, error) {
	obj := new(StorageZFSOverISCSIProperties)

	if v, ok := props["portal"].(string); ok {
		obj.Portal = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "portal")
		return nil, err
	}

	if v, ok := props["target"].(string); ok {
		obj.Target = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "target")
		return nil, err
	}

	if v, ok := props["pool"].(string); ok {
		obj.PoolName = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "pool")
		return nil, err
	}

	if v, ok := props["blocksize"].(string); ok {
		obj.BlockSize = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "blocksize")
		return nil, err
	}

	if v, ok := props["sparse"].(float64); ok {
		obj.UseSparse = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.UseSparse = DefaultStorageZFSOverISCSIUseSparse
	}

	if v, ok := props["nowritecache"].(float64); ok {
		obj.WriteCache = !internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.WriteCache = DefaultStorageZFSOverISCSIWriteCache
	}

	if v, ok := props["iscsiprovider"].(string); ok {
		if err := (&obj.ISCSIProvider).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", "iscsiprovider")
			err.AddKey("value", v)
			return nil, err
		}
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "iscsiprovider")
		return nil, err
	}

	if v, ok := props["comstar_hg"].(string); ok {
		obj.ComstarHostGroup = v
	} else {
		obj.ComstarHostGroup = DefaultStorageZFSOverISCSIComstarHostGroup
	}

	if v, ok := props["comstar_tg"].(string); ok {
		obj.ComstarTargetGroup = v
	} else {
		obj.ComstarTargetGroup = DefaultStorageZFSOverISCSIComstarTargetGroup
	}

	if v, ok := props["lio_tpg"].(string); ok {
		obj.LIOTargetPortalGroup = v
	} else {
		obj.LIOTargetPortalGroup = DefaultStorageZFSOverISCSILIOTargetPortalGroup
	}

	return obj, nil
}

const (
	StorageZFSOverISCSIKernelContents    = ContentQEMUData
	StorageZFSOverISCSIKernelImageFormat = ImageFormatRaw
	StorageZFSOverISCSIKernelShared      = AllowShareForced
	StorageZFSOverISCSIKernelSnapshots   = AllowSnapshotAll
	StorageZFSOverISCSIKernelClones      = AllowCloneAll
)

const (
	DefaultStorageZFSOverISCSIUseSparse            = false
	DefaultStorageZFSOverISCSIWriteCache           = true
	DefaultStorageZFSOverISCSIComstarHostGroup     = ""
	DefaultStorageZFSOverISCSIComstarTargetGroup   = ""
	DefaultStorageZFSOverISCSILIOTargetPortalGroup = ""
)

type ISCSIProvider uint

const (
	ISCSIProviderComstar ISCSIProvider = iota
	ISCSIProviderISTGT
	ISCSIProviderIET
	ISCSIProviderLIO
)

func (obj ISCSIProvider) Marshal() (string, error) {
	switch obj {
	case ISCSIProviderComstar:
		return "comstar", nil
	case ISCSIProviderISTGT:
		return "istgt", nil
	case ISCSIProviderIET:
		return "iet", nil
	case ISCSIProviderLIO:
		return "LIO", nil
	default:
		return "", fmt.Errorf("unknown iscsi provider")
	}
}

func (obj *ISCSIProvider) Unmarshal(s string) error {
	switch s {
	case "comstar":
		*obj = ISCSIProviderComstar
	case "istgt":
		*obj = ISCSIProviderISTGT
	case "iet":
		*obj = ISCSIProviderIET
	case "LIO":
		*obj = ISCSIProviderLIO
	default:
		return fmt.Errorf("can't unmarshal iscsi provider %s", s)
	}

	return nil
}

func (obj *ISCSIProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
