// 根据原样本 Go RTTI 恢复的数据类型，不是作者源码原文。
// RTTI 名称前缀为 wyihwaFO，uncommon_package_path 为 ZrjfjK8Y_tM9；
// 此处按后者组织包目录，保留实际混淆类型名、字段名和 JSON 标签。
package webapibase

// GIAIBDeJa25 对应 RTTI 0x35f3bc0，原布局大小 16 字节。
type GIAIBDeJa25 struct {
	RVKq_aLvfWyF string `json:"serviceName"`
}

// Re7aaf 对应 RTTI 0x35f3c38，原布局大小 16 字节。
type Re7aaf struct {
	JCRhkrluQQY string `json:"path"`
}

// D9MPLRTuA 对应 RTTI 0x3639c80，原布局大小 40 字节。
type D9MPLRTuA struct {
	EWFMfZX     string    `json:"message"`
	Ta593Rq2lSa Uhu9JY7Vj `json:"users"`
}

// WVfKv8zPvt 对应 RTTI 0x36bb6c0，原布局大小 40 字节。
type WVfKv8zPvt struct {
	IDLJXX6F    int64   `json:"id"`
	UH0HXa5O    string  `json:"uuid"`
	DwdlKBQpcgY float64 `json:"speed_limit"`
	QeuUtZ8f    int     `json:"device_limit"`
}

// Uhu9JY7Vj 对应 0x35903c0；元素描述符 0x3477d40 指向 *WVfKv8zPvt。
type Uhu9JY7Vj []*WVfKv8zPvt
