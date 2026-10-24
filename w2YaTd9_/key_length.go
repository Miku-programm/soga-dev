package w2YaTd9_

import "fmt"

// KeyLen 对应 0x10a1300；未知算法错误文本由原 decoder 解出。
func (cipher C6q3Upu) KeyLen() (int, error) {
	switch cipher {
	case "aes-128-gcm", "2022-blake3-aes-128-gcm":
		return 16, nil
	case "aes-192-gcm":
		return 24, nil
	case "aes-256-gcm", "chacha20-ietf-poly1305", "2022-blake3-aes-256-gcm":
		return 32, nil
	default:
		return 0, fmt.Errorf("unknown ss cipher <%v> key len", cipher)
	}
}
