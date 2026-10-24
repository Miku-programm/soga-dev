package MFFa9yT

import (
	"context"
	"net"
	G "recovered.local/soga/G6bHD5JA"
	J "recovered.local/soga/J1P4KNvfI_J"
	L "recovered.local/soga/LFWTTXF7JR"
	N "recovered.local/soga/nzATvajt"
	"reflect"
	"testing"
)

type sessionUserFixture struct {
	N.PgaNd3014U
	calls      []string
	alive, tcp bool
	ctx        context.Context
}

func (u *sessionUserFixture) AddAliveIp(a, b string) bool {
	u.calls = append(u.calls, "alive:"+a+":"+b)
	return u.alive
}
func (u *sessionUserFixture) DelAliveIp(a, b string)  { u.calls = append(u.calls, "delalive:"+a+":"+b) }
func (u *sessionUserFixture) AddTCP() bool            { u.calls = append(u.calls, "tcp"); return u.tcp }
func (u *sessionUserFixture) DelTCP()                 { u.calls = append(u.calls, "deltcp") }
func (u *sessionUserFixture) GetId() int64            { return 42 }
func (u *sessionUserFixture) GetCtx() context.Context { u.calls = append(u.calls, "ctx"); return u.ctx }
func (u *sessionUserFixture) RecordAliveIp(ip string) { u.calls = append(u.calls, "record:"+ip) }
func (u *sessionUserFixture) GetConnLimit() int       { return 1 }

type sessionConnFixture struct {
	net.Conn
	user   N.PgaNd3014U
	dest   *G.AHU2MRmjzQu
	kind   J.ZES6F5
	closes int
}

func (c *sessionConnFixture) Close() error {
	c.closes++
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}
func (c *sessionConnFixture) GetDestAddr() *G.AHU2MRmjzQu { return c.dest }
func (c *sessionConnFixture) GetUser() N.PgaNd3014U       { return c.user }
func (c *sessionConnFixture) NetworkType() J.ZES6F5       { return c.kind }

type limitSessionFixture struct {
	calls   *[]string
	accept  bool
	onStart func()
	mutate  *J.ZzkV7_5l
}

func (m *limitSessionFixture) StartUserConn(ctx context.Context, cancel context.CancelFunc, run func(), u N.PgaNd3014U, ip string) bool {
	*m.calls = append(*m.calls, "limit:"+ip)
	if m.onStart != nil {
		m.onStart()
	}
	if m.accept {
		run()
		if m.mutate != nil {
			m.mutate.C66t8X = "changed"
		}
	}
	return m.accept
}
func (m *limitSessionFixture) EndUserConn(u N.PgaNd3014U, ip string) {
	*m.calls = append(*m.calls, "end:"+ip)
}

func TestCentralSessionOriginalLimitsAndDeferredCleanup(t *testing.T) {
	old := sessionBlockedAt50C6C7D
	defer func() { sessionBlockedAt50C6C7D = old }()
	for _, tc := range []struct {
		name                        string
		alive, tcp, source, blocked bool
		kind                        J.ZES6F5
		want                        []string
	}{
		{"nil source", true, true, false, false, 2, nil},
		{"alive denied", false, true, true, false, 2, []string{"alive:masked:masked"}},
		{"tcp denied", true, false, true, false, 2, []string{"alive:masked:masked", "tcp", "delalive:masked:masked"}},
		{"unknown kind balanced", true, true, true, false, 2, []string{"alive:masked:masked", "tcp", "ctx", "record:masked", "deltcp", "delalive:masked:masked"}},
		{"early gate", true, true, true, true, 0, []string{"alive:masked:masked", "tcp", "deltcp", "delalive:masked:masked"}},
		{"MUX gate after record", true, true, true, true, 2, []string{"alive:masked:masked", "tcp", "ctx", "record:masked", "deltcp", "delalive:masked:masked"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sessionBlockedAt50C6C7D = tc.blocked
			u := &sessionUserFixture{alive: tc.alive, tcp: tc.tcp, ctx: context.Background()}
			b := &JWN154qJ4{zb8f3cHWW8v: context.Background()}
			b.wiy6P3.Store(&L.Wr5SVP{EMlXEc: 17})
			if err := b.init(); err != nil {
				t.Fatal(err)
			}
			addr, _ := G.U80FDt("203.0.113.9:443")
			s := &J.ZzkV7_5l{E3L0l8hZ: 19, PayAlbdv6: u, Uqzvq94e: "masked", C66t8X: "redis", YlxwftjCvz_: tc.kind, Iayb4V: addr}
			if tc.source {
				s.KeHzUwuXWh0 = net.IPv4(198, 51, 100, 1)
			}
			c := &sessionConnFixture{user: u, dest: addr, kind: tc.kind}
			b.xR518G_hBbKv(s, c)
			if c.closes != 1 || s.IcoyFqRBU != 17 || !reflect.DeepEqual(u.calls, tc.want) {
				t.Fatalf("close=%d node=%d calls=%v want=%v", c.closes, s.IcoyFqRBU, u.calls, tc.want)
			}
		})
	}
}

func TestCentralRedisRegistrationCleanupAndSecondGate(t *testing.T) {
	old := sessionBlockedAt50C6C7D
	defer func() { sessionBlockedAt50C6C7D = old }()
	for _, accept := range []bool{false, true} {
		sessionBlockedAt50C6C7D = false
		u := &sessionUserFixture{alive: true, tcp: true, ctx: context.Background()}
		addr, _ := G.U80FDt("203.0.113.9:443")
		s := &J.ZzkV7_5l{PayAlbdv6: u, Uqzvq94e: "masked", C66t8X: "redis", KeHzUwuXWh0: net.IPv4(198, 51, 100, 1), Iayb4V: addr}
		c := &sessionConnFixture{user: u, dest: addr}
		m := &limitSessionFixture{calls: &u.calls, accept: accept, mutate: s, onStart: func() { sessionBlockedAt50C6C7D = true }}
		b := &JWN154qJ4{zb8f3cHWW8v: context.Background(), af8f_J: m}
		b.wiy6P3.Store(&L.Wr5SVP{})
		if err := b.init(); err != nil {
			t.Fatal(err)
		}
		b.xR518G_hBbKv(s, c)
		want := []string{"alive:masked:masked", "tcp", "ctx", "limit:redis"}
		if accept {
			want = append(want, "record:masked", "end:changed")
		}
		want = append(want, "deltcp", "delalive:masked:masked")
		if c.closes != 1 || !reflect.DeepEqual(u.calls, want) {
			t.Fatalf("accept=%v %v want %v", accept, u.calls, want)
		}
	}
}
