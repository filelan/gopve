package firewall

import "fmt"

type Period uint

const (
	PeriodSecond Period = iota
	PeriodMinute
	PeriodHour
	PeriodDay
)

func (obj Period) Marshal() (string, error) {
	switch obj {
	case PeriodSecond:
		return "second", nil
	case PeriodMinute:
		return "minute", nil
	case PeriodHour:
		return "hour", nil
	case PeriodDay:
		return "day", nil
	default:
		return "", fmt.Errorf("unknown rate unit")
	}
}

func (obj *Period) Unmarshal(s string) error {
	switch s {
	case "second":
		*obj = PeriodSecond
	case "minute":
		*obj = PeriodMinute
	case "hour":
		*obj = PeriodHour
	case "day":
		*obj = PeriodDay
	default:
		return fmt.Errorf("can't unmarshal rate unit %s", s)
	}

	return nil
}
