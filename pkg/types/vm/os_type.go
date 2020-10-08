package vm

import (
	"encoding/json"
)

type QEMUOSType string

const (
	QEMUOSTypeOther        QEMUOSType = "other"
	QEMUOSTypeWindowsXP    QEMUOSType = "wxp"
	QEMUOSTypeWindows2000  QEMUOSType = "w2k"
	QEMUOSTypeWindows2003  QEMUOSType = "w2k3"
	QEMUOSTypeWindows2008  QEMUOSType = "w2k8"
	QEMUOSTypeWindowsVista QEMUOSType = "wvista"
	QEMUOSTypeWindows7     QEMUOSType = "win7"
	QEMUOSTypeWindows8     QEMUOSType = "win8"
	QEMUOSTypeWindows10    QEMUOSType = "win10"
	QEMUOSTypeLinux24      QEMUOSType = "l24"
	QEMUOSTypeLinux26      QEMUOSType = "l26"
	QEMUOSTypeSolaris      QEMUOSType = "solaris"
)

func (obj QEMUOSType) IsValid() bool {
	switch obj {
	case
		QEMUOSTypeOther,
		QEMUOSTypeWindowsXP,
		QEMUOSTypeWindows2000,
		QEMUOSTypeWindows2003,
		QEMUOSTypeWindows2008,
		QEMUOSTypeWindowsVista,
		QEMUOSTypeWindows7,
		QEMUOSTypeWindows8,
		QEMUOSTypeWindows10,
		QEMUOSTypeLinux24,
		QEMUOSTypeLinux26,
		QEMUOSTypeSolaris:
		return true
	default:
		return false
	}
}
func (obj QEMUOSType) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUOSType) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUOSType) Unmarshal(s string) error {
	*obj = QEMUOSType(s)
	return nil
}

func (obj *QEMUOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type LXCOSType string

const (
	LXCOSTypeUnmanaged LXCOSType = "unmanaged"
	LXCOSTypeDebian    LXCOSType = "debian"
	LXCOSTypeUbuntu    LXCOSType = "ubuntu"
	LXCOSTypeCentOS    LXCOSType = "centos"
	LXCOSTypeFedora    LXCOSType = "fedora"
	LXCOSTypeOpenSUSE  LXCOSType = "opensuse"
	LXCOSTypeArchLinux LXCOSType = "archlinux"
	LXCOSTypeAlpine    LXCOSType = "alpine"
	LXCOSTypeGentoo    LXCOSType = "gentoo"
)

func (obj LXCOSType) IsValid() bool {
	switch obj {
	case
		LXCOSTypeUnmanaged,
		LXCOSTypeDebian,
		LXCOSTypeUbuntu,
		LXCOSTypeCentOS,
		LXCOSTypeFedora,
		LXCOSTypeOpenSUSE,
		LXCOSTypeArchLinux,
		LXCOSTypeAlpine,
		LXCOSTypeGentoo:
		return true
	default:
		return false
	}
}
func (obj LXCOSType) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj LXCOSType) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *LXCOSType) Unmarshal(s string) error {
	*obj = LXCOSType(s)
	return nil
}

func (obj *LXCOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
