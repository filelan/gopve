package pool

import (
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

func (svc *Service) Create(name string, props types.PoolProperties) (types.Pool, error) {
	var form utils.RequestValues
	form.AddString("poolid", name)

	form.AddString("comment", props.Description)

	if err := svc.Client.Request(http.MethodPost, "pools", form, nil); err != nil {
		return &Pool{}, err
	}

	return svc.Get(name)
}
