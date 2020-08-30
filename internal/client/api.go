package client

import "github.com/xabinapal/gopve/pkg/service"

type API interface {
	Cluster() service.Cluster
	Node() service.Node
	Pool() service.Pool
	Storage() service.Storage
	VirtualMachine() service.VirtualMachine
	Task() service.Task
}
