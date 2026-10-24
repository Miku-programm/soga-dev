package cNT24rG4F

import "testing"

func TestLicensePanelFieldsPreserveRawValues(t *testing.T) {
	old, oldVersion := nodeConfigurationAt50A2AE0, versionAt5029EE0
	t.Cleanup(func() { nodeConfigurationAt50A2AE0 = old; versionAt5029EE0 = oldVersion })
	c := &nodeConfigurationAt50A2AE0
	c.EtRWBSEhLQc = "v2board"
	c.NJSpG6 = "webapi"
	c.VfiX6TsWZxF = "https://host/path"
	c.ApfXLZLmFXh = "::1"
	c.Mk5cDiQzaj = -5
	c.CsBHZVUxTO = "db"
	versionAt5029EE0 = "version"
	panel, api, address, host, port, name, version := LicensePanelFields()
	if panel != "v2board" || api != "webapi" || address != "https://host/path" || host != "::1" || port != -5 || name != "db" || version != "version" {
		t.Fatal("direct original fields changed")
	}
}
