package qemu

import (
	"encoding/json"
)

type HardDriveCache string

const (
	HardDriveCacheNone            HardDriveCache = "none"
	HardDriveCacheDirectSync      HardDriveCache = "directsync"
	HardDriveCacheWriteThrough    HardDriveCache = "writethrough"
	HardDriveCacheWriteBack       HardDriveCache = "writeback"
	HardDriveCacheWriteBackUnsafe HardDriveCache = "unsafe"
)

func (obj HardDriveCache) IsValid() bool {
	switch obj {
	case HardDriveCacheNone,
		HardDriveCacheDirectSync,
		HardDriveCacheWriteThrough,
		HardDriveCacheWriteBack,
		HardDriveCacheWriteBackUnsafe:
		return true
	default:
		return false
	}
}

func (obj HardDriveCache) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj HardDriveCache) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *HardDriveCache) Unmarshal(s string) error {
	*obj = HardDriveCache(s)
	return nil
}

func (obj *HardDriveCache) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
