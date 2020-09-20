package pool

import (
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/pool"
)

func (svc *Service) Create(name string, props pool.PoolProperties) error {
	form := request.Values{}
	form.AddString("poolid", name)
	form.ConditionalAddString("comment", props.Description, props.Description != "")

	if err := svc.client.Request(http.MethodPost, "pools", form, nil); err != nil {
		return err
	}

	return nil
}
