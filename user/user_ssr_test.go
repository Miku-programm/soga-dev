package user

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

func TestSSRDeviceAccounting(t *testing.T) {
	var layout LJJkaD
	if unsafe.Sizeof(layout) != 136 || unsafe.Offsetof(layout.kHH6tGdN) != 88 || unsafe.Offsetof(layout.ya8WV3ZP11A) != 120 {
		t.Fatal("SSR layout")
	}
	p := &types.Aga91y{WZ89pTAOrHK4: 42, Xmrna8: 2}
	u := &LJJkaD{AhSFeSua: cLLKJ_A31(p, nil), kHH6tGdN: make(map[int64]int)}
	if !u.AddDevice(10) || !u.AddDevice(10) || !u.AddDevice(20) || u.AddDevice(30) || u.wyY4xeau != 2 {
		t.Fatal("device reference/limit")
	}
	p.Xmrna8 = 1
	u.UpdateUser(p)
	if !u.AddDevice(10) || u.AddDevice(30) {
		t.Fatal("existing device remains allowed above limit")
	}
	u.DelDevice(999)
	u.DelDevice(10)
	u.DelDevice(10)
	if u.wyY4xeau != 2 {
		t.Fatal("premature unique decrement")
	}
	u.DelDevice(10)
	u.DelDevice(20)
	if u.wyY4xeau != 0 || len(u.kHH6tGdN) != 0 {
		t.Fatal("device count not cleared")
	}
	p.Xmrna8 = 0
	u.UpdateUser(p)
	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if !u.AddDevice(int64(i % 7)) {
				t.Error("unlimited rejected")
			}
			u.DelDevice(int64(i % 7))
		}(i)
	}
	wg.Wait()
	if u.wyY4xeau != 0 || len(u.kHH6tGdN) != 0 {
		t.Fatal("concurrent device count")
	}
	u.kHH6tGdN = nil
	func() {
		defer func() {
			if recover() == nil {
				t.Error("nil map should panic")
			}
		}()
		u.AddDevice(1)
	}()
	if !u.jS2vrvcY.TryLock() {
		t.Fatal("panic retained mutex")
	}
	u.jS2vrvcY.Unlock()
}
func TestCredentialConfigPipeline(t *testing.T) {
	// Isolate original configuration globals and its command-line registry in a
	// subprocess. Only the recovered source test executable runs, never soga.
	if os.Getenv("SOGA_CREDENTIAL_CONFIG_CHILD") != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestCredentialConfigPipeline$", "-test.count=1")
		cmd.Env = append(os.Environ(), "SOGA_CREDENTIAL_CONFIG_CHILD=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("configuration integration: %v\n%s", err, out)
		}
		return
	}
	path := filepath.Join(t.TempDir(), "soga.conf")
	config.MkwNWciPHS(path)
	load := func(panel, template, mode string) {
		t.Helper()
		cfg := fmt.Sprintf("type=%s\nmu_suffix=example.test\nmu_regex=%s\nvmess_auth_mode=%s\naccess_log_enable=false\n", panel, template, mode)
		if err := os.WriteFile(path, []byte(cfg), 0600); err != nil {
			t.Fatal(err)
		}
		if err := config.TvBECnD9fM6(); err != nil {
			t.Fatal(err)
		}
	}
	for _, panel := range []string{"sspanel-uim", "soga-v1", "proxypanel", "whmcs"} {
		load(panel, "%5m%id.%suffix", "auto")
		c := &types.QcqpSL3i6ri{YrjVl_2g_S: 12345, YyAgnMatmj: "secret", Z_taAb: "aes-256-cfb", H_y2xZht7S8: "auth_aes128_md5", ZZIDsJ: "tls1.2_ticket_auth", BskzckkE7qI: 1}
		p := &types.Aga91y{WZ89pTAOrHK4: 42, JAfXWwzYrW: c}
		u, err := PXHoe5(p, nil)
		if err != nil {
			t.Fatal(err)
		}
		wantUID := int64(42)
		if panel == "proxypanel" || panel == "whmcs" {
			wantUID = 12345
		}
		if u.GetUID() != wantUID || u.GetPassword() != "secret" || u.GetMethod() != "aes-256-cfb" || u.GetProtocol() != "auth_aes128_md5" || u.GetObfs() != "tls1.2_ticket_auth" || !u.GetSinglePortType().IsObfs() {
			t.Fatal("SSR constructor", panel)
		}
		if u.GenUserMd5() != "bd9dbe2600dc885b68d963ac8fa63933" {
			t.Fatal(u.GenUserMd5())
		}
		host, err := u.GetSSRHost()
		if err != nil || host != "bd9db42.example.test" {
			t.Fatal(host, err)
		}
		ctx := u.GetCtx()
		c.BskzckkE7qI = 2
		if changed, old := u.UpdatePanelUser(p); changed || old != "secret" || u.GetCtx() != ctx || !u.GetSinglePortType().IsProtocol() {
			t.Fatal("single-port-only update")
		}
		load(panel, "%99m-%id-%id.%suffix.%suffix", "auto")
		if host, _ := u.GetSSRHost(); host != "bd9db42.example.test" {
			t.Fatal("host cache invalidated by config alone")
		}
		c.YyAgnMatmj = "new"
		if changed, old := u.UpdatePanelUser(p); !changed || old != "secret" || u.GetCtx() == ctx || u.ya8WV3ZP11A != "" {
			t.Fatal("password update")
		}
		host, err = u.GetSSRHost()
		if err != nil || host != u.GenUserMd5()+"-42-%id.example.test.%suffix" {
			t.Fatal("first-only replacements", host, err)
		}
		p.JAfXWwzYrW = &types.JxTGUar{}
		p.Xmrna8 = 4
		ctx = u.GetCtx()
		if changed, old := u.UpdatePanelUser(p); changed || old != "new" || u.GetCtx() != ctx || u.GetConnLimit() != 4 {
			t.Fatal("wrong payload order")
		}
	}
	load("unknown", "%0m%id.%suffix", "auto")
	c := &types.QcqpSL3i6ri{YyAgnMatmj: "secret"}
	p := &types.Aga91y{WZ89pTAOrHK4: 42, JAfXWwzYrW: c}
	func() {
		defer func() {
			if got := recover(); got != "can not get ssr uid, not support panel: unknown" {
				t.Errorf("unsupported panel panic: %v", got)
			}
		}()
		PXHoe5(p, nil)
	}()
	load("sspanel-uim", "%0m%id.%suffix", "auto")
	u, _ := PXHoe5(p, nil)
	if h, e := u.GetSSRHost(); e != nil || h != "42.example.test" {
		t.Fatal(h, e)
	}
	for _, template := range []string{"no MD5 placeholder", "%999999999999999999999999999999999m"} {
		load("sspanel-uim", template, "auto")
		u.ya8WV3ZP11A = ""
		if strings.HasPrefix(template, "no") {
			func() {
				defer func() {
					if recover() == nil {
						t.Error("missing capture must panic")
					}
				}()
				u.GetSSRHost()
			}()
		} else if h, e := u.GetSSRHost(); h != "" || e == nil {
			t.Fatal("overflow capture", h, e)
		}
	}
	load("sspanel-uim", "", "auto")
	u.ya8WV3ZP11A = ""
	if h, e := u.GetSSRHost(); h != "" || e.Error() != "must config mu_suffix and mu_regex" {
		t.Fatal(h, e)
	}
	load("sspanel-uim", "%5m%id.%suffix", "aead")
	vm := &types.Yn8BJoL{SfiEGpJo: "01234567-89ab-cdef-0123-456789abcdef", FvQ2ltKOM7: 64}
	v, err := EOkO09L(&types.Aga91y{WZ89pTAOrHK4: 2, JAfXWwzYrW: vm}, nil)
	if err != nil || vm.FvQ2ltKOM7 != 0 || len(v.GetUserIds()) != 1 || v.GetBlock() == nil {
		t.Fatal("forced AEAD must mutate caller credential", err)
	}
}
