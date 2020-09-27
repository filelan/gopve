package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

var ImageFormatCases = map[string](struct {
	Object storage.ImageFormat
	Value  string
}){
	"RAW": {
		Object: storage.ImageFormatRaw,
		Value:  "raw",
	},
	"QCOW2": {
		Object: storage.ImageFormatQcow2,
		Value:  "qcow2",
	},
	"VMDK": {
		Object: storage.ImageFormatVMDK,
		Value:  "vmdk",
	},
	"SubVolume": {
		Object: storage.ImageFormatSubVolume,
		Value:  "subvol",
	},
}

func TestKindImageFormatMarshal(t *testing.T) {
	for n, tt := range ImageFormatCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			var receivedFormat storage.ImageFormat
			err := (&receivedFormat).Unmarshal(tt.Value)
			require.NoError(t, err)
			assert.Equal(t, tt.Object, receivedFormat)
		})
	}
}

func TestKindImageFormatUnmarshal(t *testing.T) {
	for n, tt := range ImageFormatCases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			receivedValue, err := tt.Object.Marshal()
			require.NoError(t, err)
			assert.Equal(t, tt.Value, receivedValue)
		})
	}
}
