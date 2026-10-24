// 依据原样本 RTTI 恢复；不代表作者源码文件。
package sogav1

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/transport/xhttp"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/ssrnames"
	"recovered.local/soga/crypto/ssmethod"
	"recovered.local/soga/panel/etagcache"
)

// CsKk5VzGb：原 RTTI 0x36392f0，amd64 大小 16 字节。
type CsKk5VzGb struct {
	cI03v8q7   *etagcache.NG0hCZyRHK
	gAaW6Z7Uh3 *config.FCEAfQJwDA7e
}

// Tw_nfz：原 RTTI 0x3639380，amd64 大小 32 字节。
type Tw_nfz struct {
	AdYIli      int64    `json:"id"`
	FmXlcnZSlSU []string `json:"ips"`
}

// V91695M：原 RTTI 0x3639410，amd64 大小 16 字节。
type V91695M struct {
	VJB4pQGJ     int64 `json:"user_id"`
	JtAQyC57KMKb int   `json:"audit_id"`
}

// HjRv_Yy3uom：原 RTTI 0x36394a0，amd64 大小 24 字节。
type HjRv_Yy3uom struct {
	RFcywH9K   int    `json:"id"`
	FI29zO1tSx string `json:"rule"`
}

// GX5_vLs：原 RTTI 0x3639530，amd64 大小 416 字节。
type GX5_vLs struct {
	XI5uZkt struct {
		H5wL4HQdqP int     "json:\"pull_interval\""
		ULhXjJaRde int     "json:\"push_interval\""
		ZZ6bu_Ki   float64 "json:\"speed_limit\""
	} `json:"basic"`
	I8ECEmdqGb struct {
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
	} `json:"config"`
}

// OFFj2pY8Z：原 RTTI 0x36395c0，amd64 大小 24 字节。
type OFFj2pY8Z struct {
	NIg3b55LCOZ int    `json:"code"`
	TKIFA83m_r  string `json:"message"`
}

// OTsrnZKa：原 RTTI 0x365e918，amd64 大小 24 字节。
type OTsrnZKa struct {
	X7cdH6hK  int64 `json:"id"`
	SszZxQ_BV int64 `json:"u"`
	CEKdEFoFV int64 `json:"d"`
}

// CgmhUmlxAlbk：原 RTTI 0x36bb540，amd64 大小 296 字节。
type CgmhUmlxAlbk struct {
	DR1aHL     string `json:"host,omitempty"`
	ZjmKtZ4NO  string `json:"path,omitempty"`
	LaGQMWdidE string `json:"mode,omitempty"`
	TvabDDO    struct {
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
	} `json:"extra"`
}

// LygoL0kC5：原 RTTI 0x36d0e38，amd64 大小 64 字节。
type LygoL0kC5 struct {
	BRhjXec float64 `json:"cpu"`
	ZmWD5iP struct {
		BcEb4PnfTS9 uint64 "json:\"used\""
		JvmfpVk     uint64 "json:\"total\""
	} `json:"mem"`
	J9mVD1 struct {
		BcEb4PnfTS9 uint64 "json:\"used\""
		JvmfpVk     uint64 "json:\"total\""
	} `json:"swap"`
	EpTvgGsJ5 struct {
		BcEb4PnfTS9 uint64 "json:\"used\""
		JvmfpVk     uint64 "json:\"total\""
	} `json:"disk"`
	EUw7Qwoa4K0 uint64 `json:"uptime"`
}

// D6uXyl88ab7：原 RTTI 0x36ef070，amd64 大小 80 字节。
type D6uXyl88ab7 struct {
	IMizGzpEPYfj int64   `json:"id"`
	S7USjTWE8o   string  `json:"uuid"`
	DHknGz       string  `json:"password"`
	WUNMcwtiDw   string  `json:"username"`
	GZJoKAlg8K9  float64 `json:"speed_limit"`
	LNDBhsiH     int     `json:"device_limit"`
	DGYgUabBB6Jx int     `json:"tcp_limit"`
}

type AQ3wazHBAA []*HjRv_Yy3uom
type DdDyLf6D []*D6uXyl88ab7
