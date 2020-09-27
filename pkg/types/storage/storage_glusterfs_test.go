package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

var GlusterFSTransportCases = map[string](struct {
	Object storage.GlusterFSTransport
	Value  string
}){
	"TCP": {
		Object: storage.GlusterFSTransportTCP,
		Value:  "tcp",
	},
	"UNIX": {
		Object: storage.GlusterFSTransportUNIX,
		Value:  "unix",
	},
	"RDMA": {
		Object: storage.GlusterFSTransportRDMA,
		Value:  "rdma",
	},
}

func TestKindGlusterFSTransportMarshal(t *testing.T) {
	for n, tt := range GlusterFSTransportCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			var receivedTransport storage.GlusterFSTransport
			err := (&receivedTransport).Unmarshal(tt.Value)
			require.NoError(t, err)
			assert.Equal(t, tt.Object, receivedTransport)
		})
	}
}

func TestKindGlusterFSTransportUnmarshal(t *testing.T) {
	for n, tt := range GlusterFSTransportCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			receivedValue, err := tt.Object.Marshal()
			require.NoError(t, err)
			assert.Equal(t, tt.Value, receivedValue)
		})
	}
}
