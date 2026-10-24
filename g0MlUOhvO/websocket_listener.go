package g0MlUOhvO

import (
	"context"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"io"
	"log"
	"net"
	"net/http"
	ws "recovered.local/soga/FgcbiePIJLm"
	"recovered.local/soga/HxKViEp2a"
	util "recovered.local/soga/UYoO4Hz"
	"recovered.local/soga/muc279Mpa6S4"
	"sync"
	"time"
)

// 原 RTTI bhXMvjeYm_m.YE2KB1I，48B；实际包路径 g0MlUOhvO。
// 以下是二进制源码表达，并非作者原始文件。
type YE2KB1I struct {
	O6B2qZ0p0
	kmDrwd       chan net.Conn
	jPywQp       string
	z35DsaN5f7ru bool
}

// 原 RTTI 64B。关闭标志和回调的原并发行为均未改变。
type V72bzf8 struct {
	DrpXao
	aSm553fxV  io.Reader
	tc_ZEEKYjB util.KTXwG8eroaXF
	vJJe8B59OX sync.Once
	e9QIo8Yf8z func()
}

// GHrg5b6zAf 对应 0x11e33c0；服务协程 0x11e6120。
func GHrg5b6zAf(listener net.Listener, path string) net.Listener {
	v := &YE2KB1I{O6B2qZ0p0: listener, kmDrwd: make(chan net.Conn, 16384)}
	go func() {
		mux := http.NewServeMux()
		mux.Handle(path, http.HandlerFunc(v.r3NKou))
		s := http.Server{ErrorLog: log.New(io.Discard, "", 1), Handler: mux}
		if e := s.Serve(listener); e != nil {
			muc279Mpa6S4.TRTbrquD(8, "http server error", e)
		}
	}()
	return v
}

// Accept 0x11e3520；两个关闭分支都创建同一原错误文本的新 error。
func (v *YE2KB1I) Accept() (net.Conn, error) {
	if v.z35DsaN5f7ru {
		return nil, errors.New("ws listener closed")
	}
	c := <-v.kmDrwd
	if v.z35DsaN5f7ru {
		if c != nil {
			c.Close()
		}
	} else if c != nil {
		return c, nil
	}
	return nil, errors.New("ws listener closed")
}

// Close 0x11e36e0，原代码直接使用 bool/channel，没有额外锁或等待。
func (v *YE2KB1I) Close() error {
	if v.z35DsaN5f7ru {
		return nil
	}
	v.z35DsaN5f7ru = true
	close(v.kmDrwd)
	return v.O6B2qZ0p0.Close()
}

// r3NKou 0x11e3740，两个闭包 0x11e6320/0x11e6340。
func (v *YE2KB1I) r3NKou(w http.ResponseWriter, r *http.Request) {
	defer func() { HxKViEp2a.SlG_Y_mg("server ws"); r.Body.Close() }()
	if v.z35DsaN5f7ru {
		return
	}
	c, _, _, err := ws.DefaultHTTPUpgrader.Upgrade(r, w)
	if err != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	wrapped := &V72bzf8{DrpXao: c, tc_ZEEKYjB: util.L9DiQX(c, ws.StateServerSide), e9QIo8Yf8z: func() { cancel() }}
	// 原调用者忽略 Upgrade 返回的 bufio.Reader；不追加不存在的缓冲拼接逻辑。
	v.kmDrwd <- wrapped
	<-ctx.Done()
}

// fAbNA17x_Ire 0x11e53c0；每条数据消息创建新帧读取器，控制帧交给原处理器。
func (v *V72bzf8) fAbNA17x_Ire() (io.Reader, error) {
	if v.aSm553fxV != nil {
		return v.aSm553fxV, nil
	}
	r := &util.HLKO3Cdqw{GX04GFKh1Gz: v.DrpXao, WXUh5psqY7: ws.StateServerSide, EbVEHBvY4lh: v.tc_ZEEKYjB}
	for {
		h, e := r.NextFrame()
		if e != nil {
			return nil, e
		}
		if h.IsHe8C&8 == 0 {
			break
		}
		if e = v.tc_ZEEKYjB(h, r); e != nil {
			return nil, e
		}
	}
	v.aSm553fxV = r
	return r, nil
}

// Read 0x11e55a0；原 pkg/errors.Cause 比较 EOF，不改为 errors.Is。
func (v *V72bzf8) Read(p []byte) (int, error) {
	for {
		r, e := v.fAbNA17x_Ire()
		if e != nil {
			return 0, e
		}
		n, e := r.Read(p)
		if pkgerrors.Cause(e) != io.EOF {
			return n, e
		}
		v.aSm553fxV = nil
		if n > 0 {
			return n, nil
		}
	}
}

// Write 0x11e56c0，单条完整二进制消息，失败时返回 0 而非局部写入长度。
func (v *V72bzf8) Write(p []byte) (int, error) {
	e := util.WriteFrameRecovered(v.DrpXao, ws.StateServerSide, ws.OpBinary, true, p)
	if e != nil {
		return 0, e
	}
	return len(p), nil
}

// Close 0x11e5800 / 闭包 0x11e63c0；每次 Close 都关闭底层并调用通知。
func (v *V72bzf8) Close() error {
	v.vJJe8B59OX.Do(func() {
		v.DrpXao.SetWriteDeadline(time.Now().Add(time.Second))
		util.WriteFrameRecovered(v.DrpXao, ws.StateServerSide, ws.OpClose, true, []byte{})
	})
	e := v.DrpXao.Close()
	v.e9QIo8Yf8z()
	return e
}
func (v *V72bzf8) GetRawConn() net.Conn { return v.DrpXao } // 0x11e58a0
