//go:build linux

// 根据原 IDA 0x153fac0、0x1540180、0x15401e0、0x15400e0、0x1544fe0 恢复。
package cert

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"time"
)

func Qr4n7sWM5di(nodeID int, certPEM, keyPEM string) (string, string, error) {
	if err := rtx.ZvmhWx(certPEM, keyPEM); err != nil {
		return "", "", fmt.Errorf("panel cert invalid: %v", err)
	}
	c08mZH44Qa(certPEM)
	certPath, keyPath := J6c4NB_R(nodeID)
	if err := wLQ7ZW(certPath, certPEM, 0644); err != nil {
		return "", "", err
	}
	if err := wLQ7ZW(keyPath, keyPEM, 0600); err != nil {
		return "", "", err
	}
	return certPath, keyPath, nil
}

func c08mZH44Qa(certPEM string) {
	if expires, err := rtx.GD_QvaWw(certPEM); err == nil {
		vW8f3PWiFGg0(expires)
	}
}

func vW8f3PWiFGg0(expires time.Time) {
	remain := time.Until(expires)
	if remain < 0 {
		logx.V8GN5Stk(slog.LevelWarn, "panel cert expired at %v, please renew it on the panel", expires)
	} else if remain < 14*24*time.Hour {
		logx.V8GN5Stk(slog.LevelWarn, "panel cert will expire in %v days (%v), please renew it on the panel", int(remain.Hours()/24), expires)
	}
}

// D18E0adGmOX 的原闭包为 0x154c140，首次检查发生于首次每日 tick。
func D18E0adGmOX(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				certPath, _ := config.ManualCertificateFiles()
				expires, err := sur_b39JM6l(certPath)
				if err != nil {
					logx.TRTbrquD(slog.LevelDebug, "check panel cert expiry:", err)
				} else {
					vW8f3PWiFGg0(*expires)
				}
			}
		}
	}()
}

// Ed4asPq5 的原闭包为 0x154c620；mtime/size 相同不检查内容。
func Ed4asPq5(ctx context.Context, certPath, keyPath string) {
	snapshot, err := eG4MicXvEEf(certPath, keyPath)
	if err != nil {
		logx.TRTbrquD(slog.LevelWarn, "watch custom cert: initial stat failed:", err)
		return
	}
	logx.V8GN5Stk(slog.LevelInfo, "start watching custom cert: %s, %s", certPath, keyPath)
	go func() {
		ticker := time.NewTicker(120 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				next, err := eG4MicXvEEf(certPath, keyPath)
				if err != nil {
					logx.TRTbrquD(slog.LevelDebug, "watch custom cert: stat failed:", err)
					continue
				}
				if next == snapshot {
					continue
				}
				if _, err = tls.LoadX509KeyPair(certPath, keyPath); err != nil {
					logx.TRTbrquD(slog.LevelWarn, "custom cert changed but invalid, skip reload:", err)
					continue
				}
				snapshot = next // 原版在发送信号前更新，失败后不会对同一快照重试。
				logx.TRTbrquD(slog.LevelInfo, "custom cert changed, reloading...")
				if err = rtx.Jx2ev8z(); err != nil {
					logx.TRTbrquD(slog.LevelWarn, "reload failed:", err)
				}
			}
		}
	}()
}
