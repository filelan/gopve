package firewall

import (
	"encoding/json"
	"fmt"
)

type Direction uint

const (
	DirectionIn Direction = iota
	DirectionOut
)

func (obj Direction) Marshal() (string, error) {
	switch obj {
	case DirectionIn:
		return "in", nil

	case DirectionOut:
		return "out", nil

	default:
		return "", fmt.Errorf("unknown direction")
	}
}

func (obj *Direction) Unmarshal(s string) error {
	switch s {
	case "in":
		*obj = DirectionIn

	case "out":
		*obj = DirectionOut

	default:
		return fmt.Errorf("can't unmarshal firewall direction %s", s)
	}

	return nil
}

func (obj *Direction) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
