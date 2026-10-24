package jAP8Cn

import (
	"fmt"
	"io"
	"os"

	"recovered.local/soga/D_zCag1qTV"
)

// kBh3sTD66VJD（0xe9e740）把流写到临时文件后重命名。
// 原函数仅 defer 归还缓冲区，没有 Close 文件、fsync 或内容验证。
// 读取失败和空文件保留 .tmp，重命名失败才尝试删除它。
func kBh3sTD66VJD(reader io.Reader, path string) error {
	temporary := path + ".tmp"
	file, err := os.OpenFile(temporary, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open geosite file failed: %v", err)
	}
	buffer := D_zCag1qTV.PoolAt50961D8.Get(8192)
	defer func() { D_zCag1qTV.PoolAt50961D8.Put(buffer) }()
	written, err := io.CopyBuffer(file, reader, buffer)
	if err != nil {
		return err
	}
	if written == 0 {
		return fmt.Errorf("download geosite data failed: empty file")
	}
	if err := os.Rename(temporary, path); err != nil {
		_ = os.Remove(temporary)
		return fmt.Errorf("rename geosite file failed: %v", err)
	}
	return nil
}

// HaC1u4AhOL2o 对应原 haC1u4AhOL2o（0xe9ec60）；首字母导出以支持跨包调用。
func HaC1u4AhOL2o(reader io.Reader) (bool, error) {
	if err := kBh3sTD66VJD(reader, idPXJg8()); err != nil {
		return false, fmt.Errorf("download geosite data failed: %v", err)
	}
	return true, nil
}

// ZrUpfY5bpa 对应原 zrUpfY5bpa（0xe9edc0），保留共用写入器的原错误文本。
func ZrUpfY5bpa(reader io.Reader) (bool, error) {
	if err := kBh3sTD66VJD(reader, "/etc/soga/geoip.dat"); err != nil {
		return false, fmt.Errorf("download geoip data failed: %v", err)
	}
	return true, nil
}
