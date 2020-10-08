package vm_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUCPUArchitecture(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.QEMUCPUArchitecture)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"X86_64": {
			Object: vm.QEMUCPUArchitectureX86_64,
			Value:  "x86_64",
		},
		"AArch64": {
			Object: vm.QEMUCPUArchitectureAArch64,
			Value:  "aarch64",
		},
	})
}

func TestLXCCPUArchitecture(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.LXCCPUArchitecture)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"AMD64": {
			Object: vm.LXCCPUArchitectureAMD64,
			Value:  "amd64",
		},
		"I386": {
			Object: vm.LXCCPUArchitectureI386,
			Value:  "i386",
		},
		"ARM64": {
			Object: vm.LXCCPUArchitectureARM64,
			Value:  "arm64",
		},
		"ARMHF": {
			Object: vm.LXCCPUArchitectureARMHF,
			Value:  "armhf",
		},
	})
}
