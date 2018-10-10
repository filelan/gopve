package internal

type JObject map[string]interface{}
type JArray []interface{}

func NewJObject(data interface{}) JObject {
	return JObject(data.(map[string]interface{}))
}

func NewJArray(data interface{}) JArray {
	return JArray(data.([]interface{}))
}

func (obj JObject) GetString(k string) string {
	return obj[k].(string)
}

func (obj JObject) GetStringDefault(k string, v string) string {
	val, ok := obj[k]
	if ok {
		return val.(string)
	} else {
		return v
	}
}

func (obj JObject) GetInt(k string) int {
	return int(obj[k].(float64))
}

func (obj JObject) GetIntDefault(k string, v int) int {
	val, ok := obj[k]
	if ok {
		return int(val.(float64))
	} else {
		return v
	}
}

func (obj JObject) GetFloat(k string) float64 {
	return obj[k].(float64)
}

func (obj JObject) GetFloatDefault(k string, v float64) float64 {
	val, ok := obj[k]
	if ok {
		return val.(float64)
	} else {
		return v
	}
}

func (obj JObject) GetBool(k string) bool {
	return obj[k].(float64) == 1
}

func (obj JObject) GetBoolDefault(k string, v bool) bool {
	val, ok := obj[k]
	if ok {
		return val.(float64) == 1
	} else {
		return v
	}
}

func (obj JObject) GetJObject(k string) JObject {
	return NewJObject(obj[k])
}
