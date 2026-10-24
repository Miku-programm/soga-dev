package rtx

import (
	"context"
	"time"

	"github.com/fsnotify/fsnotify"
	"recovered.local/soga/util/logx"
)

// RTTI noIJiSXMP.WaVcJUem 是零大小结构。选用的 fsnotify v1.7.0 与
// 原 inotify Watcher 的 64 字节、八个字段布局一致，未声称其为原精确版本。
type WaVcJUem struct{}

var fileWatcherAt5029240 = &WaVcJUem{}

// ms4MoNTpsin（0xb43c80）添加失败时等待事件、取消或三十秒后再次尝试。
func (monitor *WaVcJUem) ms4MoNTpsin(ctx context.Context, watcher *fsnotify.Watcher, path string) {
	for {
		err := watcher.AddWith(path)
		if err == nil {
			logx.V8GN5Stk(-4, "fs watcher: add %s success", path)
			return
		}
		logx.V8GN5Stk(-4, "fs watcher: add %s failed: %v, try next time", path, err)
		select {
		case <-watcher.Events:
		case <-ctx.Done():
			return
		case <-time.NewTimer(30 * time.Second).C:
		}
	}
}

// WatchChanged（0xb43f80；闭包 0xb4eca0）监视原路径，回调返回 false 时退出。
func (monitor *WaVcJUem) WatchChanged(ctx context.Context, callback func() bool, path string) error {
	watcher, err := fsnotify.NewBufferedWatcher(0)
	if err != nil {
		return err
	}
	go func() {
		defer watcher.Close()
		monitor.ms4MoNTpsin(ctx, watcher, path)
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write != 0 {
					logx.V8GN5Stk(0, "fs watcher: %v changed", path)
					if !callback() {
						return
					}
				}
				if event.Op&(fsnotify.Remove|fsnotify.Rename) != 0 {
					monitor.ms4MoNTpsin(ctx, watcher, path)
				}
			case err := <-watcher.Errors:
				logx.V8GN5Stk(0, "fs watcher: %v error: %v", path, err)
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
