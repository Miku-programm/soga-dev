package rtx

import (
	"context"
	"errors"
	"sync"
	"syscall"

	"recovered.local/soga/util/errutil"
	"recovered.local/soga/util/logx"
)

// RTTI noIJiSXMP.CCVciOvWaCEr，amd64 大小 56 字节。
type CCVciOvWaCEr struct {
	zWBFc79Q0G  sync.RWMutex // 偏移 0。
	rVSeQKqr    []byte       // 偏移 24。
	mnsdOl0CZUz int          // 偏移 48。
}

// LWzrRP（0xb45860）创建文件映射并启动原有文件变化监视。
func LWzrRP(ctx context.Context, path string) (*CCVciOvWaCEr, error) {
	// 原构造器未把文件描述符预设为 -1；首次 g9iBYoc 会先 Close 零值。
	mapped := new(CCVciOvWaCEr)
	if err := mapped.g9iBYoc(path); err != nil {
		return nil, err
	}
	mapped.hjJaaqsj(ctx, path)
	return mapped, nil
}

// hjJaaqsj（0xb45e20）及闭包（0xb4f220）；重载失败也继续监视。
func (mapped *CCVciOvWaCEr) hjJaaqsj(ctx context.Context, path string) {
	err := fileWatcherAt5029240.WatchChanged(ctx, func() bool {
		if err := mapped.g9iBYoc(path); err != nil {
			logx.V8GN5Stk(8, "refresh mmap data %s failed: %v", path, err)
		} else {
			logx.V8GN5Stk(0, "mmap data %s refreshed", path)
		}
		return true
	}, path)
	if err != nil {
		logx.V8GN5Stk(8, "monitor mmap data %s failed: %v", path, err)
	}
}

// Data（0xb45940）直接返回映射切片，没有复制底层数据。
func (mapped *CCVciOvWaCEr) Data() ([]byte, error) {
	mapped.zWBFc79Q0G.RLock()
	defer mapped.zWBFc79Q0G.RUnlock()
	if mapped.rVSeQKqr != nil {
		return mapped.rVSeQKqr, nil
	}
	return nil, errors.New("data is not initialized")
}

// g9iBYoc（0xb45b40）关闭旧映射后重新打开；重新打开失败不会保留旧数据。
func (mapped *CCVciOvWaCEr) g9iBYoc(path string) error {
	mapped.zWBFc79Q0G.Lock()
	defer mapped.zWBFc79Q0G.Unlock()
	_ = mapped.Close()
	fd, err := syscall.Open(path, syscall.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	var stat syscall.Stat_t
	if err := syscall.Fstat(fd, &stat); err != nil {
		_ = syscall.Close(fd)
		return err
	}
	data, err := syscall.Mmap(fd, 0, int(stat.Size), syscall.PROT_READ, syscall.MAP_PRIVATE)
	if err != nil {
		_ = syscall.Close(fd)
		return err
	}
	mapped.rVSeQKqr = data
	mapped.mnsdOl0CZUz = fd
	return nil
}

// Close（0xb46000）自身不取锁，因为重载方法在持有写锁时调用它。
func (mapped *CCVciOvWaCEr) Close() error {
	var unmapErr, closeErr error
	if mapped.rVSeQKqr != nil {
		unmapErr = syscall.Munmap(mapped.rVSeQKqr)
		mapped.rVSeQKqr = nil
	}
	if mapped.mnsdOl0CZUz >= 0 {
		closeErr = syscall.Close(mapped.mnsdOl0CZUz)
		mapped.mnsdOl0CZUz = -1
	}
	return errutil.T2i1ItMaa(unmapErr, closeErr)
}
