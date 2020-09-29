package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PVEStringList struct {
	Separator string

	list []string
}

func NewPVEStringList(separator string, list []string) PVEStringList {
	return PVEStringList{
		Separator: separator,

		list: list,
	}
}

func (obj PVEStringList) Len() int {
	return len(obj.list)
}

func (obj *PVEStringList) Append(elem string) {
	obj.list = append(obj.list, elem)
}

func (obj PVEStringList) Elem(index int) string {
	return obj.list[index]
}

func (obj PVEStringList) List() []string {
	return obj.list
}

func (obj PVEStringList) Marshal() (string, error) {
	return strings.Join(obj.list, ","), nil
}

func (obj *PVEStringList) Unmarshal(s string) error {
	if obj.Separator == "" {
		return fmt.Errorf("can't unmarshal, no separator defined")
	}

	obj.list = strings.Split(s, obj.Separator)

	return nil
}

func (obj *PVEStringList) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
