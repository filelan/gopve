package client_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/xabinapal/gopve/pkg/client"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func helpClientCreate(t *testing.T, exc request.Executor) *client.Client {
	t.Helper()

	cli, err := client.NewClient(client.Config{
		Executor: exc,
	})
	if err != nil {
		t.Fatalf("Unexpected client.NewClient error: %s", err.Error())
	}

	return cli
}

func TestClientAtomicBlock(t *testing.T) {
	exc := new(mocks.Executor)
	cli := helpClientCreate(t, exc)

	exc.On("StartAtomicBlock").Once().Return()
	cli.StartAtomicBlock()
	exc.AssertExpectations(t)

	exc.On("EndAtomicBlock").Once().Return()
	cli.EndAtomicBlock()
	exc.AssertExpectations(t)
}

func TestClientRequest(t *testing.T) {
	exc := new(mocks.Executor)
	cli := helpClientCreate(t, exc)

	exc.
		On("Request", http.MethodGet, "/", url.Values(nil)).
		Once().
		Return(nil, nil)

	if err := cli.Request(http.MethodGet, "/", nil, nil); err != nil {
		t.Fatalf("Unexpected client.Client.Request error: %s", err.Error())
	}

	exc.AssertExpectations(t)
}
