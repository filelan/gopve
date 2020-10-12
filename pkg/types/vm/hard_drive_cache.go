package vm

import (
	"encoding/json"
)

type QEMUHardDriveCache string

const (
	QEMUHardDriveCacheDirectSync      QEMUHardDriveCache = "directsync"
	QEMUHardDriveCacheWriteThrough    QEMUHardDriveCache = "writethrough"
	QEMUHardDriveCacheWriteBack       QEMUHardDriveCache = "writeback"
	QEMUHardDriveCacheWriteBackUnsafe QEMUHardDriveCache = "unsafe"
	QEMUHardDriveCacheNone            QEMUHardDriveCache = "none"
)

func (obj QEMUHardDriveCache) IsValid() bool {
	switch obj {
	case QEMUHardDriveCacheDirectSync,
		QEMUHardDriveCacheWriteThrough,
		QEMUHardDriveCacheWriteBack,
		QEMUHardDriveCacheWriteBackUnsafe,
		QEMUHardDriveCacheNone:
		return true
	default:
		return false
	}
}

func (obj QEMUHardDriveCache) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUHardDriveCache) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUHardDriveCache) Unmarshal(s string) error {
	*obj = QEMUHardDriveCache(s)
	return nil
}

func (obj *QEMUHardDriveCache) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
