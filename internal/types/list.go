package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PVEList struct {
	Separator string

	list []string
}

func NewPVEList(separator string, list []string) PVEList {
	return PVEList{
		Separator: separator,

		list: list,
	}
}

func (obj PVEList) Len() int {
	return len(obj.list)
}

func (obj *PVEList) Append(elem string) {
	obj.list = append(obj.list, elem)
}

func (obj PVEList) Elem(index int) string {
	return obj.list[index]
}

func (obj PVEList) List() []string {
	return obj.list
}

func (obj PVEList) Marshal() (string, error) {
	return strings.Join(obj.list, ","), nil
}

func (obj *PVEList) Unmarshal(s string) error {
	if obj.Separator == "" {
		return fmt.Errorf("can't unmarshal list, no separator defined")
	}

	obj.list = strings.Split(s, obj.Separator)

	return nil
}

func (obj *PVEList) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
