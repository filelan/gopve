package internal

import (
	"reflect"
	"strconv"
	"strings"
)

type JObject = map[string]interface{}
type JArray = []interface{}

func JSONToStruct(json JObject, s interface{}) {
	meta, ignore := GetStructMeta(s)

	for name, field := range meta {
		switch field.FieldType {
		case "dict":
			field.Set(reflect.MakeMap(field.Type()).Interface())
			for i := field.MapMinimum; i <= field.MapMaximum; i++ {
				ai := strconv.Itoa(i)
				v, ok := getJsonProperty(json, name+ai)
				if ok {
					value := InterfaceToMetaValue(v, field, field.Type().Elem())
					field.Field.SetMapIndex(reflect.ValueOf(i), reflect.ValueOf(value))
				}
			}

		case "kvdict":
			field.Set(reflect.MakeMap(field.Type()).Interface())
			for i := field.MapMinimum; i <= field.MapMaximum; i++ {
				ai := strconv.Itoa(i)
				v, ok := getJsonProperty(json, name+ai)
				if ok {
					idx := reflect.New(field.Type().Elem().Elem())
					field.Field.SetMapIndex(reflect.ValueOf(i), idx)
					KVToStruct(v.(string), idx.Interface())
				}
			}

		case "kv":
			v, ok := getJsonProperty(json, name)
			if ok {
				KVToStruct(v.(string), field.GetPtr())
			}

		default:
			v, ok := getJsonProperty(json, name)
			if !ok {
				v = field.DefaultValue
			}

			if ok || field.HasDefault {
				if field.HasHelper {
					value := field.FieldHelper.Call([]reflect.Value{reflect.ValueOf(v)})
					field.Set(value[0].Interface())
				} else {
					value := InterfaceToMetaValue(v, field, field.Type())
					field.Set(value)
				}
			}
		}
	}

	for _, field := range ignore {
		if field.HasHelper {
			value := field.FieldHelper.Call(nil)
			field.Set(value[0].Interface())
		}
	}

	m := reflect.ValueOf(s).MethodByName("UnmarshalHelper")
	if m.IsValid() {
		m.Call(nil)
	}
}

func getJsonProperty(obj JObject, key string) (interface{}, bool) {
	path := strings.Split(key, ".")
	for _, v := range path[:len(path)-1] {
		obj, ok := obj[v].(JObject)
		if !ok {
			return obj, ok
		}
	}

	val, ok := obj[path[len(path)-1]]
	return val, ok
}
