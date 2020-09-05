package node_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/node/test"
)

func TestServiceList(t *testing.T) {
	svc, exc := test.NewService()

	response, err := ioutil.ReadFile("./testdata/service_list_nodes.json")
	require.NoError(t, err)

	exc.
		On("Request", http.MethodGet, "cluster/resources", url.Values{
			"type": {"node"},
		}).
		Return(response, nil).
		Once()

	nodes, err := svc.List()
	require.NoError(t, err)
	assert.Equal(t, nil, nodes)

	exc.AssertExpectations(t)
}
