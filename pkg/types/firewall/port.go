package firewall

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
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

	var ranges types.PVEStringList
	if err := ranges.Unmarshal(s); err != nil {
		return err
	}

	for _, portRange := range ranges {
		var start, end int

		rangeValues := types.PVEStringKV{Separator: ":", AllowNoValue: true}
		if err := rangeValues.Unmarshal(portRange); err != nil {
			return err
		}

		start, err = strconv.Atoi(rangeValues.Key())
		if err != nil {
			return err
		}

		end = start

		if rangeValues.HasValue() {
			end, err = strconv.Atoi(rangeValues.Value())
			if err != nil {
				return err
			}
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
