package G7A7csgrs62p

// Len 对应原 int64/PgaNd3014U 实例 0x182b6e0；返回原子计数。
func (m *GTFAWF[K, V]) Len() int { return int(m.yaLIYjUJ.Load()) }

// Range 对应 0x182b700 及闭包 0x182b040。类型断言保留原映射行为。
func (m *GTFAWF[K, V]) Range(f func(K, V) bool) {
	m.r5_Hdcu3GxXC.Range(func(k, v any) bool { return f(k.(K), v.(V)) })
}
