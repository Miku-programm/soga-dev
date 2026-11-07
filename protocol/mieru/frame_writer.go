package mieru

import (
	"crypto/cipher"
	"fmt"
	"net"
	D "recovered.local/soga/buf/pool"
	U "recovered.local/soga/buf/bufiface"
	"sync"
	"sync/atomic"
	"time"
)

// qhFcsZC is the underlay frame writer; all writes are serialized by its mutex.
type qhFcsZC struct {
	v_7InizyTC22 sync.Mutex
	_z_j84       net.Conn
	yn5oQQkf8    cipher.AEAD
	g9zKwmF      []byte
	oynKh2Fa     bool
	rAB3kZi8     string
	// pattern 为创建连接时生效的 traffic pattern，nil 表示保持原行为。
	pattern *trafficPattern
	// clientLowEntropy 在收到客户端低熵数据帧后置位，之后服务端才发送低熵数据帧。
	clientLowEntropy atomic.Bool
	// err 记录第一次写失败；之后的写直接失败，避免在已错位的 AEAD 流上继续写。
	err error
	// deadlineRefresh 是下次刷新写超时的时间点（仅在持锁时访问）。
	deadlineRefresh time.Time
	// onError 在第一次写失败时被调用（异步），由会话用来拆除整条底层连接。
	onError func()
}

// writeTimeout 限制一次底层写操作的最长阻塞时间。对端长时间不读时，写操作
// 失败并拆除整条连接，而不是持有写锁永久阻塞所有会话。
const writeTimeout = 30 * time.Second

func vfhZtLu(conn net.Conn, aead cipher.AEAD, password string) (*qhFcsZC, error) { //0x1c535c0
	nonce, err := o0fNBhk(password)
	if err != nil {
		return nil, fmt.Errorf("generate nonce: %w", err)
	}
	// 与 mieru 一致：先按 pattern 改写 nonce 前缀，再写入依赖 nonce[:16] 的用户提示。
	pattern := currentTrafficPattern()
	if pattern != nil {
		pattern.applyNonce(nonce)
		if password != "" {
			uJghIeaJPse(nonce, password)
		}
	}
	return &qhFcsZC{_z_j84: conn, yn5oQQkf8: aead, g9zKwmF: nonce, oynKh2Fa: true, rAB3kZi8: password, pattern: pattern}, nil
}

// incrementNonce is an analyst name for the repeated original big-endian loop.
func incrementNonce(nonce []byte) {
	for i := len(nonce) - 1; i >= 0; i-- {
		nonce[i]++
		if nonce[i] != 0 {
			break
		}
	}
}

func (w *qhFcsZC) bBLmwPhjPp54(b U.Mf3phSOVqgf, payload [][]byte) { //0x1c54100
	if len(payload) == 1 {
		b.SealWrite(w.yn5oQQkf8, w.g9zKwmF, payload[0], nil)
	} else {
		start := b.Len()
		for _, part := range payload[:len(payload)-1] {
			_, _ = b.Write(part)
		}
		b.SealWriteWithStart(w.yn5oQQkf8, start, w.g9zKwmF, payload[len(payload)-1], nil)
	}
	incrementNonce(w.g9zKwmF)
}

// WriteSessionSegment @0x1c53820. Header fields are updated before locking.
// Only the error from the single connection Write is returned; a short nil
// Write is not converted to io.ErrShortWrite by the original.
func (w *qhFcsZC) WriteSessionSegment(m *IqB1foHWg, payload ...[]byte) error {
	size := 0
	for _, p := range payload {
		size += len(p)
	}
	padding := a2dHepfMMQ(size, w.rAB3kZi8)
	if limit := w.pattern.maxEnd(); limit >= 0 {
		padding = sessionPadding(size, w.rAB3kZi8, limit)
	}
	m.HPUJVQ0S = uint16(size)
	m.FO7MEJ = byte(len(padding))
	m.WyK3yKmJKe1V = uint32(time.Now().Unix() / 60)
	w.v_7InizyTC22.Lock()
	defer w.v_7InizyTC22.Unlock()
	header := m.Marshal()
	capacity := len(padding) + 48
	if w.oynKh2Fa {
		capacity += 24
	}
	if size > 0 {
		capacity += size + 16
	}
	if w.err != nil {
		return w.err
	}
	b := D.S4OQXWIZ(capacity)
	defer b.Free()
	if w.oynKh2Fa {
		_, _ = b.Write(w.g9zKwmF)
	}
	b.SealWrite(w.yn5oQQkf8, w.g9zKwmF, header[:], nil)
	incrementNonce(w.g9zKwmF)
	w.oynKh2Fa = false
	if size > 0 {
		w.bBLmwPhjPp54(b, payload)
	}
	_, _ = b.Write(padding)
	w.armDeadline()
	return w.fail(w.pattern.writeFragmented(w._z_j84, b.Bytes()))
}

// armDeadline 在持锁时刷新底层连接的写超时；每半个周期刷新一次，避免每帧一次系统调用。
func (w *qhFcsZC) armDeadline() {
	now := time.Now()
	if now.After(w.deadlineRefresh) {
		_ = w._z_j84.SetWriteDeadline(now.Add(writeTimeout))
		w.deadlineRefresh = now.Add(writeTimeout / 2)
	}
}

// fail 在持锁时记录第一次写错误并通知会话。
func (w *qhFcsZC) fail(err error) error {
	if err != nil && w.err == nil {
		w.err = err
		if w.onError != nil {
			go w.onError()
		}
	}
	return err
}

func (w *qhFcsZC) WriteDataSegment(m *IqB1foHWg, payload ...[]byte) error { //0x1c53c80
	prefix, suffix := byte(ofUaiB3MIo(256)), byte(ofUaiB3MIo(256))
	if limit := w.pattern.maxMiddle(); limit >= 0 {
		prefix = byte(min(ofUaiB3MIo(limit+1), limit))
	}
	if limit := w.pattern.maxEnd(); limit >= 0 {
		suffix = byte(min(ofUaiB3MIo(limit+1), limit))
	}
	size := 0
	for _, p := range payload {
		size += len(p)
	}
	wireSize := size
	lowEntropy := false
	if mode, rotation := w.pattern.lowEntropy(); mode != lowEntropyModeOff && m.Rf6N7kY6GVCs == 7 && w.clientLowEntropy.Load() {
		n, err := lowEntropyEncodedLen(size, mode)
		if err != nil {
			return err
		}
		lowEntropy, wireSize = true, n
		m.Rf6N7kY6GVCs = 11
		m.LowEntropyMode, m.LowEntropyRotation = mode, rotation
		m.LowEntropyMask = newLowEntropyHalfMask(mode)
		m.ExtractedPayloadLen = uint16(size)
	}
	m.B8QsA0N2yo = prefix
	m.HPUJVQ0S = uint16(wireSize)
	m.FO7MEJ = suffix
	m.WyK3yKmJKe1V = uint32(time.Now().Unix() / 60)
	w.v_7InizyTC22.Lock()
	defer w.v_7InizyTC22.Unlock()
	header := m.Marshal()
	capacity := int(prefix) + int(suffix) + 48
	if w.oynKh2Fa {
		capacity += 24
	}
	if size > 0 {
		capacity += wireSize + 16
	}
	if w.err != nil {
		return w.err
	}
	b := D.S4OQXWIZ(capacity)
	defer b.Free()
	if w.oynKh2Fa {
		_, _ = b.Write(w.g9zKwmF)
	}
	b.SealWrite(w.yn5oQQkf8, w.g9zKwmF, header[:], nil)
	incrementNonce(w.g9zKwmF)
	w.oynKh2Fa = false
	b.WriteRandom(int(prefix))
	if size > 0 && lowEntropy {
		// 低熵编码作用于密文主体，AEAD tag 原样附在后面。
		plain := make([]byte, 0, size)
		for _, p := range payload {
			plain = append(plain, p...)
		}
		sealed := w.yn5oQQkf8.Seal(nil, w.g9zKwmF, plain, nil)
		incrementNonce(w.g9zKwmF)
		_, _ = b.Write(encodeLowEntropy(sealed[:size], m.LowEntropyMode, m.LowEntropyMask, m.LowEntropyRotation))
		_, _ = b.Write(sealed[size:])
	} else if size > 0 {
		w.bBLmwPhjPp54(b, payload)
	}
	b.WriteRandom(int(suffix))
	w.armDeadline()
	_, err := w._z_j84.Write(b.Bytes())
	return w.fail(err)
}
