package I4AWsAkBLRP

import (
	"reflect"
	"testing"
)

func TestAllOriginalRTTI(t *testing.T) {
	{
		v := reflect.TypeOf(yoowfr{})
		if v.Size() != 24 {
			t.Fatalf("yoowfr size %d", v.Size())
		}
		if f, ok := v.FieldByName("b1C__L1"); !ok || f.Offset != 0 || f.Type.String() != "int" || string(f.Tag) != "" {
			t.Fatalf("yoowfr.b1C__L1: %+v", f)
		}
		if f, ok := v.FieldByName("uRSOFlo"); !ok || f.Offset != 8 || f.Type.String() != "error" || string(f.Tag) != "" {
			t.Fatalf("yoowfr.uRSOFlo: %+v", f)
		}
	}
	{
		v := reflect.TypeOf(TKUtOq{})
		if v.Size() != 32 {
			t.Fatalf("TKUtOq size %d", v.Size())
		}
		if f, ok := v.FieldByName("xlaU77zvaW7W"); !ok || f.Offset != 0 || f.Type.String() != "uint8" || string(f.Tag) != "" {
			t.Fatalf("TKUtOq.xlaU77zvaW7W: %+v", f)
		}
		if f, ok := v.FieldByName("gUa4uFKR6"); !ok || f.Offset != 1 || f.Type.String() != "uint8" || string(f.Tag) != "" {
			t.Fatalf("TKUtOq.gUa4uFKR6: %+v", f)
		}
		if f, ok := v.FieldByName("a_9mrWgHy"); !ok || f.Offset != 4 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("TKUtOq.a_9mrWgHy: %+v", f)
		}
		if f, ok := v.FieldByName("cAKGyUaj"); !ok || f.Offset != 8 || f.Type.String() != "[]uint8" || string(f.Tag) != "" {
			t.Fatalf("TKUtOq.cAKGyUaj: %+v", f)
		}
	}
	{
		v := reflect.TypeOf(vwv8dqM{})
		if v.Size() != 56 {
			t.Fatalf("vwv8dqM size %d", v.Size())
		}
		if f, ok := v.FieldByName("e9dK1WBuMa"); !ok || f.Offset != 0 || f.Type.String() != "I4AWsAkBLRP.JZpBeJdC" || string(f.Tag) != "" {
			t.Fatalf("vwv8dqM.e9dK1WBuMa: %+v", f)
		}
		if f, ok := v.FieldByName("fN47GQEn"); !ok || f.Offset != 8 || f.Type.String() != "I4AWsAkBLRP.TKUtOq" || string(f.Tag) != "" {
			t.Fatalf("vwv8dqM.fN47GQEn: %+v", f)
		}
		if f, ok := v.FieldByName("vlOxyuBq8ThB"); !ok || f.Offset != 40 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("vwv8dqM.vlOxyuBq8ThB: %+v", f)
		}
		if f, ok := v.FieldByName("wbTHvwrgbD"); !ok || f.Offset != 48 || f.Type.String() != "chan I4AWsAkBLRP.yoowfr" || string(f.Tag) != "" {
			t.Fatalf("vwv8dqM.wbTHvwrgbD: %+v", f)
		}
	}
	{
		v := reflect.TypeOf(BWpFHH3x{})
		if v.Size() != 56 {
			t.Fatalf("BWpFHH3x size %d", v.Size())
		}
		if f, ok := v.FieldByName("RTNOROJC5"); !ok || f.Offset != 0 || f.Type.String() != "int" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.RTNOROJC5: %+v", f)
		}
		if f, ok := v.FieldByName("DU0iTVAg"); !ok || f.Offset != 8 || f.Type.String() != "bool" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.DU0iTVAg: %+v", f)
		}
		if f, ok := v.FieldByName("PGgIC2"); !ok || f.Offset != 16 || f.Type.String() != "time.Duration" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.PGgIC2: %+v", f)
		}
		if f, ok := v.FieldByName("PK9EpD"); !ok || f.Offset != 24 || f.Type.String() != "time.Duration" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.PK9EpD: %+v", f)
		}
		if f, ok := v.FieldByName("XsqVEjw4"); !ok || f.Offset != 32 || f.Type.String() != "int" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.XsqVEjw4: %+v", f)
		}
		if f, ok := v.FieldByName("K3nkPQuOI"); !ok || f.Offset != 40 || f.Type.String() != "int" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.K3nkPQuOI: %+v", f)
		}
		if f, ok := v.FieldByName("CUuHLYfWnC8q"); !ok || f.Offset != 48 || f.Type.String() != "int" || string(f.Tag) != "" {
			t.Fatalf("BWpFHH3x.CUuHLYfWnC8q: %+v", f)
		}
	}
	{
		v := reflect.TypeOf(YboZfdKYbbwF{})
		if v.Size() != 200 {
			t.Fatalf("YboZfdKYbbwF size %d", v.Size())
		}
		if f, ok := v.FieldByName("aJW2r4quXUZ0"); !ok || f.Offset != 0 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.aJW2r4quXUZ0: %+v", f)
		}
		if f, ok := v.FieldByName("ax9eSX"); !ok || f.Offset != 8 || f.Type.String() != "*I4AWsAkBLRP.DZOCfwWJsy" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.ax9eSX: %+v", f)
		}
		if f, ok := v.FieldByName("dkrfsm"); !ok || f.Offset != 16 || f.Type.String() != "[][]uint8" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.dkrfsm: %+v", f)
		}
		if f, ok := v.FieldByName("cCO3_XPTNnnL"); !ok || f.Offset != 40 || f.Type.String() != "[][]uint8" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.cCO3_XPTNnnL: %+v", f)
		}
		if f, ok := v.FieldByName("jv_tE_9jqFgQ"); !ok || f.Offset != 64 || f.Type.String() != "sync.Mutex" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.jv_tE_9jqFgQ: %+v", f)
		}
		if f, ok := v.FieldByName("_oSX6fTfG"); !ok || f.Offset != 72 || f.Type.String() != "int" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF._oSX6fTfG: %+v", f)
		}
		if f, ok := v.FieldByName("fQuoCeZnfHdC"); !ok || f.Offset != 80 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.fQuoCeZnfHdC: %+v", f)
		}
		if f, ok := v.FieldByName("kPA6eSs3XGMS"); !ok || f.Offset != 88 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.kPA6eSs3XGMS: %+v", f)
		}
		if f, ok := v.FieldByName("odTLapRPCdm"); !ok || f.Offset != 96 || f.Type.String() != "sync.Once" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.odTLapRPCdm: %+v", f)
		}
		if f, ok := v.FieldByName("rcczhz"); !ok || f.Offset != 112 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.rcczhz: %+v", f)
		}
		if f, ok := v.FieldByName("cs8cXKxa"); !ok || f.Offset != 120 || f.Type.String() != "sync.Once" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.cs8cXKxa: %+v", f)
		}
		if f, ok := v.FieldByName("uWxaXR0T"); !ok || f.Offset != 136 || f.Type.String() != "atomic.Value" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.uWxaXR0T: %+v", f)
		}
		if f, ok := v.FieldByName("dDtIx_"); !ok || f.Offset != 152 || f.Type.String() != "atomic.Value" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.dDtIx_: %+v", f)
		}
		if f, ok := v.FieldByName("fTT0Ew3m"); !ok || f.Offset != 168 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.fTT0Ew3m: %+v", f)
		}
		if f, ok := v.FieldByName("cGnO8z"); !ok || f.Offset != 172 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.cGnO8z: %+v", f)
		}
		if f, ok := v.FieldByName("xwvruuYKD"); !ok || f.Offset != 176 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.xwvruuYKD: %+v", f)
		}
		if f, ok := v.FieldByName("f8xi72G1Ov"); !ok || f.Offset != 180 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.f8xi72G1Ov: %+v", f)
		}
		if f, ok := v.FieldByName("cXWhLPc"); !ok || f.Offset != 184 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.cXWhLPc: %+v", f)
		}
		if f, ok := v.FieldByName("waGBYT"); !ok || f.Offset != 192 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("YboZfdKYbbwF.waGBYT: %+v", f)
		}
	}
	{
		v := reflect.TypeOf(DZOCfwWJsy{})
		if v.Size() != 264 {
			t.Fatalf("DZOCfwWJsy size %d", v.Size())
		}
		if f, ok := v.FieldByName("mlu21WY6"); !ok || f.Offset != 0 || f.Type.String() != "io.ReadWriteCloser" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.mlu21WY6: %+v", f)
		}
		if f, ok := v.FieldByName("wpUkUgrJJHwN"); !ok || f.Offset != 16 || f.Type.String() != "*I4AWsAkBLRP.BWpFHH3x" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.wpUkUgrJJHwN: %+v", f)
		}
		if f, ok := v.FieldByName("mTqgfhNdlT82"); !ok || f.Offset != 24 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.mTqgfhNdlT82: %+v", f)
		}
		if f, ok := v.FieldByName("p2tD2AhCZnC"); !ok || f.Offset != 28 || f.Type.String() != "sync.Mutex" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.p2tD2AhCZnC: %+v", f)
		}
		if f, ok := v.FieldByName("zbny3Aedj"); !ok || f.Offset != 36 || f.Type.String() != "int32" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.zbny3Aedj: %+v", f)
		}
		if f, ok := v.FieldByName("tnmOyFJvf"); !ok || f.Offset != 40 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.tnmOyFJvf: %+v", f)
		}
		if f, ok := v.FieldByName("wjubqp"); !ok || f.Offset != 48 || f.Type.String() != "map[uint32]*I4AWsAkBLRP.YboZfdKYbbwF" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.wjubqp: %+v", f)
		}
		if f, ok := v.FieldByName("r6klva"); !ok || f.Offset != 56 || f.Type.String() != "sync.RWMutex" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.r6klva: %+v", f)
		}
		if f, ok := v.FieldByName("nHOcCGRza7V"); !ok || f.Offset != 80 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.nHOcCGRza7V: %+v", f)
		}
		if f, ok := v.FieldByName("h9RMpMidR"); !ok || f.Offset != 88 || f.Type.String() != "sync.Once" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.h9RMpMidR: %+v", f)
		}
		if f, ok := v.FieldByName("l4IUcN"); !ok || f.Offset != 104 || f.Type.String() != "atomic.Value" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.l4IUcN: %+v", f)
		}
		if f, ok := v.FieldByName("xnMhVM5Fh"); !ok || f.Offset != 120 || f.Type.String() != "atomic.Value" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.xnMhVM5Fh: %+v", f)
		}
		if f, ok := v.FieldByName("dq5F2u2"); !ok || f.Offset != 136 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.dq5F2u2: %+v", f)
		}
		if f, ok := v.FieldByName("nvN2Bu"); !ok || f.Offset != 144 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.nvN2Bu: %+v", f)
		}
		if f, ok := v.FieldByName("_5AKaibLIdIF"); !ok || f.Offset != 152 || f.Type.String() != "sync.Once" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy._5AKaibLIdIF: %+v", f)
		}
		if f, ok := v.FieldByName("gFei_ksGsjwr"); !ok || f.Offset != 164 || f.Type.String() != "sync.Once" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.gFei_ksGsjwr: %+v", f)
		}
		if f, ok := v.FieldByName("u1sgxOsIJxL"); !ok || f.Offset != 176 || f.Type.String() != "atomic.Value" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.u1sgxOsIJxL: %+v", f)
		}
		if f, ok := v.FieldByName("eLUgKTG"); !ok || f.Offset != 192 || f.Type.String() != "chan struct {}" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.eLUgKTG: %+v", f)
		}
		if f, ok := v.FieldByName("eR6WSnZ4i"); !ok || f.Offset != 200 || f.Type.String() != "sync.Once" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.eR6WSnZ4i: %+v", f)
		}
		if f, ok := v.FieldByName("cf0Ovj"); !ok || f.Offset != 216 || f.Type.String() != "chan *I4AWsAkBLRP.YboZfdKYbbwF" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.cf0Ovj: %+v", f)
		}
		if f, ok := v.FieldByName("nZwZauYdh"); !ok || f.Offset != 224 || f.Type.String() != "int32" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.nZwZauYdh: %+v", f)
		}
		if f, ok := v.FieldByName("jE8dYZ8F_"); !ok || f.Offset != 228 || f.Type.String() != "int32" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.jE8dYZ8F_: %+v", f)
		}
		if f, ok := v.FieldByName("a2A6Mhwyca"); !ok || f.Offset != 232 || f.Type.String() != "atomic.Value" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.a2A6Mhwyca: %+v", f)
		}
		if f, ok := v.FieldByName("bVqjxgmOu"); !ok || f.Offset != 248 || f.Type.String() != "uint32" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.bVqjxgmOu: %+v", f)
		}
		if f, ok := v.FieldByName("hHxKb51kQa"); !ok || f.Offset != 256 || f.Type.String() != "chan I4AWsAkBLRP.vwv8dqM" || string(f.Tag) != "" {
			t.Fatalf("DZOCfwWJsy.hHxKb51kQa: %+v", f)
		}
	}
}
