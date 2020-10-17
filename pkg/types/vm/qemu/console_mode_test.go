package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestConsoleMode(t *testing.T) {
	test.HelperTestFixedValue(
		t,
		(*qemu.ConsoleMode)(nil),
		map[string](struct {
			Object types.FixedValue
			Value  string
		}){
			"Shell": {
				Object: qemu.ConsoleModeShell,
				Value:  "shell",
			},
			"Console": {
				Object: qemu.ConsoleModeConsole,
				Value:  "console",
			},
			"TTY": {
				Object: qemu.ConsoleModeTTY,
				Value:  "tty",
			},
		},
	)
}
