package rtx

import (
	"encoding/binary"
	"io"
	"net"
)

// DCeuEa2b 来自 RTTI 0x3627b60 和 0xb41da0—0xb42400 的机器码；
// 这是二进制行为的源码表示，不是作者文件原文。显示类型前缀为 noIJiSXMP。
type DCeuEa2b struct {
	cTs3S5       []byte
	fAIkcTHidRIc int
}

// NewDCeuEa2b 是恢复工程的包间集成辅助；原调用方直接内联构造这两个字段。
func NewDCeuEa2b(data []byte) *DCeuEa2b { return &DCeuEa2b{cTs3S5: data} }
func (r *DCeuEa2b) LeftSize() int       { return len(r.cTs3S5) - r.fAIkcTHidRIc } //0xb41da0
func (r *DCeuEa2b) Len() int            { return len(r.cTs3S5) - r.fAIkcTHidRIc } //0xb41dc0
func (r *DCeuEa2b) IsEmpty() bool       { return r.LeftSize() <= 0 }              //0xb41de0
func (r *DCeuEa2b) ReadByte() (byte, error) { //0xb41e00
	if len(r.cTs3S5) < r.fAIkcTHidRIc+1 {
		return 0, io.ErrUnexpectedEOF
	}
	v := r.cTs3S5[r.fAIkcTHidRIc]
	r.fAIkcTHidRIc++
	return v, nil
}
func (r *DCeuEa2b) ReadUInt8() (uint8, error) { return r.ReadByte() } //0xb41e60
func (r *DCeuEa2b) ReadUInt16() (uint16, error) { //0xb41ec0
	b, e := r.ReadBytes(2)
	if e != nil {
		return 0, e
	}
	return binary.BigEndian.Uint16(b), nil
}
func (r *DCeuEa2b) ReadUInt32() (uint32, error) { //0xb41f20
	b, e := r.ReadBytes(4)
	if e != nil {
		return 0, e
	}
	return binary.BigEndian.Uint32(b), nil
}
func (r *DCeuEa2b) ReadUInt64() (uint64, error) { //0xb41f80
	b, e := r.ReadBytes(8)
	if e != nil {
		return 0, e
	}
	return binary.BigEndian.Uint64(b), nil
}
func (r *DCeuEa2b) ReadLeft() []byte { //0xb41fe0
	if r.LeftSize() <= 0 {
		return []byte{}
	}
	b := r.cTs3S5[r.fAIkcTHidRIc:]
	r.fAIkcTHidRIc = len(r.cTs3S5)
	return b
}
func (r *DCeuEa2b) ReadIPv4() (string, error) { //0xb42040
	b, e := r.ReadBytes(4)
	if e != nil {
		return "", e
	}
	return net.IP(b).String(), nil
}
func (r *DCeuEa2b) ReadIPv6() (string, error) { //0xb420e0
	b, e := r.ReadBytes(16)
	if e != nil {
		return "", e
	}
	return net.IP(b).String(), nil
}
func (r *DCeuEa2b) ReadBytes(size int) ([]byte, error) { //0xb42180
	end := r.fAIkcTHidRIc + size
	if len(r.cTs3S5) < end {
		return nil, io.ErrUnexpectedEOF
	}
	b := r.cTs3S5[r.fAIkcTHidRIc:end]
	r.fAIkcTHidRIc = end
	return b, nil
}
func (r *DCeuEa2b) Read(dst []byte) (int, error) { //0xb42200
	if r.LeftSize() <= 0 {
		return 0, io.EOF
	}
	n := copy(dst, r.cTs3S5[r.fAIkcTHidRIc:])
	r.fAIkcTHidRIc += n
	return n, nil
}
func (r *DCeuEa2b) Skip(size int) error { //0xb422e0
	end := r.fAIkcTHidRIc + size
	if len(r.cTs3S5) < end {
		return io.ErrUnexpectedEOF
	}
	r.fAIkcTHidRIc = end
	return nil
}
func (r *DCeuEa2b) ReadString(size int) (string, error) { //0xb42320
	end := r.fAIkcTHidRIc + size
	if len(r.cTs3S5) < end {
		return "", io.ErrUnexpectedEOF
	}
	s := string(r.cTs3S5[r.fAIkcTHidRIc:end])
	r.fAIkcTHidRIc = end
	return s, nil
}
func (r *DCeuEa2b) AllData() []byte { return r.cTs3S5 } //0xb42400
