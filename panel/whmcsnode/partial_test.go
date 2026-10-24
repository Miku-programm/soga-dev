package whmcsnode

import (
	"io"
	"net/http"
	"net/http/httptest"
	C "recovered.local/soga/config"
	"strings"
	"testing"
)

func TestOriginalPartialFetchResults(t *testing.T) {
	old := firstUsersAt378731B
	defer func() { firstUsersAt378731B = old }()
	for _, failed := range []string{"node/", "userList/", "nodeRule/"} {
		t.Run(failed, func(t *testing.T) {
			firstUsersAt378731B = true
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.Contains(r.URL.Path, failed) {
					io.WriteString(w, `{"code":409}`)
					return
				}
				switch {
				case strings.Contains(r.URL.Path, "node/"):
					io.WriteString(w, `{"code":200,"data":{"trojan_port":8443,"client_limit":9}}`)
				case strings.Contains(r.URL.Path, "userList/"):
					io.WriteString(w, `{"code":200,"data":[{"uid":1}]}`)
				case strings.Contains(r.URL.Path, "nodeRule/"):
					io.WriteString(w, `{"code":200,"data":{"rules":[{"type":"domain","pattern":"partial.test"}]}}`)
				}
			}))
			defer s.Close()
			p := uZ2az_6UM(&C.FCEAfQJwDA7e{W_ogGtF: s.URL + "/"})
			v, e := p.GetAPIData()
			if v == nil || e == nil {
				t.Fatalf("%+v %v", v, e)
			}
			if (v.H47AzZsYHB1N == nil) != (failed == "node/") {
				t.Fatal("node result lost")
			}
			if (len(v.VTme1SuHL6i_) == 0) != (failed == "userList/") {
				t.Fatal("user result lost")
			}
			if (len(v.HYnfYWa) == 0) != (failed == "nodeRule/") {
				t.Fatal("rule result lost")
			}
			if failed == "nodeRule/" && v.VTme1SuHL6i_[0].Xmrna8 != 9 {
				t.Fatal("partial result user limits not synchronized")
			}
		})
	}
}
