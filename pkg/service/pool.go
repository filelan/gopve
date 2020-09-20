package service

import "github.com/xabinapal/gopve/pkg/types/pool"

//go:generate mockery --case snake --name Pool

type Pool interface {
	List() ([]pool.Pool, error)
	Get(name string) (pool.Pool, error)
	Create(name string, props pool.PoolProperties) (pool.Pool, error)
}
