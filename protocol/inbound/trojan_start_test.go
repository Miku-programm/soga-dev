package inbound

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"testing"
	"time"
)

func TestNativeTrojanConstructStartTCPUDPStop(t *testing.T) {
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
			node := &L.Wr5SVP{EMlXEc: 32, JhIYwEukjd: L.L9AiYC{GbD_2PU: 0, T4cuhG7MTbG: "tcp", AJjaqrBjj: "none"}}
			panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{trojanPanelFixture(9, "native-password")}}
			core, e := f1mPzegGt(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: panel})
			if e != nil {
				t.Fatal(e)
			}
			v := core.(*EfaIFpp)
			defer v.Stop()
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			if e = v.Start(); e != nil {
				t.Fatal(e)
			}
			if len(v.cDyEMW2) != 1 {
				t.Fatal("listener count")
			}
			addr := v.cDyEMW2[0].Addr().String()
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
			digest := sha256.Sum224([]byte("native-password"))
			wire := []byte(hex.EncodeToString(digest[:]))
			command := byte(1)
			if kind == "udp" {
				command = 3
			}
			target := binary.BigEndian.AppendUint16([]byte{1, 127, 0, 0, 1}, uint16(port))
			wire = append(wire, 13, 10, command)
			wire = append(wire, target...)
			wire = append(wire, 13, 10)
			payload := bytes.Repeat([]byte("trojan-native-echo\x00"), 100)
			expected := payload
			if kind == "udp" {
				frame := append([]byte(nil), target...)
				frame = binary.BigEndian.AppendUint16(frame, uint16(len(payload)))
				frame = append(frame, 13, 10)
				frame = append(frame, payload...)
				wire = append(wire, frame...)
				expected = frame
			} else {
				wire = append(wire, payload...)
			}
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
