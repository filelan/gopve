package internal

import "strconv"

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
	return obj[k].(float64) == 1
}

func JBooleanDefault(obj JObject, k string, v bool) bool {
	val, ok := obj[k]
	if ok {
		return val.(float64) == 1
	} else {
		return v
	}
}
