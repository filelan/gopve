package service

import "github.com/xabinapal/gopve/pkg/types/storage"

type Storage interface {
	List() ([]storage.Storage, error)
	Get(name string) (storage.Storage, error)
}
