package hysteria

import (
	"context"
	"errors"
	quic "github.com/apernet/quic-go"
	"github.com/apernet/quic-go/quicvarint"
	"net"
	B "recovered.local/soga/net/dns"
	Pool "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	Map "recovered.local/soga/util/genmap"
	J "recovered.local/soga/net/sess"
	F "recovered.local/soga/protocol/hysteriaauth"
	C "recovered.local/soga/config"
	W "recovered.local/soga/crypto/misc"
	U "recovered.local/soga/user"
	R "recovered.local/soga/crypto/rand"
	"sync"
	"time"
)

type aK1tZAfWP2z struct {
	q3H964UHKW   context.Context
	oPadewGA     context.CancelFunc
	bibMnwQ      quic.Connection
	v7t5F5tZ     *F.YWXI4jW3
	oi0B93oah    uint32
	iXQ9zjKttm0F chan *W.ItLMumzBf
	g54qtKk8h    *time.Timer
	viQjY4UoFT6o *B.D23y2888a_
	hlcbMoqH4    U.PgaNd3014U
	iIraKZ       *G.AHU2MRmjzQu
	smLqRCJL3Msj sync.Once
	a8Q2aY       func()
}

func oaAwkKrUsl(conn quic.Connection, id uint32, in chan *W.ItLMumzBf, resolver *B.D23y2888a_, user U.PgaNd3014U, dest *G.AHU2MRmjzQu, onClose func()) J.Ki8urSHTx44 { //0x1c337c0
	ctx, cancel := context.WithCancel(conn.Context())
	return &aK1tZAfWP2z{q3H964UHKW: ctx, oPadewGA: cancel, bibMnwQ: conn, v7t5F5tZ: &F.YWXI4jW3{}, oi0B93oah: id, iXQ9zjKttm0F: in, g54qtKk8h: time.NewTimer(time.Duration(C.UDPTimeoutMinutes()) * time.Minute), viQjY4UoFT6o: resolver, hlcbMoqH4: user, iIraKZ: dest, a8Q2aY: onClose}
}

func (c *aK1tZAfWP2z) Read(b []byte) (int, error)  { n, _, err := c.ReadFromProxy(b); return n, err }   //0x1c33a40
func (c *aK1tZAfWP2z) Write(b []byte) (int, error) { return c.WriteToProxy(b, c.bibMnwQ.RemoteAddr()) } //0x1c33aa0
func (c *aK1tZAfWP2z) NetworkType() J.ZES6F5       { return 1 }                                         //0x1c33b40
func (c *aK1tZAfWP2z) ReadFromProxy(b []byte) (int, net.Addr, error) { //0x1c33b60
	for {
		select {
		case <-c.g54qtKk8h.C:
			return 0, nil, errors.New("timeout")
		case <-c.q3H964UHKW.Done():
			return 0, nil, errors.New("closed")
		case fragment := <-c.iXQ9zjKttm0F:
			packet := c.v7t5F5tZ.Feed(fragment)
			if packet == nil || len(packet.HmoFPk) > len(b) {
				continue
			}
			addr, err := G.U80FDt(packet.XrDcBXcD)
			if err != nil {
				continue
			}
			if !addr.HasIP() {
				if err = addr.ResolveHost(c.viQjY4UoFT6o); err != nil {
					continue
				}
			}
			udp, err := addr.ToUDPAddr()
			if err != nil {
				continue
			}
			return copy(b, packet.HmoFPk), udp, nil
		}
	}
}

func (c *aK1tZAfWP2z) WriteToProxy(b []byte, addr net.Addr) (int, error) { //0x1c33e20
	select {
	case <-c.q3H964UHKW.Done():
		return 0, errors.New("closed")
	default:
	}
	packet := W.ItLMumzBf{D35ycLc5aB6r: c.oi0B93oah, Kb9FLg7wuNE: 1, XrDcBXcD: addr.String(), HmoFPk: b}
	buf := Pool.PoolAt50961D8.Get(8 + quicvarint.Len(uint64(len(packet.XrDcBXcD))) + len(packet.XrDcBXcD) + len(b))
	defer Pool.PoolAt50961D8.Put(buf)
	packet.Serialize(buf)
	err := c.bibMnwQ.SendDatagram(buf)
	if err == nil {
		return len(b), nil
	}
	var tooLarge *quic.DatagramTooLargeError
	if !errors.As(err, &tooLarge) {
		return 0, err
	}
	packet.CBHlJenOD = uint16(R.Co6D_zfls(65535) + 1)
	for _, part := range F.TOF1KM6hi(&packet, int(tooLarge.MaxDataLen)) {
		n := part.Serialize(buf)
		if n >= 0 {
			if err = c.bibMnwQ.SendDatagram(buf[:n]); err != nil {
				return 0, err
			}
		}
	}
	return len(b), nil
}

func (c *aK1tZAfWP2z) Close() error {
	c.oPadewGA()
	c.smLqRCJL3Msj.Do(func() { c.a8Q2aY() })
	return nil
}                                                    //0x1c343a0
func (c *aK1tZAfWP2z) LocalAddr() net.Addr           { return c.bibMnwQ.LocalAddr() }  //0x1c34420
func (c *aK1tZAfWP2z) RemoteAddr() net.Addr          { return c.bibMnwQ.RemoteAddr() } //0x1c34460
func (c *aK1tZAfWP2z) SetDeadline(t time.Time) error { return c.SetReadDeadline(t) }   //0x1c344a0
func (c *aK1tZAfWP2z) SetReadDeadline(t time.Time) error {
	if t.IsZero() {
		c.g54qtKk8h.Stop()
	} else {
		c.g54qtKk8h.Reset(time.Until(t))
	}
	return nil
}                                                       //0x1c34500
func (c *aK1tZAfWP2z) SetWriteDeadline(time.Time) error { return nil }         //0x1c345a0
func (c *aK1tZAfWP2z) GetUser() U.PgaNd3014U            { return c.hlcbMoqH4 } //0x1c345c0
func (c *aK1tZAfWP2z) GetDestAddr() *G.AHU2MRmjzQu      { return c.iIraKZ }    //0x1c345e0

type kTVWUcbD5rU struct{ zUfp8R5e chan *W.ItLMumzBf }
type chFbgN struct {
	yoncNMg    context.Context
	eAEJVc     *B.D23y2888a_
	ykxy0e     quic.Connection
	ndlUQyGetN chan J.Ki8urSHTx44
	iuT4giLUja U.PgaNd3014U
	w7USyrd7   Map.GTFAWF[uint32, *kTVWUcbD5rU]
}

func (h *chFbgN) aMFjq2VlDD() { //0x1c34640
	for {
		data, err := h.ykxy0e.ReceiveDatagram(h.yoncNMg)
		if err != nil {
			return
		}
		packet, err := W.CY9jLhcY1f5(data)
		if err != nil {
			continue
		}
		id := packet.D35ycLc5aB6r
		s, ok := h.w7USyrd7.Get(id)
		if !ok {
			dest, err := G.U80FDt(packet.XrDcBXcD)
			if err != nil {
				continue
			}
			s = &kTVWUcbD5rU{make(chan *W.ItLMumzBf, 128)}
			h.w7USyrd7.Set(id, s)
			conn := oaAwkKrUsl(h.ykxy0e, id, s.zUfp8R5e, h.eAEJVc, h.iuT4giLUja, dest, func() { h.w7USyrd7.Delete(id) })
			h.ndlUQyGetN <- conn
		}
		s.zUfp8R5e <- packet
	}
}
