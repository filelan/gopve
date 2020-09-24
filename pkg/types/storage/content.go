package storage

import (
	"encoding/json"
	"fmt"

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
	content := types.PVEStringList{Separator: ","}

	if obj&ContentQEMUData != 0 {
		content.Append("images")
	}

	if obj&ContentContainerData != 0 {
		content.Append("rootdir")
	}

	if obj&ContentContainerTemplate != 0 {
		content.Append("vztmpl")
	}

	if obj&ContentBackup != 0 {
		content.Append("backup")
	}

	if obj&ContentISO != 0 {
		content.Append("iso")
	}

	if obj&ContentSnippet != 0 {
		content.Append("snippets")
	}

	return content.Marshal()
}

func (obj *Content) Unmarshal(s string) error {
	content := types.PVEStringList{Separator: ","}
	if err := content.Unmarshal(s); err != nil {
		return err
	}

	for _, c := range content.List() {
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
