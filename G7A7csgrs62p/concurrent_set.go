package G7A7csgrs62p

// Fxs7SY 对应原 int64 实例 RTTI 0x35f7388，包含完整56字节并发映射。
type Fxs7SY[T comparable] struct{ icFkNd GTFAWF[T, struct{}] }

func (s *Fxs7SY[T]) Add(v T)      { s.icFkNd.Set(v, struct{}{}) }         //0x1eba6e0
func (s *Fxs7SY[T]) Delete(v T)   { s.icFkNd.Delete(v) }                  //0x1eba680
func (s *Fxs7SY[T]) Has(v T) bool { _, ok := s.icFkNd.Get(v); return ok } //0x1eba740
func (s *Fxs7SY[T]) Len() int     { return s.icFkNd.Len() }               //0x1eba5c0
// Clear 0x1eba5e0 原函数直接置零底层映射，非并发 Clear 操作。
func (s *Fxs7SY[T]) Clear()               { s.icFkNd = GTFAWF[T, struct{}]{} }
func (s *Fxs7SY[T]) Range(f func(T) bool) { s.icFkNd.Range(func(v T, _ struct{}) bool { return f(v) }) } //0x1eb8da0/1eb8ae0
func (s *Fxs7SY[T]) ToList() []T { //0x1eb8ce0/1eb8960
	out := make([]T, 0, s.Len())
	s.Range(func(v T) bool { out = append(out, v); return true })
	return out
}
