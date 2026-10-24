package xhttp

import (
	"crypto/rand"
	"math"
	"net/http"
	"net/url"
	B "recovered.local/soga/buf/pool"
	"strings"
)

// 原 RTTI：AHn9pfr 64 字节，LaPHbx6B50uZ 88 字节。
type AHn9pfr struct{ AnBhAR, AVHkw5m, HxzvtHynCM, SZY0xnkLskud string }
type LaPHbx6B50uZ struct {
	TldsrRa     int
	C5zKbVfj3g_ AHn9pfr
	Yxjd2mW     AdpPM6
}

// 0x126c8c0。阈值确实截断到 byte；不扩大为 int 来改变原行为。
func gEb3WsxKBjs(n int, alphabet string) (string, bool) {
	if n <= 0 || len(alphabet) == 0 {
		return "", false
	}
	limit := byte(256 - 256%len(alphabet))
	out := B.PoolAt50961D8.Get(n)
	defer B.PoolAt50961D8.Put(out)
	random := B.PoolAt50961D8.Get(256)
	defer B.PoolAt50961D8.Put(random)
	filled := 0
	for filled < n {
		if _, e := rand.Read(random); e != nil {
			return "", false
		}
		for _, b := range random {
			if b < limit {
				out[filled] = alphabet[int(b)%len(alphabet)]
				filled++
				if filled == n {
					break
				}
			}
		}
	}
	return string(out), true
}

// 0x126cbc0；字母表经原混淆解码器受限执行取得。
func A9pjdDAvHO(n int) string {
	value, ok := gEb3WsxKBjs(max(1, int(math.Ceil(float64(n)/0.8))), "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	if !ok {
		return ""
	}
	suffix := byte('X')
	for i := 0; i < 150; i++ {
		bits := 0
		for j := 0; j < len(value); j++ {
			bits += int(paddingBitsAt3790D20[value[j]])
		}
		delta := (bits+7)/8 - n
		distance := delta
		if distance < 0 {
			distance = -distance
		}
		if distance <= 2 {
			break
		}
		if delta < 0 {
			value += string(suffix)
			if suffix == 'X' {
				suffix = 'Z'
			} else {
				suffix = 'X'
			}
		} else {
			if len(value) <= 1 {
				return value
			}
			value = value[:len(value)-1]
		}
	}
	return value
}
func D940rWda18j4(method AdpPM6, n int) string { // 0x126cd80
	if n <= 0 {
		return ""
	}
	if method == "tokenish" {
		if s := A9pjdDAvHO(n); s != "" {
			return s
		}
	}
	return strings.Repeat("X", n)
}
func VzIFuzeyXw(u *url.URL, key, value string) { // 0x126ce80
	if u == nil || key == "" || value == "" {
		return
	}
	q := u.Query()
	q.Set(key, value)
	u.RawQuery = q.Encode()
}
func (c *R0J19auWof) ApplyXPaddingToHeader(h http.Header, p LaPHbx6B50uZ) { // 0x126d020
	if h == nil {
		return
	}
	value := D940rWda18j4(p.Yxjd2mW, p.TldsrRa)
	where := p.C5zKbVfj3g_
	switch where.AnBhAR {
	case "header":
		h.Set(where.HxzvtHynCM, value)
	case "queryInHeader":
		u, e := url.Parse(where.SZY0xnkLskud)
		if e == nil && u != nil {
			u.RawQuery = where.AVHkw5m + "=" + value
			h.Set(where.HxzvtHynCM, u.String())
		}
	}
}
func (c *R0J19auWof) ApplyXPaddingToRequest(r *http.Request, p LaPHbx6B50uZ) { // 0x126d3e0
	if r == nil {
		return
	}
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	switch p.C5zKbVfj3g_.AnBhAR {
	case "header", "queryInHeader":
		c.ApplyXPaddingToHeader(r.Header, p)
		return
	}
	value := D940rWda18j4(p.Yxjd2mW, p.TldsrRa)
	switch p.C5zKbVfj3g_.AnBhAR {
	case "query":
		VzIFuzeyXw(r.URL, p.C5zKbVfj3g_.AVHkw5m, value)
	case "cookie":
		if p.C5zKbVfj3g_.AVHkw5m != "" && value != "" {
			r.AddCookie(&http.Cookie{Name: p.C5zKbVfj3g_.AVHkw5m, Value: value, Path: "/"})
		}
	}
}
func (c *R0J19auWof) ApplyXPaddingToResponse(w http.ResponseWriter, p LaPHbx6B50uZ) { // 0x126d700
	switch p.C5zKbVfj3g_.AnBhAR {
	case "header", "queryInHeader":
		c.ApplyXPaddingToHeader(w.Header(), p)
		return
	}
	value := D940rWda18j4(p.Yxjd2mW, p.TldsrRa)
	if p.C5zKbVfj3g_.AnBhAR == "cookie" && p.C5zKbVfj3g_.AVHkw5m != "" && value != "" {
		http.SetCookie(w, &http.Cookie{Name: p.C5zKbVfj3g_.AVHkw5m, Value: value, Path: "/"})
	}
}
func (c *R0J19auWof) ExtractXPaddingFromRequest(r *http.Request, obfuscated bool) (string, string) { // 0x126d980
	if r == nil {
		return "", ""
	}
	if !obfuscated {
		referer := r.Header.Get("Referer")
		if referer == "" {
			return r.URL.Query().Get("x_padding"), "query, key=x_padding"
		}
		if u, e := url.Parse(referer); e == nil {
			return u.Query().Get("x_padding"), "queryInHeader=Referer, key=x_padding"
		}
	}
	key, header := c.H6PfvFysST, c.LjJbxC6vDL
	if cookie, e := r.Cookie(key); e == nil && cookie != nil && cookie.Value != "" {
		return cookie.Value, "cookie, key=" + key
	}
	value := r.Header.Get(header)
	if value != "" {
		if c.V0qykwP == "header" {
			return value, "header=" + header
		}
		if u, e := url.Parse(value); e == nil {
			return u.Query().Get(key), "queryInHeader=" + header + ", key=" + key
		}
	}
	if value = r.URL.Query().Get(key); value != "" {
		return value, "query, key=" + key
	}
	return "", ""
}
