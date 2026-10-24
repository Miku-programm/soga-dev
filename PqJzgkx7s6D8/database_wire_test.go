package PqJzgkx7s6D8_test

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io"
	"net"
	"os"
	"path/filepath"
	_ "recovered.local/soga/GgDwT38SMQf5"
	L "recovered.local/soga/LFWTTXF7JR"
	_ "recovered.local/soga/Se1QrdYD7Aa"
	C "recovered.local/soga/cNT24rG4F"
	_ "recovered.local/soga/mIfcoDkc6"
	O "recovered.local/soga/oXGarxahaig"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	_ "recovered.local/soga/uavUP6a"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

// This independent peer speaks MySQL packets over real loopback TCP. Production
// SQL DB/driver/Query/Scan/Exec remain intact. The driver dial hook only redirects
// the original bracketed []net.IP address to this ephemeral fixture listener.
type ssPanelMySQLPeer struct {
	t                  *testing.T
	mu                 sync.Mutex
	queries, addresses []string
	mode               string
	execFailures       int
	listener           net.Listener
	wg                 sync.WaitGroup
}

func (p *ssPanelMySQLPeer) read(c net.Conn) ([]byte, error) {
	var h [4]byte
	if _, e := io.ReadFull(c, h[:]); e != nil {
		return nil, e
	}
	n := int(h[0]) | int(h[1])<<8 | int(h[2])<<16
	b := make([]byte, n)
	_, e := io.ReadFull(c, b)
	return b, e
}
func (p *ssPanelMySQLPeer) write(c net.Conn, seq byte, b []byte) error {
	out := append([]byte{byte(len(b)), byte(len(b) >> 8), byte(len(b) >> 16), seq}, b...)
	_, e := c.Write(out)
	return e
}
func ssPanelMySQLString(s string) []byte {
	if len(s) >= 251 {
		panic("fixture short string")
	}
	return append([]byte{byte(len(s))}, []byte(s)...)
}
func (p *ssPanelMySQLPeer) fail(c net.Conn, seq byte, msg string) error {
	return p.write(c, seq, append([]byte{255, 0x51, 4, '#', 'H', 'Y', '0', '0', '0'}, msg...))
}
func (p *ssPanelMySQLPeer) serve(c net.Conn) {
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

		if !strings.HasPrefix(strings.ToLower(query), "select ") {
			if p.write(c, 1, []byte{0, 1, 0, 2, 0, 0, 0}) != nil {
				return
			}
			continue
		}
		_ = fail
		var cols []string
		var raw []string
		switch {
		case strings.HasPrefix(query, "select server, node_speedlimit, mu_only"):
			cols = []string{"server", "node_speedlimit", "mu_only"}
			server := "srv;soga_port=4567|port=4568"
			if mode == "v2ray" {
				server = "srv;7788;5;ws;tls;path=/vmess"
			}
			if mode == "trojan" {
				server = "srv;port=8899"
			}
			raw = []string{server, "8.5", "1"}
		case strings.HasPrefix(query, "select node_bandwidth_limit"):
			cols = []string{"node_bandwidth_limit", "node_bandwidth", "node_group", "node_class"}
			raw = []string{"10000", "10", "2", "3"}
		case strings.Contains(query, "information_schema.COLUMNS"):
			cols = []string{"COLUMN_NAME"}
			raw = []string{"uuid"}
		case strings.HasPrefix(query, "select id, passwd"):
			cols = []string{"id", "passwd", "u", "d", "transfer_enable", "node_speedlimit", "node_connector", "port", "method", "protocol", "protocol_param", "obfs", "obfs_param", "is_multi_user", "uuid"}
			method, protocol := "aes-128-gcm", "origin"
			if mode == "ssr" {
				method = "aes-128-cfb"
				protocol = "auth_chain_a"
			}
			raw = []string{"42", "pw", "3", "4", "1000", "6", "2", "2345", method, protocol, "", "plain", "", "1", "12345678-1234-1234-1234-123456789abc"}
		case strings.HasPrefix(query, "select id, regex"):
			cols = []string{"id", "regex", "type"}
			raw = []string{"3", "full:example.org", "1"}
		case strings.HasPrefix(query, "select traffic_rate"):
			cols = []string{"traffic_rate"}
			raw = []string{"1.5"}
		case strings.Contains(query, "information_schema.TABLES"):
			cols = []string{"table_name"}
			raw = []string{"user_traffic_log"}
		default:
			p.t.Errorf("unexpected SQL %q", query)
			_ = p.fail(c, 1, "unexpected query")
			continue
		}
		values := make([]*string, len(raw))
		for i := range raw {
			values[i] = &raw[i]
		}

		seq := byte(1)
		if p.write(c, seq, []byte{byte(len(cols))}) != nil {
			return
		}
		seq++
		for _, col := range cols {
			var b []byte
			for _, s := range []string{"def", "fixture", "user", "user", col, col} {
				b = append(b, ssPanelMySQLString(s)...)
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
					row = append(row, ssPanelMySQLString(*v)...)
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
func TestFourOriginalSSPanelDatabaseAdaptersOverMySQLWire(t *testing.T) {
	l, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	p := &ssPanelMySQLPeer{t: t, listener: l}
	acceptDone := make(chan struct{})
	go func() {
		defer close(acceptDone)
		for {
			c, e := l.Accept()
			if e != nil {
				return
			}
			p.wg.Add(1)
			go p.serve(c)
		}
	}()
	defer func() { O.DB().Close(); l.Close(); <-acceptDone; p.wg.Wait() }()
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
	oldPath := C.ConfigurationPath()
	defer C.MkwNWciPHS(oldPath)
	// The configured driver option selects the text protocol for this fixture;
	// all production Query/Scan/Exec, transactions and native driver remain live.
	body := "type=sspanel-uim\nserver_type=ss\napi=db\nnode_id=9\ndb_host=127.0.0.1\ndb_port=3307\ndb_name=fixture?interpolateParams=true\ndb_user=fixture-user\ndb_password=fixture-pass\n"
	if e = os.WriteFile(path, []byte(body), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(path)
	if e = C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	for _, protocol := range []C.Mpn2JltM{"ss", "ssr", "trojan", "v2ray"} {
		t.Run(string(protocol), func(t *testing.T) {
			p.mu.Lock()
			p.mode = string(protocol)
			p.mu.Unlock()
			conf := &C.FCEAfQJwDA7e{OS91X026uML: "sspanel-uim", ZyAhkq5th8bf: protocol, F3PCiaIra: "db", Yf1KqSqJpNI7: 9, XU27D_AD: "fixture"}
			api, e := Q.SMkpVNb(conf)
			if e != nil {
				t.Fatal(e)
			}
			if api.GetAPIConfig() != conf {
				t.Fatal("configuration identity")
			}
			typ := reflect.TypeOf(api).Elem()
			if typ.Size() != 8 {
				t.Fatal(typ, typ.Size())
			}
			d, e := api.GetAPIData()
			if e != nil {
				t.Fatal(e)
			}
			if len(d.VTme1SuHL6i_) != 1 || len(d.HYnfYWa) != 1 || !d.HYnfYWa[0].Test("example.org") {
				t.Fatal(d)
			}
			u := d.VTme1SuHL6i_[0]
			if u.WZ89pTAOrHK4 != 42 || u.XuBBcTThBW != 3 || u.UWsYsmytTi != 4 || u.IUlI_15Qf_ != 1000 {
				t.Fatal(u)
			}
			switch protocol {
			case "ss":
				if d.H47AzZsYHB1N.FnMUny.Am3QoeF != 2345 || d.H47AzZsYHB1N.FnMUny.VuhWqou != "pw" || u.JAfXWwzYrW.(*L.YlfOV0DBBnx).T1oDDIgYJHh != "aes-128-gcm" {
					t.Fatal(d.H47AzZsYHB1N.FnMUny, u)
				}
			case "ssr":
				n := d.H47AzZsYHB1N.ASJxBbFS
				if n.C33jbIzI != 4567 || n.PU1Di18 != 4568 || reflect.ValueOf(n).FieldByName("kJbZdr").Len() != 1 {
					t.Fatal(n)
				}
			case "trojan":
				if d.H47AzZsYHB1N.JhIYwEukjd.GbD_2PU != 8899 {
					t.Fatal(d.H47AzZsYHB1N.JhIYwEukjd)
				}
			case "v2ray":
				n := d.H47AzZsYHB1N.SwxJNgPpg
				if n.CKsOpayd != 7788 || n.VUwf1n9g != 5 || u.JAfXWwzYrW.(*L.Yn8BJoL).FvQ2ltKOM7 != 5 {
					t.Fatal(n, u)
				}
			}
			if e = api.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 42, XJkGTQf_qkAt: 100, I6LXObODav: 200}}, JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 42, PNtn9KAV: []string{"127.0.0.1"}}}, Y0ys4ma4Tno4: L.QmiwDVjV00lZ{{GbzQUT_4f: 42, CAgUUdrv9: 3}}}); e != nil {
				t.Fatal(e)
			}
		})
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.addresses) == 0 {
		t.Fatal("no real driver dial")
	}
	for _, a := range p.addresses {
		if a != "[127.0.0.1]:3307" {
			t.Error(a)
		}
	}
	for _, prefix := range []string{"select server,", "select node_bandwidth_limit", "select id, passwd", "select id, regex", "insert into ss_node_info", "update user set u", "insert into user_traffic_log", "update ss_node set node_bandwidth", "insert into ss_node_online_log", "insert into alive_ip", "insert into detect_log"} {
		n := 0
		for _, q := range p.queries {
			if strings.HasPrefix(q, prefix) {
				n++
			}
		}
		if n != 4 {
			t.Error(fmt.Sprintf("%s: got %d calls", prefix, n))
		}
	}
}
