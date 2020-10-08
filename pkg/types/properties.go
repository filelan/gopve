package types

import (
	"reflect"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type Properties map[string]interface{}

type PropertyIntFunctions struct {
	ValidateFunc  func(int) bool
	TransformFunc func(int) int
}

func (funcs *PropertyIntFunctions) Validate(obj int) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *PropertyIntFunctions) Transform(obj int) int {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}

func (props Properties) SetInt(
	key string,
	ptr *int,
	def int,
	funcs *PropertyIntFunctions,
) error {
	err := props.SetRequiredInt(key, ptr, funcs)
	if err != nil {
		missing := errors.ErrMissingProperty
		missing.AddKey("name", key)

		if missing.Is(err) {
			*ptr = def
			return nil
		}
	}

	return err
}

func (props Properties) SetRequiredInt(
	key string,
	ptr *int,
	funcs *PropertyIntFunctions,
) error {
	if v, ok := props[key].(float64); ok {
		if v != float64(int(v)) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		val := int(v)
		if !funcs.Validate(val) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = funcs.Transform(val)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}

type PropertyUintFunctions struct {
	ValidateFunc  func(uint) bool
	TransformFunc func(uint) uint
}

func (funcs *PropertyUintFunctions) Validate(obj uint) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *PropertyUintFunctions) Transform(obj uint) uint {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}

func (props Properties) SetUint(
	key string,
	ptr *uint,
	def uint,
	funcs *PropertyUintFunctions,
) error {
	err := props.SetRequiredUint(key, ptr, funcs)
	if err != nil {
		missing := errors.ErrMissingProperty
		missing.AddKey("name", key)

		if missing.Is(err) {
			*ptr = def
			return nil
		}
	}

	return err
}

func (props Properties) SetRequiredUint(
	key string,
	ptr *uint,
	funcs *PropertyUintFunctions,
) error {
	if v, ok := props[key].(float64); ok {
		if v != float64(int(v)) || v < 0 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		val := uint(v)
		if !funcs.Validate(val) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = funcs.Transform(val)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}

func (props Properties) SetUintFromString(
	key string,
	ptr *uint,
	def uint,
	funcs *PropertyUintFunctions,
) error {

	err := props.SetRequiredUintFromString(key, ptr, funcs)
	if err != nil {
		missing := errors.ErrMissingProperty
		missing.AddKey("name", key)

		if missing.Is(err) {
			*ptr = def
			return nil
		}
	}

	return err
}

func (props Properties) SetRequiredUintFromString(
	key string,
	ptr *uint,
	funcs *PropertyUintFunctions,
) error {
	if v, ok := props[key].(string); ok {
		vv, err := strconv.Atoi(v)

		if err != nil || vv < 0 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		val := uint(vv)
		if !funcs.Validate(val) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = funcs.Transform(val)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}

type PropertyStringFunctions struct {
	ValidateFunc  func(string) bool
	TransformFunc func(string) string
}

func (funcs *PropertyStringFunctions) Validate(obj string) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *PropertyStringFunctions) Transform(obj string) string {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}

func (props Properties) SetString(
	key string,
	ptr *string,
	def string,
	funcs *PropertyStringFunctions,
) error {
	err := props.SetRequiredString(key, ptr, funcs)
	if err != nil {
		missing := errors.ErrMissingProperty
		missing.AddKey("name", key)

		if missing.Is(err) {
			*ptr = def
			return nil
		}
	}

	return err
}

func (props Properties) SetRequiredString(
	key string,
	ptr *string,
	funcs *PropertyStringFunctions,
) error {
	if v, ok := props[key].(string); ok {
		if !funcs.Validate(v) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = funcs.Transform(v)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}

type PropertyBoolFunctions struct {
	ValidateFunc  func(bool) bool
	TransformFunc func(bool) bool
}

func (funcs *PropertyBoolFunctions) Validate(obj bool) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *PropertyBoolFunctions) Transform(obj bool) bool {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}

func (props Properties) SetBool(
	key string,
	ptr *bool,
	def bool,
	funcs *PropertyBoolFunctions,
) error {
	err := props.SetRequiredBool(key, ptr, funcs)
	if err != nil {
		missing := errors.ErrMissingProperty
		missing.AddKey("name", key)

		if missing.Is(err) {
			*ptr = def
			return nil
		}
	}

	return err
}

func (props Properties) SetRequiredBool(
	key string,
	ptr *bool,
	funcs *PropertyBoolFunctions,
) error {
	if v, ok := props[key].(float64); ok {
		if v != 0 && v != 1 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		val := types.NewPVEBoolFromFloat64(v).Bool()
		if !funcs.Validate(val) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = funcs.Transform(val)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}

func (props Properties) SetFixedValue(
	key string,
	ptr FixedValuePtr,
	def FixedValue,
	validate func(FixedValue) bool,
) error {
	err := props.SetRequiredFixedValue(key, ptr, validate)
	if err != nil {
		missing := errors.ErrMissingProperty
		missing.AddKey("name", key)

		if missing.Is(err) {
			reflect.ValueOf(ptr).Elem().Set(reflect.ValueOf(def))
			return nil
		}
	}

	return err
}

func (props Properties) SetRequiredFixedValue(
	key string,
	ptr FixedValuePtr,
	validate func(FixedValue) bool,
) error {
	if v, ok := props[key].(string); ok {
		if err := ptr.Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}
