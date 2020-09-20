package test

import (
	"github.com/xabinapal/gopve/internal/service/pool"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewPool() (*pool.Pool, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()
	return pool.NewPool(svc, "test_pool", "test_description"), api, exc
}
