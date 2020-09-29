package types

import (
	"encoding/json"
)

type PVEDictionary struct {
	ListSeparator     string
	KeyValueSeparator string
	AllowNoValue      bool

	list []PVEKeyValue
}

func (obj PVEDictionary) Len() int {
	return len(obj.list)
}

func (obj *PVEDictionary) Append(elem PVEKeyValue) {
	obj.list = append(obj.list, elem)
}

func (obj PVEDictionary) Elem(index int) PVEKeyValue {
	return obj.list[index]
}

func (obj PVEDictionary) List() []PVEKeyValue {
	return obj.list
}

func (obj PVEDictionary) Marshal() (string, error) {
	list := PVEList{Separator: obj.ListSeparator}

	for _, elem := range obj.list {
		if v, err := elem.Marshal(); err == nil {
			list.Append(v)
		} else {
			return "", err
		}
	}

	return list.Marshal()
}

func (obj *PVEDictionary) Unmarshal(s string) error {
	list := PVEList{Separator: obj.ListSeparator}
	if err := (&list).Unmarshal(s); err != nil {
		return err
	}

	for _, elem := range list.List() {
		kv := PVEKeyValue{Separator: obj.KeyValueSeparator, AllowNoValue: obj.AllowNoValue}
		if err := (&kv).Unmarshal(elem); err != nil {
			return err
		}

		obj.list = append(obj.list, kv)
	}

	return nil
}

func (obj *PVEDictionary) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
