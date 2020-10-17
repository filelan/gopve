package qemu

import (
	"encoding/json"
)

type CPUType string

const (
	// KVM processor with all supported host features.
	CPUTypeHost CPUType = "host"

	// Enables all features supported by the accelerator in the current host.
	CPUTypeMax CPUType = "max"

	// QEMU Virtual CPU version 2.5+ (32 bit variant)
	CPUTypeQEMU32 CPUType = "qemu32"

	// QEMU Virtual CPU version 2.5+ (64 bit variant)
	CPUTypeQEMU64 CPUType = "qemu64"

	// Common KVM processor (32 bit variant). Legacy model just for historical compatibility with ancient QEMU versions.
	CPUTypeKVM32 CPUType = "kvm32"

	// Common KVM processor (64 bit variant). Legacy model just for historical compatibility with ancient QEMU versions.
	CPUTypeKVM64 CPUType = "kvm64"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeIntel486 CPUType = "486"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeIntelPentium CPUType = "pentium"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeIntelPentium2 CPUType = "pentium2"

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeIntelPentium3 CPUType = "pentium3"

	// Intel CPU T2600 @ 2.16. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeIntelCoreDuo CPUType = "coreduo"

	// Intel Core 2 Duo CPU T7700  @ 2.40GHz. Old Intel x86 model, its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeIntelCore2Duo CPUType = "core2duo"

	// Intel Celeron_4x0 (Conroe/Merom Class Core 2, 2006)
	CPUTypeIntelConroe CPUType = "Conroe"

	// Intel Core 2 Duo P9xxx (Penryn Class Core 2, 2007)
	CPUTypeIntelPenryn CPUType = "Penryn"

	// Intel Core i7 9xx (Nehalem Class Core i7, 2008)
	CPUTypeIntelNehalem CPUType = "Nehalem"

	// Intel Core i7 9xx (Nehalem Class Core i7, 2008, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelNehalemIBRS CPUType = "Nehalem-IBRS"

	// Westmere E56xx/L56xx/X56xx (Nehalem-C, 2010)
	CPUTypeIntelWestmere CPUType = "Westmere"

	// Westmere E56xx/L56xx/X56xx (Nehalem-C, 2010, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelWestmereIBRS CPUType = "Westmere-IBRS"

	// Intel Xeon E312xx (Sandy Bridge, 2011)
	CPUTypeIntelSandyBridge CPUType = "SandyBridge"

	// Intel Xeon E312xx (Sandy Bridge, 2011, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelSandyBridgeIBRS CPUType = "SandyBridge-IBRS"

	// Intel Xeon E3-12xx v2 (Ivy Bridge, 2012)
	CPUTypeIntelIvyBridge CPUType = "IvyBridge"

	// Intel Xeon E3-12xx v2 (Ivy Bridge, 2012, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelIvyBridgeIBRS CPUType = "IvyBridge-IBRS"

	// Intel Core Processor (Haswell, 2013)
	CPUTypeIntelHaswell CPUType = "Haswell"

	// Intel Core Processor (Haswell, 2013, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelHaswellIBRS CPUType = "Haswell-IBRS"

	// Intel Core Processor (Haswell, 2013, without Transactional Synchronization Extensions update)
	CPUTypeIntelHaswellNoTSX CPUType = "Haswell-noTSX"

	// Intel Core Processor (Haswell, 2013, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	CPUTypeIntelHaswellIBRSNoTSX CPUType = "Haswell-noTSX-IBRS"

	// Intel Core Processor (Broadwell, 2014)
	CPUTypeIntelBroadwell CPUType = "Broadwell"

	// Intel Core Processor (Broadwell, 2014, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelBroadwellIBRS CPUType = "Broadwell-IBRS"

	// Intel Core Processor (Broadwell, 2014, without Transactional Synchronization Extensions update)
	CPUTypeIntelBroadwellNoTSX CPUType = "Broadwell-noTSX"

	// Intel Core Processor (Broadwell, 2014, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	CPUTypeIntelBroadwellIBRSNoTSX CPUType = "Broadwell-noTSX-IBRS"

	// Intel Core Processor (Skylake, 2015)
	CPUTypeIntelSkyLakeClient CPUType = "Skylake-Client"

	// Intel Core Processor (Skylake, 2015, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelSkyLakeClientIBRS CPUType = "Skylake-Client-IBRS"

	// Intel Core Processor (Skylake, 2015, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	CPUTypeIntelSkyLakeClientIBRSNoTSX CPUType = "Skylake-Client-noTSX-IBRS"

	// Intel Xeon Processor (Skylake, 2016)
	CPUTypeIntelSkyLakeServer CPUType = "Skylake-Server"

	// Intel Xeon Processor (Skylake, 2016, with Indirect Branch Restricted Speculation update)
	CPUTypeIntelSkyLakeServerIBRS CPUType = "Skylake-Server-IBRS"

	// Intel Xeon Processor (Skylake, 2016, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	CPUTypeIntelSkyLakeServerIBRSNoTSX CPUType = "Skylake-Server-noTSX-IBRS"

	// Intel Xeon Processor (Cascade Lake, 2019, with “stepping” levels 6 or 7 only). The Cascade Lake Xeon processor with stepping 5 is vulnerable to MDS variants.
	CPUTypeIntelCascadeLakeServer CPUType = "Cascadelake-Server"

	// Intel Xeon Processor (Cascade Lake, 2019, with “stepping” levels 6 or 7 only, without Transactional Synchronization Extensions update). The Cascade Lake Xeon processor with stepping 5 is vulnerable to MDS variants.
	CPUTypeIntelCascadeLakeServerNoTSX CPUType = "Cascadelake-Server-noTSX"

	// Intel Xeon Phi Processor (Knights Mill)
	CPUTypeIntelKnightsMill CPUType = "KnightsMill"

	// Intel Core Processor (Icelake)
	CPUTypeIntelIceLakeClient CPUType = "Icelake-Client"

	// Intel Core Processor (Icelake, without Transactional Synchronization Extensions update)
	CPUTypeIntelIceLakeClientNoTSX CPUType = "Icelake-Client-noTSX"

	// Intel Xeon Processor (Icelake)
	CPUTypeIntelIceLakeServer CPUType = "Icelake-Server"

	// Intel Xeon Processor (Icelake, without Transactional Synchronization Extensions update)
	CPUTypeIntelIceLakeServerNoTSX CPUType = "Icelake-Server-noTSX"

	// Old AMD x86 model, its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeAMDAthlon CPUType = "athlon"

	// AMD Phenom(tm) 9550 Quad-Core Processor. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	CPUTypeAMDPhenom CPUType = "phenom"

	// AMD Opteron 240 (Gen 1 Class Opteron, 2004)
	CPUTypeAMDOpteronG1 CPUType = "Opteron_G1"

	// AMD Opteron 22xx (Gen 2 Class Opteron, 2006)
	CPUTypeAMDOpteronG2 CPUType = "Opteron_G2"

	// AMD Opteron 23xx (Gen 3 Class Opteron, 2009)
	CPUTypeAMDOpteronG3 CPUType = "Opteron_G3"

	// AMD Opteron 62xx class CPU (2011)
	CPUTypeAMDOpteronG4 CPUType = "Opteron_G4"

	// AMD Opteron 63xx class CPU (2012)
	CPUTypeAMDOpteronG5 CPUType = "Opteron_G5"

	// AMD EPYC Processor (2017)
	CPUTypeAMDEPYC CPUType = "EPYC"

	// AMD EPYC Processor (2017, with Indirect Branch Prediction Barrier update)
	CPUTypeAMDEPYCIBPB CPUType = "EPYC-IBPB"

	// AMD EPYC-Rome Processor (2019)
	CPUTypeAMDEPYCRome CPUType = "EPYC-Rome"
)

func (obj CPUType) IsValid() bool {
	switch obj {
	case CPUTypeHost,
		CPUTypeMax,
		CPUTypeQEMU32,
		CPUTypeQEMU64,
		CPUTypeKVM32,
		CPUTypeKVM64,
		CPUTypeIntel486,
		CPUTypeIntelPentium,
		CPUTypeIntelPentium2,
		CPUTypeIntelPentium3,
		CPUTypeIntelCoreDuo,
		CPUTypeIntelCore2Duo,
		CPUTypeIntelConroe,
		CPUTypeIntelPenryn,
		CPUTypeIntelNehalem,
		CPUTypeIntelNehalemIBRS,
		CPUTypeIntelWestmere,
		CPUTypeIntelWestmereIBRS,
		CPUTypeIntelSandyBridge,
		CPUTypeIntelSandyBridgeIBRS,
		CPUTypeIntelIvyBridge,
		CPUTypeIntelIvyBridgeIBRS,
		CPUTypeIntelHaswell,
		CPUTypeIntelHaswellIBRS,
		CPUTypeIntelHaswellNoTSX,
		CPUTypeIntelHaswellIBRSNoTSX,
		CPUTypeIntelBroadwell,
		CPUTypeIntelBroadwellIBRS,
		CPUTypeIntelBroadwellNoTSX,
		CPUTypeIntelBroadwellIBRSNoTSX,
		CPUTypeIntelSkyLakeClient,
		CPUTypeIntelSkyLakeClientIBRS,
		CPUTypeIntelSkyLakeClientIBRSNoTSX,
		CPUTypeIntelSkyLakeServer,
		CPUTypeIntelSkyLakeServerIBRS,
		CPUTypeIntelSkyLakeServerIBRSNoTSX,
		CPUTypeIntelCascadeLakeServer,
		CPUTypeIntelCascadeLakeServerNoTSX,
		CPUTypeIntelKnightsMill,
		CPUTypeIntelIceLakeClient,
		CPUTypeIntelIceLakeClientNoTSX,
		CPUTypeIntelIceLakeServer,
		CPUTypeIntelIceLakeServerNoTSX,
		CPUTypeAMDAthlon,
		CPUTypeAMDPhenom,
		CPUTypeAMDOpteronG1,
		CPUTypeAMDOpteronG2,
		CPUTypeAMDOpteronG3,
		CPUTypeAMDOpteronG4,
		CPUTypeAMDOpteronG5,
		CPUTypeAMDEPYC,
		CPUTypeAMDEPYCIBPB,
		CPUTypeAMDEPYCRome:
		return true
	default:
		return false
	}
}

func (obj CPUType) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj CPUType) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *CPUType) Unmarshal(s string) error {
	*obj = CPUType(s)
	return nil
}

func (obj *CPUType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
