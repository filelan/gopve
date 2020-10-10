package vm

import (
	"encoding/json"
)

type QEMUCPUKind string

const (
	// KVM processor with all supported host features.
	QEMUCPUKindHost QEMUCPUKind = "host"

	// Enables all features supported by the accelerator in the current host.
	QEMUCPUKindMax QEMUCPUKind = "max"

	// QEMU Virtual CPU version 2.5+ (32 bit variant)
	QEMUCPUKindQEMU32 QEMUCPUKind = "qemu32"

	// QEMU Virtual CPU version 2.5+ (64 bit variant)
	QEMUCPUKindQEMU64 QEMUCPUKind = "qemu64"

	// Common KVM processor (32 bit variant). Legacy model just for historical compatibility with ancient QEMU versions.
	QEMUCPUKindKVM32 QEMUCPUKind = "kvm32"

	// Common KVM processor (64 bit variant). Legacy model just for historical compatibility with ancient QEMU versions.
	QEMUCPUKindKVM64 QEMUCPUKind = "kvm64"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntel486 QEMUCPUKind = "486"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelPentium QEMUCPUKind = "pentium"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelPentium2 QEMUCPUKind = "pentium2"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelPentium3 QEMUCPUKind = "pentium3"

	// Intel CPU T2600 @ 2.16. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelCoreDuo QEMUCPUKind = "coreduo"

	// Intel Core 2 Duo CPU T7700  @ 2.40GHz. Old Intel x86 model, its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelCore2Duo QEMUCPUKind = "core2duo"

	// Intel Celeron_4x0 (Conroe/Merom Class Core 2, 2006)
	QEMUCPUKindIntelConroe QEMUCPUKind = "Conroe"

	// Intel Core 2 Duo P9xxx (Penryn Class Core 2, 2007)
	QEMUCPUKindIntelPenryn QEMUCPUKind = "Penryn"

	// Intel Core i7 9xx (Nehalem Class Core i7, 2008)
	QEMUCPUKindIntelNehalem QEMUCPUKind = "Nehalem"

	// Intel Core i7 9xx (Nehalem Class Core i7, 2008, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelNehalemIBRS QEMUCPUKind = "Nehalem-IBRS"

	// Westmere E56xx/L56xx/X56xx (Nehalem-C, 2010)
	QEMUCPUKindIntelWestmere QEMUCPUKind = "Westmere"

	// Westmere E56xx/L56xx/X56xx (Nehalem-C, 2010, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelWestmereIBRS QEMUCPUKind = "Westmere-IBRS"

	// Intel Xeon E312xx (Sandy Bridge, 2011)
	QEMUCPUKindIntelSandyBridge QEMUCPUKind = "SandyBridge"

	// Intel Xeon E312xx (Sandy Bridge, 2011, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelSandyBridgeIBRS QEMUCPUKind = "SandyBridge-IBRS"

	// Intel Xeon E3-12xx v2 (Ivy Bridge, 2012)
	QEMUCPUKindIntelIvyBridge QEMUCPUKind = "IvyBridge"

	// Intel Xeon E3-12xx v2 (Ivy Bridge, 2012, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelIvyBridgeIBRS QEMUCPUKind = "IvyBridge-IBRS"

	// Intel Core Processor (Haswell, 2013)
	QEMUCPUKindIntelHaswell QEMUCPUKind = "Haswell"

	// Intel Core Processor (Haswell, 2013, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelHaswellIBRS QEMUCPUKind = "Haswell-IBRS"

	// Intel Core Processor (Haswell, 2013, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelHaswellNoTSX QEMUCPUKind = "Haswell-noTSX"

	// Intel Core Processor (Haswell, 2013, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelHaswellIBRSNoTSX QEMUCPUKind = "Haswell-noTSX-IBRS"

	// Intel Core Processor (Broadwell, 2014)
	QEMUCPUKindIntelBroadwell QEMUCPUKind = "Broadwell"

	// Intel Core Processor (Broadwell, 2014, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelBroadwellIBRS QEMUCPUKind = "Broadwell-IBRS"

	// Intel Core Processor (Broadwell, 2014, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelBroadwellNoTSX QEMUCPUKind = "Broadwell-noTSX"

	// Intel Core Processor (Broadwell, 2014, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelBroadwellIBRSNoTSX QEMUCPUKind = "Broadwell-noTSX-IBRS"

	// Intel Core Processor (Skylake, 2015)
	QEMUCPUKindIntelSkyLakeClient QEMUCPUKind = "Skylake-Client"

	// Intel Core Processor (Skylake, 2015, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelSkyLakeClientIBRS QEMUCPUKind = "Skylake-Client-IBRS"

	// Intel Core Processor (Skylake, 2015, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelSkyLakeClientIBRSNoTSX QEMUCPUKind = "Skylake-Client-noTSX-IBRS"

	// Intel Xeon Processor (Skylake, 2016)
	QEMUCPUKindIntelSkyLakeServer QEMUCPUKind = "Skylake-Server"

	// Intel Xeon Processor (Skylake, 2016, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelSkyLakeServerIBRS QEMUCPUKind = "Skylake-Server-IBRS"

	// Intel Xeon Processor (Skylake, 2016, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelSkyLakeServerIBRSNoTSX QEMUCPUKind = "Skylake-Server-noTSX-IBRS"

	// Intel Xeon Processor (Cascade Lake, 2019, with “stepping” levels 6 or 7 only). The Cascade Lake Xeon processor with stepping 5 is vulnerable to MDS variants.
	QEMUCPUKindIntelCascadeLakeServer QEMUCPUKind = "Cascadelake-Server"

	// Intel Xeon Processor (Cascade Lake, 2019, with “stepping” levels 6 or 7 only, without Transactional Synchronization Extensions update). The Cascade Lake Xeon processor with stepping 5 is vulnerable to MDS variants.
	QEMUCPUKindIntelCascadeLakeServerNoTSX QEMUCPUKind = "Cascadelake-Server-noTSX"

	// Intel Xeon Phi Processor (Knights Mill)
	QEMUCPUKindIntelKnightsMill QEMUCPUKind = "KnightsMill"

	// Intel Core Processor (Icelake)
	QEMUCPUKindIntelIceLakeClient QEMUCPUKind = "Icelake-Client"

	// Intel Core Processor (Icelake, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelIceLakeClientNoTSX QEMUCPUKind = "Icelake-Client-noTSX"

	// Intel Xeon Processor (Icelake)
	QEMUCPUKindIntelIceLakeServer QEMUCPUKind = "Icelake-Server"

	// Intel Xeon Processor (Icelake, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelIceLakeServerNoTSX QEMUCPUKind = "Icelake-Server-noTSX"

	// Old AMD x86 model, its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindAMDAthlon QEMUCPUKind = "athlon"

	// AMD Phenom(tm) 9550 Quad-Core Processor. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindAMDPhenom QEMUCPUKind = "phenom"

	// AMD Opteron 240 (Gen 1 Class Opteron, 2004)
	QEMUCPUKindAMDOpteronG1 QEMUCPUKind = "Opteron_G1"

	// AMD Opteron 22xx (Gen 2 Class Opteron, 2006)
	QEMUCPUKindAMDOpteronG2 QEMUCPUKind = "Opteron_G2"

	// AMD Opteron 23xx (Gen 3 Class Opteron, 2009)
	QEMUCPUKindAMDOpteronG3 QEMUCPUKind = "Opteron_G3"

	// AMD Opteron 62xx class CPU (2011)
	QEMUCPUKindAMDOpteronG4 QEMUCPUKind = "Opteron_G4"

	// AMD Opteron 63xx class CPU (2012)
	QEMUCPUKindAMDOpteronG5 QEMUCPUKind = "Opteron_G5"

	// AMD EPYC Processor (2017)
	QEMUCPUKindAMDEPYC QEMUCPUKind = "EPYC"

	// AMD EPYC Processor (2017, with Indirect Branch Prediction Barrier update)
	QEMUCPUKindAMDEPYCIBPB QEMUCPUKind = "EPYC-IBPB"

	// AMD EPYC-Rome Processor (2019)
	QEMUCPUKindAMDEPYCRome QEMUCPUKind = "EPYC-Rome"
)

func (obj QEMUCPUKind) IsValid() bool {
	switch obj {
	case QEMUCPUKindHost,
		QEMUCPUKindMax,
		QEMUCPUKindQEMU32,
		QEMUCPUKindQEMU64,
		QEMUCPUKindKVM32,
		QEMUCPUKindKVM64,
		QEMUCPUKindIntel486,
		QEMUCPUKindIntelPentium,
		QEMUCPUKindIntelPentium2,
		QEMUCPUKindIntelPentium3,
		QEMUCPUKindIntelCoreDuo,
		QEMUCPUKindIntelCore2Duo,
		QEMUCPUKindIntelConroe,
		QEMUCPUKindIntelPenryn,
		QEMUCPUKindIntelNehalem,
		QEMUCPUKindIntelNehalemIBRS,
		QEMUCPUKindIntelWestmere,
		QEMUCPUKindIntelWestmereIBRS,
		QEMUCPUKindIntelSandyBridge,
		QEMUCPUKindIntelSandyBridgeIBRS,
		QEMUCPUKindIntelIvyBridge,
		QEMUCPUKindIntelIvyBridgeIBRS,
		QEMUCPUKindIntelHaswell,
		QEMUCPUKindIntelHaswellIBRS,
		QEMUCPUKindIntelHaswellNoTSX,
		QEMUCPUKindIntelHaswellIBRSNoTSX,
		QEMUCPUKindIntelBroadwell,
		QEMUCPUKindIntelBroadwellIBRS,
		QEMUCPUKindIntelBroadwellNoTSX,
		QEMUCPUKindIntelBroadwellIBRSNoTSX,
		QEMUCPUKindIntelSkyLakeClient,
		QEMUCPUKindIntelSkyLakeClientIBRS,
		QEMUCPUKindIntelSkyLakeClientIBRSNoTSX,
		QEMUCPUKindIntelSkyLakeServer,
		QEMUCPUKindIntelSkyLakeServerIBRS,
		QEMUCPUKindIntelSkyLakeServerIBRSNoTSX,
		QEMUCPUKindIntelCascadeLakeServer,
		QEMUCPUKindIntelCascadeLakeServerNoTSX,
		QEMUCPUKindIntelKnightsMill,
		QEMUCPUKindIntelIceLakeClient,
		QEMUCPUKindIntelIceLakeClientNoTSX,
		QEMUCPUKindIntelIceLakeServer,
		QEMUCPUKindIntelIceLakeServerNoTSX,
		QEMUCPUKindAMDAthlon,
		QEMUCPUKindAMDPhenom,
		QEMUCPUKindAMDOpteronG1,
		QEMUCPUKindAMDOpteronG2,
		QEMUCPUKindAMDOpteronG3,
		QEMUCPUKindAMDOpteronG4,
		QEMUCPUKindAMDOpteronG5,
		QEMUCPUKindAMDEPYC,
		QEMUCPUKindAMDEPYCIBPB,
		QEMUCPUKindAMDEPYCRome:
		return true
	default:
		return false
	}
}

func (obj QEMUCPUKind) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUCPUKind) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUCPUKind) Unmarshal(s string) error {
	*obj = QEMUCPUKind(s)
	return nil
}

func (obj *QEMUCPUKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type QEMUCPUFlags string

const (
	// Required to let the guest OS know if MDS is mitigated correctly
	QEMUCPUFlagsEnableMDClear  QEMUCPUFlags = "+md-clear"
	QEMUCPUFlagsDisableMDClear QEMUCPUFlags = "-md-clear"

	// Meltdown fix cost reduction on Westmere, SandyBride and IvyBridge Intel CPUs
	QEMUCPUFlagsEnablePCID  QEMUCPUFlags = "+pcid"
	QEMUCPUFlagsDisablePCID QEMUCPUFlags = "-pcid"

	// Allows improved Spectre mitigation with Intel CPUs
	QEMUCPUFlagsEnableSpecCtrl  QEMUCPUFlags = "+spec-ctrl"
	QEMUCPUFlagsDisableSpecCtrl QEMUCPUFlags = "-spec-ctrl"

	// Protection for "Speculative Store ByPass" for Intel models
	QEMUCPUFlagsEnableSSBD  QEMUCPUFlags = "+ssbd"
	QEMUCPUFlagsDisableSSBD QEMUCPUFlags = "-ssbd"

	// Allows improved Spectre mitigation with AMD CPUs
	QEMUCPUFlagsEnableIBPB  QEMUCPUFlags = "+ibpb"
	QEMUCPUFlagsDisableIBPB QEMUCPUFlags = "-ibpb"

	// Basis for "Speculative Store Bypass" protection for AMD models
	QEMUCPUFlagsEnableVirtSSBD  QEMUCPUFlags = "+virt-ssbd"
	QEMUCPUFlagsDisableVirtSSBD QEMUCPUFlags = "-virt-ssbd"

	// Improves Spectre mitigation performance with AMD CPUs, best used with "virt-ssbd"
	QEMUCPUFlagsEnableAMDSSBD  QEMUCPUFlags = "+amd-ssbd"
	QEMUCPUFlagsDisableAMDSSBD QEMUCPUFlags = "-amd-ssbd"

	// Notifies guest OS that host is not vulnerable for Spectre on AMD CPUs
	QEMUCPUFlagsEnableAMDNoSSB  QEMUCPUFlags = "+amd-no-ssb"
	QEMUCPUFlagsDisableAMDNoSSB QEMUCPUFlags = "-amd-no-ssb"

	// Allow guest OS to use 1GB size pages, if host HW supports it
	QEMUCPUFlagsEnablePDPE1GB  QEMUCPUFlags = "+pdpe1gb"
	QEMUCPUFlagsDisablePDPE1GB QEMUCPUFlags = "-pdpe1gb"

	// Improve performance in overcommited Windows guests. May lead to guest bluescreens on old CPUs.
	QEMUCPUFlagsEnableHVTLBFlush  QEMUCPUFlags = "+hv-tlbflush"
	QEMUCPUFlagsDisableHVTLBFlush QEMUCPUFlags = "-hv-tlbflush"

	// Improve performance for nested virtualization. Only supported on Intel CPUs.
	QEMUCPUFlagsEnableHVEVMCS  QEMUCPUFlags = "+hv-evmcs"
	QEMUCPUFlagsDisableHVEVMCS QEMUCPUFlags = "-hv-evmcs"

	// Activate AES instruction set for HW acceleration.
	QEMUCPUFlagsEnableAES  QEMUCPUFlags = "+aes"
	QEMUCPUFlagsDisableAES QEMUCPUFlags = "-aes"
)

func (obj QEMUCPUFlags) IsValid() bool {
	switch obj {
	case QEMUCPUFlagsEnableMDClear,
		QEMUCPUFlagsDisableMDClear,
		QEMUCPUFlagsEnablePCID,
		QEMUCPUFlagsDisablePCID,
		QEMUCPUFlagsEnableSpecCtrl,
		QEMUCPUFlagsDisableSpecCtrl,
		QEMUCPUFlagsEnableSSBD,
		QEMUCPUFlagsDisableSSBD,
		QEMUCPUFlagsEnableIBPB,
		QEMUCPUFlagsDisableIBPB,
		QEMUCPUFlagsEnableVirtSSBD,
		QEMUCPUFlagsDisableVirtSSBD,
		QEMUCPUFlagsEnableAMDSSBD,
		QEMUCPUFlagsDisableAMDSSBD,
		QEMUCPUFlagsEnableAMDNoSSB,
		QEMUCPUFlagsDisableAMDNoSSB,
		QEMUCPUFlagsEnablePDPE1GB,
		QEMUCPUFlagsDisablePDPE1GB,
		QEMUCPUFlagsEnableHVTLBFlush,
		QEMUCPUFlagsDisableHVTLBFlush,
		QEMUCPUFlagsEnableHVEVMCS,
		QEMUCPUFlagsDisableHVEVMCS,
		QEMUCPUFlagsEnableAES,
		QEMUCPUFlagsDisableAES:
		return true
	default:
		return false
	}
}

func (obj QEMUCPUFlags) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUCPUFlags) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUCPUFlags) Unmarshal(s string) error {
	*obj = QEMUCPUFlags(s)
	return nil
}

func (obj *QEMUCPUFlags) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
