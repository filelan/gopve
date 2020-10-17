package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestCPUType(t *testing.T) {
	test.HelperTestFixedValue(t, (*qemu.CPUType)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"Host": {
			Object: qemu.CPUTypeHost,
			Value:  "host",
		},
		"Max": {
			Object: qemu.CPUTypeMax,
			Value:  "max",
		},
		"QEMU32": {
			Object: qemu.CPUTypeQEMU32,
			Value:  "qemu32",
		},
		"QEMU64": {
			Object: qemu.CPUTypeQEMU64,
			Value:  "qemu64",
		},
		"KVM32": {
			Object: qemu.CPUTypeKVM32,
			Value:  "kvm32",
		},
		"KVM64": {
			Object: qemu.CPUTypeKVM64,
			Value:  "kvm64",
		},
		"Intel 486": {
			Object: qemu.CPUTypeIntel486,
			Value:  "486",
		},
		"Intel Pentium": {
			Object: qemu.CPUTypeIntelPentium,
			Value:  "pentium",
		},
		"Intel Pentium 2": {
			Object: qemu.CPUTypeIntelPentium2,
			Value:  "pentium2",
		},
		"Intel Pentium 3": {
			Object: qemu.CPUTypeIntelPentium3,
			Value:  "pentium3",
		},
		"Intel Core Duo": {
			Object: qemu.CPUTypeIntelCoreDuo,
			Value:  "coreduo",
		},
		"Intel Core 2 Duo": {
			Object: qemu.CPUTypeIntelCore2Duo,
			Value:  "core2duo",
		},
		"Intel Conroe": {
			Object: qemu.CPUTypeIntelConroe,
			Value:  "Conroe",
		},
		"Intep Penryn": {
			Object: qemu.CPUTypeIntelPenryn,
			Value:  "Penryn",
		},
		"Intel Nehalem": {
			Object: qemu.CPUTypeIntelNehalem,
			Value:  "Nehalem",
		},
		"Intel Nehalem with IBRS": {
			Object: qemu.CPUTypeIntelNehalemIBRS,
			Value:  "Nehalem-IBRS",
		},
		"Intel Westmere": {
			Object: qemu.CPUTypeIntelWestmere,
			Value:  "Westmere",
		},
		"Intel Westmere with IBRS": {
			Object: qemu.CPUTypeIntelWestmereIBRS,
			Value:  "Westmere-IBRS",
		},
		"Intel Sandy Bridge": {
			Object: qemu.CPUTypeIntelSandyBridge,
			Value:  "SandyBridge",
		},
		"Intel Sandy Bridge with IBRS": {
			Object: qemu.CPUTypeIntelSandyBridgeIBRS,
			Value:  "SandyBridge-IBRS",
		},
		"Intel Ivy Bridge": {
			Object: qemu.CPUTypeIntelIvyBridge,
			Value:  "IvyBridge",
		},
		"Intel Ivy Bridge with IBRS": {
			Object: qemu.CPUTypeIntelIvyBridgeIBRS,
			Value:  "IvyBridge-IBRS",
		},
		"Intel Haswell": {
			Object: qemu.CPUTypeIntelHaswell,
			Value:  "Haswell",
		},
		"Intel Haswell with IBRS": {
			Object: qemu.CPUTypeIntelHaswellIBRS,
			Value:  "Haswell-IBRS",
		},
		"Intel Haswell without TSX": {
			Object: qemu.CPUTypeIntelHaswellNoTSX,
			Value:  "Haswell-noTSX",
		},
		"Intel Haswell without TSX and with IBRS": {
			Object: qemu.CPUTypeIntelHaswellIBRSNoTSX,
			Value:  "Haswell-noTSX-IBRS",
		},
		"Intel Broadwell": {
			Object: qemu.CPUTypeIntelBroadwell,
			Value:  "Broadwell",
		},
		"Intel Broadwell with IBRS": {
			Object: qemu.CPUTypeIntelBroadwellIBRS,
			Value:  "Broadwell-IBRS",
		},
		"Intel Broadwell without TSX": {
			Object: qemu.CPUTypeIntelBroadwellNoTSX,
			Value:  "Broadwell-noTSX",
		},
		"Intel Broadwell without TSX and with IBRS": {
			Object: qemu.CPUTypeIntelBroadwellIBRSNoTSX,
			Value:  "Broadwell-noTSX-IBRS",
		},
		"Intel Sky Lake Client": {
			Object: qemu.CPUTypeIntelSkyLakeClient,
			Value:  "Skylake-Client",
		},
		"Intel Sky Lake Client with IBRS": {
			Object: qemu.CPUTypeIntelSkyLakeClientIBRS,
			Value:  "Skylake-Client-IBRS",
		},
		"Intel Sky Lake Client without TSX and with IBRS": {
			Object: qemu.CPUTypeIntelSkyLakeClientIBRSNoTSX,
			Value:  "Skylake-Client-noTSX-IBRS",
		},
		"Intel Sky Lake Server": {
			Object: qemu.CPUTypeIntelSkyLakeServer,
			Value:  "Skylake-Server",
		},
		"Intel Sky Lake Server with IBRS": {
			Object: qemu.CPUTypeIntelSkyLakeServerIBRS,
			Value:  "Skylake-Server-IBRS",
		},
		"Intel Sky Lake Server without TSX and with IBRS": {
			Object: qemu.CPUTypeIntelSkyLakeServerIBRSNoTSX,
			Value:  "Skylake-Server-noTSX-IBRS",
		},
		"Intel Cascade Lake Server": {
			Object: qemu.CPUTypeIntelCascadeLakeServer,
			Value:  "Cascadelake-Server",
		},
		"Intel Cascade Lake Server without TSX": {
			Object: qemu.CPUTypeIntelCascadeLakeServerNoTSX,
			Value:  "Cascadelake-Server-noTSX",
		},
		"Intel Knights Mill": {
			Object: qemu.CPUTypeIntelKnightsMill,
			Value:  "KnightsMill",
		},
		"Intel Ice Lake Client": {
			Object: qemu.CPUTypeIntelIceLakeClient,
			Value:  "Icelake-Client",
		},
		"Intel Ice Lake Client without TSX": {
			Object: qemu.CPUTypeIntelIceLakeClientNoTSX,
			Value:  "Icelake-Client-noTSX",
		},
		"Intel Ice Lake Server": {
			Object: qemu.CPUTypeIntelIceLakeServer,
			Value:  "Icelake-Server",
		},
		"Intel Ice Lake Server without TSX": {
			Object: qemu.CPUTypeIntelIceLakeServerNoTSX,
			Value:  "Icelake-Server-noTSX",
		},
		"AMD Athlon": {
			Object: qemu.CPUTypeAMDAthlon,
			Value:  "athlon",
		},
		"AMD Phenom": {
			Object: qemu.CPUTypeAMDPhenom,
			Value:  "phenom",
		},
		"AMD Opteron Generation 1": {
			Object: qemu.CPUTypeAMDOpteronG1,
			Value:  "Opteron_G1",
		},
		"AMD Opteron Generation 2": {
			Object: qemu.CPUTypeAMDOpteronG2,
			Value:  "Opteron_G2",
		},
		"AMD Opteron Generation 3": {
			Object: qemu.CPUTypeAMDOpteronG3,
			Value:  "Opteron_G3",
		},
		"AMD Opteron Generation 4": {
			Object: qemu.CPUTypeAMDOpteronG4,
			Value:  "Opteron_G4",
		},
		"AMD Opteron Generation 5": {
			Object: qemu.CPUTypeAMDOpteronG5,
			Value:  "Opteron_G5",
		},
		"AMD EPYC": {
			Object: qemu.CPUTypeAMDEPYC,
			Value:  "EPYC",
		},
		"AMD EPYC with IBPB": {
			Object: qemu.CPUTypeAMDEPYCIBPB,
			Value:  "EPYC-IBPB",
		},
		"AMD EPYC-Rome": {
			Object: qemu.CPUTypeAMDEPYCRome,
			Value:  "EPYC-Rome",
		},
	})
}
