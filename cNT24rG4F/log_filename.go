package cNT24rG4F

import "time"

// CpXlPLFmKQcD（0x131d5e0）提取访问日志文件名中的本地小时。
func CpXlPLFmKQcD(name string) (time.Time, bool) {
	match := accessLogPatternAt5095A98.FindStringSubmatch(name)
	if len(match) < 2 {
		return time.Time{}, false
	}
	parsed, err := time.ParseInLocation("2006-01-02 15", match[1], time.Local)
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

// YbaVop37K2p（0x1327540）提取普通日志文件名中的本地日期。
// 原正则未锚定整个名称，因此前后附加文本也可能匹配。
func YbaVop37K2p(name string) (time.Time, bool) {
	match := logPatternAt5095AB8.FindStringSubmatch(name)
	if len(match) < 3 {
		return time.Time{}, false
	}
	parsed, err := time.ParseInLocation("2006-01-02", match[2], time.Local)
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}
