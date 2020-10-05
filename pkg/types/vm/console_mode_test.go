package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func TestQEMUConsoleMode(t *testing.T) {
	QEMUConsoleModeCases := map[string](struct {
		Object vm.QEMUConsoleMode
		Value  string
	}){
		"Shell": {
			Object: vm.QEMUConsoleModeShell,
			Value:  "shell",
		},
		"Console": {
			Object: vm.QEMUConsoleModeConsole,
			Value:  "console",
		},
		"TTY": {
			Object: vm.QEMUConsoleModeTTY,
			Value:  "tty",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range QEMUConsoleModeCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.QEMUConsoleMode
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range QEMUConsoleModeCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
