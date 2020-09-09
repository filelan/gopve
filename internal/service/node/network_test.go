package node_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/node/test"
	"github.com/xabinapal/gopve/pkg/types/node"
)

func TestNodeNetworkDNS(t *testing.T) {
	n, exc := test.NewNode()

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/node_get_dns.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/dns", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedSettings := node.DNSSettings{
			FirstDNS:     "1.1.1.1",
			SecondDNS:    "1.0.0.1",
			ThirdDNS:     "208.67.222.222",
			SearchDomain: "pve.local",
		}

		settings, err := n.GetDNSSettings()
		require.NoError(t, err)
		assert.Equal(t, expectedSettings, settings)

		exc.AssertExpectations(t)
	})

	t.Run("Set", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/dns", url.Values{
				"dns1":   {"1.1.1.1"},
				"dns2":   {"1.0.0.1"},
				"dns3":   {"208.67.222.222"},
				"search": {"pve.local"},
			}).
			Return(nil, nil).
			Once()

		err := n.SetDNSSettings(node.DNSSettings{
			FirstDNS:     "1.1.1.1",
			SecondDNS:    "1.0.0.1",
			ThirdDNS:     "208.67.222.222",
			SearchDomain: "pve.local",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}

func TestNodeNetworkHosts(t *testing.T) {
	n, exc := test.NewNode()

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/node_get_hosts.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/hosts", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedFile := node.HostsFile{
			Contents: "127.0.0.1 localhost.localdomain localhost\\n10.0.0.1 test_node.pve.local test_node\\n\\n# The following lines are desirable for IPv6 capable hosts\\n\\n::1     ip6-localhost ip6-loopback\\nfe00::0 ip6-localnet\\nff00::0 ip6-mcastprefix\\nff02::1 ip6-allnodes\\nff02::2 ip6-allrouters\\nff02::3 ip6-allhosts\\n",
			Digest:   "60985c46740a60b8744b58b70533dff50f50a1a3",
		}

		file, err := n.GetHostsFile()
		require.NoError(t, err)
		assert.Equal(t, expectedFile, file)

		exc.AssertExpectations(t)
	})

	t.Run("Set", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/hosts", url.Values{
				"data":   {"127.0.0.1 localhost.localdomain localhost\\n10.0.0.1 test_node.pve.local test_node\\n\\n# The following lines are desirable for IPv6 capable hosts\\n\\n::1     ip6-localhost ip6-loopback\\nfe00::0 ip6-localnet\\nff00::0 ip6-mcastprefix\\nff02::1 ip6-allnodes\\nff02::2 ip6-allrouters\\nff02::3 ip6-allhosts\\n"},
				"digest": {"60985c46740a60b8744b58b70533dff50f50a1a3"},
			}).
			Return(nil, nil).
			Once()

		err := n.SetHostsFile(node.HostsFile{
			Contents: "127.0.0.1 localhost.localdomain localhost\\n10.0.0.1 test_node.pve.local test_node\\n\\n# The following lines are desirable for IPv6 capable hosts\\n\\n::1     ip6-localhost ip6-loopback\\nfe00::0 ip6-localnet\\nff00::0 ip6-mcastprefix\\nff02::1 ip6-allnodes\\nff02::2 ip6-allrouters\\nff02::3 ip6-allhosts\\n",
			Digest:   "60985c46740a60b8744b58b70533dff50f50a1a3",
		})
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}
