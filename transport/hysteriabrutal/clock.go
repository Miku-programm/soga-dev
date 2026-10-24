// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package hysteriabrutal

import "time"

// A Clock returns the current time
type KYybyfkN interface {
	Now() time.Time
}

// DefaultClock implements the Clock interface using the Go stdlib clock.
type MDeCv6qOB_J struct{}

var _ KYybyfkN = MDeCv6qOB_J{}

// Now gets the current time
func (MDeCv6qOB_J) Now() time.Time {
	return time.Now()
}
