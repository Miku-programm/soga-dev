package config

import "testing"

func TestOriginalContainerEnvironment(t *testing.T) {
	for _, value := range []string{"", "true", "TRUE", "True", "1", " true", "true\n", "false"} {
		t.Setenv("IS_DOCKER", value)
		if got := VSHlrRuK7_(); got != (value == "true") {
			t.Fatalf("IS_DOCKER=%q: %v", value, got)
		}
	}
}
