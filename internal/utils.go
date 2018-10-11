package internal

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type StructMeta struct {
	Field     reflect.Value
	Interface interface{}
	Default   string
	IsSet     bool
}

func GetStructMeta(s interface{}) map[string]*StructMeta {
	var meta = make(map[string]*StructMeta)
	values := reflect.ValueOf(s).Elem()
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fieldValue := values.Field(i)
		fieldType := types.Field(i)

		n := fieldType.Tag.Get("n")
		d := fieldType.Tag.Get("d")

		field := &StructMeta{
			Field:     fieldValue,
			Interface: fieldValue.Interface(),
			Default:   d,
		}

		names := strings.Split(n, ",")
		for _, name := range names {
			meta[name] = field
		}
	}

	return meta
}

func StringToMetaValue(v string, field *StructMeta) reflect.Value {
	var value reflect.Value
	switch field.Interface.(type) {
	case string:
		value = reflect.ValueOf(v)
	case int:
		val, _ := strconv.Atoi(v)
		value = reflect.ValueOf(val)
	case bool:
		val := v == "1"
		value = reflect.ValueOf(val)
	case []int:
		val := strings.Split(v, ";")
		ival := make([]int, 0)
		for _, w := range val {
			if w != "" {
				wal, _ := strconv.Atoi(w)
				ival = append(ival, wal)
			}
		}
		fmt.Println(ival)
		value = reflect.ValueOf(ival)
	}

	return value
}

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
	meta := GetStructMeta(s)
	fields := KVMap(kv)

	for name, field := range meta {
		v, ok := fields[name]

		if field.IsSet || !ok {
			continue
		}

		value := StringToMetaValue(v, field)
		field.Field.Set(value)
		field.IsSet = true
	}

	for name, field := range meta {
		_, ok := fields[name]

		if field.IsSet || ok {
			continue
		}

		value := StringToMetaValue(field.Default, field)
		field.Field.Set(value)
	}
}

func StructToForm(s interface{}, names []string) *url.Values {
	form := &url.Values{}

	structVal := reflect.ValueOf(s).Elem()
	structType := structVal.Type()

	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		fieldType := structType.Field(i)
		ignore := fieldType.Tag.Get("i")

		if ignore == "always" {
			continue
		}

		var value string
		switch field.Interface().(type) {
		case int:
			nativeVal := field.Int()
			if nativeVal == 0 && ignore == "default" {
				continue
			}
			value = strconv.FormatInt(nativeVal, 10)
		case bool:
			nativeVal := field.Bool()
			if nativeVal == false && ignore == "default" {
				continue
			}
			if nativeVal {
				value = "1"
			} else {
				value = "0"
			}
		case string:
			value = field.String()
			if value == "" && ignore == "default" {
				continue
			}
		}

		name := getFieldName(fieldType, names)
		form.Set(name, value)
	}

	return form
}

func getFieldName(f reflect.StructField, names []string) string {
	for _, v := range names {
		tag := f.Tag.Get(v)
		if tag != "" {
			return tag
		}
	}

	return f.Name
}
