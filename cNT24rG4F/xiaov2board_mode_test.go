package cNT24rG4F

import "testing"

func TestXiaoV2NodeModeExactGlobals(t *testing.T) {
	old := nodeConfigurationAt50A2AE0
	t.Cleanup(func() { nodeConfigurationAt50A2AE0 = old })
	for _, tc := range []struct {
		panel, sub string
		want       bool
	}{{"xiaov2board", "v2node", true}, {"XIAOV2BOARD", "v2node", false}, {"xiaov2board", "V2Node", false}, {"xboard", "v2node", false}, {"xiaov2board", "", false}} {
		nodeConfigurationAt50A2AE0.EtRWBSEhLQc = tc.panel
		nodeConfigurationAt50A2AE0.AG_mSFklcvj = tc.sub
		if got := XiaoV2NodeMode(); got != tc.want {
			t.Fatal(tc, got)
		}
	}
}
