package errutil

import "errors"

// 原 init@0xb25b00 在0x5097800构造的共享error；地址名由恢复工程指定。
// 文本来自原decoder，保持共享对象身份，不替换为context.Canceled。
var ErrorAt5097800 = errors.New("context done")

// 以下两项同样来自 init@0xb25b00 的原字符串解码；共享身份供各协议分支比较。
var ErrorAt5097810 = errors.New("already free")
var ErrorAt5097820 = errors.New("not enough information for making a decision")
