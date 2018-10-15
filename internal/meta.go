package internal

import (
	"reflect"
	"strconv"
	"strings"
)

type StructMeta struct {
	Field        reflect.Value
	FieldType    string
	HasHelper    bool
	HasDefault   bool
	ValueIsSet   bool
	FieldHelper  reflect.Value
	DefaultValue string
	MapMinimum   int
	MapMaximum   int
	ArraySplit   string
}

func (meta StructMeta) Type() reflect.Type {
	return meta.Field.Type()
}

func (meta StructMeta) Get() interface{} {
	return meta.Field.Interface()
}

func (meta StructMeta) GetPtr() interface{} {
	return meta.Field.Addr().Interface()
}

func (meta StructMeta) Set(val interface{}) {
	value := reflect.ValueOf(val)
	meta.Field.Set(value)
}

type StructMetaDict = map[string]*StructMeta

func GetStructMeta(s interface{}) (StructMetaDict, []*StructMeta) {
	var meta = make(StructMetaDict)
	var ignore = make([]*StructMeta, 0)

	values := reflect.ValueOf(s).Elem()
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fieldValue := values.Field(i)
		fieldType := types.Field(i)

		t, ok := fieldType.Tag.Lookup("t")
		if !ok {
			t = "field"
		}

		d, ok := fieldType.Tag.Lookup("d")

		field := &StructMeta{
			Field:        fieldValue,
			FieldType:    t,
			HasDefault:   ok,
			DefaultValue: d,
		}

		h, ok := fieldType.Tag.Lookup("h")
		if ok && h == "true" {
			field.HasHelper = true
			field.FieldHelper = reflect.ValueOf(s).MethodByName(fieldType.Name + "Helper")
		}

		min, ok := fieldType.Tag.Lookup("min")
		if ok {
			v, _ := strconv.Atoi(min)
			field.MapMinimum = v
		}

		max, ok := fieldType.Tag.Lookup("max")
		if ok {
			v, _ := strconv.Atoi(max)
			field.MapMaximum = v
		}

		split, ok := fieldType.Tag.Lookup("s")
		if ok {
			field.ArraySplit = split
		}

		n, ok := fieldType.Tag.Lookup("n")
		if ok {
			names := strings.Split(n, ",")
			for _, name := range names {
				meta[name] = field
			}
		} else {
			ignore = append(ignore, field)
		}
	}

	return meta, ignore
}

func StringToMetaValue(v string, m *StructMeta, t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.String:
		return v

	case reflect.Int:
		val, _ := strconv.Atoi(v)
		return val

	case reflect.Int64:
		val, _ := strconv.ParseInt(v, 10, 64)
		return val

	case reflect.Float64:
		val, _ := strconv.ParseFloat(v, 64)
		return val

	case reflect.Bool:
		return v == "1"

	case reflect.Array, reflect.Slice:
		val := strings.Split(v, m.ArraySplit)
		switch t.Elem().Kind() {
		case reflect.String:
			ival := make([]string, 0)
			for _, w := range val {
				if w != "" {
					ival = append(ival, w)
				}
			}
			return ival

		case reflect.Int:
			val := strings.Split(v, m.ArraySplit)
			ival := make([]int, 0)
			for _, w := range val {
				if w != "" {
					wal, _ := strconv.Atoi(w)
					ival = append(ival, wal)
				}
			}
			return ival
		}
	}

	return nil
}

func InterfaceToMetaValue(v interface{}, m *StructMeta, t reflect.Type) interface{} {
	s, ok := v.(string)
	if ok {
		return StringToMetaValue(s, m, t)
	}

	switch t.Kind() {
	case reflect.Int:
		return int(v.(float64))

	case reflect.Int64:
		return int64(v.(float64))

	case reflect.Float64:
		return v.(float64)

	case reflect.Bool:
		return v.(float64) == 1
	}

	return nil
}
