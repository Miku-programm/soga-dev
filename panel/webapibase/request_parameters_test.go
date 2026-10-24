package webapibase

import (
	"net/http"
	"net/url"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/etagcache"
	"reflect"
	"testing"
)

func TestRequestParameters(t *testing.T) {
	cfg := &config.FCEAfQJwDA7e{W_ogGtF: "https://panel.test/", Bd08Li: "secret&=测试", Yf1KqSqJpNI7: -5}
	p := &JZA3kKksUe{EA_5pkz: cfg, eyq32s: &etagcache.NG0hCZyRHK{}}
	if p.GetAPIConfig() != cfg {
		t.Fatal("must return original pointer")
	}
	for _, tc := range []struct {
		protocol config.Mpn2JltM
		want     string
	}{{"ss", "shadowsocks"}, {"v2ray", "v2ray"}, {"trojan", "trojan"}, {"SS", ""}, {" ss ", ""}, {"vless", ""}, {"", ""}} {
		cfg.ZyAhkq5th8bf = tc.protocol
		got := p.oX3U75782l()
		want := url.Values{"token": {"secret&=测试"}, "node_id": {"-5"}}
		if tc.want != "" {
			want["node_type"] = []string{tc.want}
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%q: %#v", tc.protocol, got)
		}
		decoded, err := url.ParseQuery(got.Encode())
		if err != nil || !reflect.DeepEqual(decoded, want) {
			t.Fatal(decoded, err)
		}
	}
	for _, action := range []string{"config", "user", "push", "../custom?a=1"} {
		if got := p.GetUniUrl(action); got != "https://panel.test/api/v1/server/UniProxy/"+action {
			t.Fatal(got)
		}
	}
	cfg.W_ogGtF = "https://panel.test"
	if got := p.GetUniUrl("config"); got != "https://panel.testapi/v1/server/UniProxy/config" {
		t.Fatal("must preserve concatenation", got)
	}
	if got := p.nToI_PrB("missing"); got != nil {
		t.Fatal(got)
	}
	p.eyq32s.AddETag("url", `W/"abc"`, []byte("body"))
	if got := p.nToI_PrB("url"); !reflect.DeepEqual(got, http.Header{"IF-NONE-MATCH": {`W/"abc"`}}) {
		t.Fatal(got)
	}
	p.eyq32s.AddETag("url", "", nil)
	if got := p.nToI_PrB("url"); got != nil {
		t.Fatal(got)
	}
}

func TestClientRTTILayout(t *testing.T) {
	{
		rt := reflect.TypeOf(JZA3kKksUe{})
		if rt.Size() != 16 || rt.NumField() != 2 {
			t.Fatalf("JZA3kKksUe layout: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "EA_5pkz" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "" || f.Anonymous || f.Type.String() != "*config.FCEAfQJwDA7e" {
			t.Fatalf("JZA3kKksUe field 0: %#v", f)
		}
		if f := rt.Field(1); f.Name != "eyq32s" || f.Offset != 8 || f.Type.Size() != 8 || string(f.Tag) != "" || f.Anonymous || f.Type.String() != "*etagcache.NG0hCZyRHK" {
			t.Fatalf("JZA3kKksUe field 1: %#v", f)
		}
	}
}
