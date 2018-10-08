package service

type Storage struct {
	Storage string
	Type    string
	Content []string
}

type StorageList []*Storage
