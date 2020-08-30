package pool

import (
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/pool"
)

func (svc *Service) Create(name string, props pool.PoolProperties) (pool.Pool, error) {
	var form request.Values
	form.AddString("poolid", name)

	form.AddString("comment", props.Description)

	if err := svc.client.Request(http.MethodPost, "pools", form, nil); err != nil {
		return &Pool{}, err
	}

	return svc.Get(name)
}
