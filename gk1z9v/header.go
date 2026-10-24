package gk1z9v

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"net"

	B "recovered.local/soga/D_zCag1qTV"
	G "recovered.local/soga/G6bHD5JA"
	J "recovered.local/soga/J1P4KNvfI_J"
)

type IPas09 uint8
type OtMMIvjqe uint8

// 原请求/响应头 RTTI 分别为 72/56 字节。
type Wpqfspy6zZ_B struct {
	v2VYs_ttaea  uint8
	cEcPFEa7     []byte
	ezCHRvHCePBC []byte
	daehRKryOT   uint8
	wadzfAUcy    uint8
	gnVzkTY      bool
	btI8Xvxnh    bool
	p4hZmT2VBll5 bool
	uaHWu6c      uint8
	mTRhOfKViN   IPas09
	uycGVy       OtMMIvjqe
	gicXHc       *G.AHU2MRmjzQu
}
type URTFuf2wVWf struct {
	iQlcjP47sl uint8
	mFdVdC     uint8
	j3mthjsl   uint8
	maCTf0DLoP []byte
	iqQ8HJ6OMB []byte
}

// ELa2jh3VBB @0x1d1c340：AEAD 认证用 SHA-256 的前 16 字节，旧认证用 MD5。
func ELa2jh3VBB(request *Wpqfspy6zZ_B, auth *WR4lMNfab) *URTFuf2wVWf {
	iv, key := make([]byte, 16), make([]byte, 16)
	if auth.Ih7r8HyMvOa {
		a, b := sha256.Sum256(request.cEcPFEa7), sha256.Sum256(request.ezCHRvHCePBC)
		copy(iv, a[:16])
		copy(key, b[:16])
	} else {
		a, b := md5.Sum(request.cEcPFEa7), md5.Sum(request.ezCHRvHCePBC)
		copy(iv, a[:])
		copy(key, b[:])
	}
	return &URTFuf2wVWf{iQlcjP47sl: request.daehRKryOT, maCTf0DLoP: iv, iqQ8HJ6OMB: key}
}

// Gk1af7wDMnD @0x1d15bc0 读取已解密的请求头；认证/解密由原上层连接函数执行。
func Gk1af7wDMnD(r io.Reader) (*Wpqfspy6zZ_B, error) {
	data := make([]byte, 128)
	if _, err := io.ReadFull(r, data[:38]); err != nil {
		return nil, err
	}
	if data[0] != 1 {
		return nil, fmt.Errorf("VMess version is not 1: %v", data[0])
	}
	iv := append([]byte(nil), data[1:17]...)
	key := append([]byte(nil), data[17:33]...)
	response, options, security, command := data[33], data[34], data[35], data[37]
	padding := int(security >> 4)
	var address *G.AHU2MRmjzQu
	var total int
	if command == 3 {
		total = 38 + padding + 4
		if _, err := io.ReadFull(r, data[38:total]); err != nil {
			return nil, err
		}
	} else {
		if _, err := io.ReadFull(r, data[38:41]); err != nil {
			return nil, err
		}
		port := int(binary.BigEndian.Uint16(data[38:40]))
		kind := data[40]
		var start, end int
		switch kind {
		case 1:
			start, end = 41, 45
		case 2:
			if _, err := io.ReadFull(r, data[41:42]); err != nil {
				return nil, err
			}
			start, end = 42, 42+int(data[41])
		case 3:
			start, end = 41, 57
		default:
			return nil, fmt.Errorf("unknown addr type: %v", kind)
		}
		total = end + padding + 4
		if total > 128 {
			large := B.PoolAt50961D8.Get(total)
			defer B.PoolAt50961D8.Put(large)
			copy(large, data[:start])
			data = large
		}
		if _, err := io.ReadFull(r, data[start:total]); err != nil {
			return nil, err
		}
		if kind == 2 {
			address = G.NewDomainAddress(string(data[start:end]), port)
		} else {
			ip := append(net.IP(nil), data[start:end]...)
			// 原此处内联构造与 RNFJeaRY 的 IP 分支字段相同。
			address, _ = G.RNFJeaRY(&net.TCPAddr{IP: ip, Port: port})
		}
	}
	h := fnv.New32a()
	h.Write(data[:total-4])
	if binary.BigEndian.Uint32(data[total-4:total]) != h.Sum32() {
		return nil, errors.New("VMess Header fnv invalid")
	}
	return &Wpqfspy6zZ_B{v2VYs_ttaea: 1, cEcPFEa7: iv, ezCHRvHCePBC: key, daehRKryOT: response, wadzfAUcy: options,
		gnVzkTY: options&1 != 0, btI8Xvxnh: options&1 != 0 && options&4 != 0, p4hZmT2VBll5: options&8 != 0,
		uaHWu6c: uint8(padding), mTRhOfKViN: IPas09(security & 15), uycGVy: OtMMIvjqe(command), gicXHc: address}, nil
}
func (h *Wpqfspy6zZ_B) OptS() bool                  { return h.gnVzkTY }      // 0x1d168a0
func (h *Wpqfspy6zZ_B) OptM() bool                  { return h.btI8Xvxnh }    // 0x1d168c0
func (h *Wpqfspy6zZ_B) OptPadding() bool            { return h.p4hZmT2VBll5 } // 0x1d168e0
func (h *Wpqfspy6zZ_B) GetDestAddr() *G.AHU2MRmjzQu { return h.gicXHc }       // 0x1d16900
func (h *Wpqfspy6zZ_B) NetworkType() J.ZES6F5 { // 0x1d16920
	switch h.uycGVy {
	case 1:
		return 0
	case 2:
		return 1
	case 3:
		return 2
	default:
		return 255
	}
}
func (h *URTFuf2wVWf) HeaderLen() int  { return 4 }                                             // 0x1d1c580
func (h *URTFuf2wVWf) ToBytes() []byte { return []byte{h.iQlcjP47sl, h.mFdVdC, h.j3mthjsl, 0} } // 0x1d1c5a0
