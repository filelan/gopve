package service

import "github.com/xabinapal/gopve/pkg/types/pool"

type Pool interface {
	List() ([]pool.Pool, error)
	Get(id string) (pool.Pool, error)
}
