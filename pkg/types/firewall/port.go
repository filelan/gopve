package firewall

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type PortRange struct {
	Start uint16
	End   uint16
}

type PortRanges []PortRange

func (obj PortRanges) Marshal() (string, error) {
	var s string

	for i, portRange := range obj {
		if i > 0 {
			s += ","
		}

		if portRange.Start == portRange.End {
			s += strconv.Itoa(int(portRange.Start))
		} else {
			s += fmt.Sprintf("%d:%d", portRange.Start, portRange.End)
		}
	}

	return s, nil
}

func (obj *PortRanges) Unmarshal(s string) error {
	var err error

	ranges := strings.Split(s, ",")
	for _, portRange := range ranges {
		var start, end int
		rangeValues := strings.Split(portRange, ":")
		if len(rangeValues) == 1 {
			start, err = strconv.Atoi(rangeValues[0])
			if err != nil {
				return err
			}

			end = start
		} else if len(rangeValues) == 2 {
			start, err = strconv.Atoi(rangeValues[0])
			if err != nil {
				return err
			}

			end, err = strconv.Atoi(rangeValues[1])
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unknown port range %s", s)
		}

		*obj = append(*obj, PortRange{
			Start: uint16(start),
			End:   uint16(end),
		})
	}

	return nil
}

func (obj *PortRanges) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
