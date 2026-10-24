// 本文件是对所提供 soga-v2.17.0 二进制 RTTI 的 Go 类型源码表示，
// 不是作者原始文件；未加入业务替代实现。
// 样本 SHA256：6660dcfbc7971188dcb751b440f4269aa718491b569d4c8629b751e6e339f92a。
// 证据：outputs/soga-analysis/types/configuration_structs.json、all_types.json。
// 包名采用方法与 uncommon_package_path 中的 cNT24rG4F，便于恢复未导出类型的业务函数。
// 类型名字中的 ftLBcUpmkQ 前缀与之不一致；此目录不声称是作者原工程目录。
package config

import "recovered.local/soga/crypto/ssmethodtype"

// F6DuWgn9N 是 RTTI 保留的命名字符串类型，不能替换成类型别名。
// 类型描述符虚拟地址：0x3582c40。
type F6DuWgn9N string

// cEgBVJvq：类型描述符 0x35e8298；amd64 大小 16 字节，对齐 8 字节。
type cEgBVJvq struct {
	M4qDq3KsUi string `conf:"routes_url"` // 字节偏移 0；字段记录文件偏移 0x31e82f8。
}

// o9NA6hjqxImy：类型描述符 0x35e8310；amd64 大小 16 字节，对齐 8 字节。
type o9NA6hjqxImy struct {
	RM5SFRF string `conf:"soga_key"` // 字节偏移 0；字段记录文件偏移 0x31e8370。
}

// piJb6eaS：类型描述符 0x35e8388；amd64 大小 1 字节，对齐 1 字节。
type piJb6eaS struct {
	XpffP4ai bool `conf:"mieru_force_user_hint" default:"false"` // 字节偏移 0；字段记录文件偏移 0x31e83e8。
}

// pkfwYwZ1aHu：类型描述符 0x35e8400；amd64 大小 16 字节，对齐 8 字节。
type pkfwYwZ1aHu struct {
	M4qDq3KsUi string `conf:"white_list_url"` // 字节偏移 0；字段记录文件偏移 0x31e8460。
}

// avoYN3oP9N4：类型描述符 0x3625010；amd64 大小 32 字节，对齐 8 字节。
type avoYN3oP9N4 struct {
	ZkkjP9      string `conf:"soga_panel_url"` // 字节偏移 0；字段记录文件偏移 0x3225070。
	TzpbaXGDeaz string `conf:"soga_panel_key"` // 字节偏移 16；字段记录文件偏移 0x3225088。
}

// bKCIPg：类型描述符 0x36250a0；amd64 大小 24 字节，对齐 8 字节。
type bKCIPg struct {
	BYiTyYPg3   string `conf:"trojan_remote_addr"`             // 字节偏移 0；字段记录文件偏移 0x3225100。
	EglOu0OdYex int    `conf:"trojan_remote_port" default:"0"` // 字节偏移 16；字段记录文件偏移 0x3225118。
}

// eFyNTD2_6X：类型描述符 0x3625130；amd64 大小 40 字节，对齐 8 字节。
type eFyNTD2_6X struct {
	BooNOdXcSa3U string   `conf:"block_list_url"`           // 字节偏移 0；字段记录文件偏移 0x3225190。
	U2hwBFC      []string `conf:"block_list_white_user_id"` // 字节偏移 16；字段记录文件偏移 0x32251a8。
}

// paRsja6：类型描述符 0x36251c0；amd64 大小 24 字节，对齐 8 字节。
type paRsja6 struct {
	DlaO1l       string                `conf:"ss_obfs_mode" default:"relaxed"` // 字节偏移 0；字段记录文件偏移 0x3225220。
	zdRAxv6CNPMA ssmethodtype.VE7a8rvh // 字节偏移 16；字段记录文件偏移 0x3225238。
}

// siZtoGWL：类型描述符 0x3625250；amd64 大小 24 字节，对齐 8 字节。
type siZtoGWL struct {
	BVi1B4f_v string `conf:"anytls_fallback_addr"`             // 字节偏移 0；字段记录文件偏移 0x32252b0。
	JGw3yW    int    `conf:"anytls_fallback_port" default:"0"` // 字节偏移 16；字段记录文件偏移 0x32252c8。
}

// a4aAMG：类型描述符 0x364f108；amd64 大小 32 字节，对齐 8 字节。
type a4aAMG struct {
	PVSao2JIq bool   `conf:"prometheus_enable" default:"false"` // 字节偏移 0；字段记录文件偏移 0x324f168。
	OAHNXtGh  string `conf:"prometheus_addr"`                   // 字节偏移 8；字段记录文件偏移 0x324f180。
	DXN_loYa  int    `conf:"prometheus_port" default:"19090"`   // 字节偏移 24；字段记录文件偏移 0x324f198。
}

// cp9sUcOF：类型描述符 0x364f1b0；amd64 大小 32 字节，对齐 8 字节。
type cp9sUcOF struct {
	TsRVDHP0SA   int    `conf:"ip_user_cache_time" default:"24"`             // 字节偏移 0；字段记录文件偏移 0x324f210。
	AROEPWNZd8SF bool   `conf:"ip_user_cache_save_enable" default:"true"`    // 字节偏移 8；字段记录文件偏移 0x324f228。
	FpIPwenv     string `conf:"ip_user_cache_save_dir" default:"/etc/soga/"` // 字节偏移 16；字段记录文件偏移 0x324f240。
}

// qNDGFoxnh：类型描述符 0x364f258；amd64 大小 32 字节，对齐 8 字节。
type qNDGFoxnh struct {
	PVSao2JIq bool   `conf:"pprof_enable" default:"false"` // 字节偏移 0；字段记录文件偏移 0x324f2b8。
	OAHNXtGh  string `conf:"pprof_addr"`                   // 字节偏移 8；字段记录文件偏移 0x324f2d0。
	DXN_loYa  int    `conf:"pprof_port" default:"21010"`   // 字节偏移 24；字段记录文件偏移 0x324f2e8。
}

// zu_ScZG5P4：类型描述符 0x364f300；amd64 大小 40 字节，对齐 8 字节。
type zu_ScZG5P4 struct {
	NaavbD4r35  string `conf:"log_level" default:"info"`            // 字节偏移 0；字段记录文件偏移 0x324f360。
	HSBWD7      string `conf:"log_file_dir" default:"/etc/soga/"`   // 字节偏移 16；字段记录文件偏移 0x324f378。
	G590aNjEOVK int    `conf:"log_file_retention_days" default:"7"` // 字节偏移 32；字段记录文件偏移 0x324f390。
}

// juBqfxCScme：类型描述符 0x36af600；amd64 大小 40 字节，对齐 8 字节。
type juBqfxCScme struct {
	DPqi4uoIl    string `conf:"mu_suffix" default:"microsoft.com"` // 字节偏移 0；字段记录文件偏移 0x32af660。
	NXP6HvyFvmb9 string `conf:"mu_regex" default:"%5m%id.%suffix"` // 字节偏移 16；字段记录文件偏移 0x32af678。
	XHTMQfO      bool   `conf:"ss_obfs_udp" default:"false"`       // 字节偏移 32；字段记录文件偏移 0x32af690。
	NLpXVVM0x    bool   `conf:"ssr_cid_limit" default:"true"`      // 字节偏移 33；字段记录文件偏移 0x32af6a8。
}

// uhg6MpMKJ：类型描述符 0x36af6c0；amd64 大小 64 字节，对齐 8 字节。
type uhg6MpMKJ struct {
	RHZA1kqpMSz []string `conf:"default_dns"`                       // 字节偏移 0；字段记录文件偏移 0x32af720。
	PBvoBS2NX5  int      `conf:"dns_cache_time" default:"10"`       // 字节偏移 24；字段记录文件偏移 0x32af738。
	EAOjNaQ6    string   `conf:"dns_strategy" default:"ipv4_first"` // 字节偏移 32；字段记录文件偏移 0x32af750。
	MFHBO42     string   `conf:"dns_config_url"`                    // 字节偏移 48；字段记录文件偏移 0x32af768。
}

// aAGsqO：类型描述符 0x36c6728；amd64 大小 48 字节，对齐 8 字节。
type aAGsqO struct {
	VerfnYj     string `conf:"listen"`                    // 字节偏移 0；字段记录文件偏移 0x32c6788。
	Fg5sk_      bool   `conf:"tfo" default:"false"`       // 字节偏移 16；字段记录文件偏移 0x32c67a0。
	Hni9vNgOp4g int    `conf:"tcp_timeout" default:"120"` // 字节偏移 24；字段记录文件偏移 0x32c67b8。
	JtFIOYSoCNi int    `conf:"udp_timeout" default:"2"`   // 字节偏移 32；字段记录文件偏移 0x32c67d0。
	Arfg8UjxqP  bool   `conf:"mptcp" default:"false"`     // 字节偏移 40；字段记录文件偏移 0x32c67e8。
}

// cdqXEMYJO：类型描述符 0x36da0b8；amd64 大小 96 字节，对齐 8 字节。
type cdqXEMYJO struct {
	C2BNS6_QJC   string `conf:"cert_file"`                // 字节偏移 0；字段记录文件偏移 0x32da118。
	JISdid54S5t  string `conf:"key_file"`                 // 字节偏移 16；字段记录文件偏移 0x32da130。
	MVmP0dSlHl   string `conf:"cert_mode" default:"http"` // 字节偏移 32；字段记录文件偏移 0x32da148。
	Ga_YcC       string `conf:"cert_domain"`              // 字节偏移 48；字段记录文件偏移 0x32da160。
	AQ1zOwu1K    string `conf:"cert_key_length"`          // 字节偏移 64；字段记录文件偏移 0x32da178。
	Xs7eJ3aj6Kbl string `conf:"dns_provider"`             // 字节偏移 80；字段记录文件偏移 0x32da190。
}

// jkGZ0JHJw：类型描述符 0x36da1a8；amd64 大小 88 字节，对齐 8 字节。
type jkGZ0JHJw struct {
	RS1kBzpBezso []string `conf:"tls_alpn" default:"h2,http/1.1"`                 // 字节偏移 0；字段记录文件偏移 0x32da208。
	CJri5kB26    string   `conf:"tls_min_version"`                                // 字节偏移 24；字段记录文件偏移 0x32da220。
	NH7oB3o7     string   `conf:"tls_max_version"`                                // 字节偏移 40；字段记录文件偏移 0x32da238。
	JOyIs3s9WAcA bool     `conf:"tls_prefer_server_cipher_suites" default:"true"` // 字节偏移 56；字段记录文件偏移 0x32da250。
	KVniDErTtx   string   `conf:"tls_cipher_suites"`                              // 字节偏移 64；字段记录文件偏移 0x32da268。
	CLr6mtlmQ    bool     `conf:"force_close_ssl" default:"false"`                // 字节偏移 80；字段记录文件偏移 0x32da280。
}

// k_lu9bPK：类型描述符 0x36da298；amd64 大小 80 字节，对齐 8 字节。
type k_lu9bPK struct {
	JHXrVGNjg31  bool   `conf:"tunnel_enable" default:"false"`       // 字节偏移 0；字段记录文件偏移 0x32da2f8。
	D38NXN3uV    string `conf:"tunnel_password"`                     // 字节偏移 8；字段记录文件偏移 0x32da310。
	GbyJkc       string `conf:"tunnel_type" default:"ws-tunnel"`     // 字节偏移 24；字段记录文件偏移 0x32da328。
	RtHB6a       bool   `conf:"tunnel_mux" default:"false"`          // 字节偏移 40；字段记录文件偏移 0x32da340。
	ZjL7PiTY     string `conf:"tunnel_method" default:"aes-128-gcm"` // 字节偏移 48；字段记录文件偏移 0x32da358。
	LYSz_mSWFTQp string `conf:"tunnel_ws_path" default:"/"`          // 字节偏移 64；字段记录文件偏移 0x32da370。
}

// vHzUrjrUZcY8：类型描述符 0x36da388；amd64 大小 64 字节，对齐 8 字节。
type vHzUrjrUZcY8 struct {
	KCYNbwBsZ    bool   `conf:"redis_ip_limit_enable" default:"false"` // 字节偏移 0；字段记录文件偏移 0x32da3e8。
	FXMKW6VcHH   string `conf:"redis_ip_limit_addr"`                   // 字节偏移 8；字段记录文件偏移 0x32da400。
	AhScWSQ      string `conf:"redis_ip_limit_password"`               // 字节偏移 24；字段记录文件偏移 0x32da418。
	ADewfG9Z64W  int    `conf:"redis_ip_limit_db" default:"0"`         // 字节偏移 40；字段记录文件偏移 0x32da430。
	NDRQGsmc132Y bool   `conf:"redis_ip_limit_tls" default:"false"`    // 字节偏移 48；字段记录文件偏移 0x32da448。
	BFkqcX7Dx    int    `conf:"redis_ip_limit_expiry" default:"60"`    // 字节偏移 56；字段记录文件偏移 0x32da460。
}

// x7sf3Tkb：类型描述符 0x36da478；amd64 大小 56 字节，对齐 8 字节。
type x7sf3Tkb struct {
	H2FCRVc      bool   `conf:"redis_enable" default:"false"`   // 字节偏移 0；字段记录文件偏移 0x32da4d8。
	IDyeXiEv     bool   `conf:"redis_tls" default:"false"`      // 字节偏移 1；字段记录文件偏移 0x32da4f0。
	Ij4bWLsuo_x4 string `conf:"redis_addr"`                     // 字节偏移 8；字段记录文件偏移 0x32da508。
	A9GLBE4PY    string `conf:"redis_password"`                 // 字节偏移 24；字段记录文件偏移 0x32da520。
	ZDYXh3Cuvxpw int    `conf:"redis_db" default:"0"`           // 字节偏移 40；字段记录文件偏移 0x32da538。
	Eb9vzdoE     int    `conf:"conn_limit_expiry" default:"60"` // 字节偏移 48；字段记录文件偏移 0x32da550。
}

// rXC9QSbKAoV0：类型描述符 0x36e94a8；amd64 大小 80 字节，对齐 8 字节。
type rXC9QSbKAoV0 struct {
	EFqkTZF6Wdql bool     `conf:"dy_limit_enable" default:"false"`      // 字节偏移 0；字段记录文件偏移 0x32e9508。
	CGGYGi0      string   `conf:"dy_limit_duration"`                    // 字节偏移 8；字段记录文件偏移 0x32e9520。
	PLAo2k09hxd  int      `conf:"dy_limit_trigger_time" default:"60"`   // 字节偏移 24；字段记录文件偏移 0x32e9538。
	ILhRYq68NSUC float64  `conf:"dy_limit_trigger_speed" default:"100"` // 字节偏移 32；字段记录文件偏移 0x32e9550。
	LxiE3fm4BLS  float64  `conf:"dy_limit_speed" default:"30"`          // 字节偏移 40；字段记录文件偏移 0x32e9568。
	Zao5joI9JY   int      `conf:"dy_limit_time" default:"600"`          // 字节偏移 48；字段记录文件偏移 0x32e9580。
	IrsBkXPlx    []string `conf:"dy_limit_white_user_id"`               // 字节偏移 56；字段记录文件偏移 0x32e9598。
}

// vW8lXY：类型描述符 0x36f2420；amd64 大小 64 字节，对齐 8 字节。
type vW8lXY struct {
	X4_AmA8j7    bool `conf:"ss_invalid_access_enable" default:"false"`               // 字节偏移 0；字段记录文件偏移 0x32f2480。
	UA4PZU       int  `conf:"ss_invalid_access_count" default:"30"`                   // 字节偏移 8；字段记录文件偏移 0x32f2498。
	PWFghdU      int  `conf:"ss_invalid_access_duration" default:"60"`                // 字节偏移 16；字段记录文件偏移 0x32f24b0。
	SndqWD26h    int  `conf:"ss_invalid_access_forbidden_time" default:"600"`         // 字节偏移 24；字段记录文件偏移 0x32f24c8。
	H1kZa6       bool `conf:"vmess_aead_invalid_access_enable" default:"false"`       // 字节偏移 32；字段记录文件偏移 0x32f24e0。
	Zdj5DFrb4WqO int  `conf:"vmess_aead_invalid_access_count" default:"30"`           // 字节偏移 40；字段记录文件偏移 0x32f24f8。
	BPUCSK_mS81  int  `conf:"vmess_aead_invalid_access_duration" default:"60"`        // 字节偏移 48；字段记录文件偏移 0x32f2510。
	ITmtDnIa8J   int  `conf:"vmess_aead_invalid_access_forbidden_time" default:"600"` // 字节偏移 56；字段记录文件偏移 0x32f2528。
}

// xIoYVOH：类型描述符 0x36f9f30；amd64 大小 104 字节，对齐 8 字节。
type xIoYVOH struct {
	Ff3xSe       bool      `conf:"vmess_reduce_memory" default:"false"` // 字节偏移 0；字段记录文件偏移 0x32f9f90。
	RIFsRvm1     string    `conf:"vmess_fallback_addr"`                 // 字节偏移 8；字段记录文件偏移 0x32f9fa8。
	YMm_xn_v     int       `conf:"vmess_fallback_port" default:"0"`     // 字节偏移 24；字段记录文件偏移 0x32f9fc0。
	SElnoFh_fIwr bool      `conf:"v2ray_reduce_memory" default:"false"` // 字节偏移 32；字段记录文件偏移 0x32f9fd8。
	GnOeZP5RIl   string    `conf:"v2ray_fallback_addr"`                 // 字节偏移 40；字段记录文件偏移 0x32f9ff0。
	HUt3s7n_8    int       `conf:"v2ray_fallback_port" default:"0"`     // 字节偏移 56；字段记录文件偏移 0x32fa008。
	SVuGy3CoRze  bool      `conf:"force_vmess_aead" default:"false"`    // 字节偏移 64；字段记录文件偏移 0x32fa020。
	ZFK_OBT1     string    `conf:"vmess_auth_mode" default:"auto"`      // 字节偏移 72；字段记录文件偏移 0x32fa038。
	dZWc0A2uIe   F6DuWgn9N // 字节偏移 88；字段记录文件偏移 0x32fa050。
}

// pKGHLtJsa：类型描述符 0x3708b68；amd64 大小 144 字节，对齐 8 字节。
type pKGHLtJsa struct {
	VD6q_75hbej2 int     `conf:"whmcs_port" default:"80"`    // 字节偏移 0；字段记录文件偏移 0x3308bc8。
	IlqtsyWu     float64 `conf:"whmcs_rate" default:"1"`     // 字节偏移 8；字段记录文件偏移 0x3308be0。
	X4Dd6nOZ_7   string  `conf:"whmcs_net" default:"tcp"`    // 字节偏移 16；字段记录文件偏移 0x3308bf8。
	Dd5Xza       bool    `conf:"whmcs_tls" default:"false"`  // 字节偏移 32；字段记录文件偏移 0x3308c10。
	W2WFtBWP     int     `conf:"whmcs_alter_id" default:"2"` // 字节偏移 40；字段记录文件偏移 0x3308c28。
	NQ3CpRlonU   string  `conf:"whmcs_ws_path" default:"/"`  // 字节偏移 48；字段记录文件偏移 0x3308c40。
	Ew45bIap     string  `conf:"whmcs_service_name"`         // 字节偏移 64；字段记录文件偏移 0x3308c58。
	Z9bFboOQlPiw string  `conf:"whmcs_password"`             // 字节偏移 80；字段记录文件偏移 0x3308c70。
	Et04j58qH    string  `conf:"whmcs_method"`               // 字节偏移 96；字段记录文件偏移 0x3308c88。
	EKv8Zjg7dn85 string  `conf:"whmcs_protocol"`             // 字节偏移 112；字段记录文件偏移 0x3308ca0。
	HmKQa3xt     string  `conf:"whmcs_obfs"`                 // 字节偏移 128；字段记录文件偏移 0x3308cb8。
}

// tUl2hoOYI0Ao：类型描述符 0x3708cd0；amd64 大小 144 字节，对齐 8 字节。
type tUl2hoOYI0Ao struct {
	PmIFLM       bool   `conf:"access_log_enable" default:"false"`                // 字节偏移 0；字段记录文件偏移 0x3308d30。
	WO1HZMaFg6Mx string `conf:"access_log_mode" default:"local"`                  // 字节偏移 8；字段记录文件偏移 0x3308d48。
	INBdnGu56l5  string `conf:"access_log_dir" default:"/etc/soga/access_log"`    // 字节偏移 24；字段记录文件偏移 0x3308d60。
	EuOaOO8qBn0Y int    `conf:"access_log_retention_hours" default:"72"`          // 字节偏移 40；字段记录文件偏移 0x3308d78。
	QpeDOsW      string `conf:"access_log_clickhouse_addr"`                       // 字节偏移 48；字段记录文件偏移 0x3308d90。
	CmPnvU       string `conf:"access_log_clickhouse_database" default:"default"` // 字节偏移 64；字段记录文件偏移 0x3308da8。
	G8aycKMf5aa  string `conf:"access_log_clickhouse_table" default:"access_log"` // 字节偏移 80；字段记录文件偏移 0x3308dc0。
	HnvQ_1vV     string `conf:"access_log_clickhouse_user" default:"default"`     // 字节偏移 96；字段记录文件偏移 0x3308dd8。
	ACWNbHb1C    string `conf:"access_log_clickhouse_password"`                   // 字节偏移 112；字段记录文件偏移 0x3308df0。
	DOZx3XRAm01  int    `conf:"access_log_clickhouse_batch_size" default:"1000"`  // 字节偏移 128；字段记录文件偏移 0x3308e08。
	D4WJQARA     bool   `conf:"access_log_clickhouse_auto_create" default:"true"` // 字节偏移 136；字段记录文件偏移 0x3308e20。
}

// i7BY8ClkrDX：类型描述符 0x371e5b8；amd64 大小 256 字节，对齐 8 字节。
type i7BY8ClkrDX struct {
	EtRWBSEhLQc  string   `conf:"type" default:"sspanel-uim"`     // 字节偏移 0；字段记录文件偏移 0x331e618。
	AG_mSFklcvj  string   `conf:"sub_type"`                       // 字节偏移 16；字段记录文件偏移 0x331e630。
	MnXMlyI      string   `conf:"server_type" default:"v2ray"`    // 字节偏移 32；字段记录文件偏移 0x331e648。
	NJSpG6       string   `conf:"api" default:"webapi"`           // 字节偏移 48；字段记录文件偏移 0x331e660。
	VfiX6TsWZxF  string   `conf:"webapi_url"`                     // 字节偏移 64；字段记录文件偏移 0x331e678。
	MKrU9XEJ     string   `conf:"webapi_mukey"`                   // 字节偏移 80；字段记录文件偏移 0x331e690。
	UnkaBGBuOUXV string   `conf:"webapi_key"`                     // 字节偏移 96；字段记录文件偏移 0x331e6a8。
	ApfXLZLmFXh  string   `conf:"db_host"`                        // 字节偏移 112；字段记录文件偏移 0x331e6c0。
	Mk5cDiQzaj   int      `conf:"db_port" default:"3306"`         // 字节偏移 128；字段记录文件偏移 0x331e6d8。
	CsBHZVUxTO   string   `conf:"db_name"`                        // 字节偏移 136；字段记录文件偏移 0x331e6f0。
	ZGaSmv       string   `conf:"db_user" default:"root"`         // 字节偏移 152；字段记录文件偏移 0x331e708。
	ASSy13OjR6aJ string   `conf:"db_password"`                    // 字节偏移 168；字段记录文件偏移 0x331e720。
	NawkeUTY_dJT int      `conf:"db_max_idle_time" default:"480"` // 字节偏移 184；字段记录文件偏移 0x331e738。
	EUexnK81w_   int      `conf:"db_max_idle_conn" default:"4"`   // 字节偏移 192；字段记录文件偏移 0x331e750。
	MpUqmS0      int      `conf:"db_max_conn" default:"4"`        // 字节偏移 200；字段记录文件偏移 0x331e768。
	N1Z81aDzg    []string `conf:"node_id" default:"0"`            // 字节偏移 208；字段记录文件偏移 0x331e780。
	t09QzN       []int    // 字节偏移 232；字段记录文件偏移 0x331e798。
}

// watWxZSBeVv：类型描述符 0x3730410；amd64 大小 224 字节，对齐 8 字节。
type watWxZSBeVv struct {
	Qxaw2_S      int      `conf:"check_interval" default:"60"`                                                                            // 字节偏移 0；字段记录文件偏移 0x3330470。
	Fak_7LgVV    int      `conf:"submit_interval" default:"60"`                                                                           // 字节偏移 8；字段记录文件偏移 0x3330488。
	KS63SD       int      `conf:"random_interval_range" default:"10"`                                                                     // 字节偏移 16；字段记录文件偏移 0x33304a0。
	JGgVNMp      bool     `conf:"forbidden_bit_torrent" default:"false"`                                                                  // 字节偏移 24；字段记录文件偏移 0x33304b8。
	DUpePpa39q   []string `conf:"forbidden_ports"`                                                                                        // 字节偏移 32；字段记录文件偏移 0x33304d0。
	F3BbbG_1qsA  bool     `conf:"auto_update" default:"false"`                                                                            // 字节偏移 56；字段记录文件偏移 0x33304e8。
	QgoKiaNp9    []string `conf:"domain_sniff" default:"tls,http,quic"`                                                                   // 字节偏移 64；字段记录文件偏移 0x3330500。
	Al5sHO4x90sY bool     `conf:"sniff_redirect" default:"false"`                                                                         // 字节偏移 88；字段记录文件偏移 0x3330518。
	XKO2Zot_hEk  bool     `conf:"detect_packet" default:"false"`                                                                          // 字节偏移 89；字段记录文件偏移 0x3330530。
	GozCogFPm    int      `conf:"detect_packet_max_len" default:"4096"`                                                                   // 字节偏移 96；字段记录文件偏移 0x3330548。
	Dh09He3a     bool     `conf:"auto_out_ip" default:"false"`                                                                            // 字节偏移 104；字段记录文件偏移 0x3330560。
	Osf81W0BWP   bool     `conf:"proxy_protocol" default:"false"`                                                                         // 字节偏移 105；字段记录文件偏移 0x3330578。
	UeBhd4nohvCF bool     `conf:"udp_proxy_protocol" default:"false"`                                                                     // 字节偏移 106；字段记录文件偏移 0x3330590。
	YOBYD0m      bool     `conf:"force_proxy_protocol" default:"false"`                                                                   // 字节偏移 107；字段记录文件偏移 0x33305a8。
	QzlEP1r3t    float64  `conf:"node_speed_limit" default:"0"`                                                                           // 字节偏移 112；字段记录文件偏移 0x33305c0。
	ItKGl9keQrlY int      `conf:"user_conn_limit" default:"0"`                                                                            // 字节偏移 120；字段记录文件偏移 0x33305d8。
	YyLGOC       float64  `conf:"user_speed_limit" default:"0"`                                                                           // 字节偏移 128；字段记录文件偏移 0x33305f0。
	P6KMmld      int      `conf:"user_tcp_limit" default:"0"`                                                                             // 字节偏移 136；字段记录文件偏移 0x3330608。
	JQC3mlZzpS69 int      `conf:"user_ip_limit_cidr_prefix_v4" default:"32"`                                                              // 字节偏移 144；字段记录文件偏移 0x3330620。
	BkelkzTWy    int      `conf:"user_ip_limit_cidr_prefix_v6" default:"128"`                                                             // 字节偏移 152；字段记录文件偏移 0x3330638。
	QiLulXrJUA0  bool     `conf:"submit_real_ip" default:"false"`                                                                         // 字节偏移 160；字段记录文件偏移 0x3330650。
	CVshaQaB33U  bool     `conf:"ban_private_ip" default:"false"`                                                                         // 字节偏移 161；字段记录文件偏移 0x3330668。
	JRQZn_       bool     `conf:"geo_update_enable" default:"false"`                                                                      // 字节偏移 162；字段记录文件偏移 0x3330680。
	BIdU1vYs     int      `conf:"geo_update_interval" default:"24"`                                                                       // 字节偏移 168；字段记录文件偏移 0x3330698。
	YGrHEaVoNKL  string   `conf:"geo_site_url" default:"https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat"` // 字节偏移 176；字段记录文件偏移 0x33306b0。
	YVMoJnAY     string   `conf:"geo_ip_url" default:"https://github.com/v2fly/geoip/releases/latest/download/geoip.dat"`                 // 字节偏移 192；字段记录文件偏移 0x33306c8。
	KI4eIM8J     float64  `conf:"submit_alive_ip_min_traffic" default:"0"`                                                                // 字节偏移 208；字段记录文件偏移 0x33306e0。
	DhXdMvN      float64  `conf:"submit_traffic_min_traffic" default:"0"`                                                                 // 字节偏移 216；字段记录文件偏移 0x33306f8。
}
