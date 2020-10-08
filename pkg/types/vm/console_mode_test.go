package vm_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUConsoleMode(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.QEMUConsoleMode)(nil), map[string](struct {
		Object types.FixedValue
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
	})
}
