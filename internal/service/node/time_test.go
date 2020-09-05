package node_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/node/test"
)

func TestNodeTime(t *testing.T) {
	node, exc := test.NewNode()

	response, err := ioutil.ReadFile("./testdata/node_get_time.json")
	require.NoError(t, err)

	t.Run("GetUTC", func(t *testing.T) {
		exc.
			On("Request", http.MethodGet, "nodes/test_node/time", url.Values(nil)).
			Return(response, nil).
			Once()

		location, err := time.LoadLocation("UTC")
		require.NoError(t, err)

		expectedTime := time.Unix(1609458356, 0).In(location)

		time, err := node.GetTime(false)
		require.NoError(t, err)
		assert.Equal(t, expectedTime, time)

		exc.AssertExpectations(t)
	})

	t.Run("GetLocal", func(t *testing.T) {
		exc.
			On("Request", http.MethodGet, "nodes/test_node/time", url.Values(nil)).
			Return(response, nil).
			Once()

		location, err := time.LoadLocation("Europe/Madrid")
		require.NoError(t, err)

		expectedTime := time.Unix(1609458356, 0).In(location)

		time, err := node.GetTime(true)
		require.NoError(t, err)
		assert.Equal(t, expectedTime, time)

		exc.AssertExpectations(t)
	})
}

func TestNodeTimezone(t *testing.T) {
	node, exc := test.NewNode()

	response, err := ioutil.ReadFile("./testdata/node_get_time.json")
	require.NoError(t, err)

	t.Run("Get", func(t *testing.T) {
		exc.
			On("Request", http.MethodGet, "nodes/test_node/time", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedLocation, err := time.LoadLocation("Europe/Madrid")
		require.NoError(t, err)

		location, err := node.GetTimezone()
		require.NoError(t, err)
		assert.Equal(t, expectedLocation, location)

		exc.AssertExpectations(t)
	})

	t.Run("Set", func(t *testing.T) {
		exc.
			On("Request", http.MethodPut, "nodes/test_node/time", url.Values{
				"timezone": {"Europe/Madrid"},
			}).
			Return(response, nil).
			Once()

		location, err := time.LoadLocation("Europe/Madrid")
		require.NoError(t, err)

		err = node.SetTimezone(location)
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})
}
