// 本文件是原样本机器码的 Go 源码表示，不是作者源文件原文。
// 原包元数据名 mlbRsO44；RTTI 显示包名 wqmlKj。地址及解码证据见
// work/ida-20261023/hysteria-wire/strings-decoded.json。
package mlbRsO44

import (
	"math/rand"
	"net/http"
	"strconv"
)

// bMiTafgCoM.String 的两个机器字参数对应上下界。字段名未在保留 RTTI 中出现，
// Min、Max 是恢复命名；原方法地址 0x1c28720。
type bMiTafgCoM struct{ Min, Max int }

var authPaddingAt3789480 = bMiTafgCoM{256, 2048}
var responsePaddingAt3789490 = bMiTafgCoM{128, 1024}

func (p bMiTafgCoM) String() string {
	b := make([]byte, rand.Intn(p.Max-p.Min)+p.Min)
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// GZDAjvxfC 对应 0x1c254c0。返回值按寄存器中的 string、uint64 展开；
// 原匿名请求容器没有保留可命名类型。保留 ParseUint 出错时的实际数值返回。
func GZDAjvxfC(h http.Header) (string, uint64) {
	rx, _ := strconv.ParseUint(h.Get("Hysteria-CC-RX"), 10, 64)
	return h.Get("Hysteria-Auth"), rx
}

// CIvpFjH270ca 对应 0x1c255c0。
func CIvpFjH270ca(h http.Header, udp bool, rx uint64, auto bool) {
	h.Set("Hysteria-UDP", strconv.FormatBool(udp))
	if auto {
		h.Set("Hysteria-CC-RX", "auto")
	} else {
		h.Set("Hysteria-CC-RX", strconv.FormatUint(rx, 10))
	}
	h.Set("Hysteria-Padding", authPaddingAt3789480.String())
}
