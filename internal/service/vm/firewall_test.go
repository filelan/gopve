package vm_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/internal/service/vm/test"
	"github.com/xabinapal/gopve/pkg/types/firewall"
)

func TestVirtualMachineFirewallProperties(t *testing.T) {
	virtualMachine, exc := test.NewVirtualMachine()

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_options.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/options", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedProperties := firewall.VMProperties{
			Enable:           true,
			LogLevelIncoming: firewall.LogLevelInfo,
			LogLevelOutgoing: firewall.LogLevelWarning,

			DefaultInputPolicy:  firewall.ActionDrop,
			DefaultOutputPolicy: firewall.ActionAccept,

			EnableNDP:       true,
			EnableRADV:      true,
			EnableDHCP:      true,
			EnableMACFilter: true,
			EnableIPFilter:  true,

			Digest: "0000000000000000000000000000000000000000",
		}

		properties, err := virtualMachine.GetFirewallProperties()
		require.NoError(t, err)
		assert.Equal(t, expectedProperties, properties)

		exc.AssertExpectations(t)
	})

	t.Run("Set", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/test_kind/100/firewall/options", url.Values{
				"enable":        {"1"},
				"log_level_in":  {"info"},
				"log_level_out": {"warning"},

				"policy_in":  {"DROP"},
				"policy_out": {"ACCEPT"},

				"ndp":       {"1"},
				"radv":      {"1"},
				"dhcp":      {"1"},
				"macfilter": {"1"},
				"ipfilter":  {"1"},

				"digest": {"0000000000000000000000000000000000000000"},
			}).
			Return(nil, nil).
			Once()

		err := virtualMachine.SetFirewallProperties(firewall.VMProperties{
			Enable:           true,
			LogLevelIncoming: firewall.LogLevelInfo,
			LogLevelOutgoing: firewall.LogLevelWarning,

			DefaultInputPolicy:  firewall.ActionDrop,
			DefaultOutputPolicy: firewall.ActionAccept,

			EnableNDP:       true,
			EnableRADV:      true,
			EnableDHCP:      true,
			EnableMACFilter: true,
			EnableIPFilter:  true,

			Digest: "0000000000000000000000000000000000000000",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}

func TestVirtualMachineServiceFirewallAliases(t *testing.T) {
	virtualMachine, exc := test.NewVirtualMachine()

	t.Run("List", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_aliases.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/aliases", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedAliases := []firewall.Alias{
			vm.NewFirewallAlias(virtualMachine, "local", "LAN network", "10.0.0.0/8", "0102030405060708090a0b0c0d0e0f1011121314"),
			vm.NewFirewallAlias(virtualMachine, "self", "PVE test_node", "10.0.0.1", "15161718191a1b1c1d1e1f202122232425262728"),
		}

		aliases, err := virtualMachine.ListFirewallAliases()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedAliases, aliases)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_aliases_{name}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/aliases/local", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedAlias := vm.NewFirewallAlias(virtualMachine, "local", "LAN network", "10.0.0.0/8", "0102030405060708090a0b0c0d0e0f1011121314")

		alias, err := virtualMachine.GetFirewallAlias("local")
		require.NoError(t, err)
		assert.Equal(t, expectedAlias, alias)

		exc.AssertExpectations(t)
	})
}

func TestVirtualMachineServiceFirewallIPSets(t *testing.T) {
	virtualMachine, exc := test.NewVirtualMachine()

	t.Run("List", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_ipset.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/ipset", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedIPSets := []firewall.IPSet{
			vm.NewFirewallIPSet(virtualMachine, "internal", "LAN hosts", "0102030405060708090a0b0c0d0e0f1011121314"),
			vm.NewFirewallIPSet(virtualMachine, "dns", "DNS servers", "15161718191a1b1c1d1e1f202122232425262728"),
		}

		ipSets, err := virtualMachine.ListFirewallIPSets()

		require.NoError(t, err)
		assert.ElementsMatch(t, expectedIPSets, ipSets)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_ipset.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/ipset", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedIPSet := vm.NewFirewallIPSet(virtualMachine, "internal", "LAN hosts", "0102030405060708090a0b0c0d0e0f1011121314")

		ipSet, err := virtualMachine.GetFirewallIPSet("internal")
		require.NoError(t, err)
		assert.Equal(t, expectedIPSet, ipSet)

		exc.AssertExpectations(t)
	})
}

func TestVirtualMachineServiceFirewallServiceGroups(t *testing.T) {
	virtualMachine, exc := test.NewVirtualMachine()

	t.Run("List", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_groups.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/groups", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedServiceGroups := []firewall.ServiceGroup{
			vm.NewFirewallServiceGroup(virtualMachine, "internal", "LAN hosts", "0102030405060708090a0b0c0d0e0f1011121314"),
			vm.NewFirewallServiceGroup(virtualMachine, "dns", "DNS servers", "15161718191a1b1c1d1e1f202122232425262728"),
		}

		serviceGroups, err := virtualMachine.ListFirewallServiceGroups()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedServiceGroups, serviceGroups)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_groups.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/groups", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedServiceGroup := vm.NewFirewallServiceGroup(virtualMachine, "internal", "LAN hosts", "0102030405060708090a0b0c0d0e0f1011121314")

		serviceGroup, err := virtualMachine.GetFirewallServiceGroup("internal")
		require.NoError(t, err)
		assert.Equal(t, expectedServiceGroup, serviceGroup)

		exc.AssertExpectations(t)
	})
}

func TestVirtualMachineFirewallRules(t *testing.T) {
	vm, exc := test.NewVirtualMachine()

	t.Run("List", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_rules.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/rules", url.Values(nil)).
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

		rules, err := vm.ListFirewallRules()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedRules, rules)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_firewall_rules_{rule}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/firewall/rules/0", url.Values(nil)).
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

		rule, err := vm.GetFirewallRule(0)
		require.NoError(t, err)
		assert.Equal(t, expectedRule, rule)

		exc.AssertExpectations(t)
	})

	t.Run("Add", func(t *testing.T) {
		exc.
			On("Request", http.MethodPost, "nodes/test_node/test_kind/100/firewall/rules", url.Values{
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
			}).
			Return(nil, nil).
			Once()

		err := vm.AddFirewallRule(firewall.Rule{
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
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Edit", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/test_kind/100/firewall/rules/0", url.Values{
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

		err := vm.EditFirewallRule(0, firewall.Rule{
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
			On("Request", http.MethodPut, "nodes/test_node/test_kind/100/firewall/rules/0", url.Values{
				"moveto": {"1"},
			}).
			Return(nil, nil).
			Once()

		err := vm.MoveFirewallRule(0, 1)
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		exc.
			On("Request", http.MethodDelete, "nodes/test_node/test_kind/100/firewall/rules/0", url.Values(nil)).
			Return(nil, nil).
			Once()

		err := vm.DeleteFirewallRule(0, "")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("DeleteDigest", func(t *testing.T) {
		exc.
			On("Request", http.MethodDelete, "nodes/test_node/test_kind/100/firewall/rules/0", url.Values{
				"digest": {"0102030405060708090a0b0c0d0e0f1011121314"},
			}).
			Return(nil, nil).
			Once()

		err := vm.DeleteFirewallRule(0, "0102030405060708090a0b0c0d0e0f1011121314")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}
