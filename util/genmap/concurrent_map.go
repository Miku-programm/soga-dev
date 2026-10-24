package genmap

// 以下五个方法根据 string/*uint8 实例的 0x159b8a0—0x159bc80 恢复。
// 键约束 comparable 用于表达原映射所需约束；原作者的约束声明无法还原。

func (m *GTFAWF[K, V]) Get(key K) (V, bool) {
	value, loaded := m.r5_Hdcu3GxXC.Load(key)
	if value == nil {
		var zero V
		return zero, false
	}
	return value.(V), loaded
}

func (m *GTFAWF[K, V]) Set(key K, value V) {
	_, loaded := m.r5_Hdcu3GxXC.Swap(key, value)
	if !loaded {
		m.yaLIYjUJ.Add(1)
	}
}

func (m *GTFAWF[K, V]) Delete(key K) {
	_, loaded := m.r5_Hdcu3GxXC.LoadAndDelete(key)
	if loaded {
		m.yaLIYjUJ.Add(-1)
	}
}

func (m *GTFAWF[K, V]) GetOrSet(key K, value V) (V, bool) {
	actual, loaded := m.r5_Hdcu3GxXC.LoadOrStore(key, value)
	if !loaded {
		m.yaLIYjUJ.Add(1)
	}
	return actual.(V), loaded
}

func (m *GTFAWF[K, V]) GetAndDelete(key K) (V, bool) {
	value, loaded := m.r5_Hdcu3GxXC.LoadAndDelete(key)
	if loaded {
		m.yaLIYjUJ.Add(-1)
	}
	if value == nil {
		var zero V
		return zero, loaded
	}
	return value.(V), loaded
}
