package mieru

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	N "recovered.local/soga/net/limit"
	U "recovered.local/soga/user"
	Map "recovered.local/soga/util/genmap"
)

// 模拟官方 mieru 客户端：读循环收到已关闭会话的数据帧时，会先同步写回
// closeSessionRequest 再继续读。服务端读循环若在处理控制帧时同步等待写锁/写连接，
// 双方读循环会互相等待对方读取而永久卡死（带宽打满后整条底层连接断流）。
// net.Pipe 没有内核缓冲，能稳定复现 TCP 缓冲打满后的情形。
func TestSessionReaderNeverBlocksOnWrites(t *testing.T) {
	u := testMieruUser(t, 520, "stall-user", "secret")
	users := &U.ZpNFqUVZp{GTFAWF: &Map.GTFAWF[int64, U.PgaNd3014U]{}}
	users.Add(u)
	server, client := net.Pipe()
	defer client.Close()
	defer server.Close()
	go func() {
		s, err := M_03miDCRcqa(context.Background(), server, &N.O0UldNL4Eq{RacA3he0TX: users})
		if err != nil {
			return
		}
		defer s.Close()
		for {
			a, err := s.Next()
			if err != nil {
				return
			}
			go func() {
				c, err := a.Accept()
				if err != nil {
					return
				}
				defer a.j_Ra2Blepbh.Free()
				defer c.Close()
				b := make([]byte, 16384)
				for {
					if _, err := c.Write(b); err != nil {
						return
					}
				}
			}()
		}
	}()
	client.SetDeadline(time.Now().Add(10 * time.Second))
	f := newMieruFixtureClient(t, client, u)
	var wmu sync.Mutex
	write := func(kind byte, id uint32, payload []byte) error {
		wmu.Lock()
		defer wmu.Unlock()
		return f.write(kind, id, 1, payload)
	}
	// 与官方客户端一样，会话的写出在独立 goroutine 中进行。
	go func() {
		request := []byte{5, 1, 0, 1, 127, 0, 0, 1, 0x1f, 0x90}
		for _, id := range []uint32{1, 2} {
			if write(2, id, request) != nil {
				return
			}
		}
	}()
	got := map[uint32]int{}
	closed := false
	for got[2] < 8<<20 {
		kind, id, p, err := f.read()
		if err != nil {
			t.Fatalf("client stalled after %d bytes of session 2: %v", got[2], err)
		}
		if kind == 7 || kind == 11 {
			got[id] += len(p)
		}
		if !closed && got[1] > 256<<10 {
			closed = true
			go write(4, 1, nil)
			continue
		}
		if closed && id == 1 && (kind == 7 || kind == 11) {
			// 客户端已关闭会话 1，对残留数据帧回 closeSessionRequest。
			if err := write(4, 1, nil); err != nil {
				t.Fatalf("client write stalled after %d bytes of session 2: %v", got[2], err)
			}
		}
	}
}

// 一个会话暂时没有被读取（例如服务端仍在拨号目标）时，客户端在该会话上继续上传，
// 不能让同一底层连接上的其它会话立即停摆。
func TestSessionSlowConsumerDoesNotStallOthers(t *testing.T) {
	u := testMieruUser(t, 521, "hol-user", "secret")
	users := &U.ZpNFqUVZp{GTFAWF: &Map.GTFAWF[int64, U.PgaNd3014U]{}}
	users.Add(u)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	release := make(chan struct{})
	defer close(release)
	go func() {
		c, err := ln.Accept()
		if err != nil {
			return
		}
		s, err := M_03miDCRcqa(context.Background(), c, &N.O0UldNL4Eq{RacA3he0TX: users})
		if err != nil {
			return
		}
		defer s.Close()
		for {
			a, err := s.Next()
			if err != nil {
				return
			}
			go func() {
				conn, err := a.Accept()
				if err != nil {
					return
				}
				defer a.j_Ra2Blepbh.Free()
				defer conn.Close()
				if a.j_Ra2Blepbh.xIwb7TaM == 1 {
					<-release // 会话 1 迟迟不被读取
					return
				}
				b := make([]byte, 4096)
				for {
					n, err := conn.Read(b)
					if err != nil {
						return
					}
					if _, err = conn.Write(b[:n]); err != nil {
						return
					}
				}
			}()
		}
	}()
	c, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(10 * time.Second))
	f := newMieruFixtureClient(t, c, u)
	request := []byte{5, 1, 0, 1, 127, 0, 0, 1, 0x1f, 0x90}
	if err := f.write(2, 1, 0, request); err != nil {
		t.Fatal(err)
	}
	chunk := make([]byte, 16384)
	for i := 0; i < 128; i++ { // 2 MiB，旧实现 64 帧即阻塞读循环
		if err := f.write(6, 1, uint32(i+1), chunk); err != nil {
			t.Fatal(err)
		}
	}
	if err := f.write(2, 2, 0, request); err != nil {
		t.Fatal(err)
	}
	if err := f.write(6, 2, 1, []byte("ping")); err != nil {
		t.Fatal(err)
	}
	for {
		kind, id, p, err := f.read()
		if err != nil {
			t.Fatalf("session 2 stalled behind unread session 1: %v", err)
		}
		if id == 2 && kind == 7 && string(p) == "ping" {
			return
		}
	}
}
