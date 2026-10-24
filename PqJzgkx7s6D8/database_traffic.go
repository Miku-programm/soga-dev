package PqJzgkx7s6D8

import (
	"database/sql"
	"fmt"
	A "recovered.local/soga/CAodbhZf"
	G "recovered.local/soga/G6bHD5JA"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	"strconv"
	"strings"
	"sync"
)

var trafficTableCheckedAt50C6CC8 bool
var trafficTableAvailableAt3787336 = true

// SubmitTraffic restores 0x170fa00 and its four closures 0x1735160..0x17368f4.
// Done precedes channel send; the log transaction finalizes after that send.
func (p *N3owFvm) SubmitTraffic(data L.CWfcUv) error {
	if len(data) == 0 {
		return nil
	}
	id := p.TMfNUn.Yf1KqSqJpNI7
	db := database()
	var value sql.NullFloat64
	if err := db.QueryRow("select traffic_rate from ss_node where id = ?", id).Scan(&value); err != nil {
		return err
	}
	rate := 1.0
	if value.Valid {
		rate = value.Float64
	}
	var wg sync.WaitGroup
	wg.Add(4)
	errs := make(chan error, 4)
	now := uint32(A.C5Jff5vs0())
	go func() {
		var err error
		defer func() { wg.Done(); errs <- err }()
		err = updateTrafficUsers(db, data, rate, now)
	}()
	go func() {
		var err error
		var tx *sql.Tx
		defer func() {
			wg.Done()
			errs <- err
			if tx != nil {
				if err != nil {
					_ = tx.Rollback()
				} else {
					_ = tx.Commit()
				}
			}
		}()
		if !trafficTableCheckedAt50C6CC8 {
			rows, queryErr := db.Query("SELECT table_name FROM information_schema.TABLES WHERE table_name='user_traffic_log' AND table_schema=?", p.TMfNUn.XU27D_AD)
			// The original emits an additional error here without assigning the
			// deferred error. This unusual five-send/four-receive path is preserved.
			if queryErr != nil {
				errs <- queryErr
				return
			}
			trafficTableCheckedAt50C6CC8 = true
			trafficTableAvailableAt3787336 = rows.Next()
			_ = rows.Close()
		}
		if !trafficTableAvailableAt3787336 {
			return
		}
		tx, err = db.Begin()
		if err != nil {
			return
		}
		for _, u := range data {
			traffic := G.LTaVAEf(uint64(float64(u.XJkGTQf_qkAt+u.I6LXObODav) * rate))
			_, err = tx.Exec("insert into user_traffic_log(user_id,u,d,node_id,rate,traffic,log_time) values(?,?,?,?,?,?,?)", u.Hy7Q97RROK, u.XJkGTQf_qkAt, u.I6LXObODav, id, rate, traffic, now)
			if err != nil {
				return
			}
		}
	}()
	go func() {
		var err error
		defer func() { wg.Done(); errs <- err }()
		var total int64
		for _, u := range data {
			total += u.XJkGTQf_qkAt + u.I6LXObODav
		}
		_, err = db.Exec("update ss_node set node_bandwidth = node_bandwidth + ? where id = ?", total, id)
	}()
	go func() {
		var err error
		defer func() { wg.Done(); errs <- err }()
		_, err = db.Exec("insert into ss_node_online_log(node_id,online_user,log_time) values(?,?,?)", id, len(data), now)
	}()
	wg.Wait()
	out := make([]error, 0)
	for i := 0; i < 4; i++ {
		if err := <-errs; err != nil {
			out = append(out, err)
		}
	}
	return H.T2i1ItMaa(out...)
}

// Analyst name for the original func1 body. Chunk size and SQL spelling are
// recovered from 0x1735c20; this is executed by the original first goroutine.
func updateTrafficUsers(db *sql.DB, data L.CWfcUv, rate float64, now uint32) error {
	for start := 0; start < len(data); start += 5000 {
		end := min(start+5000, len(data))
		var ids, up, down, stamp []string
		for _, u := range data[start:end] {
			ids = append(ids, strconv.FormatInt(u.Hy7Q97RROK, 10))
			up = append(up, fmt.Sprintf("when %d then u + %d", u.Hy7Q97RROK, int64(float64(u.XJkGTQf_qkAt)*rate)))
			down = append(down, fmt.Sprintf("when %d then d + %d", u.Hy7Q97RROK, int64(float64(u.I6LXObODav)*rate)))
			stamp = append(stamp, fmt.Sprintf("when %d then %d", u.Hy7Q97RROK, now))
		}
		query := "update user set u = case id\n" + strings.Join(up, "\n") + "\nend,\nd = case id\n" + strings.Join(down, "\n") + "\nend,\nt = case id\n" + strings.Join(stamp, "\n") + "\nend\nwhere id in(" + strings.Join(ids, ",") + ")"
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}
	return nil
}
