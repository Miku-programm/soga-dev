package auth

import (
	"bytes"
	"errors"
	"io"
	B "recovered.local/soga/util/bytes"
	"strconv"
)

// Auth names the observed 32-byte argument to AMwLSL6GewF: two strings.
// The original parameter type/field names were not retained in RTTI.
type Auth struct {
	Username string
	Password string
}

// LT72UC follows the original named uint8 type.
type LT72UC uint8

// Error: 0x192bc60.
func (e LT72UC) Error() string { return "SOCKS error: " + strconv.Itoa(int(e)) }

var authFailedAt5098520 = errors.New("auth failed") // init: 0x192bb80
// AMwLSL6GewF: 0x192bd00. The exact original negotiation advertises only
// username/password when auth is nonnil. It checks greeting version and auth
// status, but does not validate the final reply's VER/REP/RSV fields.
func AMwLSL6GewF(c io.ReadWriter, address []byte, command byte, auth *Auth) ([]byte, error) {
	buffer := make([]byte, 259)
	greeting := []byte{5, 1, 0}
	if auth != nil {
		greeting[2] = 2
	}
	if _, err := c.Write(greeting); err != nil {
		return nil, err
	}
	if _, err := io.ReadFull(c, buffer[:2]); err != nil {
		return nil, err
	}
	if buffer[0] != 5 {
		return nil, errors.New("SOCKS version error")
	}
	switch buffer[1] {
	case 2:
		if auth == nil {
			return nil, authFailedAt5098520
		}
		b := B.BaDekw7qe{}
		b[b.Grow(1)] = 1
		b[b.Grow(1)] = byte(len(auth.Username))
		at := b.Grow(len(auth.Username))
		copy(b[at:], auth.Username)
		b[b.Grow(1)] = byte(len(auth.Password))
		at = b.Grow(len(auth.Password))
		copy(b[at:], auth.Password)
		if _, err := c.Write(b); err != nil {
			return nil, err
		}
		if _, err := io.ReadFull(c, buffer[:2]); err != nil {
			return nil, err
		}
		if buffer[1] != 0 {
			return nil, errors.New("rejected username/password")
		}
	case 0:
	default:
		return nil, errors.New("SOCKS need auth")
	}
	request := bytes.Join([][]byte{{5, command, 0}, address}, nil)
	if _, err := c.Write(request); err != nil {
		return nil, err
	}
	if _, err := io.ReadFull(c, buffer[:3]); err != nil {
		return nil, err
	}
	return YhN2Ke2lC(c, buffer)
}

// YhN2Ke2lC: 0x192c5e0. The returned address aliases buffer, including when
// a final fixed-size read fails; only earlier framing failures return a nil slice.
func YhN2Ke2lC(r io.Reader, buffer []byte) ([]byte, error) {
	if len(buffer) < 259 {
		return nil, io.ErrShortBuffer
	}
	if _, err := io.ReadFull(r, buffer[:1]); err != nil {
		return nil, err
	}
	var length int
	switch buffer[0] {
	case 1:
		length = 7
	case 4:
		length = 19
	case 3:
		if _, err := io.ReadFull(r, buffer[1:2]); err != nil {
			return nil, err
		}
		length = int(buffer[1]) + 4
		_, err := io.ReadFull(r, buffer[2:length])
		return buffer[:length], err
	default:
		return nil, LT72UC(8)
	}
	_, err := io.ReadFull(r, buffer[1:length])
	return buffer[:length], err
}
