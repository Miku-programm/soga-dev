package sogav1

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/transport/xhttp"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/ssrnames"
	"recovered.local/soga/crypto/ssmethod"
	"recovered.local/soga/panel/etagcache"
	"reflect"
	"testing"
)

func TestOriginalRTTI(t *testing.T) {
	type field struct {
		name   string
		offset uintptr
		typ    reflect.Type
		tag    string
	}
	for _, x := range []struct {
		typ    reflect.Type
		size   uintptr
		fields []field
	}{
		{reflect.TypeFor[CsKk5VzGb](), 16, []field{
			{"cI03v8q7", 0, reflect.TypeFor[*etagcache.NG0hCZyRHK](), ""},
			{"gAaW6Z7Uh3", 8, reflect.TypeFor[*config.FCEAfQJwDA7e](), ""},
		}},
		{reflect.TypeFor[Tw_nfz](), 32, []field{
			{"AdYIli", 0, reflect.TypeFor[int64](), "json:\"id\""},
			{"FmXlcnZSlSU", 8, reflect.TypeFor[[]string](), "json:\"ips\""},
		}},
		{reflect.TypeFor[V91695M](), 16, []field{
			{"VJB4pQGJ", 0, reflect.TypeFor[int64](), "json:\"user_id\""},
			{"JtAQyC57KMKb", 8, reflect.TypeFor[int](), "json:\"audit_id\""},
		}},
		{reflect.TypeFor[HjRv_Yy3uom](), 24, []field{
			{"RFcywH9K", 0, reflect.TypeFor[int](), "json:\"id\""},
			{"FI29zO1tSx", 8, reflect.TypeFor[string](), "json:\"rule\""},
		}},
		{reflect.TypeFor[GX5_vLs](), 416, []field{
			{"XI5uZkt", 0, reflect.TypeFor[struct {
				H5wL4HQdqP int     "json:\"pull_interval\""
				ULhXjJaRde int     "json:\"push_interval\""
				ZZ6bu_Ki   float64 "json:\"speed_limit\""
			}](), "json:\"basic\""},
			{"I8ECEmdqGb", 24, reflect.TypeFor[struct {
				GJd93ieNwt   int                     "json:\"port\""
				BcHm_C3XOt   types.BoulTn31     "json:\"stream_type\""
				Gpt4dj       types.D7FHIoaa8s4a "json:\"tls_type\""
				X78zRUstxhd  string                  "json:\"path\""
				XZmYiPZisD8  string                  "json:\"service_name\""
				YDJXeuot82   ssmethod.C6q3Upu        "json:\"cipher\""
				LXsSwyq7YZ   string                  "json:\"password\""
				CWNyQiAMJQ82 string                  "json:\"host\""
				NrtZd7       string                  "json:\"obfs\""
				DDVa4fX      string                  "json:\"obfs_password\""
				MV3efn       ssrnames.ZvAxOt1l7Ar "json:\"method\""
				KRr2OnixG8   ssrnames.Mu0vMIIs    "json:\"protocol\""
				CG9XBFFHZw14 string                  "json:\"single_port_type\""
				G9GmqrKxu9iJ string                  "json:\"flow\""
				Tm4m67gYc    []string                "json:\"server_names\""
				RtH5luH9GF0t string                  "json:\"private_key\""
				Mnuh5QkiZaeP []string                "json:\"short_ids\""
				DTjgP1       string                  "json:\"dest\""
				B1fm0zbCjL   string                  "json:\"decryption\""
				HAei1i_      *CgmhUmlxAlbk           "json:\"xhttp_settings\""
				RU9mVZaMebUI float64                 "json:\"up_mbps\""
				Sj7k1l11MYq  float64                 "json:\"down_mbps\""
				JcFiIM       []string                "json:\"padding_scheme\""
				LtcfZGsn7n   string                  "json:\"transport\""
				Qz3ffZ1n4oYU string                  "json:\"traffic_pattern\""
			}](), "json:\"config\""},
		}},
		{reflect.TypeFor[OFFj2pY8Z](), 24, []field{
			{"NIg3b55LCOZ", 0, reflect.TypeFor[int](), "json:\"code\""},
			{"TKIFA83m_r", 8, reflect.TypeFor[string](), "json:\"message\""},
		}},
		{reflect.TypeFor[OTsrnZKa](), 24, []field{
			{"X7cdH6hK", 0, reflect.TypeFor[int64](), "json:\"id\""},
			{"SszZxQ_BV", 8, reflect.TypeFor[int64](), "json:\"u\""},
			{"CEKdEFoFV", 16, reflect.TypeFor[int64](), "json:\"d\""},
		}},
		{reflect.TypeFor[CgmhUmlxAlbk](), 296, []field{
			{"DR1aHL", 0, reflect.TypeFor[string](), "json:\"host,omitempty\""},
			{"ZjmKtZ4NO", 16, reflect.TypeFor[string](), "json:\"path,omitempty\""},
			{"LaGQMWdidE", 32, reflect.TypeFor[string](), "json:\"mode,omitempty\""},
			{"TvabDDO", 48, reflect.TypeFor[struct {
				JyvbKKy      map[string]string "json:\"headers,omitempty\""
				RBXr5h4RdG   *xhttp.MUVoWl0   "json:\"xPaddingBytes,omitempty\""
				Zs29Z4s_     bool              "json:\"noSSEHeader,omitempty\""
				XlNMXRvv2Xv  *xhttp.MUVoWl0   "json:\"scMaxEachPostBytes,omitempty\""
				BKEah4PHj6   int64             "json:\"scMaxBufferedPosts,omitempty\""
				PqnRBGUA     *xhttp.MUVoWl0   "json:\"scStreamUpServerSecs,omitempty\""
				Gz8Z5EHS37zf bool              "json:\"xPaddingObfsMode,omitempty\""
				Qxwak5       string            "json:\"xPaddingKey,omitempty\""
				WFCdvAcCb    string            "json:\"xPaddingHeader,omitempty\""
				ScFUrULw0u   string            "json:\"xPaddingPlacement,omitempty\""
				NT4IRoa6GK2R string            "json:\"xPaddingMethod,omitempty\""
				Q3GsIzXFb    string            "json:\"uplinkHTTPMethod,omitempty\""
				LU0dA0xkmKyH string            "json:\"sessionPlacement,omitempty\""
				MY2CflJOwKzG string            "json:\"sessionKey,omitempty\""
				S3Go8lWC     string            "json:\"seqPlacement,omitempty\""
				TdWIFCHZ5w9E string            "json:\"seqKey,omitempty\""
				XGJef7FJgZ   string            "json:\"uplinkDataPlacement,omitempty\""
				UdR4T5STN    string            "json:\"uplinkDataKey,omitempty\""
				SsIGB4WX     *xhttp.MUVoWl0   "json:\"uplinkChunkSize,omitempty\""
				XkM_38AWGOZ5 int32             "json:\"serverMaxHeaderBytes,omitempty\""
			}](), "json:\"extra\""},
		}},
		{reflect.TypeFor[LygoL0kC5](), 64, []field{
			{"BRhjXec", 0, reflect.TypeFor[float64](), "json:\"cpu\""},
			{"ZmWD5iP", 8, reflect.TypeFor[struct {
				BcEb4PnfTS9 uint64 "json:\"used\""
				JvmfpVk     uint64 "json:\"total\""
			}](), "json:\"mem\""},
			{"J9mVD1", 24, reflect.TypeFor[struct {
				BcEb4PnfTS9 uint64 "json:\"used\""
				JvmfpVk     uint64 "json:\"total\""
			}](), "json:\"swap\""},
			{"EpTvgGsJ5", 40, reflect.TypeFor[struct {
				BcEb4PnfTS9 uint64 "json:\"used\""
				JvmfpVk     uint64 "json:\"total\""
			}](), "json:\"disk\""},
			{"EUw7Qwoa4K0", 56, reflect.TypeFor[uint64](), "json:\"uptime\""},
		}},
		{reflect.TypeFor[D6uXyl88ab7](), 80, []field{
			{"IMizGzpEPYfj", 0, reflect.TypeFor[int64](), "json:\"id\""},
			{"S7USjTWE8o", 8, reflect.TypeFor[string](), "json:\"uuid\""},
			{"DHknGz", 24, reflect.TypeFor[string](), "json:\"password\""},
			{"WUNMcwtiDw", 40, reflect.TypeFor[string](), "json:\"username\""},
			{"GZJoKAlg8K9", 56, reflect.TypeFor[float64](), "json:\"speed_limit\""},
			{"LNDBhsiH", 64, reflect.TypeFor[int](), "json:\"device_limit\""},
			{"DGYgUabBB6Jx", 72, reflect.TypeFor[int](), "json:\"tcp_limit\""},
		}},
	} {
		if x.typ.Size() != x.size || x.typ.NumField() != len(x.fields) {
			t.Fatalf("layout %v", x.typ)
		}
		for i, w := range x.fields {
			f := x.typ.Field(i)
			if f.Name != w.name || f.Offset != w.offset || f.Type != w.typ || string(f.Tag) != w.tag {
				t.Fatalf("field %v.%s", x.typ, f.Name)
			}
		}
	}
}
