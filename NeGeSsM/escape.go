package NeGeSsM

import "strings"

// v4V_xxf: 0x185a740。只处理竖线与 CR/LF；协议字段不调用它。
func v4V_xxf(s string) string {
	if strings.IndexAny(s, "|\r\n") < 0 {
		return s
	}
	return strings.NewReplacer("|", "_", "\r", "", "\n", "").Replace(s)
}

// oRjSReB: 0x185eb20。按字节转义，保留原代码对无效 UTF-8、HTML 字符、
// U+2028/U+2029 不替换的行为，不能用 encoding/json 或 strconv.Quote 替代。
func oRjSReB(b *strings.Builder, s string) {
	b.WriteByte('"')
	const hex = "0123456789abcdef"
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			if c < 0x20 {
				b.WriteString(`\u00`)
				b.WriteByte(hex[c>>4])
				b.WriteByte(hex[c&15])
			} else {
				b.WriteByte(c)
			}
		}
	}
	b.WriteByte('"')
}
