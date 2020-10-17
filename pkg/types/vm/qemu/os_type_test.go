package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestOSType(t *testing.T) {
	test.HelperTestFixedValue(t, (*qemu.OSType)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"Other": {
			Object: qemu.OSTypeOther,
			Value:  "other",
		},
		"WindowsXP": {
			Object: qemu.OSTypeWindowsXP,
			Value:  "wxp",
		},
		"Windows2000": {
			Object: qemu.OSTypeWindows2000,
			Value:  "w2k",
		},
		"Windows2003": {
			Object: qemu.OSTypeWindows2003,
			Value:  "w2k3",
		},
		"Windows2008": {
			Object: qemu.OSTypeWindows2008,
			Value:  "w2k8",
		},
		"WindowsVista": {
			Object: qemu.OSTypeWindowsVista,
			Value:  "wvista",
		},
		"Windows7": {
			Object: qemu.OSTypeWindows7,
			Value:  "win7",
		},
		"Windows8": {
			Object: qemu.OSTypeWindows8,
			Value:  "win8",
		},
		"Windows10": {
			Object: qemu.OSTypeWindows10,
			Value:  "win10",
		},
		"Linux2.4": {
			Object: qemu.OSTypeLinux24,
			Value:  "l24",
		},
		"Linux2.6": {
			Object: qemu.OSTypeLinux26,
			Value:  "l26",
		},
		"Solaris": {
			Object: qemu.OSTypeSolaris,
			Value:  "solaris",
		},
	})
}
