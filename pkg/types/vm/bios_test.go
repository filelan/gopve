package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func TestQEMUBIOS(t *testing.T) {
	QEMUBIOSCases := map[string](struct {
		Object vm.QEMUBIOS
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
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range QEMUBIOSCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.QEMUBIOS
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range QEMUBIOSCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
