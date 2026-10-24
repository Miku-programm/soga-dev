package nXZ99wahk

import (
	"io"
	"net"
	"os"
	"recovered.local/soga/BaEWsPy"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/J1P4KNvfI_J"
)

// WSnkIcJKU 为原24字节结构；通过真实嵌入保留原自动生成方法。
type WSnkIcJKU struct {
	*Ldrr25M
	oNZFwhbp85ce *BaEWsPy.D23y2888a_
	dBLfTvhGz    bool
}

// ReadFromProxy 对应0x1ceb060；解析失败的包被释放并继续等待，短目标缓冲丢弃剩余包内容。
func (c *WSnkIcJKU) ReadFromProxy(p []byte) (int, net.Addr, error) {
	for {
		select {
		case <-c.s1w1RA.Wait():
			return 0, nil, os.ErrDeadlineExceeded
		case f, ok := <-c.fj8yHOR:
			if !ok {
				return 0, nil, io.EOF
			}
			addr := f.hUHpeGfV
			if addr != nil {
				c.dBLfTvhGz = true
			} else {
				addr = c.yFmjxzi
			}
			if err := addr.ResolveHost(c.oNZFwhbp85ce); err != nil {
				f.FreeData()
				continue
			}
			udp, err := addr.ToUDPAddr()
			if err != nil {
				f.FreeData()
				continue
			}
			n, _ := f.c8AtRj7hb.Read(p)
			f.FreeData()
			return n, udp, nil
		}
	}
}

// WriteToProxy 对应0x1ceb320。保留原两个分支不同的错误处理；带地址时忽略连接写入错误。
func (c *WSnkIcJKU) WriteToProxy(p []byte, addr net.Addr) (int, error) {
	if !c.dBLfTvhGz {
		return c.Ldrr25M.Write(p)
	}
	a, err := G6bHD5JA.RNFJeaRY(addr)
	if err != nil {
		return 0, err
	}
	b := D_zCag1qTV.S4OQXWIZ(len(p) + 30)
	defer b.Free()
	if a.GetFirstIp().To4() != nil {
		b.WriteUint16(12)
	} else {
		b.WriteUint16(24)
	}
	b.WriteUint16(c.gwE4Or1X8JW)
	b.WriteByte(2)
	b.WriteByte(1)
	b.WriteByte(2)
	if _, err = G6bHD5JA.HfFNot(b, a); err != nil {
		return 0, err
	}
	b.WriteUint16(uint16(len(p)))
	b.Write(p)
	c.hiMx_c.Write(b.Bytes())
	return len(p), nil
}
func (c *WSnkIcJKU) NetworkType() J1P4KNvfI_J.ZES6F5 { return 1 } //0x1cebb40
