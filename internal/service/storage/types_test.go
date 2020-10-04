package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/storage"
	"github.com/xabinapal/gopve/pkg/types"
	storage_types "github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func helperCreateStorage(
	kind storage_types.Kind,
	props types.Properties,
) (storage_types.Storage, error) {
	return storage.NewDynamicStorage(
		nil,
		"test_storage",
		kind,
		&storage_types.Properties{},
		props,
	)
}

func TestStorageNew(t *testing.T) {
}

func TestStorageDir(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"path":          "test_path",
		"mkdir":         1,
		"is_mountpoint": 1,
	})

	obj, err := helperCreateStorage(storage_types.KindDir, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageDir)(nil), obj)

	concreteStorage := obj.(storage_types.StorageDir)

	assert.Equal(t, "test_path", concreteStorage.LocalPath())
	assert.Equal(t, true, concreteStorage.LocalPathCreate())
	assert.Equal(t, true, concreteStorage.LocalPathIsManaged())
}

func TestStorageLVM(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"base":                  "test_base",
		"vgname":                "test_vg",
		"saferemove":            1,
		"saferemove_throughput": 1024,
		"tagged_only":           1,
	})

	obj, err := helperCreateStorage(storage_types.KindLVM, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageLVM)(nil), obj)

	concreteStorage := obj.(storage_types.StorageLVM)

	assert.Equal(t, "test_base", concreteStorage.BaseStorage())
	assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
	assert.Equal(t, true, concreteStorage.SafeRemove())
	assert.Equal(t, 1024, concreteStorage.SafeRemoveThroughput())
	assert.Equal(t, true, concreteStorage.TaggedOnly())
}

func TestStorageLVMThin(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"thinpool": "test_pool",
		"vgname":   "test_vg",
	})

	obj, err := helperCreateStorage(storage_types.KindLVMThin, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageLVMThin)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageLVMThin)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_pool", concreteStorage.ThinPool())
	assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
}

func TestStorageZFS(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"pool":       "test_pool",
		"blocksize":  "1024",
		"sparse":     1,
		"mountpoint": "test_mountpoint",
	})

	obj, err := helperCreateStorage(storage_types.KindZFS, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageZFS)(nil), obj)

	concreteStorage := obj.(storage_types.StorageZFS)

	assert.Equal(t, "test_pool", concreteStorage.PoolName())
	assert.Equal(t, "1024", concreteStorage.BlockSize())
	assert.Equal(t, true, concreteStorage.UseSparse())
	assert.Equal(t, "test_mountpoint", concreteStorage.LocalPath())
}

func TestStorageNFS(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"server":  "test_server",
		"options": "vers=4.2",
		"export":  "/test_export",
		"path":    "/test_path",
		"mkdir":   1,
	})

	obj, err := helperCreateStorage(storage_types.KindNFS, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageNFS)(nil), obj)

	concreteStorage := obj.(storage_types.StorageNFS)

	assert.Equal(t, "test_server", concreteStorage.Server())
	assert.Equal(t, storage_types.NFSVersion42, concreteStorage.NFSVersion())
	assert.Equal(t, "/test_export", concreteStorage.ServerPath())
	assert.Equal(t, "/test_path", concreteStorage.LocalPath())
	assert.Equal(t, true, concreteStorage.LocalPathCreate())
}

func TestStorageNewCIFS(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"server":     "test_server",
		"smbversion": "2.1",
		"domain":     "test_domain",
		"username":   "test_username",
		"password":   "test_password",
		"share":      "test_share",
		"path":       "/test_path",
		"mkdir":      1,
	})

	obj, err := helperCreateStorage(storage_types.KindCIFS, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageCIFS)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageCIFS)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_server", concreteStorage.Server())
	assert.Equal(t, storage_types.SMBVersion21, concreteStorage.SMBVersion())
	assert.Equal(t, "test_domain", concreteStorage.Domain())
	assert.Equal(t, "test_username", concreteStorage.Username())
	assert.Equal(t, "test_password", concreteStorage.Password())
	assert.Equal(t, "test_share", concreteStorage.ServerShare())
	assert.Equal(t, "/test_path", concreteStorage.LocalPath())
	assert.Equal(t, true, concreteStorage.LocalPathCreate())
}

func TestStorageNewGlusterFS(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"server":    "test_server",
		"server2":   "test_backup",
		"transport": "unix",
		"volume":    "test_volume",
	})

	obj, err := helperCreateStorage(storage_types.KindGlusterFS, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageGlusterFS)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageGlusterFS)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_server", concreteStorage.MainServer())
	assert.Equal(t, "test_backup", concreteStorage.BackupServer())
	assert.Equal(
		t,
		storage_types.GlusterFSTransportUNIX,
		concreteStorage.Transport(),
	)
	assert.Equal(t, "test_volume", concreteStorage.Volume())
}

func TestStorageNewISCSIKernel(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"portal": "test_portal",
		"target": "test_target",
	})

	obj, err := helperCreateStorage(storage_types.KindISCSIKernel, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageISCSIKernel)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageISCSIKernel)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_portal", concreteStorage.Portal())
	assert.Equal(t, "test_target", concreteStorage.Target())
}

func TestStorageNewISCSIUser(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"portal": "test_portal",
		"target": "test_target",
	})

	obj, err := helperCreateStorage(storage_types.KindISCSIUser, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageISCSIUser)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageISCSIUser)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_portal", concreteStorage.Portal())
	assert.Equal(t, "test_target", concreteStorage.Target())
}

func TestStorageNewCephFS(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"monhost":  "test_host_1 test_host_2 test_host_3",
		"username": "test_username",
		"fuse":     1,
		"subdir":   "/test_subdir",
		"path":     "/test_path",
	})

	obj, err := helperCreateStorage(storage_types.KindCephFS, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageCephFS)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageCephFS)
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
}

func TestStorageNewRBD(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"monhost":  "test_host_1 test_host_2 test_host_3",
		"username": "test_username",
		"krbd":     1,
		"pool":     "test_pool",
	})

	obj, err := helperCreateStorage(storage_types.KindRBD, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageRBD)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageRBD)
	require.Equal(t, true, ok)

	assert.ElementsMatch(
		t,
		[]string{"test_host_1", "test_host_2", "test_host_3"},
		concreteStorage.MonitorHosts(),
	)
	assert.Equal(t, "test_username", concreteStorage.Username())
	assert.Equal(t, true, concreteStorage.UseKRBD())
	assert.Equal(t, "test_pool", concreteStorage.PoolName())
}

func TestStorageNewDRBD(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"redundancy": 16,
	})

	obj, err := helperCreateStorage(storage_types.KindDRBD, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageDRBD)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageDRBD)
	require.Equal(t, true, ok)

	assert.Equal(t, uint(16), concreteStorage.Redundancy())
}

func TestStorageNewZFSOverISCSI(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
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
	})

	obj, err := helperCreateStorage(storage_types.KindZFSOverISCSI, props)
	require.NoError(t, err)
	require.Implements(t, (*storage_types.StorageZFSOverISCSI)(nil), obj)

	concreteStorage, ok := obj.(storage_types.StorageZFSOverISCSI)
	require.Equal(t, true, ok)

	assert.Equal(t, "test_portal", concreteStorage.Portal())
	assert.Equal(t, "test_target", concreteStorage.Target())
	assert.Equal(t, "test_pool", concreteStorage.PoolName())
	assert.Equal(t, "1024", concreteStorage.BlockSize())
	assert.Equal(t, true, concreteStorage.UseSparse())
	assert.Equal(t, false, concreteStorage.WriteCache())
	assert.Equal(
		t,
		storage_types.ISCSIProviderIET,
		concreteStorage.ISCSIProvider(),
	)
	assert.Equal(t, "test_comstar_hg", concreteStorage.ComstarHostGroup())
	assert.Equal(t, "test_comstar_tg", concreteStorage.ComstarTargetGroup())
	assert.Equal(t, "test_lio_tpg", concreteStorage.LIOTargetPortalGroup())
}
