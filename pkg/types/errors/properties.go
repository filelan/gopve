package errors

var ErrMissingProperty = NewKeyedClientError("500 - missing property!", nil)

func NewErrMissingProperty(key string) KeyedClientError {
	err := ErrMissingProperty
	err.AddKey("name", key)
	return err
}

var ErrInvalidProperty = NewKeyedClientError("500 - invalid property value!", nil)

func NewErrInvalidProperty(key string, value interface{}) KeyedClientError {
	err := ErrInvalidProperty
	err.AddKey("name", key)
	err.AddKey("value", value)
	return err
}
