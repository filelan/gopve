package storage

import (
	"encoding/json"
	"fmt"
)

type ImageFormat int

const (
	ImageFormatUnknown ImageFormat = 1 << iota
	ImageFormatRaw
	ImageFormatQcow2
	ImageFormatVMDK
	ImageFormatSubVolume
)

func (obj ImageFormat) Marshal() (string, error) {
	switch obj {
	case ImageFormatRaw:
		return "raw", nil
	case ImageFormatQcow2:
		return "qcow2", nil
	case ImageFormatVMDK:
		return "vmdk", nil
	case ImageFormatSubVolume:
		return "subvol", nil

	default:
		return "", fmt.Errorf("unknown storage kind")
	}
}

func (obj *ImageFormat) Unmarshal(s string) error {
	switch s {
	case "raw":
		*obj = ImageFormatRaw
	case "qcow2":
		*obj = ImageFormatQcow2
	case "vmdk":
		*obj = ImageFormatVMDK
	case "subvol":
		*obj = ImageFormatSubVolume

	default:
		return fmt.Errorf("unknown storage image format %s", s)
	}

	return nil
}

func (obj *ImageFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
