package config

// CertificateIssuanceSettings 是新增跨包桥接，表达原证书包对 0x509f9e0..0x509fa18
// 四个全局字符串的内联读取；不是原样本中独立存在的函数符号。
func CertificateIssuanceSettings() (mode, domain, keyLength, dnsProvider string) {
	c := &configAt509F9C0
	return c.MVmP0dSlHl, c.Ga_YcC, c.AQ1zOwu1K, c.Xs7eJ3aj6Kbl
}

// SetCertificateFiles 是原 TFwUsh 对 0x509f9c0/0x509f9d0 写入的跨包桥接。
func SetCertificateFiles(cert, key string) {
	configAt509F9C0.C2BNS6_QJC = cert
	configAt509F9C0.JISdid54S5t = key
}

// CertificateDNSValue 桥接原 cKh3OZ4rAaWw 对字符串映射 0x5095aa8 的内联读取。
// 原调用者先 strings.ToLower，读取本身不执行标准化或环境查询。
func CertificateDNSValue(key string) string { return configurationAt5095AA8[key] }
