package storage

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type getResponseJSON struct {
	Name string       `json:"storage"`
	Type storage.Kind `json:"type"`

	Content  storage.Content `json:"content"`
	Shared   types.PVEBool   `json:"shared"`
	Disabled types.PVEBool   `json:"disable"`

	ImageFormat     storage.ImageFormat `json:"format"`
	MaxBackupsPerVM uint                `json:"maxfiles"`

	Nodes types.PVEStringList `json:"nodes"`

	ExtraProperties ExtraProperties `json:"-"`
}

func (res *getResponseJSON) UnmarshalJSON(b []byte) error {
	type UnmarshalJSON getResponseJSON
	var x UnmarshalJSON
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	*res = getResponseJSON(x)
	return nil
}

func (res getResponseJSON) Map(
	svc *Service,
	full bool,
) (storage.Storage, error) {
	return NewDynamicStorage(svc, res.Name, res.Type, storage.Properties{
		Content:  res.Content,
		Shared:   res.Shared.Bool(),
		Disabled: res.Disabled.Bool(),

		ImageFormat:     res.ImageFormat,
		MaxBackupsPerVM: res.MaxBackupsPerVM,

		Nodes: res.Nodes.List(),

		ExtraProperties: res.ExtraProperties,
	})
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
