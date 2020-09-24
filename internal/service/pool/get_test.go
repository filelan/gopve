package pool_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/pool"
	"github.com/xabinapal/gopve/internal/service/pool/test"
	types "github.com/xabinapal/gopve/pkg/types/pool"
)

func TestServiceList(t *testing.T) {
	svc, _, exc := test.NewService()

	response, err := ioutil.ReadFile("./testdata/get_pools.json")
	require.NoError(t, err)

	expectedPools := []types.Pool{
		pool.NewPool(svc, "test_pool", "test_description"),
		pool.NewPool(svc, "test_pool2", ""),
	}

	exc.
		On("Request", http.MethodGet, "pools", url.Values(nil)).
		Return(response, nil).
		Once()

	pools, err := svc.List()
	require.NoError(t, err)
	assert.ElementsMatch(t, expectedPools, pools)

	exc.AssertExpectations(t)
}

func TestServiceGet(t *testing.T) {
	svc, _, exc := test.NewService()

	response, err := ioutil.ReadFile("./testdata/get_pools_{poolid}.json")
	require.NoError(t, err)

	member1, err := pool.NewPoolMemberVirtualMachine(svc, "qemu/100")
	require.NoError(t, err)

	member2, err := pool.NewPoolMemberVirtualMachine(svc, "lxc/101")
	require.NoError(t, err)

	member3, err := pool.NewPoolMemberStorage(svc, "storage/test_node/local")
	require.NoError(t, err)

	expectedPool := pool.NewFullPool(
		svc,
		"test_pool",
		"test_description",
		[]types.PoolMember{
			member1,
			member2,
			member3,
		},
	)

	exc.
		On("Request", http.MethodGet, "pools/test_pool", url.Values(nil)).
		Return(response, nil).
		Once()

	pool, err := svc.Get("test_pool")
	require.NoError(t, err)
	assert.Equal(t, expectedPool, pool)

	exc.AssertExpectations(t)
}
