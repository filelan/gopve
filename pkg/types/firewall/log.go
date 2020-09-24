package firewall

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
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

type LogEntries map[int]string

type LogLimit struct {
	Enable        bool
	RateMessages  uint
	RatePeriod    Period
	BurstMessages uint
}

func (obj LogLimit) Marshal() (string, error) {
	var s string

	if obj.Enable {
		s += "enable=1"
	} else {
		s += "enable=0"
	}

	ratePeriod, err := obj.RatePeriod.Marshal()
	if err != nil {
		return "", err
	}

	s += fmt.Sprintf(",rate=%d/%s", obj.RateMessages, ratePeriod)
	s += fmt.Sprintf(",burst=%d", obj.BurstMessages)

	return s, nil
}

func (obj *LogLimit) Unmarshal(s string) error {
	props := types.PVEStringList{Separator: ","}
	if err := props.Unmarshal(s); err != nil {
		return err
	}

	for _, prop := range props.List() {
		kv := types.PVEStringKV{Separator: "=", AllowNoValue: false}
		if err := kv.Unmarshal(prop); err != nil {
			return err
		}

		switch kv.Key() {
		case "enable":
			switch kv.Value() {
			case "0":
				obj.Enable = false
			case "1":
				obj.Enable = true
			default:
				return fmt.Errorf("can't unmarshal log limit %s", s)
			}

		case "rate":
			v := types.PVEStringKV{Separator: "/", AllowNoValue: false}
			v.Unmarshal(kv.Value())

			rateMessage, err := strconv.Atoi(v.Key())
			if err != nil {
				return err
			}
			obj.RateMessages = uint(rateMessage)

			if err := (&obj.RatePeriod).Unmarshal(v.Value()); err != nil {
				return err
			}

		case "burst":
			v, err := strconv.Atoi(kv.Value())
			if err != nil {
				return err
			}

			obj.BurstMessages = uint(v)
		}
	}

	return nil
}

func (obj *LogLimit) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type GetLogOptions struct {
	LineStart uint
	LineLimit uint
}
