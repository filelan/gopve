package firewall

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
	props := strings.Split(s, ",")

	for _, prop := range props {
		kv := strings.Split(prop, "=")
		if len(kv) == 2 {
			switch kv[0] {
			case "enable":
				if kv[1] == "0" {
					obj.Enable = false
				} else if kv[1] == "1" {
					obj.Enable = true
				} else {
					fmt.Errorf("can't unmarshal log limit %s", s)
				}
			case "rate":
				v := strings.Split(kv[1], "/")
				if len(v) == 2 {
					rateMessage, err := strconv.Atoi(v[0])
					if err != nil {
						return err
					}
					obj.RateMessages = uint(rateMessage)

					if err := (&obj.RatePeriod).Unmarshal(v[1]); err != nil {
						return err
					}
				} else {
					fmt.Errorf("can't unmarshal log limit %s", s)
				}
			case "burst":
				v, err := strconv.Atoi(kv[1])
				if err != nil {
					return err
				}

				obj.BurstMessages = uint(v)
			default:
				fmt.Errorf("can't unmarshal log limit %s", s)
			}
		} else {
			return fmt.Errorf("can't unmarshal log limit %s", s)
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
