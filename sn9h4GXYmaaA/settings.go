package sn9h4GXYmaaA

import "strings"

type L4jWYITSyXGq map[string]string

// DsbUyzhj9 @0x188afa0：只按换行和第一个等号分隔，不修剪空白。
func DsbUyzhj9(b []byte) L4jWYITSyXGq {
	m := make(L4jWYITSyXGq)
	for _, line := range strings.Split(string(b), "\n") {
		if p := strings.SplitN(line, "=", 2); len(p) == 2 {
			m[p[0]] = p[1]
		}
	}
	return m
}

// ToBytes @0x188ada0 保留 Go map 遍历顺序，不添加排序或末尾换行。
func (m L4jWYITSyXGq) ToBytes() []byte {
	var lines []string
	for k, v := range m {
		lines = append(lines, k+"="+v)
	}
	return []byte(strings.Join(lines, "\n"))
}
