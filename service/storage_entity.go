package service

type Storage struct {
	Storage string   `n:"storage"`
	Type    string   `n:"type"`
	Content []string `n:"content" s:","`
}

type StorageList []*Storage
