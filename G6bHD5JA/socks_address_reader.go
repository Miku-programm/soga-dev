package G6bHD5JA

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// SzhW4tOzUZg 对应 0x15a03a0；原 SOCKS 格式为类型、地址、网络序端口。
// 域名分支不会把 IP 文本转换为 IP 地址；空域名同样保留。
func SzhW4tOzUZg(r io.Reader) (*AHU2MRmjzQu, error) {
	kind, err := TBUcsYQLu25q(r)
	if err != nil {
		return nil, err
	}
	var size int
	var isIP bool
	switch kind {
	case 1:
		size = 6
		isIP = true
	case 4:
		size = 18
		isIP = true
	case 3:
		n, err := TBUcsYQLu25q(r)
		if err != nil {
			return nil, err
		}
		size = int(n) + 2
	default:
		return nil, fmt.Errorf("unknown addr type: %v", kind)
	}
	data := make([]byte, size)
	if _, err = io.ReadFull(r, data); err != nil {
		return nil, err
	}
	port := int(binary.BigEndian.Uint16(data[size-2:]))
	if isIP {
		ip := append(net.IP(nil), data[:size-2]...)
		return &AHU2MRmjzQu{dlcy_1: ip.String(), oKcMaQDwbH: []net.IP{ip}, pDnAz_: port, r_VdNaPA: true}, nil
	}
	return &AHU2MRmjzQu{dlcy_1: string(data[:size-2]), pDnAz_: port}, nil
}
