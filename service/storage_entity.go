package service

import "fmt"

type StorageError struct {
	Storage string
}

func (e *StorageError) Error() string {
	return fmt.Sprintf("Storage %s does not exist", e.Storage)
}

type Storage struct {
	Storage string   `n:"storage"`
	Type    string   `n:"type"`
	Content []string `n:"content" s:","`
}

type StorageList map[string]*Storage
