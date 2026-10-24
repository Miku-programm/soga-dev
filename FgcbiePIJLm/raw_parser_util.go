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
// 原 ELF 0x11b8760 的 raw Upgrade 及解析链逐分支对照 gobwas/ws v1.4.0 恢复。
// 只恢复机器码中已证实的函数；未存活的局部类型和字段使用明确恢复名。
package FgcbiePIJLm

import (
	"bufio"
	"bytes"
	"fmt"
)

func fm9r93qPp(bts []byte) (ret int, err error) {

	var n int
	if n = len(bts); n < 1 {
		return 0, fmt.Errorf("converting empty bytes to int")
	}
	for i := 0; i < n; i++ {
		if bts[i]&0xf0 != 0x30 {
			return 0, fmt.Errorf("%s is not a numeric character", string(bts[i]))
		}
		ret += int(bts[i]&0xf) * pow(10, n-i-1)
	}
	return ret, nil
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func cU_pjSUO(bts []byte, sep byte) (b1, b2, b3 []byte) {
	a := bytes.IndexByte(bts, sep)
	b := bytes.IndexByte(bts[a+1:], sep)
	if a == -1 || b == -1 {
		return bts, nil, nil
	}
	b += a + 1
	return bts[:a], bts[a+1 : b], bts[b+1:]
}

func btrim(bts []byte) []byte {
	var i, j int
	for i = 0; i < len(bts) && (bts[i] == ' ' || bts[i] == '\t'); {
		i++
	}
	for j = len(bts); j > i && (bts[j-1] == ' ' || bts[j-1] == '\t'); {
		j--
	}
	return bts[i:j]
}

func canonicalizeHeaderKey(k []byte) {
	upper := true
	for i, c := range k {
		if upper && 'a' <= c && c <= 'z' {
			k[i] &= toUpper
		} else if !upper && 'A' <= c && c <= 'Z' {
			k[i] |= toLower
		}
		upper = c == '-'
	}
}

func aqxVoILI83(br *bufio.Reader) ([]byte, error) {
	var line []byte
	for {
		bts, err := br.ReadSlice('\n')
		if err == bufio.ErrBufferFull {

			line = append(line, bts...)
			continue
		}

		if line == nil {
			line = bts
		} else {
			line = append(line, bts...)
		}

		if err != nil {
			return line, err
		}

		n := len(line)

		if n > 1 && line[n-2] == '\r' {
			line = line[:n-2]
		} else {
			line = line[:n-1]
		}

		return line, nil
	}
}

func nonZero(a, b int) int {
	if a != 0 {
		return a
	}
	return b
}
