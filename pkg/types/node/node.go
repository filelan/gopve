package node

import (
	"time"

	"github.com/xabinapal/gopve/pkg/types/task"
)

type Node interface {
	Name() string
	Status() Status

	Shutdown() error
	Reboot() error
	WakeOnLAN() (task.Task, error)

	GetSyslog(opts GetSyslogOptions) (SyslogEntries, error)

	GetDNSSettings() (DNSSettings, error)
	SetDNSSettings(settings DNSSettings) error
	GetHostsFile() (HostsFile, error)
	SetHostsFile(file HostsFile) error

	GetTime(local bool) (time.Time, error)
	GetTimezone() (*time.Location, error)
	SetTimezone(timezone *time.Location) error
}
