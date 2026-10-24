// 原样本 RTTI 字段及 JSON 标签；不是作者源码原文。
package webapi

import (
	"encoding/json/jsontext"
	C "recovered.local/soga/config"
	W "recovered.local/soga/panel/etagcache"
)

// 原类型大小 16 字节。
type Zkt5w_2o struct {
	iAFjt_tT9ZmG *C.FCEAfQJwDA7e
	oDI12LC      *W.NG0hCZyRHK
}

// 原类型大小 24 字节。
type LC4dqy struct {
	M9gh81OOM4ca int64  `json:"uid"`
	A6Zb5j5      string `json:"ip"`
}

// 原类型大小 24 字节。
type XnaEdZLZi_Lz struct {
	FD3yBPArxu   int    `json:"port"`
	AMjK_y00H3Vf string `json:"obfs_password"`
}

// 原类型大小 56 字节。
type Yq4D8wXbIQyY struct {
	OWYIa8fDu struct {
		F2WCOEc5Eu6 int "json:\"pull_interval\""
		I8kigte     int "json:\"push_interval\""
	} `json:"basic"`
	B0XVtyvn0k31 string         `json:"protocol"`
	OWRl87iaah   jsontext.Value `json:"config"`
}

// 原类型大小 24 字节。
type DjMeiycBVU struct {
	I4dYSFy5bqX  int64 `json:"uid"`
	QdZ1mt63MOLf int64 `json:"upload"`
	BJNQsH4BpW2  int64 `json:"download"`
}

// 原类型大小 40 字节。
type JV07s34JhN6z struct {
	SMh2QKAfSzk  int    `json:"port"`
	Afx4JqDX6KD4 string `json:"method"`
	DaDgO0Sh     string `json:"server_key"`
}

// 原类型大小 40 字节。
type MACjYTsGk7G struct {
	AnbiaLml0Bs  int64   `json:"id"`
	WtNHeT4      string  `json:"uuid"`
	YQbaJusq0g6N float64 `json:"speed_limit"`
	D7H3u6wL2q5Y int     `json:"device_limit"`
}

// 原类型大小 72 字节。
type Fp4V6S84pbeq struct {
	ASTH4uK7   int    `json:"port"`
	VsSEJhwCJT string `json:"security"`
	AnWvvG4Vva string `json:"transport"`
	ZgnJBd5rnF struct {
		NjZ_Np string "json:\"path\""
		Uce3uX string "json:\"service_name\""
	} `json:"transport_config"`
}

// 原类型大小 72 字节。
type PtTswARD struct {
	ASTH4uK7   int    `json:"port"`
	VsSEJhwCJT string `json:"security"`
	AnWvvG4Vva string `json:"transport"`
	ZgnJBd5rnF struct {
		NjZ_Np string "json:\"path\""
		Uce3uX string "json:\"service_name\""
	} `json:"transport_config"`
}

// 原类型大小 160 字节。
type LuuJNqNq1YTl struct {
	KzTz1UaVet int    `json:"port"`
	WbrpN6f5U  string `json:"flow"`
	U7TxoE88w8 string `json:"security"`
	SYTfrb     struct {
		HX23KFdaHxy string "json:\"reality_private_key\""
		StpFTNt5U   string "json:\"reality_server_addr\""
		T2S3gmIMZ   int    "json:\"reality_server_port\""
		KpKrDeAYv   string "json:\"reality_short_id\""
		HbXeOi732YO string "json:\"sni\""
	} `json:"security_config"`
	UePfBt4Gx string `json:"transport"`
	V51ymS4   struct {
		NjZ_Np string "json:\"path\""
		Uce3uX string "json:\"service_name\""
	} `json:"transport_config"`
}
type T8lmI8Ic []*MACjYTsGk7G
type QFeT8QtEgxS []*LC4dqy
type IodE2Y []*DjMeiycBVU
