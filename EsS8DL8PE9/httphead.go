/*
The MIT License (MIT)

Copyright (c) 2017 Sergey Kamardin

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
// 根据原 ELF/IDA 和 gobwas/httphead v0.1.0 逐函数对照恢复；非作者原始文件。
// 存活类型保留 RTTI 名称；未存活的值类型及内联函数保留明确对照名称。
package EsS8DL8PE9

import (
	"bytes"
)

func PgwXhS7T(data []byte, it func([]byte) bool) bool {
	lexer := &K1cBAl{data: data}

	var ok bool
	for lexer.Next() {
		switch lexer.Type() {
		case ItemToken:
			ok = true
			if !it(lexer.Bytes()) {
				return true
			}
		case ItemSeparator:
			if !isComma(lexer.Bytes()) {
				return false
			}
		default:
			return false
		}
	}

	return ok && !lexer.err
}

type SelectFlag byte

const (
	SelectCopy SelectFlag = 1 << iota

	SelectUnique
)

type ExK8bKI2Xox struct {
	Check func(BIMZTmbYHb) bool

	Flags SelectFlag

	Alloc func(n int) []byte
}

func (s ExK8bKI2Xox) Select(data []byte, options []BIMZTmbYHb) ([]BIMZTmbYHb, bool) {
	var current BIMZTmbYHb
	var has bool
	index := -1

	alloc := s.Alloc
	if alloc == nil {
		alloc = b75Ovh9a
	}
	check := s.Check
	if check == nil {
		check = hkFfEI2qM
	}

	ok := LU7plsn8k(data, func(idx int, name, attr, val []byte) Control {
		if idx != index {
			if has && check(current) {
				if s.Flags&SelectCopy != 0 {
					current = current.Copy(alloc(current.Size()))
				}
				options = append(options, current)
				has = false
			}
			if s.Flags&SelectUnique != 0 {
				for i := len(options) - 1; i >= 0; i-- {
					if bytes.Equal(options[i].KXa0Kt2I, name) {
						return ControlSkip
					}
				}
			}
			index = idx
			current = BIMZTmbYHb{KXa0Kt2I: name}
			has = true
		}
		if attr != nil {
			current.Jd2j6aD7.Set(attr, val)
		}

		return ControlContinue
	})
	if has && check(current) {
		if s.Flags&SelectCopy != 0 {
			current = current.Copy(alloc(current.Size()))
		}
		options = append(options, current)
	}

	return options, ok
}

func b75Ovh9a(n int) []byte     { return make([]byte, n) }
func hkFfEI2qM(BIMZTmbYHb) bool { return true }

type Control byte

const (
	ControlContinue Control = iota

	ControlBreak

	ControlSkip
)

func LU7plsn8k(data []byte, it func(index int, option, attribute, value []byte) Control) bool {
	lexer := &K1cBAl{data: data}

	var ok bool
	var state int
	const (
		stateKey = iota
		stateParamBeforeName
		stateParamName
		stateParamBeforeValue
		stateParamValue
	)

	var (
		index             int
		key, param, value []byte
		mustCall          bool
	)
	for lexer.Next() {
		var (
			call      bool
			growIndex int
		)

		t := lexer.Type()
		v := lexer.Bytes()

		switch t {
		case ItemToken:
			switch state {
			case stateKey, stateParamBeforeName:
				key = v
				state = stateParamBeforeName
				mustCall = true
			case stateParamName:
				param = v
				state = stateParamBeforeValue
				mustCall = true
			case stateParamValue:
				value = v
				state = stateParamBeforeName
				call = true
			default:
				return false
			}

		case ItemString:
			if state != stateParamValue {
				return false
			}
			value = v
			state = stateParamBeforeName
			call = true

		case ItemSeparator:
			switch {
			case isComma(v) && state == stateKey:

			case isComma(v) && state == stateParamBeforeName:
				state = stateKey

				call = mustCall
				if !call {

					index++
				} else {

					growIndex = 1
				}

			case isComma(v) && state == stateParamBeforeValue:
				state = stateKey
				growIndex = 1
				call = true

			case isSemicolon(v) && state == stateParamBeforeName:
				state = stateParamName

			case isSemicolon(v) && state == stateParamBeforeValue:
				state = stateParamName
				call = true

			case isEquality(v) && state == stateParamBeforeValue:
				state = stateParamValue

			default:
				return false
			}

		default:
			return false
		}

		if call {
			switch it(index, key, param, value) {
			case ControlBreak:

				return true

			case ControlSkip:

				state = stateKey
				lexer.SkipEscaped(',')

			case ControlContinue:

			default:
				panic("unexpected control value")
			}
			ok = true
			param = nil
			value = nil
			mustCall = false
			index += growIndex
		}
	}
	if mustCall {
		ok = true
		it(index, key, param, value)
	}

	return ok && !lexer.err
}

func isComma(b []byte) bool {
	return len(b) == 1 && b[0] == ','
}
func isSemicolon(b []byte) bool {
	return len(b) == 1 && b[0] == ';'
}
func isEquality(b []byte) bool {
	return len(b) == 1 && b[0] == '='
}
