package ZrjfjK8Y_tM9

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchResultsEndpointsAndJSON(t *testing.T) {
	response := `{"server_port":443,"base_config":{"pull_interval":"12.9","push_interval":20}}`
	path := ""
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { path = r.URL.Path; _, _ = w.Write([]byte(response)) }))
	defer server.Close()
	p := panelForRequestTests()
	p.EA_5pkz.W_ogGtF = server.URL + "/"
	node, err := p.GetUniNode()
	if err != nil || path != "/api/v1/server/UniProxy/config" || node.BDLnnhaR5 != 443 || node.WNvzoiMPp.F2WCOEc5Eu6.Value() != 12 {
		t.Fatalf("node=%+v path=%q err=%v", node, path, err)
	}
	response = `{"users":[{"id":12,"uuid":"u","speed_limit":3.5,"device_limit":2}]}`
	users, err := p.GetUsersResult()
	if err != nil || path != "/api/v1/server/UniProxy/user" || len(users.Ta593Rq2lSa) != 1 || users.Ta593Rq2lSa[0].IDLJXX6F != 12 {
		t.Fatalf("users=%+v path=%q err=%v", users, path, err)
	}
	response = "null"
	node, err = p.GetUniNode()
	if node == nil || err != nil {
		t.Fatalf("null节点应为零值对象: %v %v", node, err)
	}
	users, err = p.GetUsersResult()
	if users == nil || err != nil || users.Ta593Rq2lSa != nil {
		t.Fatalf("null用户应为零值对象: %v %v", users, err)
	}
}
func TestFetchResultsFailures(t *testing.T) {
	response := `{"message":"denied"}`
	status := 200
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(status); _, _ = w.Write([]byte(response)) }))
	defer server.Close()
	p := panelForRequestTests()
	p.EA_5pkz.W_ogGtF = server.URL + "/"
	if node, err := p.GetUniNode(); node != nil || err == nil || err.Error() != `parse node failed: <nil>, result: {"message":"denied"}` {
		t.Fatalf("message: %v %v", node, err)
	}
	if users, err := p.GetUsersResult(); users != nil || err == nil || err.Error() != `parse users failed: <nil>, result: {"message":"denied"}` {
		t.Fatalf("message: %v %v", users, err)
	}
	response = "invalid"
	if node, err := p.GetUniNode(); node != nil || err == nil || !strings.HasPrefix(err.Error(), "parse node failed: ") || !strings.HasSuffix(err.Error(), ", result: invalid") {
		t.Fatalf("JSON: %v %v", node, err)
	}
	if users, err := p.GetUsersResult(); users != nil || err == nil || !strings.HasPrefix(err.Error(), "parse users failed: ") {
		t.Fatalf("JSON: %v %v", users, err)
	}
	status = 500
	response = "failed"
	if node, err := p.GetUniNode(); node != nil || err == nil || err.Error() != "get node failed: request failed, result: failed, status code: 500" {
		t.Fatalf("HTTP: %v %v", node, err)
	}
	if users, err := p.GetUsersResult(); users != nil || err == nil || err.Error() != "get users failed: request failed, result: failed, status code: 500" {
		t.Fatalf("HTTP: %v %v", users, err)
	}
}
