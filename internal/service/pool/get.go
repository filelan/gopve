package pool

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/pool"
)

type getResponseJSON struct {
	Name        string `json:"poolid"`
	Description string `json:"comment"`
}

func (res getResponseJSON) Map(svc *Service, full bool) (pool.Pool, error) {
	pool := &Pool{
		svc:  svc,
		full: full,

		name:        res.Name,
		description: res.Description,
	}

	if full {
		// TODO
	}

	return pool, nil
}

func (svc *Service) List() ([]pool.Pool, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "pools", nil, &res); err != nil {
		return nil, err
	}

	pools := make([]pool.Pool, len(res))
	for i, node := range res {
		out, err := node.Map(svc, false)
		if err != nil {
			return nil, err
		}

		pools[i] = out
	}

	return pools, nil
}

func (svc *Service) Get(name string) (pool.Pool, error) {
	var res getResponseJSON
	if err := svc.client.Request(http.MethodGet, fmt.Sprintf("pools/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.Map(svc, true)
}
