package qemu_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestCPUFlags(t *testing.T) {
	test.HelperTestFixedValue(t, (*qemu.CPUFlags)(nil), map[string](struct {
		Object types.FixedValue
		Value  string
	}){
		"EnableMDClear": {
			Object: qemu.CPUFlagsEnableMDClear,
			Value:  "+md-clear",
		},
		"DisableMDClear": {
			Object: qemu.CPUFlagsDisableMDClear,
			Value:  "-md-clear",
		},
		"EnablePCID": {
			Object: qemu.CPUFlagsEnablePCID,
			Value:  "+pcid",
		},
		"DisablePCID": {
			Object: qemu.CPUFlagsDisablePCID,
			Value:  "-pcid",
		},
		"EnableSpecCtrl": {
			Object: qemu.CPUFlagsEnableSpecCtrl,
			Value:  "+spec-ctrl",
		},
		"DisableSpecCtrl": {
			Object: qemu.CPUFlagsDisableSpecCtrl,
			Value:  "-spec-ctrl",
		},
		"EnableSSBD": {
			Object: qemu.CPUFlagsEnableSSBD,
			Value:  "+ssbd",
		},
		"DisableSSBD": {
			Object: qemu.CPUFlagsDisableSSBD,
			Value:  "-ssbd",
		},
		"EnableIBPB": {
			Object: qemu.CPUFlagsEnableIBPB,
			Value:  "+ibpb",
		},
		"DisableIBPB": {
			Object: qemu.CPUFlagsDisableIBPB,
			Value:  "-ibpb",
		},
		"EnableVirtSSBD": {
			Object: qemu.CPUFlagsEnableVirtSSBD,
			Value:  "+virt-ssbd",
		},
		"DisableVirtSSBD": {
			Object: qemu.CPUFlagsDisableVirtSSBD,
			Value:  "-virt-ssbd",
		},
		"EnableAMDSSBD": {
			Object: qemu.CPUFlagsEnableAMDSSBD,
			Value:  "+amd-ssbd",
		},
		"DisableAMDSSBD": {
			Object: qemu.CPUFlagsDisableAMDSSBD,
			Value:  "-amd-ssbd",
		},
		"EnableAMDNoSSB": {
			Object: qemu.CPUFlagsEnableAMDNoSSB,
			Value:  "+amd-no-ssb",
		},
		"DisableAMDNoSSB": {
			Object: qemu.CPUFlagsDisableAMDNoSSB,
			Value:  "-amd-no-ssb",
		},
		"EnablePDPE1GB": {
			Object: qemu.CPUFlagsEnablePDPE1GB,
			Value:  "+pdpe1gb",
		},
		"DisablePDPE1GB": {
			Object: qemu.CPUFlagsDisablePDPE1GB,
			Value:  "-pdpe1gb",
		},
		"EnableHVTLBFlush": {
			Object: qemu.CPUFlagsEnableHVTLBFlush,
			Value:  "+hv-tlbflush",
		},
		"DisableHVTLBFlush": {
			Object: qemu.CPUFlagsDisableHVTLBFlush,
			Value:  "-hv-tlbflush",
		},
		"EnableHVEVMCS": {
			Object: qemu.CPUFlagsEnableHVEVMCS,
			Value:  "+hv-evmcs",
		},
		"DisableHVEVMCS": {
			Object: qemu.CPUFlagsDisableHVEVMCS,
			Value:  "-hv-evmcs",
		},
		"EnableEnableAES": {
			Object: qemu.CPUFlagsEnableAES,
			Value:  "+aes",
		},
		"DisableAES": {
			Object: qemu.CPUFlagsDisableAES,
			Value:  "-aes",
		},
	})
}
