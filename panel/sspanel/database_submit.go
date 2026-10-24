package sspanel

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
	A "recovered.local/soga/util/rtx"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	M "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
)

func (p *N3owFvm) SubmitNodeInfo() error {
	db := database()
	now := uint32(A.C5Jff5vs0())
	id := p.TMfNUn.Yf1KqSqJpNI7
	avg, err := hostLoad()
	if err != nil {
		M.TRTbrquD(4, "get load avg err:", err)
		avg = &load.AvgStat{}
	}
	uptime, err := hostUptime()
	if err != nil {
		M.TRTbrquD(4, "Get uptime err:", err)
	}
	_, err = db.Exec("insert into ss_node_info(node_id, uptime, `load`, log_time)values(?,?,?,?)", id, uptime, fmt.Sprintf("%.2f %.2f %.2f", avg.Load1, avg.Load5, avg.Load15), now)
	return err
}
func (p *N3owFvm) SubmitAliveIP(data L.CQhZrJ4dK) error {
	if len(data) == 0 {
		return nil
	}
	db := database()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if tx != nil {
				_ = tx.Rollback()
			}
		} else {
			_ = tx.Commit()
		}
	}()
	id := p.TMfNUn.Yf1KqSqJpNI7
	now := uint32(A.C5Jff5vs0())
	for _, u := range data {
		for _, ip := range u.PNtn9KAV {
			_, err = tx.Exec("insert into alive_ip(nodeid,userid,ip,datetime) values(?, ?, ?, ?)", id, u.Vy9ybu, ip, now)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *N3owFvm) SubmitDetectLog(data L.QmiwDVjV00lZ) error {
	if len(data) == 0 {
		return nil
	}
	db := database()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// The original retries the status-column form after rollback, but still
	// returns the first error and deliberately discards the retry result.
	defer func() {
		if err != nil {
			if tx != nil {
				_ = tx.Rollback()
			}
			_ = p._rsq9Xmy9Vtv(data)
		} else {
			_ = tx.Commit()
		}
	}()
	id := p.TMfNUn.Yf1KqSqJpNI7
	now := uint32(A.C5Jff5vs0())
	for _, d := range data {
		_, err = tx.Exec("insert into detect_log(user_id,list_id,datetime,node_id) values(?,?,?,?)", d.GbzQUT_4f, d.CAgUUdrv9, now, id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *N3owFvm) _rsq9Xmy9Vtv(data L.QmiwDVjV00lZ) error {
	if len(data) == 0 {
		return nil
	}
	db := database()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if tx != nil {
				_ = tx.Rollback()
			}
		} else {
			_ = tx.Commit()
		}
	}()
	id := p.TMfNUn.Yf1KqSqJpNI7
	now := uint32(A.C5Jff5vs0())
	for _, d := range data {
		_, err = tx.Exec("insert into detect_log(user_id,list_id,datetime,node_id,status) values(?,?,?,?,?)", d.GbzQUT_4f, d.CAgUUdrv9, now, id, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *N3owFvm) SubmitData(data *Q.FGmNba4MhioY) error {
	return H.D5aDHlQDLp(p.SubmitNodeInfo, func() error { return p.SubmitTraffic(data.HTgPyE3VytZd) }, func() error { return p.SubmitDetectLog(data.Y0ys4ma4Tno4) }, func() error { return p.SubmitAliveIP(data.JNcHHwTZ7) })
}
