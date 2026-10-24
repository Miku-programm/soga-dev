// 根据 soga-v2.17.0 的 IDA/RTTI 及受限字符串解码恢复；非作者源码原文。
package a3PdV_bmpRUK

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/muc279Mpa6S4"
)

// 原 init@0x149fb80 设置字符串全局 0x5096d90；互斥锁在 0x50c6d98。
var storageDirectoryAt5096D90 = "/etc/soga"
var selfCertificateMutexAt50C6D98 sync.Mutex

// oWCp7YUAFId 原 RTTI 大小 64 字节；两个文件的修改时间和大小。
type oWCp7YUAFId struct {
	dGHfAx      time.Time
	fErFXoSf7Sv int64
	bcInmv8PX2r time.Time
	eMARy7uR    int64
}

// J6c4NB_R 对应 0x153f220；这里使用固定目录，不读取自签证书缓存目录变量。
func J6c4NB_R(nodeID int) (string, string) {
	return fmt.Sprintf("/etc/soga/panel_%v.crt", nodeID), fmt.Sprintf("/etc/soga/panel_%v.key", nodeID)
}

// sur_b39JM6l 对应 0x14ab880。只解析第一个 PEM 块，不检查块名。
func sur_b39JM6l(path string) (*time.Time, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cert file not exist: %v", path)
	}
	block, _ := pem.Decode(contents)
	if block == nil {
		muc279Mpa6S4.TRTbrquD(slog.LevelError, "cert file invalid:", path)
		return nil, fmt.Errorf("cert file invalid: %v", path)
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("cert file invalid: %v", path)
	}
	return &cert.NotAfter, nil
}

// RUKs8pSwE4 对应 0x153fd00。不检查信任链、用途或域名。
func RUKs8pSwE4(certPath, keyPath string) error {
	if !CAodbhZf.VSCJLy2No(certPath) || !CAodbhZf.VSCJLy2No(keyPath) {
		return errors.New("cached panel cert not found")
	}
	if _, err := tls.LoadX509KeyPair(certPath, keyPath); err != nil {
		return fmt.Errorf("cached panel cert invalid: %v", err)
	}
	expires, err := sur_b39JM6l(certPath)
	if err != nil {
		return err
	}
	if time.Now().After(*expires) {
		return fmt.Errorf("cached panel cert expired at %v", expires)
	}
	return nil
}

// EEGj3Ra 对应 0x153f380。缓存失效时生成十年自签证书；锁涵盖生成和读取。
func EEGj3Ra(domain string) (string, string, error) {
	if domain == "" || domain == "." || domain == ".." || filepath.Base(domain) != domain {
		return "", "", fmt.Errorf("invalid domain: %q", domain)
	}
	selfCertificateMutexAt50C6D98.Lock()
	defer selfCertificateMutexAt50C6D98.Unlock()
	certPath := filepath.Join(storageDirectoryAt5096D90, fmt.Sprintf("panel_self_%v.crt", domain))
	keyPath := filepath.Join(storageDirectoryAt5096D90, fmt.Sprintf("panel_self_%v.key", domain))
	if RUKs8pSwE4(certPath, keyPath) != nil {
		if err := lLTDzR(domain, certPath, keyPath, 10); err != nil {
			return "", "", err
		}
	}
	cert, err := os.ReadFile(certPath)
	if err != nil {
		return "", "", err
	}
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return "", "", err
	}
	return string(cert), string(key), nil
}

// eG4MicXvEEf 对应 0x1544d80。任一 Stat 失败均返回全零快照。
func eG4MicXvEEf(certPath, keyPath string) (oWCp7YUAFId, error) {
	cert, err := os.Stat(certPath)
	if err != nil {
		return oWCp7YUAFId{}, err
	}
	key, err := os.Stat(keyPath)
	if err != nil {
		return oWCp7YUAFId{}, err
	}
	return oWCp7YUAFId{cert.ModTime(), cert.Size(), key.ModTime(), key.Size()}, nil
}

// wLQ7ZW 对应 0x1540400。内容相同时不改变权限或修改时间。
// 使用同目录临时文件后 Rename；原函数没有 fsync，也没有证书/私钥成对事务。
func wLQ7ZW(path, content string, mode os.FileMode) error {
	if old, err := os.ReadFile(path); err == nil && string(old) == content {
		return nil
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	file, err := os.CreateTemp(dir, filepath.Base(path)+".tmp*")
	if err != nil {
		return err
	}
	_, writeErr := file.Write([]byte(content))
	closeErr := file.Close()
	if writeErr != nil {
		os.Remove(file.Name())
		return writeErr
	}
	if closeErr != nil {
		os.Remove(file.Name())
		return closeErr
	}
	if err = os.Chmod(file.Name(), mode); err == nil {
		err = os.Rename(file.Name(), path)
	}
	if err != nil {
		os.Remove(file.Name())
	}
	return err
}
