package J1P4KNvfI_J

import (
	"net"
	"sync"
)

// 原 RTTI 24 字节的异步 UDP 连接结果。
type HEzs3hFmtN struct {
	hp2D4MaFj    chan struct{}
	etSDxqn41oqm net.PacketConn
}

func (r *HEzs3hFmtN) Ready() <-chan struct{} { return r.hp2D4MaFj }    //0x187e6a0
func (r *HEzs3hFmtN) Conn() net.PacketConn   { return r.etSDxqn41oqm } //0x187e6c0
// Complete @0x187e6e0，原方法只允许完成一次；再次关闭通道会 panic。
func (r *HEzs3hFmtN) Complete(c net.PacketConn) { r.etSDxqn41oqm = c; close(r.hp2D4MaFj) }

// TaQ7Prswa 保留原 RTTI 48 字节单 sync.Map 字段。
type TaQ7Prswa struct{ c28_iQWu sync.Map }

func (m *TaQ7Prswa) Get(key string) (*HEzs3hFmtN, bool) {
	v, ok := m.c28_iQWu.Load(key)
	if !ok {
		return nil, false
	}
	return v.(*HEzs3hFmtN), true
} //0x187e760
func (m *TaQ7Prswa) GetOrCreate(key string) (*HEzs3hFmtN, bool) { //0x187e800
	if r, ok := m.Get(key); ok {
		return r, true
	}
	r := &HEzs3hFmtN{hp2D4MaFj: make(chan struct{})}
	v, loaded := m.c28_iQWu.LoadOrStore(key, r)
	return v.(*HEzs3hFmtN), loaded
}

// Delete @0x187e920。未完成的项目只移除，不等候、不关闭以后完成的连接。
func (m *TaQ7Prswa) Delete(key string) {
	v, loaded := m.c28_iQWu.LoadAndDelete(key)
	if !loaded {
		return
	}
	r := v.(*HEzs3hFmtN)
	select {
	case <-r.hp2D4MaFj:
		if r.etSDxqn41oqm != nil {
			r.etSDxqn41oqm.Close()
		}
	default:
	}
}

// CloseAll @0x187e9e0 + func1 @0x1888b00：关闭已完成连接，原映射不清空。
func (m *TaQ7Prswa) CloseAll() {
	m.c28_iQWu.Range(func(_, v any) bool {
		r := v.(*HEzs3hFmtN)
		select {
		case <-r.hp2D4MaFj:
			if r.etSDxqn41oqm != nil {
				r.etSDxqn41oqm.Close()
			}
		default:
		}
		return true
	})
}
