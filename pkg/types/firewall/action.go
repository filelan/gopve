package firewall

import (
	"encoding/json"
	"fmt"
)

type Action uint

const (
	ActionAccept Action = iota
	ActionReject
	ActionDrop
)

func (obj Action) Marshal() (string, error) {
	switch obj {
	case ActionAccept:
		return "ACCEPT", nil
	case ActionReject:
		return "REJECT", nil
	case ActionDrop:
		return "DROP", nil
	default:
		return "", fmt.Errorf("unknown action")
	}
}

func (obj *Action) Unmarshal(s string) error {
	switch s {
	case "ACCEPT":
		*obj = ActionAccept
	case "REJECT":
		*obj = ActionReject
	case "DROP":
		*obj = ActionDrop
	default:
		return fmt.Errorf("can't unmarshal firewall action %s", s)
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
