package client_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/client"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func TestClientAtomicBlock(t *testing.T) {
	exc := new(mocks.Executor)
	cli := client.NewClientWithExecutor(exc, 0)

	exc.On("StartAtomicBlock").Return().Once()
	cli.StartAtomicBlock()
	exc.AssertExpectations(t)

	exc.On("EndAtomicBlock").Return().Once()
	cli.EndAtomicBlock()
	exc.AssertExpectations(t)
}

func TestClientRequest(t *testing.T) {
	exc := new(mocks.Executor)
	cli := client.NewClientWithExecutor(exc, 0)

	exc.
		On("Request", http.MethodGet, "/", url.Values(nil)).
		Return(nil, nil).
		Once()

	err := cli.Request(http.MethodGet, "/", nil, nil)
	require.NoError(t, err)

	exc.AssertExpectations(t)
}
