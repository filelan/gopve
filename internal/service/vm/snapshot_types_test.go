package vm_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/internal/service/vm/test"
	types "github.com/xabinapal/gopve/pkg/types/vm"
)

func TestSnapshot(t *testing.T) {
	virtualMachine, _, exc := test.NewVirtualMachine()

	loc, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	getSnapshot := func(parent string) *vm.Snapshot {
		return vm.NewSnapshot(virtualMachine, "test_snapshot", "test_description", time.Unix(1609372800, 0).In(loc), true, parent)
	}

	t.Run("Name", func(t *testing.T) {
		snapshot := getSnapshot("")
		expectedName := "test_snapshot"
		name := snapshot.Name()
		assert.Equal(t, expectedName, name)
	})

	t.Run("Description", func(t *testing.T) {
		snapshot := getSnapshot("")
		expectedDescription := "test_description"
		description := snapshot.Description()
		assert.Equal(t, expectedDescription, description)
	})

	t.Run("Timestamp", func(t *testing.T) {
		snapshot := getSnapshot("")
		expectedTimestamp := time.Unix(1609372800, 0).In(loc)
		timestamp := snapshot.Timestamp()
		assert.Equal(t, expectedTimestamp, timestamp)
	})

	t.Run("Digest", func(t *testing.T) {
		snapshot := getSnapshot("")
		expectedWithRAM := true
		withRAM := snapshot.WithRAM()
		assert.Equal(t, expectedWithRAM, withRAM)
	})

	t.Run("Parent", func(t *testing.T) {
		snapshot := getSnapshot("first")
		expectedParent := "first"
		parent := snapshot.Parent()
		assert.Equal(t, expectedParent, parent)
	})

	t.Run("GetParent", func(t *testing.T) {
		snapshot := getSnapshot("first")

		response, err := ioutil.ReadFile("./testdata/get_nodes_{node}_{kind}_{vmid}_snapshot.json")
		require.NoError(t, err)

		exc.
			On("Request", http.MethodGet, "nodes/test_node/test_kind/100/snapshot", url.Values(nil)).
			Return(response, nil).
			Once()

		expectedParent := vm.NewSnapshot(virtualMachine, "first", "First snapshot", time.Unix(1609372800, 0).In(loc), true, "")

		parent, err := snapshot.GetParent()
		require.NoError(t, err)
		assert.Equal(t, expectedParent, parent)
	})

	t.Run("GetRootParent", func(t *testing.T) {
		snapshot := getSnapshot("")
		_, err := snapshot.GetParent()
		require.EqualError(t, err, types.ErrRootParentSnapshot.Error())
	})

	t.Run("GetProperties", func(t *testing.T) {
		snapshot := getSnapshot("")

		expectedProperties := types.SnapshotProperties{
			Description: "test_description",
		}

		properties, err := snapshot.GetProperties()
		require.NoError(t, err)
		assert.Equal(t, expectedProperties, properties)
	})

	t.Run("SetProperties", func(t *testing.T) {
		snapshot := getSnapshot("")

		exc.
			On("Request", http.MethodPut, "nodes/test_node/test_kind/100/snapshot/test_snapshot/config", url.Values{
				"description": {"new_description"},
			}).
			Return(nil, nil).
			Once()

		err := snapshot.SetProperties(types.SnapshotProperties{
			Description: "new_description",
		})
		require.NoError(t, err)

		assert.Equal(t, "new_description", snapshot.Description())

		exc.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		snapshot := getSnapshot("")

		exc.
			On("Request", http.MethodDelete, "nodes/test_node/test_kind/100/snapshot/test_snapshot", url.Values(nil)).
			Return(nil, nil).
			Once()

		err := snapshot.Delete()
		require.NoError(t, err)

		exc.AssertExpectations(t)
	})

	t.Run("DeleteCurrent", func(t *testing.T) {
		snapshot := vm.NewCurrentSnapshot(virtualMachine, "")

		err := snapshot.Delete()
		require.EqualError(t, err, types.ErrDeleteCurrentSnapshot.Error())

		exc.AssertExpectations(t)
	})
}
