package test

import (
	"github.com/xabinapal/gopve/internal/service/storage"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
	types "github.com/xabinapal/gopve/pkg/types/storage"
)

func NewStorage() (*storage.Storage, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()
	return storage.NewStorage(
		svc,
		"test_storage",
		"test_kind",
		types.ContentQEMUData,
	), api, exc
}
