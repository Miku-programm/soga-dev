package VasCJz

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestRequestPayloadChunking(t *testing.T) {
	for _, key := range []string{"", "X-payload"} {
		c := &R0J19auWof{F0XSwqNYoaP: map[string]string{"User-Agent": ""}, HFEkUe4a9Sas: key, JvaHmm: &MUVoWl0{64, 64}}
		data := bytes.Repeat([]byte{0xff, 0x00, 0xfb, 0xfe, 0x01}, 41)
		encoded := base64.RawURLEncoding.EncodeToString(data)
		h := c.GetRequestHeaderWithPayload(data)
		cookies := c.GetRequestCookiesWithPayload(data)
		for i, offset := 0, 0; offset < len(encoded); i, offset = i+1, offset+64 {
			want := encoded[offset:min(offset+64, len(encoded))]
			if got := h.Get(fmt.Sprintf("%s-%d", key, i)); got != want {
				t.Fatalf("header chunk %d: %q", i, got)
			}
			expected := &http.Cookie{Name: fmt.Sprintf("%s_%d", key, i), Value: want}
			if !reflect.DeepEqual(cookies[i], expected) {
				t.Fatalf("cookie chunk %d: %#v", i, cookies[i])
			}
		}
		if got := c.GetRequestCookiesWithPayload(nil); got == nil || len(got) != 0 {
			t.Fatalf("empty cookies: %#v", got)
		}
		if got := c.GetRequestHeaderWithPayload(nil); len(got) != 1 || !reflect.DeepEqual(got["User-Agent"], []string{""}) {
			t.Fatalf("empty headers: %#v", got)
		}
	}
	c := &R0J19auWof{F0XSwqNYoaP: map[string]string{"X-Ab": "first", "x-ab": "second", "User-Agent": "custom"}}
	if h := c.GetRequestHeader(); len(h.Values("X-Ab")) != 2 {
		t.Fatalf("canonical collision lost: %#v", h)
	}
}

// 在真实 HTTP 连接上传分片，接到已恢复原服务端的 packet-up 队列。
func TestGeneratedPayloadThroughNativeXHTTP(t *testing.T) {
	for _, place := range []string{"header", "cookie"} {
		t.Run(place, func(t *testing.T) {
			cfg := &R0J19auWof{HoBz4V: "/generated", MfWFUwd1: "packet-up", Aonc19Hp33E: place, HFEkUe4a9Sas: "P", JvaHmm: &MUVoWl0{64, 97}}
			l, base := xhttpFixture(t, cfg)
			payload := bytes.Repeat([]byte("\x00\xffnative-generated-payload"), 41)
			r := paddedRequest(t, "POST", base+"/generated/session/0", nil)
			if place == "header" {
				h := cfg.GetRequestHeaderWithPayload(payload)
				h.Set("Referer", r.Header.Get("Referer"))
				r.Header = h
			} else {
				for _, cookie := range cfg.GetRequestCookiesWithPayload(payload) {
					r.AddCookie(cookie)
				}
			}
			client := &http.Client{Timeout: 3 * time.Second}
			response, err := client.Do(r)
			if err != nil {
				t.Fatal(err)
			}
			io.Copy(io.Discard, response.Body)
			response.Body.Close()
			if response.StatusCode != 200 {
				t.Fatal(response.Status)
			}
			response, err = client.Do(paddedRequest(t, "GET", base+"/generated/session", nil))
			if err != nil {
				t.Fatal(err)
			}
			defer response.Body.Close()
			conn := xhttpAccept(t, l)
			defer conn.Close()
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			got := make([]byte, len(payload))
			if _, err = io.ReadFull(conn, got); err != nil || !bytes.Equal(got, payload) {
				t.Fatalf("payload %d: %v", len(got), err)
			}
			if strings.Contains(base64.RawURLEncoding.EncodeToString(payload), "=") {
				t.Fatal("padded encoding")
			}
		})
	}
}
