package xiaov2boardnode

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	CA "recovered.local/soga/util/rtx"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	Web "recovered.local/soga/panel/nodeconvert"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	O "recovered.local/soga/db"
	Q "recovered.local/soga/panel/registry"
	"strconv"
	"strings"
)

// These entries restore the original package's registered constructors and methods.
func gBG7jqqBxj(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &MrOlpocy{pBXis_w: c} }
func init() {
	Q.PanelFactoriesAt5095E50["whmcs-v2ray-db"] = gBG7jqqBxj
	Q.PanelFactoriesAt5095E50["whmcs-v2ray-webapi"] = Web.D_5dOciQK_A
}
func (p *MrOlpocy) GetAPIConfig() *C.FCEAfQJwDA7e { return p.pBXis_w }
func (p *MrOlpocy) GetUsers() (L.Sxcf3wu3L, error) {
	rows, err := O.DB().QueryContext(context.Background(), "select id, u, d, transfer_enable, uuid from user where enable = 1 and u + d < transfer_enable")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := QgFZpfpg8DH{}
	for rows.Next() {
		u := new(NzbHXtDtzF)
		var secret sql.NullString
		if err := rows.Scan(&u.LMCS6ltENGT, &u.OpTNDA2LU, &u.ObCFHYm9Op, &u.GfoBcVVws1j_, &secret); err != nil {
			return nil, err
		}
		if !secret.Valid {
			Log.V8GN5Stk(slog.LevelWarn, "user id %v: uuid is null", u.LMCS6ltENGT)
			continue
		}
		u.DhIHEmq = secret.String
		list = append(list, u)
	}
	// Original intentionally does not inspect Rows.Err after Next returns false.
	return list.GetUsers(), nil
}
func (p *MrOlpocy) SubmitData(in *Q.FGmNba4MhioY) error {
	if in == nil {
		return nil
	}
	return p.SubmitTraffic(in.HTgPyE3VytZd)
}

// Original batches5000 records and continues after failed Execs, preserving duplicates.
func (p *MrOlpocy) SubmitTraffic(in L.CWfcUv) error {
	if len(in) == 0 {
		return nil
	}
	db := O.DB()
	rate := C.WHMCSSettings().IlqtsyWu
	now := uint32(CA.C5Jff5vs0())
	var errs []error
	for start := 0; start < len(in); start += 5000 {
		end := start + 5000
		if end > len(in) {
			end = len(in)
		}
		ids, ups, downs, times := []string{}, []string{}, []string{}, []string{}
		for _, v := range in[start:end] {
			ids = append(ids, strconv.FormatInt(v.Hy7Q97RROK, 10))
			ups = append(ups, fmt.Sprintf("when %d then u + %d", v.Hy7Q97RROK, int64(float64(v.XJkGTQf_qkAt)*rate)))
			downs = append(downs, fmt.Sprintf("when %d then d + %d", v.Hy7Q97RROK, int64(float64(v.I6LXObODav)*rate)))
			times = append(times, fmt.Sprintf("when %d then %d", v.Hy7Q97RROK, now))
		}
		statement := "update user set u = case id\n" + strings.Join(ups, "\n") + "\nend,\nd = case id\n" + strings.Join(downs, "\n") + "\nend,\nt = case id\n" + strings.Join(times, "\n") + "\nend\nwhere id in(" + strings.Join(ids, ",") + ")"
		if _, err := db.ExecContext(context.Background(), statement); err != nil {
			errs = append(errs, err)
		}
	}
	return H.T2i1ItMaa(errs...)
}
func (p *MrOlpocy) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	node, e := p.GetNode()
	if e != nil {
		return nil, e
	}
	users, e := p.GetUsers()
	if e != nil {
		return nil, e
	}
	for _, u := range users {
		if v, ok := u.JAfXWwzYrW.(*L.Yn8BJoL); ok {
			v.FvQ2ltKOM7 = node.SwxJNgPpg.VUwf1n9g
		}
	}
	return &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: users}, nil
}
