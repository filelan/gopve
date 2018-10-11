package internal

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type StructMeta struct {
	Field   reflect.Value
	Default string
	IsSet   bool
}

func (meta StructMeta) Get() interface{} {
	return meta.Field.Interface()
}

func (meta StructMeta) Set(val interface{}) {
	value := reflect.ValueOf(val)
	meta.Field.Set(value)
}

type StructMetaDict = map[string]*StructMeta

func GetStructMeta(s interface{}) StructMetaDict {
	var meta = make(StructMetaDict)
	values := reflect.ValueOf(s).Elem()
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fieldValue := values.Field(i)
		fieldType := types.Field(i)

		n := fieldType.Tag.Get("n")
		d := fieldType.Tag.Get("d")

		field := &StructMeta{
			Field:   fieldValue,
			Default: d,
		}

		names := strings.Split(n, ",")
		for _, name := range names {
			meta[name] = field
		}
	}

	return meta
}

func StringToMetaValue(v string, field *StructMeta) interface{} {
	switch field.Get().(type) {
	case string:
		return v
	case int:
		val, _ := strconv.Atoi(v)
		return val
	case bool:
		return v == "1"
	case []string:
		val := strings.Split(v, ";")
		ival := make([]string, 0)
		for _, w := range val {
			if w != "" {
				ival = append(ival, w)
			}
		}
		return ival
	case []int:
		val := strings.Split(v, ";")
		ival := make([]int, 0)
		for _, w := range val {
			if w != "" {
				wal, _ := strconv.Atoi(w)
				ival = append(ival, wal)
			}
		}
		return ival
	}

	return nil
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

func KVToStruct(kv string, s interface{}, helpers ...func(StructMetaDict)) {
	meta := GetStructMeta(s)
	fields := KVMap(kv)

	for name, field := range meta {
		v, ok := fields[name]
		if !field.IsSet && ok {
			value := StringToMetaValue(v, field)
			field.Set(value)
			field.IsSet = true
		}
	}

	for name, field := range meta {
		_, ok := fields[name]
		if !field.IsSet && !ok {
			value := StringToMetaValue(field.Default, field)
			field.Set(value)
		}
	}

	for _, helper := range helpers {
		helper(meta)
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
