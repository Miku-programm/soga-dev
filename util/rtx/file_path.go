// 根据原样本 IDA 函数体恢复；非作者源码原文。
package rtx

import (
	"os"
	"strings"
)

// VSCJLy2No 对应 0xb43b60。错误时只检查 os.ErrExist，不按 ErrNotExist 取反。
// os.init 将 io/fs 的 "file already exists" 错误复制到原全局 0x5096ef0。
func VSCJLy2No(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// EVZvjBCf 对应 0xb43be0。只移除一个结尾斜杠，再截到最后的斜杠之后。
// 该函数不会清理路径、识别反斜杠，也不会把没有目录的路径改成点号。
func EVZvjBCf(path string) string {
	if len(path) != 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path[:strings.LastIndex(path, "/")+1]
}
