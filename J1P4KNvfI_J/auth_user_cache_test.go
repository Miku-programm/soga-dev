package J1P4KNvfI_J

import (
	"bytes"
	"container/list"
	"encoding/hex"
	"errors"
	"io"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/jellydator/ttlcache/v3"
	B "recovered.local/soga/D_zCag1qTV"
	M "recovered.local/soga/G7A7csgrs62p"
	C "recovered.local/soga/cNT24rG4F"
	U "recovered.local/soga/nzATvajt"
)

type cacheUser struct {
	U.PgaNd3014U
	id int64
}

func (u *cacheUser) GetId() int64 { return u.id }
func cacheFixture() *EafxyUZD {
	return &EafxyUZD{dd_Dh_6Zne: time.Hour, eGhvKTJIaDi: ttlcache.New(ttlcache.WithTTL[string, *Elkpyuoj](time.Hour))}
}
func TestAuthCacheOriginalLayout(t *testing.T) {
	for _, tc := range []struct {
		v any
		n uintptr
	}{{HRRGg0nld{}, 24}, {Elkpyuoj{}, 48}, {EafxyUZD{}, 40}, {YHksCjPM{}, 16}, {B07Fb7qv{}, 48}} {
		if n := reflect.TypeOf(tc.v).Size(); n != tc.n {
			t.Fatalf("%T size %d", tc.v, n)
		}
	}
}
func TestAuthCacheHitOrderExpiryAndDuplicateIdentity(t *testing.T) {
	c := cacheFixture()
	a, b := &cacheUser{id: 1}, &cacheUser{id: 2}
	c.AddUser("127.0.0.1", a, 1)
	c.AddUser("127.0.0.1", b, time.Now().Add(time.Hour).UnixMilli())
	c.AddUser("127.0.0.1", &cacheUser{id: 1}, time.Now().Add(2*time.Hour).UnixMilli())
	var seen []int64
	c.RangeUser("127.0.0.1", func(u U.PgaNd3014U) bool {
		seen = append(seen, u.GetId())
		if u.GetId() == 1 && u != a {
			t.Fatal("duplicate replaced original identity")
		}
		return u.GetId() != 1
	})
	if !reflect.DeepEqual(seen, []int64{2, 1}) {
		t.Fatal(seen)
	}
	l := c.eGhvKTJIaDi.Get("127.0.0.1").Value()
	if l.Count() != 2 || l.gvgrVQca.Front().Value.(*HRRGg0nld).jMdUpB <= time.Now().UnixMilli() {
		t.Fatal("expired match must be refreshed and moved first")
	}
	seen = nil
	c.RangeUser("127.0.0.1", func(u U.PgaNd3014U) bool { seen = append(seen, u.GetId()); return true })
	if !reflect.DeepEqual(seen, []int64{1, 2}) {
		t.Fatal(seen)
	}
	c.AddUser("::1", a, 0)
	c.RemoveUsers([]U.PgaNd3014U{a})
	if l.Count() != 1 || c.eGhvKTJIaDi.Get("::1").Value().Count() != 0 {
		t.Fatal("removal must visit every IP bucket")
	}
}
func TestAuthCacheVisitsExpiredBeforeRemoving(t *testing.T) {
	l := &Elkpyuoj{ite4WB58: time.Hour, gvgrVQca: list.New(), zerMi8z: make(map[int64]*list.Element)}
	l.AddUser(&cacheUser{id: 1}, 1)
	l.AddUser(&cacheUser{id: 2}, 1)
	calls := 0
	l.Range(func(v *HRRGg0nld) bool { calls++; return true }, false)
	if calls != 2 || l.Count() != 0 || len(l.zerMi8z) != 0 {
		t.Fatal(calls, l.Count())
	}
	l.AddUser(&cacheUser{id: 3}, 1)
	l.DeleteExpired()
	if l.Count() != 0 {
		t.Fatal("cleanup")
	}
}
func TestAuthCacheRecordExactWireAndShortRecordQuirk(t *testing.T) {
	raw, _ := hex.DecodeString("01c00002010000000101020304050607080000000000000009")
	r := &B07Fb7qv{net.ParseIP("192.0.2.1"), CXTczIZ5hN7o{&YHksCjPM{0x0102030405060708, 9}}}
	b := B.S4OQXWIZ(0)
	defer b.Free()
	r.Write(b)
	if !bytes.Equal(b.Bytes(), raw) {
		t.Fatalf("wire %x", b.Bytes())
	}
	out, err := LAQrKU(bytes.NewBuffer(raw))
	if err != nil || !out.SRR4_uk.Equal(r.SRR4_uk) || *out.ZFFCVgG8d[0] != *r.ZFFCVgG8d[0] {
		t.Fatal(out, err)
	}
	for n := 0; n < len(raw); n++ {
		out, err := LAQrKU(bytes.NewBuffer(raw[:n]))
		if out != nil {
			t.Fatalf("short %d returned record", n)
		}
		switch {
		case n == 0:
			if !errors.Is(err, io.EOF) {
				t.Fatal(n, err)
			}
		case n < 9:
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatal(n, err)
			}
		default:
			if err != nil {
				t.Fatalf("original nil error branch at %d: %v", n, err)
			}
		}
	}
	if _, err := LAQrKU(bytes.NewBuffer([]byte{3})); err == nil || err.Error() != "unknown ip type" {
		t.Fatal(err)
	}
	b.Reset()
	r.SRR4_uk = net.ParseIP("2001:db8::1")
	r.Write(b)
	out, err = LAQrKU(bytes.NewBuffer(b.Bytes()))
	if err != nil || !out.SRR4_uk.Equal(r.SRR4_uk) || b.Byte(0) != 2 {
		t.Fatal(out, err)
	}
}
func TestAuthCachePersistAndLoadFilters(t *testing.T) {
	// 本测试不初始化全局配置；确保 Save 的原目录副作用不会落到系统路径。
	if C.IPUserCacheDirectory() != "" {
		t.Fatal("fixture requires untouched configuration")
	}
	a, b := &cacheUser{id: 1}, &cacheUser{id: 2}
	c := cacheFixture()
	c.AddUser("192.0.2.1", a, time.Now().Add(time.Hour).UnixMilli())
	c.AddUser("2001:db8::1", b, 1)
	c.AddUser("not-an-ip", a, 0)
	p := filepath.Join(t.TempDir(), "users.bin")
	if err := c.Save(p); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(data)
	records := 0
	for {
		r, e := LAQrKU(buf)
		if errors.Is(e, io.EOF) {
			break
		}
		if e != nil {
			t.Fatal(e)
		}
		records++
		if r.SRR4_uk == nil {
			t.Fatal(r)
		}
	}
	if records != 2 {
		t.Fatal(records)
	}
	users := &U.ZpNFqUVZp{GTFAWF: new(M.GTFAWF[int64, U.PgaNd3014U])}
	users.Add(a)
	users.Add(b)
	restored := cacheFixture()
	ok, err := restored.Load(p, users)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ids := []int64{}
	restored.RangeUser("192.0.2.1", func(u U.PgaNd3014U) bool { ids = append(ids, u.GetId()); return true })
	restored.RangeUser("2001:db8::1", func(u U.PgaNd3014U) bool { t.Fatal("expired persisted entry accepted"); return true })
	if !reflect.DeepEqual(ids, []int64{1}) {
		t.Fatal(ids)
	}
	if ok, err := restored.Load(p+".missing", users); ok || err != nil {
		t.Fatal(ok, err)
	}
	if err := os.WriteFile(p, []byte{3}, 0600); err != nil {
		t.Fatal(err)
	}
	if ok, err := restored.Load(p, users); ok || err == nil {
		t.Fatal(ok, err)
	}
}
