package qwn0vjA

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/G6bHD5JA"
)

// 以下错误和盐来自原 init@0x1872540 的受限字符串/字节解码仿真。
var errorAt5098380 = errors.New("not quic")
var errorAt5098390 = errors.New("not initial packet")
var saltAt509A1B0 = []byte{0xaf, 0xbf, 0xec, 0x28, 0x99, 0x93, 0xd2, 0x4c, 0x9e, 0x97, 0x86, 0xf1, 0x9c, 0x61, 0x11, 0xe0, 0x43, 0x90, 0xa8, 0x99}
var saltAt509A1D0 = []byte{0x38, 0x76, 0x2c, 0xf7, 0xf5, 0x59, 0x34, 0xb3, 0x4d, 0x17, 0x9a, 0xe6, 0xa4, 0xc8, 0x0c, 0xad, 0xcc, 0xbb, 0x7f, 0x0a}

type naaHnaatZCex struct{}
type jO2kbfaNQJ struct{ dunzUGo string }

func (*jO2kbfaNQJ) Type() MCe3PxlaMYM          { return "quic" }    //0x1876400
func (v *jO2kbfaNQJ) Value() string            { return v.dunzUGo } //0x1876420
func (*naaHnaatZCex) Type() MCe3PxlaMYM        { return "quic" }    //0x1876440
func (*naaHnaatZCex) Network() G6bHD5JA.CjhtWb { return 1 }         //0x1876460

// Sniff 对应 0x1876480 及恢复 panic 的 func1。保留原实现的版本、包号、
// 帧类型限制及输入原地修改行为；不是完整 QUIC 传输实现，也不是作者源码原文。
func (*naaHnaatZCex) Sniff(data []byte) (result FlCuFjj_V1u, err error) {
	defer func() {
		if recover() != nil {
			err = errorAt5098380
		}
	}()
	r := CAodbhZf.NewDCeuEa2b(data)
	first, e := r.ReadUInt8()
	if e != nil {
		return nil, errorAt5098380
	}
	if first&0x80 == 0 || first&0x40 == 0 {
		return nil, errorAt5098390
	}
	version, e := r.ReadUInt32()
	if e != nil {
		return nil, errorAt5098380
	}
	if version != 1 && version != 0xff00001d {
		return nil, errorAt5098380
	}
	if (first&0x30)>>4 != 0 {
		return nil, errorAt5098390
	}
	dcidLength, e := r.ReadUInt8()
	if e != nil {
		return nil, errorAt5098380
	}
	dcid, e := r.ReadBytes(int(dcidLength))
	if e != nil {
		return nil, errorAt5098380
	}
	scidLength, e := r.ReadUInt8()
	if e != nil {
		return nil, errorAt5098380
	}
	if _, e = r.ReadBytes(int(scidLength)); e != nil {
		return nil, errorAt5098380
	}
	tokenLength, e := vFUaVI1hy(r)
	if e != nil || tokenLength > uint64(len(data)) {
		return nil, errorAt5098380
	}
	if _, e = r.ReadBytes(int(tokenLength)); e != nil {
		return nil, errorAt5098380
	}
	packetLength, e := vFUaVI1hy(r)
	if e != nil {
		return nil, errorAt5098380
	}
	pnOffset := len(data) - r.LeftSize()
	var saved [4]byte
	copy(saved[:], data[pnOffset:pnOffset+4])
	salt := saltAt509A1B0
	if version == 1 {
		salt = saltAt509A1D0
	}
	initial := hkdf.Extract(crypto.SHA256.New, dcid, salt)
	secret := qhXufgHMZAdy(crypto.SHA256, initial, nil, "client in", 32)
	hp := qhXufgHMZAdy(crypto.SHA256, secret, nil, "quic hp", 16)
	block, e := aes.NewCipher(hp)
	if e != nil {
		return nil, e
	}
	scratch := D_zCag1qTV.S4OQXWIZ(64)
	scratch.SkipEnd(block.BlockSize())
	mask := scratch.BytesTo(block.BlockSize())
	block.Encrypt(mask, data[pnOffset+4:pnOffset+20])
	data[0] ^= mask[0] & 15
	for i := 0; i < 4; i++ {
		data[pnOffset+i] ^= mask[i+1]
	}
	if data[0]&3 != 0 {
		return nil, errorAt5098390
	}
	pn, e := r.ReadUInt8()
	if e != nil {
		return nil, e
	}
	if pn|1 != 1 {
		return nil, errorAt5098390
	}
	copy(data[pnOffset+1:pnOffset+4], saved[1:])
	encrypted := data[pnOffset+1 : uint64(pnOffset)+packetLength]
	key := qhXufgHMZAdy(crypto.SHA256, secret, nil, "quic key", 16)
	iv := qhXufgHMZAdy(crypto.SHA256, secret, nil, "quic iv", 12)
	// 原 crypto/tls.aeadAESGCMTLS13 使用固定 IV 异或 8 字节包序号。
	// scratch 的掩码和序号分配保留；原函数没有释放该临时缓冲区。
	keyBlock, e := aes.NewCipher(key)
	if e != nil {
		panic(e)
	}
	aead, e := cipher.NewGCM(keyBlock)
	if e != nil {
		panic(e)
	}
	scratch.SkipEnd(8)
	sequence := scratch.BytesRange(block.BlockSize(), block.BlockSize()+8)
	clear(sequence)
	sequence[7] = pn
	for i := range sequence {
		iv[4+i] ^= sequence[i]
	}
	plain, e := aead.Open(encrypted[:0], iv, encrypted, data[:pnOffset+1])
	if e != nil {
		return nil, e
	}
	r = CAodbhZf.NewDCeuEa2b(plain)
	assembly := D_zCag1qTV.PoolAt50961D8.Get(r.LeftSize())
	defer D_zCag1qTV.PoolAt50961D8.Put(assembly)
	var maxEnd uint64
	for !r.IsEmpty() {
		var kind byte
		for !r.IsEmpty() {
			kind, _ = r.ReadUInt8()
			if kind != 0 {
				break
			}
		}
		switch kind {
		case 0, 1:
		case 2, 3:
			if _, e = vFUaVI1hy(r); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			if _, e = vFUaVI1hy(r); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			count, e := vFUaVI1hy(r)
			if e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			if _, e = vFUaVI1hy(r); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			for ; count > 0; count-- {
				if _, e = vFUaVI1hy(r); e != nil {
					return nil, io.ErrUnexpectedEOF
				}
				if _, e = vFUaVI1hy(r); e != nil {
					return nil, io.ErrUnexpectedEOF
				}
			}
			if kind == 3 {
				for i := 0; i < 3; i++ {
					if _, e = vFUaVI1hy(r); e != nil {
						return nil, io.ErrUnexpectedEOF
					}
				}
			}
		case 6:
			offset, e := vFUaVI1hy(r)
			if e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			length, e := vFUaVI1hy(r)
			if e != nil || length > uint64(r.LeftSize()) {
				return nil, io.ErrUnexpectedEOF
			}
			end := offset + length
			if end > maxEnd {
				maxEnd = end
			}
			if _, e = r.Read(assembly[offset:end]); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
		case 28:
			if _, e = vFUaVI1hy(r); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			if _, e = vFUaVI1hy(r); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			reasonLength, e := vFUaVI1hy(r)
			if e != nil {
				return nil, io.ErrUnexpectedEOF
			}
			if _, e = r.ReadBytes(int(reasonLength)); e != nil {
				return nil, io.ErrUnexpectedEOF
			}
		default:
			return nil, errorAt5098390
		}
	}
	var host yLT2n11mN
	if e = mCawThJ(assembly[:maxEnd], &host); e != nil {
		return nil, e
	}
	return &jO2kbfaNQJ{dunzUGo: host.dunzUGo}, nil
}

// NHzYZHDSd 对应 0x1872c40；原名称逐字匹配，不做大小写或空格归一化。
func NHzYZHDSd(kind MCe3PxlaMYM) (S5TGvY4, error) {
	switch kind {
	case "bt":
		return &ezgFiQP115F{}, nil
	case "http":
		return &q7GP8uwJQpnC{}, nil
	case "tls":
		return &aln4LlVj3z{}, nil
	case "quic":
		return &naaHnaatZCex{}, nil
	default:
		return nil, fmt.Errorf("sniffer type <%v> not exist", kind)
	}
}
