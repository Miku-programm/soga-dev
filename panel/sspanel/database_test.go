package sspanel

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
)

// A controlled SQL driver records arguments and injects database failures.
// This verifies call/transaction semantics, not compatibility with a live server.
type sqlCall struct {
	query string
	args  []driver.NamedValue
}
type sqlFixture struct {
	mu                 sync.Mutex
	calls              []sqlCall
	query              func(string, []driver.NamedValue) (driver.Rows, error)
	exec               func(string, []driver.NamedValue) error
	commitErr          error
	commits, rollbacks int
	ended              chan struct{}
	beginErr           error
}
type fixtureConnector struct{ f *sqlFixture }

func (c fixtureConnector) Connect(context.Context) (driver.Conn, error) {
	return &fixtureConn{f: c.f}, nil
}
func (c fixtureConnector) Driver() driver.Driver { return fixtureDriver{c.f} }

type fixtureDriver struct{ f *sqlFixture }

func (d fixtureDriver) Open(string) (driver.Conn, error) { return &fixtureConn{f: d.f}, nil }

type fixtureConn struct{ f *sqlFixture }

func (c *fixtureConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("unexpected prepare")
}
func (c *fixtureConn) Close() error { return nil }
func (c *fixtureConn) Begin() (driver.Tx, error) {
	c.f.record("BEGIN", nil)
	if c.f.beginErr != nil {
		return nil, c.f.beginErr
	}
	return &fixtureTx{c.f}, nil
}
func (c *fixtureConn) QueryContext(_ context.Context, q string, a []driver.NamedValue) (driver.Rows, error) {
	c.f.record(q, a)
	if c.f.query == nil {
		return nil, fmt.Errorf("unexpected query %s", q)
	}
	return c.f.query(q, a)
}
func (c *fixtureConn) ExecContext(_ context.Context, q string, a []driver.NamedValue) (driver.Result, error) {
	c.f.record(q, a)
	if c.f.exec != nil {
		if e := c.f.exec(q, a); e != nil {
			return nil, e
		}
	}
	return driver.RowsAffected(1), nil
}
func (f *sqlFixture) record(q string, a []driver.NamedValue) {
	f.mu.Lock()
	f.calls = append(f.calls, sqlCall{q, append([]driver.NamedValue(nil), a...)})
	f.mu.Unlock()
}

type fixtureTx struct{ f *sqlFixture }

func (x *fixtureTx) Commit() error {
	x.f.record("COMMIT", nil)
	x.f.mu.Lock()
	x.f.commits++
	x.f.mu.Unlock()
	if x.f.ended != nil {
		x.f.ended <- struct{}{}
	}
	return x.f.commitErr
}
func (x *fixtureTx) Rollback() error {
	x.f.record("ROLLBACK", nil)
	x.f.mu.Lock()
	x.f.rollbacks++
	x.f.mu.Unlock()
	if x.f.ended != nil {
		x.f.ended <- struct{}{}
	}
	return nil
}

type fixtureRows struct {
	names []string
	data  [][]driver.Value
	tail  error
}

func (r *fixtureRows) Columns() []string { return r.names }
func (r *fixtureRows) Close() error      { return nil }
func (r *fixtureRows) Next(v []driver.Value) error {
	if len(r.data) == 0 {
		if r.tail != nil {
			return r.tail
		}
		return io.EOF
	}
	copy(v, r.data[0])
	r.data = r.data[1:]
	return nil
}
func rowsOf(n int, data ...[]driver.Value) *fixtureRows {
	names := make([]string, n)
	for i := range names {
		names[i] = fmt.Sprint(i)
	}
	return &fixtureRows{names: names, data: data}
}
func withDatabase(t *testing.T, f *sqlFixture) (*N3owFvm, *sql.DB) {
	t.Helper()
	db := sql.OpenDB(fixtureConnector{f})
	old := database
	database = func() *sql.DB { return db }
	oldC, oldU, oldT, oldA := columnsCheckedAt50C6CC6, uuidColumnAt50C6CC7, trafficTableCheckedAt50C6CC8, trafficTableAvailableAt3787336
	columnsCheckedAt50C6CC6 = false
	uuidColumnAt50C6CC7 = false
	trafficTableCheckedAt50C6CC8 = false
	trafficTableAvailableAt3787336 = true
	t.Cleanup(func() {
		db.Close()
		database = old
		columnsCheckedAt50C6CC6, uuidColumnAt50C6CC7, trafficTableCheckedAt50C6CC8, trafficTableAvailableAt3787336 = oldC, oldU, oldT, oldA
	})
	return &N3owFvm{TMfNUn: &C.FCEAfQJwDA7e{Yf1KqSqJpNI7: 9, XU27D_AD: "schema", ZyAhkq5th8bf: "ss"}}, db
}
func waitEvent(t *testing.T, ch <-chan struct{}) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(3 * time.Second):
		t.Fatal("SQL event did not finish")
	}
}
func argsValues(a []driver.NamedValue) []any {
	v := make([]any, len(a))
	for i := range a {
		v[i] = a[i].Value
	}
	return v
}

func TestDatabaseUsersQueriesNullableRowsAndSchemaCache(t *testing.T) {
	for _, uuid := range []bool{false, true} {
		t.Run(fmt.Sprint(uuid), func(t *testing.T) {
			heart := make(chan struct{}, 4)
			group := 2
			probes := 0
			f := &sqlFixture{exec: func(q string, a []driver.NamedValue) error {
				if !strings.HasPrefix(q, "update ss_node set node_heartbeat") {
					t.Error(q)
				}
				if len(a) != 2 || a[1].Value != int64(9) {
					t.Error(a)
				}
				heart <- struct{}{}
				return nil
			}}
			f.query = func(q string, a []driver.NamedValue) (driver.Rows, error) {
				switch {
				case strings.HasPrefix(q, "select node_bandwidth_limit"):
					return rowsOf(4, []driver.Value{int64(100), int64(20), int64(group), int64(7)}), nil
				case strings.Contains(q, "information_schema.COLUMNS"):
					probes++
					if !reflect.DeepEqual(argsValues(a), []any{"schema"}) {
						t.Error(a)
					}
					if uuid {
						return rowsOf(1, []driver.Value{"passwd"}, []driver.Value{"uuid"}), nil
					}
					return rowsOf(1, []driver.Value{"passwd"}), nil
				case strings.HasPrefix(q, "select id, passwd"):
					if strings.Contains(q, ", uuid from") != uuid {
						t.Error(q)
					}
					suffix := "(class >= ? or is_admin = 1)"
					want := []any{a[0].Value, int64(7)}
					if group != 0 {
						suffix = "((class >= ? and node_group >= ?) or is_admin = 1)"
						want = append(want, int64(2))
					}
					if !strings.HasSuffix(q, suffix) || !reflect.DeepEqual(argsValues(a), want) {
						t.Error(q, a)
					}
					if _, e := time.ParseInLocation("2006-01-02 15:04:05", a[0].Value.(string), time.Local); e != nil {
						t.Fatal(e)
					}
					values := []driver.Value{int64(4), "pw", int64(3), int64(5), int64(100), float64(8.5), int64(2), int64(2345), "aes-128-gcm", "origin", nil, "plain", nil, int64(1)}
					if uuid {
						values = append(values, "known-uuid")
					}
					r := rowsOf(len(values), values)
					r.tail = errors.New("row tail is deliberately ignored")
					return r, nil
				}
				return nil, fmt.Errorf("unexpected %s", q)
			}
			p, _ := withDatabase(t, f)
			for i := 0; i < 2; i++ {
				r, e := p.GetUsersResult()
				waitEvent(t, heart)
				if e != nil || len(r.L8xOe2JrpGDq) != 1 {
					t.Fatal(r, e)
				}
				u := r.L8xOe2JrpGDq[0]
				if u.BD16miI0 != 4 || u.ESm7aAEddE != "pw" || u.DpaCpPFc != 3 || u.VsaEfGizz2x != 5 || u.TnjwdH != 8.5 || u.FPf5h1940BAv != "" || u.HX3uIcF != "" {
					t.Fatal(u)
				}
				if uuid && u.Tiybx6 != "known-uuid" {
					t.Fatal(u)
				}
				group = 0
			}
			if probes != 1 {
				t.Fatal(probes)
			}
		})
	}
}
func TestDatabaseQuotaAndQueryError(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(fmt.Sprint(fail), func(t *testing.T) {
			heart := make(chan struct{}, 1)
			sentinel := errors.New("node missing")
			f := &sqlFixture{exec: func(string, []driver.NamedValue) error { heart <- struct{}{}; return nil }, query: func(q string, a []driver.NamedValue) (driver.Rows, error) {
				if fail {
					return nil, sentinel
				}
				return rowsOf(4, []driver.Value{int64(100), int64(100), int64(0), int64(0)}), nil
			}}
			p, _ := withDatabase(t, f)
			r, e := p.GetUsersResult()
			waitEvent(t, heart)
			if fail {
				if e == nil || e.Error() != "node id: 9 not exist: node missing" {
					t.Fatal(e)
				}
			} else if e != nil || r.L8xOe2JrpGDq == nil || len(r.L8xOe2JrpGDq) != 0 {
				t.Fatal(r, e)
			}
		})
	}
}

func TestDatabaseTransactionFailureAndFallback(t *testing.T) {
	first := errors.New("old schema needs status")
	f := &sqlFixture{exec: func(q string, a []driver.NamedValue) error {
		if strings.Contains(q, "detect_log") && !strings.Contains(q, "status") {
			return first
		}
		return nil
	}, commitErr: errors.New("ignored commit")}
	p, _ := withDatabase(t, f)
	if e := p.SubmitDetectLog(L.QmiwDVjV00lZ{{GbzQUT_4f: 4, CAgUUdrv9: 6}}); e != first {
		t.Fatal(e)
	}
	if f.commits != 1 || f.rollbacks != 1 {
		t.Fatal(f.commits, f.rollbacks)
	}
	f.mu.Lock()
	calls := append([]sqlCall(nil), f.calls...)
	f.mu.Unlock()
	if len(calls) != 6 || calls[2].query != "ROLLBACK" || calls[3].query != "BEGIN" || calls[5].query != "COMMIT" {
		t.Fatal(calls)
	}
	if a := calls[4].args; len(a) != 5 || a[0].Value != int64(4) || a[1].Value != int64(6) || a[3].Value != int64(9) || a[4].Value != int64(0) {
		t.Fatal(a)
	}
	if e := p.SubmitAliveIP(L.CQhZrJ4dK{{Vy9ybu: 4, PNtn9KAV: []string{"1.2.3.4", "::1"}}}); e != nil {
		t.Fatal(e)
	}
	if f.commits != 2 {
		t.Fatal(f.commits)
	}
	var nilPanel *N3owFvm
	if nilPanel.SubmitAliveIP(nil) != nil || nilPanel.SubmitDetectLog(nil) != nil || nilPanel.SubmitTraffic(nil) != nil {
		t.Fatal("empty path")
	}
}

func TestDatabaseTrafficExactSQLAndFourEffects(t *testing.T) {
	ended := make(chan struct{}, 2)
	f := &sqlFixture{ended: ended, commitErr: errors.New("ignored commit")}
	f.query = func(q string, a []driver.NamedValue) (driver.Rows, error) {
		if strings.HasPrefix(q, "select traffic_rate") {
			return rowsOf(1, []driver.Value{float64(1.5)}), nil
		}
		if strings.Contains(q, "information_schema.TABLES") {
			return rowsOf(1, []driver.Value{"user_traffic_log"}), nil
		}
		return nil, fmt.Errorf("unexpected query %s", q)
	}
	p, _ := withDatabase(t, f)
	data := L.CWfcUv{{Hy7Q97RROK: 7, XJkGTQf_qkAt: 3, I6LXObODav: 1024}, {Hy7Q97RROK: 8, XJkGTQf_qkAt: 1, I6LXObODav: 2}}
	if e := p.SubmitTraffic(data); e != nil {
		t.Fatal(e)
	}
	waitEvent(t, ended)
	f.mu.Lock()
	calls := append([]sqlCall(nil), f.calls...)
	f.mu.Unlock()
	counts := map[string]int{}
	for _, c := range calls {
		switch {
		case strings.HasPrefix(c.query, "update user set"):
			counts["users"]++
			if !strings.Contains(c.query, "when 7 then u + 4\nwhen 8 then u + 1\nend,\nd = case id\nwhen 7 then d + 1536\nwhen 8 then d + 3\nend,") || !strings.HasSuffix(c.query, "\nend\nwhere id in(7,8)") || len(c.args) != 0 {
				t.Error(c)
			}
		case strings.HasPrefix(c.query, "update ss_node set node_bandwidth"):
			counts["bandwidth"]++
			if !reflect.DeepEqual(argsValues(c.args), []any{int64(1030), int64(9)}) {
				t.Error(c)
			}
		case strings.HasPrefix(c.query, "insert into ss_node_online_log"):
			counts["online"]++
			if c.args[0].Value != int64(9) || c.args[1].Value != int64(2) {
				t.Error(c)
			}
		case strings.HasPrefix(c.query, "insert into user_traffic_log"):
			counts["traffic"]++
			if c.args[0].Value == int64(7) && !reflect.DeepEqual(argsValues(c.args)[:6], []any{int64(7), int64(3), int64(1024), int64(9), float64(1.5), "1.50KB"}) {
				t.Error(c)
			}
		}
	}
	if !reflect.DeepEqual(counts, map[string]int{"users": 1, "bandwidth": 1, "online": 1, "traffic": 2}) {
		t.Fatal(counts)
	}
}
func TestDatabaseTrafficChunkBoundaryAndFirstFailure(t *testing.T) {
	f := &sqlFixture{}
	_, db := withDatabase(t, f)
	data := make(L.CWfcUv, 5001)
	for i := range data {
		data[i] = &L.DLxslMwYJXQw{Hy7Q97RROK: int64(i + 1), XJkGTQf_qkAt: 1, I6LXObODav: 2}
	}
	if e := updateTrafficUsers(db, data, 1, 123); e != nil {
		t.Fatal(e)
	}
	if len(f.calls) != 2 || !strings.HasSuffix(f.calls[0].query, ",5000)") || f.calls[1].query != "update user set u = case id\nwhen 5001 then u + 1\nend,\nd = case id\nwhen 5001 then d + 2\nend,\nt = case id\nwhen 5001 then 123\nend\nwhere id in(5001)" {
		t.Fatal("wrong chunk boundary")
	}
	sentinel := errors.New("stop first chunk")
	f.exec = func(string, []driver.NamedValue) error { return sentinel }
	f.calls = nil
	if e := updateTrafficUsers(db, data, 1, 123); e != sentinel || len(f.calls) != 1 {
		t.Fatal(e, len(f.calls))
	}
}

func TestDatabaseBeginErrorPrecedesReceiverRead(t *testing.T) {
	sentinel := errors.New("begin failure")
	f := &sqlFixture{beginErr: sentinel}
	_, _ = withDatabase(t, f)
	var p *N3owFvm
	if e := p.SubmitAliveIP(L.CQhZrJ4dK{{}}); e != sentinel {
		t.Fatal(e)
	}
	if e := p.SubmitDetectLog(L.QmiwDVjV00lZ{{}}); e != sentinel {
		t.Fatal(e)
	}
	if e := p._rsq9Xmy9Vtv(L.QmiwDVjV00lZ{{}}); e != sentinel {
		t.Fatal(e)
	}
	if len(f.calls) != 3 || f.commits != 0 || f.rollbacks != 0 {
		t.Fatal(f.calls)
	}
}
func TestDatabaseNodeAndDetectorScanFailures(t *testing.T) {
	mode := "normal"
	f := &sqlFixture{query: func(q string, _ []driver.NamedValue) (driver.Rows, error) {
		if strings.HasPrefix(q, "select server") {
			if mode == "bad" {
				return rowsOf(3, []driver.Value{"x", "not float", int64(1)}), nil
			}
			return rowsOf(3, []driver.Value{nil, float64(2.5), int64(1)}), nil
		}
		if mode == "bad" {
			return rowsOf(3, []driver.Value{"not int", "x", int64(1)}), nil
		}
		r := rowsOf(3, []driver.Value{int64(3), "full:example.org", int64(1)})
		r.tail = errors.New("ignored tail")
		return r, nil
	}}
	p, _ := withDatabase(t, f)
	n, e := p.GetNodeResult()
	if e != nil || n.L8xOe2JrpGDq.HIgDb6THe1A9 != "" || n.L8xOe2JrpGDq.OFIQtbK5h != 2.5 {
		t.Fatal(n, e)
	}
	rules, e := p.GetDetectRulesResult()
	if e != nil || len(rules.L8xOe2JrpGDq) != 1 {
		t.Fatal(rules, e)
	}
	mode = "bad"
	if _, e = p.GetNodeResult(); e == nil {
		t.Fatal("missing node scan error")
	}
	if _, e = p.GetDetectRulesResult(); e == nil {
		t.Fatal("missing detector scan error")
	}
}
func TestDatabaseTrafficSchemaFailureCompletesOriginalExtraSend(t *testing.T) {
	fail := errors.New("table probe failed")
	f := &sqlFixture{query: func(q string, _ []driver.NamedValue) (driver.Rows, error) {
		if strings.HasPrefix(q, "select traffic_rate") {
			return rowsOf(1, []driver.Value{nil}), nil
		}
		return nil, fail
	}}
	p, _ := withDatabase(t, f)
	done := make(chan error, 1)
	go func() { done <- p.SubmitTraffic(L.CWfcUv{{Hy7Q97RROK: 1, XJkGTQf_qkAt: 1, I6LXObODav: 2}}) }()
	select {
	case e := <-done:
		if e == nil || !strings.Contains(e.Error(), fail.Error()) {
			t.Fatal(e)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("four receive path blocked")
	}
	if trafficTableCheckedAt50C6CC8 {
		t.Fatal("failed table probe must not set cache")
	}
}
