package lxc_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/test"
)

func TestOSType(t *testing.T) {
	test.HelperTestFixedValue(t, (*lxc.OSType)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"Ummanaged": {
			Object: lxc.OSTypeUnmanaged,
			Value:  "unmanaged",
		},
		"Debian": {
			Object: lxc.OSTypeDebian,
			Value:  "debian",
		},
		"Ubuntu": {
			Object: lxc.OSTypeUbuntu,
			Value:  "ubuntu",
		},
		"CentOS": {
			Object: lxc.OSTypeCentOS,
			Value:  "centos",
		},
		"Fedora": {
			Object: lxc.OSTypeFedora,
			Value:  "fedora",
		},
		"OpenSUSE": {
			Object: lxc.OSTypeOpenSUSE,
			Value:  "opensuse",
		},
		"ArchLinux": {
			Object: lxc.OSTypeArchLinux,
			Value:  "archlinux",
		},
		"Alpine": {
			Object: lxc.OSTypeAlpine,
			Value:  "alpine",
		},
		"Gentoo": {
			Object: lxc.OSTypeGentoo,
			Value:  "gentoo",
		},
	})
}
