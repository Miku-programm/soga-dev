package ss2022

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"lukechampine.com/blake3"
	"net"
	"os"
	"path/filepath"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
	"testing"
	"testing/synctest"
	"time"
	"unsafe"
)

func userFixture(t *testing.T, method ssmethod.C6q3Upu, id int64, fill byte) *user.WtK9HQSy66 {
	t.Helper()
	n, err := method.KeyLen()
	if err != nil {
		t.Fatal(err)
	}
	user, err := user.MbpJoYH(&types.Aga91y{WZ89pTAOrHK4: id, JAfXWwzYrW: &types.YlfOV0DBBnx{T1oDDIgYJHh: method, A_wNUTkMh6pH: base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{fill}, n))}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	return user
}
func managerFixture(t *testing.T, ctx context.Context, method ssmethod.C6q3Upu) *N07NaA {
	t.Helper()
	n, err := method.KeyLen()
	if err != nil {
		t.Fatal(err)
	}
	m, err := DkAuKfGU(ctx, &types.AAMPrdcYTL6{HMkc4BgR: method, VuhWqou: base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{7}, n))})
	if err != nil {
		t.Fatal(err)
	}
	return m
}
func TestNativeSessionLayouts(t *testing.T) {
	var m N07NaA
	var s RGGcUu1R
	var w KHKdG8c
	if unsafe.Sizeof(m) != 96 || unsafe.Offsetof(m.su5Rzwte) != 64 || unsafe.Sizeof(s) != 1104 || unsafe.Offsetof(s.paeyZKdn) != 72 || unsafe.Sizeof(w) != 1032 {
		t.Fatal("RTTI layout")
	}
}
func TestReplayWindowOriginalBoundariesAndPartialReset(t *testing.T) {
	var w KHKdG8c
	if !w.Check(0) {
		t.Fatal("zero initially accepted")
	}
	w.Add(0)
	if w.Check(0) {
		t.Fatal("replay")
	}
	w.Add(63)
	w.Add(64)
	if w.Check(63) || w.Check(64) || !w.Check(65) {
		t.Fatal("word boundary")
	}
	w.Add(9000)
	if !w.Check(9000-8128) || w.Check(9000-8129) {
		t.Fatal("inclusive backward distance")
	}
	w.Add(9000 - 8128)
	if w.Check(9000 - 8128) {
		t.Fatal("backward replay")
	}
	w.nypzxyc8qDxT[10] = 123
	w.Reset()
	if w.rsPnryPvPJ != 0 || w.nypzxyc8qDxT[0] != 0 || w.nypzxyc8qDxT[10] != 123 {
		t.Fatal("original Reset clears only 16 bytes")
	}
	w.Add(^uint64(0))
	if w.Check(^uint64(0)) || !w.Check(^uint64(0)-1) {
		t.Fatal("maximum uint64")
	}
}
func TestOriginalSessionDirectionalKeysAndAddressIndex(t *testing.T) {
	for _, method := range []ssmethod.C6q3Upu{"2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		m := managerFixture(t, ctx, method)
		u := userFixture(t, method, 12, 3)
		m.UpdateUser(u)
		full := blake3.Sum512(u.PSK())
		var hash [16]byte
		copy(hash[:], full[:16])
		if h, ok := m.GetHashById(12); !ok || h != hash {
			t.Fatal("user hash")
		}
		if found, ok := m.GetUserByHash(hash); !ok || found != u {
			t.Fatal("user index")
		}
		old := hash
		u.PSK()[0] ^= 1
		m.UpdateUser(u)
		if _, ok := m.GetUserByHash(old); ok {
			t.Fatal("old identity retained")
		}
		a := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1000}
		s, err := m.GetOrNewSession(123, u, a)
		if err != nil {
			t.Fatal(err)
		}
		if s.fl_h6QOtKast == 123 || s.e5Lk2aYUf != 123 || s.c9dMiAM0t89w.Add(1) != 0 {
			t.Fatal("IDs/counter")
		}
		var salt [8]byte
		for _, pair := range []struct {
			id     uint64
			server bool
		}{{123, false}, {s.fl_h6QOtKast, true}} {
			binary.BigEndian.PutUint64(salt[:], pair.id)
			peer, err := method.NewAEAD(method.SubKey(u.PSK(), salt[:]))
			if err != nil {
				t.Fatal(err)
			}
			aead := s.ul7vaxxk4q4
			if pair.server {
				aead = s.qXprYj6U9AK
			}
			sealed := aead.Seal(nil, make([]byte, 12), []byte("native session frame"), nil)
			got, err := peer.Open(nil, make([]byte, 12), sealed, nil)
			if err != nil || string(got) != "native session frame" {
				t.Fatal(err)
			}
		}
		b := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 2000}
		again, err := m.GetOrNewSession(123, nil, b)
		if err != nil || again != s || s.ob06bMoWqrSi != b {
			t.Fatal("ID reuse")
		}
		if found, ok := m.GetSessionByClientAddr(a); !ok || found != s {
			t.Fatal("original old address mapping retained")
		}
		if found, ok := m.GetSessionByClientAddr(b); !ok || found != s {
			t.Fatal("new address mapping")
		}
		m.RemoveUser(12)
		if _, ok := m.GetHashById(12); ok {
			t.Fatal("user removal")
		}
	}
}
func TestSaltTTLDoesNotExtendOnHit(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		m := managerFixture(t, ctx, "2022-blake3-aes-128-gcm")
		salt := []byte("salt")
		m.AddSalt(salt)
		salt[0] = 'x'
		if !m.CheckSalt([]byte("salt")) || m.CheckSalt(salt) {
			t.Fatal("AddSalt must copy")
		}
		time.Sleep(50 * time.Second)
		if !m.CheckSalt([]byte("salt")) {
			t.Fatal("early expiration")
		}
		time.Sleep(11 * time.Second)
		if m.CheckSalt([]byte("salt")) {
			t.Fatal("lookup extended salt TTL")
		}
		defer func() {
			if recover() == nil {
				t.Fatal("empty CheckSalt must panic")
			}
		}()
		m.CheckSalt(nil)
	})
}

// 原配置默认值在显式加载时才赋到结构体；测试也遵循该启动顺序。
func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "soga-native-ss-test-")
	if err != nil {
		panic(err)
	}
	path := filepath.Join(dir, "soga.conf")
	if err = os.WriteFile(path, []byte("udp_timeout=2\nlog_file=false\naccess_log_enable=false\n"), 0600); err != nil {
		panic(err)
	}
	config.MkwNWciPHS(path)
	if err = config.TvBECnD9fM6(); err != nil {
		panic(err)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}
