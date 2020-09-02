package client_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/client"
	"github.com/xabinapal/gopve/pkg/request"
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
