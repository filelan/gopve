package types

import (
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type Properties map[string]interface{}

func (props Properties) SetUint(
	key string,
	ptr *uint,
	def uint,
	validate func(uint) bool,
) error {
	if v, ok := props[key].(float64); ok {
		if v != float64(int(v)) || v < 0 ||
			(validate != nil && !validate(uint(v))) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = uint(v)
	} else {
		*ptr = def
	}

	return nil
}

func (props Properties) SetRequiredUint(
	key string,
	ptr *uint,
	validate func(uint) bool,
) error {
	if v, ok := props[key].(float64); ok {
		if v != float64(int(v)) || v < 0 ||
			(validate != nil && !validate(uint(v))) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = uint(v)
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
	validate func(uint) bool,
) error {
	if v, ok := props[key].(string); ok {
		vv, err := strconv.Atoi(v)

		if err != nil || vv < 0 ||
			(validate != nil && !validate(uint(vv))) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = uint(vv)
	} else {
		*ptr = def
	}

	return nil
}

func (props Properties) SetRequiredUintFromString(
	key string,
	ptr *uint,
	validate func(uint) bool,
) error {
	if v, ok := props[key].(string); ok {
		vv, err := strconv.Atoi(v)

		if err != nil || vv < 0 ||
			(validate != nil && !validate(uint(vv))) {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = uint(vv)
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
	validate func(bool) bool,
) error {
	if v, ok := props[key].(float64); ok {
		if v != 0 && v != 1 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		*ptr = def
	}

	return nil
}

func (props Properties) SetRequiredBool(
	key string,
	ptr *bool,
	validate func(bool) bool,
) error {
	if v, ok := props[key].(float64); ok {
		if v != 0 && v != 1 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", key)
			err.AddKey("value", v)
			return err
		}

		*ptr = types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", key)
		return err
	}

	return nil
}
