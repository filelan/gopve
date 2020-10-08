package vm_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUBIOS(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.QEMUBIOS)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"SeaBIOS": {
			Object: vm.QEMUBIOSSeaBIOS,
			Value:  "seabios",
		},
		"OVMF": {
			Object: vm.QEMUBIOSOVMF,
			Value:  "ovmf",
		},
	})
}
