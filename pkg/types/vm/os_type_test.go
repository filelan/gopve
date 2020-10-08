package vm_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUOSType(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.QEMUOSType)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"Other": {
			Object: vm.QEMUOSTypeOther,
			Value:  "other",
		},
		"WindowsXP": {
			Object: vm.QEMUOSTypeWindowsXP,
			Value:  "wxp",
		},
		"Windows2000": {
			Object: vm.QEMUOSTypeWindows2000,
			Value:  "w2k",
		},
		"Windows2003": {
			Object: vm.QEMUOSTypeWindows2003,
			Value:  "w2k3",
		},
		"Windows2008": {
			Object: vm.QEMUOSTypeWindows2008,
			Value:  "w2k8",
		},
		"WindowsVista": {
			Object: vm.QEMUOSTypeWindowsVista,
			Value:  "wvista",
		},
		"Windows7": {
			Object: vm.QEMUOSTypeWindows7,
			Value:  "win7",
		},
		"Windows8": {
			Object: vm.QEMUOSTypeWindows8,
			Value:  "win8",
		},
		"Windows10": {
			Object: vm.QEMUOSTypeWindows10,
			Value:  "win10",
		},
		"Linux2.4": {
			Object: vm.QEMUOSTypeLinux24,
			Value:  "l24",
		},
		"Linux2.6": {
			Object: vm.QEMUOSTypeLinux26,
			Value:  "l26",
		},
		"Solaris": {
			Object: vm.QEMUOSTypeSolaris,
			Value:  "solaris",
		},
	})
}

func TestLXCOSType(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.LXCOSType)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"Ummanaged": {
			Object: vm.LXCOSTypeUnmanaged,
			Value:  "unmanaged",
		},
		"Debian": {
			Object: vm.LXCOSTypeDebian,
			Value:  "debian",
		},
		"Ubuntu": {
			Object: vm.LXCOSTypeUbuntu,
			Value:  "ubuntu",
		},
		"CentOS": {
			Object: vm.LXCOSTypeCentOS,
			Value:  "centos",
		},
		"Fedora": {
			Object: vm.LXCOSTypeFedora,
			Value:  "fedora",
		},
		"OpenSUSE": {
			Object: vm.LXCOSTypeOpenSUSE,
			Value:  "opensuse",
		},
		"ArchLinux": {
			Object: vm.LXCOSTypeArchLinux,
			Value:  "archlinux",
		},
		"Alpine": {
			Object: vm.LXCOSTypeAlpine,
			Value:  "alpine",
		},
		"Gentoo": {
			Object: vm.LXCOSTypeGentoo,
			Value:  "gentoo",
		},
	})
}
