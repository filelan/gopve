package client_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func TestClientAtomicBlock(t *testing.T) {
	exc := new(mocks.Executor)
	cli := helpClientCreate(t, exc)

	exc.On("StartAtomicBlock").Return().Once()
	cli.StartAtomicBlock()
	exc.AssertExpectations(t)

	exc.On("EndAtomicBlock").Return().Once()
	cli.EndAtomicBlock()
	exc.AssertExpectations(t)
}

func TestClientRequest(t *testing.T) {
	exc := new(mocks.Executor)
	cli := helpClientCreate(t, exc)

	exc.
		On("Request", http.MethodGet, "/", url.Values(nil)).
		Return(nil, nil).
		Once()

	if err := cli.Request(http.MethodGet, "/", nil, nil); err != nil {
		t.Fatalf("Unexpected client.Client.Request error: %s", err.Error())
	}

	exc.AssertExpectations(t)
}
