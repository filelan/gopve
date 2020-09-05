package node_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/node/test"
)

func TestNodePower(t *testing.T) {
	node, exc := test.NewNode()

	t.Run("Shutdown", func(t *testing.T) {
		exc.
			On("Request", http.MethodPost, "nodes/test_node/status", url.Values{
				"command": {"shutdown"},
			}).
			Return([]byte{}, nil).
			Once()

		err := node.Shutdown()
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("Reboot", func(t *testing.T) {
		exc.
			On("Request", http.MethodPost, "nodes/test_node/status", url.Values{
				"command": {"reboot"},
			}).
			Return([]byte{}, nil).
			Once()

		err := node.Reboot()
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}
