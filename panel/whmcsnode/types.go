// 本文件依据样本类型元数据恢复，不是作者原文。包目录按 uncommon_package_path；
// RTTI 显示前缀经过独立混淆，不将其误称为原目录。
package whmcsnode

import (
	"recovered.local/soga/audit/detector"
)

// J5tD2c9：原 RTTI 0x35eeee0，显示名称 rES_2TTolaU.J5tD2c9。
type J5tD2c9 struct {
	detector.GMLWI3
}

// PqISXuLk：原 RTTI 0x36584b0，显示名称 rES_2TTolaU.PqISXuLk。
type PqISXuLk struct {
	Jrq3niigv int64   `json:"uid"`
	Jm6azqIw  string  `json:"password"`
	Bexa45zkL float64 `json:"speed_limit"`
}

// BoM67db：原 RTTI 0x3658558，显示名称 rES_2TTolaU.BoM67db。
type BoM67db struct {
	IdNatD8MGczW int         `json:"code"`
	UzQ0WhgKfU   string      `json:"message"`
	GPETCf2      []*PqISXuLk `json:"data"`
}

// XhDo0vyZH：原 RTTI 0x3658600，显示名称 rES_2TTolaU.XhDo0vyZH。
type XhDo0vyZH struct {
	JLAJhSUZPpQv float64 `json:"speed_limit"`
	JQsJVj       int     `json:"client_limit"`
	Jbn82w0HZp   int     `json:"trojan_port"`
}

// WmjoeO：原 RTTI 0x36586a8，显示名称 rES_2TTolaU.WmjoeO。
type WmjoeO struct {
	IdNatD8MGczW int       `json:"code"`
	UzQ0WhgKfU   string    `json:"message"`
	GPETCf2      XhDo0vyZH `json:"data"`
}
