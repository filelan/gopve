package vm

import (
	"encoding/json"
	"fmt"
)

type QEMUOSType uint

const (
	QEMUOSTypeOther QEMUOSType = iota
	QEMUOSTypeWindowsXP
	QEMUOSTypeWindows2000
	QEMUOSTypeWindows2003
	QEMUOSTypeWindows2008
	QEMUOSTypeWindowsVista
	QEMUOSTypeWindows7
	QEMUOSTypeWindows8
	QEMUOSTypeWindows10
	QEMUOSTypeLinux24
	QEMUOSTypeLinux26
	QEMUOSTypeSolaris
)

func (obj QEMUOSType) Marshal() (string, error) {
	switch obj {
	case QEMUOSTypeOther:
		return "other", nil
	case QEMUOSTypeWindowsXP:
		return "wxp", nil
	case QEMUOSTypeWindows2000:
		return "w2k", nil
	case QEMUOSTypeWindows2003:
		return "w2k3", nil
	case QEMUOSTypeWindows2008:
		return "w2k8", nil
	case QEMUOSTypeWindowsVista:
		return "wvista", nil
	case QEMUOSTypeWindows7:
		return "win7", nil
	case QEMUOSTypeWindows8:
		return "win8", nil
	case QEMUOSTypeWindows10:
		return "win10", nil
	case QEMUOSTypeLinux24:
		return "l24", nil
	case QEMUOSTypeLinux26:
		return "l26", nil
	case QEMUOSTypeSolaris:
		return "solaris", nil
	default:
		return "", fmt.Errorf("unknown qemu os type")
	}
}

func (obj *QEMUOSType) Unmarshal(s string) error {
	switch s {
	case "other":
		*obj = QEMUOSTypeOther
	case "wxp":
		*obj = QEMUOSTypeWindowsXP
	case "w2k":
		*obj = QEMUOSTypeWindows2000
	case "w2k3":
		*obj = QEMUOSTypeWindows2003
	case "w2k8":
		*obj = QEMUOSTypeWindows2008
	case "wvista":
		*obj = QEMUOSTypeWindowsVista
	case "win7":
		*obj = QEMUOSTypeWindows7
	case "win8":
		*obj = QEMUOSTypeWindows8
	case "win10":
		*obj = QEMUOSTypeWindows10
	case "l24":
		*obj = QEMUOSTypeLinux24
	case "l26":
		*obj = QEMUOSTypeLinux26
	case "solaris":
		*obj = QEMUOSTypeSolaris
	default:
		return fmt.Errorf("can't unmarshal qemu os type %s", s)
	}

	return nil
}

func (obj *QEMUOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type LXCOSType uint

const (
	LXCOSTypeUnmanaged LXCOSType = iota
	LXCOSTypeDebian
	LXCOSTypeUbuntu
	LXCOSTypeCentOS
	LXCOSTypeFedora
	LXCOSTypeOpenSUSE
	LXCOSTypeArchLinux
	LXCOSTypeAlpine
	LXCOSTypeGentoo
)

func (obj LXCOSType) Marshal() (string, error) {
	switch obj {
	case LXCOSTypeUnmanaged:
		return "unmanaged", nil
	case LXCOSTypeDebian:
		return "debian", nil
	case LXCOSTypeUbuntu:
		return "ubuntu", nil
	case LXCOSTypeCentOS:
		return "centos", nil
	case LXCOSTypeFedora:
		return "fedora", nil
	case LXCOSTypeOpenSUSE:
		return "opensuse", nil
	case LXCOSTypeArchLinux:
		return "archlinux", nil
	case LXCOSTypeAlpine:
		return "alpine", nil
	case LXCOSTypeGentoo:
		return "gentoo", nil
	default:
		return "", fmt.Errorf("unknown lxc os type")
	}
}

func (obj *LXCOSType) Unmarshal(s string) error {
	switch s {
	case "unmanaged":
		*obj = LXCOSTypeUnmanaged
	case "debian":
		*obj = LXCOSTypeDebian
	case "ubuntu":
		*obj = LXCOSTypeUbuntu
	case "centos":
		*obj = LXCOSTypeCentOS
	case "fedora":
		*obj = LXCOSTypeFedora
	case "opensuse":
		*obj = LXCOSTypeOpenSUSE
	case "archlinux":
		*obj = LXCOSTypeArchLinux
	case "alpine":
		*obj = LXCOSTypeAlpine
	case "gentoo":
		*obj = LXCOSTypeGentoo
	default:
		return fmt.Errorf("can't unmarshal lxc os type %s", s)
	}

	return nil
}

func (obj *LXCOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
