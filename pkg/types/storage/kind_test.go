package storage_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestKind(t *testing.T) {
	test.HelperTestFixedValue(t, (*storage.Kind)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"Dir": {
			Object: storage.KindDir,
			Value:  "dir",
		},
		"LVM": {
			Object: storage.KindLVM,
			Value:  "lvm",
		},
		"LVMThin": {
			Object: storage.KindLVMThin,
			Value:  "lvmthin",
		},
		"ZFS": {
			Object: storage.KindZFS,
			Value:  "zfspool",
		},
		"NFS": {
			Object: storage.KindNFS,
			Value:  "nfs",
		},
		"CIFS": {
			Object: storage.KindCIFS,
			Value:  "cifs",
		},
		"GlusterFS": {
			Object: storage.KindGlusterFS,
			Value:  "glusterfs",
		},
		"ISCSIKernel": {
			Object: storage.KindISCSIKernel,
			Value:  "iscsi",
		},
		"ISCSIUser": {
			Object: storage.KindISCSIUser,
			Value:  "iscsidirect",
		},
		"CephFS": {
			Object: storage.KindCephFS,
			Value:  "cephfs",
		},
		"RBD": {
			Object: storage.KindRBD,
			Value:  "rbd",
		},
		"DRBD": {
			Object: storage.KindDRBD,
			Value:  "drbd",
		},
		"ZFSOverISCSI": {
			Object: storage.KindZFSOverISCSI,
			Value:  "zfs",
		},
	})
}
