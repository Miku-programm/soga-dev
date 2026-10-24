package ss2022

import (
	"crypto/cipher"
	"net"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/buf/bufiface"
	"recovered.local/soga/user"
)

// UjmhRG8aj 对应 RTTI 0x36ecb50，104 字节。
type UjmhRG8aj struct {
	fzA4TxUwy    *sess.Mdn09F6aMrpa
	cquf6qnPiBAg *user.WtK9HQSy66
	zG6YhL       []byte
	dr1CJBOWODc  cipher.AEAD
	aT_8mD       []byte
	tuFFaE9i53   bufiface.Mf3phSOVqgf
	jmLEXkZdhWHC bool
}

func (w *UjmhRG8aj) init() error { // 0x1c6f880
	method := w.cquf6qnPiBAg.GetCipher()
	n, err := method.KeyLen()
	if err != nil {
		return err
	}
	w.tuFFaE9i53.WriteRandom(n)
	a, err := method.NewAEAD(method.SubKey(w.cquf6qnPiBAg.PSK(), w.tuFFaE9i53.Bytes()))
	if err != nil {
		return err
	}
	w.dr1CJBOWODc = a
	w.aT_8mD = make([]byte, 12)
	return nil
}
func (w *UjmhRG8aj) a8UACSShTx(size int) error { // 0x1c6fa00
	if err := w.init(); err != nil {
		return err
	}
	start := w.tuFFaE9i53.Len()
	w.tuFFaE9i53.WriteByte(1)
	w.tuFFaE9i53.WriteUint64(uint64(uint32(rtx.C5Jff5vs0())))
	w.tuFFaE9i53.Write(w.zG6YhL)
	w.tuFFaE9i53.WriteUint16(uint16(size))
	w.tuFFaE9i53.SealWriteWithStart(w.dr1CJBOWODc, start, w.aT_8mD, nil, nil)
	incrementNonce(w.aT_8mD)
	return nil
}

// Write 对应 0x1c6fb80；原函数不分片，长度字段转换为 uint16，短写且无 error 仍返回明文全长。
func (w *UjmhRG8aj) Write(data []byte) (int, error) {
	w.tuFFaE9i53.Reset()
	if w.jmLEXkZdhWHC {
		w.tuFFaE9i53.WriteUint16(uint16(len(data)))
		w.tuFFaE9i53.Seal(w.dr1CJBOWODc, w.aT_8mD, nil)
		incrementNonce(w.aT_8mD)
	} else {
		if err := w.a8UACSShTx(len(data)); err != nil {
			return 0, err
		}
		w.jmLEXkZdhWHC = true
	}
	w.tuFFaE9i53.SealWrite(w.dr1CJBOWODc, w.aT_8mD, data, nil)
	incrementNonce(w.aT_8mD)
	if _, err := w.fzA4TxUwy.Write(w.tuFFaE9i53.Bytes()); err != nil {
		return 0, err
	}
	return len(data), nil
}
func (w *UjmhRG8aj) WriteBuffer(b bufiface.Mf3phSOVqgf) (int, error) { return w.Write(b.Bytes()) } // 0x1c6fd80
func (w *UjmhRG8aj) Free()                                              { w.tuFFaE9i53.Free() }       // 0x1c6fe00

type DrpXao = net.Conn

// DjXTT7iwk 对应 RTTI 0x37158d0；保持读写器提升的用户/目标/缓冲方法。
type DjXTT7iwk struct {
	DrpXao
	*XWewqTAtRxH2
	*UjmhRG8aj
}

func (c *DjXTT7iwk) Read(p []byte) (int, error)  { return c.XWewqTAtRxH2.Read(p) }             // 0x1c64da0
func (c *DjXTT7iwk) Write(p []byte) (int, error) { return c.UjmhRG8aj.Write(p) }               // 0x1c64e00
func (c *DjXTT7iwk) Free()                       { c.XWewqTAtRxH2.Free(); c.UjmhRG8aj.Free() } // 0x1c64e60
// YZbzDDrJX 对应 0x1c64b80，握手失败不额外关闭传入连接。
func YZbzDDrJX(conn net.Conn, m *N07NaA) (sess.Ki8urSHTx44, error) {
	counted := &sess.Mdn09F6aMrpa{DrpXao: conn}
	reader, err := yGCzb2S(counted, m)
	if err != nil {
		return nil, err
	}
	writer := &UjmhRG8aj{fzA4TxUwy: counted, cquf6qnPiBAg: reader.v6K00RA5CJ2, zG6YhL: reader.aWzYHn7, tuFFaE9i53: pool.S4OQXWIZ(0x4000)}
	return &DjXTT7iwk{DrpXao: conn, XWewqTAtRxH2: reader, UjmhRG8aj: writer}, nil
}
