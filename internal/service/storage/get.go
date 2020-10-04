package storage

import (
	"encoding/json"
	"fmt"
	"net/http"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type getResponseJSON struct {
	Name string       `json:"storage"`
	Type storage.Kind `json:"type"`

	Content  storage.Content        `json:"content"`
	Shared   internal_types.PVEBool `json:"shared"`
	Disabled internal_types.PVEBool `json:"disable"`

	ImageFormat     storage.ImageFormat `json:"format"`
	MaxBackupsPerVM uint                `json:"maxfiles"`

	Nodes internal_types.PVEList `json:"nodes"`

	Digest string `json:"digest"`

	ExtraProperties types.Properties `json:"-"`
}

func (res *getResponseJSON) UnmarshalJSON(b []byte) error {
	type UnmarshalJSON getResponseJSON

	var x UnmarshalJSON
	x.Nodes.Separator = ","

	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	if err := json.Unmarshal(b, &x.ExtraProperties); err != nil {
		return err
	}

	props := []string{
		"storage",
		"type",
		"content",
		"shared",
		"disable",
		"format",
		"maxfiles",
		"nodes",
		"digest",
	}
	for _, prop := range props {
		delete(x.ExtraProperties, prop)
	}

	*res = getResponseJSON(x)

	return nil
}

func (res getResponseJSON) Map(
	svc *Service,
) (storage.Storage, error) {
	props := &storage.Properties{
		Content:  res.Content,
		Shared:   res.Shared.Bool(),
		Disabled: res.Disabled.Bool(),

		ImageFormat:     res.ImageFormat,
		MaxBackupsPerVM: res.MaxBackupsPerVM,

		Nodes: res.Nodes.List(),

		Digest: res.Digest,
	}

	return NewDynamicStorage(
		svc,
		res.Name,
		res.Type,
		props,
		res.ExtraProperties,
	)
}

func (svc *Service) List() ([]storage.Storage, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "storage", nil, &res); err != nil {
		return nil, err
	}

	storages := make([]storage.Storage, len(res))
	for i, storage := range res {
		out, err := storage.Map(svc)
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

	return res.Map(svc)
}
