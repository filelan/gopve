package pool

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/pool"
)

type Service struct {
	client client.Client
	api    client.API
}

func NewService(cli client.Client, api client.API) *Service {
	return &Service{
		client: cli,
		api:    api,
	}
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

	switch x := pool.(type) {
	case *Pool:
		*obj = *x
	default:
		panic("This should never happen")
	}

	return nil
}

func (obj *Pool) Name() string {
	return obj.name
}

func (obj *Pool) Description() (string, error) {
	return obj.description, nil
}

func (obj *Pool) GetProperties() (pool.PoolProperties, error) {
	if err := obj.Load(); err != nil {
		return pool.PoolProperties{}, err
	}

	return pool.PoolProperties{
		Description: obj.description,
	}, nil
}

func (obj *Pool) SetProperties(props pool.PoolProperties) error {
	var form request.Values
	form.AddString("comment", props.Description)

	if err := obj.svc.client.Request(http.MethodPut, fmt.Sprintf("pools/%s", obj.name), form, nil); err != nil {
		return err
	}

	obj.description = props.Description

	return nil
}

func (obj *Pool) Delete(force bool) error {
	return fmt.Errorf("not implemented")
}
