package VasCJz

import (
	"io"
	"net/http"
	"recovered.local/soga/xray/done"
	"sync"
)

type Eu1uXZE = http.ResponseWriter

// 原 RTTI 48 字节。Reader 和响应接口的提升方法由 Go 编译器生成。
type aaktBBMB struct {
	sync.Mutex
	*done.Instance
	io.Reader
	Eu1uXZE
}

// 0x125e280：只有写成功才刷新；原断言不是可选 Flusher 检查。
func (w *aaktBBMB) Write(p []byte) (int, error) {
	w.Lock()
	defer w.Unlock()
	if w.Done() {
		return 0, io.ErrClosedPipe
	}
	n, err := w.Eu1uXZE.Write(p)
	if err == nil {
		w.Eu1uXZE.(http.Flusher).Flush()
	}
	return n, err
}

// 0x125e4e0：仅发关闭信号，不能顺带关闭嵌入的请求 Body。
func (w *aaktBBMB) Close() error {
	w.Lock()
	defer w.Unlock()
	return w.Instance.Close()
}
