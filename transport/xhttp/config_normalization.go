package xhttp

import "strings"

// 按样本 0x1251280 恢复。样本无条件补末尾斜杠；不套用上游后来
// 按 session/seq placement 决定是否补斜杠的行为。
func (c *R0J19auWof) GetNormalizedPath() string {
	p := strings.SplitN(c.HoBz4V, "?", 2)[0]
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	if !strings.HasSuffix(p, "/") {
		p += "/"
	}
	return p
}

// 原入口 0x1251340；保留问号后的原文，不执行 URL 解码。
func (c *R0J19auWof) GetNormalizedQuery() string {
	p := strings.SplitN(c.HoBz4V, "?", 2)
	if len(p) > 1 {
		return p[1]
	}
	return ""
}

// 原入口 0x12525e0。
func (c *R0J19auWof) GetNormalizedUplinkHTTPMethod() string {
	if c.KJMkBJ != "" {
		return c.KJMkBJ
	}
	return "POST"
}

// 原入口 0x1252620。样本只检查上界为零，不检查负值或区间倒置。
func (c *R0J19auWof) GetNormalizedScMaxEachPostBytes() MUVoWl0 {
	if c.ULnN_iDWLIXN != nil && c.ULnN_iDWLIXN.HHTPwX0N != 0 {
		return *c.ULnN_iDWLIXN
	}
	return MUVoWl0{1000000, 1000000}
}

// 原入口 0x1252640。
func (c *R0J19auWof) GetNormalizedScMaxBufferedPosts() int64 {
	if c.LEaRo0 != 0 {
		return c.LEaRo0
	}
	return 30
}

// 原入口 0x1252660。
func (c *R0J19auWof) GetNormalizedScStreamUpServerSecs() MUVoWl0 {
	if c.Iv950ouQ5AW != nil && c.Iv950ouQ5AW.HHTPwX0N != 0 {
		return *c.Iv950ouQ5AW
	}
	return MUVoWl0{20, 80}
}

// 原入口 0x1252680。只在下界小于 64 时同时提升上界。
func (c *R0J19auWof) GetNormalizedUplinkChunkSize() MUVoWl0 {
	if c.JvaHmm != nil && c.JvaHmm.HHTPwX0N != 0 {
		r := *c.JvaHmm
		if r.AS9zG362 < 64 {
			r.AS9zG362 = 64
			r.HHTPwX0N = max(64, r.HHTPwX0N)
		}
		return r
	}
	switch c.Aonc19Hp33E {
	case "cookie":
		return MUVoWl0{2048, 3072}
	case "header":
		return MUVoWl0{3000, 4000}
	default:
		return c.GetNormalizedScMaxEachPostBytes()
	}
}

// 原入口 0x1252720。
func (c *R0J19auWof) GetNormalizedServerMaxHeaderBytes() int {
	if c.Diwr8ey > 0 {
		return int(c.Diwr8ey)
	}
	return 8192
}

// 原入口 0x1252740。
func (c *R0J19auWof) GetNormalizedSessionPlacement() string {
	if c.XEKp9SbqKQ != "" {
		return c.XEKp9SbqKQ
	}
	return "path"
}

// 原入口 0x1252780。
func (c *R0J19auWof) GetNormalizedSeqPlacement() string {
	if c.HqtZJaRZNw != "" {
		return c.HqtZJaRZNw
	}
	return "path"
}

// 原入口 0x12527c0。
func (c *R0J19auWof) GetNormalizedUplinkDataPlacement() string {
	if c.Aonc19Hp33E != "" {
		return c.Aonc19Hp33E
	}
	return "body"
}

// 原入口 0x1252800。两个混淆字符串由原指令受限执行取回。
func (c *R0J19auWof) GetNormalizedSessionKey() string {
	if c.G9oU9diBqr != "" {
		return c.G9oU9diBqr
	}
	switch c.GetNormalizedSessionPlacement() {
	case "header":
		return "X-Session"
	case "cookie", "query":
		return "x_session"
	default:
		return ""
	}
}

// 原入口 0x1252960。
func (c *R0J19auWof) GetNormalizedSeqKey() string {
	if c.NPLvCM != "" {
		return c.NPLvCM
	}
	switch c.GetNormalizedSeqPlacement() {
	case "header":
		return "X-Seq"
	case "cookie", "query":
		return "x_seq"
	default:
		return ""
	}
}

// 原入口 0x126d000。
func (c *R0J19auWof) GetNormalizedXPaddingBytes() MUVoWl0 {
	if c.N2C4XMihD != nil && c.N2C4XMihD.HHTPwX0N != 0 {
		return *c.N2C4XMihD
	}
	return MUVoWl0{100, 1000}
}
