//go:build linux

package inbound

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	E "recovered.local/soga/net/outbound"
	A "recovered.local/soga/protocol/ssactual"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	F "recovered.local/soga/panel/ssrnames"
	Q "recovered.local/soga/panel/registry"
	"sort"
	"testing"
	"time"
	"unsafe"
)

func ssrMultiConfig(t *testing.T, udp ...bool) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "soga.conf")
	text := "type=sspanel-uim\nmu_suffix=example.test\nmu_regex=%16m%id.%suffix\nlisten=127.0.0.1\naccess_log_enable=false\nlog_file_dir=\nban_private_ip=false\nforbidden_bit_torrent=false\nproxy_protocol=false\nudp_proxy_protocol=false\ndy_limit_enable=false\nsoga_key=\n"
	if len(udp) > 0 && udp[0] {
		text += "ss_obfs_udp=true\n"
	}
	if err := os.WriteFile(path, []byte(text), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(path)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
}
func ssrMultiPanel(id int64, password string) *L.Aga91y {
	p := ssrPanelFixture(5000+int(id), "none", password)
	p.WZ89pTAOrHK4 = id
	return p
}
func ssrMultiHost(id int64, password string) string {
	sum := md5.Sum([]byte(fmt.Sprint(id) + password + "none" + "plain" + "origin"))
	return hex.EncodeToString(sum[:])[:16] + fmt.Sprint(id) + ".example.test"
}
func ssrMultiCore(t *testing.T, mode int) (*Ze0fOaAko, *Q.D5gTIZ5aVg) {
	t.Helper()
	carrier := ssrPanelFixture(0, "aes-128-cfb", "shared-carrier")
	carrier.WZ89pTAOrHK4 = 999
	cred := carrier.JAfXWwzYrW.(*L.QcqpSL3i6ri)
	cred.BskzckkE7qI = L.LaTpsgYQ7(mode)
	if mode == 1 {
		cred.ZZIDsJ = "http_simple"
	} else {
		cred.H_y2xZht7S8 = "auth_aes128_md5"
	}
	node := &L.Wr5SVP{EMlXEc: 69, ASJxBbFS: L.Dv6QJyFW6bU{QawbKS2x: true}}
	// Original RTTI's unexported carrier slice at432; no added production API.
	*(*[]*L.Aga91y)(unsafe.Add(unsafe.Pointer(node), 432)) = []*L.Aga91y{carrier}
	data := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{ssrMultiPanel(42, "alpha"), ssrMultiPanel(43, "bravo")}}
	raw, err := x4oEZV(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: data})
	if err != nil {
		t.Fatal(err)
	}
	v := raw.(*Ze0fOaAko)
	t.Cleanup(v.Stop)
	if err = v.Start(); err != nil {
		t.Fatal(err)
	}
	return v, data
}
func ssrMultiAdapter(t *testing.T, v *Ze0fOaAko, mode int, udp bool, id int64, password, carrier string) *A.RpNX5wlEsF {
	t.Helper()
	state := v.mc0MXUU.sSKJ79jFt[0]
	port := state.woqAbWjQH_A[0].Addr().(*net.TCPAddr).Port
	if udp {
		port = state.aaBpAQ44Zfs[0].LocalAddr().(*net.UDPAddr).Port
	}
	opt := A.ShadowsocksROptions{Name: "multiuser-wire", Server: "127.0.0.1", Port: port, Cipher: "aes-128-cfb", Password: carrier, Obfs: "plain", Protocol: "auth_aes128_md5", ProtocolParam: fmt.Sprintf("%d:%s", id, password), UDP: true}
	if mode == 1 {
		opt.Obfs = "http_simple"
		opt.ObfsParam = ssrMultiHost(id, password)
		opt.Protocol = "origin"
		opt.ProtocolParam = ""
	}
	a, err := A.CR6bC5(opt)
	if err != nil {
		t.Fatal(err)
	}
	return a
}
func ssrMultiTCP(t *testing.T, a *A.RpNX5wlEsF, port int) net.Conn {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c, err := a.DialContext(ctx, &E.D0mLgQF{ZouEpmjw8: net.IPv4(127, 0, 0, 1), JRXMVvg: E.AN18DDo(port)})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Close() })
	return c
}
func ssrMultiTCPExchange(c net.Conn) error {
	c.SetDeadline(time.Now().Add(350 * time.Millisecond))
	if _, err := c.Write([]byte("ping")); err != nil {
		return err
	}
	p := make([]byte, 5)
	if _, err := io.ReadFull(c, p); err != nil {
		return err
	}
	if string(p) != "Aping" {
		return fmt.Errorf("unexpected echo %q", p)
	}
	return nil
}
func ssrMultiExpect(t *testing.T, c net.Conn, accept bool) {
	t.Helper()
	err := ssrMultiTCPExchange(c)
	if accept && err != nil {
		t.Fatal(err)
	}
	if !accept && err == nil {
		t.Fatal("rejected credential still transfers")
	}
}
func ssrMultiUDP(t *testing.T, a *A.RpNX5wlEsF) net.PacketConn {
	t.Helper()
	c, err := a.ListenPacketContext(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Close() })
	return c
}
func ssrMultiUDPExpect(t *testing.T, c net.PacketConn, port int, accept bool) {
	t.Helper()
	c.SetDeadline(time.Now().Add(350 * time.Millisecond))
	target := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: port}
	_, err := c.WriteTo([]byte("ping"), target)
	if err == nil {
		p := make([]byte, 1024)
		var n int
		var from net.Addr
		n, from, err = c.ReadFrom(p)
		if err == nil && (string(p[:n]) != "Aping" || from.String() != target.String()) {
			t.Fatalf("UDP wrong target/payload %v %x", from, p[:n])
		}
	}
	if accept && err != nil {
		t.Fatal(err)
	}
	if !accept && err == nil {
		t.Fatal("rejected UDP credential still transfers")
	}
}
func TestNativeSSRSinglePortMultiuserLiveLifecycle(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrMultiConfig(t)
	for _, mode := range []int{1, 2} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			v, d := ssrMultiCore(t, mode)
			state := v.mc0MXUU.sSKJ79jFt[0]
			target := liveTarget(t, "127.0.0.1", 0, 'A')
			if mode == 1 && len(state.aaBpAQ44Zfs) != 0 {
				t.Fatal("mode1 default must not invent UDP")
			}
			a := ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 42, "alpha", "shared-carrier"), target)
			b := ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 43, "bravo", "shared-carrier"), target)
			ssrMultiExpect(t, a, true)
			ssrMultiExpect(t, b, true)
			var ua, ub net.PacketConn
			if mode == 2 {
				ua = ssrMultiUDP(t, ssrMultiAdapter(t, v, mode, true, 42, "alpha", "shared-carrier"))
				ub = ssrMultiUDP(t, ssrMultiAdapter(t, v, mode, true, 43, "bravo", "shared-carrier"))
				ssrMultiUDPExpect(t, ua, target, true)
				ssrMultiUDPExpect(t, ub, target, true)
			}
			// Invalid host/UID credentials must be rejected by real native admission.
			ssrMultiExpect(t, ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 44, "unknown", "shared-carrier"), target), false)
			d.VTme1SuHL6i_ = L.Sxcf3wu3L{ssrMultiPanel(42, "alpha-new"), ssrMultiPanel(43, "bravo")}
			if v.UpdateAPIData(d) {
				t.Fatal("node restart")
			}
			if v.mc0MXUU.sSKJ79jFt[0] != state {
				t.Fatal("ordinary user update restarted shared listener")
			}
			ssrMultiExpect(t, a, false)
			ssrMultiExpect(t, b, true)
			ssrMultiExpect(t, ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 42, "alpha", "shared-carrier"), target), false)
			a2 := ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 42, "alpha-new", "shared-carrier"), target)
			ssrMultiExpect(t, a2, true)
			if mode == 2 {
				ssrMultiUDPExpect(t, ua, target, false)
				ssrMultiUDPExpect(t, ub, target, true)
				ssrMultiUDPExpect(t, ssrMultiUDP(t, ssrMultiAdapter(t, v, mode, true, 42, "alpha-new", "shared-carrier")), target, true)
			}
			d.VTme1SuHL6i_ = L.Sxcf3wu3L{ssrMultiPanel(42, "alpha-new")}
			v.UpdateAPIData(d)
			ssrMultiExpect(t, b, false)
			ssrMultiExpect(t, a2, true)
			if mode == 2 {
				ssrMultiUDPExpect(t, ub, target, false)
			}
			ssrMultiExpect(t, ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 43, "bravo", "shared-carrier"), target), false)
			// Carrier replacement retires accepts but port-context is ignored by the
			// already authenticated TCP handler; ordinary-user streams remain valid.
			carrier := ssrPanelFixture(0, "aes-128-cfb", "new-carrier")
			carrier.WZ89pTAOrHK4 = 999
			c := carrier.JAfXWwzYrW.(*L.QcqpSL3i6ri)
			c.BskzckkE7qI = L.LaTpsgYQ7(mode)
			if mode == 1 {
				c.ZZIDsJ = "http_simple"
			} else {
				c.H_y2xZht7S8 = F.Mu0vMIIs("auth_aes128_md5")
			}
			oldAddr := state.woqAbWjQH_A[0].Addr().String()
			*(*[]*L.Aga91y)(unsafe.Add(unsafe.Pointer(d.H47AzZsYHB1N), 432)) = []*L.Aga91y{carrier}
			v.UpdateAPIData(d)
			replacement := v.mc0MXUU.sSKJ79jFt[0]
			if replacement == state || state.ot1kuknWf35.Err() != context.Canceled {
				t.Fatal("carrier did not retire old port")
			}
			if raw, err := net.DialTimeout("tcp", oldAddr, 100*time.Millisecond); err == nil {
				raw.Close()
				t.Fatal("retired TCP listener accepts")
			}
			ssrMultiExpect(t, a2, true)
			a3 := ssrMultiTCP(t, ssrMultiAdapter(t, v, mode, false, 42, "alpha-new", "new-carrier"), target)
			ssrMultiExpect(t, a3, true)
			d.VTme1SuHL6i_ = nil
			v.UpdateAPIData(d)
			ssrMultiExpect(t, a2, false)
			ssrMultiExpect(t, a3, false)
			if v.mc0MXUU.sSKJ79jFt[0] != replacement {
				t.Fatal("removing all users removed carrier")
			}
			addr := replacement.woqAbWjQH_A[0].Addr().String()
			v.Stop()
			if replacement.ot1kuknWf35.Err() != context.Canceled {
				t.Fatal("Stop did not cancel port")
			}
			if raw, err := net.DialTimeout("tcp", addr, 100*time.Millisecond); err == nil {
				raw.Close()
				t.Fatal("stopped listener accepts")
			}
		})
	}
}

// Test-only independent client arithmetic. It does not call any Roa/ad5
// protocol, list or padding implementation; original machine fixtures below
// separately validate this reference before it is used as a wire oracle.
type ssrPaddingRNG struct{ a, b uint64 }

func (r *ssrPaddingRNG) next() uint64 {
	a, b := r.a, r.b
	a ^= a << 23
	r.a = b
	r.b = a ^ b ^ (a >> 17) ^ (b >> 26)
	return r.b + b
}
func ssrPaddingSeed(seed []byte) ssrPaddingRNG {
	var b [16]byte
	copy(b[:], seed)
	return ssrPaddingRNG{binary.LittleEndian.Uint64(b[:8]), binary.LittleEndian.Uint64(b[8:])}
}
func ssrPaddingList(protocol byte, key []byte, period uint64) []int {
	seed := bytes.Clone(key)
	if protocol == 'f' {
		for i := 0; i < 8; i++ {
			seed[i] ^= byte(period >> uint(56-8*i))
		}
	}
	r := ssrPaddingSeed(seed)
	list := make([]int, int(r.next()%24)+12)
	for i := range list {
		list[i] = int(r.next() % 2340 % 2040 % 1440)
	}
	sort.Ints(list)
	if protocol != 'c' {
		for list[len(list)-1] < 1300 && len(list) < 64 {
			list = append(list, int(r.next()%2340%2040%1440))
		}
		sort.Ints(list)
	}
	return list
}
func ssrPaddingSize(protocol byte, list []int, length int, seed []byte) (int, ssrPaddingRNG) {
	target := length + 4
	if protocol == 'd' && list[len(list)-1] <= target {
		return 0, ssrPaddingRNG{}
	}
	var b [16]byte
	copy(b[:], seed)
	binary.LittleEndian.PutUint16(b[:], uint16(length))
	r := ssrPaddingSeed(b[:])
	for i := 0; i < 4; i++ {
		r.next()
	}
	if list[len(list)-1] <= target {
		if protocol != 'c' || target >= 1440 {
			return 0, r
		}
		modulus := uint64(1021)
		if target > 1300 {
			modulus = 31
		} else if target > 900 {
			modulus = 127
		} else if target > 400 {
			modulus = 521
		}
		return int(r.next() % modulus), r
	}
	index := sort.SearchInts(list, target)
	if protocol == 'c' || protocol == 'd' {
		index += int(r.next() % uint64(len(list)-index))
	}
	return list[index] - target, r
}
func TestSSRIndependentPaddingOriginalFixtures(t *testing.T) {
	key := make([]byte, 16)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for _, tc := range []struct {
		protocol byte
		list     []int
	}{
		{'c', []int{36, 45, 116, 172, 255, 366, 472, 557, 590, 715, 782, 797, 935, 1112, 1163}},
		{'d', []int{36, 45, 116, 172, 222, 255, 289, 311, 344, 366, 472, 483, 557, 590, 692, 715, 782, 782, 797, 876, 935, 1112, 1163, 1209, 1397}},
	} {
		if got := ssrPaddingList(tc.protocol, key, 0); fmt.Sprint(got) != fmt.Sprint(tc.list) {
			t.Fatalf("original list %c: %v", tc.protocol, got)
		}
	}
	// Original C fallback traces at length branch boundaries, with list {0}.
	for _, tc := range []struct{ n, want int }{{0, 538}, {1, 528}, {396, 857}, {397, 200}, {896, 370}, {897, 89}, {1296, 88}, {1297, 1}, {1435, 14}, {1436, 0}} {
		got, _ := ssrPaddingSize('c', []int{0}, tc.n, key)
		if got != tc.want {
			t.Fatalf("original pad %d: %d != %d", tc.n, got, tc.want)
		}
	}
}

type ssrPaddedWire struct {
	c                              net.Conn
	protocol                       byte
	key, requestSeed, responseSeed []byte
	list                           []int
	enc, dec                       *rc4.Cipher
	sent, received                 uint32
	firstResponse                  bool
	requestPads, responsePads      []int
}

func ssrPaddedClient(t *testing.T, endpoint, protocol string) *ssrPaddedWire {
	t.Helper()
	c, err := net.Dial("tcp", endpoint)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Close() })
	c.SetDeadline(time.Now().Add(2 * time.Second))
	key := ssrIndependentKey("native-matrix-key", 16)
	header := make([]byte, 36)
	rand.Read(header[:4])
	recvSeed := ssrChainMAC(key, header[:4])
	copy(header[4:12], recvSeed[:8])
	binary.LittleEndian.PutUint32(header[12:16], 42^binary.LittleEndian.Uint32(recvSeed[8:12]))
	auth := make([]byte, 16)
	binary.LittleEndian.PutUint32(auth, uint32(time.Now().Unix()))
	rand.Read(auth[4:12])
	binary.LittleEndian.PutUint16(auth[12:], 4)
	block, _ := aes.NewCipher(ssrIndependentKey(base64.StdEncoding.EncodeToString(key)+protocol, 16))
	cipher.NewCBCEncrypter(block, make([]byte, 16)).CryptBlocks(header[16:32], auth)
	sendSeed := ssrChainMAC(key, header[12:32])
	copy(header[32:], sendSeed[:4])
	sessionKey := ssrIndependentKey(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(recvSeed), 16)
	enc, _ := rc4.NewCipher(sessionKey)
	dec, _ := rc4.NewCipher(sessionKey)
	if _, err = c.Write(header); err != nil {
		t.Fatal(err)
	}
	variant := protocol[len(protocol)-1]
	period := (uint64(uint32(time.Now().Unix())) >> 1) * 3257812231 >> 47
	return &ssrPaddedWire{c: c, protocol: variant, key: key, requestSeed: recvSeed, responseSeed: sendSeed, list: ssrPaddingList(variant, key, period), enc: enc, dec: dec, sent: 1, received: 1, firstResponse: true}
}
func (c *ssrPaddedWire) send(t *testing.T, p []byte) {
	t.Helper()
	pad, r := ssrPaddingSize(c.protocol, c.list, len(p), c.requestSeed)
	c.requestPads = append(c.requestPads, pad)
	offset := 0
	if len(p) > 0 && pad > 0 {
		offset = int(r.next() % 0x200000011 % uint64(pad))
	}
	frame := make([]byte, len(p)+pad+4)
	rand.Read(frame)
	binary.LittleEndian.PutUint16(frame, uint16(len(p))^binary.LittleEndian.Uint16(c.requestSeed[14:]))
	c.enc.XORKeyStream(frame[2+offset:], p)
	key := binary.LittleEndian.AppendUint32(bytes.Clone(c.key), c.sent)
	c.sent++
	mac := ssrChainMAC(key, frame[:len(frame)-2])
	copy(frame[len(frame)-2:], mac[:2])
	c.requestSeed = mac
	if _, err := c.c.Write(frame); err != nil {
		t.Fatal(err)
	}
}
func (c *ssrPaddedWire) receive(t *testing.T) []byte {
	t.Helper()
	var header [2]byte
	if _, err := io.ReadFull(c.c, header[:]); err != nil {
		t.Fatal("frame header", err)
	}
	length := int(binary.LittleEndian.Uint16(header[:]) ^ binary.LittleEndian.Uint16(c.responseSeed[14:]))
	if length > 4096 {
		t.Fatalf("decoded frame length %d", length)
	}
	pad, r := ssrPaddingSize(c.protocol, c.list, length, c.responseSeed)
	c.responsePads = append(c.responsePads, pad)
	body := make([]byte, length+pad+2)
	if _, err := io.ReadFull(c.c, body); err != nil {
		t.Fatal("frame body", length, pad, err)
	}
	key := binary.LittleEndian.AppendUint32(bytes.Clone(c.key), c.received)
	c.received++
	mac := ssrChainMAC(key, header[:], body[:len(body)-2])
	if !bytes.Equal(body[len(body)-2:], mac[:2]) {
		t.Fatal("response HMAC")
	}
	c.responseSeed = mac
	offset := 0
	if length > 0 && pad > 0 {
		offset = int(r.next() % 0x200000011 % uint64(pad))
	}
	plain := make([]byte, length)
	c.dec.XORKeyStream(plain, body[offset:offset+length])
	if c.firstResponse {
		if len(plain) < 2 || binary.LittleEndian.Uint16(plain) != 1460 {
			t.Fatal("MSS prefix", plain)
		}
		c.firstResponse = false
		plain = plain[2:]
	}
	return plain
}
func TestNativeSSRIndependentSmallPaddedChains(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, variant := range []byte{'c', 'd', 'e', 'f'} {
		t.Run(string(variant), func(t *testing.T) {
			protocol := "auth_chain_" + string(variant)
			v := ssrNativeMatrixCore(t, "none", protocol, "plain")
			listener, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			defer listener.Close()
			done := make(chan error, 1)
			go func() {
				c, err := listener.Accept()
				if err != nil {
					done <- err
					return
				}
				defer c.Close()
				c.SetDeadline(time.Now().Add(3 * time.Second))
				_, err = io.Copy(c, c)
				done <- err
			}()
			client := ssrPaddedClient(t, v.mc0MXUU.sSKJ79jFt[0].woqAbWjQH_A[0].Addr().String(), protocol)
			address := ssFixtureAddress(1, listener.Addr().(*net.TCPAddr).Port)
			// Cross list and fallback boundaries on one authenticated stream, advancing
			// both RC4 streams, MAC chains and independent padding PRNG per frame.
			for i, size := range []int{1, 16, 96, 396, 397, 896, 897, 1296, 1297, 1435, 1436, 1440} {
				payload := make([]byte, size)
				for j := range payload {
					payload[j] = byte(j*31 + i)
				}
				request := bytes.Clone(payload)
				if i == 0 {
					request = append(bytes.Clone(address), request...)
				}
				client.send(t, request)
				var got []byte
				for len(got) < len(payload) {
					got = append(got, client.receive(t)...)
				}
				if !bytes.Equal(got, payload) {
					t.Fatalf("payload boundary %d differs (%d)", size, len(got))
				}
			}
			requestNonzero, responseNonzero := false, false
			for _, p := range client.requestPads {
				requestNonzero = requestNonzero || p > 0
			}
			for _, p := range client.responsePads {
				responseNonzero = responseNonzero || p > 0
			}
			if !requestNonzero || !responseNonzero {
				t.Fatalf("missing nonzero padding: request%v response%v", client.requestPads, client.responsePads)
			}
			t.Logf("request padding%v; response padding%v", client.requestPads, client.responsePads)
			client.c.Close()
			if err := <-done; err != nil {
				t.Fatal("echo target", err)
			}
		})
	}
}

// Optional mode1 UDP does not carry HTTP Host authentication: the original
// receiveSSRDatagrams falls back to the shared carrier user when ReadFromProxy
// returns nil user. Preserve and test this distinction from mode2 UID auth.
func TestNativeSSRSinglePortMode1OptionalUDP(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrMultiConfig(t, true)
	v, d := ssrMultiCore(t, 1)
	state := v.mc0MXUU.sSKJ79jFt[0]
	if len(state.aaBpAQ44Zfs) == 0 {
		t.Fatal("explicit mode1 UDP not started")
	}
	target := liveTarget(t, "127.0.0.1", 0, 'A')
	a := ssrMultiUDP(t, ssrMultiAdapter(t, v, 1, true, 42, "alpha", "shared-carrier"))
	unknown := ssrMultiUDP(t, ssrMultiAdapter(t, v, 1, true, 44, "unknown", "shared-carrier"))
	ssrMultiUDPExpect(t, a, target, true)
	ssrMultiUDPExpect(t, unknown, target, true)
	d.VTme1SuHL6i_ = nil
	v.UpdateAPIData(d)
	ssrMultiUDPExpect(t, a, target, true)
	ssrMultiUDPExpect(t, unknown, target, true)
	// Incorrect shared cipher password still fails; ordinary Host/UID fields
	// were simply never present on this protocol's UDP wire.
	bad := ssrMultiUDP(t, ssrMultiAdapter(t, v, 1, true, 44, "unknown", "wrong-carrier"))
	ssrMultiUDPExpect(t, bad, target, false)
	v.Stop()
	ssrMultiUDPExpect(t, a, target, false)
}
