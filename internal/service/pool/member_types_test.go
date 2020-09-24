package pool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/pool"
	"github.com/xabinapal/gopve/internal/service/pool/test"
	storage_test "github.com/xabinapal/gopve/internal/service/storage/test"
	vm_test "github.com/xabinapal/gopve/internal/service/vm/test"
	types "github.com/xabinapal/gopve/pkg/types/pool"
)

func TestPoolMemberVirtualMachine(t *testing.T) {
	svc, api, exc := test.NewService()

	getPoolMember := func() types.PoolMemberVirtualMachine {
		poolMember, err := pool.NewPoolMemberVirtualMachine(
			svc,
			"test_kind/100",
		)
		require.NoError(t, err)
		require.IsType(t, new(pool.PoolMemberVirtualMachine), poolMember)

		return poolMember.(types.PoolMemberVirtualMachine)
	}

	t.Run("ID", func(t *testing.T) {
		poolMember := getPoolMember()

		expectedID := "test_kind/100"
		id := poolMember.ID()
		assert.Equal(t, expectedID, id)
	})

	t.Run("Kind", func(t *testing.T) {
		poolMember := getPoolMember()

		expectedKind := types.MemberKindVirtualMachine
		kind := poolMember.Kind()
		assert.Equal(t, expectedKind, kind)
	})

	t.Run("Get", func(t *testing.T) {
		poolMember := getPoolMember()

		expectedVirtualMachine, _, _ := vm_test.NewVirtualMachine()

		api.VirtualMachineService.
			On("Get", uint(100)).
			Return(expectedVirtualMachine, nil).
			Once()

		virtualMachine, err := poolMember.Get()
		require.NoError(t, err)
		assert.Equal(t, expectedVirtualMachine, virtualMachine)

		exc.AssertExpectations(t)
	})
}

func TestPoolMemberStorage(t *testing.T) {
	svc, api, exc := test.NewService()

	getPoolMember := func() types.PoolMemberStorage {
		poolMember, err := pool.NewPoolMemberStorage(
			svc,
			"storage/test_storage",
		)
		require.NoError(t, err)
		require.IsType(t, new(pool.PoolMemberStorage), poolMember)

		return poolMember.(types.PoolMemberStorage)
	}

	t.Run("ID", func(t *testing.T) {
		poolMember := getPoolMember()

		expectedID := "storage/test_storage"
		id := poolMember.ID()
		assert.Equal(t, expectedID, id)
	})

	t.Run("Kind", func(t *testing.T) {
		poolMember := getPoolMember()

		expectedKind := types.MemberKindStorage
		kind := poolMember.Kind()
		assert.Equal(t, expectedKind, kind)
	})

	t.Run("Get", func(t *testing.T) {
		poolMember := getPoolMember()

		expectedStorage, _, _ := storage_test.NewStorage()

		api.StorageService.
			On("Get", "test_storage").
			Return(expectedStorage, nil).
			Once()

		storage, err := poolMember.Get()
		require.NoError(t, err)
		assert.Equal(t, expectedStorage, storage)

		exc.AssertExpectations(t)
	})
}
