package ssrcipher

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	A "recovered.local/soga/crypto/replaywindow"
	J "recovered.local/soga/util/ssrhelper"
	S "recovered.local/soga/protocol/ssrshared"
	L "recovered.local/soga/util/logx"
)

// Original init 0x1c9c220; all four texts are independently decoded in
// work evidence/native-ssr-cipher-strings.json.
var errorAt5098DA0 = errors.New("salt repeat")
var errorAt5098DB0 = errors.New("empty key")
var errorAt5098DC0 = errors.New("empty method")
var errorAt5098DD0 = errors.New("iv repeat")

// VavyqXsqy9V restores 0x1ca3160 (native TCP cipher factory).
func VavyqXsqy9V(info *S.L7E7H2aW8KVs, rw io.ReadWriter) (*KrOWYtzUYft, error) {
	if info.MY3lagqY == "" {
		return nil, errorAt5098DB0
	}
	if info.L6BrlCXd == "" {
		return nil, errorAt5098DC0
	}
	spec, ok := cipherMapAt5096960[info.L6BrlCXd]
	if !ok {
		return nil, errors.New("Unsupported encryption method: " + string(info.L6BrlCXd))
	}
	key := J.Is8IqhjVE6(info.MY3lagqY, spec.geqiu3e)
	c := &KrOWYtzUYft{xOgDLUH_POa: spec, k4heEtCdxndh: info}
	if spec.bsYWIvPVb != nil {
		c.yQxNctj5qy3 = rw
		c.cTNCvpoT = key
	} else {
		c.s_53ql = &jSUbF0{eMYicxBI: rw, obzsFnGuS: spec, fkUrkngrG1x: key, g9tmWrw57CtG: -1, pY3EEt3: info}
	}
	info.IawCjSkuTM4 = c.GetKey()
	info.SP2KC4lkkl3C = spec.geqiu3e
	return c, nil
}

// DP5vq3 restores 0x1ca3580 (native UDP cipher factory).
func DP5vq3(info *S.L7E7H2aW8KVs, pc net.PacketConn) (*KrOWYtzUYft, error) {
	if info.MY3lagqY == "" {
		return nil, errorAt5098DB0
	}
	if info.L6BrlCXd == "" {
		return nil, errorAt5098DC0
	}
	spec, ok := cipherMapAt5096960[info.L6BrlCXd]
	if !ok {
		return nil, errors.New("Unsupported encryption method: " + string(info.L6BrlCXd))
	}
	key := J.Is8IqhjVE6(info.MY3lagqY, spec.geqiu3e)
	c := &KrOWYtzUYft{xOgDLUH_POa: spec, k4heEtCdxndh: info}
	if spec.bsYWIvPVb != nil {
		c.WCApOca6IQW7 = pc
		c.cTNCvpoT = key
	} else {
		c.s_53ql = &jSUbF0{WCApOca6IQW7: pc, obzsFnGuS: spec, fkUrkngrG1x: key, g9tmWrw57CtG: -1, pY3EEt3: info}
	}
	info.IawCjSkuTM4 = c.GetKey()
	info.SP2KC4lkkl3C = spec.geqiu3e
	return c, nil
}
func (c *KrOWYtzUYft) GetKey() []byte {
	if c.s_53ql != nil {
		return c.s_53ql.fkUrkngrG1x
	}
	return c.cTNCvpoT
}                                  // 0x1ca3aa0
func (c *KrOWYtzUYft) KeyLen() int { return c.xOgDLUH_POa.geqiu3e } // 0x1ca3ae0
func (c *KrOWYtzUYft) SetServerInfo(info *S.L7E7H2aW8KVs) {
	if c.s_53ql != nil {
		c.s_53ql.pY3EEt3 = info
	} else {
		c.k4heEtCdxndh = info
	}
}                                                    // 0x1ca3b00
func (c *jSUbF0) SetServerInfo(info *S.L7E7H2aW8KVs) { c.pY3EEt3 = info } // 0x1c9ed00
func (c *KrOWYtzUYft) bzNeWM(iv []byte) (cipher.Stream, error) {
	rand.Read(iv[:c.xOgDLUH_POa.fzh2VK4Wf])
	return c.xOgDLUH_POa.bsYWIvPVb(c.cTNCvpoT, iv[:c.xOgDLUH_POa.fzh2VK4Wf], 1)
} // 0x1ca39c0

// Read restores 0x1ca3b80. The initial IV uses exactly one Read; a short
// successful read is an error. Data paired with a reader error is discarded.
func (c *KrOWYtzUYft) Read(p []byte) (int, error) {
	if c.s_53ql != nil {
		return c.s_53ql.Read(p)
	}
	if c.gzDImprSY == nil {
		iv := make([]byte, c.xOgDLUH_POa.fzh2VK4Wf)
		if len(iv) > 0 {
			n, e := c.yQxNctj5qy3.Read(iv)
			if e != nil {
				return 0, e
			}
			if n != len(iv) {
				return 0, fmt.Errorf("iv len not enough: %v", n)
			}
			if !A.HFlnXWty(iv) {
				L.TRTbrquD(slog.LevelDebug, "tcp iv repeat")
			}
		}
		stream, e := c.xOgDLUH_POa.bsYWIvPVb(c.cTNCvpoT, iv, 0)
		if e != nil {
			return 0, e
		}
		c.h34BSfWne__V = stream
		c.gzDImprSY = iv
		c.k4heEtCdxndh.Ds7eTY7Qp66N = iv
		c.k4heEtCdxndh.DvdB7JZX = c.xOgDLUH_POa.fzh2VK4Wf
	}
	n, e := c.yQxNctj5qy3.Read(p[:min(len(p), 8192)])
	if e != nil {
		return 0, e
	}
	c.h34BSfWne__V.XORKeyStream(p, p[:n])
	return n, nil
}

// Write restores 0x1ca3fa0, including successful short-write handling.
func (c *KrOWYtzUYft) Write(p []byte) (int, error) {
	if c.s_53ql != nil {
		return c.s_53ql.Write(p)
	}
	first := c.hnwrdS == nil
	if first {
		iv := make([]byte, c.xOgDLUH_POa.fzh2VK4Wf)
		stream, e := c.bzNeWM(iv)
		if e != nil {
			return 0, e
		}
		c.bCboKyQCMB = stream
		c.hnwrdS = iv
		c.k4heEtCdxndh.CM7BT9 = iv
		c.k4heEtCdxndh.Dtnph_j45t = c.xOgDLUH_POa.fzh2VK4Wf
	}
	prefix := 0
	if first {
		prefix = c.xOgDLUH_POa.fzh2VK4Wf
	}
	out, e := c.GetWriteBuf(len(p) + prefix)
	if e != nil {
		return 0, e
	}
	c.bCboKyQCMB.XORKeyStream(out[prefix:], p)
	if first {
		copy(out, c.hnwrdS)
	}
	if _, e = c.yQxNctj5qy3.Write(out[:len(p)+prefix]); e != nil {
		return 0, e
	}
	return len(p), nil
}

// ReadFrom restores 0x1ca42a0. IV replay is logged without rejection here.
func (c *KrOWYtzUYft) ReadFrom(p []byte) (int, net.Addr, error) {
	if c.s_53ql != nil {
		return c.s_53ql.ReadFrom(p)
	}
	n, addr, e := c.WCApOca6IQW7.ReadFrom(p)
	if e != nil {
		return 0, nil, e
	}
	size := c.xOgDLUH_POa.fzh2VK4Wf
	if n < size {
		return 0, nil, fmt.Errorf("iv len not enough: %v", n)
	}
	if size > 0 && !A.HFlnXWty(p[:size]) {
		L.TRTbrquD(slog.LevelDebug, "udp iv repeat")
	}
	stream, e := c.xOgDLUH_POa.bsYWIvPVb(c.cTNCvpoT, p[:size], 0)
	if e != nil {
		return 0, nil, e
	}
	copy(p, p[size:n])
	stream.XORKeyStream(p, p[:n-size])
	return n - size, addr, nil
}
func (c *KrOWYtzUYft) WriteTo(p []byte, addr net.Addr) (int, error) { // 0x1ca4660
	if c.s_53ql != nil {
		return c.s_53ql.WriteTo(p, addr)
	}
	size := c.xOgDLUH_POa.fzh2VK4Wf
	out, e := c.GetWriteBuf(size + len(p))
	if e != nil {
		return 0, e
	}
	stream, e := c.bzNeWM(out[:size])
	if e != nil {
		return 0, e
	}
	stream.XORKeyStream(out[size:], p)
	if _, e = c.WCApOca6IQW7.WriteTo(out[:size+len(p)], addr); e != nil {
		return 0, e
	}
	return len(p), nil
}
