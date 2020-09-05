package service

import "github.com/xabinapal/gopve/pkg/types/pool"

//go:generate mockery --case snake --name Pool

type Pool interface {
	List() ([]pool.Pool, error)
	Get(id string) (pool.Pool, error)
}
