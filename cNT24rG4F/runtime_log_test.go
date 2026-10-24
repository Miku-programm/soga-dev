package cNT24rG4F

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"recovered.local/soga/muc279Mpa6S4"
)

func TestRuntimeLogPathAndSamePathLevel(t *testing.T) {
	oldConfig, oldNode, oldPath := configAt509B6A0, nodeConfigurationAt50A2AE0, logPathAt5095AC0.Load()
	t.Cleanup(func() {
		configAt509B6A0 = oldConfig
		nodeConfigurationAt50A2AE0 = oldNode
		logPathAt5095AC0.Store(oldPath)
		muc279Mpa6S4.LQ3CVLXgCaCF(0, "")
	})
	dir := filepath.Join(t.TempDir(), "new", "logs")
	configAt509B6A0.HSBWD7 = dir
	configAt509B6A0.NaavbD4r35 = "warn"
	nodeConfigurationAt50A2AE0.MnXMlyI = "Shadowsocks"
	want := dir + "/soga-ss-" + time.Now().Format("2006-01-02") + ".log"
	if got := _RTOfcJE0BpE(); got != want {
		t.Fatalf("%q != %q", got, want)
	}
	logPathAt5095AC0.Store(nil)
	W2uyDtOBN2()
	if got := logPathAt5095AC0.Load(); got == nil || *got != want {
		t.Fatal(got)
	}
	muc279Mpa6S4.V8GN5Stk(0, "hidden-info")
	muc279Mpa6S4.V8GN5Stk(4, "first-warning")
	configAt509B6A0.NaavbD4r35 = "error"
	previous := logPathAt5095AC0.Load()
	W2uyDtOBN2()
	if logPathAt5095AC0.Load() != previous {
		t.Fatal("same path unnecessarily replaced")
	}
	muc279Mpa6S4.V8GN5Stk(4, "still-warning")
	data, err := os.ReadFile(want)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "hidden-info") || !strings.Contains(string(data), "first-warning") || !strings.Contains(string(data), "still-warning") {
		t.Fatalf("log level/path behavior: %s", data)
	}
	// 目录创建失败时保存空路径并回到标准输出。
	configAt509B6A0.HSBWD7 = filepath.Join(want, "impossible")
	W2uyDtOBN2()
	if got := logPathAt5095AC0.Load(); got == nil || *got != "" {
		t.Fatal("mkdir error did not fall back")
	}
}
