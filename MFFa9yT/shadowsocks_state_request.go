package MFFa9yT

import (
	jsonv1 "encoding/json"
	"encoding/json/v2"
	"errors"
	"fmt"
	"strconv"

	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/tkbUwvW3ex"
	"recovered.local/soga/uLofn6j"
)

// lCtbLGhP (0x1dcb5e0) builds the original 17-field request. Its domain
// helper 0x1dc6de0, metrics helper 0x1dc9140 and all 17 map values were
// independently traced against the same ELF. They match licensePayload;
// Evidence: ss-payload-traces.json and ss-request-logic-traces.json.
// This shares independently proved field construction, not a substitute request.
func lCtbLGhP() (*tkbUwvW3ex.KUD0DUCag, error) {
	data, err := json.Marshal(licensePayload(), jsonv1.DefaultOptionsV1())
	if err != nil {
		return nil, fmt.Errorf("marshal data: %w", err)
	}
	return tkbUwvW3ex.LBab24XA(data)
}

// nxUOMW0sgb is 0x1dcdfc0: JSON/decrypt/JSON failures each return 100005.
func nxUOMW0sgb(data []byte, request *tkbUwvW3ex.KUD0DUCag) string {
	response := new(tkbUwvW3ex.D4bEdSC7On)
	if json.Unmarshal(data, response, jsonv1.DefaultOptionsV1()) != nil {
		return "100005"
	}
	plain, err := response.Decode(request)
	if err != nil {
		return "100005"
	}
	result := new(tkbUwvW3ex.UnopiKFbiaEJ)
	if json.Unmarshal(plain, result, jsonv1.DefaultOptionsV1()) != nil {
		return "100005"
	}
	return strconv.Itoa(result.IJ9cKUVV2nxt)
}

// fzRe58 (0x1dcf180) retains independent request construction on every
// attempt and stops on the first error-free HTTP 200 even if decoding fails.
func fzRe58() (string, error) {
	endpoints := []string{
		"https://agdfgkdfg.singledance.xyz/soga_license",
		"https://boiudihsdfk.singledance.xyz/soga_license",
		"https://sxp7hu5t.cilema.xyz/soga_license",
		"https://sksjd.vaxilu.com/soga_license",
	}
	uLofn6j.Tij6hx(len(endpoints), func(i, j int) { endpoints[i], endpoints[j] = endpoints[j], endpoints[i] })
	return requestShadowsocksConnectionState(endpoints)
}

// Extracted original retry loop; production supplies only the shuffled original hosts.
func requestShadowsocksConnectionState(endpoints []string) (string, error) {
	for _, endpoint := range endpoints {
		request, err := lCtbLGhP()
		if err != nil {
			continue
		}
		body, status, err := G6bHD5JA.ZBovtj1xky(endpoint, nil, request, nil)
		if status != 200 || err != nil {
			continue
		}
		return nxUOMW0sgb(body, request), nil
	}
	return "", errors.New("check soga failed")
}
