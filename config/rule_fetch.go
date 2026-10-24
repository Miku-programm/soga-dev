package config

import (
	"io"
	"net/http"

	"recovered.local/soga/crypto/hashutil"
	"recovered.local/soga/audit/matcher"
)

// 原 off_5029350 指向 0x509c010 的零值 48 字节 HTTP Client。
// Get@0x9cf860 和 do@0x9cfa20 对应标准 HTTP 客户端调用链；这里使用
// Go 标准库实现该依赖，不声称已找回原编译器的标准库补丁版本。
var ruleHTTPClientAt5029350 = new(http.Client)

// Klz7sa_01 对应 0x131a9a0。远程内容不按状态码过滤，不移除注释。
func Klz7sa_01() (matcher.VHWq49iRat, error) {
	if url := configAt509B660.BooNOdXcSa3U; url != "" {
		return fetchRuleLines(url)
	}
	return *blockListAt5095A90.Load(), nil
}

// Cf33JCugNk1 对应 0x133eb80；本地返回共享切片，保留原别名行为。
func Cf33JCugNk1() (matcher.VHWq49iRat, error) {
	if url := configAt5096E00.M4qDq3KsUi; url != "" {
		return fetchRuleLines(url)
	}
	return *whiteListAt5095AD8.Load(), nil
}

// 两个原函数中相同 GET/ReadAll/分行步骤的源码提取。
func fetchRuleLines(url string) (matcher.VHWq49iRat, error) {
	response, err := ruleHTTPClientAt5029350.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return matcher.VHWq49iRat(hashutil.DAaBs7_C(string(body), "\n", -1)), nil
}
