package cNT24rG4F

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestOriginalCacheIdentityBranches(t *testing.T) {
	old := nodeConfigurationAt50A2AE0
	defer func() { nodeConfigurationAt50A2AE0 = old }()
	c := &nodeConfigurationAt50A2AE0
	c.EtRWBSEhLQc = "v2board"
	c.NJSpG6 = "webapi"
	c.UnkaBGBuOUXV = "fixture-key"
	c.VfiX6TsWZxF = "https://first.invalid"
	first := AVktbLizjJl()
	want := fmt.Sprintf("%x", md5.Sum([]byte("v2board-fixture-key")))
	if first != want {
		t.Fatal("webapi identity")
	}
	c.VfiX6TsWZxF = "https://second.invalid"
	if AVktbLizjJl() != first {
		t.Fatal("original identity excludes URL")
	}
	c.UnkaBGBuOUXV = ""
	c.MKrU9XEJ = "fixture-key"
	if AVktbLizjJl() != first {
		t.Fatal("legacy key fallback")
	}
	c.NJSpG6 = "db"
	c.ApfXLZLmFXh, c.Mk5cDiQzaj, c.CsBHZVUxTO = "::1", 3306, "fixture"
	if AVktbLizjJl() != fmt.Sprintf("%x", md5.Sum([]byte("v2board-::1-3306-fixture"))) {
		t.Fatal("database identity")
	}
	c.NJSpG6 = "unknown"
	if AVktbLizjJl() != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Fatal("unknown API empty identity")
	}
}
