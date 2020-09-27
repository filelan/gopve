package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

var NFSVersionCases = map[string](struct {
	Object storage.NFSVersion
	Value  string
}){
	"3": {
		Object: storage.NFSVersion30,
		Value:  "3",
	},
	"4": {
		Object: storage.NFSVersion40,
		Value:  "4",
	},
	"4.1": {
		Object: storage.NFSVersion41,
		Value:  "4.1",
	},
	"4.2": {
		Object: storage.NFSVersion42,
		Value:  "4.2",
	},
}

func TestKindNFSVersionMarshal(t *testing.T) {
	for n, tt := range NFSVersionCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			var receivedTransport storage.NFSVersion
			err := (&receivedTransport).Unmarshal(tt.Value)
			require.NoError(t, err)
			assert.Equal(t, tt.Object, receivedTransport)
		})
	}
}

func TestKindNFSVersionUnmarshal(t *testing.T) {
	for n, tt := range NFSVersionCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			receivedValue, err := tt.Object.Marshal()
			require.NoError(t, err)
			assert.Equal(t, tt.Value, receivedValue)
		})
	}
}
