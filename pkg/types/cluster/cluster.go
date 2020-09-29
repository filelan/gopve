package cluster

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
)

type Mode string

const (
	ModeStandalone Mode = "standalone"
	ModeCluster    Mode = "cluster"
)

type Cluster interface {
	Mode() Mode
	Name() string
}

type NodeProperties struct {
	ID    uint
	Votes uint

	Link0 NodeLink
	Link1 NodeLink
	Link2 NodeLink
	Link3 NodeLink
	Link4 NodeLink
	Link5 NodeLink
	Link6 NodeLink
	Link7 NodeLink
}

func (obj NodeProperties) MapToValues() (request.Values, error) {
	values := make(request.Values)

	values.AddUint("nodeid", obj.ID)
	values.AddUint("votes", obj.Votes)

	values.ConditionalAddObject("link0", obj.Link0, obj.Link0.Address != "")
	values.ConditionalAddObject("link1", obj.Link1, obj.Link1.Address != "")
	values.ConditionalAddObject("link2", obj.Link2, obj.Link2.Address != "")
	values.ConditionalAddObject("link3", obj.Link3, obj.Link3.Address != "")
	values.ConditionalAddObject("link4", obj.Link4, obj.Link4.Address != "")
	values.ConditionalAddObject("link5", obj.Link5, obj.Link5.Address != "")
	values.ConditionalAddObject("link6", obj.Link6, obj.Link6.Address != "")
	values.ConditionalAddObject("link7", obj.Link7, obj.Link7.Address != "")

	return values, nil
}

type NodeLink struct {
	Address  string
	Priority uint
}

func (obj NodeLink) Marshal() (string, error) {
	var content string

	content += fmt.Sprintf("address=%s", obj.Address)
	if obj.Priority != 0 {
		content += fmt.Sprintf(",priority=%d", obj.Priority)
	}

	return content, nil
}

func (obj *NodeLink) Unmarshal(s string) error {
	content := types.PVEList{Separator: ","}
	if err := content.Unmarshal(s); err != nil {
		return err
	}

	var addressIsSet bool

	for _, c := range content.List() {
		kv := types.PVEKeyValue{Separator: "=", AllowNoValue: true}
		if err := kv.Unmarshal(c); err != nil {
			return err
		}

		if !kv.HasValue() {
			if addressIsSet {
				return fmt.Errorf("can't unmarshal %s", s)
			}

			obj.Address = kv.Key()
			addressIsSet = true
		} else {
			switch kv.Key() {
			case "address":
				obj.Address = kv.Value()
				addressIsSet = true

			case "priority":
				priority, err := strconv.Atoi(kv.Value())
				if err != nil {
					return err
				}

				obj.Priority = uint(priority)

			default:
				return fmt.Errorf("unknown key %s", kv.Key())
			}
		}
	}

	return nil
}

func (obj *NodeLink) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
