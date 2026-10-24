package cNT24rG4F

// InspectionConfiguration reads 0x50a2959, 0x50a2960 and 0x50a29a1.
func InspectionConfiguration() (detectPacket bool, maxLength int, banPrivateIP bool) {
	return configAt50A2900.XKO2Zot_hEk, configAt50A2900.GozCogFPm, configAt50A2900.CVshaQaB33U
}

// IsInspectionExempt uses the same set populated by init.0.func1@0x1346e80.
// Original Sniff/TestAddr/TestPacket reference this set through 0x5029168.
func IsInspectionExempt(id int64) bool {
	_, ok := blockListWhiteIDsAt5095AE0[id]
	return ok
}

// DomainSniffProtocols reads the slice header at 0x50a2940 without a copy.
func DomainSniffProtocols() []string { return configAt50A2900.QgoKiaNp9 }

// ForceCloseSSL reads the original bool at 0x509ec30.
func ForceCloseSSL() bool { return configAt509EBE0.CLr6mtlmQ }

func SniffRedirect() bool { return configAt50A2900.Al5sHO4x90sY }

func AccessLogEnabled() bool { return configAt50A2260.PmIFLM }
