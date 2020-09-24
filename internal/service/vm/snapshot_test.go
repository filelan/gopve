package vm_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	task "github.com/xabinapal/gopve/internal/service/task/test"
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/internal/service/vm/test"
	types "github.com/xabinapal/gopve/pkg/types/vm"
)

func TestVirtualMachineSnapshot(t *testing.T) {
	virtualMachine, api, exc := test.NewVirtualMachine()

	t.Run("List", func(t *testing.T) {
		response, err := ioutil.ReadFile(
			"./testdata/get_nodes_{node}_{kind}_{vmid}_snapshot.json",
		)
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/snapshot", url.Values(nil)).
			Return(response, nil).
			Once()

		loc, err := time.LoadLocation("UTC")
		require.NoError(t, err)

		expectedSnapshots := []types.Snapshot{
			vm.NewSnapshot(
				virtualMachine,
				"first",
				"First snapshot",
				time.Unix(1609372800, 0).In(loc),
				true,
				"",
			),
			vm.NewSnapshot(
				virtualMachine,
				"second",
				"Second snapshot",
				time.Unix(1609416000, 0).In(loc),
				false,
				"first",
			),
			vm.NewCurrentSnapshot(virtualMachine, "first"),
		}

		snapshots, err := virtualMachine.ListSnapshots()
		require.NoError(t, err)
		assert.ElementsMatch(t, expectedSnapshots, snapshots)

		exc.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		response, err := ioutil.ReadFile(
			"./testdata/get_nodes_{node}_{kind}_{vmid}_snapshot.json",
		)
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/snapshot", url.Values(nil)).
			Return(response, nil).
			Once()

		loc, err := time.LoadLocation("UTC")
		require.NoError(t, err)

		expectedSnapshot := vm.NewSnapshot(
			virtualMachine,
			"first",
			"First snapshot",
			time.Unix(1609372800, 0).In(loc),
			true,
			"",
		)

		snapshot, err := virtualMachine.GetSnapshot("first")
		require.NoError(t, err)
		assert.Equal(t, expectedSnapshot, snapshot)

		exc.AssertExpectations(t)
	})

	t.Run("Create", func(t *testing.T) {
		exc.
			On("Request", http.MethodPost, "nodes/test_node/test_kind/100/snapshot", url.Values{
				"snapname":    {"first"},
				"description": {"First snapshot"},
			}).
			Return(
				[]byte(
					"{\"data\":\"UPID:test_node::::qmsnapshot:100:root@pam:\"}",
				),
				nil,
			).
			Once()

		expectedTask, _, _ := task.NewTask(
			"test_node",
			"::",
			"qmsnapshot",
			"100",
			"root@pam",
			"",
		)

		api.TaskService.
			On("Get", "UPID:test_node::::qmsnapshot:100:root@pam:").
			Return(expectedTask, nil)

		task, err := virtualMachine.CreateSnapshot(
			"first",
			types.SnapshotProperties{
				Description: "First snapshot",
			},
		)
		require.NoError(t, err)
		assert.Equal(t, expectedTask, task)

		exc.AssertExpectations(t)
	})

	t.Run("Rollback", func(t *testing.T) {
		exc.
			On("Request", http.MethodPost, "nodes/test_node/test_kind/100/snapshot/first/rollback", url.Values(nil)).
			Return(
				[]byte(
					"{\"data\":\"UPID:test_node::::qmrollback:100:root@pam:\"}",
				),
				nil,
			).
			Once()

		expectedTask, _, _ := task.NewTask(
			"test_node",
			"::",
			"qmrollback",
			"100",
			"root@pam",
			"",
		)

		api.TaskService.
			On("Get", "UPID:test_node::::qmrollback:100:root@pam:").
			Return(expectedTask, nil)

		task, err := virtualMachine.RollbackToSnapshot("first")
		require.NoError(t, err)
		assert.Equal(t, expectedTask, task)

		exc.AssertExpectations(t)
	})
}
