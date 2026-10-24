package MFFa9yT

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/q5bqCgTdxaTM"
	"recovered.local/soga/xx5fRyrFw"
	"testing"
	"unsafe"
)

// 每个用例在隔离测试进程加载真实配置，保留生产全局初始化要求。
func vlessFixtureProcess(t *testing.T) bool {
	t.Helper()
	if os.Getenv("RECOVERED_VLESS_USERS_FIXTURE") != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^"+t.Name()+"$", "-test.timeout=20s")
		cmd.Env = append(os.Environ(), "RECOVERED_VLESS_USERS_FIXTURE=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("fixture %v\n%s", err, out)
		}
		return false
	}
	conf := filepath.Join(t.TempDir(), "soga.conf")
	if err := os.WriteFile(conf, []byte("dy_limit_enable=false\naccess_log_enable=false\nlog_file=false\n"), 0600); err != nil {
		t.Fatal(err)
	}
	cNT24rG4F.MkwNWciPHS(conf)
	if err := cNT24rG4F.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	return true
}
func newVlessUsersFixture(t *testing.T) *KHonGaqbIy_6 {
	t.Helper()
	b, e := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(b.yPUbKO)
	return &KHonGaqbIy_6{JWN154qJ4: b, vU51MPl_at_: &G7A7csgrs62p.GTFAWF[uuid.UUID, *nzATvajt.A6CvRbtca]{}}
}
func panelVless(id int64, key string) *LFWTTXF7JR.Aga91y {
	return &LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: id, JAfXWwzYrW: &LFWTTXF7JR.IImIwacp{Dh5iiakU: key}}
}
func TestVlessStateLayoutAndConstruction(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	var v KHonGaqbIy_6
	var arg DV_XPJ4z0
	if unsafe.Sizeof(v) != 48 || unsafe.Offsetof(v.kX5i8U) != 8 || unsafe.Offsetof(v.vU51MPl_at_) != 32 || unsafe.Offsetof(v.i420mh4Yr) != 40 || unsafe.Sizeof(arg) != 40 || unsafe.Offsetof(arg.POpyjFWF8v) != 16 || unsafe.Offsetof(arg.HG5k0365) != 32 {
		t.Fatal("original RTTI")
	}
	v2 := newVlessUsersFixture(t)
	if v2.jc_QPN6 == nil || v2.dwfzZM == nil || v2.hS8xXFfcBr == nil || v2.qE9aU9 == nil || v2.oSsZZT4EmeG == nil || v2.cEUo2gYKI == nil {
		t.Fatal("constructor missing real state")
	}
	v2.yPUbKO()
	if v2.zb8f3cHWW8v.Err() != context.Canceled {
		t.Fatal("child cancellation")
	}
}
func TestVlessUserUpdateFailureDeletionAndCounts(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVlessUsersFixture(t)
	key1 := "11111111-1111-4111-8111-111111111111"
	key2 := "22222222-2222-4222-8222-222222222222"
	v.spcNlUL(LFWTTXF7JR.Sxcf3wu3L{panelVless(1, key1), panelVless(2, "bad")})
	first, ok := v.jc_QPN6.Get(1)
	if !ok || v.jc_QPN6.GTFAWF.Len() != 1 || v.fUxO7UYAH != 2 || userCountAt50C6DA0 != 2 {
		t.Fatal("failed creates still count")
	}
	user := first.(*nzATvajt.A6CvRbtca)
	oldCtx := user.GetCtx()
	v.spcNlUL(LFWTTXF7JR.Sxcf3wu3L{panelVless(1, key2)})
	if oldCtx.Err() != context.Canceled {
		t.Fatal("UUID update must cancel old session")
	}
	if _, ok := v.vU51MPl_at_.Get(uuid.Must(uuid.FromString(key1))); ok {
		t.Fatal("stale UUID")
	}
	if got, ok := v.vU51MPl_at_.Get(uuid.Must(uuid.FromString(key2))); !ok || got != user {
		t.Fatal("new UUID index")
	}
	bad := panelVless(1, "invalid")
	bad.LmOKwjha = 7
	v.spcNlUL(LFWTTXF7JR.Sxcf3wu3L{bad})
	if got, ok := v.vU51MPl_at_.Get(uuid.Must(uuid.FromString(key2))); !ok || got != user || user.GetSpeedLimit() != 7 {
		t.Fatal("failed UUID update must retain UUID and update common fields")
	}
	oldCtx = user.GetCtx()
	v.spcNlUL(nil)
	if oldCtx.Err() != context.Canceled || v.vU51MPl_at_.Len() != 0 || v.jc_QPN6.GTFAWF.Len() != 0 || v.fUxO7UYAH != 0 {
		t.Fatal("removed user not detached")
	}
}
func TestVlessOriginalRestrictedUserCeiling(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	prior := cNT24rG4F.Q0SDIN()
	cNT24rG4F.T15Jifo(true)
	defer cNT24rG4F.T15Jifo(prior)
	v := newVlessUsersFixture(t)
	panels := make(LFWTTXF7JR.Sxcf3wu3L, 90)
	for i := range panels {
		panels[i] = panelVless(int64(i), fmt.Sprintf("%08x-1111-4111-8111-111111111111", i))
	}
	v.spcNlUL(panels)
	if v.fUxO7UYAH != 88 || v.jc_QPN6.GTFAWF.Len() != 88 || v.vU51MPl_at_.Len() != 88 {
		t.Fatal("original 88-user restriction")
	}
	if _, ok := v.jc_QPN6.Get(88); ok {
		t.Fatal("truncated tail accepted")
	}
}
func TestVlessStopPreservesMapsAndClosesListener(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVlessUsersFixture(t)
	v.spcNlUL(LFWTTXF7JR.Sxcf3wu3L{panelVless(1, "11111111-1111-4111-8111-111111111111")})
	user, _ := v.jc_QPN6.Get(1)
	oldCtx := user.GetCtx()
	l, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer l.Close()
	v.kX5i8U = []net.Listener{l}
	v.Stop()
	if _, e = l.Accept(); e == nil {
		t.Fatal("listener remains open")
	}
	if v.zb8f3cHWW8v.Err() != context.Canceled || oldCtx.Err() != context.Canceled || user.GetCtx().Err() != nil || v.jc_QPN6.GTFAWF.Len() != 1 || v.vU51MPl_at_.Len() != 1 || len(v.kX5i8U) != 1 {
		t.Fatal("stop reset/preservation")
	}
}
func TestVlessAPIUpdateRestartAndRules(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVlessUsersFixture(t)
	data := &q5bqCgTdxaTM.D5gTIZ5aVg{H47AzZsYHB1N: &LFWTTXF7JR.Wr5SVP{Z1TIgHOQ: "one"}, ATGYGi0: xx5fRyrFw.VHWq49iRat{"full:block.test"}, NjjWbjH: xx5fRyrFw.VHWq49iRat{"full:white.test"}}
	if v.UpdateAPIData(data) {
		t.Fatal("first init requests restart")
	}
	black := v.oIgEuicmPL.Load()
	white := v._IoxtH.Load()
	if v.UpdateAPIData(data) || black != v.oIgEuicmPL.Load() || white != v._IoxtH.Load() {
		t.Fatal("equal slices should retain current rules")
	}
	data.ATGYGi0 = xx5fRyrFw.VHWq49iRat{"full:changed.test"}
	data.BeqZvr6 = fmt.Errorf("fixture")
	v.UpdateAPIData(data)
	if v.oIgEuicmPL.Load() != black {
		t.Fatal("error result must not replace rules")
	}
	data.H47AzZsYHB1N = &LFWTTXF7JR.Wr5SVP{Z1TIgHOQ: "two"}
	data.BeqZvr6 = nil
	if !v.UpdateAPIData(data) || v.oIgEuicmPL.Load() != black {
		t.Fatal("restart must precede rule updates")
	}
}
