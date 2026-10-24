/*
The MIT License (MIT)

Copyright (c) 2017-2021 Sergey Kamardin <gobwas@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
// 原 ELF 的 0x11ab0a0、0x11b78e0 及调用链与 gobwas/ws v1.4.0 对照恢复。
// HTTP 仅在已核对的服务/劫持调用域使用 Go 标准类型；不宣称整包 fork 完全相同。
package FgcbiePIJLm

import (
	"bytes"

	httphead "recovered.local/soga/EsS8DL8PE9"
)

func uEdadz_EwJwD(header, token string) (has bool) {
	return btsHasToken(strToBytes(header), strToBytes(token))
}

func btsHasToken(header, token []byte) (has bool) {
	httphead.PgwXhS7T(header, func(v []byte) bool {
		has = bytes.EqualFold(v, token)
		return !has
	})
	return has
}

const (
	toLower  = 'a' - 'A'
	toUpper  = ^byte(toLower)
	toLower8 = uint64(toLower) |
		uint64(toLower)<<8 |
		uint64(toLower)<<16 |
		uint64(toLower)<<24 |
		uint64(toLower)<<32 |
		uint64(toLower)<<40 |
		uint64(toLower)<<48 |
		uint64(toLower)<<56
)
