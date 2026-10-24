package J1P4KNvfI_J

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
)

// 原嵌入字段名DrpXao来自混淆后的net.Conn；别名仅保持该字段名称与提升方法。
type DrpXao = net.Conn

// Mdn09F6aMrpa对应RTTI0x36fc768，64字节，原前缀bPyexLTWaQN。
type Mdn09F6aMrpa struct {
	DrpXao
	dqmafe    int
	e1uNaRZ   int
	eGVVb0FCf nzATvajt.PgaNd3014U
	g5cCyo    context.Context
}

// SetUser@0x18855c0先绑定用户与上下文，再结算已缓存流量；nil用户保留原panic。
func (c *Mdn09F6aMrpa) SetUser(user nzATvajt.PgaNd3014U) {
	c.eGVVb0FCf = user
	c.g5cCyo = user.GetCtx()
	user.AddUp(c.dqmafe)
	user.AddDown(c.e1uNaRZ)
	c.dqmafe = 0
	c.e1uNaRZ = 0
}
func (c *Mdn09F6aMrpa) GetUser() nzATvajt.PgaNd3014U { return c.eGVVb0FCf } //0x18856c0
func (c *Mdn09F6aMrpa) GetRawConn() net.Conn         { return c.DrpXao }    //0x1885ca0
func (c *Mdn09F6aMrpa) AddUp(n int) {
	if c.eGVVb0FCf != nil {
		c.eGVVb0FCf.AddUp(n)
	} else {
		c.dqmafe += n
	}
} //0x1885a40
func (c *Mdn09F6aMrpa) AddDown(n int) {
	if c.eGVVb0FCf != nil {
		c.eGVVb0FCf.AddDown(n)
	} else {
		c.e1uNaRZ += n
	}
} //0x1885aa0
// zFM714L9G@0x18856e0不主动Close；原代码只记录日志并返回错误。
func (c *Mdn09F6aMrpa) zFM714L9G() error {
	if c.eGVVb0FCf == nil {
		return nil
	}
	if c.eGVVb0FCf.IsTrafficOut() {
		muc279Mpa6S4.TRTbrquD(slog.LevelDebug, "user traffic out:", c.eGVVb0FCf.GetId())
		muc279Mpa6S4.TRTbrquD(slog.LevelDebug, "close user connection:", c.eGVVb0FCf.GetId())
		return fmt.Errorf("user traffic out")
	}
	select {
	case <-c.g5cCyo.Done():
		muc279Mpa6S4.TRTbrquD(slog.LevelDebug, "user context done")
		return HxKViEp2a.ErrorAt5097800
	default:
		return nil
	}
}

// Read@0x1885b00读取后计数，随后检查失败会丢弃本次n及底层error。
func (c *Mdn09F6aMrpa) Read(data []byte) (int, error) {
	n, err := c.DrpXao.Read(data)
	c.AddUp(n)
	if checked := c.zFM714L9G(); checked != nil {
		return 0, checked
	}
	return n, err
}

// Write@0x1885be0按请求长度提前计数；检查失败时不调用底层Write。
func (c *Mdn09F6aMrpa) Write(data []byte) (int, error) {
	c.AddDown(len(data))
	if err := c.zFM714L9G(); err != nil {
		return 0, err
	}
	return c.DrpXao.Write(data)
}
