package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestCPUArchitecture(t *testing.T) {
	test.HelperTestFixedValue(
		t,
		(*qemu.CPUArchitecture)(nil),
		map[string](struct {
			Object types.FixedValue
			Value  string
		}){
			"AMD64": {
				Object: qemu.CPUArchitectureAMD64,
				Value:  "x86_64",
			},
			"ARM64": {
				Object: qemu.CPUArchitectureARM64,
				Value:  "aarch64",
			},
		},
	)
}
