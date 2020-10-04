package errors

var ErrMissingProperty = NewKeyedClientError("500 - missing property!", nil)

var ErrInvalidProperty = NewKeyedClientError("500 - invalid property value!", nil)
