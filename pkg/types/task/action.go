package task

import (
	"encoding/json"
	"fmt"
)

type Action int

const (
	ActionUnknown Action = iota
	ActionClusterCreate
	ActionQMCreate
	ActionVZCreate
)

func (obj *Action) Unmarshal(s string) error {
	switch s {
	case "qmcreate":
		*obj = ActionQMCreate

	case "vzcreate":
		*obj = ActionVZCreate

	default:
		return fmt.Errorf("unknown task kind %s", s)
	}

	return nil
}

func (obj *Action) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
