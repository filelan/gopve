package errors

import (
	"fmt"
	"strings"
)

type KeyedClientError struct {
	msg  string
	keys map[string]interface{}
}

func NewKeyedClientError(
	msg string,
	keys map[string]interface{},
) KeyedClientError {
	return KeyedClientError{
		msg:  msg,
		keys: keys,
	}
}

func (err KeyedClientError) Error() string {
	var b strings.Builder

	b.WriteString(err.msg)

	if err.keys != nil {
		for k, v := range err.keys {
			fmt.Fprintf(&b, " %s=%v", k, v)
		}
	}

	return b.String()
}

func (err KeyedClientError) Is(target error) bool {
	ts := target.Error()
	es := err.Error()
	return ts == es
}

func (err KeyedClientError) IsBase(target error) bool {
	if t, ok := target.(KeyedClientError); ok {
		return t.msg == err.msg
	}

	return false
}

func (err *KeyedClientError) AddKey(key string, value interface{}) {
	if err.keys == nil {
		err.keys = make(map[string]interface{})
	}

	err.keys[key] = value
}
