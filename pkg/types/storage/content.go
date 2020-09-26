package storage

import (
	"encoding/json"
	"fmt"

	"github.com/xabinapal/gopve/internal/types"
)

type Content int

const (
	ContentNone Content = 0

	ContentUnknown Content = 1 << (iota - 1)

	// ContentQEMUData represents QEMU virtual machine image files. It's treated as the internal "images" type, which shows up as "Disk image" in the UI.
	ContentQEMUData

	// ContentContainerData represents LXC container filesystems. It's treated as the internal "rootdir" type, which shows up  as "ISO image" in the UI.
	ContentContainerData

	// ContentISO represents ISO image files. It's treated as the internal "iso" type, which shows up as "ISO files" in the UI.
	ContentISO

	// ContentContainerTemplate represents LXC container template files. It's treated as the internal "vztml" type, which shows up as  "Container template" in the UI.
	ContentContainerTemplate

	// ContentBackup represents QEMU and LXC backup files. It's treated as the internal "backup" type, which shows up as "VZDump backup file" in the UI.
	ContentBackup

	// ContentSnippet represents snippet files like guest hook scripts. It's treated as the internal "snippets" type, which shows up as "Snippets" in the UI.
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
