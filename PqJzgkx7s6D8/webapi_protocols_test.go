package PqJzgkx7s6D8_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	_ "recovered.local/soga/GgDwT38SMQf5"
	L "recovered.local/soga/LFWTTXF7JR"
	_ "recovered.local/soga/Se1QrdYD7Aa"
	C "recovered.local/soga/cNT24rG4F"
	_ "recovered.local/soga/mIfcoDkc6"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	_ "recovered.local/soga/uavUP6a"
	"reflect"
	"sync"
	"testing"
)

func TestFourOriginalSSPanelWebAPIRegistrations(t *testing.T) {
	for _, protocol := range []string{"ss", "ssr", "v2ray", "trojan"} {
		t.Run(protocol, func(t *testing.T) {
			var mu sync.Mutex
			paths := map[string]int{}
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				paths[r.URL.Path]++
				mu.Unlock()
				if r.URL.Query().Get("key") != "k" || r.URL.Query().Get("node_id") != "9" {
					t.Error(r.URL)
				}
				switch r.URL.Path {
				case "/mod_mu/nodes/9/info":
					fmt.Fprint(w, `{"ret":1,"data":{"server":"server;8081;5;ws;tls;path=/raw","mu_only":1,"custom_config":{"offset_port_node":"5555","method":"aes-128-gcm","server_key":"serverkey","v2_port":"7777","trojan_port":"8888"}}}`)
				case "/mod_mu/users":
					fmt.Fprint(w, `{"ret":1,"data":[{"id":4,"passwd":"pass","uuid":"12345678-1234-1234-1234-123456789abc","port":2222,"method":"aes-128-cfb","protocol":"auth_chain_a","obfs":"plain","is_multi_user":2}]}`)
				case "/mod_mu/func/detect_rules":
					fmt.Fprint(w, `{"ret":1,"data":[{"id":3,"regex":"full:example.org","type":1}]}`)
				case "/mod_mu/nodes":
					fmt.Fprint(w, `{"ret":1,"data":[{"id":9,"server":"srv;soga_port=9999|port=3333"}]}`)
				default:
					t.Error(r.URL.Path)
					http.NotFound(w, r)
				}
			}))
			defer srv.Close()
			p, e := Q.SMkpVNb(&C.FCEAfQJwDA7e{OS91X026uML: "sspanel-uim", ZyAhkq5th8bf: C.Mpn2JltM(protocol), F3PCiaIra: "webapi", W_ogGtF: srv.URL + "/", Bd08Li: "k", Yf1KqSqJpNI7: 9})
			if e != nil {
				t.Fatal(e)
			}
			d, e := p.GetAPIData()
			if e != nil {
				t.Fatal(e)
			}
			if len(d.VTme1SuHL6i_) != 1 || len(d.HYnfYWa) != 1 || !d.HYnfYWa[0].Test("example.org") {
				t.Fatal(d)
			}
			switch protocol {
			case "ss":
				if n := d.H47AzZsYHB1N.FnMUny; n.Am3QoeF != 5555 || n.HMkc4BgR != "aes-128-gcm" || d.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.YlfOV0DBBnx).T1oDDIgYJHh != "aes-128-gcm" {
					t.Fatal(n)
				}
			case "ssr":
				n := d.H47AzZsYHB1N.ASJxBbFS
				if n.PU1Di18 != 3333 || n.C33jbIzI != 9999 || reflect.ValueOf(n).FieldByName("kJbZdr").Len() != 1 {
					t.Fatal(n)
				}
			case "v2ray":
				if n := d.H47AzZsYHB1N.SwxJNgPpg; n.CKsOpayd != 7777 || n.VUwf1n9g != 5 || !n.UVeahqDBT || d.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.Yn8BJoL).FvQ2ltKOM7 != 5 {
					t.Fatal(n)
				}
			case "trojan":
				if n := d.H47AzZsYHB1N.JhIYwEukjd; n.GbD_2PU != 8888 || n.T4cuhG7MTbG != "tcp" {
					t.Fatal(n)
				}
			}
			mu.Lock()
			defer mu.Unlock()
			if paths["/mod_mu/users"] != 1 || paths["/mod_mu/func/detect_rules"] != 1 {
				t.Fatal(paths)
			}
		})
	}
}
