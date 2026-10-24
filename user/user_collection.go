package user

import "recovered.local/soga/util/genmap"

// ZpNFqUVZp 为原 RTTI 的 8 字节用户集合，嵌入字段名及实例化类型保持。
type ZpNFqUVZp struct {
	*genmap.GTFAWF[int64, PgaNd3014U]
}

// Add @0x181f140 按真实用户 ID 写入共享映射；相同 ID 覆盖原值。
func (m *ZpNFqUVZp) Add(u PgaNd3014U) { m.Set(u.GetId(), u) }
