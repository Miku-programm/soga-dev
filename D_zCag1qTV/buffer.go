package D_zCag1qTV

import (
	"crypto/rand"
	"io"
)

// 原 RTTI dy6KhjOW1Tn.X_aXX5：40 字节，切片/读位置/写位置偏移 0/24/32。
type X_aXX5 struct {
	kNItjuFiwnR []byte
	cnahaWq6    int
	__XpzgdOCbm int
}

// S4OQXWIZ @ 0xe96c80；底层切片长度扩为池返回的容量，可用数据初始为零。
func S4OQXWIZ(size int) *X_aXX5 {
	data := PoolAt50961D8.Get(size)
	return &X_aXX5{kNItjuFiwnR: data[:cap(data)]}
}

// LdHEmSs9 @ 0xe96d20 按值返回 40 字节结构；ASM 的五个返回寄存器
// 是 array/cap/cap/0/0，Hex-Rays 将该返回值误标成了单个指针。
func LdHEmSs9(size int) X_aXX5 {
	data := PoolAt50961D8.Get(size)
	return X_aXX5{kNItjuFiwnR: data[:cap(data)]}
}

func (b *X_aXX5) Reset()                        { b.cnahaWq6 = 0; b.__XpzgdOCbm = 0 }              // 0xe96da0
func (b *X_aXX5) Byte(index int) byte           { return b.kNItjuFiwnR[b.cnahaWq6+index] }         // 0xe96dc0
func (b *X_aXX5) SetByte(index int, value byte) { b.kNItjuFiwnR[b.cnahaWq6+index] = value }        // 0xe96de0
func (b *X_aXX5) Bytes() []byte                 { return b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm] } // 0xe96e00

// GrowTo @ 0xe96e60；仅真正扩容时压紧剩余内容。
func (b *X_aXX5) GrowTo(size int) {
	if len(b.kNItjuFiwnR) < size {
		remaining := b.__XpzgdOCbm - b.cnahaWq6
		data := PoolAt50961D8.Get(size)
		data = data[:cap(data)]
		copy(data, b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm])
		PoolAt50961D8.Put(b.kNItjuFiwnR)
		b.kNItjuFiwnR = data
		b.cnahaWq6 = 0
		b.__XpzgdOCbm = remaining
	}
}

// GrowBy @ 0xe96fc0；扩容请求使用旧切片长度+增量，非剩余长度+增量。
func (b *X_aXX5) GrowBy(size int) {
	if len(b.kNItjuFiwnR) < b.__XpzgdOCbm+size {
		remaining := b.__XpzgdOCbm - b.cnahaWq6
		if len(b.kNItjuFiwnR) < remaining+size {
			data := PoolAt50961D8.Get(len(b.kNItjuFiwnR) + size)
			data = data[:cap(data)]
			copy(data, b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm])
			PoolAt50961D8.Put(b.kNItjuFiwnR)
			b.kNItjuFiwnR = data
		} else {
			copy(b.kNItjuFiwnR, b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm])
		}
		b.cnahaWq6 = 0
		b.__XpzgdOCbm = remaining
	}
}

// 原切片方法按底层容量检查，不把 end 限制为当前可读长度。
func (b *X_aXX5) BytesRange(start, end int) []byte {
	return b.kNItjuFiwnR[b.cnahaWq6+start : b.cnahaWq6+end]
}                                            // 0xe97180
func (b *X_aXX5) BytesFrom(start int) []byte { return b.kNItjuFiwnR[b.cnahaWq6+start : b.__XpzgdOCbm] } // 0xe971e0
func (b *X_aXX5) BytesTo(end int) []byte     { return b.kNItjuFiwnR[b.cnahaWq6 : b.cnahaWq6+end] }      // 0xe97240
func (b *X_aXX5) Skip(size int) {
	b.cnahaWq6 += size
	if b.__XpzgdOCbm <= b.cnahaWq6 {
		b.Reset()
	}
}                                  // 0xe97280
func (b *X_aXX5) SkipEnd(size int) { b.GrowBy(size); b.__XpzgdOCbm += size } // 0xe972a0
func (b *X_aXX5) ReverseEnd(size int) int { // 0xe97300
	b.__XpzgdOCbm -= size
	if b.cnahaWq6 > b.__XpzgdOCbm {
		size -= b.cnahaWq6 - b.__XpzgdOCbm
		b.__XpzgdOCbm = b.cnahaWq6
	}
	return size
}
func (b *X_aXX5) Len() int      { return b.__XpzgdOCbm - b.cnahaWq6 }  // 0xe97340
func (b *X_aXX5) Cap() int      { return len(b.kNItjuFiwnR) }          // 0xe97360
func (b *X_aXX5) IsEmpty() bool { return b.cnahaWq6 == b.__XpzgdOCbm } // 0xe97380

func (b *X_aXX5) Write(data []byte) (int, error) { // 0xe973a0
	if len(data) == 0 {
		return 0, nil
	}
	b.GrowBy(len(data))
	n := copy(b.kNItjuFiwnR[b.__XpzgdOCbm:], data)
	b.__XpzgdOCbm += n
	return n, nil
}
func (b *X_aXX5) WriteRandom(size int) []byte { // 0xe974a0
	b.GrowBy(size)
	data := b.kNItjuFiwnR[b.__XpzgdOCbm : b.__XpzgdOCbm+size]
	_, _ = rand.Read(data)
	b.__XpzgdOCbm += size
	return data
}
func (b *X_aXX5) WriteByte(value byte) {
	b.GrowBy(1)
	b.kNItjuFiwnR[b.__XpzgdOCbm] = value
	b.__XpzgdOCbm++
} // 0xe97560
func (b *X_aXX5) WriteString(value string) { // 0xe979a0
	b.GrowBy(len(value))
	b.__XpzgdOCbm += copy(b.kNItjuFiwnR[b.__XpzgdOCbm:], value)
}
func (b *X_aXX5) Read(data []byte) (int, error) { // 0xe97a60
	if b.cnahaWq6 == b.__XpzgdOCbm {
		return 0, io.EOF
	}
	n := copy(data, b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm])
	b.Skip(n)
	return n, nil
}
func (b *X_aXX5) ReadBytes(size int) ([]byte, error) { // 0xe97b40
	if size > b.__XpzgdOCbm-b.cnahaWq6 {
		return nil, io.ErrUnexpectedEOF
	}
	data := b.kNItjuFiwnR[b.cnahaWq6 : b.cnahaWq6+size]
	b.Skip(size)
	return data, nil
}

// 原指针类型与缓冲接口的 RTTI 均确认返回 int；它不是 io.ReaderFrom。
func (b *X_aXX5) ReadFrom(reader io.Reader) (int, error) {
	return b.ReadAllFrom(reader, 4096)
} // 0xe97bc0
func (b *X_aXX5) ReadAllFrom(reader io.Reader, increment int) (int, error) { // 0xe97c20
	total := 0
	for {
		b.GrowBy(increment)
		n, err := reader.Read(b.kNItjuFiwnR[b.__XpzgdOCbm:])
		b.__XpzgdOCbm += n
		total += n
		if err == io.EOF {
			return total, nil
		}
		if err != nil {
			return total, err
		}
	}
}
func (b *X_aXX5) ReadFullFrom(reader io.Reader, size int) (int, error) { // 0xe97d60
	b.GrowBy(size)
	n, err := io.ReadFull(reader, b.kNItjuFiwnR[b.__XpzgdOCbm:b.__XpzgdOCbm+size])
	b.__XpzgdOCbm += n
	return n, err
}
func (b *X_aXX5) ReadMaxFrom(reader io.Reader, size int) (int, error) { // 0xe97e40
	b.GrowBy(size)
	n, err := reader.Read(b.kNItjuFiwnR[b.__XpzgdOCbm : b.__XpzgdOCbm+size])
	b.__XpzgdOCbm += n
	return n, err
}
func (b *X_aXX5) String() string { return string(b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm]) } // 0xe982c0
func (b *X_aXX5) Free() { // 0xe98ee0
	if b.kNItjuFiwnR != nil {
		PoolAt50961D8.Put(b.kNItjuFiwnR)
		b.kNItjuFiwnR = nil
		b.Reset()
	}
}
