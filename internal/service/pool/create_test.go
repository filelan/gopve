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

func TestServiceCreate(t *testing.T) {
	svc, exc := test.NewService()

	response, err := ioutil.ReadFile("./testdata/get_pools_{poolid}.json")
	require.NoError(t, err)

	expectedPool := pool.NewFullPool(svc, "test_pool", "test_description", []types.PoolMember{
		pool.NewPoolMemberVirtualMachine(svc, "qemu/100"),
		pool.NewPoolMemberVirtualMachine(svc, "lxc/101"),
		pool.NewPoolMemberStorage(svc, "storage/test_node/local"),
	})

	exc.
		On("Request", http.MethodPost, "pools", url.Values{
			"poolid":  {"test_pool"},
			"comment": {"test_description"},
		}).
		Return(nil, nil).
		Once()

	exc.
		On("Request", http.MethodGet, "pools/test_pool", url.Values(nil)).
		Return(response, nil).
		Once()

	pool, err := svc.Create("test_pool", types.PoolProperties{
		Description: "test_description",
	})
	require.NoError(t, err)
	assert.Equal(t, expectedPool, pool)

	exc.AssertExpectations(t)
}
