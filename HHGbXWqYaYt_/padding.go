package HHGbXWqYaYt_

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math/big"
	CA "recovered.local/soga/CAodbhZf"
	S "recovered.local/soga/sn9h4GXYmaaA"
	"strconv"
	"strings"
)

// FZBbzkJPWbrk 保留原始 56 字节字段布局（RTTI 0x36b6740）。
type FZBbzkJPWbrk struct {
	zUYCa31rxj  S.L4jWYITSyXGq
	KruzWnr9zxA []byte
	ZJzf4Os     uint32
	GtWiNfVs7j  string
}

// 原全局值；名称中的地址用于保留无符号变量与二进制的对应关系。
var DefaultAt5098340 CA.EahpvU_iseQ[*FZBbzkJPWbrk]

// 此默认字符串由原 init @0x188bac0 的解码器执行结果恢复。
var DefaultBytesAt509A190 = []byte("stop=8\n0=30-30\n1=100-400\n2=400-500,c,500-1000,c,500-1000,c,500-1000,c,500-1000\n3=9-9,500-1000\n4=500-1000\n5=500-1000\n6=500-1000\n7=500-1000")

func init() { TrBr1o4z(DefaultBytesAt509A190) }

func TrBr1o4z(b []byte) bool { // 0x188bba0
	if p := R_GY_FGx(b); p != nil {
		DefaultAt5098340.Store(p)
		return true
	}
	return false
}

func R_GY_FGx(b []byte) *FZBbzkJPWbrk { // 0x188bc20
	p := &FZBbzkJPWbrk{KruzWnr9zxA: b, GtWiNfVs7j: fmt.Sprintf("%x", md5.Sum(b))}
	m := S.DsbUyzhj9(b)
	if len(m) == 0 {
		return nil
	}
	stop, err := strconv.Atoi(m["stop"])
	if err != nil {
		return nil
	}
	p.ZJzf4Os = uint32(stop)
	p.zUYCa31rxj = m
	return p
}

func (p *FZBbzkJPWbrk) GenerateRecordPayloadSizes(record uint32) []int { // 0x188bde0
	rule, ok := p.zUYCa31rxj[strconv.FormatInt(int64(record), 10)]
	if !ok {
		return nil
	}
	var result []int
	for _, item := range strings.Split(rule, ",") {
		bounds := strings.Split(item, "-")
		if len(bounds) != 2 {
			if item == "c" {
				result = append(result, -1)
			}
			continue
		}
		a, err := strconv.ParseInt(bounds[0], 10, 64)
		if err != nil {
			continue
		}
		b, err := strconv.ParseInt(bounds[1], 10, 64)
		if err != nil {
			continue
		}
		lo, hi := min(a, b), max(a, b)
		if lo <= 0 {
			continue
		}
		if lo == hi {
			result = append(result, int(lo))
			continue
		}
		n, _ := rand.Int(rand.Reader, big.NewInt(hi-lo))
		result = append(result, int(n.Int64()+lo))
	}
	return result
}
