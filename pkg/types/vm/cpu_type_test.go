package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func TestQEMUCPUKind(t *testing.T) {
	QEMUCPUKindCases := map[string](struct {
		Object vm.QEMUCPUKind
		Value  string
	}){
		"Host": {
			Object: vm.QEMUCPUKindHost,
			Value:  "host",
		},
		"Max": {
			Object: vm.QEMUCPUKindMax,
			Value:  "max",
		},
		"QEMU32": {
			Object: vm.QEMUCPUKindQEMU32,
			Value:  "qemu32",
		},
		"QEMU64": {
			Object: vm.QEMUCPUKindQEMU64,
			Value:  "qemu64",
		},
		"KVM32": {
			Object: vm.QEMUCPUKindKVM32,
			Value:  "kvm32",
		},
		"KVM64": {
			Object: vm.QEMUCPUKindKVM64,
			Value:  "kvm64",
		},
		"Intel 486": {
			Object: vm.QEMUCPUKindIntel486,
			Value:  "486",
		},
		"Intel Pentium": {
			Object: vm.QEMUCPUKindIntelPentium,
			Value:  "pentium",
		},
		"Intel Pentium 2": {
			Object: vm.QEMUCPUKindIntelPentium2,
			Value:  "pentium2",
		},
		"Intel Pentium 3": {
			Object: vm.QEMUCPUKindIntelPentium3,
			Value:  "pentium3",
		},
		"Intel Core Duo": {
			Object: vm.QEMUCPUKindIntelCoreDuo,
			Value:  "coreduo",
		},
		"Intel Core 2 Duo": {
			Object: vm.QEMUCPUKindIntelCore2Duo,
			Value:  "core2duo",
		},
		"Intel Conroe": {
			Object: vm.QEMUCPUKindIntelConroe,
			Value:  "Conroe",
		},
		"Intep Penryn": {
			Object: vm.QEMUCPUKindIntelPenryn,
			Value:  "Penryn",
		},
		"Intel Nehalem": {
			Object: vm.QEMUCPUKindIntelNehalem,
			Value:  "Nehalem",
		},
		"Intel Nehalem with IBRS": {
			Object: vm.QEMUCPUKindIntelNehalemIBRS,
			Value:  "Nehalem-IBRS",
		},
		"Intel Westmere": {
			Object: vm.QEMUCPUKindIntelWestmere,
			Value:  "Westmere",
		},
		"Intel Westmere with IBRS": {
			Object: vm.QEMUCPUKindIntelWestmereIBRS,
			Value:  "Westmere-IBRS",
		},
		"Intel Sandy Bridge": {
			Object: vm.QEMUCPUKindIntelSandyBridge,
			Value:  "SandyBridge",
		},
		"Intel Sandy Bridge with IBRS": {
			Object: vm.QEMUCPUKindIntelSandyBridgeIBRS,
			Value:  "SandyBridge-IBRS",
		},
		"Intel Ivy Bridge": {
			Object: vm.QEMUCPUKindIntelIvyBridge,
			Value:  "IvyBridge",
		},
		"Intel Ivy Bridge with IBRS": {
			Object: vm.QEMUCPUKindIntelIvyBridgeIBRS,
			Value:  "IvyBridge-IBRS",
		},
		"Intel Haswell": {
			Object: vm.QEMUCPUKindIntelHaswell,
			Value:  "Haswell",
		},
		"Intel Haswell with IBRS": {
			Object: vm.QEMUCPUKindIntelHaswellIBRS,
			Value:  "Haswell-IBRS",
		},
		"Intel Haswell without TSX": {
			Object: vm.QEMUCPUKindIntelHaswellNoTSX,
			Value:  "Haswell-noTSX",
		},
		"Intel Haswell without TSX and with IBRS": {
			Object: vm.QEMUCPUKindIntelHaswellIBRSNoTSX,
			Value:  "Haswell-noTSX-IBRS",
		},
		"Intel Broadwell": {
			Object: vm.QEMUCPUKindIntelBroadwell,
			Value:  "Broadwell",
		},
		"Intel Broadwell with IBRS": {
			Object: vm.QEMUCPUKindIntelBroadwellIBRS,
			Value:  "Broadwell-IBRS",
		},
		"Intel Broadwell without TSX": {
			Object: vm.QEMUCPUKindIntelBroadwellNoTSX,
			Value:  "Broadwell-noTSX",
		},
		"Intel Broadwell without TSX and with IBRS": {
			Object: vm.QEMUCPUKindIntelBroadwellIBRSNoTSX,
			Value:  "Broadwell-noTSX-IBRS",
		},
		"Intel Sky Lake Client": {
			Object: vm.QEMUCPUKindIntelSkyLakeClient,
			Value:  "Skylake-Client",
		},
		"Intel Sky Lake Client with IBRS": {
			Object: vm.QEMUCPUKindIntelSkyLakeClientIBRS,
			Value:  "Skylake-Client-IBRS",
		},
		"Intel Sky Lake Client without TSX and with IBRS": {
			Object: vm.QEMUCPUKindIntelSkyLakeClientIBRSNoTSX,
			Value:  "Skylake-Client-noTSX-IBRS",
		},
		"Intel Sky Lake Server": {
			Object: vm.QEMUCPUKindIntelSkyLakeServer,
			Value:  "Skylake-Server",
		},
		"Intel Sky Lake Server with IBRS": {
			Object: vm.QEMUCPUKindIntelSkyLakeServerIBRS,
			Value:  "Skylake-Server-IBRS",
		},
		"Intel Sky Lake Server without TSX and with IBRS": {
			Object: vm.QEMUCPUKindIntelSkyLakeServerIBRSNoTSX,
			Value:  "Skylake-Server-noTSX-IBRS",
		},
		"Intel Cascade Lake Server": {
			Object: vm.QEMUCPUKindIntelCascadeLakeServer,
			Value:  "Cascadelake-Server",
		},
		"Intel Cascade Lake Server without TSX": {
			Object: vm.QEMUCPUKindIntelCascadeLakeServerNoTSX,
			Value:  "Cascadelake-Server-noTSX",
		},
		"Intel Knights Mill": {
			Object: vm.QEMUCPUKindIntelKnightsMill,
			Value:  "KnightsMill",
		},
		"Intel Ice Lake Client": {
			Object: vm.QEMUCPUKindIntelIceLakeClient,
			Value:  "Icelake-Client",
		},
		"Intel Ice Lake Client without TSX": {
			Object: vm.QEMUCPUKindIntelIceLakeClientNoTSX,
			Value:  "Icelake-Client-noTSX",
		},
		"Intel Ice Lake Server": {
			Object: vm.QEMUCPUKindIntelIceLakeServer,
			Value:  "Icelake-Server",
		},
		"Intel Ice Lake Server without TSX": {
			Object: vm.QEMUCPUKindIntelIceLakeServerNoTSX,
			Value:  "Icelake-Server-noTSX",
		},
		"AMD Athlon": {
			Object: vm.QEMUCPUKindAMDAthlon,
			Value:  "athlon",
		},
		"AMD Phenom": {
			Object: vm.QEMUCPUKindAMDPhenom,
			Value:  "phenom",
		},
		"AMD Opteron Generation 1": {
			Object: vm.QEMUCPUKindAMDOpteronG1,
			Value:  "Opteron_G1",
		},
		"AMD Opteron Generation 2": {
			Object: vm.QEMUCPUKindAMDOpteronG2,
			Value:  "Opteron_G2",
		},
		"AMD Opteron Generation 3": {
			Object: vm.QEMUCPUKindAMDOpteronG3,
			Value:  "Opteron_G3",
		},
		"AMD Opteron Generation 4": {
			Object: vm.QEMUCPUKindAMDOpteronG4,
			Value:  "Opteron_G4",
		},
		"AMD Opteron Generation 5": {
			Object: vm.QEMUCPUKindAMDOpteronG5,
			Value:  "Opteron_G5",
		},
		"AMD EPYC": {
			Object: vm.QEMUCPUKindEPYC,
			Value:  "EPYC",
		},
		"AMD EPYC with IBPB": {
			Object: vm.QEMUCPUKindEPYCIBPB,
			Value:  "EPYC-IBPB",
		},
		"AMD EPYC-Rome": {
			Object: vm.QEMUCPUKindEPYCRome,
			Value:  "EPYC-Rome",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range QEMUCPUKindCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedObject vm.QEMUCPUKind
				err := (&receivedObject).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedObject)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range QEMUCPUKindCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
