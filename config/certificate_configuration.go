package config

// ManualCertificateFiles reads the two original strings at 0x509f9c0/0x509f9d0.
func ManualCertificateFiles() (cert, key string) {
	return configAt509F9C0.C2BNS6_QJC, configAt509F9C0.JISdid54S5t
}

// TunnelSettings reads 0x509e300 and the string at 0x509e318.
func TunnelSettings() (enabled bool, mode string) {
	return configAt509E300.JHXrVGNjg31, configAt509E300.GbyJkc
}
