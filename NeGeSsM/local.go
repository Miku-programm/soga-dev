package NeGeSsM

import (
	"bufio"
	"os"
	"path/filepath"
	C "recovered.local/soga/cNT24rG4F"
	L "recovered.local/soga/muc279Mpa6S4"
	"strconv"
	"strings"
	"time"
)

// ajkzPuuQ: 0x186dd40。原样输出八个竖线分隔字段。
func (s *cToP4Y8) ajkzPuuQ(b *strings.Builder, r *cUwNbDJ6opWy) {
	if !s._3tRkKvkg(r.pv0C8wVKLJh) {
		return
	}
	b.Reset()
	b.WriteString(r.pv0C8wVKLJh.Format("2006-01-02 15:04:05"))
	b.WriteByte('|')
	b.WriteString(r.uUbcC0nMK1Z9)
	b.WriteByte('|')
	b.WriteString(strconv.Itoa(r.iJYI5hR))
	b.WriteByte('|')
	b.WriteString(strconv.FormatInt(r.zBzHxRx4qP0, 10))
	b.WriteByte('|')
	b.WriteString(v4V_xxf(r.c7a4q7l7VIYD))
	b.WriteByte('|')
	b.WriteString(v4V_xxf(r.mv3eQCXbV))
	b.WriteByte('|')
	b.WriteString(v4V_xxf(r.vz3eAFzt))
	b.WriteByte('|')
	b.WriteString(strconv.Itoa(r.g8o8V93EQtw))
	b.WriteByte('\n')
	if _, err := s.wBou6BBU0y2S.WriteString(b.String()); err != nil {
		L.V8GN5Stk(4, "write access log failed: %v", err)
	}
}

// _3tRkKvkg: 0x186ea40。目录或小时改变才轮换；打开失败时保留空状态。
func (s *cToP4Y8) _3tRkKvkg(t time.Time) bool {
	_, dir, _ := C.AccessLogRetentionConfiguration()
	if dir == "" {
		s.t_DddcBjF7()
		return false
	}
	name := C.O__f2DDogBIK(t)
	if s.uzBEsO7m != nil && name == s.wvBqLYlW && dir == s.kZamhKNX {
		return true
	}
	s.t_DddcBjF7()
	if err := os.MkdirAll(dir, 0755); err != nil {
		L.V8GN5Stk(8, "make access log dir <%v> failed: %v", dir, err)
		return false
	}
	path := filepath.Join(dir, name)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		L.V8GN5Stk(8, "open access log file <%v> failed: %v", path, err)
		return false
	}
	s.uzBEsO7m = f
	s.wBou6BBU0y2S = bufio.NewWriterSize(f, 65536)
	s.wvBqLYlW = name
	s.kZamhKNX = dir
	return true
}

// eyAgzSwv7d: 0x186ef20。
func (s *cToP4Y8) eyAgzSwv7d() {
	if s.wBou6BBU0y2S != nil {
		if err := s.wBou6BBU0y2S.Flush(); err != nil {
			L.V8GN5Stk(4, "flush access log failed: %v", err)
		}
	}
}

// t_DddcBjF7: 0x186f000。刷新、关闭并清零完整状态，忽略 Close 错误。
func (s *cToP4Y8) t_DddcBjF7() {
	s.eyAgzSwv7d()
	if s.uzBEsO7m != nil {
		_ = s.uzBEsO7m.Close()
	}
	*s = cToP4Y8{}
}
