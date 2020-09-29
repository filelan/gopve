package firewall

import (
	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
)

const (
	DefaultMaxTrackedConnections         = 262144
	DefaultMaxConnectionEstablishTimeout = 432000
	DefaultMaxConectionSYNACKTimeout     = 60
	DefaultSYNFloodProtectionRate        = 200
	DefaultSYNFloodProtectionBurst       = 1000
)

type ClusterProperties struct {
	Enable         bool
	EnableEbtables bool

	DefaultInputPolicy  Action
	DefaultOutputPolicy Action

	LogLimit LogLimit

	Digest string
}

func (props ClusterProperties) MapToValues() (request.Values, error) {
	values := make(request.Values)

	values.AddBool("enable", props.Enable)
	values.AddBool("ebtables", props.EnableEbtables)

	values.AddObject("policy_in", props.DefaultInputPolicy)
	values.AddObject("policy_out", props.DefaultOutputPolicy)

	values.AddObject("log_ratelimit", props.LogLimit)

	values.ConditionalAddString("digest", props.Digest, props.Digest != "")

	return values, nil
}

type NodeProperties struct {
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

func (props NodeProperties) MapToValues() (request.Values, error) {
	delete := types.PVEList{Separator: ","}
	values := make(request.Values)

	values.AddBool("enable", props.Enable)
	values.AddObject("log_level_in", props.LogLevelIncoming)
	values.AddObject("log_level_out", props.LogLevelOutgoing)

	values.AddBool("log_nf_conntrack", props.LogTrackedConnections)
	values.AddBool(
		"nf_conntrack_allow_invalid",
		props.AllowInvalidConnectionPackets,
	)

	if props.MaxTrackedConnections == 0 ||
		props.MaxTrackedConnections == DefaultMaxTrackedConnections {
		delete.Append("nf_conntrack_max")
	} else {
		values.AddUint("nf_conntrack_max", props.MaxTrackedConnections)
	}
	if props.MaxConnectionEstablishTimeout == 0 ||
		props.MaxConnectionEstablishTimeout == DefaultMaxConnectionEstablishTimeout {
		delete.Append("nf_conntrack_tcp_timeout_established")
	} else {
		values.AddUint("nf_conntrack_tcp_timeout_established", props.MaxConnectionEstablishTimeout)
	}
	if props.MaxConnectionSYNACKTimeout == 0 ||
		props.MaxConnectionSYNACKTimeout == DefaultMaxConectionSYNACKTimeout {
		delete.Append("nf_conntrack_tcp_timeout_syn_recv")
	} else {
		values.AddUint("nf_conntrack_tcp_timeout_syn_recv", props.MaxConnectionSYNACKTimeout)
	}

	values.AddBool("ndp", props.EnableNDP)

	values.AddBool("nosmurfs", props.EnableSMURFS)
	values.AddObject("smurf_log_level", props.SMURFSLogLevel)

	values.AddBool("tcpflags", props.EnableTCPFlagsFilter)
	values.AddObject("tcp_flags_log_level", props.TCPFlagsFilterLogLevel)

	values.AddBool("protection_synflood", props.EnableSYNFloodProtection)
	if props.SYNFloodProtectionRate == 0 ||
		props.SYNFloodProtectionRate == DefaultSYNFloodProtectionRate {
		delete.Append("protection_synflood_rate")
	} else {
		values.AddUint("protection_synflood_rate", props.SYNFloodProtectionRate)
	}

	if props.SYNFloodProtectionBurst == 0 ||
		props.SYNFloodProtectionBurst == DefaultSYNFloodProtectionBurst {
		delete.Append("protection_synflood_burst")
	} else {
		values.AddUint("protection_synflood_burst", props.SYNFloodProtectionBurst)
	}

	values.ConditionalAddString("digest", props.Digest, props.Digest != "")

	values.ConditionalAddObject(
		"delete",
		delete,
		delete.Len() != 0,
	)

	return values, nil
}

type VMProperties struct {
	Enable           bool
	LogLevelIncoming LogLevel
	LogLevelOutgoing LogLevel

	DefaultInputPolicy  Action
	DefaultOutputPolicy Action

	EnableNDP       bool
	EnableRADV      bool
	EnableDHCP      bool
	EnableMACFilter bool
	EnableIPFilter  bool

	Digest string
}

func (props VMProperties) MapToValues() (request.Values, error) {
	values := make(request.Values)

	values.AddBool("enable", props.Enable)
	values.AddObject("log_level_in", props.LogLevelIncoming)
	values.AddObject("log_level_out", props.LogLevelOutgoing)

	values.AddObject("policy_in", props.DefaultInputPolicy)
	values.AddObject("policy_out", props.DefaultOutputPolicy)

	values.AddBool("ndp", props.EnableNDP)
	values.AddBool("radv", props.EnableRADV)
	values.AddBool("dhcp", props.EnableDHCP)
	values.AddBool("macfilter", props.EnableMACFilter)
	values.AddBool("ipfilter", props.EnableIPFilter)

	values.ConditionalAddString("digest", props.Digest, props.Digest != "")

	return values, nil
}
