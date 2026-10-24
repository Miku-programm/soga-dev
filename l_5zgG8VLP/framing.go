package l_5zgG8VLP

import (
	"encoding/binary"
	"fmt"
	"io"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/G6bHD5JA"
)

// 以下三类型按原RTTI布局，分别为24/32/32字节。
type KhhcaT7 struct {
	io.Reader
	dyM549K_ [2]byte
}
type CZs766affqOH struct {
	io.Writer
	g7F22C    *ZfmpaHtmu
	zp2gBeBaA bool
}
type OuMHpcpymd1 struct {
	io.Writer
	g7F22C    *ZfmpaHtmu
	zp2gBeBaA bool
}

// iP6eRhiZ @0x1d07fe0。
func iP6eRhiZ(reader io.Reader, header *ZfmpaHtmu) (io.Reader, error) {
	switch header.aj8eMgsHQRY {
	case 1, 3:
		return reader, nil
	case 2:
		return &KhhcaT7{Reader: reader}, nil
	default:
		return nil, fmt.Errorf("unknown command: %v", header.aj8eMgsHQRY)
	}
}

// yPiHcQrp4Vs @0x1d081a0。
func yPiHcQrp4Vs(writer io.Writer, header *ZfmpaHtmu) (io.Writer, error) {
	switch header.aj8eMgsHQRY {
	case 1, 3:
		return &CZs766affqOH{Writer: writer, g7F22C: header}, nil
	case 2:
		return &OuMHpcpymd1{Writer: writer, g7F22C: header}, nil
	default:
		return nil, fmt.Errorf("unknown command: %v", header.aj8eMgsHQRY)
	}
}

// Write @0x1d083a0。首次写入前设置标记，失败后也不再次发送响应头。
// 首次底层Write无错误时返回完整payload长度，即使底层报告短写。
func (w *CZs766affqOH) Write(data []byte) (int, error) {
	if w.zp2gBeBaA {
		return w.Writer.Write(data)
	}
	w.zp2gBeBaA = true
	buffer := D_zCag1qTV.S4OQXWIZ(len(data) + 2)
	defer buffer.Free()
	w.g7F22C.WriteResponseHeader(buffer)
	_, _ = buffer.Write(data)
	if _, err := w.Writer.Write(buffer.Bytes()); err != nil {
		return 0, err
	}
	return len(data), nil
}

// Read @0x1d08b60。短目标缓冲会丢弃当前UDP帧剩余数据，保持下一帧对齐。
func (r *KhhcaT7) Read(data []byte) (int, error) {
	if _, err := io.ReadFull(r.Reader, r.dyM549K_[:]); err != nil {
		return 0, err
	}
	length := int(binary.BigEndian.Uint16(r.dyM549K_[:]))
	target := data
	if len(target) > length {
		target = target[:length]
	}
	n, err := io.ReadFull(r.Reader, target)
	if err != nil {
		return n, err
	}
	if n < length {
		if err := G6bHD5JA.KlTAqrF42mR(r.Reader, length-n); err != nil {
			return n, err
		}
	}
	return n, nil
}

// Write @0x1d08c80。原长度字段只保留低16位，但所有payload字节仍写入。
func (w *OuMHpcpymd1) Write(data []byte) (int, error) {
	var buffer *D_zCag1qTV.X_aXX5
	if w.zp2gBeBaA {
		buffer = D_zCag1qTV.S4OQXWIZ(len(data) + 2)
	} else {
		buffer = D_zCag1qTV.S4OQXWIZ(len(data) + 4)
		w.g7F22C.WriteResponseHeader(buffer)
		w.zp2gBeBaA = true
	}
	defer buffer.Free()
	buffer.WriteUint16(uint16(len(data)))
	_, _ = buffer.Write(data)
	if _, err := w.Writer.Write(buffer.Bytes()); err != nil {
		return 0, err
	}
	return len(data), nil
}
