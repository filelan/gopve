package cluster_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/cluster"
	"github.com/xabinapal/gopve/internal/service/cluster/test"
	"github.com/xabinapal/gopve/pkg/types/firewall"
)

func TestFirewallAliases(t *testing.T) {
	svc, _, exc := test.NewService()

	getAlias := func() *cluster.FirewallAlias {
		return cluster.NewFirewallAlias(svc, "test_alias", "test_description", "127.0.0.1", "test_digest")
	}

	t.Run("Name", func(t *testing.T) {
		alias := getAlias()
		expectedName := "test_alias"
		name := alias.Name()
		assert.Equal(t, expectedName, name)
	})

	t.Run("Description", func(t *testing.T) {
		alias := getAlias()
		expectedDescription := "test_description"
		description := alias.Description()
		assert.Equal(t, expectedDescription, description)
	})

	t.Run("Address", func(t *testing.T) {
		alias := getAlias()
		expectedAddress := "127.0.0.1"
		address := alias.Address()
		assert.Equal(t, expectedAddress, address)
	})

	t.Run("Digest", func(t *testing.T) {
		alias := getAlias()
		expectedDigest := "test_digest"
		digest := alias.Digest()
		assert.Equal(t, expectedDigest, digest)
	})

	t.Run("Rename", func(t *testing.T) {
		alias := getAlias()

		exc.
			On("Request", http.MethodPut, "cluster/firewall/aliases/test_alias", url.Values{
				"rename": {"new_name"},
			}).
			Return(nil, nil).
			Once()

		err := alias.Rename("new_name")
		require.NoError(t, err)

		assert.Equal(t, "new_name", alias.Name())

		exc.AssertExpectations(t)
	})

	t.Run("GetProperties", func(t *testing.T) {
		alias := getAlias()
		expectedProperties := firewall.AliasProperties{
			Description: "test_description",
			Address:     "127.0.0.1",
			Digest:      "test_digest",
		}

		properties, err := alias.GetProperties()
		require.NoError(t, err)
		assert.Equal(t, expectedProperties, properties)
	})

	t.Run("SetProperties", func(t *testing.T) {
		alias := getAlias()
		exc.
			On("Request", http.MethodPut, "cluster/firewall/aliases/test_alias", url.Values{
				"comment": {"new_description"},
				"cidr":    {"10.0.0.1"},
				"digest":  {"new_digest"},
			}).
			Return(nil, nil).
			Once()

		err := alias.SetProperties(firewall.AliasProperties{
			Description: "new_description",
			Address:     "10.0.0.1",
			Digest:      "new_digest",
		})
		require.NoError(t, err)

		assert.Equal(t, "new_description", alias.Description())
		assert.Equal(t, "10.0.0.1", alias.Address())
		assert.Equal(t, "new_digest", alias.Digest())

		exc.AssertExpectations(t)
	})
}

func TestFirewallIPSets(t *testing.T) {
	svc, _, exc := test.NewService()

	getIPSet := func() *cluster.FirewallIPSet {
		return cluster.NewFirewallIPSet(svc, "test_ipset", "test_description", "test_digest")
	}

	t.Run("Name", func(t *testing.T) {
		ipSet := getIPSet()
		expectedName := "test_ipset"
		name := ipSet.Name()
		assert.Equal(t, expectedName, name)
	})

	t.Run("Description", func(t *testing.T) {
		ipSet := getIPSet()
		expectedDescription := "test_description"
		description := ipSet.Description()
		assert.Equal(t, expectedDescription, description)
	})

	t.Run("Digest", func(t *testing.T) {
		ipSet := getIPSet()
		expectedDigest := "test_digest"
		digest := ipSet.Digest()
		assert.Equal(t, expectedDigest, digest)
	})

	t.Run("Rename", func(t *testing.T) {
		ipSet := getIPSet()

		exc.
			On("Request", http.MethodPost, "cluster/firewall/ipset", url.Values{
				"group":  {"test_ipset"},
				"rename": {"new_name"},
			}).
			Return(nil, nil).
			Once()

		err := ipSet.Rename("new_name")
		require.NoError(t, err)

		assert.Equal(t, "new_name", ipSet.Name())

		exc.AssertExpectations(t)
	})

	t.Run("GetProperties", func(t *testing.T) {
		ipSet := getIPSet()
		expectedProperties := firewall.IPSetProperties{
			Description: "test_description",
			Digest:      "test_digest",
		}

		properties, err := ipSet.GetProperties()
		require.NoError(t, err)
		assert.Equal(t, expectedProperties, properties)
	})

	t.Run("SetProperties", func(t *testing.T) {
		ipSet := getIPSet()

		exc.
			On("Request", http.MethodPost, "cluster/firewall/ipset", url.Values{
				"group":   {"test_ipset"},
				"comment": {"new_description"},
				"digest":  {"new_digest"},
			}).
			Return(nil, nil).
			Once()

		err := ipSet.SetProperties(firewall.IPSetProperties{
			Description: "new_description",
			Digest:      "new_digest",
		})
		require.NoError(t, err)

		assert.Equal(t, "new_description", ipSet.Description())
		assert.Equal(t, "new_digest", ipSet.Digest())

		exc.AssertExpectations(t)
	})

	t.Run("List", func(t *testing.T) {
		ipSet := getIPSet()

		response, err := ioutil.ReadFile("./testdata/get_cluster_firewall_ipset_{name}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/firewall/ipset/test_ipset", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedAddresses := []firewall.IPSetAddress{
			{
				Address:     "10.0.0.1",
				Description: "PVE test_node",
				NoMatch:     false,
				Digest:      "0102030405060708090a0b0c0d0e0f1011121314",
			},
			{
				Address:     "10.0.0.2",
				Description: "PVE test_node2",
				NoMatch:     false,
				Digest:      "15161718191a1b1c1d1e1f202122232425262728",
			},
			{
				Address:     "10.0.0.3",
				Description: "PVE test_node3",
				NoMatch:     false,
				Digest:      "292a2b2c2d2e2f303132333435363738393a3b3c",
			},
			{
				Address:     "10.0.0.254",
				Description: "Gateway",
				NoMatch:     true,
				Digest:      "3d3e3f404142434445464748494a4b4c4d4e4f50",
			},
		}

		addresses, err := ipSet.ListAddresses()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedAddresses, addresses)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		ipSet := getIPSet()

		response, err := ioutil.ReadFile("./testdata/get_cluster_firewall_ipset_{name}_{cidr}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/firewall/ipset/test_ipset/10.0.0.1", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedAddress := firewall.IPSetAddress{
			Address:     "10.0.0.1",
			Description: "PVE test_node",
			NoMatch:     false,
			Digest:      "0102030405060708090a0b0c0d0e0f1011121314",
		}

		address, err := ipSet.GetAddress("10.0.0.1")
		require.NoError(t, err)
		assert.Equal(t, expectedAddress, address)

		exc.AssertExpectations(t)
	})

	t.Run("Add", func(t *testing.T) {
		ipSet := getIPSet()

		exc.
			On("Request", http.MethodPost, "cluster/firewall/ipset/test_ipset", url.Values{
				"cidr":    {"127.0.0.1"},
				"comment": {"test_description"},
				"nomatch": {"0"},
			}).
			Return(nil, nil).
			Once()

		err := ipSet.AddAddress(firewall.IPSetAddress{
			Address:     "127.0.0.1",
			Description: "test_description",
			NoMatch:     false,
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Edit", func(t *testing.T) {
		ipSet := getIPSet()

		exc.
			On("Request", http.MethodPut, "cluster/firewall/ipset/test_ipset/127.0.0.1", url.Values{
				"comment": {"new_description"},
				"nomatch": {"1"},
				"digest":  {"0102030405060708090a0b0c0d0e0f1011121314"},
			}).
			Return(nil, nil).
			Once()

		err := ipSet.EditAddress(firewall.IPSetAddress{
			Address:     "127.0.0.1",
			Description: "new_description",
			NoMatch:     true,
			Digest:      "0102030405060708090a0b0c0d0e0f1011121314",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		ipSet := getIPSet()

		exc.
			On("Request", http.MethodDelete, "cluster/firewall/ipset/test_ipset/127.0.0.1", url.Values(nil)).
			Return(nil, nil).
			Once()

		err := ipSet.DeleteAddress("127.0.0.1", "")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("DeleteDigest", func(t *testing.T) {
		ipSet := getIPSet()

		exc.
			On("Request", http.MethodDelete, "cluster/firewall/ipset/test_ipset/127.0.0.1", url.Values{
				"digest": {"0102030405060708090a0b0c0d0e0f1011121314"},
			}).
			Return(nil, nil).
			Once()

		err := ipSet.DeleteAddress("127.0.0.1", "0102030405060708090a0b0c0d0e0f1011121314")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}

func TestFirewallServiceGroupRules(t *testing.T) {
	svc, _, exc := test.NewService()

	getServiceGroup := func() *cluster.FirewallServiceGroup {
		return cluster.NewFirewallServiceGroup(svc, "test_sg", "test_description", "test_digest")
	}

	t.Run("Name", func(t *testing.T) {
		serviceGroup := getServiceGroup()
		expectedName := "test_sg"
		name := serviceGroup.Name()
		assert.Equal(t, expectedName, name)
	})

	t.Run("Description", func(t *testing.T) {
		serviceGroup := getServiceGroup()
		expectedDescription := "test_description"
		description := serviceGroup.Description()
		assert.Equal(t, expectedDescription, description)
	})

	t.Run("Digest", func(t *testing.T) {
		serviceGroup := getServiceGroup()
		expectedDigest := "test_digest"
		digest := serviceGroup.Digest()
		assert.Equal(t, expectedDigest, digest)
	})

	t.Run("Rename", func(t *testing.T) {
		serviceGroup := getServiceGroup()

		exc.
			On("Request", http.MethodPost, "cluster/firewall/groups", url.Values{
				"group":  {"test_sg"},
				"rename": {"new_name"},
			}).
			Return(nil, nil).
			Once()

		err := serviceGroup.Rename("new_name")
		require.NoError(t, err)

		assert.Equal(t, "new_name", serviceGroup.Name())

		exc.AssertExpectations(t)
	})

	t.Run("GetProperties", func(t *testing.T) {
		serviceGroup := getServiceGroup()
		expectedProperties := firewall.ServiceGroupProperties{
			Description: "test_description",
			Digest:      "test_digest",
		}

		properties, err := serviceGroup.GetProperties()
		require.NoError(t, err)
		assert.Equal(t, expectedProperties, properties)
	})

	t.Run("SetProperties", func(t *testing.T) {
		serviceGroup := getServiceGroup()
		exc.
			On("Request", http.MethodPost, "cluster/firewall/groups", url.Values{
				"group":   {"test_sg"},
				"comment": {"new_description"},
				"digest":  {"new_digest"},
			}).
			Return(nil, nil).
			Once()

		err := serviceGroup.SetProperties(firewall.ServiceGroupProperties{
			Description: "new_description",
			Digest:      "new_digest",
		})
		require.NoError(t, err)

		assert.Equal(t, "new_description", serviceGroup.Description())
		assert.Equal(t, "new_digest", serviceGroup.Digest())

		exc.AssertExpectations(t)
	})

	t.Run("List", func(t *testing.T) {
		serviceGroup := getServiceGroup()

		response, err := ioutil.ReadFile("./testdata/get_cluster_firewall_groups_{group}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/firewall/groups/test_sg", url.Values(nil)).
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

		rules, err := serviceGroup.ListFirewallRules()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedRules, rules)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		serviceGroup := getServiceGroup()

		response, err := ioutil.ReadFile("./testdata/get_cluster_firewall_groups_{group}_{rule}.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/firewall/groups/test_sg/0", url.Values(nil)).
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

		rule, err := serviceGroup.GetFirewallRule(0)
		require.NoError(t, err)
		assert.Equal(t, expectedRule, rule)

		exc.AssertExpectations(t)
	})

	t.Run("Add", func(t *testing.T) {
		serviceGroup := getServiceGroup()

		exc.
			On("Request", http.MethodPost, "cluster/firewall/groups/test_sg", url.Values{
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

		err := serviceGroup.AddFirewallRule(firewall.Rule{
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
		serviceGroup := getServiceGroup()

		exc.
			On("Request", http.MethodPut, "cluster/firewall/groups/test_sg/0", url.Values{
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

		err := serviceGroup.EditFirewallRule(0, firewall.Rule{
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
		serviceGroup := getServiceGroup()

		exc.
			On("Request", http.MethodPut, "cluster/firewall/groups/test_sg/0", url.Values{
				"moveto": {"1"},
			}).
			Return(nil, nil).
			Once()

		err := serviceGroup.MoveFirewallRule(0, 1)
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		serviceGroup := getServiceGroup()

		exc.
			On("Request", http.MethodDelete, "cluster/firewall/groups/test_sg/0", url.Values(nil)).
			Return(nil, nil).
			Once()

		err := serviceGroup.DeleteFirewallRule(0, "")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("DeleteDigest", func(t *testing.T) {
		serviceGroup := getServiceGroup()

		exc.
			On("Request", http.MethodDelete, "cluster/firewall/groups/test_sg/0", url.Values{
				"digest": {"0102030405060708090a0b0c0d0e0f1011121314"},
			}).
			Return(nil, nil).
			Once()

		err := serviceGroup.DeleteFirewallRule(0, "0102030405060708090a0b0c0d0e0f1011121314")
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}
