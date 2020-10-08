package vm_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUCPUKind(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.QEMUCPUKind)(nil), map[string](struct {
		Object types.FixedValue
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
			Object: vm.QEMUCPUKindAMDEPYC,
			Value:  "EPYC",
		},
		"AMD EPYC with IBPB": {
			Object: vm.QEMUCPUKindAMDEPYCIBPB,
			Value:  "EPYC-IBPB",
		},
		"AMD EPYC-Rome": {
			Object: vm.QEMUCPUKindAMDEPYCRome,
			Value:  "EPYC-Rome",
		},
	})
}

func TestQEMUCPUFlags(t *testing.T) {
	test.HelperTestFixedValue(t, (*vm.QEMUCPUFlags)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"EnableMDClear": {
			Object: vm.QEMUCPUFlagsEnableMDClear,
			Value:  "+md-clear",
		},
		"DisableMDClear": {
			Object: vm.QEMUCPUFlagsDisableMDClear,
			Value:  "-md-clear",
		},
		"EnablePCID": {
			Object: vm.QEMUCPUFlagsEnablePCID,
			Value:  "+pcid",
		},
		"DisablePCID": {
			Object: vm.QEMUCPUFlagsDisablePCID,
			Value:  "-pcid",
		},
		"EnableSpecCtrl": {
			Object: vm.QEMUCPUFlagsEnableSpecCtrl,
			Value:  "+spec-ctrl",
		},
		"DisableSpecCtrl": {
			Object: vm.QEMUCPUFlagsDisableSpecCtrl,
			Value:  "-spec-ctrl",
		},
		"EnableSSBD": {
			Object: vm.QEMUCPUFlagsEnableSSBD,
			Value:  "+ssbd",
		},
		"DisableSSBD": {
			Object: vm.QEMUCPUFlagsDisableSSBD,
			Value:  "-ssbd",
		},
		"EnableIBPB": {
			Object: vm.QEMUCPUFlagsEnableIBPB,
			Value:  "+ibpb",
		},
		"DisableIBPB": {
			Object: vm.QEMUCPUFlagsDisableIBPB,
			Value:  "-ibpb",
		},
		"EnableVirtSSBD": {
			Object: vm.QEMUCPUFlagsEnableVirtSSBD,
			Value:  "+virt-ssbd",
		},
		"DisableVirtSSBD": {
			Object: vm.QEMUCPUFlagsDisableVirtSSBD,
			Value:  "-virt-ssbd",
		},
		"EnableAMDSSBD": {
			Object: vm.QEMUCPUFlagsEnableAMDSSBD,
			Value:  "+amd-ssbd",
		},
		"DisableAMDSSBD": {
			Object: vm.QEMUCPUFlagsDisableAMDSSBD,
			Value:  "-amd-ssbd",
		},
		"EnableAMDNoSSB": {
			Object: vm.QEMUCPUFlagsEnableAMDNoSSB,
			Value:  "+amd-no-ssb",
		},
		"DisableAMDNoSSB": {
			Object: vm.QEMUCPUFlagsDisableAMDNoSSB,
			Value:  "-amd-no-ssb",
		},
		"EnablePDPE1GB": {
			Object: vm.QEMUCPUFlagsEnablePDPE1GB,
			Value:  "+pdpe1gb",
		},
		"DisablePDPE1GB": {
			Object: vm.QEMUCPUFlagsDisablePDPE1GB,
			Value:  "-pdpe1gb",
		},
		"EnableHVTLBFlush": {
			Object: vm.QEMUCPUFlagsEnableHVTLBFlush,
			Value:  "+hv-tlbflush",
		},
		"DisableHVTLBFlush": {
			Object: vm.QEMUCPUFlagsDisableHVTLBFlush,
			Value:  "-hv-tlbflush",
		},
		"EnableHVEVMCS": {
			Object: vm.QEMUCPUFlagsEnableHVEVMCS,
			Value:  "+hv-evmcs",
		},
		"DisableHVEVMCS": {
			Object: vm.QEMUCPUFlagsDisableHVEVMCS,
			Value:  "-hv-evmcs",
		},
		"EnableEnableAES": {
			Object: vm.QEMUCPUFlagsEnableAES,
			Value:  "+aes",
		},
		"DisableAES": {
			Object: vm.QEMUCPUFlagsDisableAES,
			Value:  "-aes",
		},
	})
}
