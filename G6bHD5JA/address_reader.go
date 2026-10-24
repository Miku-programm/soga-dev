package G6bHD5JA

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// TBUcsYQLu25q 对应 0x15a0260；失败不返回已经读入的部分字节。
func TBUcsYQLu25q(r io.Reader) (byte, error) {
	var b [1]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, err
	}
	return b[0], nil
}

// HoymNl_VBq 对应 0x15a0300，两个字节采用网络序。
func HoymNl_VBq(r io.Reader) (uint16, error) {
	var b [2]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b[:]), nil
}

// CHYbTjleKL 对应 0x15a07c0：先端口，后地址类型及地址字节。
// 原实现域名原样保留，即使其文本看起来像 IP 或长度为零。
func CHYbTjleKL(r io.Reader) (*AHU2MRmjzQu, error) {
	port, err := HoymNl_VBq(r)
	if err != nil {
		return nil, err
	}
	kind, err := TBUcsYQLu25q(r)
	if err != nil {
		return nil, err
	}
	var length int
	var isIP bool
	switch kind {
	case 1:
		length = 4
		isIP = true
	case 2:
		size, err := TBUcsYQLu25q(r)
		if err != nil {
			return nil, err
		}
		length = int(size)
	case 3:
		length = 16
		isIP = true
	default:
		return nil, fmt.Errorf("unknown addr type: %v", kind)
	}
	data := make([]byte, length)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, err
	}
	if isIP {
		ip := append(net.IP(nil), data...)
		return &AHU2MRmjzQu{dlcy_1: ip.String(), oKcMaQDwbH: []net.IP{ip}, pDnAz_: int(port), r_VdNaPA: true}, nil
	}
	return &AHU2MRmjzQu{dlcy_1: string(data), pDnAz_: int(port)}, nil
}
