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
	types "github.com/xabinapal/gopve/pkg/types/cluster"
)

func TestClusterServiceGet(t *testing.T) {
	svc, _, exc := test.NewService()

	t.Run("Standalone", func(t *testing.T) {
		exc.
			On("Request", http.MethodGet, "cluster/config/join", url.Values(nil)).
			Return(nil, types.ErrNotInCluster).
			Once()

		expectedCluster := cluster.NewCluster(svc, types.ModeStandalone, "")

		cluster, err := svc.Get()
		require.NoError(t, err)
		assert.Equal(t, expectedCluster, cluster)

		exc.AssertExpectations(t)
	})

	t.Run("Cluster", func(t *testing.T) {
		response, err := ioutil.ReadFile("./testdata/cluster_config_join.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "cluster/config/join", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedCluster := cluster.NewCluster(svc, types.ModeCluster, "test_cluster")

		cluster, err := svc.Get()
		require.NoError(t, err)
		assert.Equal(t, expectedCluster, cluster)

		exc.AssertExpectations(t)
	})
}
