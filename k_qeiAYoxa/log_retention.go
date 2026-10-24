package k_qeiAYoxa

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/dig"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
)

// 两个原 RTTI 均为零大小结构。
type AaBrngvix_ struct{}
type W2JC8JgOW struct{}

func init() {
	// 原 init.0@0x1ee49a0 和 init.2@0x1ee6460，组名来自 0x2123d40。
	cNT24rG4F.Bnm1qf(rjabCj7Pjm, dig.Group("JOBS"))
	cNT24rG4F.Bnm1qf(nrTbX9QD, dig.Group("JOBS"))
}

func rjabCj7Pjm() YNfyrQAMx      { return &AaBrngvix_{} } // 0x1ee4a00。
func nrTbX9QD() YNfyrQAMx        { return &W2JC8JgOW{} }  // 0x1ee64c0。
func (*AaBrngvix_) Cron() string { return "0 1 * * * ?" } // 0x1ee4a20。
func (*W2JC8JgOW) Cron() string  { return "0 0 * * * ?" } // 0x1ee64e0。

// Run 对应 0x1ee4a80，不检查 access_log_enable，检查 mode、目录和保留小时数。
func (*AaBrngvix_) Run() {
	mode, dir, hours := cNT24rG4F.AccessLogRetentionConfiguration()
	if mode != "local" || dir == "" || hours <= 0 {
		return
	}
	if _, err := os.Stat(dir); err != nil {
		return
	}
	now := time.Now()
	currentHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.Local)
	remove := []string{}
	err := fs.WalkDir(os.DirFS(dir), ".", retentionWalker(currentHour, hours, time.Hour, cNT24rG4F.CpXlPLFmKQcD, &remove))
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(4, "check access log dir %v failed: %v", dir, err)
	}
	for _, name := range remove {
		if err := os.Remove(filepath.Join(dir, name)); err != nil {
			muc279Mpa6S4.V8GN5Stk(4, "remove access log file <%v> failed: %v", name, err)
		} else {
			muc279Mpa6S4.V8GN5Stk(0, "remove access log file: %v", name)
		}
	}
}

// Run 对应 0x1ee6540：先切换当日日志，再执行清理。
func (*W2JC8JgOW) Run() {
	cNT24rG4F.W2uyDtOBN2()
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	dir, days := cNT24rG4F.LogRetentionConfiguration()
	if days <= 0 || dir == "" {
		return
	}
	remove := []string{}
	err := fs.WalkDir(os.DirFS(dir), ".", retentionWalker(today, days, 24*time.Hour, cNT24rG4F.YbaVop37K2p, &remove))
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(4, "check log file dir %v failed: %v", dir, err)
	}
	for _, name := range remove {
		if err := os.Remove(filepath.Join(dir, name)); err != nil {
			muc279Mpa6S4.V8GN5Stk(4, "remove log file <%v> failed: %v", name, err)
		} else {
			muc279Mpa6S4.V8GN5Stk(0, "remove log file: %v", name)
		}
	}
}

// retentionWalker 合并原闭包 0x1eea3a0、0x1eea840 的相同流程（分析命名）。
// 使用相对路径整体匹配文件名；小时与天数均以整数除法截断，严格大于保留值才删除。
func retentionWalker(now time.Time, keep int, unit time.Duration, parse func(string) (time.Time, bool), remove *[]string) fs.WalkDirFunc {
	return func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		date, ok := parse(path)
		if ok && int64(now.Sub(date)/unit) > int64(keep) {
			*remove = append(*remove, path)
		}
		return nil
	}
}
