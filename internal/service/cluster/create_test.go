package cluster_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/cluster/test"
	task "github.com/xabinapal/gopve/internal/service/task/test"
	types "github.com/xabinapal/gopve/pkg/types/cluster"
)

func TestClusterServiceCreate(t *testing.T) {
	svc, api, exc := test.NewService()

	t.Run("Standalone", func(t *testing.T) {
		exc.
			On("Request", http.MethodGet, "cluster/config/join", url.Values(nil)).
			Return(nil, types.ErrNotInCluster).
			Once()

		exc.
			On("Request", http.MethodPost, "cluster/config", url.Values{
				"clustername": {"test_cluster"},
				"nodeid":      {"1"},
				"votes":       {"1"},
				"link0":       {"address=10.0.0.1,priority=1"},
			}).
			Return(
				[]byte(
					"{\"data\":\"UPID:test_node::::clustercreate:test_cluster:root@pam:\"}",
				),
				nil,
			).
			Once()

		expectedTask, _, _ := task.NewTask(
			"test_node",
			"::",
			"clustercreate",
			"test_cluster",
			"root@pam",
			"",
		)

		api.TaskService.
			On("Get", "UPID:test_node::::clustercreate:test_cluster:root@pam:").
			Return(expectedTask, nil)

		task, err := svc.Create("test_cluster", types.NodeProperties{
			ID:    1,
			Votes: 1,

			Link0: types.NodeLink{
				Address:  "10.0.0.1",
				Priority: 1,
			},
		})

		require.NoError(t, err)
		assert.Equal(t, expectedTask, task)

		exc.AssertExpectations(t)
	})

	t.Run("Cluster", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/cluster_config_join.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/config/join", url.Values(nil)).
			Return(response, nil).
			Once()

		_, err = svc.Create("test_cluster", types.NodeProperties{
			ID:    1,
			Votes: 1,

			Link0: types.NodeLink{
				Address:  "10.0.0.1",
				Priority: 1,
			},
		})
		require.EqualError(t, err, types.ErrAlreadyInCluster.Error())

		exc.AssertExpectations(t)
	})
}

func TestClusterServiceJoin(t *testing.T) {
	svc, api, exc := test.NewService()

	t.Run("Standalone", func(t *testing.T) {
		exc.
			On("Request", http.MethodGet, "cluster/config/join", url.Values(nil)).
			Return(nil, types.ErrNotInCluster).
			Once()

		exc.
			On("Request", http.MethodPost, "cluster/config/join", url.Values{
				"hostname":    {"10.0.0.1"},
				"password":    {"test_password"},
				"fingerprint": {"test_fingerprint"},
				"nodeid":      {"2"},
				"votes":       {"1"},
				"link0":       {"address=10.0.0.2,priority=1"},
			}).
			Return(
				[]byte(
					"{\"data\":\"UPID:test_node2::::clustercreate:test_cluster:root@pam:\"}",
				),
				nil,
			).
			Once()

		expectedTask, _, _ := task.NewTask(
			"test_node2",
			"::",
			"clustercreate",
			"test_cluster",
			"root@pam",
			"",
		)

		api.TaskService.
			On("Get", "UPID:test_node2::::clustercreate:test_cluster:root@pam:").
			Return(expectedTask, nil)

		task, err := svc.Join(
			"10.0.0.1",
			"test_password",
			"test_fingerprint",
			types.NodeProperties{
				ID:    2,
				Votes: 1,

				Link0: types.NodeLink{
					Address:  "10.0.0.2",
					Priority: 1,
				},
			},
		)

		require.NoError(t, err)
		assert.Equal(t, expectedTask, task)

		exc.AssertExpectations(t)
	})

	t.Run("Cluster", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/cluster_config_join.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/config/join", url.Values(nil)).
			Return(response, nil).
			Once()

		_, err = svc.Join(
			"10.0.0.1",
			"test_password",
			"test_fingerprint",
			types.NodeProperties{
				ID:    2,
				Votes: 1,

				Link0: types.NodeLink{
					Address:  "10.0.0.2",
					Priority: 1,
				},
			},
		)
		require.EqualError(t, err, types.ErrAlreadyInCluster.Error())

		exc.AssertExpectations(t)
	})
}
