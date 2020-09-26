package storage

import (
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type Storage struct {
	svc  *Service
	full bool

	name    string
	kind    storage.Kind
	shared  bool
	content storage.Content

	imageFormat     storage.ImageFormat
	maxBackupsPerVM uint

	nodes []string
}

func NewStorage(
	svc *Service,
	name string,
	kind storage.Kind,
	content storage.Content,
) *Storage {
	return &Storage{
		svc: svc,

		name:    name,
		kind:    kind,
		content: content,
	}
}

func NewFullStorage(
	svc *Service,
	name string,
	kind storage.Kind,
	content storage.Content,
	nodes []string,
) *Storage {
	obj := NewStorage(svc, name, kind, content)
	obj.full = true
	obj.nodes = nodes

	return obj
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

func (obj *Storage) Kind() (storage.Kind, error) {
	if err := obj.Load(); err != nil {
		return storage.KindUnknown, err
	}

	return obj.kind, nil
}

func (obj *Storage) Shared() (bool, error) {
	if err := obj.Load(); err != nil {
		return false, err
	}

	return obj.shared, nil
}

func (obj *Storage) Content() (storage.Content, error) {
	if err := obj.Load(); err != nil {
		return storage.ContentUnknown, err
	}

	return obj.content, nil
}

func (obj *Storage) ImageFormat() (storage.ImageFormat, error) {
	if err := obj.Load(); err != nil {
		return storage.ImageFormatUnknown, err
	}

	return obj.imageFormat, nil
}

func (obj *Storage) MaxBackupsPerVM() (uint, error) {
	if err := obj.Load(); err != nil {
		return 0, err
	}

	return obj.maxBackupsPerVM, nil
}

func (obj *Storage) Nodes() ([]string, error) {
	if err := obj.Load(); err != nil {
		return nil, err
	}

	return obj.nodes, nil
}
