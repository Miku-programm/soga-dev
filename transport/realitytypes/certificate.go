// 由原样本 RTTI 恢复的类型表示，并非作者原始文件。
package realitytypes

// RTTI 显示前缀 qih_GaofTcgg，实际 uncommon package path 为本包名。
type XBL6HZJUA3O string

// ZK5THfX 对应原 120 字节证书记录，保留实际字段及序列化标签。
type ZK5THfX struct {
	QvNk7x1dJQ int         `json:"id" gorm:"column:id;primaryKey;autoIncrement"`            // 原偏移 0
	Pgpq5sNc   string      `json:"domain" gorm:"column:domain;uniqueIndex:idx_cert_domain"` // 原偏移 8
	IFa4vG     bool        `json:"auto_renew" gorm:"column:auto_renew"`                     // 原偏移 24
	FHaYRa     string      `json:"email" gorm:"column:email"`                               // 原偏移 32
	VemC_yd    XBL6HZJUA3O `json:"provider_type" gorm:"column:provider_type"`               // 原偏移 48
	JdrlDaD1k  string      `json:"provider_config" gorm:"column:provider_config"`           // 原偏移 64
	Y1mdfY     string      `json:"certContent" gorm:"column:cert_content"`                  // 原偏移 80
	SzI2MbNsk  string      `json:"keyContent" gorm:"column:key_content"`                    // 原偏移 96
	ICiRoQoM1K int64       `json:"expiryTime" gorm:"-"`                                     // 原偏移 112
}
