package sogav1node_test

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"io"
	"net"
	"os"
	"path/filepath"
	_ "recovered.local/soga/panel/xboardnode"
	L "recovered.local/soga/panel/types"
	_ "recovered.local/soga/panel/sogav1node"
	_ "recovered.local/soga/panel/xiaov2boardnode"
	C "recovered.local/soga/config"
	_ "recovered.local/soga/panel/ppanelnode"
	O "recovered.local/soga/db"
	Q "recovered.local/soga/panel/registry"
	"strings"
	"sync"
	"testing"
	"time"
)

// This independent peer speaks MySQL packets over real loopback TCP. Production
// SQL DB/driver/Query/Scan/Exec remain intact. The driver dial hook only redirects
// the original bracketed []net.IP address to this ephemeral fixture listener.
type mysqlPeer struct {
	t                  *testing.T
	mu                 sync.Mutex
	queries, addresses []string
	mode               string
	execFailures       int
	listener           net.Listener
	wg                 sync.WaitGroup
}

func (p *mysqlPeer) read(c net.Conn) ([]byte, error) {
	var h [4]byte
	if _, e := io.ReadFull(c, h[:]); e != nil {
		return nil, e
	}
	n := int(h[0]) | int(h[1])<<8 | int(h[2])<<16
	b := make([]byte, n)
	_, e := io.ReadFull(c, b)
	return b, e
}
func (p *mysqlPeer) write(c net.Conn, seq byte, b []byte) error {
	out := append([]byte{byte(len(b)), byte(len(b) >> 8), byte(len(b) >> 16), seq}, b...)
	_, e := c.Write(out)
	return e
}
func mysqlString(s string) []byte {
	if len(s) >= 251 {
		panic("fixture short string")
	}
	return append([]byte{byte(len(s))}, []byte(s)...)
}
func (p *mysqlPeer) fail(c net.Conn, seq byte, msg string) error {
	return p.write(c, seq, append([]byte{255, 0x51, 4, '#', 'H', 'Y', '0', '0', '0'}, msg...))
}
func (p *mysqlPeer) serve(c net.Conn) {
	defer p.wg.Done()
	defer c.Close()
	_ = c.SetDeadline(time.Now().Add(20 * time.Second))
	const capabilities uint32 = 0x200 | 0x8000 | 0x80000 | 8 | 0x2000
	h := append([]byte{10}, []byte("5.7.0-fixture\x00")...)
	h = append(h, 1, 0, 0, 0)
	h = append(h, []byte("12345678\x00")...)
	h = append(h, byte(capabilities&255), byte((capabilities>>8)&255), 33, 2, 0, byte((capabilities>>16)&255), byte((capabilities>>24)&255), 21)
	h = append(h, make([]byte, 10)...)
	h = append(h, []byte("abcdefghijkl\x00mysql_native_password\x00")...)
	if e := p.write(c, 0, h); e != nil {
		return
	}
	login, e := p.read(c)
	if e != nil {
		return
	}
	if len(login) < 33 || !strings.HasPrefix(string(login[32:]), "fixture-user\x00") {
		p.t.Errorf("bad login %x", login)
		return
	}
	if e = p.write(c, 2, []byte{0, 0, 0, 2, 0, 0, 0}); e != nil {
		return
	}
	for {
		b, e := p.read(c)
		if e != nil {
			return
		}
		if len(b) == 0 {
			return
		}
		switch b[0] {
		case 1:
			return
		case 14:
			if p.write(c, 1, []byte{0, 0, 0, 2, 0, 0, 0}) != nil {
				return
			}
			continue
		case 3:
		default:
			p.t.Errorf("unexpected command %d", b[0])
			return
		}
		query := string(b[1:])
		p.mu.Lock()
		p.queries = append(p.queries, query)
		mode := p.mode
		fail := p.execFailures > 0
		if strings.HasPrefix(query, "update ") && fail {
			p.execFailures--
		}
		p.mu.Unlock()
		if strings.HasPrefix(query, "update ") {
			if fail {
				e = p.fail(c, 1, "fixture exec failure")
			} else {
				e = p.write(c, 1, []byte{0, 1, 0, 2, 0, 0, 0})
			}
			if e != nil {
				return
			}
			continue
		}
		expected := "select id, u, d, transfer_enable, uuid from user where enable = 1 and u + d < transfer_enable"
		if strings.Contains(query, "speedlimitss") {
			expected = "select id, u, d, transfer_enable, uuid, speedlimitss from user where enable = 1 and u + d < transfer_enable"
		}
		if strings.Contains(query, "pid,") {
			expected = "select pid, u, d, transfer_enable, port, passwd from user where enable = 1 and u + d < transfer_enable"
		}
		if query != expected {
			p.t.Errorf("original SQL mismatch: %q", query)
		}
		if mode == "queryerror" {
			if p.fail(c, 1, "fixture query failure") != nil {
				return
			}
			continue
		}
		cols := []string{"id", "u", "d", "transfer_enable", "uuid"}
		values := []*string{}
		for _, s := range []string{"42", "3", "4", "1000", "fixture-secret"} {
			v := s
			values = append(values, &v)
		}
		if strings.Contains(query, "speedlimitss") {
			cols = append(cols, "speedlimitss")
			v := "8"
			values = append(values, &v)
		} else if strings.Contains(query, "pid,") {
			cols = []string{"pid", "u", "d", "transfer_enable", "port", "passwd"}
			port, pw := "4567", "fixture-secret"
			values[4] = &port
			values = append(values, &pw)
		}
		if mode == "emptysecret" {
			empty := ""
			at := 4
			if strings.Contains(query, "pid,") {
				at = 5
			}
			values[at] = &empty
		}
		if mode == "scanerror" {
			v := "not-an-integer"
			values[0] = &v
		}
		if mode == "null" {
			values[len(cols)-1] = nil
			if strings.Contains(query, "speedlimitss") {
				v := "8"
				values[len(cols)-1] = &v
				values[4] = nil
			}
		}
		seq := byte(1)
		if p.write(c, seq, []byte{byte(len(cols))}) != nil {
			return
		}
		seq++
		for _, col := range cols {
			var b []byte
			for _, s := range []string{"def", "fixture", "user", "user", col, col} {
				b = append(b, mysqlString(s)...)
			}
			b = append(b, 12, 33, 0, 0, 4, 0, 0, 253, 0, 0, 0, 0, 0)
			if p.write(c, seq, b) != nil {
				return
			}
			seq++
		}
		if p.write(c, seq, []byte{254, 0, 0, 2, 0}) != nil {
			return
		}
		seq++
		if mode != "empty" {
			var row []byte
			for _, v := range values {
				if v == nil {
					row = append(row, 251)
				} else {
					row = append(row, mysqlString(*v)...)
				}
			}
			if p.write(c, seq, row) != nil {
				return
			}
			seq++
		}
		if mode == "lateerror" {
			e = p.fail(c, seq, "fixture row-tail failure")
		} else {
			e = p.write(c, seq, []byte{254, 0, 0, 2, 0})
		}
		if e != nil {
			return
		}
	}
}
func TestFourOriginalDatabaseAdaptersOverMySQL(t *testing.T) {
	l, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	p := &mysqlPeer{t: t, listener: l}
	defer func() { _ = O.DB().Close(); l.Close(); p.wg.Wait() }()
	go func() {
		for {
			c, e := l.Accept()
			if e != nil {
				return
			}
			p.wg.Add(1)
			go p.serve(c)
		}
	}()
	mysql.RegisterDialContext("tcp", func(ctx context.Context, address string) (net.Conn, error) {
		p.mu.Lock()
		p.addresses = append(p.addresses, address)
		p.mu.Unlock()
		return (&net.Dialer{}).DialContext(ctx, "tcp", l.Addr().String())
	})
	defer mysql.RegisterDialContext("tcp", func(ctx context.Context, address string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "tcp", address)
	})
	path := filepath.Join(t.TempDir(), "soga.conf")
	base := "type=whmcs\nserver_type=ss\napi=db\nnode_id=71\nsoga_key=\ndb_host=127.0.0.1\ndb_port=3307\ndb_name=fixture\ndb_user=fixture-user\ndb_password=fixture-pass\nwhmcs_rate=1.5\nwhmcs_port=8443\nwhmcs_method=aes-128-gcm\nwhmcs_password=carrier\nwhmcs_protocol=auth_aes128_md5\nwhmcs_obfs=plain\nwhmcs_net=grpc\nwhmcs_tls=true\nwhmcs_alter_id=2\nwhmcs_ws_path=/ws\nwhmcs_service_name=service\nuser_speed_limit=12\nuser_conn_limit=4\nuser_tcp_limit=7\n"
	load := func(body string) {
		t.Helper()
		if e := os.WriteFile(path, []byte(body), 0600); e != nil {
			t.Fatal(e)
		}
		C.MkwNWciPHS(path)
		if e := C.TvBECnD9fM6(); e != nil {
			t.Fatal(e)
		}
	}
	load(base)
	for _, kind := range []C.Mpn2JltM{"ss", "ssr", "trojan", "v2ray"} {
		t.Run(string(kind), func(t *testing.T) {
			conf := &C.FCEAfQJwDA7e{ZyAhkq5th8bf: kind}
			f, ok := Q.PanelFactoriesAt5095E50["whmcs-"+string(kind)+"-db"]
			if !ok {
				t.Fatal("factory")
			}
			a := f(conf)
			if a.GetAPIConfig() != conf {
				t.Fatal("config identity")
			}
			p.mu.Lock()
			p.mode = "normal"
			p.mu.Unlock()
			out, e := a.GetAPIData()
			if e != nil {
				t.Fatal(e)
			}
			if len(out.VTme1SuHL6i_) != 1 {
				t.Fatal(out)
			}
			u := out.VTme1SuHL6i_[0]
			if u.WZ89pTAOrHK4 != 42 || u.XuBBcTThBW != 3 || u.UWsYsmytTi != 4 || u.IUlI_15Qf_ != 1000 || u.Xmrna8 != 4 || u.B4DYSxuA31za != 7 {
				t.Fatalf("user %+v", u)
			}
			if out.H47AzZsYHB1N.EMlXEc != 0 || out.HYnfYWa != nil {
				t.Fatal("original no node id / detector writes")
			}
			switch kind {
			case "ss":
				v := u.JAfXWwzYrW.(*L.YlfOV0DBBnx)
				if v.A_wNUTkMh6pH != "fixture-secret" || v.T1oDDIgYJHh != "aes-128-gcm" || u.LmOKwjha != 8 {
					t.Fatal(v, u)
				}
			case "ssr":
				v := u.JAfXWwzYrW.(*L.QcqpSL3i6ri)
				if v.YrjVl_2g_S != 4567 || v.YyAgnMatmj != "fixture-secret" || !out.H47AzZsYHB1N.ASJxBbFS.QawbKS2x {
					t.Fatal(v)
				}
			case "trojan":
				if u.JAfXWwzYrW.(*L.JxTGUar).ARARP1z0yA != "fixture-secret" || out.H47AzZsYHB1N.JhIYwEukjd.T4cuhG7MTbG != "tcp" {
					t.Fatal(u)
				}
			case "v2ray":
				v := u.JAfXWwzYrW.(*L.Yn8BJoL)
				if v.SfiEGpJo != "fixture-secret" || v.FvQ2ltKOM7 != 2 || out.H47AzZsYHB1N.SwxJNgPpg.F2GXyZZ != "grpc" {
					t.Fatal(v)
				}
			}
			for _, mode := range []string{"null", "empty", "emptysecret", "scanerror", "queryerror", "lateerror"} {
				p.mu.Lock()
				p.mode = mode
				p.mu.Unlock()
				v, e := a.GetAPIData()
				switch mode {
				case "scanerror", "queryerror":
					if v != nil || e == nil {
						t.Fatal(mode, v, e)
					}
				case "null", "empty":
					if e != nil || v.VTme1SuHL6i_ == nil || len(v.VTme1SuHL6i_) != 0 {
						t.Fatal(mode, v, e)
					}
				case "lateerror", "emptysecret":
					if e != nil || len(v.VTme1SuHL6i_) != 1 {
						t.Fatal("original ignores Rows.Err", v, e)
					}
				}
			}
			p.mu.Lock()
			untouched := len(p.queries)
			p.mu.Unlock()
			if e = a.SubmitData(nil); e != nil {
				t.Fatal(e)
			}
			if e = a.SubmitData(&Q.FGmNba4MhioY{}); e != nil {
				t.Fatal(e)
			}
			p.mu.Lock()
			p.mode = "normal"
			if len(p.queries) != untouched {
				t.Error("empty submit queried database")
			}
			before := len(p.queries)
			p.execFailures = 1
			p.mu.Unlock()
			traffic := make(L.CWfcUv, 5001)
			for i := range traffic {
				traffic[i] = &L.DLxslMwYJXQw{Hy7Q97RROK: int64(i + 1), XJkGTQf_qkAt: 3, I6LXObODav: 5}
			}
			if e = a.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: traffic}); e == nil || !strings.Contains(e.Error(), "fixture exec failure") {
				t.Fatal(e)
			}
			p.mu.Lock()
			queries := append([]string(nil), p.queries[before:]...)
			p.mu.Unlock()
			if len(queries) != 2 {
				t.Fatal("batch failure must continue", len(queries))
			}
			col := "id"
			if kind == "ssr" {
				col = "pid"
			}
			if !strings.HasPrefix(queries[0], "update user set u = case "+col+"\nwhen 1 then u + 4\n") || !strings.Contains(queries[0], "when 1 then d + 7") || !strings.Contains(queries[1], "where "+col+" in(5001)") {
				t.Fatal("SQL batch/rate shape")
			}
		})
	}
	p.mu.Lock()
	for _, a := range p.addresses {
		if a != "[127.0.0.1]:3307" {
			t.Error(a)
		}
	}
	p.mu.Unlock()
	for _, kind := range []string{"ss", "trojan", "v2ray"} {
		if Q.PanelFactoriesAt5095E50["whmcs-"+kind+"-webapi"] == nil {
			t.Fatal("web factory", kind)
		}
	}
	if Q.PanelFactoriesAt5095E50["whmcs-ssr-webapi"] != nil {
		t.Fatal("invented SSR WebAPI")
	}
	// DB original validation order: SS accepts port0; other three reject it before query.
	load(strings.Replace(base, "whmcs_port=8443", "whmcs_port=0", 1))
	for _, kind := range []string{"ssr", "trojan", "v2ray"} {
		a := Q.PanelFactoriesAt5095E50["whmcs-"+kind+"-db"](&C.FCEAfQJwDA7e{})
		if _, e := a.GetAPIData(); e == nil || e.Error() != "please set whmcs_port" {
			t.Fatal(kind, e)
		}
	}
	ss := Q.PanelFactoriesAt5095E50["whmcs-ss-db"](&C.FCEAfQJwDA7e{}).(interface{ GetNode() (*L.Wr5SVP, error) })
	if n, e := ss.GetNode(); e != nil || n.FnMUny.Am3QoeF != 0 {
		t.Fatal("original SS accepts zero port", n, e)
	}
	for _, field := range []string{"password", "method", "protocol", "obfs"} {
		// Empty one field while all preceding validations retain valid values.
		lines := strings.Split(base, "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "whmcs_"+field+"=") {
				lines[i] = "whmcs_" + field + "="
			}
		}
		load(strings.Join(lines, "\n"))
		ssr := Q.PanelFactoriesAt5095E50["whmcs-ssr-db"](&C.FCEAfQJwDA7e{}).(interface{ GetNode() (*L.Wr5SVP, error) })
		if _, e := ssr.GetNode(); e == nil || e.Error() != "please set whmcs_"+field {
			t.Fatal(field, e)
		}
	}

}
