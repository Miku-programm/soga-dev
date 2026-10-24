package crypto

import (
	"crypto/rand"
	"math/big"
)

// RandBetween 0x124c5c0。上界不包含在内，反向区间会交换；原调用忽略随机源错误。
func RandBetween(from, to int64) int64 {
	if from == to {
		return from
	}
	if from > to {
		from, to = to, from
	}
	n, _ := rand.Int(rand.Reader, big.NewInt(to-from))
	return from + n.Int64()
}
