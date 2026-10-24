// 按原 ELF 0x1c28820..0x1c29a80 恢复。错误文本来自受限机器码解码；
// quicvarint 的原七个保留函数与 apernet/quic-go cb103fcecfe 的机器码已逐条核对。
package mlbRsO44

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/apernet/quic-go/quicvarint"
	"io"
)

// ItLMumzBf 对应 RTTI wqmlKj.ItLMumzBf，原尺寸 48 字节。
type ItLMumzBf struct {
	D35ycLc5aB6r uint32
	CBHlJenOD    uint16
	QmDvxK_8Q    uint8
	Kb9FLg7wuNE  uint8
	XrDcBXcD     string
	HmoFPk       []byte
}

// T7aXhP 对应 0x1c28820。
func T7aXhP(r io.Reader) (string, error) {
	br := quicvarint.NewReader(r)
	n, err := quicvarint.Read(br)
	if err != nil {
		return "", err
	}
	if n == 0 || n > 2048 {
		return "", errors.New("invalid address length")
	}
	b := make([]byte, n)
	if _, err = io.ReadFull(r, b); err != nil {
		return "", err
	}
	padding, err := quicvarint.Read(br)
	if err != nil {
		return "", err
	}
	if padding > 4096 {
		return "", errors.New("invalid padding length")
	}
	if padding > 0 {
		if _, err = io.CopyN(io.Discard, r, int64(padding)); err != nil {
			return "", err
		}
	}
	return string(b), nil
}

// C3FDNGT2zI 对应 0x1c28c40；原函数只返回 Write 的 error，不检查短写。
func C3FDNGT2zI(w io.Writer, ok bool, msg string) error {
	padding := responsePaddingAt3789490.String()
	b := make([]byte, 1+quicvarint.Len(uint64(len(msg)))+len(msg)+quicvarint.Len(uint64(len(padding)))+len(padding))
	if !ok {
		b[0] = 1
	}
	i := 1
	i += vVaooWJ93Bp(b[i:], uint64(len(msg)))
	i += copy(b[i:], msg)
	i += vVaooWJ93Bp(b[i:], uint64(len(padding)))
	copy(b[i:], padding)
	_, err := w.Write(b)
	return err
}

// HeaderSize / Size / Serialize 分别对应 0x1c29000 / 0x1c290c0 / 0x1c29180。
func (m *ItLMumzBf) HeaderSize() int {
	return 8 + quicvarint.Len(uint64(len(m.XrDcBXcD))) + len(m.XrDcBXcD)
}
func (m *ItLMumzBf) Size() int { return m.HeaderSize() + len(m.HmoFPk) }
func (m *ItLMumzBf) Serialize(b []byte) int {
	if len(b) < m.Size() {
		return -1
	}
	binary.BigEndian.PutUint32(b, m.D35ycLc5aB6r)
	binary.BigEndian.PutUint16(b[4:], m.CBHlJenOD)
	b[6], b[7] = m.QmDvxK_8Q, m.Kb9FLg7wuNE
	i := 8
	i += vVaooWJ93Bp(b[i:], uint64(len(m.XrDcBXcD)))
	i += copy(b[i:], m.XrDcBXcD)
	i += copy(b[i:], m.HmoFPk)
	return i
}

// CY9jLhcY1f5 对应 0x1c29440。原实现使用 bytes.Buffer 与四次 binary.Read；
// 不接受空 UDP 载荷，不在此处检查分片序号。Data 与输入共用底层数组。
func CY9jLhcY1f5(data []byte) (*ItLMumzBf, error) {
	m := new(ItLMumzBf)
	b := bytes.NewBuffer(data)
	if err := binary.Read(b, binary.BigEndian, &m.D35ycLc5aB6r); err != nil {
		return nil, err
	}
	if err := binary.Read(b, binary.BigEndian, &m.CBHlJenOD); err != nil {
		return nil, err
	}
	if err := binary.Read(b, binary.BigEndian, &m.QmDvxK_8Q); err != nil {
		return nil, err
	}
	if err := binary.Read(b, binary.BigEndian, &m.Kb9FLg7wuNE); err != nil {
		return nil, err
	}
	n, err := quicvarint.Read(b)
	if err != nil {
		return nil, err
	}
	if n == 0 || n > 2048 {
		return nil, errors.New("invalid address length")
	}
	rest := b.Bytes()
	if n >= uint64(len(rest)) {
		return nil, errors.New("invalid message length")
	}
	m.XrDcBXcD = string(rest[:n])
	m.HmoFPk = rest[n:]
	return m, nil
}

// vVaooWJ93Bp 对应 0x1c29860；逐字节赋值保留短缓冲区 panic 前的写入。
func vVaooWJ93Bp(b []byte, v uint64) int {
	if v <= 63 {
		b[0] = byte(v)
		return 1
	}
	if v <= 16383 {
		b[0] = byte(v>>8) | 0x40
		b[1] = byte(v)
		return 2
	}
	if v <= 1073741823 {
		b[0] = byte(v>>24) | 0x80
		b[1] = byte(v >> 16)
		b[2] = byte(v >> 8)
		b[3] = byte(v)
		return 4
	}
	if v <= 4611686018427387903 {
		b[0] = byte(v>>56) | 0xc0
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
		b[3] = byte(v >> 32)
		b[4] = byte(v >> 24)
		b[5] = byte(v >> 16)
		b[6] = byte(v >> 8)
		b[7] = byte(v)
		return 8
	}
	panic(fmt.Sprintf("%#x doesn't fit into 62 bits", v))
}
