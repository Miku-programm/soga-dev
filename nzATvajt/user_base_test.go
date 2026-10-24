package nzATvajt

import (
	"math"
	"sync"
	"testing"
	"unsafe"

	"recovered.local/soga/LFWTTXF7JR"
)

func TestBaseOriginalLayoutAndInitialization(t *testing.T) {
	u := cLLKJ_A31(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 9, LmOKwjha: 8, B4DYSxuA31za: 2}, nil)
	defer u.kI9WvUIev()
	if unsafe.Sizeof(*u) != 152 || unsafe.Offsetof(u.rQTIvZ) != 112 || unsafe.Offsetof(u.pSde23G4D) != 120 || unsafe.Offsetof(u.mVxQSdX) != 144 {
		t.Fatal("layout")
	}
	if u.GetId() != 9 || u.ckKNgBSf != 0 || u.GetSpeedLimit() != 0 {
		t.Fatal("base constructor must defer limit activation")
	}
	u.UpdateUser(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 99, LmOKwjha: 8, B4DYSxuA31za: 2})
	if u.GetId() != 9 || u.GetSpeedLimit() != 8 || !u.AddTCP() || !u.AddTCP() || u.AddTCP() {
		t.Fatal("update or TCP limit")
	}
	u.DelTCP()
	if !u.AddTCP() {
		t.Fatal("release")
	}
	u.DelTCP()
	u.DelTCP()
	u.DelTCP()
	if u.rQTIvZ.Load() != -1 {
		t.Fatal("original allows negative count")
	}
	u.SetSpeedLimit(2)
	if u.GetSpeedLimit() != 2 {
		t.Fatal("dynamic speed")
	}
	u.SetSpeedLimit(0)
	if u.GetSpeedLimit() != 8 {
		t.Fatal("restore configured speed")
	}
	u.SetSpeedLimit(math.NaN())
	if u.GetSpeedLimit() != 0 {
		t.Fatal("amd64 NaN conversion and clamp")
	}
	u.SetSpeedLimit(-1)
	if u.GetSpeedLimit() != 8 {
		t.Fatal("restore after NaN")
	}
}

func TestOriginalAccountingAndQuota(t *testing.T) {
	p := &LFWTTXF7JR.Aga91y{XuBBcTThBW: 10, UWsYsmytTi: 20, IUlI_15Qf_: 37}
	u := cLLKJ_A31(p, nil)
	defer u.kI9WvUIev()
	u.tKYtJcjS.UpdateLimit(1)
	if u.AllowAndAddUp(2) || u.AllowAndAddDown(3) {
		t.Fatal("oversize must reject")
	}
	u.AddUp(2)
	if !u.IsTrafficOut() {
		t.Fatal("rejected traffic still counts toward quota")
	}
	if up, down := u.GetAndClearTraffic(); up != 4 || down != 3 {
		t.Fatal(up, down)
	}
	if u.IsTrafficOut() {
		t.Fatal("cleared counters no longer contribute")
	}
	if up, down := u.GetAndClearDyTraffic(); up != 4 || down != 3 {
		t.Fatal(up, down)
	}
	if up, down := u.GetAndClearTraffic(); up != 0 || down != 0 {
		t.Fatal(up, down)
	}
	u.AddDown(-7)
	if _, down := u.GetAndClearTraffic(); down != -7 {
		t.Fatal("signed original counters", down)
	}
	u.hBZll21 = math.MaxInt64
	u.uIawOCHn = 1
	u.pq4dsrgDKl9 = 1
	if u.IsTrafficOut() {
		t.Fatal("original signed overflow")
	}
}

func TestAtomicCountersAndContext(t *testing.T) {
	u := cLLKJ_A31(&LFWTTXF7JR.Aga91y{}, nil)
	defer func() { u.kI9WvUIev() }()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				u.AddUp(1)
				u.AddDown(2)
			}
		}()
	}
	wg.Wait()
	if up, down := u.GetAndClearTraffic(); up != 8000 || down != 16000 {
		t.Fatal(up, down)
	}
	old := u.GetCtx()
	u.ResetCtx()
	if old.Err() == nil || u.GetCtx().Err() != nil || old == u.GetCtx() {
		t.Fatal("context renewal")
	}
}

func TestAliveIPReportingThroughUser(t *testing.T) {
	u := cLLKJ_A31(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 7, Xmrna8: 1}, nil)
	defer u.kI9WvUIev()
	if !u.AddAliveIp("one", "one") || u.AddAliveIp("two", "two") {
		t.Fatal("IP limit")
	}
	if u.GetAliveIpData() != nil {
		t.Fatal("count does not record reporting traffic")
	}
	u.RecordAliveIp("one")
	for i := 0; i < 2; i++ {
		data := u.GetAliveIpData()
		if data == nil || data.Vy9ybu != 7 || len(data.PNtn9KAV) != 1 || data.PNtn9KAV[0] != "one" {
			t.Fatal(data)
		}
	}
	u.DelAliveIp("one", "one")
	if u.GetAliveIpData() == nil || u.GetAliveIpData() != nil {
		t.Fatal("offline final report")
	}
}
