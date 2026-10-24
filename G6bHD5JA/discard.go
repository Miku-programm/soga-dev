package G6bHD5JA

import (
	"io"
	"recovered.local/soga/D_zCag1qTV"
)

// KlTAqrF42mR @0x15b34e0 使用原共享池的1024字节缓冲。
func KlTAqrF42mR(r io.Reader, size int) error {
	buffer := D_zCag1qTV.PoolAt50961D8.Get(1024)
	defer D_zCag1qTV.BE1pnyFd(buffer)
	return NFCqwGJM4(r, buffer, size)
}

// NFCqwGJM4 @0x15b3600：size<=0时持续读取直到错误；EOF也原样返回。
// size>0且最后一次Read同时返回足量数据和error时，仍优先返回error。
func NFCqwGJM4(r io.Reader, buffer []byte, size int) error {
	if size <= 0 {
		for {
			if _, err := r.Read(buffer); err != nil {
				return err
			}
		}
	}
	for size > 0 {
		target := buffer
		if len(target) > size {
			target = target[:size]
		}
		n, err := r.Read(target)
		size -= n
		if err != nil {
			return err
		}
	}
	return nil
}
