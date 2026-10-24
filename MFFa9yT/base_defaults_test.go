package MFFa9yT

import "testing"

// 对照原 ASM 的返回方式和 base-default-strings.json 的实际解码结果。
func TestNativeBaseDefaultMethods(t *testing.T) {
	for _, b := range []*JWN154qJ4{nil, {}} {
		first, second := b.Start(), b.Start()
		if first == nil || first.Error() != "not implement error" || first == second {
			t.Fatalf("original Start error allocation: %v / %v", first, second)
		}
		for name, call := range map[string]func(){"Stop": b.Stop, "UpdateAPIData": func() { b.UpdateAPIData(nil) }} {
			t.Run(name, func(t *testing.T) {
				defer func() {
					v := recover()
					if text, ok := v.(string); !ok || text != "not implement error" {
						t.Fatalf("original panic value: %#v", v)
					}
				}()
				call()
			})
		}
	}
}
