package types

import (
	"math"
	"reflect"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/schema"
)

type Properties map[string]interface{}

func (obj Properties) GetAsList(key, separator string) (types.PVEList, error) {
	prop := types.PVEList{
		Separator: separator,
	}

	var ok bool

	var v interface{}
	if v, ok = obj[key]; !ok {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return prop, err
	}

	var vv string
	if vv, ok = v.(string); !ok {
		err := errors.ErrInvalidProperty
		err.AddKey("name", key)
		err.AddKey("value", v)
		return prop, err
	}

	if err := (&prop).Unmarshal(vv); err != nil {
		return prop, err
	}

	return prop, nil
}

func (obj Properties) GetAsDict(
	key, listSeparator, keyValueSeparator string,
	allowNoValue bool,
) (types.PVEDictionary, error) {
	prop := types.PVEDictionary{
		ListSeparator:     listSeparator,
		KeyValueSeparator: keyValueSeparator,
		AllowNoValue:      allowNoValue,
	}

	var ok bool

	var v interface{}
	if v, ok = obj[key]; !ok {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return prop, err
	}

	var vv string
	if vv, ok = v.(string); !ok {
		err := errors.ErrInvalidProperty
		err.AddKey("name", key)
		err.AddKey("value", v)
		return prop, err
	}

	if err := (&prop).Unmarshal(vv); err != nil {
		return prop, err
	}

	return prop, nil
}

func (props Properties) SetInt(
	key string,
	ptr *int,
	def int,
	funcs *schema.IntFunctions,
) error {
	err := props.SetRequiredInt(key, ptr, funcs)
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		*ptr = def
		return nil
	}

	return err
}

func (props Properties) SetRequiredInt(
	key string,
	ptr *int,
	funcs *schema.IntFunctions,
) error {
	if v, ok := props[key].(float64); ok {
		if v != float64(int(v)) || v < math.MinInt32 || v > math.MaxInt32 {
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

func (props Properties) SetUint(
	key string,
	ptr *uint,
	def uint,
	funcs *schema.UintFunctions,
) error {
	err := props.SetRequiredUint(key, ptr, funcs)
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		*ptr = def
		return nil
	}

	return err
}

func (props Properties) SetRequiredUint(
	key string,
	ptr *uint,
	funcs *schema.UintFunctions,
) error {
	if v, ok := props[key].(float64); ok {
		if v != float64(int(v)) || v < 0 || v > math.MaxInt32 {
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
	funcs *schema.UintFunctions,
) error {
	err := props.SetRequiredUintFromString(key, ptr, funcs)
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		*ptr = def
		return nil
	}

	return err
}

func (props Properties) SetRequiredUintFromString(
	key string,
	ptr *uint,
	funcs *schema.UintFunctions,
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

func (props Properties) SetString(
	key string,
	ptr *string,
	def string,
	funcs *schema.StringFunctions,
) error {
	err := props.SetRequiredString(key, ptr, funcs)
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		*ptr = def
		return nil
	}

	return err
}

func (props Properties) SetRequiredString(
	key string,
	ptr *string,
	funcs *schema.StringFunctions,
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

func (props Properties) SetBool(
	key string,
	ptr *bool,
	def bool,
	funcs *schema.BoolFunctions,
) error {
	err := props.SetRequiredBool(key, ptr, funcs)
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		*ptr = def
		return nil
	}

	return err
}

func (props Properties) SetRequiredBool(
	key string,
	ptr *bool,
	funcs *schema.BoolFunctions,
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
	if err != nil && errors.ErrMissingProperty.IsBase(err) {
		reflect.ValueOf(ptr).Elem().Set(reflect.ValueOf(def))
		return nil
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
