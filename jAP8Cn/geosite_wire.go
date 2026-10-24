package jAP8Cn

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"recovered.local/soga/GYzZ5v"
)

// mcrrOY（0xe9a080）顺序查找 GeoSiteList 中的首个匹配项。
// 长度前缀之后直接切片，越界时保留原程序的 panic 行为。
func mcrrOY(data []byte, code string) (*GYzZ5v.P3tOWWIHzX, error) {
	for len(data) != 0 {
		number, kind, n := protowire.ConsumeTag(data)
		if n < 0 {
			return nil, fmt.Errorf("failed to decode tag: %v", protowire.ParseError(n))
		}
		data = data[n:]
		if number == 1 && kind == protowire.BytesType {
			length, n := protowire.ConsumeVarint(data)
			if n < 0 {
				return nil, fmt.Errorf("failed to decode length: %v", protowire.ParseError(n))
			}
			data = data[n:]
			site, err := sUZtoMEpw_(data[:length], code)
			data = data[length:]
			if err != nil {
				return nil, err
			}
			if site != nil {
				return site, nil
			}
		} else {
			// 原程序先移除了 tag，再从当前位置调用 ConsumeField。
			// 此处保留该顺序，不改成 ConsumeFieldValue。
			_, _, n := protowire.ConsumeField(data)
			if n < 0 {
				return nil, fmt.Errorf("failed to consume field: %v", protowire.ParseError(n))
			}
			data = data[n:]
		}
	}
	return nil, nil
}

// sUZtoMEpw_（0xe9a740）仅检查消息的第一个字段是否为 country_code。
func sUZtoMEpw_(data []byte, code string) (*GYzZ5v.P3tOWWIHzX, error) {
	if len(data) == 0 {
		return nil, nil
	}
	number, kind, n := protowire.ConsumeTag(data)
	if n < 0 {
		return nil, fmt.Errorf("failed to decode geoSite tag: %v", protowire.ParseError(n))
	}
	if number != 1 || kind != protowire.BytesType {
		return nil, nil
	}
	data = data[n:]
	name, n := protowire.ConsumeBytes(data)
	if n < 0 {
		return nil, fmt.Errorf("failed to decode country code: %v", protowire.ParseError(n))
	}
	data = data[n:]
	if !strings.EqualFold(string(name), code) {
		return nil, nil
	}
	site := &GYzZ5v.P3tOWWIHzX{I0dBJNoqAhc: code}
	// 原调用是 Unmarshal，不是 Merge；已消耗的 country_code 字段
	// 没有传给它，也没有在解码后重新写入。
	if err := proto.Unmarshal(data, site); err != nil {
		return nil, fmt.Errorf("failed to unmarshal GeoSite: %v", err)
	}
	return site, nil
}
