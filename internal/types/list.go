package types

import (
	"encoding/json"
	"strings"
)

type PVEStringList []string

func (obj PVEStringList) String() string {
	return strings.Join(obj, ",")
}

func (obj *PVEStringList) Unmarshal(s string) error {
	content := strings.Split(s, ",")

	for _, c := range content {
		*obj = append(*obj, c)
	}

	return nil
}

func (obj *PVEStringList) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
