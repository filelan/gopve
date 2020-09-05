package storage

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Storage interface {
	Name() string
	Kind() (string, error)
	Content() (StorageContent, error)
}

type StorageContent int

const (
	StorageContentQEMUData StorageContent = 1 << iota
	StorageContentContainerData
	StorageContentContainerTemplate
	StorageContentISO
	StorageContentBackup
	StorageContentSnippet
)

func (obj StorageContent) MarshalJSON() ([]byte, error) {
	var content []string

	if obj&StorageContentQEMUData != 0 {
		content = append(content, "images")
	}

	if obj&StorageContentContainerData != 0 {
		content = append(content, "rootdir")
	}

	if obj&StorageContentContainerTemplate != 0 {
		content = append(content, "vztmpl")
	}

	if obj&StorageContentBackup != 0 {
		content = append(content, "backup")
	}

	if obj&StorageContentISO != 0 {
		content = append(content, "iso")
	}

	if obj&StorageContentSnippet != 0 {
		content = append(content, "snippets")
	}

	return json.Marshal(strings.Join(content, ","))
}

func (obj *StorageContent) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	content := strings.Split(s, ",")

	for _, c := range content {
		switch c {
		case "images":
			*obj |= StorageContentQEMUData

		case "rootdir":
			*obj |= StorageContentContainerData

		case "vztmpl":
			*obj |= StorageContentContainerTemplate

		case "backup":
			*obj |= StorageContentBackup

		case "iso":
			*obj |= StorageContentISO

		case "snippets":
			*obj |= StorageContentSnippet

		default:
			return fmt.Errorf("unknown storage kind %s", c)
		}
	}

	return nil
}
