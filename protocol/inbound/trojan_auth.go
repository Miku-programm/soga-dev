package inbound

import (
	"net"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	"time"
)

// s3mqrqd_ @0x1d8fc60：原单次非空Read需要完整56字节；短读直接回落。
func (v *EfaIFpp) s3mqrqd_(conn net.Conn) (*U.DWaKIV1T, []byte) {
	var auth [56]byte
	conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	var n int
	var err error
	for {
		n, err = conn.Read(auth[:])
		if n > 0 || err != nil {
			break
		}
	}
	conn.SetReadDeadline(time.Time{})
	if err != nil || n != 56 {
		Log.TRTbrquD(-4, "can not read trojan auth data:", auth[:n])
		return nil, auth[:n]
	}
	if user, ok := v.d9Cxrbh1g7.Load(auth); ok {
		return user.(*U.DWaKIV1T), nil
	}
	Log.TRTbrquD(-4, "auth failed:", auth)
	return nil, auth[:]
}
