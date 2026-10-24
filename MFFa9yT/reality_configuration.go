package MFFa9yT

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"recovered.local/soga/reality"
)

// fwKvAvIIe 由原0x1d67ca0恢复。依次校验short ID、private key、SNI、dest；
// 短ID复制进8字节键并右侧补零，空ID字符串有效；不额外施加文档约束。
func fwKvAvIIe(ctx context.Context, names []string, key, dest string, ids []string) (*reality.Config, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("reality short id 不能为空")
	}
	short := make(map[[8]byte]bool, len(ids))
	for _, s := range ids {
		b, e := hex.DecodeString(s)
		if e != nil {
			return nil, fmt.Errorf("reality short id 格式错误: %v: %v", s, e)
		}
		if len(b) > 8 {
			return nil, fmt.Errorf("reality short id 超过 8 字节: %v", s)
		}
		var id [8]byte
		copy(id[:], b)
		short[id] = true
	}
	private, e := base64.RawURLEncoding.DecodeString(key)
	if e != nil {
		return nil, fmt.Errorf("reality private key 格式错误: %v: %v", key, e)
	}
	if len(private) != 32 {
		return nil, fmt.Errorf("reality private key 解码后应为 32 字节，实际 %v 字节: %v", len(private), key)
	}
	if len(names) == 0 {
		return nil, fmt.Errorf("reality sni (serverNames) 不能为空")
	}
	serverNames := make(map[string]bool, len(names))
	for _, s := range names {
		serverNames[s] = true
	}
	if dest == "" {
		return nil, fmt.Errorf("reality dest 不能为空")
	}
	return &reality.Config{DialContext: gzasC2TjpYF(ctx, dest), Type: "tcp", Dest: dest, ServerNames: serverNames, PrivateKey: private, ShortIds: short, SessionTicketsDisabled: true}, nil
}
