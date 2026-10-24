package realitykdf

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestDecodedTimeRangeStrings(t *testing.T) {
	// 黄金字节来自原 ELF 常量与指令的独立静态推导；包含被回调丢弃的字节。
	check := func(wantHex string, wantLen, wantCap int) func([]byte) string {
		return func(data []byte) string {
			t.Helper()
			if got := hex.EncodeToString(data); got != wantHex {
				t.Fatalf("完整解码结果 %s，期望 %s", got, wantHex)
			}
			if len(data) != wantLen || cap(data) != wantCap {
				t.Fatalf("回调切片大小 (%d,%d)，期望 (%d,%d)", len(data), cap(data), wantLen, wantCap)
			}
			return "callback-result"
		}
	}
	if got := haZl2FW(0x9a7a3b4e, 0xfabd, 0x2e32, 0xb0f3,
		check("83b0cc253032643a253032642d253032643a253032648e6e", 24, 24)); got != "callback-result" {
		t.Fatalf("格式解码回调返回值 %q", got)
	}
	if got := xybQauJZY_n(0xf2a99ecb, 0xa002, 0x65fb, 0x74,
		check("e697b6e997b4e6aeb5e8a784e58899e6a0bce5bc8fe99499e8afaf3a20257658f7", 33, 33)); got != "callback-result" {
		t.Fatalf("错误解码回调返回值 %q", got)
	}
	if got := haZl2FW(0x9a7a3b4e, 0xfabd, 0x2e32, 0xb0f3, qxmp4I0); got != "%02d:%02d-%02d:%02d" {
		t.Fatalf("格式字符串 %q", got)
	}
	if got := xybQauJZY_n(0xf2a99ecb, 0xa002, 0x65fb, 0x74, aWXs7MI); got != "时间段规则格式错误: %v" {
		t.Fatalf("错误格式 %q", got)
	}
}

func TestTimeRangeListParsing(t *testing.T) {
	for _, tc := range []struct {
		input, wantText, wantError string
		want                       ITMfakhJ
	}{
		{"", "00:00-24:00", "", ITMfakhJ{{BZwEA15n: 24}}},
		{"\u2003, ,\t,", "00:00-24:00", "", ITMfakhJ{{BZwEA15n: 24}}},
		{" 20:00-06:00 , , 06:00-20:00 ,24:00-24:00 ", "06:00-20:00,06:00-20:00,24:00-24:00", "",
			ITMfakhJ{{LXcSWg26p: 6, BZwEA15n: 20}, {LXcSWg26p: 6, BZwEA15n: 20}, {LXcSWg26p: 24, BZwEA15n: 24}}},
		{" 01:02-03:04 , bad , 04:00-05:00", "01:02-03:04", "时间段规则格式错误: bad",
			ITMfakhJ{{LXcSWg26p: 1, RBNUY7G: 2, BZwEA15n: 3, G1LUbU: 4}}},
		{" \t无效%v\t ,01:00-02:00", "", "时间段规则格式错误: 无效%v", ITMfakhJ{}},
		{"24:01-24:00", "", "时间段规则格式错误: 24:01-24:00", ITMfakhJ{}},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Z4FKfucI_A5(tc.input)
			errText := ""
			if err != nil {
				errText = err.Error()
			}
			if errText != tc.wantError || !reflect.DeepEqual(got, tc.want) || got.String() != tc.wantText {
				t.Fatalf("解析 %q 得到 %#v / %q / %q；期望 %#v / %q / %q", tc.input, got, got.String(), errText, tc.want, tc.wantText, tc.wantError)
			}
		})
	}
}

func TestTimeRangeStringDoesNotValidateOrNormalize(t *testing.T) {
	r := &D1Teq5S1f{LXcSWg26p: 123, RBNUY7G: -1, BZwEA15n: 0, G1LUbU: 99}
	if got := r.String(); got != "123:-1-00:99" {
		t.Fatalf("字符串 %q", got)
	}
	if got := (ITMfakhJ(nil)).String(); got != "" {
		t.Fatalf("空列表 %q", got)
	}
}

func TestTimeRangeNilElementPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("原函数不跳过空指针，应当触发 panic")
		}
	}()
	_ = (ITMfakhJ{nil}).String()
}

func TestTimeDecoderCallbacksUseCapacity(t *testing.T) {
	// 原回调检查 cap，允许扩展零长度切片后转换成字符串。
	if got := qxmp4I0(make([]byte, 0, 22)); len(got) != 19 {
		t.Fatalf("长度 %d", len(got))
	}
	if got := aWXs7MI(make([]byte, 0, 31)); len(got) != 31 {
		t.Fatalf("长度 %d", len(got))
	}
	for _, test := range []struct {
		name string
		run  func()
	}{
		{"format", func() { qxmp4I0(make([]byte, 21)) }},
		{"error", func() { aWXs7MI(make([]byte, 30)) }},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("容量不足应当触发 panic")
				}
			}()
			test.run()
		})
	}
}
