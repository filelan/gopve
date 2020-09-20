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
	content storage.Content
}

func NewStorage(svc *Service, name string, kind string, content storage.Content) *Storage {
	return &Storage{
		svc: svc,

		name:    name,
		kind:    kind,
		content: content,
	}
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

func (obj *Storage) Content() (storage.Content, error) {
	if err := obj.Load(); err != nil {
		return storage.Content(0), err
	}

	return obj.content, nil
}
