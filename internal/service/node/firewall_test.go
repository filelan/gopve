package node_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/node/test"
	"github.com/xabinapal/gopve/pkg/types/firewall"
)

func TestNodeFirewallProperties(t *testing.T) {
	n, exc := test.NewNode()

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_firewall_options.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/firewall/options", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedProperties := firewall.NodeProperties{
			Enable:                        true,
			LogLevelIncoming:              firewall.LogLevelInfo,
			LogLevelOutgoing:              firewall.LogLevelWarning,
			LogTrackedConnections:         true,
			AllowInvalidConnectionPackets: true,
			MaxTrackedConnections:         262144,
			MaxConnectionEstablishTimeout: 432000,
			MaxConnectionSYNACKTimeout:    60,
			EnableNDP:                     true,
			EnableSMURFS:                  true,
			SMURFSLogLevel:                firewall.LogLevelDebug,
			EnableTCPFlagsFilter:          true,
			TCPFlagsFilterLogLevel:        firewall.LogLevelCritical,
			EnableSYNFloodProtection:      true,
			SYNFloodProtectionRate:        200,
			SYNFloodProtectionBurst:       1000,
			Digest:                        "0000000000000000000000000000000000000000",
		}

		properties, err := n.GetFirewallProperties()
		require.NoError(t, err)
		assert.Equal(t, expectedProperties, properties)

		exc.AssertExpectations(t)
	})

	t.Run("Set", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/firewall/options", url.Values{
				"enable":                               {"1"},
				"log_level_in":                         {"info"},
				"log_level_out":                        {"warning"},
				"log_nf_conntrack":                     {"1"},
				"nf_conntrack_allow_invalid":           {"1"},
				"nf_conntrack_max":                     {"131072"},
				"nf_conntrack_tcp_timeout_established": {"216000"},
				"nf_conntrack_tcp_timeout_syn_recv":    {"30"},
				"ndp":                                  {"1"},
				"nosmurfs":                             {"1"},
				"smurf_log_level":                      {"debug"},
				"tcpflags":                             {"1"},
				"tcp_flags_log_level":                  {"crit"},
				"protection_synflood":                  {"1"},
				"protection_synflood_rate":             {"100"},
				"protection_synflood_burst":            {"500"},
				"digest":                               {"0000000000000000000000000000000000000000"},
			}).
			Return(nil, nil).
			Once()

		err := n.SetFirewallProperties(firewall.NodeProperties{
			Enable:                        true,
			LogLevelIncoming:              firewall.LogLevelInfo,
			LogLevelOutgoing:              firewall.LogLevelWarning,
			LogTrackedConnections:         true,
			AllowInvalidConnectionPackets: true,
			MaxTrackedConnections:         131072,
			MaxConnectionEstablishTimeout: 216000,
			MaxConnectionSYNACKTimeout:    30,
			EnableNDP:                     true,
			EnableSMURFS:                  true,
			SMURFSLogLevel:                firewall.LogLevelDebug,
			EnableTCPFlagsFilter:          true,
			TCPFlagsFilterLogLevel:        firewall.LogLevelCritical,
			EnableSYNFloodProtection:      true,
			SYNFloodProtectionRate:        100,
			SYNFloodProtectionBurst:       500,
			Digest:                        "0000000000000000000000000000000000000000",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}

func TestNodeFirewallRules(t *testing.T) {
	n, exc := test.NewNode()

	t.Run("List", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_firewall_rules.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/firewall/rules", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedRules := []firewall.Rule{
			{
				Enable:             true,
				Description:        "test_rule_1",
				SecurityGroup:      "",
				Interface:          "eth0",
				Direction:          firewall.DirectionIn,
				Action:             firewall.ActionAccept,
				SourceAddress:      "0.0.0.0/0",
				DestinationAddress: "10.0.0.0-10.0.0.255",
				Macro:              firewall.MacroNone,
				Protocol:           firewall.ProtocolTCP,
				SourcePorts: []firewall.PortRange{
					{Start: 0, End: 65535},
				},
				DestinationPorts: []firewall.PortRange{
					{Start: 80, End: 80},
					{Start: 443, End: 443},
					{Start: 8080, End: 8083},
				},
				LogLevel: firewall.LogLevelEmergency,
				Digest:   "0102030405060708090a0b0c0d0e0f1011121314",
			},
			{
				Enable:             false,
				Description:        "test_rule_2",
				SecurityGroup:      "",
				Interface:          "eth1",
				Direction:          firewall.DirectionOut,
				Action:             firewall.ActionDrop,
				SourceAddress:      "10.0.0.0-10.255.255.255",
				DestinationAddress: "0.0.0.0",
				Macro:              firewall.MacroPing,
				Protocol:           firewall.ProtocolNone,
				SourcePorts:        nil,
				DestinationPorts:   nil,
				LogLevel:           firewall.LogLevelCritical,
				Digest:             "15161718191a1b1c1d1e1f202122232425262728",
			},
			{
				Enable:        true,
				Description:   "test_rule_3",
				SecurityGroup: "test_security_group",
				Interface:     "eth2",
				Digest:        "292a2b2c2d2e2f303132333435363738393a3b3c",
			},
		}

		rules, err := n.ListFirewallRules()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedRules, rules)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_firewall_rules_{rule}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/firewall/rules/0", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedRule := firewall.Rule{
			Enable:             true,
			Description:        "test_rule_1",
			SecurityGroup:      "",
			Interface:          "eth0",
			Direction:          firewall.DirectionIn,
			Action:             firewall.ActionAccept,
			SourceAddress:      "0.0.0.0/0",
			DestinationAddress: "10.0.0.0-10.0.0.255",
			Macro:              firewall.MacroNone,
			Protocol:           firewall.ProtocolTCP,
			SourcePorts: []firewall.PortRange{
				{Start: 0, End: 65535},
			},
			DestinationPorts: []firewall.PortRange{
				{Start: 80, End: 80},
				{Start: 443, End: 443},
				{Start: 8080, End: 8083},
			},
			LogLevel: firewall.LogLevelEmergency,
			Digest:   "0102030405060708090a0b0c0d0e0f1011121314",
		}

		rule, err := n.GetFirewallRule(0)
		require.NoError(t, err)
		assert.Equal(t, expectedRule, rule)

		exc.AssertExpectations(t)
	})

	t.Run("Add", func(t *testing.T) {
		exc.
			On("Request", http.MethodPost, "nodes/test_node/firewall/rules", url.Values{
				"enable":  {"1"},
				"comment": {"test_rule_1"},
				"log":     {"emerg"},
				"type":    {"in"},
				"action":  {"ACCEPT"},
				"iface":   {"eth0"},
				"source":  {"0.0.0.0/0"},
				"dest":    {"10.0.0.0-10.0.0.255"},
				"proto":   {"tcp"},
				"sport":   {"0:65535"},
				"dport":   {"80,443,8080:8083"},
				"digest":  {"0102030405060708090a0b0c0d0e0f1011121314"},
			}).
			Return(nil, nil).
			Once()

		err := n.AddFirewallRule(firewall.Rule{
			Enable:             true,
			Description:        "test_rule_1",
			SecurityGroup:      "",
			Interface:          "eth0",
			Direction:          firewall.DirectionIn,
			Action:             firewall.ActionAccept,
			SourceAddress:      "0.0.0.0/0",
			DestinationAddress: "10.0.0.0-10.0.0.255",
			Macro:              firewall.MacroNone,
			Protocol:           firewall.ProtocolTCP,
			SourcePorts: []firewall.PortRange{
				{Start: 0, End: 65535},
			},
			DestinationPorts: []firewall.PortRange{
				{Start: 80, End: 80},
				{Start: 443, End: 443},
				{Start: 8080, End: 8083},
			},
			LogLevel: firewall.LogLevelEmergency,
			Digest:   "0102030405060708090a0b0c0d0e0f1011121314",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Edit", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/firewall/rules/0", url.Values{
				"enable":  {"1"},
				"comment": {"test_rule_1"},
				"log":     {"emerg"},
				"type":    {"in"},
				"action":  {"ACCEPT"},
				"iface":   {"eth0"},
				"source":  {"0.0.0.0/0"},
				"dest":    {"10.0.0.0-10.0.0.255"},
				"proto":   {"tcp"},
				"sport":   {"0:65535"},
				"dport":   {"80,443,8080:8083"},
				"digest":  {"0102030405060708090a0b0c0d0e0f1011121314"},
				"delete":  {"macro"},
			}).
			Return(nil, nil).
			Once()

		err := n.EditFirewallRule(0, firewall.Rule{
			Enable:             true,
			Description:        "test_rule_1",
			SecurityGroup:      "",
			Interface:          "eth0",
			Direction:          firewall.DirectionIn,
			Action:             firewall.ActionAccept,
			SourceAddress:      "0.0.0.0/0",
			DestinationAddress: "10.0.0.0-10.0.0.255",
			Macro:              firewall.MacroNone,
			Protocol:           firewall.ProtocolTCP,
			SourcePorts: []firewall.PortRange{
				{Start: 0, End: 65535},
			},
			DestinationPorts: []firewall.PortRange{
				{Start: 80, End: 80},
				{Start: 443, End: 443},
				{Start: 8080, End: 8083},
			},
			LogLevel: firewall.LogLevelEmergency,
			Digest:   "0102030405060708090a0b0c0d0e0f1011121314",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Move", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/firewall/rules/0", url.Values{
				"moveto": {"1"},
			}).
			Return(nil, nil).
			Once()

		err := n.MoveFirewallRule(0, 1)
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		exc.
			On("Request", http.MethodDelete, "nodes/test_node/firewall/rules/0", url.Values(nil)).
			Return(nil, nil).
			Once()

		err := n.DeleteFirewallRule(0, "")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("DeleteDigest", func(t *testing.T) {
		exc.
			On("Request", http.MethodDelete, "nodes/test_node/firewall/rules/0", url.Values{
				"digest": {"0102030405060708090a0b0c0d0e0f1011121314"},
			}).
			Return(nil, nil).
			Once()

		err := n.DeleteFirewallRule(0, "0102030405060708090a0b0c0d0e0f1011121314")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}
