// 根据原样本 RTTI、内联元数据与 IDA 指令恢复的集合类型。
// RTTI 名字前缀 dpV9yXd2B，方法包路径 G7A7csgrs62p；采用后者组织源码。
package G7A7csgrs62p

// Fw_lgKgam_jg 的 int64 实例描述符为 0x35f6938，仅有一个未导出 map 字段。
// comparable 表达恢复代码需要的键约束；原作者约束的源码写法不可由实例确定。
type Fw_lgKgam_jg[T comparable] struct {
	icFkNd map[T]struct{}
}

// PQQKui 的名称由 0x1326360 的内联调用树确认，调用点偏移 0xc4。
func PQQKui[T comparable]() *Fw_lgKgam_jg[T] {
	return &Fw_lgKgam_jg[T]{icFkNd: make(map[T]struct{})}
}

// Add 的 int64 实例位于 0x1347040；零值集合不自动初始化映射。
func (s *Fw_lgKgam_jg[T]) Add(value T) { s.icFkNd[value] = struct{}{} }

// Del 的 int64 实例位于 0x1346fe0。
func (s *Fw_lgKgam_jg[T]) Del(value T) { delete(s.icFkNd, value) }

// Has 的 int64 实例位于 0x13470a0。
func (s *Fw_lgKgam_jg[T]) Has(value T) bool {
	_, ok := s.icFkNd[value]
	return ok
}

// Len 的 int64 实例位于 0x1346fc0。
func (s *Fw_lgKgam_jg[T]) Len() int { return len(s.icFkNd) }

// Range 的 int64 实例位于 0x1346f20；原映射遍历不排序，回调返回假时停止。
func (s *Fw_lgKgam_jg[T]) Range(callback func(T) bool) {
	for value := range s.icFkNd {
		if !callback(value) {
			break
		}
	}
}
