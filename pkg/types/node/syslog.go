package node

import (
	"time"
)

type SyslogEntries map[uint]string

type GetSyslogOptions struct {
	LineStart uint
	LineLimit uint

	Since time.Time
	Until time.Time

	Service string
}
