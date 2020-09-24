package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/task"
	"github.com/xabinapal/gopve/internal/service/task/test"
	types "github.com/xabinapal/gopve/pkg/types/task"
)

func TestServiceGet(t *testing.T) {
	svc, _, _ := test.NewService()

	t.Run("Valid", func(t *testing.T) {
		expectedTask := task.NewTask(svc, "test_node", "00000000:00000000:00000000", "test_action", "test_id", "root@pam", "")

		receivedTask, err := svc.Get("UPID:test_node:00000000:00000000:00000000:test_action:test_id:root@pam:")
		require.NoError(t, err)

		assert.Equal(t, expectedTask, receivedTask)
	})

	options := map[string]string{
		"InvalidLength":  "UPID:test_node::::test_action:test_id:root@pam",
		"InvalidContent": ":test_node::::test_action:test_id:root@pam:",
	}

	for n, tt := range options {
		tt := tt

		t.Run(n, func(t *testing.T) {
			_, err := svc.Get(tt)
			assert.EqualError(t, err, types.ErrInvalidUPID.Error())
		})
	}
}

func TestServiceGetSpecialized(t *testing.T) {
	svc, _, _ := test.NewService()

	type x struct {
		UPID string
		Type interface{}
	}
	options := map[string]x{
		"QMCreate": {
			"UPID:test_node:00000000:00000000:00000000:qmcreate:100:root@pam:",
			(*types.VirtualMachineTask)(nil),
		},
		"VZCreate": {
			"UPID:test_node:00000000:00000000:00000000:vzcreate:100:root@pam:",
			(*types.VirtualMachineTask)(nil),
		},
	}

	for n, tt := range options {
		tt := tt

		t.Run(n, func(t *testing.T) {
			receivedTask, err := svc.Get(tt.UPID)
			require.NoError(t, err)

			assert.Implements(t, tt.Type, receivedTask)
		})
	}
}
