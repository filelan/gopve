package vm

import (
	"encoding/json"
	"fmt"
)

type QEMUCPUKind uint

const (
	// KVM processor with all supported host features.
	QEMUCPUKindHost QEMUCPUKind = iota

	// Enables all features supported by the accelerator in the current host.
	QEMUCPUKindMax

	// Common KVM processor (32 bit variant). Legacy model just for historical compatibility with ancient QEMU versions.
	QEMUCPUKindKVM32

	// Common KVM processor (64 bit variant). Legacy model just for historical compatibility with ancient QEMU versions.
	QEMUCPUKindKVM64

	// QEMU Virtual CPU version 2.5+ (32 bit variant)
	QEMUCPUKindQEMU32

	// QEMU Virtual CPU version 2.5+ (64 bit variant)
	QEMUCPUKindQEMU64

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntel486

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelPentium

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelPentium2

	// Old Intel x86 model. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelPentium3

	// Intel CPU T2600 @ 2.16. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelCoreDuo

	// Intel Core 2 Duo CPU T7700  @ 2.40GHz. Old Intel x86 model, its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindIntelCore2Duo

	// Intel Celeron_4x0 (Conroe/Merom Class Core 2, 2006)
	QEMUCPUKindIntelConroe

	// Intel Core 2 Duo P9xxx (Penryn Class Core 2, 2007)
	QEMUCPUKindIntelPenryn

	// Intel Core i7 9xx (Nehalem Class Core i7, 2008)
	QEMUCPUKindIntelNehalem

	// Intel Core i7 9xx (Nehalem Class Core i7, 2008, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelNehalemIBRS

	// Westmere E56xx/L56xx/X56xx (Nehalem-C, 2010)
	QEMUCPUKindIntelWestmere

	// Westmere E56xx/L56xx/X56xx (Nehalem-C, 2010, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelWestmereIBRS

	// Intel Xeon E312xx (Sandy Bridge, 2011)
	QEMUCPUKindIntelSandyBridge

	// Intel Xeon E312xx (Sandy Bridge, 2011, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelSandyBridgeIBRS

	// Intel Xeon E3-12xx v2 (Ivy Bridge, 2012)
	QEMUCPUKindIntelIvyBridge

	// Intel Xeon E3-12xx v2 (Ivy Bridge, 2012, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelIvyBridgeIBRS

	// Intel Core Processor (Haswell, 2013)
	QEMUCPUKindIntelHaswell

	// Intel Core Processor (Haswell, 2013, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelHaswellIBRS

	// Intel Core Processor (Haswell, 2013, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelHaswellNoTSX

	// Intel Core Processor (Haswell, 2013, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelHaswellIBRSNoTSX

	// Intel Core Processor (Broadwell, 2014)
	QEMUCPUKindIntelBroadwell

	// Intel Core Processor (Broadwell, 2014, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelBroadwellIBRS

	// Intel Core Processor (Broadwell, 2014, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelBroadwellNoTSX

	// Intel Core Processor (Broadwell, 2014, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelBroadwellIBRSNoTSX

	// Intel Core Processor (Skylake, 2015)
	QEMUCPUKindIntelSkyLakeClient

	// Intel Core Processor (Skylake, 2015, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelSkyLakeClientIBRS

	// Intel Core Processor (Skylake, 2015, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelSkyLakeClientIBRSNoTSX

	// Intel Xeon Processor (Skylake, 2016)
	QEMUCPUKindIntelSkyLakeServer

	// Intel Xeon Processor (Skylake, 2016, with Indirect Branch Restricted Speculation update)
	QEMUCPUKindIntelSkyLakeServerIBRS

	// Intel Xeon Processor (Skylake, 2016, with Indirect Branch Restricted Speculation update and without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelSkyLakeServerIBRSNoTSX

	// Intel Xeon Processor (Cascade Lake, 2019, with “stepping” levels 6 or 7 only). The Cascade Lake Xeon processor with stepping 5 is vulnerable to MDS variants.
	QEMUCPUKindIntelCascadeLakeServer

	// Intel Xeon Processor (Cascade Lake, 2019, with “stepping” levels 6 or 7 only, without Transactional Synchronization Extensions update). The Cascade Lake Xeon processor with stepping 5 is vulnerable to MDS variants.
	QEMUCPUKindIntelCascadeLakeServerNoTSX

	// Intel Xeon Phi Processor (Knights Mill)
	QEMUCPUKindIntelKnightsMill

	// Intel Core Processor (Icelake)
	QEMUCPUKindIntelIceLakeClient

	// Intel Core Processor (Icelake, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelIceLakeClientNoTSX

	// Intel Xeon Processor (Icelake)
	QEMUCPUKindIntelIceLakeServer

	// Intel Xeon Processor (Icelake, without Transactional Synchronization Extensions update)
	QEMUCPUKindIntelIceLakeServerNoTSX

	// Old AMD x86 model, its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindAMDAthlon

	// AMD Phenom(tm) 9550 Quad-Core Processor. Its usage is discouraged, as it exposes a very limited featureset, which prevents guests having optimal performance.
	QEMUCPUKindAMDPhenom

	// AMD Opteron 240 (Gen 1 Class Opteron, 2004)
	QEMUCPUKindAMDOpteronG1

	// AMD Opteron 22xx (Gen 2 Class Opteron, 2006)
	QEMUCPUKindAMDOpteronG2

	// AMD Opteron 23xx (Gen 3 Class Opteron, 2009)
	QEMUCPUKindAMDOpteronG3

	// AMD Opteron 62xx class CPU (2011)
	QEMUCPUKindAMDOpteronG4

	// AMD Opteron 63xx class CPU (2012)
	QEMUCPUKindAMDOpteronG5

	// AMD EPYC Processor (2017)
	QEMUCPUKindEPYC

	// AMD EPYC Processor (2017, with Indirect Branch Prediction Barrier update)
	QEMUCPUKindEPYCIBPB

	// AMD EPYC-Rome Processor (2019)
	QEMUCPUKindEPYCRome
)

func (obj QEMUCPUKind) Marshal() (string, error) {
	switch obj {
	case QEMUCPUKindHost:
		return "host", nil
	case QEMUCPUKindMax:
		return "max", nil
	case QEMUCPUKindQEMU32:
		return "qemu32", nil
	case QEMUCPUKindQEMU64:
		return "qemu64", nil
	case QEMUCPUKindKVM32:
		return "kvm32", nil
	case QEMUCPUKindKVM64:
		return "kvm64", nil
	case QEMUCPUKindIntel486:
		return "486", nil
	case QEMUCPUKindIntelPentium:
		return "pentium", nil
	case QEMUCPUKindIntelPentium2:
		return "pentium2", nil
	case QEMUCPUKindIntelPentium3:
		return "pentium3", nil
	case QEMUCPUKindIntelCoreDuo:
		return "coreduo", nil
	case QEMUCPUKindIntelCore2Duo:
		return "core2duo", nil
	case QEMUCPUKindIntelConroe:
		return "Conroe", nil
	case QEMUCPUKindIntelPenryn:
		return "Penryn", nil
	case QEMUCPUKindIntelNehalem:
		return "Nehalem", nil
	case QEMUCPUKindIntelNehalemIBRS:
		return "Nehalem-IBRS", nil
	case QEMUCPUKindIntelWestmere:
		return "Westmere", nil
	case QEMUCPUKindIntelWestmereIBRS:
		return "Westmere-IBRS", nil
	case QEMUCPUKindIntelSandyBridge:
		return "SandyBridge", nil
	case QEMUCPUKindIntelSandyBridgeIBRS:
		return "SandyBridge-IBRS", nil
	case QEMUCPUKindIntelIvyBridge:
		return "IvyBridge", nil
	case QEMUCPUKindIntelIvyBridgeIBRS:
		return "IvyBridge-IBRS", nil
	case QEMUCPUKindIntelHaswell:
		return "Haswell", nil
	case QEMUCPUKindIntelHaswellIBRS:
		return "Haswell-IBRS", nil
	case QEMUCPUKindIntelHaswellNoTSX:
		return "Haswell-noTSX", nil
	case QEMUCPUKindIntelHaswellIBRSNoTSX:
		return "Haswell-noTSX-IBRS", nil
	case QEMUCPUKindIntelBroadwell:
		return "Broadwell", nil
	case QEMUCPUKindIntelBroadwellIBRS:
		return "Broadwell-IBRS", nil
	case QEMUCPUKindIntelBroadwellNoTSX:
		return "Broadwell-noTSX", nil
	case QEMUCPUKindIntelBroadwellIBRSNoTSX:
		return "Broadwell-noTSX-IBRS", nil
	case QEMUCPUKindIntelSkyLakeClient:
		return "Skylake-Client", nil
	case QEMUCPUKindIntelSkyLakeClientIBRS:
		return "Skylake-Client-IBRS", nil
	case QEMUCPUKindIntelSkyLakeClientIBRSNoTSX:
		return "Skylake-Client-noTSX-IBRS", nil
	case QEMUCPUKindIntelSkyLakeServer:
		return "Skylake-Server", nil
	case QEMUCPUKindIntelSkyLakeServerIBRS:
		return "Skylake-Server-IBRS", nil
	case QEMUCPUKindIntelSkyLakeServerIBRSNoTSX:
		return "Skylake-Server-noTSX-IBRS", nil
	case QEMUCPUKindIntelCascadeLakeServer:
		return "Cascadelake-Server", nil
	case QEMUCPUKindIntelCascadeLakeServerNoTSX:
		return "Cascadelake-Server-noTSX", nil
	case QEMUCPUKindIntelKnightsMill:
		return "KnightsMill", nil
	case QEMUCPUKindIntelIceLakeClient:
		return "Icelake-Client", nil
	case QEMUCPUKindIntelIceLakeClientNoTSX:
		return "Icelake-Client-noTSX", nil
	case QEMUCPUKindIntelIceLakeServer:
		return "Icelake-Server", nil
	case QEMUCPUKindIntelIceLakeServerNoTSX:
		return "Icelake-Server-noTSX", nil
	case QEMUCPUKindAMDAthlon:
		return "athlon", nil
	case QEMUCPUKindAMDPhenom:
		return "phenom", nil
	case QEMUCPUKindAMDOpteronG1:
		return "Opteron_G1", nil
	case QEMUCPUKindAMDOpteronG2:
		return "Opteron_G2", nil
	case QEMUCPUKindAMDOpteronG3:
		return "Opteron_G3", nil
	case QEMUCPUKindAMDOpteronG4:
		return "Opteron_G4", nil
	case QEMUCPUKindAMDOpteronG5:
		return "Opteron_G5", nil
	case QEMUCPUKindEPYC:
		return "EPYC", nil
	case QEMUCPUKindEPYCIBPB:
		return "EPYC-IBPB", nil
	case QEMUCPUKindEPYCRome:
		return "EPYC-Rome", nil
	default:
		return "", fmt.Errorf("unknown qemu cpu type")
	}
}

func (obj *QEMUCPUKind) Unmarshal(s string) error {
	switch s {
	case "host":
		*obj = QEMUCPUKindHost
	case "max":
		*obj = QEMUCPUKindMax
	case "qemu32":
		*obj = QEMUCPUKindQEMU32
	case "qemu64":
		*obj = QEMUCPUKindQEMU64
	case "kvm32":
		*obj = QEMUCPUKindKVM32
	case "kvm64":
		*obj = QEMUCPUKindKVM64
	case "486":
		*obj = QEMUCPUKindIntel486
	case "pentium":
		*obj = QEMUCPUKindIntelPentium
	case "pentium2":
		*obj = QEMUCPUKindIntelPentium2
	case "pentium3":
		*obj = QEMUCPUKindIntelPentium3
	case "coreduo":
		*obj = QEMUCPUKindIntelCoreDuo
	case "core2duo":
		*obj = QEMUCPUKindIntelCore2Duo
	case "Conroe":
		*obj = QEMUCPUKindIntelConroe
	case "Penryn":
		*obj = QEMUCPUKindIntelPenryn
	case "Nehalem":
		*obj = QEMUCPUKindIntelNehalem
	case "Nehalem-IBRS":
		*obj = QEMUCPUKindIntelNehalemIBRS
	case "Westmere":
		*obj = QEMUCPUKindIntelWestmere
	case "Westmere-IBRS":
		*obj = QEMUCPUKindIntelWestmereIBRS
	case "SandyBridge":
		*obj = QEMUCPUKindIntelSandyBridge
	case "SandyBridge-IBRS":
		*obj = QEMUCPUKindIntelSandyBridgeIBRS
	case "IvyBridge":
		*obj = QEMUCPUKindIntelIvyBridge
	case "IvyBridge-IBRS":
		*obj = QEMUCPUKindIntelIvyBridgeIBRS
	case "Haswell":
		*obj = QEMUCPUKindIntelHaswell
	case "Haswell-IBRS":
		*obj = QEMUCPUKindIntelHaswellIBRS
	case "Haswell-noTSX":
		*obj = QEMUCPUKindIntelHaswellNoTSX
	case "Haswell-noTSX-IBRS":
		*obj = QEMUCPUKindIntelHaswellIBRSNoTSX
	case "Broadwell":
		*obj = QEMUCPUKindIntelBroadwell
	case "Broadwell-IBRS":
		*obj = QEMUCPUKindIntelBroadwellIBRS
	case "Broadwell-noTSX":
		*obj = QEMUCPUKindIntelBroadwellNoTSX
	case "Broadwell-noTSX-IBRS":
		*obj = QEMUCPUKindIntelBroadwellIBRSNoTSX
	case "Skylake-Client":
		*obj = QEMUCPUKindIntelSkyLakeClient
	case "Skylake-Client-IBRS":
		*obj = QEMUCPUKindIntelSkyLakeClientIBRS
	case "Skylake-Client-noTSX-IBRS":
		*obj = QEMUCPUKindIntelSkyLakeClientIBRSNoTSX
	case "Skylake-Server":
		*obj = QEMUCPUKindIntelSkyLakeServer
	case "Skylake-Server-IBRS":
		*obj = QEMUCPUKindIntelSkyLakeServerIBRS
	case "Skylake-Server-noTSX-IBRS":
		*obj = QEMUCPUKindIntelSkyLakeServerIBRSNoTSX
	case "Cascadelake-Server":
		*obj = QEMUCPUKindIntelCascadeLakeServer
	case "Cascadelake-Server-noTSX":
		*obj = QEMUCPUKindIntelCascadeLakeServerNoTSX
	case "KnightsMill":
		*obj = QEMUCPUKindIntelKnightsMill
	case "Icelake-Client":
		*obj = QEMUCPUKindIntelIceLakeClient
	case "Icelake-Client-noTSX":
		*obj = QEMUCPUKindIntelIceLakeClientNoTSX
	case "Icelake-Server":
		*obj = QEMUCPUKindIntelIceLakeServer
	case "Icelake-Server-noTSX":
		*obj = QEMUCPUKindIntelIceLakeServerNoTSX
	case "athlon":
		*obj = QEMUCPUKindAMDAthlon
	case "phenom":
		*obj = QEMUCPUKindAMDPhenom
	case "Opteron_G1":
		*obj = QEMUCPUKindAMDOpteronG1
	case "Opteron_G2":
		*obj = QEMUCPUKindAMDOpteronG2
	case "Opteron_G3":
		*obj = QEMUCPUKindAMDOpteronG3
	case "Opteron_G4":
		*obj = QEMUCPUKindAMDOpteronG4
	case "Opteron_G5":
		*obj = QEMUCPUKindAMDOpteronG5
	case "EPYC":
		*obj = QEMUCPUKindEPYC
	case "EPYC-IBPB":
		*obj = QEMUCPUKindEPYCIBPB
	case "EPYC-Rome":
		*obj = QEMUCPUKindEPYCRome
	default:
		return fmt.Errorf("can't unmarshal qemu cpu type %s", s)
	}

	return nil
}

func (obj *QEMUCPUKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
