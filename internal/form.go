package internal

import (
	"net/url"
	"reflect"
	"strconv"
)

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
