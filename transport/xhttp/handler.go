package xhttp

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	B "recovered.local/soga/buf/pool"
	Log "recovered.local/soga/util/logx"
	Q "recovered.local/soga/util/deadline2"
	XCrypto "recovered.local/soga/xray/crypto"
	"recovered.local/soga/xray/done"
	XErrors "recovered.local/soga/xray/errors"
	XHTTP "recovered.local/soga/xray/http"
	"recovered.local/soga/xray/internet"
	XNet "recovered.local/soga/xray/net"
	"strconv"
	"strings"
	"time"
)

// ServeHTTP 0x125a820。下方辅助函数名由分析者拆分，分支及执行顺序来自此单函数。
func (h *e8Q7HURSMg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := h.jgg_3z8gKTCK
	if h.taRd4Q3PP5 != "" && !internet.IsValidHTTPHost(r.Host, h.taRd4Q3PP5) {
		Log.TRTbrquD(0, "failed to validate host, request:", r.Host, ", config:", h.taRd4Q3PP5)
		w.WriteHeader(404)
		return
	}
	if !strings.HasPrefix(r.URL.Path, h.jD05lxT) {
		Log.TRTbrquD(0, "failed to validate path, request:", r.URL.Path, ", config:", h.jD05lxT)
		w.WriteHeader(404)
		return
	}
	c.WriteResponseHeader(w, r.Method, r.Header)
	bounds := c.GetNormalizedXPaddingBytes()
	padding := LaPHbx6B50uZ{TldsrRa: int(XCrypto.RandBetween(int64(bounds.AS9zG362), int64(bounds.HHTPwX0N)))}
	if c.PnRF9ifkXKD {
		padding.C5zKbVfj3g_ = AHn9pfr{AnBhAR: c.V0qykwP, AVHkw5m: c.H6PfvFysST, HxzvtHynCM: c.LjJbxC6vDL}
		padding.Yxjd2mW = AdpPM6(c.MLlGWkds)
	} else {
		padding.C5zKbVfj3g_ = AHn9pfr{AnBhAR: "header", HxzvtHynCM: "X-Padding"}
	}
	c.ApplyXPaddingToResponse(w, padding)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	value, location := c.ExtractXPaddingFromRequest(r, c.PnRF9ifkXKD)
	if !c.IsPaddingValid(value, bounds.AS9zG362, bounds.HHTPwX0N, AdpPM6(c.MLlGWkds)) {
		Log.TRTbrquD(0, "invalid padding ("+location+") length:", int32(len(value)))
		w.WriteHeader(400)
		return
	}
	id, sequence := c.ExtractMetaFromRequest(r, h.jD05lxT)
	if id == "" && c.MfWFUwd1 != "" && c.MfWFUwd1 != "auto" && c.MfWFUwd1 != "stream-one" && c.MfWFUwd1 != "stream-up" {
		Log.TRTbrquD(0, "stream-one mode is not allowed")
		w.WriteHeader(400)
		return
	}
	remote := h.requestRemoteAddr(r)
	var session *f30vZAC4M2Oh
	if id != "" {
		session = h.d39ITgdaw(id)
	}
	if (r.Method != "GET" || sequence != "") && id != "" {
		if sequence != "" {
			h.servePacketUpload(w, r, sequence, session)
			return
		}
		h.serveStreamUpload(w, r, session)
		return
	}
	if r.Method == "GET" {
		if id != "" {
			session.zYk4hazu.Close()
			defer h.wLTHxb.Delete(id)
		}
	} else if id != "" {
		Log.TRTbrquD(0, "unsupported method: ", r.Method)
		w.WriteHeader(405)
		return
	}
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Cache-Control", "no-store")
	if !c.A6NxppvA {
		w.Header().Set("Content-Type", "text/event-stream")
	}
	w.WriteHeader(200)
	w.(http.Flusher).Flush()
	stream := &aaktBBMB{Instance: done.New(), Reader: r.Body, Eu1uXZE: w}
	conn := &_0_9jH{scAGARMF0XsA: stream, xxabacttFOn: stream, n6EztFBWOT: remote, gSBX2NrxZ_8y: h._7JXu_SO6K}
	if id != "" {
		conn.xxabacttFOn = session.jg638CNSsPu
	}
	conn.pcYQQ0 = Q.NewReader(conn.xxabacttFOn)
	ctx, cancel := context.WithCancel(r.Context())
	conn.gvvJIQwQyA = cancel
	select {
	case h.vat5bP_nCH <- conn:
	default:
		conn.Close()
		return
	}
	select {
	case <-ctx.Done():
	case <-r.Context().Done():
	case <-stream.Wait():
	}
	conn.Close()
}

// 原函数先解析 RemoteAddr，失败时使用 0.0.0.0:0；XFF 仅首项为 IP 时覆盖。
func (h *e8Q7HURSMg) requestRemoteAddr(r *http.Request) net.Addr {
	var forwarded []XNet.Address
	if h.epe3GXboMP1 != nil && len(h.epe3GXboMP1.TrustedXForwardedFor) > 0 {
		for _, key := range h.epe3GXboMP1.TrustedXForwardedFor {
			if len(r.Header[http.CanonicalHeaderKey(key)]) != 0 {
				forwarded = XHTTP.ParseXForwardedFor(r.Header)
				break
			}
		}
	} else {
		forwarded = XHTTP.ParseXForwardedFor(r.Header)
	}
	tcp, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
	if err != nil {
		tcp = &net.TCPAddr{IP: net.IP{0, 0, 0, 0}}
	}
	var addr net.Addr = tcp
	if r.ProtoMajor == 3 {
		addr = &net.UDPAddr{IP: tcp.IP, Port: tcp.Port}
	}
	if len(forwarded) > 0 && forwarded[0].Family().IsIP() {
		addr = &net.TCPAddr{IP: forwarded[0].IP()}
	}
	return addr
}

func (h *e8Q7HURSMg) servePacketUpload(w http.ResponseWriter, r *http.Request, sequence string, s *f30vZAC4M2Oh) {
	c := h.jgg_3z8gKTCK
	if c.MfWFUwd1 != "" && c.MfWFUwd1 != "auto" && c.MfWFUwd1 != "packet-up" {
		Log.TRTbrquD(0, "packet-up mode is not allowed")
		w.WriteHeader(400)
		return
	}
	place := c.GetNormalizedUplinkDataPlacement()
	var headerPayload, cookiePayload []byte
	if place == "auto" || place == "header" {
		var chunks []string
		for i := 0; ; i++ {
			v := r.Header.Get(fmt.Sprintf("%s-%d", c.HFEkUe4a9Sas, i))
			if v == "" {
				break
			}
			chunks = append(chunks, v)
		}
		var err error
		headerPayload, err = base64.RawURLEncoding.DecodeString(strings.Join(chunks, ""))
		if err != nil {
			Log.TRTbrquD(0, "Invalid base64 in header's payload: ", err.Error())
			w.WriteHeader(400)
			return
		}
	}
	if place == "auto" || place == "cookie" {
		var chunks []string
		for i := 0; ; i++ {
			v, e := r.Cookie(fmt.Sprintf("%s_%d", c.HFEkUe4a9Sas, i))
			if e != nil {
				break
			}
			chunks = append(chunks, v.Value)
		}
		var err error
		cookiePayload, err = base64.RawURLEncoding.DecodeString(strings.Join(chunks, ""))
		if err != nil {
			Log.TRTbrquD(0, "Invalid base64 in cookies' payload: ", err.Error())
			w.WriteHeader(400)
			return
		}
	}
	var body, owned []byte
	defer func() {
		if owned != nil {
			B.PoolAt50961D8.Put(owned)
		}
		if body != nil {
			B.PoolAt50961D8.Put(body)
		}
	}()
	maxBytes := int(h.mZcQrY.fBC_MC1GOIw.GetNormalizedScMaxEachPostBytes().HHTPwX0N)
	tooLarge := func() {
		Log.TRTbrquD(0, "Too large upload. scMaxEachPostBytes is set to ", maxBytes, "but request size exceed it. Adjust scMaxEachPostBytes on the server to be at least as large as client.")
		w.WriteHeader(413)
	}
	if place == "auto" || place == "body" {
		if r.ContentLength > int64(maxBytes) {
			tooLarge()
			return
		}
		size := maxBytes + 1
		if r.ContentLength > 0 {
			size = int(r.ContentLength)
		}
		body = B.PoolAt50961D8.Get(size)
		// 0x125d32a—0x125d487：样本两支均 ReadFull；未知长度也不将 EOF 当成功。
		if _, err := io.ReadFull(r.Body, body); err != nil {
			XErrors.LogInfoInner(context.Background(), err, "failed to read body payload")
			w.WriteHeader(400)
			return
		}
	}
	noBody := len(body) == 0
	var payload []byte
	switch place {
	case "auto":
		total := len(headerPayload) + len(cookiePayload) + len(body)
		if total > 0 {
			owned = B.PoolAt50961D8.Get(total)
			n := copy(owned, headerPayload)
			n += copy(owned[n:], cookiePayload)
			copy(owned[n:], body)
			payload = owned
		}
		if body != nil {
			B.PoolAt50961D8.Put(body)
			body = nil
		}
	case "body":
		owned = body
		body = nil
		payload = owned
	case "cookie":
		payload = cookiePayload
	case "header":
		payload = headerPayload
	}
	if len(payload) > maxBytes {
		tooLarge()
		return
	}
	seq, err := strconv.ParseUint(sequence, 10, 64)
	if err != nil {
		XErrors.LogInfoInner(context.Background(), err, "failed to upload (ParseUint)")
		w.WriteHeader(500)
		return
	}
	if err := s.jg638CNSsPu.Push(UmYbKhCpX{F6dEJnE3: payload, Au22AAVkjZY: owned, JMZzwbXWi: seq}); err != nil {
		XErrors.LogInfoInner(context.Background(), err, "failed to upload (PushPayload)")
		w.WriteHeader(500)
		return
	}
	owned = nil
	if noBody {
		w.Header().Set("Cache-Control", "no-store")
	}
	w.WriteHeader(200)
}

func (h *e8Q7HURSMg) serveStreamUpload(w http.ResponseWriter, r *http.Request, s *f30vZAC4M2Oh) {
	c := h.jgg_3z8gKTCK
	if c.MfWFUwd1 != "" && c.MfWFUwd1 != "auto" && c.MfWFUwd1 != "stream-up" {
		Log.TRTbrquD(0, "stream-up mode is not allowed")
		w.WriteHeader(400)
		return
	}
	stream := &aaktBBMB{Instance: done.New(), Reader: r.Body, Eu1uXZE: w}
	if err := s.jg638CNSsPu.Push(UmYbKhCpX{ACMt1h1: stream}); err != nil {
		XErrors.LogInfoInner(context.Background(), err, "failed to upload (PushReader)")
		w.WriteHeader(409)
	} else {
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(200)
		secs := c.GetNormalizedScStreamUpServerSecs()
		if r.Header.Get("Referer") != "" && secs.HHTPwX0N > 0 {
			go func() { //0x1272f80：先写填充，再等待；不通过 ticker 或 context 提前退出。
				for {
					pad := c.GetNormalizedXPaddingBytes()
					if _, err := stream.Write(bytes.Repeat([]byte{'X'}, int(int32(XCrypto.RandBetween(int64(pad.AS9zG362), int64(pad.HHTPwX0N)))))); err != nil {
						return
					}
					time.Sleep(time.Duration(int32(XCrypto.RandBetween(int64(secs.AS9zG362), int64(secs.HHTPwX0N)))) * time.Second)
				}
			}()
		}
		select {
		case <-stream.Wait():
		case <-r.Context().Done():
		}
	}
	stream.Close()
}
