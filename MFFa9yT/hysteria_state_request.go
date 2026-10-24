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

// hCXJN9YU (0x1dc0a00) builds the original 17-field request. Its domain
// helper 0x1dbdbe0, metrics helper 0x1dbf980 and all 17 map values were
// independently traced against the same ELF. They match licensePayload;
// Evidence: hysteria-payload-traces.json and hysteria-request-logic-traces.json.
// This shares independently proved field construction, not a substitute request.
func hCXJN9YU() (*tkbUwvW3ex.KUD0DUCag, error) {
	data, err := json.Marshal(licensePayload(), jsonv1.DefaultOptionsV1())
	if err != nil {
		return nil, fmt.Errorf("marshal data: %w", err)
	}
	return tkbUwvW3ex.LBab24XA(data)
}

// rWeAdhA96 is 0x1dc3280: JSON/decrypt/JSON failures each return 100005.
func rWeAdhA96(data []byte, request *tkbUwvW3ex.KUD0DUCag) string {
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

// vLhTYZafwY is the independent request entry used by the Hysteria workers.
// (0x1dc53e0) retains independent request construction on every
// attempt and stops on the first error-free HTTP 200 even if decoding fails.
func vLhTYZafwY() (string, error) {
	endpoints := []string{
		"https://agdfgkdfg.singledance.xyz/soga_license",
		"https://boiudihsdfk.singledance.xyz/soga_license",
		"https://sxp7hu5t.cilema.xyz/soga_license",
		"https://sksjd.vaxilu.com/soga_license",
	}
	uLofn6j.Tij6hx(len(endpoints), func(i, j int) { endpoints[i], endpoints[j] = endpoints[j], endpoints[i] })
	return requestHysteriaConnectionState(endpoints)
}

// Extracted original retry loop; production supplies only the shuffled original hosts.
func requestHysteriaConnectionState(endpoints []string) (string, error) {
	for _, endpoint := range endpoints {
		request, err := hCXJN9YU()
		if err != nil {
			continue
		}
		body, status, err := G6bHD5JA.ZBovtj1xky(endpoint, nil, request, nil)
		if status != 200 || err != nil {
			continue
		}
		return rWeAdhA96(body, request), nil
	}
	return "", errors.New("check soga failed")
}
