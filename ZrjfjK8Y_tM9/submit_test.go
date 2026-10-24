package ZrjfjK8Y_tM9

import (
	"io"
	"net/http"
	"net/http/httptest"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/q5bqCgTdxaTM"
	"testing"
)

func TestSubmitTrafficWireFormatAndEmptyData(t *testing.T) {
	calls := 0
	body := ""
	path := ""
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		path = r.URL.Path
		b, _ := io.ReadAll(r.Body)
		body = string(b)
	}))
	defer server.Close()
	p := panelForRequestTests()
	p.EA_5pkz.W_ogGtF = server.URL + "/"
	if err := p.SubmitData(&q5bqCgTdxaTM.FGmNba4MhioY{JNcHHwTZ7: LFWTTXF7JR.CQhZrJ4dK{{Vy9ybu: 1, PNtn9KAV: []string{"127.0.0.1"}}}}); err != nil || calls != 0 {
		t.Fatal("原方法不单独提交IP列表")
	}
	traffic := LFWTTXF7JR.CWfcUv{{Hy7Q97RROK: 7, XJkGTQf_qkAt: 10, I6LXObODav: 20}, {Hy7Q97RROK: -1, XJkGTQf_qkAt: -3, I6LXObODav: 0}, {Hy7Q97RROK: 7, XJkGTQf_qkAt: 30, I6LXObODav: 40}}
	if err := p.SubmitData(&q5bqCgTdxaTM.FGmNba4MhioY{HTgPyE3VytZd: traffic}); err != nil {
		t.Fatal(err)
	}
	if calls != 1 || path != "/api/v1/server/UniProxy/push" || body != `{"-1":[-3,0],"7":[30,40]}` {
		t.Fatalf("calls=%d path=%q body=%q", calls, path, body)
	}
}
func TestSubmitRetainsNilEntryPanic(t *testing.T) {
	p := panelForRequestTests()
	defer func() {
		if recover() == nil {
			t.Fatal("nil流量记录应保留原panic")
		}
	}()
	_ = p.eALpAMr(LFWTTXF7JR.CWfcUv{nil})
}
