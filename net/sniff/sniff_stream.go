package sniff

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/errutil"
	"strings"
)

// 错误文本由原 init@0x1872540 的字符串解码器静态仿真取得。
var errorAt5098350 = errors.New("sniff failed")
var errorAt5098360 = errors.New("not bt protocol")
var errorAt5098370 = errors.New("not an HTTP method")
var errorAt50983A0 = errors.New("not enough data")
var errorAt50983B0 = errors.New("not client hello")
var errorAt50983C0 = errors.New("not TLS header")

// Sniff 对应 0x1872b00；只运行指定网络的嗅探器，忽略各次失败，返回首个无错误结果。
func (s KZgBMVcpL) Sniff(data []byte, network relay.CjhtWb) (FlCuFjj_V1u, error) {
	for _, v := range s {
		if v.Network() == network {
			if result, err := v.Sniff(data); err == nil {
				return result, nil
			}
		}
	}
	return nil, errorAt5098350
}

type ezgFiQP115F struct{}
type dHZyaNv0Y2Bz struct{}

func (*dHZyaNv0Y2Bz) Type() MCe3PxlaMYM       { return "bt" } //0x18740e0
func (*dHZyaNv0Y2Bz) Value() string           { return "bt" } //0x1874100
func (*ezgFiQP115F) Type() MCe3PxlaMYM        { return "bt" } //0x1874120
func (*ezgFiQP115F) Network() relay.CjhtWb { return 0 }    //0x1874140
// Sniff 对应 0x1874160，仅核对前20字节，后续握手字段不参与此嗅探。
func (*ezgFiQP115F) Sniff(data []byte) (FlCuFjj_V1u, error) {
	if len(data) >= 20 && data[0] == 19 && string(data[1:20]) == "BitTorrent protocol" {
		return &dHZyaNv0Y2Bz{}, nil
	}
	return nil, errorAt5098360
}

type q7GP8uwJQpnC struct{}
type q_HzS3 struct{ hpf1qO2oj string }

func (*q_HzS3) Type() MCe3PxlaMYM              { return "http" }      //0x18755e0
func (v *q_HzS3) Value() string                { return v.hpf1qO2oj } //0x1875600
func (*q7GP8uwJQpnC) Type() MCe3PxlaMYM        { return "http" }      //0x1875620
func (*q7GP8uwJQpnC) Network() relay.CjhtWb { return 0 }           //0x1875640
// ParseHost 对应 0x1875660；仅忽略 AddrError.Err 包含 missing port 的错误。
func (*q7GP8uwJQpnC) ParseHost(host string) (string, error) {
	parsed, _, err := net.SplitHostPort(host)
	if err == nil {
		return parsed, nil
	}
	if addrErr, ok := err.(*net.AddrError); ok && strings.Contains(addrErr.Err, "missing port") {
		return host, nil
	}
	return "", err
}

// 原静态 off_50368f0 数组按此顺序保存这些方法前缀（没有要求随后是空格）。
var methodPrefixesAt50368F0 = [][]byte{[]byte("get"), []byte("post"), []byte("head"), []byte("put"), []byte("delete"), []byte("options"), []byte("connect")}

// tTU63Fb 对应 0x1875760。数据短于当前前缀即返回原共享 already free 错误。
func (*q7GP8uwJQpnC) tTU63Fb(data []byte) error {
	for _, method := range methodPrefixesAt50368F0 {
		if len(data) < len(method) {
			return errutil.ErrorAt5097810
		}
		if bytes.EqualFold(data[:len(method)], method) {
			return nil
		}
	}
	return errorAt5098370
}

// Sniff 对应 0x1875860。以单个换行拆分；只在空字节切片处停止，
// 因而 CRLF 的空行内容仍有一个 CR，原函数会继续检查后续行。
func (s *q7GP8uwJQpnC) Sniff(data []byte) (FlCuFjj_V1u, error) {
	if err := s.tTU63Fb(data); err != nil {
		return nil, err
	}
	result := &q_HzS3{}
	for _, line := range bytes.Split(data, []byte{'\n'})[1:] {
		if len(line) == 0 {
			break
		}
		parts := bytes.SplitN(line, []byte{':'}, 2)
		if len(parts) == 2 && strings.ToLower(string(parts[0])) == "host" {
			host, err := s.ParseHost(strings.ToLower(string(bytes.TrimSpace(parts[1]))))
			if err != nil {
				return nil, err
			}
			result.hpf1qO2oj = host
		}
	}
	if result.hpf1qO2oj == "" {
		return nil, errutil.ErrorAt5097820
	}
	return result, nil
}

type aln4LlVj3z struct{}
type yLT2n11mN struct{ dunzUGo string }

func (*yLT2n11mN) Type() MCe3PxlaMYM         { return "tls" }     //0x187aa20
func (v *yLT2n11mN) Value() string           { return v.dunzUGo } //0x187aa40
func (*aln4LlVj3z) Type() MCe3PxlaMYM        { return "tls" }     //0x187aa60
func (*aln4LlVj3z) Network() relay.CjhtWb { return 0 }         //0x187aa80
// Sniff 对应 0x187aea0；首条记录需要完整，后续记录不参与。
func (*aln4LlVj3z) Sniff(data []byte) (FlCuFjj_V1u, error) {
	if len(data) < 5 {
		return nil, errorAt50983A0
	}
	if data[0] != 22 || data[1] != 3 {
		return nil, errorAt50983C0
	}
	n := int(binary.BigEndian.Uint16(data[3:5]))
	if len(data) < n+5 {
		return nil, errorAt50983A0
	}
	r := &yLT2n11mN{}
	if err := mCawThJ(data[5:5+n], r); err != nil {
		return nil, err
	}
	return r, nil
}

// mCawThJ 对应 0x187aaa0。保留原边界分类、首个 host_name 返回，
// 不额外校验握手类型/握手长度，且空名称可成功、末尾点号会失败。
func mCawThJ(data []byte, result *yLT2n11mN) error {
	if len(data) < 42 {
		return errorAt50983A0
	}
	sessionLen := int(data[38])
	if sessionLen > 32 || len(data) < 39+sessionLen {
		return errorAt50983A0
	}
	data = data[39+sessionLen:]
	if len(data) < 2 {
		return errorAt50983A0
	}
	cipherLen := int(binary.BigEndian.Uint16(data))
	if cipherLen%2 != 0 || len(data) < 2+cipherLen {
		return errorAt50983B0
	}
	data = data[2+cipherLen:]
	if len(data) == 0 {
		return errorAt50983A0
	}
	compressLen := int(data[0])
	if len(data) < 1+compressLen {
		return errorAt50983A0
	}
	data = data[1+compressLen:]
	if len(data) < 2 {
		return errorAt50983B0
	}
	extLen := int(binary.BigEndian.Uint16(data))
	data = data[2:]
	if extLen != len(data) {
		return errorAt50983B0
	}
	for len(data) > 0 {
		if len(data) < 4 {
			return errorAt50983B0
		}
		kind := binary.BigEndian.Uint16(data)
		n := int(binary.BigEndian.Uint16(data[2:]))
		data = data[4:]
		if n > len(data) {
			return errorAt50983B0
		}
		if kind == 0 {
			if n < 2 {
				return errorAt50983B0
			}
			names := data[:n]
			listLen := int(binary.BigEndian.Uint16(names))
			names = names[2:]
			if listLen != len(names) {
				return errorAt50983B0
			}
			for len(names) > 0 {
				if len(names) < 3 {
					return errorAt50983B0
				}
				nameType := names[0]
				nameLen := int(binary.BigEndian.Uint16(names[1:]))
				names = names[3:]
				if nameLen > len(names) {
					return errorAt50983B0
				}
				if nameType == 0 {
					host := string(names[:nameLen])
					if strings.HasSuffix(host, ".") {
						return errorAt50983B0
					}
					result.dunzUGo = host
					return nil
				}
				names = names[nameLen:]
			}
		}
		data = data[n:]
	}
	return errorAt50983C0
}
