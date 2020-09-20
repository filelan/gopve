package pool_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/pool/test"
)

func TestServiceDelete(t *testing.T) {
	svc, _, exc := test.NewService()

	exc.
		On("Request", http.MethodDelete, "pools/test_pool", url.Values(nil)).
		Return(nil, nil).
		Once()

	err := svc.Delete("test_pool")
	require.NoError(t, err)

	exc.AssertExpectations(t)
}
