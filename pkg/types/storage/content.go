package storage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xabinapal/gopve/internal/types"
)

type Content int

const (
	ContentQEMUData Content = 1 << iota
	ContentContainerData
	ContentContainerTemplate
	ContentISO
	ContentBackup
	ContentSnippet
)

func (obj Content) Marshal() (string, error) {
	var content []string

	if obj&ContentQEMUData != 0 {
		content = append(content, "images")
	}

	if obj&ContentContainerData != 0 {
		content = append(content, "rootdir")
	}

	if obj&ContentContainerTemplate != 0 {
		content = append(content, "vztmpl")
	}

	if obj&ContentBackup != 0 {
		content = append(content, "backup")
	}

	if obj&ContentISO != 0 {
		content = append(content, "iso")
	}

	if obj&ContentSnippet != 0 {
		content = append(content, "snippets")
	}

	return strings.Join(content, ","), nil
}

func (obj *Content) Unmarshal(s string) error {
	var content types.PVEStringList
	if err := content.Unmarshal(s); err != nil {
		return err
	}

	for _, c := range content {
		switch c {
		case "images":
			*obj |= ContentQEMUData

		case "rootdir":
			*obj |= ContentContainerData

		case "vztmpl":
			*obj |= ContentContainerTemplate

		case "backup":
			*obj |= ContentBackup

		case "iso":
			*obj |= ContentISO

		case "snippets":
			*obj |= ContentSnippet

		default:
			return fmt.Errorf("unknown storage kind %s", c)
		}
	}

	return nil
}

func (obj *Content) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
