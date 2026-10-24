//  soga-v2.17.0 开源版
// 字符串来自该函数原调用参数的受限静态仿真，未执行原程序。
package update

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	G6 "recovered.local/soga/net/relay"
	C "recovered.local/soga/config"
	logx "recovered.local/soga/util/logx"
)

// 原 RTTI 类型名前缀为 j48VWvOD5LA9，uncommon_package_path 为 I46ita；
// 使用实际函数所属包表达该类型，不伪称恢复了作者原目录。
type e5s7yjcoNuB6 struct {
	UlWjg9iuG string `json:"tag_name"`
	KecFp709H string `json:"published_at"`
}

// PublishTime 对应 0x1f0b6a0，解析失败返回零时间。
func (r *e5s7yjcoNuB6) PublishTime() time.Time {
	v, e := time.Parse(time.RFC3339, r.KecFp709H)
	if e != nil {
		return time.Time{}
	}
	return v
}

// m_33LUULr2 对应 0x1f0b720。只重试 HTTP/读取错误，最多三次；
// HTTP 状态本身不判失败。原 json/v2 调用使用 DefaultOptionsV1。
func m_33LUULr2() (string, time.Time, error) {
	var data []byte
	var err error
	for i := 0; i < 3; i++ {
		_, data, _, err = G6.IlWBbp1xdPw("https://api.github.com/repos/vaxilu/soga/releases/latest", nil, nil)
		if err == nil {
			break
		}
	}
	if data == nil {
		return "", time.Time{}, err
	}
	var release e5s7yjcoNuB6
	if err = json.Unmarshal(data, &release); err != nil {
		return "", time.Time{}, err
	}
	return release.UlWjg9iuG, release.PublishTime(), nil
}

// bYg0ELm 对应 0x1f0b8a0。linux/amd64 是此样本实际编译常量。
func bYg0ELm(version string) (string, error) {
	address := fmt.Sprintf("https://github.com/vaxilu/soga/releases/download/%s/soga-%v-%v.tar.gz", version, "linux", "amd64")
	response, err := G6.RequestFromRecoveredCall(address, "GET", nil, nil, nil, "")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	name := "soga.tar.gz"
	_ = os.Remove(name)
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()
	buffer := make([]byte, 8192)
	for {
		n, readErr := response.Body.Read(buffer)
		if n > 0 {
			// 原 0x1f0bbb1 的 Write 返回错误覆盖前面的 Read 错误，不能改成 io.Copy。
			_, readErr = file.Write(buffer[:n])
			if readErr != nil {
				logx.TRTbrquD(4, "write file err:", readErr)
				_ = file.Close()
				_ = os.Remove(name)
				return "", readErr
			}
		}
		if readErr == io.EOF {
			return name, nil
		}
		if readErr != nil {
			_ = file.Close()
			_ = os.Remove(name)
			return "", readErr
		}
	}
}

// pUbOA95m8Gea 对应 0x1f0bec0。按文件名末尾 '/' 分目录，
// 不使用 Typeflag/Mode 判定，也不额外改写原路径处理。
func pUbOA95m8Gea(archive, destination string) error {
	file, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer file.Close()
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()
	reader := tar.NewReader(gzipReader)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		target := path.Join(destination, header.Name)
		if strings.HasSuffix(header.Name, "/") {
			if err = os.MkdirAll(target, 0777); err != nil {
				return err
			}
		} else {
			output, err := kT0645CUQ9PS(target)
			if err != nil {
				return err
			}
			_, err = io.Copy(output, reader)
			_ = output.Close()
			if err != nil && err != io.EOF {
				return err
			}
		}
	}
}

// kT0645CUQ9PS 对应 0x1f0c440。原代码把字节 LastIndex 用于 rune 切片，
// 原无 '/' 时会 panic，非 ASCII 路径也可能出现偏移差异；此处保留实际算法。
func kT0645CUQ9PS(name string) (*os.File, error) {
	runes := []rune(name)
	directory := string(runes[:strings.LastIndex(name, "/")])
	if err := os.MkdirAll(directory, 0777); err != nil {
		return nil, err
	}
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

// g0x1jn 对应 0x1f0c540 及 func1@0x1f17220；顺序和错误级别按原码保留。
func g0x1jn(version string) {
	archive, err := bYg0ELm(version)
	if err != nil {
		logx.TRTbrquD(0, "download soga file failed:", err)
		return
	}
	logx.TRTbrquD(0, "download soga success")
	destination := path.Join(os.TempDir(), "temp")
	_ = os.RemoveAll(destination)
	defer func() {
		err = os.Remove(archive)
		if err != nil {
			logx.V8GN5Stk(4, "remove <%v> failed: %v", archive, err)
		}
		err = os.RemoveAll(destination)
		if err != nil {
			logx.V8GN5Stk(4, "remove <%v> failed: %v", destination, err)
		}
	}()
	err = pUbOA95m8Gea(archive, destination)
	if err != nil {
		logx.TRTbrquD(4, "decompress soga err:", err)
		return
	}
	err = os.RemoveAll("/usr/local/soga/")
	if err != nil {
		logx.TRTbrquD(8, "remove old soga directory failed:", err)
		return
	}
	err = os.Rename(path.Join(destination, "soga"), "/usr/local/soga")
	if err != nil {
		logx.TRTbrquD(8, "move soga directory to '/usr/local/soga' failed:", err)
		return
	}
	_ = os.Chmod("/usr/local/soga/soga", 0777)
	if C.VSHlrRuK7_() {
		logx.TRTbrquD(0, "update soga success, please manual restart docker")
		return
	}
	logx.TRTbrquD(0, "update soga success, restart now")
	err = exec.Command("soga", "restart").Run()
	if err != nil {
		logx.TRTbrquD(8, "soga restart failed:", err)
	}
}

// ZAWxLZe5zfNN 对应 0x1f0cd40 / func1@0x1f17480。
// 首次检查在 24 小时后；取消上下文停止 ticker，不立即请求版本。
func ZAWxLZe5zfNN(ctx context.Context) {
	version := C.StartupVersion()
	if strings.Contains(version, "dev") || strings.Contains(version, "pre") {
		logx.TRTbrquD(4, "dev or pre version can not auto update")
		return
	}
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				current := C.StartupVersion()
				latest, published, err := m_33LUULr2()
				if err != nil {
					continue
				}
				logx.V8GN5Stk(0, "latest version: %v, current version: %v", latest, current)
				if latest != current && time.Now().Sub(published) >= 72*time.Hour {
					logx.V8GN5Stk(0, "start update soga version to v%v", latest)
					g0x1jn(latest)
				}
			}
		}
	}()
}
