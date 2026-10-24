// 根据 0x1274440 恢复；字段次序由原 RTTI 0x3651280 交叉核对。
package LFWTTXF7JR

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/cryptobyte"
	"strings"
)

// YUH4nAWc 读取私钥、配置各带两字节长度前缀的 ECH 密钥序列。
func YUH4nAWc(text string) ([]tls.EncryptedClientHelloKey, error) {
	var data []byte
	if block, _ := pem.Decode([]byte(text)); block != nil {
		data = block.Bytes
	} else {
		var err error
		data, err = base64.StdEncoding.DecodeString(strings.TrimSpace(text))
		if err != nil {
			return nil, fmt.Errorf("ech key 格式错误: %v", err)
		}
	}
	input := cryptobyte.String(data)
	var keys []tls.EncryptedClientHelloKey
	for len(input) != 0 {
		var privateKey, config cryptobyte.String
		if !input.ReadUint16LengthPrefixed(&privateKey) || !input.ReadUint16LengthPrefixed(&config) {
			return nil, errors.New("ech key 解析失败")
		}
		keys = append(keys, tls.EncryptedClientHelloKey{Config: config, PrivateKey: privateKey, SendAsRetry: true})
	}
	if len(keys) == 0 {
		return nil, errors.New("ech key 为空")
	}
	return keys, nil
}
