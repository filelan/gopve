package storage

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/storage"
)

type getResponseJSON struct {
	Name    string                 `json:"storage"`
	Kind    string                 `json:"type"`
	Content storage.StorageContent `json:"content"`
}

func (res getResponseJSON) Map(svc *Service, full bool) (storage.Storage, error) {
	storage := &Storage{
		svc:  svc,
		full: full,

		name:    res.Name,
		kind:    res.Kind,
		content: res.Content,
	}

	if full {
		// TODO
	}

	return storage, nil
}

func (svc *Service) List() ([]storage.Storage, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "storage", nil, &res); err != nil {
		return nil, err
	}

	storages := make([]storage.Storage, len(res))
	for i, storage := range res {
		out, err := storage.Map(svc, false)
		if err != nil {
			return nil, err
		}

		storages[i] = out
	}

	return storages, nil
}

func (svc *Service) Get(name string) (storage.Storage, error) {
	var res getResponseJSON
	if err := svc.client.Request(http.MethodGet, fmt.Sprintf("storage/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.Map(svc, true)
}
