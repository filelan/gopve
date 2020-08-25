package api

import (
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
)

func (api *API) Version() (*types.Version, error) {
	var res types.Version
	err := api.client.Request(http.MethodGet, "/version", nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
