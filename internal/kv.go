package internal

import (
	"reflect"
	"strings"
)

func KVMap(kv string) map[string]string {
	meta := make(map[string]string)
	split := strings.Split(kv, ",")
	for _, field := range split {
		fieldSplit := strings.SplitN(field, "=", 2)
		if len(fieldSplit) == 1 {
			meta[""] = fieldSplit[0]
		} else {
			meta[fieldSplit[0]] = fieldSplit[1]
		}
	}
	return meta
}

func KVToStruct(kv string, s interface{}) {
	meta, _ := GetStructMeta(s)
	fields := KVMap(kv)

	for name, field := range meta {
		v, ok := fields[name]
		if !field.ValueIsSet && ok {
			if field.HasHelper {
				value := field.FieldHelper.Call([]reflect.Value{reflect.ValueOf(v)})
				field.Set(value[0].Interface())
			} else {
				value := StringToMetaValue(v, field, field.Type())
				field.Set(value)
			}
			field.ValueIsSet = true
		}
	}

	for name, field := range meta {
		_, ok := fields[name]
		if !field.ValueIsSet && !ok && field.HasDefault {
			value := StringToMetaValue(field.DefaultValue, field, field.Type())
			field.Set(value)
		}
	}

	m := reflect.ValueOf(s).MethodByName("UnmarshalHelper")
	if m.IsValid() {
		m.Call(nil)
	}
}
