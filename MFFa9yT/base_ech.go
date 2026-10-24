package MFFa9yT

import (
	"crypto/tls"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/muc279Mpa6S4"
)

// dINraX_uOx5T 对应0x1d42fe0，原tls.Config+368为EncryptedClientHelloKeys。
// 空文本或解码后的空列表不覆盖已有密钥。
func (b *JWN154qJ4) dINraX_uOx5T(config *tls.Config) error {
	if config == nil {
		return nil
	}
	text := b.wiy6P3.Load().LKMSis7zkX
	var keys []tls.EncryptedClientHelloKey
	var err error
	if text != "" {
		keys, err = LFWTTXF7JR.YUH4nAWc(text)
	}
	if err != nil {
		return err
	}
	if len(keys) != 0 {
		config.EncryptedClientHelloKeys = keys
		muc279Mpa6S4.V8GN5Stk(0, "ech enabled, keys: %v", len(keys))
	}
	return nil
}
