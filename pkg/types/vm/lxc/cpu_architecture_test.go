package lxc_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/test"
)

func TestCPUArchitecture(t *testing.T) {
	test.HelperTestFixedValue(
		t,
		(*lxc.CPUArchitecture)(nil),
		map[string](struct {
			Object types.FixedValue
			Value  string
		}){
			"I386": {
				Object: lxc.CPUArchitectureI386,
				Value:  "i386",
			},
			"AMD64": {
				Object: lxc.CPUArchitectureAMD64,
				Value:  "amd64",
			},
			"ARM64": {
				Object: lxc.CPUArchitectureARM64,
				Value:  "arm64",
			},
			"ARMHF": {
				Object: lxc.CPUArchitectureARMHF,
				Value:  "armhf",
			},
		},
	)
}
