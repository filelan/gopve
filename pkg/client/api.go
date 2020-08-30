package client

import (
	"time"

	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/internal/service/cluster"
	"github.com/xabinapal/gopve/internal/service/node"
	"github.com/xabinapal/gopve/internal/service/pool"
	"github.com/xabinapal/gopve/internal/service/storage"
	"github.com/xabinapal/gopve/internal/service/task"
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/pkg/service"
)

type API client.API

type api struct {
	cluster        service.Cluster
	node           service.Node
	pool           service.Pool
	storage        service.Storage
	virtualMachine service.VirtualMachine
	task           service.Task
}

func NewAPI(cli client.Client, poolingInterval time.Duration) API {
	api := &api{}
	api.cluster = cluster.NewService(cli, api)
	api.node = node.NewService(cli, api)
	api.pool = pool.NewService(cli, api)
	api.storage = storage.NewService(cli, api)
	api.virtualMachine = vm.NewService(cli, api)
	api.task = task.NewService(cli, api, poolingInterval)

	return api
}

func (api api) Cluster() service.Cluster {
	return api.cluster
}

func (api api) Node() service.Node {
	return api.node
}

func (api api) Pool() service.Pool {
	return api.pool
}

func (api api) Storage() service.Storage {
	return api.storage
}

func (api api) VirtualMachine() service.VirtualMachine {
	return api.virtualMachine
}

func (api api) Task() service.Task {
	return api.task
}
