package test

import (
	"github.com/xabinapal/gopve/internal/service/pool"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewPool() (*pool.Pool, *mocks.Executor) {
	svc, exc := NewService()
	return pool.NewPool(svc, "test_pool", "test_description"), exc
}
