package oXGarxahaig

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/go-sql-driver/mysql"
	"recovered.local/soga/cNT24rG4F"
)

func TestOriginalDatabaseInitializationIsLazyAndOnce(t *testing.T) {
	path := filepath.Join(t.TempDir(), "soga.conf")
	oldPath := cNT24rG4F.ConfigurationPath()
	t.Cleanup(func() { cNT24rG4F.MkwNWciPHS(oldPath) })
	load := func(host, name string) {
		t.Helper()
		text := fmt.Sprintf("type=sspanel-uim\nserver_type=ss\nnode_id=1\napi=db\ndb_host=%s\ndb_port=3307\ndb_name=%s\ndb_user=fixture-user\ndb_password=fixture-pass\ndb_max_idle_time=2\ndb_max_idle_conn=2\ndb_max_conn=3\n", host, name)
		if err := os.WriteFile(path, []byte(text), 0600); err != nil {
			t.Fatal(err)
		}
		cNT24rG4F.MkwNWciPHS(path)
		if err := cNT24rG4F.TvBECnD9fM6(); err != nil {
			t.Fatal(err)
		}
	}
	var addresses []string
	stop := errors.New("fixture prevents network access")
	mysql.RegisterDialContext("tcp", func(_ context.Context, address string) (net.Conn, error) {
		addresses = append(addresses, address)
		return nil, stop
	})
	load("127.0.0.1", "fixture")
	db := DB()
	if db == nil {
		t.Fatal("original initialization returned nil")
	}
	t.Cleanup(func() { _ = db.Close() })
	if len(addresses) != 0 || db.Stats().MaxOpenConnections != 3 {
		t.Fatalf("initialization must not dial: %v, %+v", addresses, db.Stats())
	}
	if err := db.PingContext(context.Background()); !errors.Is(err, stop) {
		t.Fatalf("unexpected driver result: %v", err)
	}
	// %v produces [127.0.0.1], preserved by mysql.ParseDSN and handed to DialContext.
	if len(addresses) != 1 || addresses[0] != "[127.0.0.1]:3307" {
		t.Fatalf("original full-slice DSN address: %v", addresses)
	}
	load("::1", "changed")
	if DB() != db || len(addresses) != 1 {
		t.Fatal("second initialization must reuse the original global pointer")
	}
	_ = db.Close()
	onceAt50C7CF0 = sync.Once{}
	load("127.0.0.1", "fixture?timeout=invalid")
	if DB() != nil {
		t.Fatal("sql.Open parse error must leave the original nil pointer")
	}
	load("127.0.0.1", "valid-again")
	if DB() != nil {
		t.Fatal("failed initial Open must still consume sync.Once")
	}
}
