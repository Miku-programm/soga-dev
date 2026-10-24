package encmux

import (
	"context"
	"net"
	"net/http"
	tlsclient "recovered.local/soga/transport/utlsconfig"
	ws "recovered.local/soga/transport/websocket"
	util "recovered.local/soga/transport/wsutil"
	"time"
)

// WrapClientConn restores original VA 0x11DB840. Its presence in RTTI does
// not establish that the recovered application startup invokes this helper.
func (c *VTl94pzz2IkB) WrapClientConn(ctx context.Context, raw net.Conn) (net.Conn, error) {
	if c.hZORuZstvS07 {
		return raw, nil
	}
	if c.rj1DhSBHZ7O {
		raw = tlsclient.JZejsfshGB(raw, c.psdXJ2Mi, c.p1Wtyln2GJ6)
	}
	if c.ufiMqR6x {
		var err error
		raw, err = SYSb8haLOir(ctx, raw, c.p1Wtyln2GJ6, c.lyXfKPaz4m23)
		if err != nil {
			return nil, err
		}
	}
	return &Cb7vj3KQA8{DrpXao: raw, hp3cHCehD2pe: c, npFKkBd4j: nil}, nil
}

// SYSb8haLOir restores VA 0x11E2AA0 and closure 0x11E6000. The original
// host argument is unused. It also discards Dial's buffered reader: retaining
// such bytes here would change the original behavior.
func SYSb8haLOir(ctx context.Context, raw net.Conn, host, path string) (net.Conn, error) {
	d := ws.Z4oWaKU5v{ReadBufferSize: 4096, WriteBufferSize: 4096, Timeout: 30 * time.Second,
		Header:  ws.NCYj6EtaAgQ(make(http.Header)),
		NetDial: func(context.Context, string, string) (net.Conn, error) { return raw, nil },
	}
	var err error
	raw, _, _, err = d.Dial(ctx, "ws://"+raw.RemoteAddr().String()+path)
	if err != nil {
		return nil, err
	}
	return &F2Eb1b{DrpXao: raw, jsKotLrECn: util.L9DiQX(raw, ws.StateClientSide)}, nil
}
