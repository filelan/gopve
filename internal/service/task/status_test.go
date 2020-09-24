package task_test

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/task/test"
	types "github.com/xabinapal/gopve/pkg/types/task"
)

func TestTaskGetStatus(t *testing.T) {
	obj, _, exc := test.NewTask(
		"test_node",
		"00000000:00000000:00000000",
		"test_action",
		"test_id",
		"test_user",
		"test_extra",
	)

	options := map[string]struct {
		GoldenFile string
		TaskStatus types.Status
	}{
		"Running": {
			GoldenFile: "./testdata/get_nodes_{node}_tasks_{upid}_status__running.json",
			TaskStatus: types.StatusRunning,
		},
		"Stopped": {
			GoldenFile: "./testdata/get_nodes_{node}_tasks_{upid}_status__stopped.json",
			TaskStatus: types.StatusStopped,
		},
	}

	for n, tt := range options {
		tt := tt

		t.Run(n, func(t *testing.T) {
			response, err := ioutil.ReadFile(tt.GoldenFile)
			require.NoError(t, err)

			exc.
				On("Request", "GET", "nodes/test_node/tasks/UPID:test_node:00000000:00000000:00000000:test_action:test_id:test_user:test_extra/status", url.Values(nil)).
				Return(response, nil).
				Once()

			receivedStatus, err := obj.GetStatus()
			require.NoError(t, err)

			assert.Equal(t, tt.TaskStatus, receivedStatus)
		})
	}

	t.Run("Error", func(t *testing.T) {
		exc.
			On("Request", "GET", "nodes/test_node/tasks/UPID:test_node:00000000:00000000:00000000:test_action:test_id:test_user:test_extra/status", url.Values(nil)).
			Return(nil, fmt.Errorf("test_error")).
			Once()

		expectedStatus := types.StatusStopped
		receivedStatus, err := obj.GetStatus()

		assert.Equal(t, expectedStatus, receivedStatus)
		assert.EqualError(t, err, "test_error")
	})
}
