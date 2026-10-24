//go:build linux

package cert

import (
	"context"
	"encoding/pem"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"recovered.local/soga/util/rtx"
	"recovered.local/soga/config"
)

func TestPEMValidationAndPanelEarlyFailure(t *testing.T) {
	cert, key, err := func() (string, string, error) {
		d := t.TempDir()
		c, k := filepath.Join(d, "c"), filepath.Join(d, "k")
		if err := lLTDzR("test.example", c, k, 1); err != nil {
			return "", "", err
		}
		a, _ := os.ReadFile(c)
		b, _ := os.ReadFile(k)
		return string(a), string(b), nil
	}()
	if err != nil {
		t.Fatal(err)
	}
	if err = rtx.ZvmhWx(cert, key); err != nil {
		t.Fatal(err)
	}
	expires, err := rtx.GD_QvaWw(cert)
	if err != nil || !expires.After(time.Now()) {
		t.Fatal(expires, err)
	}
	block, _ := pem.Decode([]byte(cert))
	block.Type = "UNRELATED"
	expires, err = rtx.GD_QvaWw(string(pem.EncodeToMemory(block)))
	if err != nil || expires.IsZero() {
		t.Fatal("original accepts PEM block type", err)
	}
	for _, bad := range []string{"not PEM", string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: []byte("bad DER")}))} {
		expires, err = rtx.GD_QvaWw(bad)
		if err == nil || !expires.IsZero() {
			t.Fatal(expires, err)
		}
		c, k, err := Qr4n7sWM5di(1, bad, key)
		if err == nil || c != "" || k != "" || !strings.HasPrefix(err.Error(), "panel cert invalid: invalid cert/key pair: ") {
			t.Fatal(c, k, err)
		}
	}
}

func TestCertificateWatchersCancelAndRejectInvalidChanges(t *testing.T) {
	// 受测试虚拟时钟控制，涵盖每日检查、两分钟检查、文件丢失、无效变更与退出。
	// 不触发成功证书变更分支，因此测试不会向宿主进程发送重载信号。
	oldCert, oldKey := config.ManualCertificateFiles()
	defer config.SetCertificateFiles(oldCert, oldKey)
	d := t.TempDir()
	cert, key := filepath.Join(d, "cert"), filepath.Join(d, "key")
	os.WriteFile(cert, []byte("invalid cert"), 0600)
	os.WriteFile(key, []byte("invalid key"), 0600)
	config.SetCertificateFiles(cert, key)
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		D18E0adGmOX(ctx)
		Ed4asPq5(ctx, cert, key)
		synctest.Wait()
		time.Sleep(121 * time.Second)
		synctest.Wait()
		os.WriteFile(cert, []byte("changed and still invalid"), 0600)
		time.Sleep(121 * time.Second)
		synctest.Wait()
		os.Remove(cert)
		time.Sleep(25 * time.Hour)
		synctest.Wait()
		cancel()
		synctest.Wait()
	})
	// 原初始 Stat 错误直接返回，不启动守护协程。
	Ed4asPq5(context.Background(), cert, key)
}
