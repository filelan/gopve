package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

var SMBVersionCases = map[string](struct {
	Object storage.SMBVersion
	Value  string
}){
	"2.0": {
		Object: storage.SMBVersion20,
		Value:  "2.0",
	},
	"2.1": {
		Object: storage.SMBVersion21,
		Value:  "2.1",
	},
	"3.0": {
		Object: storage.SMBVersion30,
		Value:  "3.0",
	},
}

func TestKindSMBVersionMarshal(t *testing.T) {
	for n, tt := range SMBVersionCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			var receivedTransport storage.SMBVersion
			err := (&receivedTransport).Unmarshal(tt.Value)
			require.NoError(t, err)
			assert.Equal(t, tt.Object, receivedTransport)
		})
	}
}

func TestKindSMBVersionUnmarshal(t *testing.T) {
	for n, tt := range SMBVersionCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			receivedValue, err := tt.Object.Marshal()
			require.NoError(t, err)
			assert.Equal(t, tt.Value, receivedValue)
		})
	}
}
