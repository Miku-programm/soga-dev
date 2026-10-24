package cNT24rG4F

import (
	"testing"
	"time"
)

func TestSSInvalidAuthInlineDefaults(t *testing.T) {
	old := configAt50C8440
	t.Cleanup(func() { configAt50C8440 = old })
	configAt50C8440.X4_AmA8j7 = true
	configAt50C8440.UA4PZU = -1
	configAt50C8440.PWFghdU = 0
	configAt50C8440.SndqWD26h = -10
	c := SSInvalidAccessSettings()
	if !c.SEU_Z687ugp || c.GxuEkr != 30 || c.QS255wlN7m2W != time.Minute || c.NjiRkFvidOaA != 10*time.Minute {
		t.Fatalf("defaults %+v", c)
	}
	configAt50C8440.UA4PZU = 7
	configAt50C8440.PWFghdU = 11
	configAt50C8440.SndqWD26h = 13
	c = SSInvalidAccessSettings()
	if c.GxuEkr != 7 || c.QS255wlN7m2W != 11*time.Second || c.NjiRkFvidOaA != 13*time.Second {
		t.Fatalf("values %+v", c)
	}
}
