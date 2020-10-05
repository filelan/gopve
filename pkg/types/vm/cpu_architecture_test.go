package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func TestQEMUCPUArchitecture(t *testing.T) {
	QEMUCPUArchitectureCases := map[string](struct {
		Object vm.QEMUCPUArchitecture
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
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range QEMUCPUArchitectureCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.QEMUCPUArchitecture
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range QEMUCPUArchitectureCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedObject, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedObject)
			})
		}
	})
}

func TestLXCCPUArchitecture(t *testing.T) {
	LXCCPUArchitectureCases := map[string](struct {
		Object vm.LXCCPUArchitecture
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
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range LXCCPUArchitectureCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.LXCCPUArchitecture
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range LXCCPUArchitectureCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
