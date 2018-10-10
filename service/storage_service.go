package service

import (
	"strings"

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
	for _, storage := range internal.NewJArray(data) {
		val := internal.NewJObject(storage)
		row := &Storage{
			Storage: val.GetString("storage"),
			Type:    val.GetString("type"),
			Content: strings.Split(val.GetString("content"), ","),
		}

		res = append(res, row)
	}

	return &res, nil
}

func (s *StorageService) Get(storage string) (*Storage, error) {
	data, err := s.client.Get("storage/" + storage)
	if err != nil {
		return nil, err
	}

	val := internal.NewJObject(data)
	res := Storage{
		Storage: val.GetString("storage"),
		Type:    val.GetString("type"),
		Content: strings.Split(val.GetString("content"), ","),
	}

	return &res, nil
}
