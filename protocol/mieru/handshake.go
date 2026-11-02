package mieru

import (
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"math"
	mathrand "math/rand"
	"net"
	"sync/atomic"
	D "recovered.local/soga/buf/pool"
	C "recovered.local/soga/config"
	U "recovered.local/soga/user"
	"time"
)

var trafficOutAt50983D0 = errors.New("user traffic out")
var userHintRequiredAt50983E0 = errors.New("user hint required but not present")

// panelForceUserHint 记录面板下发的 user_hint_is_mandatory。
var panelForceUserHint atomic.Bool

// SetPanelForceUserHint 由 mieru inbound 启动时按节点配置设置。
func SetPanelForceUserHint(v bool) { panelForceUserHint.Store(v) }

// ofDJn56sP0 @0x1c4dc00 and closure @0x1c5c940. Match hints first;
// the optional fallback group is searched afterwards. The hint is the username,
// while the password derives the encryption key with a NUL and username salt.
func ofDJn56sP0(nonce, ciphertext []byte, iterate func(func(*U.WIZQ28) bool)) (U.PgaNd3014U, cipher.AEAD, []byte, error) {
	force := C.MieruForceUserHint() || panelForceUserHint.Load()
	var matching, remaining []*U.WIZQ28
	iterate(func(user *U.WIZQ28) bool {
		if pyUlp642nl(user.GetUsername(), nonce) {
			matching = append(matching, user)
		} else if !force {
			remaining = append(remaining, user)
		}
		return true
	})
	if force && len(matching) == 0 {
		return nil, nil, nil, userHintRequiredAt50983E0
	}
	salts := pVQO7MqTE4ao(time.Now())
	for _, group := range [2][]*U.WIZQ28{matching, remaining} {
		for _, user := range group {
			secret := gUBL99(user.GetPassword(), []byte(user.GetUsername()))
			for _, salt := range salts {
				key := pbkdf2.Key(secret[:], salt, 64, 32, sha256.New)
				aead, err := chacha20poly1305.NewX(key)
				if err != nil {
					continue
				}
				plain, err := aead.Open(nil, nonce, ciphertext, nil)
				if err == nil {
					return user, aead, plain, nil
				}
			}
		}
	}
	return nil, nil, nil, fmt.Errorf("decrypt failed: no matching user")
}

// cMGKnW @0x1c4c480. The replay fingerprint is recorded before authentication;
// failed authentication does not undo it. A required but missing hint skips the
// extra read delay, while replay errors otherwise override decrypt errors.
func cMGKnW(conn net.Conn, iterate func(func(*U.WIZQ28) bool)) (*rZDq_kinV4G, U.PgaNd3014U, *IqB1foHWg, error) {
	_ = conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	defer conn.SetReadDeadline(time.Time{})
	first := make([]byte, 72)
	if _, err := io.ReadFull(conn, first); err != nil {
		return nil, nil, nil, fmt.Errorf("read first segment: %w", err)
	}
	replay := yCMyVg().n7hmYhVTVA(first)
	user, aead, plain, err := ofDJn56sP0(first[:24], first[24:], iterate)
	if err != nil {
		if errors.Is(err, userHintRequiredAt50983E0) {
			return nil, nil, nil, err
		}
		xvkpqpGhETY(conn)
		if replay {
			return nil, nil, nil, fmt.Errorf("replay attack detected (decrypt failed)")
		}
		return nil, nil, nil, err
	}
	if replay {
		xvkpqpGhETY(conn)
		return nil, nil, nil, fmt.Errorf("replay attack detected (decrypt succeeded)")
	}
	m, err := Ox0UfI5kIfcR(plain)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unmarshal first metadata: %w", err)
	}
	return &rZDq_kinV4G{aMKXGpHBfg: conn, n49jm3JWYo: aead, xUtUMCF3grV: append([]byte(nil), first[:24]...)}, user, m, nil
}

func vR0ruI(minimum, maximum int) int { //0x1c4e420
	if maximum <= minimum {
		return minimum
	}
	amount := float64(mathrand.Intn(maximum - minimum + 1))
	r := mathrand.Float64()
	return int(amount*math.Sqrt(r*r*r)) + minimum
}
func xvkpqpGhETY(conn net.Conn) { //0x1c4e200
	delay := vR0ruI(1000, 10000) + fsAE63SXO(50000)
	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(delay) * time.Millisecond))
	size := 1 << (fsAE63SXO(4) + 12)
	b := D.PoolAt50961D8.Get(size)
	defer D.PoolAt50961D8.Put(b)
	amount := vR0ruI(0, size-254) + fsAE63SXO(256)
	if amount > 0 {
		_, _ = io.ReadAtLeast(conn, b, amount)
	}
}
