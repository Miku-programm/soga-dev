package cNT24rG4F

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// AVktbLizjJl @0x1321820：缓存命名空间使用面板类型与对接身份的MD5。
// webapi分支读取原密钥选择器H7eU0uGXU，不读取URL；未知api哈希空串。
func AVktbLizjJl() string {
	c := &nodeConfigurationAt50A2AE0
	var identity string
	switch c.NJSpG6 {
	case "db":
		identity = fmt.Sprintf("%s-%s-%d-%s", c.EtRWBSEhLQc, c.ApfXLZLmFXh, c.Mk5cDiQzaj, c.CsBHZVUxTO)
	case "webapi":
		identity = fmt.Sprintf("%s-%s", c.EtRWBSEhLQc, H7eU0uGXU())
	}
	sum := md5.Sum([]byte(identity))
	return hex.EncodeToString(sum[:])
}
