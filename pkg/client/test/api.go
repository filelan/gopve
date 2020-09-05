package test

import (
	"github.com/xabinapal/gopve/pkg/service"
	"github.com/xabinapal/gopve/pkg/service/mocks"
)

type API struct {
	ClusterService        *mocks.Cluster
	NodeService           *mocks.Node
	PoolService           *mocks.Pool
	StorageService        *mocks.Storage
	VirtualMachineService *mocks.VirtualMachine
	TaskService           *mocks.Task
}

func NewAPI() *API {
	return &API{
		ClusterService:        new(mocks.Cluster),
		NodeService:           new(mocks.Node),
		PoolService:           new(mocks.Pool),
		StorageService:        new(mocks.Storage),
		VirtualMachineService: new(mocks.VirtualMachine),
		TaskService:           new(mocks.Task),
	}
}

func (api API) Cluster() service.Cluster {
	return api.ClusterService
}

func (api API) Node() service.Node {
	return api.NodeService
}

func (api API) Pool() service.Pool {
	return api.PoolService
}

func (api API) Storage() service.Storage {
	return api.StorageService
}

func (api API) VirtualMachine() service.VirtualMachine {
	return api.VirtualMachineService
}

func (api API) Task() service.Task {
	return api.TaskService
}
