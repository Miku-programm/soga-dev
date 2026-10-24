/*
MIT License

Copyright (c) 2016-2017 Daniel Fu

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

// 依据原 ELF（SHA256 6660dcfb…）IDA/RTTI 恢复；并非作者原文件。
// 以 MIT 许可 smux v1.5.24 为可比对基线，原 fork 的批量发送、锁范围、
// 版本协商、异步窗口更新及缓冲池差异均按机器码恢复，不能用上游包直接替代。
// RTTI 显示前缀 ji2LadBZQZW6，uncommon package path 为 I4AWsAkBLRP。
package realityconn

import (
	"errors"
	"fmt"
	"io"
	"math"
	"time"
)

type BWpFHH3x struct {
	RTNOROJC5 int

	DU0iTVAg bool

	PGgIC2 time.Duration

	PK9EpD time.Duration

	XsqVEjw4 int

	K3nkPQuOI int

	CUuHLYfWnC8q int
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func defaultConfigRecovered() *BWpFHH3x {
	return &BWpFHH3x{
		RTNOROJC5:    1,
		PGgIC2:       10 * time.Second,
		PK9EpD:       30 * time.Second,
		XsqVEjw4:     32768,
		K3nkPQuOI:    4194304,
		CUuHLYfWnC8q: 65536,
	}
}

// 原函数 0x118bee0。
func CaHL4gb(wpUkUgrJJHwN *BWpFHH3x) error {
	if !(wpUkUgrJJHwN.RTNOROJC5 == 1 || wpUkUgrJJHwN.RTNOROJC5 == 2) {
		return errors.New("unsupported protocol version")
	}
	if !wpUkUgrJJHwN.DU0iTVAg {
		if wpUkUgrJJHwN.PGgIC2 == 0 {
			return errors.New("keep-alive interval must be positive")
		}
		if wpUkUgrJJHwN.PK9EpD < wpUkUgrJJHwN.PGgIC2 {
			return fmt.Errorf("keep-alive timeout must be larger than keep-alive interval")
		}
	}
	if wpUkUgrJJHwN.XsqVEjw4 <= 0 {
		return errors.New("max frame size must be positive")
	}
	if wpUkUgrJJHwN.XsqVEjw4 > 65535 {
		return errors.New("max frame size must not be larger than 65535")
	}
	if wpUkUgrJJHwN.K3nkPQuOI <= 0 {
		return errors.New("max receive buffer must be positive")
	}
	if wpUkUgrJJHwN.CUuHLYfWnC8q <= 0 {
		return errors.New("max stream buffer must be positive")
	}
	if wpUkUgrJJHwN.CUuHLYfWnC8q > wpUkUgrJJHwN.K3nkPQuOI {
		return errors.New("max stream buffer must not be larger than max receive buffer")
	}
	if wpUkUgrJJHwN.CUuHLYfWnC8q > math.MaxInt32 {
		return errors.New("max stream buffer cannot be larger than 2147483647")
	}
	return nil
}

// 原函数 0x118c560。
func CXFDjl_TYx8(mlu21WY6 io.ReadWriteCloser, wpUkUgrJJHwN *BWpFHH3x) (*DZOCfwWJsy, error) {
	if wpUkUgrJJHwN == nil {
		wpUkUgrJJHwN = defaultConfigRecovered()
	}
	if err := CaHL4gb(wpUkUgrJJHwN); err != nil {
		return nil, err
	}
	copyConfig := *wpUkUgrJJHwN
	return yfgupp(&copyConfig, mlu21WY6, false), nil
}
