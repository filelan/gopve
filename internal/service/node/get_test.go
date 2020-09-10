package node_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/node"
	"github.com/xabinapal/gopve/internal/service/node/test"
	types "github.com/xabinapal/gopve/pkg/types/node"
)

func TestServiceList(t *testing.T) {
	svc, exc := test.NewService()

	response, err := ioutil.ReadFile("./testdata/service_list_nodes.json")
	require.NoError(t, err)

	expectedNodes := []types.Node{
		node.NewNode(svc, "test_node", types.StatusOnline),
		node.NewNode(svc, "test_node2", types.StatusOffline),
		node.NewNode(svc, "test_node3", types.StatusUnknown),
	}

	exc.
		On("Request", http.MethodGet, "cluster/resources", url.Values{
			"type": {"node"},
		}).
		Return(response, nil).
		Once()

	nodes, err := svc.List()
	require.NoError(t, err)
	assert.ElementsMatch(t, expectedNodes, nodes)

	exc.AssertExpectations(t)
}
