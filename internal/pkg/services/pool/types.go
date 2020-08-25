package pool

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

type Service struct {
	utils.Service
}

func NewService(api utils.API, client utils.Client) *Service {
	return &Service{utils.Service{api, client}}
}

type Pool struct {
	svc  *Service
	full bool

	name        string
	description string
}

func (obj *Pool) Load() error {
	if obj.full {
		return nil
	}

	pool, err := obj.svc.Get(obj.name)
	if err != nil {
		return nil
	}

	obj.description, _ = pool.Description()

	return nil
}

func (obj *Pool) Name() string {
	return obj.name
}

func (obj *Pool) Description() (string, error) {
	if err := obj.Load(); err != nil {
		return "", err
	}

	return obj.description, nil
}

func (obj *Pool) GetProperties() (types.PoolProperties, error) {
	if err := obj.Load(); err != nil {
		return types.PoolProperties{}, err
	}

	return types.PoolProperties{
		Description: obj.description,
	}, nil
}

func (obj *Pool) SetProperties(props types.PoolProperties) error {
	var form utils.RequestValues
	form.AddString("comment", props.Description)

	if err := obj.svc.Client.Request(http.MethodPut, fmt.Sprintf("pools/%s", obj.name), form, nil); err != nil {
		return err
	}

	obj.description = props.Description

	return nil
}

func (obj *Pool) Delete(force bool) error {
	return fmt.Errorf("not implemented")
}
