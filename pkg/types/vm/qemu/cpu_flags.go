package qemu

import (
	"encoding/json"
)

type CPUFlags string

const (
	// Required to let the guest OS know if MDS is mitigated correctly
	CPUFlagsEnableMDClear  CPUFlags = "+md-clear"
	CPUFlagsDisableMDClear CPUFlags = "-md-clear"

	// Meltdown fix cost reduction on Westmere, SandyBride and IvyBridge Intel CPUs
	CPUFlagsEnablePCID  CPUFlags = "+pcid"
	CPUFlagsDisablePCID CPUFlags = "-pcid"

	// Allows improved Spectre mitigation with Intel CPUs
	CPUFlagsEnableSpecCtrl  CPUFlags = "+spec-ctrl"
	CPUFlagsDisableSpecCtrl CPUFlags = "-spec-ctrl"

	// Protection for "Speculative Store ByPass" for Intel models
	CPUFlagsEnableSSBD  CPUFlags = "+ssbd"
	CPUFlagsDisableSSBD CPUFlags = "-ssbd"

	// Allows improved Spectre mitigation with AMD CPUs
	CPUFlagsEnableIBPB  CPUFlags = "+ibpb"
	CPUFlagsDisableIBPB CPUFlags = "-ibpb"

	// Basis for "Speculative Store Bypass" protection for AMD models
	CPUFlagsEnableVirtSSBD  CPUFlags = "+virt-ssbd"
	CPUFlagsDisableVirtSSBD CPUFlags = "-virt-ssbd"

	// Improves Spectre mitigation performance with AMD CPUs, best used with "virt-ssbd"
	CPUFlagsEnableAMDSSBD  CPUFlags = "+amd-ssbd"
	CPUFlagsDisableAMDSSBD CPUFlags = "-amd-ssbd"

	// Notifies guest OS that host is not vulnerable for Spectre on AMD CPUs
	CPUFlagsEnableAMDNoSSB  CPUFlags = "+amd-no-ssb"
	CPUFlagsDisableAMDNoSSB CPUFlags = "-amd-no-ssb"

	// Allow guest OS to use 1GB size pages, if host HW supports it
	CPUFlagsEnablePDPE1GB  CPUFlags = "+pdpe1gb"
	CPUFlagsDisablePDPE1GB CPUFlags = "-pdpe1gb"

	// Improve performance in overcommited Windows guests. May lead to guest bluescreens on old CPUs.
	CPUFlagsEnableHVTLBFlush  CPUFlags = "+hv-tlbflush"
	CPUFlagsDisableHVTLBFlush CPUFlags = "-hv-tlbflush"

	// Improve performance for nested virtualization. Only supported on Intel CPUs.
	CPUFlagsEnableHVEVMCS  CPUFlags = "+hv-evmcs"
	CPUFlagsDisableHVEVMCS CPUFlags = "-hv-evmcs"

	// Activate AES instruction set for HW acceleration.
	CPUFlagsEnableAES  CPUFlags = "+aes"
	CPUFlagsDisableAES CPUFlags = "-aes"
)

func (obj CPUFlags) IsValid() bool {
	switch obj {
	case CPUFlagsEnableMDClear,
		CPUFlagsDisableMDClear,
		CPUFlagsEnablePCID,
		CPUFlagsDisablePCID,
		CPUFlagsEnableSpecCtrl,
		CPUFlagsDisableSpecCtrl,
		CPUFlagsEnableSSBD,
		CPUFlagsDisableSSBD,
		CPUFlagsEnableIBPB,
		CPUFlagsDisableIBPB,
		CPUFlagsEnableVirtSSBD,
		CPUFlagsDisableVirtSSBD,
		CPUFlagsEnableAMDSSBD,
		CPUFlagsDisableAMDSSBD,
		CPUFlagsEnableAMDNoSSB,
		CPUFlagsDisableAMDNoSSB,
		CPUFlagsEnablePDPE1GB,
		CPUFlagsDisablePDPE1GB,
		CPUFlagsEnableHVTLBFlush,
		CPUFlagsDisableHVTLBFlush,
		CPUFlagsEnableHVEVMCS,
		CPUFlagsDisableHVEVMCS,
		CPUFlagsEnableAES,
		CPUFlagsDisableAES:
		return true
	default:
		return false
	}
}

func (obj CPUFlags) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj CPUFlags) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *CPUFlags) Unmarshal(s string) error {
	*obj = CPUFlags(s)
	return nil
}

func (obj *CPUFlags) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
