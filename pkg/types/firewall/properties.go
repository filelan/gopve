package firewall

import (
	"strings"

	"github.com/xabinapal/gopve/pkg/request"
)

type GetOptions struct {
	LineStart uint
	LineLimit uint
}

const (
	DefaultMaxTrackedConnections         = 262144
	DefaultMaxConnectionEstablishTimeout = 432000
	DefaultMaxConectionSYNACKTimeout     = 60
	DefaultSYNFloodProtectionRate        = 200
	DefaultSYNFloodProtectionBurst       = 1000
)

type Properties struct {
	Enable           bool
	LogLevelIncoming LogLevel
	LogLevelOutgoing LogLevel

	LogTrackedConnections         bool
	AllowInvalidConnectionPackets bool
	MaxTrackedConnections         uint
	MaxConnectionEstablishTimeout uint
	MaxConnectionSYNACKTimeout    uint

	EnableNDP bool

	EnableSMURFS   bool
	SMURFSLogLevel LogLevel

	EnableTCPFlagsFilter   bool
	TCPFlagsFilterLogLevel LogLevel

	EnableSYNFloodProtection bool
	SYNFloodProtectionRate   uint
	SYNFloodProtectionBurst  uint

	Digest string
}

func (props Properties) MapToValues() (request.Values, error) {
	var delete []string
	form := make(request.Values)

	form.AddBool("enable", props.Enable)
	form.AddObject("log_level_in", props.LogLevelIncoming)
	form.AddObject("log_level_out", props.LogLevelOutgoing)

	form.AddBool("log_nf_conntrack", props.LogTrackedConnections)
	form.AddBool("nf_conntrack_allow_invalid", props.AllowInvalidConnectionPackets)

	if props.MaxTrackedConnections == 0 || props.MaxTrackedConnections == DefaultMaxTrackedConnections {
		delete = append(delete, "nf_conntrack_max")
	} else {
		form.AddUint("nf_conntrack_max", props.MaxTrackedConnections)
	}
	if props.MaxConnectionEstablishTimeout == 0 || props.MaxConnectionEstablishTimeout == DefaultMaxConnectionEstablishTimeout {
		delete = append(delete, "nf_conntrack_tcp_timeout_established")
	} else {
		form.AddUint("nf_conntrack_tcp_timeout_established", props.MaxConnectionEstablishTimeout)
	}
	if props.MaxConnectionSYNACKTimeout == 0 || props.MaxConnectionSYNACKTimeout == DefaultMaxConectionSYNACKTimeout {
		delete = append(delete, "nf_conntrack_tcp_timeout_syn_recv")
	} else {
		form.AddUint("nf_conntrack_tcp_timeout_syn_recv", props.MaxConnectionSYNACKTimeout)
	}

	form.AddBool("ndp", props.EnableNDP)

	form.AddBool("nosmurfs", !props.EnableSMURFS)
	form.AddObject("smurf_log_level", props.SMURFSLogLevel)

	form.AddBool("tcpflags", props.EnableTCPFlagsFilter)
	form.AddObject("tcp_flags_log_level", props.TCPFlagsFilterLogLevel)

	form.AddBool("protection_synflood", props.EnableSYNFloodProtection)
	if props.SYNFloodProtectionRate == 0 || props.SYNFloodProtectionRate == DefaultSYNFloodProtectionRate {
		delete = append(delete, "protection_synflood_rate")
	} else {
		form.AddUint("protection_synflood_rate", props.SYNFloodProtectionRate)
	}

	if props.SYNFloodProtectionBurst == 0 || props.SYNFloodProtectionBurst == DefaultSYNFloodProtectionBurst {
		delete = append(delete, "protection_synflood_burst")
	} else {
		form.AddUint("protection_synflood_burst", props.SYNFloodProtectionBurst)
	}

	form.ConditionalAddString("digest", props.Digest, props.Digest != "")

	form.ConditionalAddString("delete", strings.Join(delete, ","), len(delete) != 0)

	return form, nil
}

type LogEntries map[int]string
