package internal

import (
	"reflect"
	"strconv"
)

type JObject = map[string]interface{}
type JArray = []interface{}

func JString(obj JObject, k string) string {
	return obj[k].(string)
}

func JStringDefault(obj JObject, k string, v string) string {
	val, ok := obj[k]
	if ok {
		return val.(string)
	} else {
		return v
	}
}

func AsJInt(obj JObject, k string) int {
	val, err := strconv.Atoi(obj[k].(string))
	if err != nil {
		panic(err)
	}
	return val
}

func JInt(obj JObject, k string) int {
	return int(obj[k].(float64))
}

func JIntDefault(obj JObject, k string, v int) int {
	val, ok := obj[k]
	if ok {
		return int(val.(float64))
	} else {
		return v
	}
}

func JFloat(obj JObject, k string) float64 {
	return obj[k].(float64)
}

func JFloatDefault(obj JObject, k string, v float64) float64 {
	val, ok := obj[k]
	if ok {
		return val.(float64)
	} else {
		return v
	}
}

func JBoolean(obj JObject, k string) bool {
	switch obj[k].(type) {
	case string:
		return obj[k].(string) == "1"
	case float64:
		return obj[k].(float64) == 1
	}
	return false
}

func JBooleanDefault(obj JObject, k string, v bool) bool {
	val, ok := obj[k]
	if ok {
		switch val.(type) {
		case string:
			return val.(string) == "1"
		case float64:
			return val.(float64) == 1
		}
		return false
	} else {
		return v
	}
}

func JSONToStruct(json JObject, s interface{}) {
	meta, ignore := GetStructMeta(s)

	for name, field := range meta {
		if field.FieldType == "dict" {
			field.Set(reflect.MakeMap(field.Type()).Interface())
			for i := field.MapMinimum; i <= field.MapMaximum; i++ {
				ai := strconv.Itoa(i)
				v, ok := json[name+ai]
				if ok {
					value := InterfaceToMetaValue(v, field, field.Type().Elem())
					field.Field.SetMapIndex(reflect.ValueOf(i), reflect.ValueOf(value))
				}
			}
		} else if field.FieldType == "kvdict" {
			field.Set(reflect.MakeMap(field.Type()).Interface())
			for i := field.MapMinimum; i <= field.MapMaximum; i++ {
				ai := strconv.Itoa(i)
				v, ok := json[name+ai]
				if ok {
					idx := reflect.New(field.Type().Elem().Elem())
					field.Field.SetMapIndex(reflect.ValueOf(i), idx)
					KVToStruct(v.(string), idx.Interface())
				}
			}
		} else if field.FieldType == "kv" {
			v, ok := json[name]
			if ok {
				KVToStruct(v.(string), field.GetPtr())
			}
		} else {
			v, ok := json[name]
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
