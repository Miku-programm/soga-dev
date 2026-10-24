package pool

import (
	"encoding/binary"
	"io"
)

// ReadUint8 @ 0xe97f20 使用与多字节读取不同的空判断。
func (b *X_aXX5) ReadUint8() (byte, error) {
	if b.cnahaWq6 == b.__XpzgdOCbm {
		return 0, io.ErrUnexpectedEOF
	}
	value := b.kNItjuFiwnR[b.cnahaWq6]
	b.Skip(1)
	return value, nil
}

func (b *X_aXX5) WriteUint16(value uint16) {
	b.GrowBy(2)
	binary.BigEndian.PutUint16(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
	b.__XpzgdOCbm += 2
}
func (b *X_aXX5) ReadUint16() (uint16, error) {
	if b.__XpzgdOCbm-b.cnahaWq6 < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	value := binary.BigEndian.Uint16(b.kNItjuFiwnR[b.cnahaWq6:])
	b.Skip(2)
	return value, nil
}

func (b *X_aXX5) WriteUint32(value uint32) {
	b.GrowBy(4)
	binary.BigEndian.PutUint32(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
	b.__XpzgdOCbm += 4
}
func (b *X_aXX5) ReadUint32() (uint32, error) {
	if b.__XpzgdOCbm-b.cnahaWq6 < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	value := binary.BigEndian.Uint32(b.kNItjuFiwnR[b.cnahaWq6:])
	b.Skip(4)
	return value, nil
}

func (b *X_aXX5) WriteUint64(value uint64) {
	b.GrowBy(8)
	binary.BigEndian.PutUint64(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
	b.__XpzgdOCbm += 8
}
func (b *X_aXX5) ReadUint64() (uint64, error) {
	if b.__XpzgdOCbm-b.cnahaWq6 < 8 {
		return 0, io.ErrUnexpectedEOF
	}
	value := binary.BigEndian.Uint64(b.kNItjuFiwnR[b.cnahaWq6:])
	b.Skip(8)
	return value, nil
}

func (b *X_aXX5) WriteUint16Le(value uint16) {
	b.GrowBy(2)
	binary.LittleEndian.PutUint16(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
	b.__XpzgdOCbm += 2
}
func (b *X_aXX5) ReadUint16Le() (uint16, error) {
	if b.__XpzgdOCbm-b.cnahaWq6 < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	value := binary.LittleEndian.Uint16(b.kNItjuFiwnR[b.cnahaWq6:])
	b.Skip(2)
	return value, nil
}

func (b *X_aXX5) WriteUint32Le(value uint32) {
	b.GrowBy(4)
	binary.LittleEndian.PutUint32(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
	b.__XpzgdOCbm += 4
}
func (b *X_aXX5) ReadUint32Le() (uint32, error) {
	if b.__XpzgdOCbm-b.cnahaWq6 < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	value := binary.LittleEndian.Uint32(b.kNItjuFiwnR[b.cnahaWq6:])
	b.Skip(4)
	return value, nil
}

func (b *X_aXX5) WriteUint64Le(value uint64) {
	b.GrowBy(8)
	binary.LittleEndian.PutUint64(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
	b.__XpzgdOCbm += 8
}
func (b *X_aXX5) ReadUint64Le() (uint64, error) {
	if b.__XpzgdOCbm-b.cnahaWq6 < 8 {
		return 0, io.ErrUnexpectedEOF
	}
	value := binary.LittleEndian.Uint64(b.kNItjuFiwnR[b.cnahaWq6:])
	b.Skip(8)
	return value, nil
}
