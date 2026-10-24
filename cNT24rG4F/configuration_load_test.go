package cNT24rG4F

import (
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func isolateLoader(t *testing.T) {
	t.Helper()
	isolateRegistry(t)
	path, commandLine := configurationPathAt5096DB0, flag.CommandLine
	t.Cleanup(func() { configurationPathAt5096DB0, flag.CommandLine = path, commandLine })
	configurationPathAt5096DB0 = filepath.Join(t.TempDir(), "soga.conf")
	flag.CommandLine = flag.NewFlagSet("recovered-config-test", flag.ContinueOnError)
}

func TestCompleteConfigurationPrecedence(t *testing.T) {
	isolateLoader(t)
	type config struct {
		Value  string `conf:"recovered_value" default:"default"`
		Env    string `conf:"recovered_env" default:"default"`
		Case   string `conf:"recovered_case" default:"default"`
		Empty  string `conf:"recovered_empty" default:"default"`
		Quoted string `conf:"recovered_quoted"`
	}
	c := config{}
	_NA1kft(&c)
	input := "recovered_value=file\nrecovered_env=file\nrecovered_case=file\nrecovered_empty=file\nrecovered_quoted=\"x # y\"\nunknown=kept\n"
	if err := os.WriteFile(configurationPathAt5096DB0, []byte(input), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("recovered_value", "env")
	t.Setenv("recovered_env", "from-env # comment")
	t.Setenv("RECOVERED_CASE", "must-not-match")
	t.Setenv("recovered_empty", "")
	t.Setenv("DNS_RECOVERED_TOKEN", " token=with=equals # comment ")
	flag.CommandLine.String("recovered_value", "unvisited", "")
	flag.CommandLine.String("recovered_empty", "unvisited", "")
	if err := flag.CommandLine.Parse([]string{"-recovered_value=flag # comment", "-recovered_empty="}); err != nil {
		t.Fatal(err)
	}
	var seen []string
	afterLoadAt5099300 = []func(){func() { seen = append(seen, c.Value) }, func() { seen = append(seen, c.Env) }}
	if err := TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	if c != (config{Value: "flag", Env: "from-env", Case: "file", Empty: "", Quoted: `"x`}) {
		t.Fatalf("merged configuration: %#v", c)
	}
	if !reflect.DeepEqual(seen, []string{"flag", "from-env"}) || configurationAt5095AA8["dns_recovered_token"] != "token=with=equals" || configurationAt5095AA8["unknown"] != "kept" {
		t.Fatalf("callbacks or extra keys differ: %v %#v", seen, configurationAt5095AA8)
	}
	if defaultsAt5095AA0["recovered_value"] != "default" {
		t.Fatal("loader must clone defaults")
	}
}

func TestMissingFileCreationAndFailureOrder(t *testing.T) {
	isolateLoader(t)
	configurationPathAt5096DB0 = filepath.Join(t.TempDir(), "new", "nested", "soga.conf")
	c := struct {
		Value int `conf:"recovered_number" default:"7"`
	}{}
	_NA1kft(&c)
	if err := TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(configurationPathAt5096DB0)
	if err != nil || len(data) != 0 || c.Value != 7 {
		t.Fatalf("missing file: %q %v %v", data, err, c)
	}
	configurationPathAt5096DB0 = t.TempDir()
	configurationAt5095AA8 = map[string]string{"old": "value"}
	called := false
	afterLoadAt5099300 = []func(){func() { called = true }}
	if err := TvBECnD9fM6(); err == nil {
		t.Fatal("reading directory must fail")
	}
	if called || !reflect.DeepEqual(configurationAt5095AA8, defaultsAt5095AA0) {
		t.Fatal("read failure order differs")
	}
	configurationPathAt5096DB0 += "/absent/"
	configurationAt5095AA8 = map[string]string{"old": "value"}
	if err := TvBECnD9fM6(); err == nil {
		t.Fatal("file creation with trailing slash must fail")
	}
	if configurationAt5095AA8["old"] != "value" {
		t.Fatal("creation failure must precede resetting the map")
	}
}

func TestConfigPathPrefixDoesNotClean(t *testing.T) {
	old := configurationPathAt5096DB0
	t.Cleanup(func() { configurationPathAt5096DB0 = old })
	for _, tc := range []struct{ input, want string }{
		{"", "/etc/soga/"}, {"node/soga.conf", "/etc/soga/node/soga.conf"},
		{"../soga.conf", "/etc/soga/../soga.conf"}, {"/tmp//soga.conf", "/tmp//soga.conf"},
	} {
		MkwNWciPHS(tc.input)
		if configurationPathAt5096DB0 != tc.want {
			t.Errorf("path %q => %q", tc.input, configurationPathAt5096DB0)
		}
	}
}
