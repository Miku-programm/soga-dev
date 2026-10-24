package gk1z9v

import "fmt"

type k7S_Ie2j func([]byte, []byte) (BwCT6jfasz6, error)
type yarvofa4fXjT func([]byte, []byte) (G1Wopa, error)

// 原 map 0x5096338/0x5096340；init.0/1/2 的注册值及函数指针已核对。
var encryptAt5096338 = map[IPas09]k7S_Ie2j{
	1: func(key, iv []byte) (BwCT6jfasz6, error) { // 0x1d22680
		c, err := txowtYpKbXXG(key, iv)
		if err != nil {
			return nil, err
		}
		return c, nil
	},
	3: hYUXBKqsWpGa, 4: q4ORlubuA2c, 5: rOYZKFEuVKV,
}
var decryptAt5096340 = map[IPas09]yarvofa4fXjT{
	1: func(key, iv []byte) (G1Wopa, error) { // 0x1d22720
		c, err := txowtYpKbXXG(key, iv)
		if err != nil {
			return nil, err
		}
		return c, nil
	},
	3: bYH58jCAX, 4: a1jSJ0g, 5: cSfaA7LJF,
}

func rOYZKFEuVKV(_, _ []byte) (BwCT6jfasz6, error) { return &P8gDoqFWbS{}, nil } // 0x1d1c100
func cSfaA7LJF(_, _ []byte) (G1Wopa, error)        { return &P8gDoqFWbS{}, nil } // 0x1d1c120
func RsWvpjES(sec IPas09, key, iv []byte) (BwCT6jfasz6, error) { // 0x1d1ac00
	if f, ok := encryptAt5096338[sec]; ok {
		return f(key, iv)
	}
	return nil, fmt.Errorf("sec is not exist: %v", sec)
}
func DtSw_eAbIQ(sec IPas09, key, iv []byte) (G1Wopa, error) { // 0x1d1adc0
	if f, ok := decryptAt5096340[sec]; ok {
		return f(key, iv)
	}
	return nil, fmt.Errorf("sec is not exist: %v", sec)
}
