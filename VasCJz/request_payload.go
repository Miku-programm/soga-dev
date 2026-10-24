package VasCJz

import (
	"encoding/base64"
	"fmt"
	"net/http"

	XCrypto "recovered.local/soga/xray/crypto"
	"recovered.local/soga/xray/utils"
)

// GetRequestHeader 对应 0x12513a0；原调用 Add，保留标准化后同名的多个值。
func (c *R0J19auWof) GetRequestHeader() http.Header {
	h := make(http.Header)
	for key, value := range c.F0XSwqNYoaP {
		h.Add(key, value)
	}
	utils.TryDefaultHeadersWith(h, "fetch")
	return h
}

// GetRequestHeaderWithPayload 对应 0x1251540。原样本使用未经补默认值的 key。
func (c *R0J19auWof) GetRequestHeaderWithPayload(payload []byte) http.Header {
	h := c.GetRequestHeader()
	key := c.HFEkUe4a9Sas
	data := base64.RawURLEncoding.EncodeToString(payload)
	for i := 0; len(data) != 0; i++ {
		r := c.GetNormalizedUplinkChunkSize()
		n := min(len(data), int(int32(XCrypto.RandBetween(int64(r.AS9zG362), int64(r.HHTPwX0N)))))
		part := data[:n]
		data = data[n:]
		h.Set(fmt.Sprintf("%s-%d", key, i), part)
	}
	return h
}

// GetRequestCookiesWithPayload 对应 0x1251880；空载荷返回非 nil 空切片。
func (c *R0J19auWof) GetRequestCookiesWithPayload(payload []byte) []*http.Cookie {
	cookies := []*http.Cookie{}
	key := c.HFEkUe4a9Sas
	data := base64.RawURLEncoding.EncodeToString(payload)
	for i := 0; len(data) != 0; i++ {
		r := c.GetNormalizedUplinkChunkSize()
		n := min(len(data), int(int32(XCrypto.RandBetween(int64(r.AS9zG362), int64(r.HHTPwX0N)))))
		part := data[:n]
		data = data[n:]
		cookies = append(cookies, &http.Cookie{Name: fmt.Sprintf("%s_%d", key, i), Value: part})
	}
	return cookies
}
