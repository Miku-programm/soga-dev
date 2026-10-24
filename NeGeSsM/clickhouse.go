package NeGeSsM

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	C "recovered.local/soga/cNT24rG4F"
	L "recovered.local/soga/muc279Mpa6S4"
	"strconv"
	"strings"
	"time"
)

// z12Z6AW7z: 0x185c8a0。客户端超时 10 秒；配置在构造时快照。
func z12Z6AW7z() c5PkGn {
	c := C.E7kAu2lOVtXY()
	s := &aGuDZ8NsNT{jVe51FV: c, yw1h7z1aNajs: &http.Client{Timeout: 10 * time.Second}, q3mZAt: !c.EVFQh90cvJK}
	if c.ASvYULtJ7 != "" {
		query := "INSERT INTO " + c.FullTable() + " (time, network, node_id, user_id, user_ip, host, dest_ip, dest_port) FORMAT JSONEachRow"
		s.bV1jN0ia = c.ASvYULtJ7 + "/?query=" + url.QueryEscape(query)
	} else {
		L.V8GN5Stk(8, "access_log_clickhouse_addr is empty, access log will be dropped")
	}
	return s
}

// ajkzPuuQ: 0x185cbc0。时间是 Unix 秒；字段顺序与逐字节 JSON 转义保持原样。
func (s *aGuDZ8NsNT) ajkzPuuQ(b *strings.Builder, r *cUwNbDJ6opWy) {
	if s.bV1jN0ia == "" {
		return
	}
	b.Reset()
	b.WriteString(`{"time":`)
	b.WriteString(strconv.FormatInt(r.pv0C8wVKLJh.Unix(), 10))
	b.WriteString(`,"network":`)
	oRjSReB(b, r.uUbcC0nMK1Z9)
	b.WriteString(`,"node_id":`)
	b.WriteString(strconv.Itoa(r.iJYI5hR))
	b.WriteString(`,"user_id":`)
	b.WriteString(strconv.FormatInt(r.zBzHxRx4qP0, 10))
	b.WriteString(`,"user_ip":`)
	oRjSReB(b, r.c7a4q7l7VIYD)
	b.WriteString(`,"host":`)
	oRjSReB(b, r.mv3eQCXbV)
	b.WriteString(`,"dest_ip":`)
	oRjSReB(b, r.vz3eAFzt)
	b.WriteString(`,"dest_port":`)
	b.WriteString(strconv.Itoa(r.g8o8V93EQtw))
	b.WriteString("}\n")
	s.egrmDIRRx5.WriteString(b.String())
	s.peQJTFa++
	if s.peQJTFa >= s.jVe51FV.BaLySDum3K {
		s.eyAgzSwv7d()
	}
}

// eyAgzSwv7d: 0x185da40。初始化失败或提交失败仍丢弃当前批次，原代码没有重试队列。
func (s *aGuDZ8NsNT) eyAgzSwv7d() {
	if s.peQJTFa == 0 {
		return
	}
	n := s.peQJTFa
	data := s.egrmDIRRx5.Bytes()
	if s.init() {
		if err := s.kByXwtYgYQ(s.bV1jN0ia, data); err != nil {
			L.V8GN5Stk(4, "write %v access log records to clickhouse failed: %v", n, err)
		}
	}
	s.egrmDIRRx5.Reset()
	s.peQJTFa = 0
}

// t_DddcBjF7: 0x185dc00。
func (s *aGuDZ8NsNT) t_DddcBjF7() { s.eyAgzSwv7d(); s.yw1h7z1aNajs.CloseIdleConnections() }

// init: 0x185dc40。失败后至少间隔一分钟再次尝试建库/建表。
func (s *aGuDZ8NsNT) init() bool {
	if s.q3mZAt {
		return true
	}
	now := time.Now()
	if !s.ixZxkij70Gar.IsZero() && now.Sub(s.ixZxkij70Gar) < time.Minute {
		return false
	}
	s.ixZxkij70Gar = now
	if err := s.kByXwtYgYQ(s.jVe51FV.ASvYULtJ7+"/", []byte("CREATE DATABASE IF NOT EXISTS `"+s.jVe51FV.ZH4VO_H4JZ+"`")); err != nil {
		L.V8GN5Stk(4, "create clickhouse database <%v> failed: %v", s.jVe51FV.ZH4VO_H4JZ, err)
		return false
	}
	if err := s.kByXwtYgYQ(s.jVe51FV.ASvYULtJ7+"/", []byte(s.fJMhCf())); err != nil {
		L.V8GN5Stk(4, "create clickhouse table <%v> failed: %v", s.jVe51FV.FullTable(), err)
		return false
	}
	s.q3mZAt = true
	return true
}

// fJMhCf: 0x185e0a0。
func (s *aGuDZ8NsNT) fJMhCf() string {
	return "CREATE TABLE IF NOT EXISTS " + s.jVe51FV.FullTable() + ` (
	time DateTime,
	network LowCardinality(String),
	node_id Int32,
	user_id Int64,
	user_ip String,
	host String,
	dest_ip String,
	dest_port UInt16
) ENGINE = MergeTree PARTITION BY toYYYYMMDD(time) ORDER BY (time, user_id)`
}

// kByXwtYgYQ: 0x185e400；func1 0x1871d20 负责排空响应后关闭。
// 原代码只接受 HTTP 200，其它状态读取最多 512 字节用于错误，再排空剩余部分。
func (s *aGuDZ8NsNT) kByXwtYgYQ(endpoint string, data []byte) error {
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(data))
	if err != nil {
		return err
	}
	if s.jVe51FV.H0x74Y != "" {
		req.Header.Set("X-ClickHouse-User", s.jVe51FV.H0x74Y)
	}
	if s.jVe51FV.SNWlSe != "" {
		req.Header.Set("X-ClickHouse-Key", s.jVe51FV.SNWlSe)
	}
	resp, err := s.yw1h7z1aNajs.Do(req)
	if err != nil {
		return err
	}
	defer func() { _, _ = io.Copy(io.Discard, resp.Body); _ = resp.Body.Close() }()
	if resp.StatusCode == 200 {
		return nil
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
	return fmt.Errorf("clickhouse response %v: %v", resp.StatusCode, strings.TrimSpace(string(body)))
}
