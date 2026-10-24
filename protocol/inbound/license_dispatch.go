package inbound

import (
	jsonv1 "encoding/json"
	"encoding/json/v2"
	"errors"
	"strconv"

	"recovered.local/soga/net/relay"
	"recovered.local/soga/transport/realitykdf"
	"recovered.local/soga/crypto/rand"
)

// aNU8Vqc_q recovers 0x1db88c0. All three decoding failures return the
// original status 100005. A valid decrypted response returns its code only.
func aNU8Vqc_q(data []byte, request *realitykdf.KUD0DUCag) string {
	response := new(realitykdf.D4bEdSC7On)
	if json.Unmarshal(data, response, jsonv1.DefaultOptionsV1()) != nil {
		return "100005"
	}
	plain, err := response.Decode(request)
	if err != nil {
		return "100005"
	}
	result := new(realitykdf.UnopiKFbiaEJ)
	if json.Unmarshal(plain, result, jsonv1.DefaultOptionsV1()) != nil {
		return "100005"
	}
	return strconv.Itoa(result.IJ9cKUVV2nxt)
}

// ogDnYO recovers 0x1dbab40. The shuffled endpoints and per-attempt request
// construction are original. Transport/build failures try the next endpoint;
// the first HTTP 200 without an error is decoded without trying another host.
func ogDnYO() (string, error) {
	endpoints := []string{
		"https://agdfgkdfg.singledance.xyz/soga_license",
		"https://boiudihsdfk.singledance.xyz/soga_license",
		"https://sxp7hu5t.cilema.xyz/soga_license",
		"https://sksjd.vaxilu.com/soga_license",
	}
	// Original crypto-random Fisher-Yates shuffle at 0xe8e520.
	rand.Tij6hx(len(endpoints), func(i, j int) { endpoints[i], endpoints[j] = endpoints[j], endpoints[i] })
	for _, endpoint := range endpoints {
		request, err := bj4dWcdriU()
		if err != nil {
			continue
		}
		body, status, err := relay.ZBovtj1xky(endpoint, nil, request, nil)
		if status != 200 || err != nil {
			continue
		}
		return aNU8Vqc_q(body, request), nil
	}
	return "", errors.New("check soga failed")
}

// Nya2yw2fbyOg recovers 0x1d67660 and the original session-state dispatch.
func Nya2yw2fbyOg() error {
	code, err := ogDnYO()
	if err != nil {
		return err
	}
	if handler, ok := euYBu_M()[code]; ok {
		return handler()
	}
	return errors.New("unknown error")
}
