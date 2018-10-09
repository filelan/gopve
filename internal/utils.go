package internal

import (
	"net/url"
	"reflect"
	"strconv"
)

func BoolToForm(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func AddStructToForm(form *url.Values, s interface{}, names []string) {
	structVal := reflect.ValueOf(s).Elem()
	structType := structVal.Type()

	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		fieldType := structType.Field(i)
		ignoreDefault := fieldType.Tag.Get("i")

		var value string
		switch field.Interface().(type) {
		case int:
			nativeVal := field.Int()
			if nativeVal == 0 && ignoreDefault != "f" {
				continue
			}
			value = strconv.FormatInt(nativeVal, 10)
		case bool:
			nativeVal := field.Bool()
			if nativeVal == false && ignoreDefault != "f" {
				continue
			}
			value = BoolToForm(nativeVal)
		case string:
			value = field.String()
			if value == "" && ignoreDefault != "f" {
				continue
			}
		}

		name := getFieldName(fieldType, names)
		form.Set(name, value)
	}
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
