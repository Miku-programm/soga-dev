// 根据原 ELF RTTI/uncommon 包路径恢复，不是作者源码文件原文。
package nodeconvert

import (
	C "recovered.local/soga/config"
)

// RTTI 0x35f3da0 size 8
type DrEoaRCt2 struct {
	pBXis_w *C.FCEAfQJwDA7e
}

// RTTI 0x35f3e18 size 8
type EWcuZ1Kttay struct {
	IiyqiyH0 int64 `json:"uid"`
}

// RTTI 0x3639f50 size 24
type ZbwhxuO8Dq9a struct {
	DqE8i6s5lS  int64  `json:"uid"`
	JqoB2xKgVmg string `json:"ip"`
}

// RTTI 0x3639fe0 size 48
type AnlFZbNq struct {
	Dps5ZWL  []*PNf9Em87 `json:"inbounds"`
	YhC3piOh struct {
		UwS4pcm HidyXaJnC `json:"rules"`
	} `json:"routing"`
}

// RTTI 0x363a070 size 40
type Oh1oNuTp struct {
	DCbciSyJ_c string          `json:"msg"`
	XH32wVY    []*AenyiikKLY67 `json:"data"`
}

// RTTI 0x365eb10 size 24
type EMisRMdmr8d struct {
	W0KlV7rEtxgz int64 `json:"uid"`
	QavUxYiaO    int64 `json:"u"`
	Vdf1KaaiPt   int64 `json:"d"`
}

// RTTI 0x365ebb8 size 64
type Mwi3dEVZ struct {
	F3hpMd0 []string `json:"domain"`
	DaFBjr  []string `json:"protocol"`
	Ab0SdSa string   `json:"outboundTag"`
}

// RTTI 0x36bb840 size 120
type PNf9Em87 struct {
	SioIDq    int    `json:"port"`
	QxNm4UvLN string `json:"protocol"`
	VhYAkS    struct {
		P9ydbTayFu string `json:"method"`
		XUfl9V     string `json:"password"`
	} `json:"settings"`
	Jedod2gPZ4UT struct {
		WZVYFl     string `json:"network"`
		Id1HxAg5Oy string `json:"security"`
		D5bsRALb   struct {
			JCRhkrluQQY string `json:"path"`
		} `json:"wsSettings"`
		Ol5ZxI8vO struct {
			RVKq_aLvfWyF string `json:"serviceName"`
		} `json:"grpcSettings"`
	} `json:"streamingSettings"`
}

// RTTI 0x36d0fe8 size 72
type Vg6Nm38GB9 struct {
	EUjZVPy7nt string `json:"cpu"`
	Fpzi5hx    string `json:"mem"`
	Zn5cIP_V   string `json:"net"`
	JzgG45gXrE string `json:"disk"`
	PJJnV8     uint64 `json:"uptime"`
}

// RTTI 0x36e26c8 size 72
type AenyiikKLY67 struct {
	PqTXyjlHNDD_ int64   `json:"id"`
	SDB5DM3E68Y  string  `json:"uuid"`
	HZO6BjGa     string  `json:"secret"`
	NRLcM3Xm2zgj string  `json:"password"`
	C2TUmBU9gT   int     `json:"dt"`
	WEkiWDf5si   float64 `json:"st"`
}

// RTTI 0x3590528
type Ga4NII7c12I []*EWcuZ1Kttay

// RTTI 0x3590570
type IvwLwWM []*EMisRMdmr8d

// RTTI 0x35905b8
type X22YvSZk []*ZbwhxuO8Dq9a

// RTTI 0x35da358
type HidyXaJnC []*Mwi3dEVZ
