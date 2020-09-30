package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/storage"
	types "github.com/xabinapal/gopve/pkg/types/storage"
)

func helperCreateStorage(
	kind types.Kind,
	props types.ExtraProperties,
) (types.Storage, error) {
	return storage.NewDynamicStorage(
		nil,
		"test_storage",
		kind,
		types.Properties{
			ExtraProperties: props,
		},
	)
}

func helperFilterOptionalProperties(
	props types.ExtraProperties,
	requiredProps []string,
) types.ExtraProperties {
	finalProps := make(types.ExtraProperties)
	for _, v := range requiredProps {
		finalProps[v] = props[v]
	}

	return finalProps
}

func helperTestRequiredProperties(
	t *testing.T,
	kind types.Kind,
	props types.ExtraProperties,
	requiredProps []string,
) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(types.ExtraProperties, len(props))
			for k, v := range props {
				finalProps[k] = v
			}

			delete(finalProps, prop)

			_, err := helperCreateStorage(kind, finalProps)

			expectedError := types.ErrMissingProperty
			expectedError.AddKey("name", prop)

			assert.EqualError(t, err, expectedError.Error())
		}
	}
}

func TestStorageNew(t *testing.T) {
}

func TestStorageNewDir(t *testing.T) {
	props := types.ExtraProperties{
		"path":          "test_path",
		"mkdir":         1,
		"is_mountpoint": 1,
	}

	obj, err := helperCreateStorage(types.KindDir, props)
	require.NoError(t, err)
	require.Implements(t, (*types.StorageDir)(nil), obj)

	concreteStorage := obj.(types.StorageDir)

	assert.Equal(t, "test_path", concreteStorage.LocalPath())
	assert.Equal(t, true, concreteStorage.LocalPathCreate())
	assert.Equal(t, true, concreteStorage.LocalPathIsManaged())
}

func TestStorageNewLVM(t *testing.T) {
	props := types.ExtraProperties{
		"base":                  "test_base",
		"vgname":                "test_vg",
		"saferemove":            1,
		"saferemove_throughput": 1024,
		"tagged_only":           1,
	}

	obj, err := helperCreateStorage(types.KindLVM, props)
	require.NoError(t, err)
	require.Implements(t, (*types.StorageLVM)(nil), obj)

	concreteStorage := obj.(types.StorageLVM)

	assert.Equal(t, "test_base", concreteStorage.BaseStorage())
	assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
	assert.Equal(t, true, concreteStorage.SafeRemove())
	assert.Equal(t, 1024, concreteStorage.SafeRemoveThroughput())
	assert.Equal(t, true, concreteStorage.TaggedOnly())
}

func TestStorageNewLVMThin(t *testing.T) {
	props := types.ExtraProperties{
		"thinpool": "test_pool",
		"vgname":   "test_vg",
	}

	obj, err := helperCreateStorage(types.KindLVMThin, props)
	require.NoError(t, err)
	require.Implements(t, (*types.StorageLVMThin)(nil), obj)

	concreteStorage, ok := obj.(types.StorageLVMThin)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_pool", concreteStorage.ThinPool())
	assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
}

func TestStorageNewZFS(t *testing.T) {
	props := types.ExtraProperties{
		"pool":       "test_pool",
		"blocksize":  "1024",
		"sparse":     1,
		"mountpoint": "test_mountpoint",
	}

	obj, err := helperCreateStorage(types.KindZFS, props)
	require.NoError(t, err)
	require.Implements(t, (*types.StorageZFS)(nil), obj)

	concreteStorage := obj.(types.StorageZFS)

	assert.Equal(t, "test_pool", concreteStorage.PoolName())
	assert.Equal(t, "1024", concreteStorage.BlockSize())
	assert.Equal(t, true, concreteStorage.UseSparse())
	assert.Equal(t, "test_mountpoint", concreteStorage.LocalPath())

}

func TestStorageNewNFS(t *testing.T) {
	kind := types.KindNFS

	props := types.ExtraProperties{
		"server":  "test_server",
		"options": "vers=4.2",
		"export":  "/test_export",
		"path":    "/test_path",
		"mkdir":   1,
	}

	requiredProps := []string{"server", "export", "path"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageNFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageNFS)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_server", concreteStorage.Server())
		assert.Equal(t, types.NFSVersion42, concreteStorage.NFSVersion())
		assert.Equal(t, "/test_export", concreteStorage.ServerPath())
		assert.Equal(t, "/test_path", concreteStorage.LocalPath())
		assert.Equal(t, true, concreteStorage.LocalPathCreate())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageNFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageNFS)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageNFSVersion,
			concreteStorage.NFSVersion(),
		)

		assert.Equal(
			t,
			types.DefaultStorageNFSLocalPathCreate,
			concreteStorage.LocalPathCreate(),
		)
	})
}

func TestStorageNewCIFS(t *testing.T) {
	kind := types.KindCIFS

	props := types.ExtraProperties{
		"server":     "test_server",
		"smbversion": "2.1",
		"domain":     "test_domain",
		"username":   "test_username",
		"password":   "test_password",
		"share":      "test_share",
		"path":       "/test_path",
		"mkdir":      1,
	}

	requiredProps := []string{"server", "share", "path"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageCIFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageCIFS)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_server", concreteStorage.Server())
		assert.Equal(t, types.SMBVersion21, concreteStorage.SMBVersion())
		assert.Equal(t, "test_domain", concreteStorage.Domain())
		assert.Equal(t, "test_username", concreteStorage.Username())
		assert.Equal(t, "test_password", concreteStorage.Password())
		assert.Equal(t, "test_share", concreteStorage.ServerShare())
		assert.Equal(t, "/test_path", concreteStorage.LocalPath())
		assert.Equal(t, true, concreteStorage.LocalPathCreate())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageCIFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageCIFS)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageCIFSSMBVersion,
			concreteStorage.SMBVersion(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCIFSDomain,
			concreteStorage.Domain(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCIFSUsername,
			concreteStorage.Username(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCIFSPassword,
			concreteStorage.Password(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCIFSLocalPathCreate,
			concreteStorage.LocalPathCreate(),
		)
	})
}

func TestStorageNewGlusterFS(t *testing.T) {
	kind := types.KindGlusterFS

	props := types.ExtraProperties{
		"server":    "test_server",
		"server2":   "test_backup",
		"transport": "unix",
		"volume":    "test_volume",
	}

	requiredProps := []string{"server", "volume"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageGlusterFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageGlusterFS)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_server", concreteStorage.MainServer())
		assert.Equal(t, "test_backup", concreteStorage.BackupServer())
		assert.Equal(
			t,
			types.GlusterFSTransportUNIX,
			concreteStorage.Transport(),
		)
		assert.Equal(t, "test_volume", concreteStorage.Volume())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageGlusterFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageGlusterFS)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageGlusterFSBackupServer,
			concreteStorage.BackupServer(),
		)

		assert.Equal(
			t,
			types.DefaultStorageGlusterFSTransport,
			concreteStorage.Transport(),
		)
	})
}

func TestStorageNewISCSIKernel(t *testing.T) {
	kind := types.KindISCSIKernel

	props := types.ExtraProperties{
		"portal": "test_portal",
		"target": "test_target",
	}

	requiredProps := []string{"portal", "target"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageISCSIKernel)(nil), obj)

		concreteStorage, ok := obj.(types.StorageISCSIKernel)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_portal", concreteStorage.Portal())
		assert.Equal(t, "test_target", concreteStorage.Target())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)
}

func TestStorageNewISCSIUser(t *testing.T) {
	kind := types.KindISCSIUser

	props := types.ExtraProperties{
		"portal": "test_portal",
		"target": "test_target",
	}

	requiredProps := []string{"portal", "target"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageISCSIUser)(nil), obj)

		concreteStorage, ok := obj.(types.StorageISCSIUser)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_portal", concreteStorage.Portal())
		assert.Equal(t, "test_target", concreteStorage.Target())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)
}

func TestStorageNewCephFS(t *testing.T) {
	kind := types.KindCephFS

	props := types.ExtraProperties{
		"monhost":  "test_host_1 test_host_2 test_host_3",
		"username": "test_username",
		"fuse":     1,
		"subdir":   "/test_subdir",
		"path":     "/test_path",
	}

	requiredProps := []string{"path"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageCephFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageCephFS)
		require.Equal(t, true, ok)

		assert.ElementsMatch(
			t,
			[]string{"test_host_1", "test_host_2", "test_host_3"},
			concreteStorage.MonitorHosts(),
		)
		assert.Equal(t, "test_username", concreteStorage.Username())
		assert.Equal(t, true, concreteStorage.UseFUSE())
		assert.Equal(t, "/test_subdir", concreteStorage.ServerPath())
		assert.Equal(t, "/test_path", concreteStorage.LocalPath())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageCephFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageCephFS)
		require.Equal(t, true, ok)

		assert.ElementsMatch(
			t,
			types.DefaultStorageCephFSMonitorHosts,
			concreteStorage.MonitorHosts(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCephFSUsername,
			concreteStorage.Username(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCephFSUseFUSE,
			concreteStorage.UseFUSE(),
		)

		assert.Equal(
			t,
			types.DefaultStorageCephFSServerPath,
			concreteStorage.ServerPath(),
		)
	})
}

func TestStorageNewRBD(t *testing.T) {
	kind := types.KindRBD

	props := types.ExtraProperties{
		"monhost":  "test_host_1 test_host_2 test_host_3",
		"username": "test_username",
		"krbd":     1,
		"pool":     "test_pool",
	}

	requiredProps := []string{}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageRBD)(nil), obj)

		concreteStorage, ok := obj.(types.StorageRBD)
		require.Equal(t, true, ok)

		assert.ElementsMatch(
			t,
			[]string{"test_host_1", "test_host_2", "test_host_3"},
			concreteStorage.MonitorHosts(),
		)
		assert.Equal(t, "test_username", concreteStorage.Username())
		assert.Equal(t, true, concreteStorage.UseKRBD())
		assert.Equal(t, "test_pool", concreteStorage.PoolName())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageRBD)(nil), obj)

		concreteStorage, ok := obj.(types.StorageRBD)
		require.Equal(t, true, ok)

		assert.ElementsMatch(
			t,
			types.DefaultStorageRBDMonitorHosts,
			concreteStorage.MonitorHosts(),
		)

		assert.Equal(
			t,
			types.DefaultStorageRBDUsername,
			concreteStorage.Username(),
		)

		assert.Equal(
			t,
			types.DefaultStorageRBDUseKRBD,
			concreteStorage.UseKRBD(),
		)

		assert.Equal(
			t,
			types.DefaultStorageRBDPoolName,
			concreteStorage.PoolName(),
		)
	})
}

func TestStorageNewDRBD(t *testing.T) {
	kind := types.KindDRBD

	props := types.ExtraProperties{
		"redundancy": 16,
	}

	requiredProps := []string{}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageDRBD)(nil), obj)

		concreteStorage, ok := obj.(types.StorageDRBD)
		require.Equal(t, true, ok)

		assert.Equal(t, uint(16), concreteStorage.Redundancy())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageDRBD)(nil), obj)

		concreteStorage, ok := obj.(types.StorageDRBD)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageDRBDRedundancy,
			concreteStorage.Redundancy(),
		)
	})
}

func TestStorageNewZFSOverISCSI(t *testing.T) {
	kind := types.KindZFSOverISCSI

	props := types.ExtraProperties{
		"portal":        "test_portal",
		"target":        "test_target",
		"pool":          "test_pool",
		"blocksize":     "1024",
		"sparse":        1,
		"nowritecache":  1,
		"iscsiprovider": "iet",
		"comstar_hg":    "test_comstar_hg",
		"comstar_tg":    "test_comstar_tg",
		"lio_tpg":       "test_lio_tpg",
	}

	requiredProps := []string{
		"portal",
		"target",
		"pool",
		"blocksize",
		"iscsiprovider",
	}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageZFSOverISCSI)(nil), obj)

		concreteStorage, ok := obj.(types.StorageZFSOverISCSI)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_portal", concreteStorage.Portal())
		assert.Equal(t, "test_target", concreteStorage.Target())
		assert.Equal(t, "test_pool", concreteStorage.PoolName())
		assert.Equal(t, "1024", concreteStorage.BlockSize())
		assert.Equal(t, true, concreteStorage.UseSparse())
		assert.Equal(t, false, concreteStorage.WriteCache())
		assert.Equal(t, types.ISCSIProviderIET, concreteStorage.ISCSIProvider())
		assert.Equal(t, "test_comstar_hg", concreteStorage.ComstarHostGroup())
		assert.Equal(t, "test_comstar_tg", concreteStorage.ComstarTargetGroup())
		assert.Equal(t, "test_lio_tpg", concreteStorage.LIOTargetPortalGroup())
	})

	t.Run(
		"RequiredProperties",
		helperTestRequiredProperties(t, kind, props, requiredProps),
	)

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, requiredProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageZFSOverISCSI)(nil), obj)

		concreteStorage, ok := obj.(types.StorageZFSOverISCSI)
		require.Equal(t, true, ok)

		assert.ElementsMatch(
			t,
			types.DefaultStorageZFSOverISCSIUseSparse,
			concreteStorage.UseSparse(),
		)

		assert.Equal(
			t,
			types.DefaultStorageZFSOverISCSIWriteCache,
			concreteStorage.WriteCache(),
		)

		assert.Equal(
			t,
			types.DefaultStorageZFSOverISCSIComstarHostGroup,
			concreteStorage.ComstarHostGroup(),
		)

		assert.Equal(
			t,
			types.DefaultStorageZFSOverISCSIComstarTargetGroup,
			concreteStorage.ComstarTargetGroup(),
		)

		assert.Equal(
			t,
			types.DefaultStorageZFSOverISCSILIOTargetPortalGroup,
			concreteStorage.LIOTargetPortalGroup(),
		)
	})
}
