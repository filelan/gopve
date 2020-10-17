package qemu

import (
	"encoding/json"
)

type OSType string

const (
	OSTypeOther        OSType = "other"
	OSTypeWindowsXP    OSType = "wxp"
	OSTypeWindows2000  OSType = "w2k"
	OSTypeWindows2003  OSType = "w2k3"
	OSTypeWindows2008  OSType = "w2k8"
	OSTypeWindowsVista OSType = "wvista"
	OSTypeWindows7     OSType = "win7"
	OSTypeWindows8     OSType = "win8"
	OSTypeWindows10    OSType = "win10"
	OSTypeLinux24      OSType = "l24"
	OSTypeLinux26      OSType = "l26"
	OSTypeSolaris      OSType = "solaris"
)

func (obj OSType) IsValid() bool {
	switch obj {
	case
		OSTypeOther,
		OSTypeWindowsXP,
		OSTypeWindows2000,
		OSTypeWindows2003,
		OSTypeWindows2008,
		OSTypeWindowsVista,
		OSTypeWindows7,
		OSTypeWindows8,
		OSTypeWindows10,
		OSTypeLinux24,
		OSTypeLinux26,
		OSTypeSolaris:
		return true
	default:
		return false
	}
}

func (obj OSType) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj OSType) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *OSType) Unmarshal(s string) error {
	*obj = OSType(s)
	return nil
}

func (obj *OSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
