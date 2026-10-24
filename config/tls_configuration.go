package config

import (
	"crypto/tls"
	L "recovered.local/soga/util/logx"
	"strconv"
	"strings"
)

// B02ZlILA8i 对应0x132ea80，仅认可原tls.CipherSuites安全套件集合。
// 名称区分大小写，十六进制仅接受小写0x前缀，保留顺序及重复值。
func B02ZlILA8i() []uint16 {
	byName := make(map[string]uint16)
	byID := make(map[uint16]bool)
	for _, s := range tls.CipherSuites() {
		byName[s.Name] = s.ID
		byID[s.ID] = true
	}
	var result []uint16
	for _, token := range strings.Split(configAt509EBE0.KVniDErTtx, ":") {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		if id, ok := byName[token]; ok {
			result = append(result, id)
			continue
		}
		if strings.HasPrefix(token, "0x") {
			token = token[2:]
			id, err := strconv.ParseUint(token, 16, 16)
			if err == nil && byID[uint16(id)] {
				result = append(result, uint16(id))
				continue
			}
		}
		L.TRTbrquD(4, "not support tls cipher suite:", token)
	}
	return result
}

// c2dSGQJ4q 对应0x132f500。证书与ALPN均保留原切片引用关系。
func c2dSGQJ4q(cert tls.Certificate) *tls.Config {
	version := func(s string) uint16 {
		switch s {
		case "1.0":
			return tls.VersionTLS10
		case "1.1":
			return tls.VersionTLS11
		case "1.2":
			return tls.VersionTLS12
		case "1.3":
			return tls.VersionTLS13
		}
		return 0
	}
	min, max := version(configAt509EBE0.CJri5kB26), version(configAt509EBE0.NH7oB3o7)
	alpn := configAt509EBE0.RS1kBzpBezso
	suites := B02ZlILA8i()
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}, CipherSuites: suites, PreferServerCipherSuites: configAt509EBE0.JOyIs3s9WAcA}
	if min != 0 {
		cfg.MinVersion = min
	}
	if max != 0 {
		cfg.MaxVersion = max
	}
	if len(alpn) != 0 {
		cfg.NextProtos = alpn
	}
	return cfg
}

// Ks6Wk5LnO 对应0x132f1a0，载入实际配置指定的证书文件。
func Ks6Wk5LnO() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(configAt509F9C0.C2BNS6_QJC, configAt509F9C0.JISdid54S5t)
	if err != nil {
		return nil, err
	}
	return c2dSGQJ4q(cert), nil
}

// XmxN0NDuhcA 对应0x132f300；仅当两份PEM均非空时使用传入证书。
func XmxN0NDuhcA(certPEM, keyPEM string) (*tls.Config, error) {
	if certPEM == "" || keyPEM == "" {
		return Ks6Wk5LnO()
	}
	cert, err := tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		return nil, err
	}
	return c2dSGQJ4q(cert), nil
}
