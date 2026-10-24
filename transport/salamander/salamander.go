package salamander

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
	"math/rand"
	"sync"
	"time"
)

type PSIAOZXXoq interface {
	Deobfuscate([]byte, []byte) int
	Obfuscate([]byte, []byte) int
}

// DOqsrH8iC 的 40 字节布局来自原类型 0x365f980。
type DOqsrH8iC struct {
	EDkKnfGR8w []byte
	FY8Ki3jl   *rand.Rand
	pO1dNfho_  sync.Mutex
}

// 错误格式由 0x1c2c292 的原字符串解码器恢复；最小长度常量为 4。
var shortPasswordAt5098D90 = fmt.Errorf("PSK must be at least %d bytes", 4)

func Ea9POnRF(key []byte) (*DOqsrH8iC, error) { // 0x1c2c9c0
	if len(key) < 4 {
		return nil, shortPasswordAt5098D90
	}
	return &DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(rand.NewSource(time.Now().UnixNano()))}, nil
}

func (o *DOqsrH8iC) Obfuscate(input, output []byte) int { // 0x1c2cbc0
	if len(output) < len(input)+8 {
		return 0
	}
	o.pO1dNfho_.Lock()
	o.FY8Ki3jl.Read(output[:8])
	o.pO1dNfho_.Unlock()
	mask := blake2b.Sum256(append(o.EDkKnfGR8w, output[:8]...))
	for i, b := range input {
		output[i+8] = b ^ mask[i%32]
	}
	return len(input) + 8
}

func (o *DOqsrH8iC) Deobfuscate(input, output []byte) int { // 0x1c2ce60
	n := len(input) - 8
	if n <= 0 || len(output) < n {
		return 0
	}
	mask := blake2b.Sum256(append(o.EDkKnfGR8w, input[:8]...))
	for i, b := range input[8:] {
		output[i] = b ^ mask[i%32]
	}
	return n
}
