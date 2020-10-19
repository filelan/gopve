package types

import (
	"encoding/json"
	"strconv"

	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/schema"
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

func (obj PVEDictionary) Elem(key string) (string, bool) {
	for _, elem := range obj.list {
		if elem.key == key {
			return elem.value, true
		}
	}

	return "", false
}

func (obj PVEDictionary) List() []PVEKeyValue {
	return obj.list
}

func (obj PVEDictionary) InjectInt(
	key string,
	ptr *int,
	def int,
	funcs *schema.IntFunctions,
) error {
	err := obj.InjectRequiredInt(key, ptr, funcs)
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		*ptr = def
		return nil
	}

	return err
}

func (obj PVEDictionary) InjectRequiredInt(
	key string,
	ptr *int,
	funcs *schema.IntFunctions,
) error {
	if v, ok := obj.Elem(key); ok {
		if val, err := strconv.Atoi(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		} else if !funcs.Validate(val) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		} else {
			*ptr = funcs.Transform(val)
		}
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
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
		kv := PVEKeyValue{
			Separator:    obj.KeyValueSeparator,
			AllowNoValue: obj.AllowNoValue,
		}
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
