// 依据样本 soga-go/xray/utils 的六个函数恢复；不是上游文件的直接替代。
package utils

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/klauspost/cpuid/v2"
)

var (
	chromeVersion = ChromeVersion()
	chromeUA      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + strconv.Itoa(chromeVersion) + ".0.0.0 Safari/537.36"
	chromeCHUA    = getGreasedChUa(chromeVersion, "chrome")
	// 原 0x124c77c 的 Edg/ 前没有空格。
	edgeUA   = chromeUA + "Edg/" + strconv.Itoa(chromeVersion) + ".0.0.0"
	edgeCHUA = getGreasedChUa(chromeVersion, "edge")
)

// ChromeVersion 对应 0x124c840。种子来自原 CPUInfo 的五个字段，
// 由 0x1094120 各写入位置确定；不把尚未初始化的 BSS 零值当作种子。
func ChromeVersion() int {
	c := &cpuid.CPU
	seed := int64(c.PhysicalCores + c.LogicalCores + c.Family + c.Model + c.CacheLine)
	return chromeVersionAt(time.Now(), seed)
}

func chromeVersionAt(now time.Time, seed int64) int {
	r := rand.New(rand.NewSource(seed))
	t := time.Date(2026, 1, 13, 0, 0, 0, 0, time.UTC)
	v := 144
	for t.Before(now) {
		t = t.AddDate(0, 0, r.Intn(21)+25)
		v++
	}
	return v - 1
}

// 原静态数据 0x5039af0、0x5039b10、0x5039b30、0x5039b50。
var greaseChars = []string{" ", "(", ":", "-", ".", "/", ")", ";", "=", "?", "_"}
var greaseVersions = []string{"8", "99", "24"}
var permutations3 = [][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
var permutations4 = [][4]int{
	{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1},
	{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0},
	{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0},
	{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0},
}

// 原入口 0x124ca60；负版本数的越界行为保留。
func getUngreasedChUa(version int, browser string) []string {
	result := make([]string, 0, 4)
	result = append(result, `"Not`+greaseChars[version%len(greaseChars)]+"A"+greaseChars[(version+1)%len(greaseChars)]+`Brand";v="`+greaseVersions[version%len(greaseVersions)]+`"`)
	v := strconv.Itoa(version)
	result = append(result, `"Chromium";v="`+v+`"`)
	switch browser {
	case "edge":
		result = append(result, `"Microsoft Edge";v="`+v+`"`)
	case "chrome":
		result = append(result, `"Google Chrome";v="`+v+`"`)
	}
	return result
}

// 原入口 0x124ce20：置换下标是写入位置，不是读取位置。
func getGreasedChUa(version int, browser string) string {
	values := getUngreasedChUa(version, browser)
	var order []int
	switch len(values) {
	case 1:
		order = []int{0}
	case 2:
		order = []int{version % 2, (version + 1) % 2}
	case 3:
		order = permutations3[version%len(permutations3)][:]
	default:
		order = permutations4[version%len(permutations4)][:]
	}
	result := make([]string, len(values))
	for i, j := range order {
		result[j] = values[i]
	}
	return strings.Join(result, ", ")
}

// 原入口 0x124d0a0。直接写入的 CH/DNT 大小写和 Header.Set 的
// 标准化结果并不相同，须保留这个可观察差异。
func applyMasqueradedHeaders(h http.Header, browser, kind string) {
	chromium := browser == "chrome" || browser == "edge"
	switch browser {
	case "chrome", "edge":
		ua, ch := chromeUA, chromeCHUA
		if browser == "edge" {
			ua, ch = edgeUA, edgeCHUA
		}
		h["Sec-CH-UA"] = []string{ch}
		h["Sec-CH-UA-Mobile"] = []string{"?0"}
		h["Sec-CH-UA-Platform"] = []string{`"Windows"`}
		h["DNT"] = []string{"1"}
		h.Set("User-Agent", ua)
		h.Set("Accept-Language", "en-US,en;q=0.9")
	case "firefox":
		h.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:140.0) Gecko/20100101 Firefox/140.0")
		h["DNT"] = []string{"1"}
		h.Set("Accept-Language", "en-US,en;q=0.5")
	case "golang":
		h.Del("User-Agent")
		return
	}
	setEmpty := func(key, value string) {
		if h.Get(key) == "" {
			h.Set(key, value)
		}
	}
	switch kind {
	case "ws":
		h.Set("Sec-Fetch-Mode", "websocket")
		h.Set("Sec-Fetch-Dest", "empty")
		h.Set("Sec-Fetch-Site", "same-origin")
		setEmpty("Cache-Control", "no-cache")
		setEmpty("Pragma", "no-cache")
		setEmpty("Accept", "*/*")
	case "nav":
		if chromium {
			setEmpty("Cache-Control", "max-age=0")
		}
		h.Set("Upgrade-Insecure-Requests", "1")
		if chromium {
			setEmpty("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/jxl,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		} else if browser == "firefox" {
			setEmpty("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		}
		h.Set("Sec-Fetch-Site", "none")
		h.Set("Sec-Fetch-Mode", "navigate")
		h.Set("Sec-Fetch-User", "?1")
		h.Set("Sec-Fetch-Dest", "document")
		h.Set("Priority", "u=0, i")
	case "fetch":
		h.Set("Sec-Fetch-Mode", "cors")
		h.Set("Sec-Fetch-Dest", "empty")
		h.Set("Sec-Fetch-Site", "same-origin")
		if chromium {
			setEmpty("Priority", "u=1, i")
		} else if browser == "firefox" {
			setEmpty("Priority", "u=4")
		}
		setEmpty("Cache-Control", "no-cache")
		setEmpty("Pragma", "no-cache")
		setEmpty("Accept", "*/*")
	}
}

// TryDefaultHeadersWith 对应 0x124ecc0。空值 UA 与缺失 UA 有不同语义。
func TryDefaultHeadersWith(h http.Header, kind string) {
	if len(h.Values("User-Agent")) == 0 {
		applyMasqueradedHeaders(h, "chrome", kind)
		return
	}
	switch ua := h.Get("User-Agent"); ua {
	case "edge", "chrome", "golang", "firefox":
		applyMasqueradedHeaders(h, ua, kind)
	}
}
