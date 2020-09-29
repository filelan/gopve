package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

var KindCases = map[string](struct {
	Object storage.Kind
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
}

func TestKindMarshal(t *testing.T) {
	for n, tt := range KindCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			var receivedKind storage.Kind
			err := (&receivedKind).Unmarshal(tt.Value)
			require.NoError(t, err)
			assert.Equal(t, tt.Object, receivedKind)
		})
	}
}

func TestKindUnmarshal(t *testing.T) {
	for n, tt := range KindCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			receivedValue, err := tt.Object.Marshal()
			require.NoError(t, err)
			assert.Equal(t, tt.Value, receivedValue)
		})
	}
}
