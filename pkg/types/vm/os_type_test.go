package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func TestQEMUOSType(t *testing.T) {
	QEMUOSTypeCases := map[string](struct {
		Object vm.QEMUOSType
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
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range QEMUOSTypeCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.QEMUOSType
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range QEMUOSTypeCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}

func TestLXCOSType(t *testing.T) {
	LXCOSTypeCases := map[string](struct {
		Object vm.LXCOSType
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
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range LXCOSTypeCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.LXCOSType
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range LXCOSTypeCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
