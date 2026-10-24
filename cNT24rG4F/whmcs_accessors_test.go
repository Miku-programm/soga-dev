package cNT24rG4F

import "testing"

func TestWHMCSSettingsPreservesOriginalValue(t *testing.T) {
	old := configAt50A2300
	defer func() { configAt50A2300 = old }()
	configAt50A2300 = pKGHLtJsa{VD6q_75hbej2: 8443, IlqtsyWu: 1.5, X4Dd6nOZ_7: "grpc", Dd5Xza: true, W2WFtBWP: 2, NQ3CpRlonU: "/path", Ew45bIap: "service", Z9bFboOQlPiw: "password", Et04j58qH: "aes-128-gcm", EKv8Zjg7dn85: "auth_aes128_md5", HmKQa3xt: "plain"}
	got := WHMCSSettings()
	if got != configAt50A2300 {
		t.Fatal("mapping")
	}
	got.VD6q_75hbej2 = 1
	if configAt50A2300.VD6q_75hbej2 != 8443 {
		t.Fatal("value must not alias global")
	}
}
