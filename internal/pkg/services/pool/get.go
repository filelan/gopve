package pool

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
)

type getResponseJSON struct {
	Name    string   `json:"id"`
	Comment string   `json:"comment"`
	Members []string `json:"members"`
}

func (res getResponseJSON) ConvertToEntity(svc *Service, full bool) (types.Pool, error) {
	pool := &Pool{
		svc:  svc,
		full: full,

		name: res.Name,
	}

	if full {
		pool.description = res.Comment
	}

	return pool, nil
}

func (svc *Service) List() ([]types.Pool, error) {
	var res []getResponseJSON
	if err := svc.Client.Request(http.MethodGet, "pools", nil, &res); err != nil {
		return nil, err
	}

	pools := make([]types.Pool, len(res))
	for i, node := range res {
		out, err := node.ConvertToEntity(svc, false)
		if err != nil {
			return nil, err
		}

		pools[i] = out
	}

	return pools, nil
}

func (svc *Service) Get(name string) (types.Pool, error) {
	var res getResponseJSON
	if err := svc.Client.Request(http.MethodGet, fmt.Sprintf("pools/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.ConvertToEntity(svc, true)
}
