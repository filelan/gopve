package service

import (
	"strings"

	"github.com/xabinapal/gopve/internal"
)

type StorageServiceProvider interface {
	List() (StorageList, error)
}

type StorageService struct {
	client *internal.Client
}

type Storage struct {
	Storage string
	Type    string
	Content []string
}

type StorageList []Storage

func NewStorageService(c *internal.Client) *StorageService {
	storage := &StorageService{client: c}
	return storage
}

func (s *StorageService) List() (StorageList, error) {
	data, err := s.client.Get("storage")
	if err != nil {
		return nil, err
	}

	var storages StorageList
	for _, storage := range data {
		value := storage.(map[string]interface{})
		storages = append(storages, Storage{
			Storage: value["storage"].(string),
			Type:    value["type"].(string),
			Content: strings.Split(value["content"].(string), ","),
		})
	}

	return storages, nil
}
