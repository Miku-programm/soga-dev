package tw357_KzLSN2

import (
	"bytes"
	E "recovered.local/soga/hfOuV8fOe4L"
	"reflect"
	"testing"
)

func TestNativeObfsFactory(t *testing.T) {
	for name, want := range map[E.PMorzD7u]any{"none": &rqkJYP_xJtF{}, "plain": &rqkJYP_xJtF{}, "http_simple": &nArN8vNMTaR{}, "http_post": &nArN8vNMTaR{}, "simple_obfs_http": &k7yarZ7PgV6{}, "simple_obfs_tls": &cqVDRTx{}, "tls1.2_ticket_auth": &g7X2glAwFd91{}} {
		got, e := Z_Mqu9of(name, new(bytes.Buffer))
		if e != nil || reflect.TypeOf(got) != reflect.TypeOf(want) {
			t.Fatal(name, got, e)
		}
	}
	for _, name := range []E.PMorzD7u{"", "None", "random_head", "tls1.2_ticket_fastauth"} {
		if got, e := Z_Mqu9of(name, new(bytes.Buffer)); got != nil || e == nil || e.Error() != "obfs "+string(name)+" not support" {
			t.Fatal(got, e)
		}
	}
}
