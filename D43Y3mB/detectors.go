package D43Y3mB

import (
	"encoding/json"
	"fmt"
	Log "recovered.local/soga/muc279Mpa6S4"
	X "recovered.local/soga/xx5fRyrFw"
)

// ToDetector 对应 0x16b87c0；只接受原 reg/domain 标签。
// D43Y3mB.(*Lv22mM8eYq).ToDetector：原样本入口 0x16b87c0。
func (p *Lv22mM8eYq) ToDetector() (*X.IXZNQqqXDsV, error) {
	switch p.USQvAO4WtdOX {
	case "reg":
		return X.NewPanelDetector(p.DPx7K7u, "regexp:"+p.LgZENghw), nil
	case "domain":
		return X.NewPanelDetector(p.DPx7K7u, "domain:"+p.LgZENghw), nil
	}
	return nil, fmt.Errorf("unknown detect rule type: %v", p.USQvAO4WtdOX)
}

// ToDetectors 对应 0x16b8a20；mode不参与选择，未知规则记录后跳过。
// D43Y3mB.(*Y4Pw3Ojd6HO).ToDetectors：原样本入口 0x16b8a20。
func (p *Y4Pw3Ojd6HO) ToDetectors() X.Ij9r21E0UOt3 {
	out := make(X.Ij9r21E0UOt3, 0)
	for _, r := range p.GPETCf2.NPpVYG5M9 {
		d, err := r.ToDetector()
		if err != nil {
			Log.TRTbrquD(4, err)
			continue
		}
		out = append(out, d)
	}
	return out
}

// GetDetectors 对应 0x16ba320。
// D43Y3mB.(*GMLWI3).GetDetectors：原样本入口 0x16ba320。
func (p *GMLWI3) GetDetectors(address string) (X.Ij9r21E0UOt3, error) {
	body, err := p.GetCommon(address)
	if err != nil {
		return nil, fmt.Errorf("get detect rule failed: %v", err)
	}
	var r Y4Pw3Ojd6HO
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse detect rule failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get detect rule failed: %v", string(body))
	}
	return r.ToDetectors(), nil
}
