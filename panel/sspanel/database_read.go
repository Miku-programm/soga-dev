package sspanel

import (
	"database/sql"
	"fmt"
	A "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	U "recovered.local/soga/panel/sspaneltypes"
	C "recovered.local/soga/config"
	F "recovered.local/soga/panel/ssrnames"
	M "recovered.local/soga/util/logx"
	O "recovered.local/soga/db"
	X "recovered.local/soga/audit/matcher"
)

// 原 RTTI 8 字节；两个包全局字节保持原跨实例列探测状态。
type N3owFvm struct{ TMfNUn *C.FCEAfQJwDA7e }

var columnsCheckedAt50C6CC6 bool
var uuidColumnAt50C6CC7 bool

// 实际生产读取仍走原 oXGarxahaig once 初始化和全局 DB。
var database = O.DB

func (p *N3owFvm) GetAPIConfig() *C.FCEAfQJwDA7e { return p.TMfNUn }
func (p *N3owFvm) GetNodeResult() (*U.H0izB2agvG, error) {
	id := p.TMfNUn.Yf1KqSqJpNI7
	db := database()
	var server sql.NullString
	var speed float64
	var mu int
	if err := db.QueryRow("select server, node_speedlimit, mu_only from ss_node where id = ?", id).Scan(&server, &speed, &mu); err != nil {
		return nil, err
	}
	return &U.H0izB2agvG{L8xOe2JrpGDq: U.NNrxxXeBt{OFIQtbK5h: speed, HIgDb6THe1A9: server.String, A0zacuR6Nzb: mu}}, nil
}
func (p *N3owFvm) GetUsersResult() (*U.O8T6nD3uPzb, error) {
	id := p.TMfNUn.Yf1KqSqJpNI7
	db := database()
	now := uint32(A.C5Jff5vs0())
	go func() {
		if _, err := db.Exec("update ss_node set node_heartbeat = ? where id = ?", now, id); err != nil {
			M.TRTbrquD(4, "node heart err:", err)
		}
	}()
	var limit, used int64
	var group, class int
	if err := db.QueryRow("select node_bandwidth_limit, node_bandwidth, node_group, node_class from ss_node where id = ?", id).Scan(&limit, &used, &group, &class); err != nil {
		return nil, fmt.Errorf("node id: %v not exist: %v", id, err)
	}
	if limit > 0 && used >= limit {
		M.TRTbrquD(0, "node bandwidthLimit")
		return &U.O8T6nD3uPzb{L8xOe2JrpGDq: make([]*U.QhOI5ti, 0)}, nil
	}
	if !columnsCheckedAt50C6CC6 {
		rows, err := db.Query("select `COLUMN_NAME` from information_schema.COLUMNS where table_name = 'user' and table_schema = ?", p.TMfNUn.XU27D_AD)
		if err != nil {
			return nil, err
		}
		var name sql.NullString
		for rows.Next() {
			if err := rows.Scan(&name); err != nil {
				return nil, err
			}
			if name.String == "uuid" {
				uuidColumnAt50C6CC7 = true
				break
			}
		}
		_ = rows.Close()
		columnsCheckedAt50C6CC6 = true
	}
	fields := "id, passwd, u, d, transfer_enable, node_speedlimit, node_connector, port, method, protocol, protocol_param, obfs, obfs_param, is_multi_user"
	query := fmt.Sprintf("select %s from user", fields)
	if uuidColumnAt50C6CC7 {
		query = fmt.Sprintf("select %s, uuid from user", fields)
	}
	query += " where enable = 1 and expire_in > ? and u + d < transfer_enable and "
	var rows *sql.Rows
	var err error
	if group != 0 {
		query += "((class >= ? and node_group >= ?) or is_admin = 1)"
		rows, err = db.Query(query, A.XMMIhT97lTE7(), class, group)
	} else {
		query += "(class >= ? or is_admin = 1)"
		rows, err = db.Query(query, A.XMMIhT97lTE7(), class)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]*U.QhOI5ti, 0)
	for rows.Next() {
		var id, upload, download, quota, connector sql.NullInt64
		var password, method, protocol, protocolParam, obfs, obfsParam, uuid sql.NullString
		var speed sql.NullFloat64
		var port, mu sql.NullInt32
		dest := []any{&id, &password, &upload, &download, &quota, &speed, &connector, &port, &method, &protocol, &protocolParam, &obfs, &obfsParam, &mu}
		if uuidColumnAt50C6CC7 {
			dest = append(dest, &uuid)
		}
		if err = rows.Scan(dest...); err != nil {
			return nil, err
		}
		out = append(out, &U.QhOI5ti{BD16miI0: id.Int64, ESm7aAEddE: password.String, DpaCpPFc: upload.Int64, VsaEfGizz2x: download.Int64, DhjDXGJ1F2i: quota.Int64, TnjwdH: speed.Float64, Zsascdnac: int(connector.Int64), TdPoGXzK: int(port.Int32), BAyDRja4o: F.ZvAxOt1l7Ar(method.String), KzG7RsMW_4yx: F.Mu0vMIIs(protocol.String), FPf5h1940BAv: protocolParam.String, LEtQeEkSu: F.PMorzD7u(obfs.String), HX3uIcF: obfsParam.String, RiD2Bnl: L.LaTpsgYQ7(mu.Int32), Tiybx6: uuid.String})
	}
	// 原函数不读取 Rows.Err；保留已扫描部分结果。
	return &U.O8T6nD3uPzb{L8xOe2JrpGDq: out}, nil
}
func (p *N3owFvm) GetDetectRulesResult() (*U.Nq0XWDbiQ, error) {
	rows, err := database().Query("select id, regex, `type` from detect_list")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]*U.E3s7VeiKCij, 0)
	for rows.Next() {
		var id sql.NullInt64
		var pattern sql.NullString
		var kind sql.NullInt32
		if err = rows.Scan(&id, &pattern, &kind); err != nil {
			return nil, err
		}
		out = append(out, &U.E3s7VeiKCij{Uyj0tudIZN: int(id.Int64), XEmS9xCt: pattern.String, Bk26FTPOD: int(kind.Int32)})
	}
	return &U.Nq0XWDbiQ{L8xOe2JrpGDq: out}, nil
}
func (p *N3owFvm) GetUsers() (L.Sxcf3wu3L, error) {
	r, e := p.GetUsersResult()
	if e != nil {
		return nil, fmt.Errorf("get users failed: %v", e)
	}
	return r.GetUsers(p.TMfNUn.ZyAhkq5th8bf), nil
}
func (p *N3owFvm) GetDetectors() (X.Ij9r21E0UOt3, error) {
	r, e := p.GetDetectRulesResult()
	if e != nil {
		return nil, fmt.Errorf("get detect rules failed: %v", e)
	}
	return r.ToDetectors(), nil
}
