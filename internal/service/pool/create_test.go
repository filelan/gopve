package pool_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/pool/test"
	types "github.com/xabinapal/gopve/pkg/types/pool"
)

func TestServiceCreate(t *testing.T) {
	svc, _, exc := test.NewService()

	exc.
		On("Request", http.MethodPost, "pools", url.Values{
			"poolid":  {"test_pool"},
			"comment": {"test_description"},
		}).
		Return(nil, nil).
		Once()

	err := svc.Create("test_pool", types.PoolProperties{
		Description: "test_description",
	})
	require.NoError(t, err)

	exc.AssertExpectations(t)
}
