package inbound

import (
	"context"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	U "recovered.local/soga/user"
	Q "recovered.local/soga/panel/registry"
	"testing"
	"time"
)

func TestNativeVMessConstructStartRelayStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	for _, mode := range []string{"aead", "auto"} {
		t.Run(mode, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "soga.conf")
			os.WriteFile(path, []byte("listen=127.0.0.1\nvmess_auth_mode="+mode+"\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\n"), 0600)
			C.MkwNWciPHS(path)
			if e := C.TvBECnD9fM6(); e != nil {
				t.Fatal(e)
			}
			node := &L.Wr5SVP{EMlXEc: 31, SwxJNgPpg: L.JW0CeHo3O{CKsOpayd: 0, F2GXyZZ: "tcp"}}
			panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")}}
			core, e := ag90n_rZjJO(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: panel})
			if e != nil {
				t.Fatal(e)
			}
			v := core.(*K4avFiYOPi)
			defer v.Stop()
			if (v.dj_cTnaUp6f == nil) != (mode == "aead") || v.p54XvyV == nil || v.l_ZoZQDxG7 == nil || v.qgz9PfTS == nil {
				t.Fatal("original constructor mode branches")
			}
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			if e = v.Start(); e != nil {
				t.Fatal(e)
			}
			if len(v.l8Uq4M) != 1 {
				t.Fatal("listener count", len(v.l8Uq4M))
			}
			addr := v.l8Uq4M[0].Addr().String()
			echo, e := net.Listen("tcp4", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer echo.Close()
			done := make(chan error, 1)
			go func() {
				c, e := echo.Accept()
				if e != nil {
					done <- e
					return
				}
				defer c.Close()
				c.SetDeadline(time.Now().Add(5 * time.Second))
				_, e = io.Copy(c, c)
				done <- e
			}()
			c, e := net.DialTimeout("tcp", addr, time.Second)
			if e != nil {
				t.Fatal(e)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(5 * time.Second))
			u, _ := v.jc_QPN6.Get(7)
			exerciseNativeVMessWire(t, c, u.(*U.QKaIWV2), echo.Addr().(*net.TCPAddr).Port, "tcp")
			c.Close()
			v.Stop()
			if v.jc_QPN6.Len() != 0 || v.zb8f3cHWW8v.Err() != context.Canceled {
				t.Fatal("Stop retained active users/context")
			}
			if c, e = net.DialTimeout("tcp", addr, 100*time.Millisecond); e == nil {
				c.Close()
				t.Fatal("listener remained open")
			}
			select {
			case e := <-done:
				if e != nil {
					t.Fatal(e)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("destination remained open")
			}
		})
	}
}
