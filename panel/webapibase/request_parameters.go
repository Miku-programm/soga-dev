// 根据原样本指令恢复的请求参数函数，不是作者源码原文。
package webapibase

import (
	"net/http"
	"net/url"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/etagcache"
	"strconv"
)

// JZA3kKksUe 对应0x3639bf0，原名称前缀wyihwaFO，大小16字节。
// 配置读取与请求参数在此文件；获取、转换及提交在同包其他源码文件。
type JZA3kKksUe struct {
	EA_5pkz *config.FCEAfQJwDA7e
	eyq32s  *etagcache.NG0hCZyRHK
}

// GetAPIConfig 对应0x1745e80，返回原配置指针。
func (p *JZA3kKksUe) GetAPIConfig() *config.FCEAfQJwDA7e { return p.EA_5pkz }

// m8JdgCTbwl7 对应0x1745ea0，匹配原值，不做大小写或空白规范化。
func (p *JZA3kKksUe) m8JdgCTbwl7() string {
	switch p.EA_5pkz.ZyAhkq5th8bf {
	case "ss":
		return "shadowsocks" // 0x1747fe0 + 0x1747f80静态解码结果。
	case "v2ray":
		return "v2ray"
	case "trojan":
		return "trojan"
	}
	return ""
}

// oX3U75782l 对应0x1745f80；node_type由0x1748220、0x174dd40、0x17481c0解出。
func (p *JZA3kKksUe) oX3U75782l() url.Values {
	values := url.Values{"token": {p.EA_5pkz.Bd08Li}, "node_id": {strconv.Itoa(p.EA_5pkz.Yf1KqSqJpNI7)}}
	if nodeType := p.m8JdgCTbwl7(); nodeType != "" {
		values["node_type"] = []string{nodeType}
	}
	return values
}

// GetUniUrl 对应0x17463c0，直接连接三段字符串；不自动补斜杠或转义。
func (p *JZA3kKksUe) GetUniUrl(action string) string {
	// 原0x17486e0及0x1748680解码结果，不含前导斜杠。
	return p.EA_5pkz.W_ogGtF + "api/v1/server/UniProxy/" + action
}

// nToI_PrB 对应0x1746240。原0x1748500/0x17484a0解出大写请求头。
func (p *JZA3kKksUe) nToI_PrB(key string) http.Header {
	etag := p.eyq32s.GetETag(key)
	if etag == "" {
		return nil
	}
	return http.Header{"IF-NONE-MATCH": {etag}}
}
