package xiaov2board

import (
	"encoding/json"
	L "recovered.local/soga/panel/types"
	"testing"
)

// 响应格式与 xiaov2b MieruSettings::nodeConfig 及 UniProxy user 接口一致。
func TestMieruNodeAndUsers(t *testing.T) {
	var n BdT6k1HU8L
	data := `{"listen_ip":"0.0.0.0","server_port":2012,"network":"tcp","network_settings":{"mtu":1400,"multiplexing":"MULTIPLEXING_LOW","server_port_range":"2012-2022"},"port_bindings":[{"portRange":"2012-2022","protocol":"TCP"}],"mtu":1400,"traffic_pattern":"CAESAA==","user_hint_is_mandatory":true,"base_config":{"push_interval":60,"pull_interval":60}}`
	if err := json.Unmarshal([]byte(data), &n); err != nil {
		t.Fatal(err)
	}
	out, err := n.ToPanelNode(3)
	if err != nil {
		t.Fatal(err)
	}
	want := L.Xh4IiOa0{MQE3K4mq: 2012, GGSLzD: "TCP", CLWr9bQsb: "CAESAA==", PortRange: "2012-2022", ForceUserHint: true}
	if out.ABA4hq != want {
		t.Fatalf("%+v", out.ABA4hq)
	}

	n = BdT6k1HU8L{}
	if err := json.Unmarshal([]byte(`{"server_port":3000,"network":"udp","port_bindings":[{"port":3000,"protocol":"UDP"}]}`), &n); err != nil {
		t.Fatal(err)
	}
	out, _ = n.ToPanelNode(3)
	if want := (L.Xh4IiOa0{MQE3K4mq: 3000, GGSLzD: "UDP"}); out.ABA4hq != want {
		t.Fatalf("%+v", out.ABA4hq)
	}

	n = BdT6k1HU8L{}
	if err := json.Unmarshal([]byte(`{"protocol":"mieru","server_port":2012,"network":"tcp","port_bindings":[{"port":2012,"protocol":"TCP"}]}`), &n); err != nil {
		t.Fatal(err)
	}
	if err := n.ocVy8aVcKm("mieru"); err != nil || !n.edYRMY8 {
		t.Fatal(n, err)
	}

	old := stateAt378733C
	t.Cleanup(func() { stateAt378733C = old })
	var r IMyQ5CxaikH
	if err := json.Unmarshal([]byte(`{"users":[{"id":7,"uuid":"u-7","username":"7","password":"u-7"},{"id":8,"uuid":"u-8"}]}`), &r); err != nil {
		t.Fatal(err)
	}
	users := r.ToPanelUsers("mieru")
	for i, w := range []L.DO9qqM9{{Ejujyaj: "7", GltkJGvV: "u-7"}, {Ejujyaj: "8", GltkJGvV: "u-8"}} {
		got, ok := users[i].JAfXWwzYrW.(*L.DO9qqM9)
		if !ok || *got != w {
			t.Fatalf("user %d: %+v", i, users[i].JAfXWwzYrW)
		}
	}
}
