package service

import (
	"github.com/xabinapal/gopve/internal"
)

type StorageServiceProvider interface {
	List() (*StorageList, error)
	Get(storage string) (*Storage, error)
}

type StorageService struct {
	client *internal.Client
}

func NewStorageService(c *internal.Client) *StorageService {
	storage := &StorageService{client: c}
	return storage
}

func (s *StorageService) List() (*StorageList, error) {
	data, err := s.client.Get("storage")
	if err != nil {
		return nil, err
	}

	var res StorageList
	for _, storage := range data.(internal.JArray) {
		val := storage.(internal.JObject)
		row := &Storage{}
		internal.JSONToStruct(val, row)
		res = append(res, row)
	}

	return &res, nil
}

func (s *StorageService) Get(storage string) (*Storage, error) {
	data, err := s.client.Get("storage/" + storage)
	if err != nil {
		return nil, err
	}

	val := data.(internal.JObject)
	res := &Storage{}

	internal.JSONToStruct(val, res)
	return res, nil
}
