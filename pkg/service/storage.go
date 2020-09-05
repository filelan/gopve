package service

import "github.com/xabinapal/gopve/pkg/types/storage"

//go:generate mockery --case snake --name Storage

type Storage interface {
	List() ([]storage.Storage, error)
	Get(name string) (storage.Storage, error)
}
