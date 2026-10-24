package inbound

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/binary"
	"hash/crc32"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"

	G "recovered.local/soga/util/genmap"
	J "recovered.local/soga/net/sess"
	L "recovered.local/soga/panel/types"
	N "recovered.local/soga/net/limit"
	C "recovered.local/soga/config"
	V "recovered.local/soga/protocol/vmess"
	U "recovered.local/soga/user"
)

func vmessPanelFixture(id int64, key string) *L.Aga91y {
	return &L.Aga91y{WZ89pTAOrHK4: id, JAfXWwzYrW: &L.Yn8BJoL{SfiEGpJo: key}}
}
func newVmessAuthFixture(t *testing.T) *K4avFiYOPi {
	t.Helper()
	b, err := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(b.yPUbKO)
	return &K4avFiYOPi{JWN154qJ4: b, p54XvyV: &U.ZpNFqUVZp{GTFAWF: &G.GTFAWF[int64, U.PgaNd3014U]{}}, l_ZoZQDxG7: J.YCcacra(b.zb8f3cHWW8v), qgz9PfTS: N.Azb3Zr(b.zb8f3cHWW8v, &C.Nl62Lnd_{SEU_Z687ugp: true, GxuEkr: 2, QS255wlN7m2W: time.Minute, NjiRkFvidOaA: time.Minute})}
}
func vmessAuthID(user *U.QKaIWV2, timestamp int64, validCRC bool) (id [16]byte) {
	var plain [16]byte
	binary.BigEndian.PutUint64(plain[:8], uint64(timestamp))
	copy(plain[8:12], []byte{13, 27, 39, 51})
	binary.BigEndian.PutUint32(plain[12:], crc32.ChecksumIEEE(plain[:12]))
	if !validCRC {
		plain[15] ^= 1
	}
	user.GetBlock().Encrypt(id[:], plain[:])
	return
}
func TestVmessNativeUserAuthAndRemoval(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	var layout K4avFiYOPi
	if unsafe.Sizeof(layout) != 64 || unsafe.Offsetof(layout.dj_cTnaUp6f) != 32 || unsafe.Offsetof(layout.qgz9PfTS) != 56 {
		t.Fatal("original RTTI")
	}
	v := newVmessAuthFixture(t)
	key := "11111111-1111-4111-8111-111111111111"
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(7, key)})
	base, ok := v.jc_QPN6.Get(7)
	if !ok {
		t.Fatal("missing user")
	}
	u := base.(*U.QKaIWV2)
	id := vmessAuthID(u, time.Now().Unix(), true)
	auth := v.y80zV6("192.0.2.1", id)
	if auth == nil || auth.Xwwpafq != u || !auth.Ih7r8HyMvOa || auth.B2pr6gMaPO != [16]byte{} || string(auth.T_ul7Q_0j) != string(id[:]) {
		t.Fatal("AEAD auth result")
	}
	// 主表移除后仍能命中原 IP 缓存；正式面板删除必须调用 RemoveUsers。
	v.jc_QPN6.Delete(7)
	if v.y80zV6("192.0.2.1", id) == nil || v.y80zV6("192.0.2.2", id) != nil {
		t.Fatal("cache-first lookup")
	}
	v.jc_QPN6.Add(u)
	before := u.GetCtx()
	v.spcNlUL(nil)
	if v.y80zV6("192.0.2.1", id) != nil || before.Err() != context.Canceled || v.p54XvyV.Len() != 0 {
		t.Fatal("panel removal did not invalidate auth cache")
	}
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(7, key)})
	beforeUser, _ := v.jc_QPN6.Get(7)
	before = beforeUser.GetCtx()
	v.Stop()
	if v.jc_QPN6.Len() != 0 || before.Err() != context.Canceled {
		t.Fatal("VMess Stop must delete users")
	}
}
func TestVmessNativeAuthValidationAndBan(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVmessAuthFixture(t)
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(1, "11111111-1111-4111-8111-111111111111")})
	b, _ := v.jc_QPN6.Get(1)
	u := b.(*U.QKaIWV2)
	for _, tc := range []struct {
		name  string
		delta int64
		crc   bool
	}{{"CRC", 0, false}, {"old", -180, true}, {"future", 180, true}} {
		t.Run(tc.name, func(t *testing.T) {
			ip := "192.0.2." + tc.name
			bad := vmessAuthID(u, time.Now().Unix()+tc.delta, tc.crc)
			for i := 0; i < 3; i++ {
				if v.y80zV6(ip, bad) != nil {
					t.Fatal("accepted invalid auth")
				}
			}
			if !v.qgz9PfTS.IsForbidden(ip) {
				t.Fatal("no ban after failed attempts")
			}
			good := vmessAuthID(u, time.Now().Unix(), true)
			if v.y80zV6(ip, good) != nil {
				t.Fatal("ban bypass")
			}
			v.qgz9PfTS.Valid(ip)
			if v.y80zV6(ip, good) == nil {
				t.Fatal("auth after clearing ban")
			}
		})
	}
}

type vmessReadFixture struct {
	net.Conn
	reads     [][]byte
	readErr   error
	deadlines []time.Time
}

func (c *vmessReadFixture) Read(p []byte) (int, error) {
	if len(c.reads) == 0 {
		return 0, io.EOF
	}
	b := c.reads[0]
	c.reads = c.reads[1:]
	return copy(p, b), c.readErr
}
func (c *vmessReadFixture) SetReadDeadline(d time.Time) error {
	c.deadlines = append(c.deadlines, d)
	return nil
}
func (c *vmessReadFixture) RemoteAddr() net.Addr {
	return &net.TCPAddr{IP: net.IPv4(192, 0, 2, 9), Port: 2000}
}
func TestVmessNativeAuthInitialReadSemantics(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVmessAuthFixture(t)
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(1, "11111111-1111-4111-8111-111111111111")})
	b, _ := v.jc_QPN6.Get(1)
	id := vmessAuthID(b.(*U.QKaIWV2), time.Now().Unix(), true)
	for _, n := range []int{0, 1, 15, 16} {
		c := &vmessReadFixture{reads: [][]byte{nil, id[:n]}}
		if n == 0 {
			c.readErr = io.EOF
		}
		auth, fallback := v.s3mqrqd_(c)
		if n == 16 {
			if auth == nil || fallback != nil {
				t.Fatal("full auth")
			}
		} else if auth != nil || string(fallback) != string(id[:n]) {
			t.Fatalf("short read %d", n)
		}
		if len(c.deadlines) != 2 || c.deadlines[0].IsZero() || !c.deadlines[1].IsZero() {
			t.Fatal("deadline reset")
		}
	}
	c := &vmessReadFixture{reads: [][]byte{id[:]}, readErr: io.EOF}
	if auth, fallback := v.s3mqrqd_(c); auth != nil || string(fallback) != string(id[:]) {
		t.Fatal("data plus EOF must use fallback")
	}
}

func TestVmessNativeMD5SelectionAndUserRemoval(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVmessAuthFixture(t)
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(42, "11111111-1111-4111-8111-111111111111")})
	b, _ := v.jc_QPN6.Get(42)
	u := b.(*U.QKaIWV2)
	v.dj_cTnaUp6f = V.NewMD5AuthManager(v.zb8f3cHWW8v, v.jc_QPN6, true)
	v.dj_cTnaUp6f.StartTask()
	stamp := uint32(time.Now().Unix())
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], uint64(stamp))
	idBytes := u.GetUserIds()[0].Yfd33z6YYdJa
	h := hmac.New(md5.New, idBytes[:])
	h.Write(data[:])
	id := [16]byte(h.Sum(nil))
	until := time.Now().Add(2 * time.Second)
	for {
		if _, _, ok := v.dj_cTnaUp6f.Get(id); ok {
			break
		}
		if time.Now().After(until) {
			t.Fatal("MD5 window not initialized")
		}
		time.Sleep(time.Millisecond)
	}
	// 同一入口先走MD5，不触发AEAD错误计数。
	c := &vmessReadFixture{reads: [][]byte{id[:]}}
	auth, fallback := v.s3mqrqd_(c)
	if auth == nil || fallback != nil || auth.Ih7r8HyMvOa || auth.Xwwpafq != u || auth.QAkNeaNePGH != stamp || auth.B2pr6gMaPO != id {
		t.Fatal("MD5 selection")
	}
	v.jc_QPN6.Delete(42)
	if v.iUB_i9oBP_68(id) != nil {
		t.Fatal("stale MD5 bucket authenticated removed user")
	}
}

func TestVmessNativeRestrictedCeiling(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	before := C.Q0SDIN()
	C.T15Jifo(true)
	defer C.T15Jifo(before)
	v := newVmessAuthFixture(t)
	panels := make(L.Sxcf3wu3L, 90)
	for i := range panels {
		panels[i] = vmessPanelFixture(int64(i), "11111111-1111-4111-8111-111111111111")
	}
	v.spcNlUL(panels)
	if v.jc_QPN6.Len() != 88 || v.p54XvyV.Len() != 88 || v.fUxO7UYAH != 88 || userCountAt50C6DA0 != 88 {
		t.Fatal("original restricted ceiling")
	}
}
