// 根据原样本 IDA 指令及解码字符串恢复；非作者文件原文。
package oXGarxahaig

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"recovered.local/soga/BaEWsPy"
	"recovered.local/soga/cNT24rG4F"
)

var onceAt50C7CF0 sync.Once
var databaseAt5096918 *sql.DB

// sTLjv2K_RUL 对应 0x170c0a0；闭包对应 0x170d580。
func sTLjv2K_RUL() {
	onceAt50C7CF0.Do(func() {
		host, port, name, user, password := cNT24rG4F.DatabaseConfiguration()
		ips, err := BaEWsPy.ResolveDefault(host)
		if err != nil {
			log.Fatal("can not resolve db host:", host)
		}
		// 原代码将完整地址切片交给 %v，未取首地址，也未额外 Ping。
		dsn := fmt.Sprintf("%s:%s@tcp(%v:%d)/%s", user, password, ips, port, name)
		databaseAt5096918, err = sql.Open("mysql", dsn)
		if err == nil {
			idleMinutes, idleConnections, openConnections := cNT24rG4F.DatabasePoolConfiguration()
			databaseAt5096918.SetConnMaxIdleTime(time.Duration(idleMinutes) * time.Minute)
			databaseAt5096918.SetMaxIdleConns(idleConnections)
			databaseAt5096918.SetMaxOpenConns(openConnections)
		}
	})
}

// DB 保持原调用点的“初始化后读取全局指针”顺序。
func DB() *sql.DB {
	sTLjv2K_RUL()
	return databaseAt5096918
}
