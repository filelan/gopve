package storage

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type getResponseJSON struct {
	Name     string          `json:"storage"`
	Type     storage.Kind    `json:"type"`
	Content  storage.Content `json:"content"`
	Nodes    string          `json:"nodes"`
	Disabled types.PVEBool   `json:"disable"`

	Shared types.PVEBool `json:"shared"`

	ImageFormat     storage.ImageFormat `json:"format"`
	MaxBackupsPerVM uint                `json:"maxfiles"`
}

func (res getResponseJSON) Map(
	svc *Service,
	full bool,
) (storage.Storage, error) {
	var out *Storage

	if full {
		nodes := types.PVEStringList{Separator: ","}
		if err := nodes.Unmarshal(res.Nodes); err != nil {
			return nil, err
		}

		out = NewFullStorage(svc, res.Name, res.Type, res.Content, nodes.List())
		out.nodes = nodes.List()
		return out, nil
	}

	out = NewStorage(svc, res.Name, res.Type, res.Content)
	return out, nil
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
