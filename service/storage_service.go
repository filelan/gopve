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
	for _, storage := range data.(internal.JArray) {
		val := storage.(internal.JObject)
		row := &Storage{
			Storage: internal.JString(val, "storage"),
			Type:    internal.JString(val, "type"),
			Content: strings.Split(internal.JString(val, "content"), ","),
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

	val := data.(internal.JObject)
	res := Storage{
		Storage: internal.JString(val, "storage"),
		Type:    internal.JString(val, "type"),
		Content: strings.Split(internal.JString(val, "content"), ","),
	}

	return &res, nil
}
