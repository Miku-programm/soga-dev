package config

// LicensePanelFields is a new cross-package bridge for the direct global reads
// in inbound.gS8fWl/bj4dWcdriU. It preserves the original unnormalized values.
// Fields correspond to 0x50a2ae0, 0x50a2b10, 0x50a2b20, 0x50a2b50,
// 0x50a2b60 and 0x50a2b68; version is the shared global at 0x5029ee0.
func LicensePanelFields() (panel, api, webURL, dbHost string, dbPort int, dbName, version string) {
	c := &nodeConfigurationAt50A2AE0
	return c.EtRWBSEhLQc, c.NJSpG6, c.VfiX6TsWZxF, c.ApfXLZLmFXh, c.Mk5cDiQzaj, c.CsBHZVUxTO, versionAt5029EE0
}
