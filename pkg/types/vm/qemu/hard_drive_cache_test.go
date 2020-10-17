package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestHardDriveCache(t *testing.T) {
	test.HelperTestFixedValue(t, (*qemu.HardDriveCache)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"None": {
			Object: qemu.HardDriveCacheNone,
			Value:  "none",
		},
		"DirectSync": {
			Object: qemu.HardDriveCacheDirectSync,
			Value:  "directsync",
		},
		"WriteThrough": {
			Object: qemu.HardDriveCacheWriteThrough,
			Value:  "writethrough",
		},
		"WriteBack": {
			Object: qemu.HardDriveCacheWriteBack,
			Value:  "writeback",
		},
		"WriteBackUnsafe": {
			Object: qemu.HardDriveCacheWriteBackUnsafe,
			Value:  "unsafe",
		},
	})
}
