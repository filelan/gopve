package storage

import (
	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type Service struct {
	client client.Client
	api    client.API
}

func NewService(cli client.Client, api client.API) *Service {
	return &Service{
		client: cli,
		api:    api,
	}
}

type Storage struct {
	svc  *Service
	full bool

	name    string
	kind    string
	content storage.StorageContent
}

func (obj *Storage) Load() error {
	if obj.full {
		return nil
	}

	storage, err := obj.svc.Get(obj.name)
	if err != nil {
		return nil
	}

	switch x := storage.(type) {
	case *Storage:
		*obj = *x
	default:
		panic("This should never happen")
	}

	return nil
}

func (obj *Storage) Name() string {
	return obj.name
}

func (obj *Storage) Kind() (string, error) {
	if err := obj.Load(); err != nil {
		return "", err
	}

	return obj.kind, nil
}

func (obj *Storage) Content() (storage.StorageContent, error) {
	if err := obj.Load(); err != nil {
		return storage.StorageContent(0), err
	}

	return obj.content, nil
}
