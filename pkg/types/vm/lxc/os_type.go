package lxc

import "encoding/json"

type OSType string

const (
	OSTypeUnmanaged OSType = "unmanaged"
	OSTypeDebian    OSType = "debian"
	OSTypeUbuntu    OSType = "ubuntu"
	OSTypeCentOS    OSType = "centos"
	OSTypeFedora    OSType = "fedora"
	OSTypeOpenSUSE  OSType = "opensuse"
	OSTypeArchLinux OSType = "archlinux"
	OSTypeAlpine    OSType = "alpine"
	OSTypeGentoo    OSType = "gentoo"
)

func (obj OSType) IsValid() bool {
	switch obj {
	case
		OSTypeUnmanaged,
		OSTypeDebian,
		OSTypeUbuntu,
		OSTypeCentOS,
		OSTypeFedora,
		OSTypeOpenSUSE,
		OSTypeArchLinux,
		OSTypeAlpine,
		OSTypeGentoo:
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
