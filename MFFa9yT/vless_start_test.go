package MFFa9yT

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/gofrs/uuid"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"testing"
	"time"
)

func TestNativeVLESSConstructStartTCPUDPStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	path := filepath.Join(t.TempDir(), "soga.conf")
	os.WriteFile(path, []byte("listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600)
	C.MkwNWciPHS(path)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	for _, kind := range []string{"tcp", "udp"} {
		t.Run(kind, func(t *testing.T) {
			node := &L.Wr5SVP{EMlXEc: 32, Ck6DxjB: L.VkkFLH{GJtOdKOWBiFP: 0, XmIaPj3CzOF: "tcp"}}
			panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panelVless(9, "11111111-1111-4111-8111-111111111111")}}
			core, e := oHr1LeATmS(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: panel})
			if e != nil {
				t.Fatal(e)
			}
			v := core.(*KHonGaqbIy_6)
			defer v.Stop()
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			if e = v.Start(); e != nil {
				t.Fatal(e)
			}
			if len(v.kX5i8U) != 1 {
				t.Fatal("listener count")
			}
			addr := v.kX5i8U[0].Addr().String()
			var port int
			done := make(chan error, 1)
			if kind == "tcp" {
				echo, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer echo.Close()
				port = echo.Addr().(*net.TCPAddr).Port
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
			} else {
				echo, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer echo.Close()
				port = echo.LocalAddr().(*net.UDPAddr).Port
				go func() {
					echo.SetDeadline(time.Now().Add(5 * time.Second))
					b := make([]byte, 4096)
					n, a, e := echo.ReadFrom(b)
					if e == nil {
						_, e = echo.WriteTo(b[:n], a)
					}
					done <- e
				}()
			}
			c, e := net.DialTimeout("tcp", addr, time.Second)
			if e != nil {
				t.Fatal(e)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(5 * time.Second))
			id := uuid.FromStringOrNil("11111111-1111-4111-8111-111111111111")
			wire := append([]byte{0}, id[:]...)
			command := byte(1)
			if kind == "udp" {
				command = 2
			}
			wire = append(wire, 0, command)
			wire = binary.BigEndian.AppendUint16(wire, uint16(port))
			wire = append(wire, 1, 127, 0, 0, 1)
			payload := bytes.Repeat([]byte("vless-native-echo\x00"), 100)
			expected := []byte{0, 0}
			if kind == "udp" {
				wire = binary.BigEndian.AppendUint16(wire, uint16(len(payload)))
				expected = binary.BigEndian.AppendUint16(expected, uint16(len(payload)))
			}
			wire = append(wire, payload...)
			expected = append(expected, payload...)

			if _, e = c.Write(wire); e != nil {
				t.Fatal(e)
			}
			got := make([]byte, len(expected))
			if _, e = io.ReadFull(c, got); e != nil || !bytes.Equal(got, expected) {
				t.Fatalf("%s native relay %v", kind, e)
			}
			c.Close()
			user, _ := v.jc_QPN6.Get(9)
			previousCtx := user.GetCtx()
			v.Stop()
			if v.jc_QPN6.Len() != 1 || previousCtx.Err() != context.Canceled || user.GetCtx().Err() != nil || v.zb8f3cHWW8v.Err() != context.Canceled {
				t.Fatal("original stop state")
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

func TestNativeVLESSStartRejectsBeforeListening(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	for _, tc := range []struct{ flow, decryption, want string }{
		{"invalid", "", "not support vless flow: invalid"},
		{"", "invalid", "invalid vless decryption: invalid"},
	} {
		node := &L.Wr5SVP{Ck6DxjB: L.VkkFLH{XmIaPj3CzOF: "tcp", NMzMUQ2: tc.flow, DlDK5Y9AS: tc.decryption}}
		core, e := oHr1LeATmS(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node}})
		if e != nil {
			t.Fatal(e)
		}
		v := core.(*KHonGaqbIy_6)
		e = v.Start()
		v.Stop()
		if e == nil || e.Error() != tc.want || len(v.kX5i8U) != 0 {
			t.Fatalf("%+v %v", tc, e)
		}
	}
}
