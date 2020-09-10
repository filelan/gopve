package firewall

import (
	"encoding/json"
	"fmt"
)

type LogLevel int

const (
	LogLevelNone LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelNotice
	LogLevelWarning
	LogLevelError
	LogLevelCritical
	LogLevelAlert
	LogLevelEmergency
)

func (obj LogLevel) Marshal() (string, error) {
	switch obj {
	case LogLevelNone:
		return "nolog", nil

	case LogLevelDebug:
		return "debug", nil

	case LogLevelInfo:
		return "info", nil

	case LogLevelNotice:
		return "notice", nil

	case LogLevelWarning:
		return "warning", nil

	case LogLevelError:
		return "err", nil

	case LogLevelCritical:
		return "crit", nil

	case LogLevelAlert:
		return "alert", nil

	case LogLevelEmergency:
		return "emerg", nil

	default:
		return "", fmt.Errorf("unknown firewall level")
	}
}

func (obj *LogLevel) Unmarshal(s string) error {
	switch s {
	case "nolog":
		*obj = LogLevelNone

	case "debug":
		*obj = LogLevelDebug

	case "info":
		*obj = LogLevelInfo

	case "notice":
		*obj = LogLevelNotice

	case "warning":
		*obj = LogLevelWarning

	case "err":
		*obj = LogLevelError

	case "crit":
		*obj = LogLevelCritical

	case "alert":
		*obj = LogLevelAlert

	case "emerg":
		*obj = LogLevelEmergency

	default:
		return fmt.Errorf("can't unmarshal firewall level %s", s)
	}

	return nil
}

func (obj *LogLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
