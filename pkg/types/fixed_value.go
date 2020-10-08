package types

import (
	"encoding/json"
)

type FixedValue interface {
	IsValid() bool
	IsUnknown() bool

	Marshaler
}

type FixedValuePtr interface {
	FixedValue

	Unmarshaler
	json.Unmarshaler
}
