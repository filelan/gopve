package vm

import (
	"time"

	"github.com/xabinapal/gopve/pkg/request"
)

type Snapshot interface {
	Name() string
	Description() string
	Timestamp() time.Time

	WithRAM() bool

	Parent() string
	GetParent() (Snapshot, error)

	GetProperties() (SnapshotProperties, error)
	SetProperties(props SnapshotProperties) error

	Delete() error
}

type SnapshotProperties struct {
	Description string
}

func (obj SnapshotProperties) MapToValues() (request.Values, error) {
	values := request.Values{
		"description": {obj.Description},
	}

	return values, nil
}
