package sess

import (
	"context"
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/net/sniff"
	"sync/atomic"
)

// ZzkV7_5l 和管道接口按原 RTTI 恢复；全部字段和顺序保持原样。
type ZzkV7_5l struct {
	E3L0l8hZ    int
	IcoyFqRBU   int
	UWItZyY66   *limit.O0UldNL4Eq
	L4SIW3qbjS  net.IP
	X1gBTWHQO99 int
	KeHzUwuXWh0 net.IP
	Y72kDV35    int
	Uqzvq94e    string
	R61M7ZTr    net.IP
	N15sqBj36   int
	IksfOg      net.IP
	IqEsiK      int
	TQWc2J      sniff.FlCuFjj_V1u
	PayAlbdv6   user.PgaNd3014U
	Iayb4V      *relay.AHU2MRmjzQu
	YlxwftjCvz_ ZES6F5
	C66t8X      string
}
type ZXQDsvH func()
type Yew2AaR func()
type Stcofg interface {
	AfterRemote(*ZzkV7_5l) error
	SessionStart(context.Context, context.CancelFunc, *ZzkV7_5l) (ZXQDsvH, error)
}
type C7y6el2j interface {
	Handle(*ZzkV7_5l, DqwVny0NcXb) (DqwVny0NcXb, bool, Yew2AaR, error)
}
type ZbXMdM3YH4jF []C7y6el2j
type PH7Th2Ejpz []Stcofg

// 地址命名保留原 qword_50C7908 的全局会话序号；原静态初值为零。
var sessionSequenceAt50C7908 atomic.Int64

// MHLL7QOie 对应 0x187ea20。原代码仅接受实际 *net.TCPAddr/*net.UDPAddr，
// 不解析任意 Addr.String；本地/远端 IP 切片共享底层存储。
func MHLL7QOie(c Ki8urSHTx44) *ZzkV7_5l {
	s := &ZzkV7_5l{E3L0l8hZ: int(sessionSequenceAt50C7908.Add(1)), PayAlbdv6: c.GetUser(), Iayb4V: c.GetDestAddr(), YlxwftjCvz_: c.NetworkType()}
	s.L4SIW3qbjS, s.X1gBTWHQO99, _ = sessionAddressParts(c.LocalAddr())
	var ok bool
	s.KeHzUwuXWh0, s.Y72kDV35, ok = sessionAddressParts(c.RemoteAddr())
	if ok {
		v4, v6, _ := config.SessionIPConfiguration()
		s.Uqzvq94e = relay.HYeSYzPzN(s.KeHzUwuXWh0, v4, v6).String()
		s.C66t8X = s.Uqzvq94e
	}
	return s
}

// sessionAddressParts 是用于表达多个原函数内联地址分支的集成辅助，不计原函数。
func sessionAddressParts(a net.Addr) (net.IP, int, bool) {
	switch a := a.(type) {
	case *net.TCPAddr:
		return a.IP, a.Port, true
	case *net.UDPAddr:
		return a.IP, a.Port, true
	}
	return nil, 0, false
}

// GetSubmitIp 对应 0x187ece0；每次读取全局配置，不缓存配置开关。
func (s *ZzkV7_5l) GetSubmitIp() string {
	_, _, realIP := config.SessionIPConfiguration()
	if realIP {
		return s.KeHzUwuXWh0.String()
	}
	return s.Uqzvq94e
}

// Clone 对应 0x187f260：浅拷贝全部字段，并分配新的全局序号。
func (s *ZzkV7_5l) Clone() *ZzkV7_5l {
	id := int(sessionSequenceAt50C7908.Add(1))
	r := *s
	r.E3L0l8hZ = id
	return &r
}
