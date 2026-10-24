package ZrjfjK8Y_tM9

import (
	"encoding/json"
	"recovered.local/soga/LFWTTXF7JR"
	"strings"
	"testing"
)

func TestPanelNodeTransportAndSnapshot(t *testing.T) {
	for _, tc := range []struct{ network, raw, path, service string }{
		{"", "bad ignored", "", ""}, {"ws", `{"path":"/socket"}`, "/socket", ""}, {"grpc", `{"serviceName":"rpc"}`, "", "rpc"}, {"grpc", "null", "", ""},
	} {
		n := &U3nQ9iaQj{BDLnnhaR5: 8443, PqqmTY8L: -1, FYeHum: "aes-128-gcm", D07VWhUsih: "secret"}
		n.X6Jwaaf492C = LFWTTXF7JR.BoulTn31(tc.network)
		n.MDoXQFch8aWk = []byte(tc.raw)
		// 非 JSON 的 settings 对非 ws/grpc 也会使快照 Marshal 失败；先独立验证此路径。
		result, err := n.ToPanelNode(37)
		if err != nil {
			t.Fatal(err)
		}
		if result.EMlXEc != 37 || result.SwxJNgPpg.CKsOpayd != 8443 || !result.SwxJNgPpg.UVeahqDBT || result.SwxJNgPpg.ChJnngzb != tc.path || result.SwxJNgPpg.KWffzRDF != tc.service || result.JhIYwEukjd.PBDb3Wc != tc.path || result.JhIYwEukjd.QVlaSCZ4RCK != tc.service || result.FnMUny.VuhWqou != "secret" {
			t.Fatalf("wrong mapping: %+v", result)
		}
		if tc.network == "" && n.X6Jwaaf492C != "tcp" {
			t.Fatal("network mutation absent")
		}
		if result.AjXcBfRcAXR == nil {
			t.Fatal("original returns nonnil empty DNS routes")
		}
	}
	n := &U3nQ9iaQj{}
	node, err := n.ToPanelNode(1)
	if err != nil {
		t.Fatal(err)
	}
	var saved U3nQ9iaQj
	if err := json.Unmarshal([]byte(node.Z1TIgHOQ), &saved); err != nil || saved.X6Jwaaf492C != "tcp" {
		t.Fatalf("snapshot: %v %q", err, node.Z1TIgHOQ)
	}
}

func TestPanelNodeSettingsErrors(t *testing.T) {
	for _, network := range []string{"ws", "grpc"} {
		n := &U3nQ9iaQj{X6Jwaaf492C: LFWTTXF7JR.BoulTn31(network), MDoXQFch8aWk: []byte("{")}
		out, err := n.ToPanelNode(1)
		if out != nil || err == nil || !strings.HasPrefix(err.Error(), "parse "+network+" settings failed: ") || !strings.HasSuffix(err.Error(), ", result: {") {
			t.Fatalf("%v %v", out, err)
		}
	}
}
func TestPanelRouteConversionPreservesIDsAndFallback(t *testing.T) {
	n := &U3nQ9iaQj{RqE21w: IUXU5ZXqeiRC{
		{CEfy7Qj: 9, ReTngXeu: "block", Fu2GAU2hpa1b: []byte(`["domain:example.com",""]`)},
		{CEfy7Qj: 42, ReTngXeu: "block", Fu2GAU2hpa1b: []byte("one,two")},
		{ReTngXeu: "dns", CoShU6j: "1.1.1.1", Fu2GAU2hpa1b: []byte(`["full:a"]`)},
		{ReTngXeu: "dns", CoShU6j: "", Fu2GAU2hpa1b: []byte(`["x"]`)},
		{ReTngXeu: "DNS", CoShU6j: "2.2.2.2", Fu2GAU2hpa1b: []byte(`["x"]`)},
	}}
	detectors := n.GetDetectors()
	if len(detectors) != 4 {
		t.Fatal(len(detectors))
	}
	for i, id := range []int{9, 9, 42, 42} {
		if detectors[i].GetId() != id {
			t.Fatal("route ID lost")
		}
	}
	if !detectors[0].Test("sub.example.com") || detectors[1].Test("") || detectors[2].GetRule() != "one" {
		t.Fatal("match behavior lost")
	}
	out, err := n.ToPanelNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if len(out.AjXcBfRcAXR) != 1 {
		t.Fatalf("DNS: %+v", out.AjXcBfRcAXR)
	}
	dns := out.AjXcBfRcAXR[0]
	if dns.CM0eMsqyP != "ipv4_first" || dns.K8BNfziIgb[0] != "1.1.1.1" || dns.CZ80ik[0] != "full:a" {
		t.Fatalf("%+v", dns)
	}
	if (&U3nQ9iaQj{}).GetDetectors() == nil {
		t.Fatal("nil detectors")
	}
}
