package sess

import (
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/user"
)

// AFui66gdB is the original 80-byte inbound datagram (RTTI 0x36cc680).
type AFui66gdB struct {
	Y8YsJ7QnEI3h KsKg3yls
	AIO3Qb       net.Addr
	PNliwRWkrU   *relay.AHU2MRmjzQu
	S5eGpRGjp5oo user.PgaNd3014U
	AGHsUHa      []byte
}

// TgFd4ag1 @0x187ed40. Only concrete TCP/UDP addresses populate the IP
// fields. Unlike the stream constructor, no submit-IP normalization occurs.
func TgFd4ag1(p AFui66gdB) *ZzkV7_5l {
	s := &ZzkV7_5l{E3L0l8hZ: int(sessionSequenceAt50C7908.Add(1)), PayAlbdv6: p.S5eGpRGjp5oo, Iayb4V: p.PNliwRWkrU, YlxwftjCvz_: 1}
	s.KeHzUwuXWh0, s.Y72kDV35, _ = sessionAddressParts(p.AIO3Qb)
	if p.Y8YsJ7QnEI3h != nil {
		s.L4SIW3qbjS, s.X1gBTWHQO99, _ = sessionAddressParts(p.Y8YsJ7QnEI3h.LocalAddr())
	}
	return s
}
