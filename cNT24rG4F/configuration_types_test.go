// 期望数据逐项转录自原样本 RTTI，未从待测声明计算。
// 类型前缀按源码包映射：ftLBcUpmkQ → cNT24rG4F。
// 此测试只验证源码的类型布局与记录一致，不证明业务行为等价。
package cNT24rG4F

import (
	"recovered.local/soga/ykfASHa7uK3k"
	"reflect"
	"runtime"
	"testing"
)

type fieldEvidence struct {
	name, typ, tag     string
	offset, size       uintptr
	align              int
	embedded, exported bool
}
type structEvidence struct {
	typ    reflect.Type
	size   uintptr
	align  int
	fields []fieldEvidence
}

func TestConfigurationRTTILayout(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("样本为 amd64；其他架构不作布局等价声明")
	}
	cases := []structEvidence{
		{reflect.TypeOf(cEgBVJvq{}), 16, 8, []fieldEvidence{
			{"M4qDq3KsUi", "string", "conf:\"routes_url\"", 0, 16, 8, false, true},
		}},
		{reflect.TypeOf(o9NA6hjqxImy{}), 16, 8, []fieldEvidence{
			{"RM5SFRF", "string", "conf:\"soga_key\"", 0, 16, 8, false, true},
		}},
		{reflect.TypeOf(piJb6eaS{}), 1, 1, []fieldEvidence{
			{"XpffP4ai", "bool", "conf:\"mieru_force_user_hint\" default:\"false\"", 0, 1, 1, false, true},
		}},
		{reflect.TypeOf(pkfwYwZ1aHu{}), 16, 8, []fieldEvidence{
			{"M4qDq3KsUi", "string", "conf:\"white_list_url\"", 0, 16, 8, false, true},
		}},
		{reflect.TypeOf(avoYN3oP9N4{}), 32, 8, []fieldEvidence{
			{"ZkkjP9", "string", "conf:\"soga_panel_url\"", 0, 16, 8, false, true},
			{"TzpbaXGDeaz", "string", "conf:\"soga_panel_key\"", 16, 16, 8, false, true},
		}},
		{reflect.TypeOf(bKCIPg{}), 24, 8, []fieldEvidence{
			{"BYiTyYPg3", "string", "conf:\"trojan_remote_addr\"", 0, 16, 8, false, true},
			{"EglOu0OdYex", "int", "conf:\"trojan_remote_port\" default:\"0\"", 16, 8, 8, false, true},
		}},
		{reflect.TypeOf(eFyNTD2_6X{}), 40, 8, []fieldEvidence{
			{"BooNOdXcSa3U", "string", "conf:\"block_list_url\"", 0, 16, 8, false, true},
			{"U2hwBFC", "[]string", "conf:\"block_list_white_user_id\"", 16, 24, 8, false, true},
		}},
		{reflect.TypeOf(paRsja6{}), 24, 8, []fieldEvidence{
			{"DlaO1l", "string", "conf:\"ss_obfs_mode\" default:\"relaxed\"", 0, 16, 8, false, true},
			{"zdRAxv6CNPMA", "ykfASHa7uK3k.VE7a8rvh", "", 16, 1, 1, false, false},
		}},
		{reflect.TypeOf(siZtoGWL{}), 24, 8, []fieldEvidence{
			{"BVi1B4f_v", "string", "conf:\"anytls_fallback_addr\"", 0, 16, 8, false, true},
			{"JGw3yW", "int", "conf:\"anytls_fallback_port\" default:\"0\"", 16, 8, 8, false, true},
		}},
		{reflect.TypeOf(a4aAMG{}), 32, 8, []fieldEvidence{
			{"PVSao2JIq", "bool", "conf:\"prometheus_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"OAHNXtGh", "string", "conf:\"prometheus_addr\"", 8, 16, 8, false, true},
			{"DXN_loYa", "int", "conf:\"prometheus_port\" default:\"19090\"", 24, 8, 8, false, true},
		}},
		{reflect.TypeOf(cp9sUcOF{}), 32, 8, []fieldEvidence{
			{"TsRVDHP0SA", "int", "conf:\"ip_user_cache_time\" default:\"24\"", 0, 8, 8, false, true},
			{"AROEPWNZd8SF", "bool", "conf:\"ip_user_cache_save_enable\" default:\"true\"", 8, 1, 1, false, true},
			{"FpIPwenv", "string", "conf:\"ip_user_cache_save_dir\" default:\"/etc/soga/\"", 16, 16, 8, false, true},
		}},
		{reflect.TypeOf(qNDGFoxnh{}), 32, 8, []fieldEvidence{
			{"PVSao2JIq", "bool", "conf:\"pprof_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"OAHNXtGh", "string", "conf:\"pprof_addr\"", 8, 16, 8, false, true},
			{"DXN_loYa", "int", "conf:\"pprof_port\" default:\"21010\"", 24, 8, 8, false, true},
		}},
		{reflect.TypeOf(zu_ScZG5P4{}), 40, 8, []fieldEvidence{
			{"NaavbD4r35", "string", "conf:\"log_level\" default:\"info\"", 0, 16, 8, false, true},
			{"HSBWD7", "string", "conf:\"log_file_dir\" default:\"/etc/soga/\"", 16, 16, 8, false, true},
			{"G590aNjEOVK", "int", "conf:\"log_file_retention_days\" default:\"7\"", 32, 8, 8, false, true},
		}},
		{reflect.TypeOf(juBqfxCScme{}), 40, 8, []fieldEvidence{
			{"DPqi4uoIl", "string", "conf:\"mu_suffix\" default:\"microsoft.com\"", 0, 16, 8, false, true},
			{"NXP6HvyFvmb9", "string", "conf:\"mu_regex\" default:\"%5m%id.%suffix\"", 16, 16, 8, false, true},
			{"XHTMQfO", "bool", "conf:\"ss_obfs_udp\" default:\"false\"", 32, 1, 1, false, true},
			{"NLpXVVM0x", "bool", "conf:\"ssr_cid_limit\" default:\"true\"", 33, 1, 1, false, true},
		}},
		{reflect.TypeOf(uhg6MpMKJ{}), 64, 8, []fieldEvidence{
			{"RHZA1kqpMSz", "[]string", "conf:\"default_dns\"", 0, 24, 8, false, true},
			{"PBvoBS2NX5", "int", "conf:\"dns_cache_time\" default:\"10\"", 24, 8, 8, false, true},
			{"EAOjNaQ6", "string", "conf:\"dns_strategy\" default:\"ipv4_first\"", 32, 16, 8, false, true},
			{"MFHBO42", "string", "conf:\"dns_config_url\"", 48, 16, 8, false, true},
		}},
		{reflect.TypeOf(aAGsqO{}), 48, 8, []fieldEvidence{
			{"VerfnYj", "string", "conf:\"listen\"", 0, 16, 8, false, true},
			{"Fg5sk_", "bool", "conf:\"tfo\" default:\"false\"", 16, 1, 1, false, true},
			{"Hni9vNgOp4g", "int", "conf:\"tcp_timeout\" default:\"120\"", 24, 8, 8, false, true},
			{"JtFIOYSoCNi", "int", "conf:\"udp_timeout\" default:\"2\"", 32, 8, 8, false, true},
			{"Arfg8UjxqP", "bool", "conf:\"mptcp\" default:\"false\"", 40, 1, 1, false, true},
		}},
		{reflect.TypeOf(cdqXEMYJO{}), 96, 8, []fieldEvidence{
			{"C2BNS6_QJC", "string", "conf:\"cert_file\"", 0, 16, 8, false, true},
			{"JISdid54S5t", "string", "conf:\"key_file\"", 16, 16, 8, false, true},
			{"MVmP0dSlHl", "string", "conf:\"cert_mode\" default:\"http\"", 32, 16, 8, false, true},
			{"Ga_YcC", "string", "conf:\"cert_domain\"", 48, 16, 8, false, true},
			{"AQ1zOwu1K", "string", "conf:\"cert_key_length\"", 64, 16, 8, false, true},
			{"Xs7eJ3aj6Kbl", "string", "conf:\"dns_provider\"", 80, 16, 8, false, true},
		}},
		{reflect.TypeOf(jkGZ0JHJw{}), 88, 8, []fieldEvidence{
			{"RS1kBzpBezso", "[]string", "conf:\"tls_alpn\" default:\"h2,http/1.1\"", 0, 24, 8, false, true},
			{"CJri5kB26", "string", "conf:\"tls_min_version\"", 24, 16, 8, false, true},
			{"NH7oB3o7", "string", "conf:\"tls_max_version\"", 40, 16, 8, false, true},
			{"JOyIs3s9WAcA", "bool", "conf:\"tls_prefer_server_cipher_suites\" default:\"true\"", 56, 1, 1, false, true},
			{"KVniDErTtx", "string", "conf:\"tls_cipher_suites\"", 64, 16, 8, false, true},
			{"CLr6mtlmQ", "bool", "conf:\"force_close_ssl\" default:\"false\"", 80, 1, 1, false, true},
		}},
		{reflect.TypeOf(k_lu9bPK{}), 80, 8, []fieldEvidence{
			{"JHXrVGNjg31", "bool", "conf:\"tunnel_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"D38NXN3uV", "string", "conf:\"tunnel_password\"", 8, 16, 8, false, true},
			{"GbyJkc", "string", "conf:\"tunnel_type\" default:\"ws-tunnel\"", 24, 16, 8, false, true},
			{"RtHB6a", "bool", "conf:\"tunnel_mux\" default:\"false\"", 40, 1, 1, false, true},
			{"ZjL7PiTY", "string", "conf:\"tunnel_method\" default:\"aes-128-gcm\"", 48, 16, 8, false, true},
			{"LYSz_mSWFTQp", "string", "conf:\"tunnel_ws_path\" default:\"/\"", 64, 16, 8, false, true},
		}},
		{reflect.TypeOf(vHzUrjrUZcY8{}), 64, 8, []fieldEvidence{
			{"KCYNbwBsZ", "bool", "conf:\"redis_ip_limit_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"FXMKW6VcHH", "string", "conf:\"redis_ip_limit_addr\"", 8, 16, 8, false, true},
			{"AhScWSQ", "string", "conf:\"redis_ip_limit_password\"", 24, 16, 8, false, true},
			{"ADewfG9Z64W", "int", "conf:\"redis_ip_limit_db\" default:\"0\"", 40, 8, 8, false, true},
			{"NDRQGsmc132Y", "bool", "conf:\"redis_ip_limit_tls\" default:\"false\"", 48, 1, 1, false, true},
			{"BFkqcX7Dx", "int", "conf:\"redis_ip_limit_expiry\" default:\"60\"", 56, 8, 8, false, true},
		}},
		{reflect.TypeOf(x7sf3Tkb{}), 56, 8, []fieldEvidence{
			{"H2FCRVc", "bool", "conf:\"redis_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"IDyeXiEv", "bool", "conf:\"redis_tls\" default:\"false\"", 1, 1, 1, false, true},
			{"Ij4bWLsuo_x4", "string", "conf:\"redis_addr\"", 8, 16, 8, false, true},
			{"A9GLBE4PY", "string", "conf:\"redis_password\"", 24, 16, 8, false, true},
			{"ZDYXh3Cuvxpw", "int", "conf:\"redis_db\" default:\"0\"", 40, 8, 8, false, true},
			{"Eb9vzdoE", "int", "conf:\"conn_limit_expiry\" default:\"60\"", 48, 8, 8, false, true},
		}},
		{reflect.TypeOf(rXC9QSbKAoV0{}), 80, 8, []fieldEvidence{
			{"EFqkTZF6Wdql", "bool", "conf:\"dy_limit_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"CGGYGi0", "string", "conf:\"dy_limit_duration\"", 8, 16, 8, false, true},
			{"PLAo2k09hxd", "int", "conf:\"dy_limit_trigger_time\" default:\"60\"", 24, 8, 8, false, true},
			{"ILhRYq68NSUC", "float64", "conf:\"dy_limit_trigger_speed\" default:\"100\"", 32, 8, 8, false, true},
			{"LxiE3fm4BLS", "float64", "conf:\"dy_limit_speed\" default:\"30\"", 40, 8, 8, false, true},
			{"Zao5joI9JY", "int", "conf:\"dy_limit_time\" default:\"600\"", 48, 8, 8, false, true},
			{"IrsBkXPlx", "[]string", "conf:\"dy_limit_white_user_id\"", 56, 24, 8, false, true},
		}},
		{reflect.TypeOf(vW8lXY{}), 64, 8, []fieldEvidence{
			{"X4_AmA8j7", "bool", "conf:\"ss_invalid_access_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"UA4PZU", "int", "conf:\"ss_invalid_access_count\" default:\"30\"", 8, 8, 8, false, true},
			{"PWFghdU", "int", "conf:\"ss_invalid_access_duration\" default:\"60\"", 16, 8, 8, false, true},
			{"SndqWD26h", "int", "conf:\"ss_invalid_access_forbidden_time\" default:\"600\"", 24, 8, 8, false, true},
			{"H1kZa6", "bool", "conf:\"vmess_aead_invalid_access_enable\" default:\"false\"", 32, 1, 1, false, true},
			{"Zdj5DFrb4WqO", "int", "conf:\"vmess_aead_invalid_access_count\" default:\"30\"", 40, 8, 8, false, true},
			{"BPUCSK_mS81", "int", "conf:\"vmess_aead_invalid_access_duration\" default:\"60\"", 48, 8, 8, false, true},
			{"ITmtDnIa8J", "int", "conf:\"vmess_aead_invalid_access_forbidden_time\" default:\"600\"", 56, 8, 8, false, true},
		}},
		{reflect.TypeOf(xIoYVOH{}), 104, 8, []fieldEvidence{
			{"Ff3xSe", "bool", "conf:\"vmess_reduce_memory\" default:\"false\"", 0, 1, 1, false, true},
			{"RIFsRvm1", "string", "conf:\"vmess_fallback_addr\"", 8, 16, 8, false, true},
			{"YMm_xn_v", "int", "conf:\"vmess_fallback_port\" default:\"0\"", 24, 8, 8, false, true},
			{"SElnoFh_fIwr", "bool", "conf:\"v2ray_reduce_memory\" default:\"false\"", 32, 1, 1, false, true},
			{"GnOeZP5RIl", "string", "conf:\"v2ray_fallback_addr\"", 40, 16, 8, false, true},
			{"HUt3s7n_8", "int", "conf:\"v2ray_fallback_port\" default:\"0\"", 56, 8, 8, false, true},
			{"SVuGy3CoRze", "bool", "conf:\"force_vmess_aead\" default:\"false\"", 64, 1, 1, false, true},
			{"ZFK_OBT1", "string", "conf:\"vmess_auth_mode\" default:\"auto\"", 72, 16, 8, false, true},
			{"dZWc0A2uIe", "cNT24rG4F.F6DuWgn9N", "", 88, 16, 8, false, false},
		}},
		{reflect.TypeOf(pKGHLtJsa{}), 144, 8, []fieldEvidence{
			{"VD6q_75hbej2", "int", "conf:\"whmcs_port\" default:\"80\"", 0, 8, 8, false, true},
			{"IlqtsyWu", "float64", "conf:\"whmcs_rate\" default:\"1\"", 8, 8, 8, false, true},
			{"X4Dd6nOZ_7", "string", "conf:\"whmcs_net\" default:\"tcp\"", 16, 16, 8, false, true},
			{"Dd5Xza", "bool", "conf:\"whmcs_tls\" default:\"false\"", 32, 1, 1, false, true},
			{"W2WFtBWP", "int", "conf:\"whmcs_alter_id\" default:\"2\"", 40, 8, 8, false, true},
			{"NQ3CpRlonU", "string", "conf:\"whmcs_ws_path\" default:\"/\"", 48, 16, 8, false, true},
			{"Ew45bIap", "string", "conf:\"whmcs_service_name\"", 64, 16, 8, false, true},
			{"Z9bFboOQlPiw", "string", "conf:\"whmcs_password\"", 80, 16, 8, false, true},
			{"Et04j58qH", "string", "conf:\"whmcs_method\"", 96, 16, 8, false, true},
			{"EKv8Zjg7dn85", "string", "conf:\"whmcs_protocol\"", 112, 16, 8, false, true},
			{"HmKQa3xt", "string", "conf:\"whmcs_obfs\"", 128, 16, 8, false, true},
		}},
		{reflect.TypeOf(tUl2hoOYI0Ao{}), 144, 8, []fieldEvidence{
			{"PmIFLM", "bool", "conf:\"access_log_enable\" default:\"false\"", 0, 1, 1, false, true},
			{"WO1HZMaFg6Mx", "string", "conf:\"access_log_mode\" default:\"local\"", 8, 16, 8, false, true},
			{"INBdnGu56l5", "string", "conf:\"access_log_dir\" default:\"/etc/soga/access_log\"", 24, 16, 8, false, true},
			{"EuOaOO8qBn0Y", "int", "conf:\"access_log_retention_hours\" default:\"72\"", 40, 8, 8, false, true},
			{"QpeDOsW", "string", "conf:\"access_log_clickhouse_addr\"", 48, 16, 8, false, true},
			{"CmPnvU", "string", "conf:\"access_log_clickhouse_database\" default:\"default\"", 64, 16, 8, false, true},
			{"G8aycKMf5aa", "string", "conf:\"access_log_clickhouse_table\" default:\"access_log\"", 80, 16, 8, false, true},
			{"HnvQ_1vV", "string", "conf:\"access_log_clickhouse_user\" default:\"default\"", 96, 16, 8, false, true},
			{"ACWNbHb1C", "string", "conf:\"access_log_clickhouse_password\"", 112, 16, 8, false, true},
			{"DOZx3XRAm01", "int", "conf:\"access_log_clickhouse_batch_size\" default:\"1000\"", 128, 8, 8, false, true},
			{"D4WJQARA", "bool", "conf:\"access_log_clickhouse_auto_create\" default:\"true\"", 136, 1, 1, false, true},
		}},
		{reflect.TypeOf(i7BY8ClkrDX{}), 256, 8, []fieldEvidence{
			{"EtRWBSEhLQc", "string", "conf:\"type\" default:\"sspanel-uim\"", 0, 16, 8, false, true},
			{"AG_mSFklcvj", "string", "conf:\"sub_type\"", 16, 16, 8, false, true},
			{"MnXMlyI", "string", "conf:\"server_type\" default:\"v2ray\"", 32, 16, 8, false, true},
			{"NJSpG6", "string", "conf:\"api\" default:\"webapi\"", 48, 16, 8, false, true},
			{"VfiX6TsWZxF", "string", "conf:\"webapi_url\"", 64, 16, 8, false, true},
			{"MKrU9XEJ", "string", "conf:\"webapi_mukey\"", 80, 16, 8, false, true},
			{"UnkaBGBuOUXV", "string", "conf:\"webapi_key\"", 96, 16, 8, false, true},
			{"ApfXLZLmFXh", "string", "conf:\"db_host\"", 112, 16, 8, false, true},
			{"Mk5cDiQzaj", "int", "conf:\"db_port\" default:\"3306\"", 128, 8, 8, false, true},
			{"CsBHZVUxTO", "string", "conf:\"db_name\"", 136, 16, 8, false, true},
			{"ZGaSmv", "string", "conf:\"db_user\" default:\"root\"", 152, 16, 8, false, true},
			{"ASSy13OjR6aJ", "string", "conf:\"db_password\"", 168, 16, 8, false, true},
			{"NawkeUTY_dJT", "int", "conf:\"db_max_idle_time\" default:\"480\"", 184, 8, 8, false, true},
			{"EUexnK81w_", "int", "conf:\"db_max_idle_conn\" default:\"4\"", 192, 8, 8, false, true},
			{"MpUqmS0", "int", "conf:\"db_max_conn\" default:\"4\"", 200, 8, 8, false, true},
			{"N1Z81aDzg", "[]string", "conf:\"node_id\" default:\"0\"", 208, 24, 8, false, true},
			{"t09QzN", "[]int", "", 232, 24, 8, false, false},
		}},
		{reflect.TypeOf(watWxZSBeVv{}), 224, 8, []fieldEvidence{
			{"Qxaw2_S", "int", "conf:\"check_interval\" default:\"60\"", 0, 8, 8, false, true},
			{"Fak_7LgVV", "int", "conf:\"submit_interval\" default:\"60\"", 8, 8, 8, false, true},
			{"KS63SD", "int", "conf:\"random_interval_range\" default:\"10\"", 16, 8, 8, false, true},
			{"JGgVNMp", "bool", "conf:\"forbidden_bit_torrent\" default:\"false\"", 24, 1, 1, false, true},
			{"DUpePpa39q", "[]string", "conf:\"forbidden_ports\"", 32, 24, 8, false, true},
			{"F3BbbG_1qsA", "bool", "conf:\"auto_update\" default:\"false\"", 56, 1, 1, false, true},
			{"QgoKiaNp9", "[]string", "conf:\"domain_sniff\" default:\"tls,http,quic\"", 64, 24, 8, false, true},
			{"Al5sHO4x90sY", "bool", "conf:\"sniff_redirect\" default:\"false\"", 88, 1, 1, false, true},
			{"XKO2Zot_hEk", "bool", "conf:\"detect_packet\" default:\"false\"", 89, 1, 1, false, true},
			{"GozCogFPm", "int", "conf:\"detect_packet_max_len\" default:\"4096\"", 96, 8, 8, false, true},
			{"Dh09He3a", "bool", "conf:\"auto_out_ip\" default:\"false\"", 104, 1, 1, false, true},
			{"Osf81W0BWP", "bool", "conf:\"proxy_protocol\" default:\"false\"", 105, 1, 1, false, true},
			{"UeBhd4nohvCF", "bool", "conf:\"udp_proxy_protocol\" default:\"false\"", 106, 1, 1, false, true},
			{"YOBYD0m", "bool", "conf:\"force_proxy_protocol\" default:\"false\"", 107, 1, 1, false, true},
			{"QzlEP1r3t", "float64", "conf:\"node_speed_limit\" default:\"0\"", 112, 8, 8, false, true},
			{"ItKGl9keQrlY", "int", "conf:\"user_conn_limit\" default:\"0\"", 120, 8, 8, false, true},
			{"YyLGOC", "float64", "conf:\"user_speed_limit\" default:\"0\"", 128, 8, 8, false, true},
			{"P6KMmld", "int", "conf:\"user_tcp_limit\" default:\"0\"", 136, 8, 8, false, true},
			{"JQC3mlZzpS69", "int", "conf:\"user_ip_limit_cidr_prefix_v4\" default:\"32\"", 144, 8, 8, false, true},
			{"BkelkzTWy", "int", "conf:\"user_ip_limit_cidr_prefix_v6\" default:\"128\"", 152, 8, 8, false, true},
			{"QiLulXrJUA0", "bool", "conf:\"submit_real_ip\" default:\"false\"", 160, 1, 1, false, true},
			{"CVshaQaB33U", "bool", "conf:\"ban_private_ip\" default:\"false\"", 161, 1, 1, false, true},
			{"JRQZn_", "bool", "conf:\"geo_update_enable\" default:\"false\"", 162, 1, 1, false, true},
			{"BIdU1vYs", "int", "conf:\"geo_update_interval\" default:\"24\"", 168, 8, 8, false, true},
			{"YGrHEaVoNKL", "string", "conf:\"geo_site_url\" default:\"https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat\"", 176, 16, 8, false, true},
			{"YVMoJnAY", "string", "conf:\"geo_ip_url\" default:\"https://github.com/v2fly/geoip/releases/latest/download/geoip.dat\"", 192, 16, 8, false, true},
			{"KI4eIM8J", "float64", "conf:\"submit_alive_ip_min_traffic\" default:\"0\"", 208, 8, 8, false, true},
			{"DhXdMvN", "float64", "conf:\"submit_traffic_min_traffic\" default:\"0\"", 216, 8, 8, false, true},
		}},
	}
	count := 0
	for _, c := range cases {
		t.Run(c.typ.Name(), func(t *testing.T) {
			if c.typ.Size() != c.size || c.typ.Align() != c.align {
				t.Errorf("结构布局：得到 size=%d align=%d，期望 %d %d", c.typ.Size(), c.typ.Align(), c.size, c.align)
			}
			if c.typ.NumField() != len(c.fields) {
				t.Fatalf("字段数量：得到 %d，期望 %d", c.typ.NumField(), len(c.fields))
			}
			for i, w := range c.fields {
				f := c.typ.Field(i)
				count++
				if f.Name != w.name || f.Type.String() != w.typ || string(f.Tag) != w.tag || f.Offset != w.offset || f.Type.Size() != w.size || f.Type.Align() != w.align || f.Anonymous != w.embedded || f.IsExported() != w.exported {
					t.Errorf("字段 %d 与 RTTI 不符：得到 %+v；期望 %+v", i, f, w)
				}
			}
		})
	}
	if len(cases) != 28 || count != 160 {
		t.Fatalf("证据覆盖错误：结构 %d，字段 %d", len(cases), count)
	}
}

func TestNamedUnderlyingTypes(t *testing.T) {
	a := reflect.TypeOf(F6DuWgn9N(""))
	b := reflect.TypeOf(ykfASHa7uK3k.VE7a8rvh(0))
	if a.Name() != "F6DuWgn9N" || a.Kind() != reflect.String || a.PkgPath() != "recovered.local/soga/cNT24rG4F" {
		t.Fatalf("字符串命名类型身份错误：%v", a)
	}
	if b.Name() != "VE7a8rvh" || b.Kind() != reflect.Uint8 || b.PkgPath() != "recovered.local/soga/ykfASHa7uK3k" || b.Size() != 1 || b.Align() != 1 {
		t.Fatalf("字节命名类型身份错误：%v", b)
	}
	if runtime.GOARCH == "amd64" && (a.Size() != 16 || a.Align() != 8) {
		t.Fatalf("字符串布局错误：size=%d align=%d", a.Size(), a.Align())
	}
}
