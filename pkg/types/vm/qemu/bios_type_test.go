package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestBIOSType(t *testing.T) {
	test.HelperTestFixedValue(t, (*qemu.BIOSType)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"SeaBIOS": {
			Object: qemu.BIOSTypeSeaBIOS,
			Value:  "seabios",
		},
		"OVMF": {
			Object: qemu.BIOSTypeOVMF,
			Value:  "ovmf",
		},
	})
}
