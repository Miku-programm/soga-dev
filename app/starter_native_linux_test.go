//go:build linux

package app

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	L "recovered.local/soga/panel/types"
	M "recovered.local/soga/protocol/inbound"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"reflect"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func TestNativeStarterFactoryAndDeferredReceiverRead(t *testing.T) {
	if os.Getenv("RECOVERED_NATIVE_STARTER_FIXTURE") != "1" {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestNativeStarterFactoryAndDeferredReceiverRead$", "-test.v")
		cmd.Env = append(os.Environ(), "RECOVERED_NATIVE_STARTER_FIXTURE=1")
		if out, e := cmd.CombinedOutput(); e != nil {
			t.Fatalf("native fixture: %v\n%s", e, out)
		}
		return
	}
	if unsafe.Sizeof(WrdrARwuhuha{}) != 8 || unsafe.Offsetof(WrdrARwuhuha{}.zqoWpLRIH) != 0 {
		t.Fatal("RTTI layout")
	}
	var nilStarter *WrdrARwuhuha
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if got, e := nilStarter.vDpj7g(canceled, nil); got != nil || e != context.Canceled {
		t.Fatal("cancellation must precede receiver/data reads", got, e)
	}
	writeConfig := func(protocol string) {
		path := filepath.Join(t.TempDir(), "soga.conf")
		text := "type=sspanel-uim\nserver_type=" + protocol + "\nlisten=127.0.0.1\nlog_file=false\naccess_log_enable=false\ndy_limit_enable=false\nmu_suffix=example.test\nmu_regex=%0m%id.%suffix\n"
		if e := os.WriteFile(path, []byte(text), 0600); e != nil {
			t.Fatal(e)
		}
		C.MkwNWciPHS(path)
		if e := C.TvBECnD9fM6(); e != nil {
			t.Fatal(e)
		}
	}
	data := &Q.D5gTIZ5aVg{H47AzZsYHB1N: &L.Wr5SVP{EMlXEc: 713}}
	writeConfig("unknown-native")
	if got, e := nilStarter.vDpj7g(context.Background(), data); got != nil || e == nil || e.Error() != "unknown server type: unknown-native" {
		t.Fatal("factory error must precede receiver read", got, e)
	}
	writeConfig("ssr")
	starterOnceAt50C7AA0 = sync.Once{}
	called := make(chan int, 1)
	starter := eXC63toqu(func(n int) { called <- n }).(*WrdrARwuhuha)
	core, e := starter.vDpj7g(context.Background(), data)
	if e != nil {
		t.Fatal(e)
	}
	defer core.Stop()
	if _, ok := core.(*M.Ze0fOaAko); !ok {
		t.Fatalf("actual native factory returned %T", core)
	}
	select {
	case n := <-called:
		if n != 0 {
			t.Fatal(n)
		}
	case <-time.After(time.Second):
		t.Fatal("initial callback missing")
	}
	// Once was consumed. A nil receiver now remains unobserved even when the real
	// constructor and Start succeed; evaluating callback at method entry is wrong.
	other, e := nilStarter.vDpj7g(context.Background(), data)
	if e != nil {
		t.Fatal(e)
	}
	other.Stop()
}

func TestNativeStarterStartIsAsynchronous(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := make(chan struct{}, 1)
	panel := &pullLoopPanel{fetch: func() (*Q.D5gTIZ5aVg, error) { return nil, context.Canceled }}
	panel.report = func(*Q.FGmNba4MhioY) error { started <- struct{}{}; return nil }
	// The original starts the empty report goroutine even for canceled contexts;
	// no callback field is read and no inbound listener is created here.
	var starter *WrdrARwuhuha
	starter.Start(ctx, panel)
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("asynchronous starter did not submit initial data")
	}
}

func TestNativeStarterTenOriginalRegistrations(t *testing.T) {
	expected := map[C.FSOZSH7g]_Zpu8_d{"dj": u1nzch, "ppanel": aGvHzoD3fr, "proxypanel": yyCM1ej, "soga-v1": jQGVEC1W, "sspanel-uim": joTwyrx, "v2board": lag2OyJwFxd, "v2raysocks": ssDTthclJN, "whmcs": oXJigpaEaR, "xboard": qu_iEPVMzK9e, "xiaov2board": uaxpQZ}
	if len(starterConstructorsAt5095C50) != 10 || len(starterCallbacksAt5095C58) != 10 {
		t.Fatal("original registration count")
	}
	for key, want := range expected {
		got, e := AewQql2U(key)
		if e != nil {
			t.Fatal(e)
		}
		var cb _Zpu8_d
		if key == "dj" {
			d, ok := got.(*AiB7DQuj)
			if !ok || d.WrdrARwuhuha == nil {
				t.Fatalf("dj type %T", got)
			}
			cb = d.zqoWpLRIH
		} else {
			w, ok := got.(*WrdrARwuhuha)
			if !ok {
				t.Fatalf("%s type %T", key, got)
			}
			cb = w.zqoWpLRIH
		}
		if reflect.ValueOf(cb).Pointer() != reflect.ValueOf(want).Pointer() {
			t.Fatal("callback mismatch", key)
		}
	}
	for _, key := range []C.FSOZSH7g{"", "V2BOARD", "xboard ", "v2board-vless-webapi", "unknown%v\x00"} {
		got, e := AewQql2U(key)
		if got != nil || e == nil || e.Error() != "unknown panel: "+string(key) {
			t.Fatal(key, got, e)
		}
	}
	if unsafe.Sizeof(AiB7DQuj{}) != 8 {
		t.Fatal("dj RTTI")
	}
}

type djStatusFixture struct {
	interval func() time.Duration
	status   func() error
}

func (p *djStatusFixture) GetSubmitInterval() time.Duration { return p.interval() }
func (p *djStatusFixture) SubmitStatus() error              { return p.status() }
func TestNativeDJStatusLoopOriginalOrder(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	intervals, reports := 0, 0
	p := &djStatusFixture{interval: func() time.Duration { intervals++; return time.Millisecond }, status: func() error {
		reports++
		if reports == 2 {
			cancel()
		}
		return context.Canceled
	}}
	djSubmitStatusLoop(ctx, p)
	if intervals != 1 || reports < 2 {
		t.Fatal("original interval snapshot/error continuation", intervals, reports)
	}
	ctx, cancel = context.WithCancel(context.Background())
	cancel()
	reports = 0
	p.interval = func() time.Duration { return time.Hour }
	p.status = func() error { reports++; return nil }
	djSubmitStatusLoop(ctx, p)
	if reports != 0 {
		t.Fatal("unexpected immediate status report")
	}
	p.interval = func() time.Duration { return 0 }
	defer func() {
		if recover() == nil {
			t.Fatal("nonpositive ticker interval must panic before ctx cancellation")
		}
	}()
	djSubmitStatusLoop(ctx, p)
}
