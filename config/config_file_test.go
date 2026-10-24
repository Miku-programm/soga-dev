package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestBinaryFileParsingBranches(t *testing.T) {
	previous := configurationAt5095AA8
	t.Cleanup(func() { configurationAt5095AA8 = previous })
	configurationAt5095AA8 = map[string]string{"retained": "yes"}
	path := filepath.Join(t.TempDir(), "soga.conf")
	input := "\r\n\t# comment\nmissing separator\nTYPE=first\ntype = second\nTOKEN=a=b=c\nINLINE=value # later\nHASH=a#b\nTAB=a\t#b\n=empty key\nQUOTED=\"literal\"\n\u2003ÄKEY\u2003=\u3000值\u3000\n"
	if err := os.WriteFile(path, []byte(input), 0600); err != nil {
		t.Fatal(err)
	}
	if err := zyxKcF82B(path); err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"retained": "yes", "type": "second", "token": "a=b=c",
		"inline": "value # later", "hash": "a#b", "tab": "a\t#b",
		"": "empty key", "quoted": "\"literal\"", "äkey": "值",
	}
	if !reflect.DeepEqual(configurationAt5095AA8, want) {
		t.Fatalf("文件解析与已确认分支不符：%#v", configurationAt5095AA8)
	}
}

func TestBinaryReadFailurePreservesMap(t *testing.T) {
	previous := configurationAt5095AA8
	t.Cleanup(func() { configurationAt5095AA8 = previous })
	configurationAt5095AA8 = map[string]string{"key": "unchanged"}
	if err := zyxKcF82B(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("文件读取失败必须返回错误")
	}
	if !reflect.DeepEqual(configurationAt5095AA8, map[string]string{"key": "unchanged"}) {
		t.Fatal("读取失败改变了配置映射")
	}
}

func TestBinaryInlineCommentBranches(t *testing.T) {
	cases := []struct{ input, want string }{
		{"  value # comment  ", "value"},
		{"value#data", "value#data"},
		{"value\t#data", "value\t#data"},
		{"#data", "#data"},
		{" # comment", ""},
		{"\u2003value\u3000", "value"},
		{"\"a # b\"", "\"a"},
		{"", ""},
	}
	for _, tc := range cases {
		if got := zuOmYdCNcr(tc.input); got != tc.want {
			t.Errorf("输入 %q：得到 %q，预期 %q", tc.input, got, tc.want)
		}
	}
}
