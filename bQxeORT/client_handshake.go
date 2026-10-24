package bQxeORT

import (
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"
)

// Handshake recovers 0x1cf3880: chained public-key exchanges, an ephemeral
// X25519/ML-KEM exchange, padding, and the original cached-session fast path.
func (s *GPFxUP0_) Handshake(conn net.Conn) (*RfBfjVN4XE, error) {
	if s.QUnOciyca6eb == nil {
		return nil, errors.New("uninitialized")
	}
	result := &RfBfjVN4XE{DrpXao: conn, HiUDG1Fl: true}
	paddingTotal, segments, delays := OODk_yfADN(s.Q8W_Oz, s.PVif2PW)
	request := make([]byte, s.JhDzI4ehWF+1266+paddingTotal)
	rand.Read(request[:16])
	remaining := request[16 : 16+s.JhDzI4ehWF]
	var shared []byte
	var previous cipher.Stream
	for i, public := range s.QUnOciyca6eb {
		encodedLength := 32
		if p, ok := public.(*ecdh.PublicKey); ok {
			private, _ := ecdh.X25519().GenerateKey(rand.Reader)
			copy(remaining, private.PublicKey().Bytes())
			var err error
			shared, err = private.ECDH(p)
			if err != nil {
				return nil, err
			}
		}
		if p, ok := public.(*mlkem.EncapsulationKey768); ok {
			var ciphertext []byte
			shared, ciphertext = p.Encapsulate()
			copy(remaining, ciphertext)
			encodedLength = 1088
		}
		if s.S799IXI != 0 {
			GdvNyLM(s.JTmPRiuw[i], request[:16]).XORKeyStream(remaining, remaining[:encodedLength])
		}
		if previous != nil {
			previous.XORKeyStream(remaining, remaining[:32])
		}
		if i == len(s.QUnOciyca6eb)-1 {
			break
		}
		previous = GdvNyLM(shared, request[:16])
		remaining = remaining[encodedLength:]
		previous.XORKeyStream(remaining, s.J2battR3S[i+1][:])
		remaining = remaining[32:]
	}
	handshakeCipher := WrnUSnC6(request[:16], shared, result.HiUDG1Fl)
	if s.BsCV7uFJFW != 0 {
		s.QgCGo4Q6lLF.RLock()
		if time.Now().Before(s.EQa3Gaz) {
			result.R5MQKRg26Z = s
			result.ENEXKXNftX = append(s.Wliqugg3J89, shared...)
			handshakeCipher.Seal(request[:s.JhDzI4ehWF+16], nil, []byte{0, 32}, nil)
			handshakeCipher.Seal(request[:s.JhDzI4ehWF+34], nil, s.HMnsiUUhd, nil)
			s.QgCGo4Q6lLF.RUnlock()
			result.A9M03eHa = request[:s.JhDzI4ehWF+66]
			result.AVDV5UPQ = WrnUSnC6(request[s.JhDzI4ehWF+34:s.JhDzI4ehWF+66], result.ENEXKXNftX, result.HiUDG1Fl)
			if s.S799IXI == 2 {
				result.DrpXao = &LUKUT5W0R6{DrpXao: conn, D1_SMs0WWKz: GdvNyLM(result.ENEXKXNftX, request[:16]), IdEXNQK5b1k: len(result.A9M03eHa), DhRIe1pBzWdS: make([]byte, 0, 5), LMy55MYg_l: 16, TfVMo2pTBnt: make([]byte, 0, 5)}
			}
			return result, nil
		}
		s.QgCGo4Q6lLF.RUnlock()
	}
	requestPart := request[s.JhDzI4ehWF+16 : s.JhDzI4ehWF+1266]
	handshakeCipher.Seal(requestPart[:0], nil, []byte{4, 208}, nil)
	kemPrivate, _ := mlkem.GenerateKey768()
	xPrivate, _ := ecdh.X25519().GenerateKey(rand.Reader)
	publicPair := append(kemPrivate.EncapsulationKey().Bytes(), xPrivate.PublicKey().Bytes()...)
	handshakeCipher.Seal(requestPart[:18], nil, publicPair, nil)
	padding := request[s.JhDzI4ehWF+1266:]
	handshakeCipher.Seal(padding[:0], nil, []byte{byte((paddingTotal - 18) >> 8), byte(paddingTotal - 18)}, nil)
	handshakeCipher.Seal(padding[:18], nil, padding[18:paddingTotal-16], nil)
	segments[0] += s.JhDzI4ehWF + 1266
	unsent := request
	for i, n := range segments {
		if n > 0 {
			if _, err := conn.Write(unsent[:n]); err != nil {
				return nil, err
			}
			unsent = unsent[n:]
		}
		if i < len(delays) {
			time.Sleep(delays[i])
		}
	}
	serverPublic := make([]byte, 1136)
	if _, err := io.ReadFull(conn, serverPublic); err != nil {
		return nil, err
	}
	// Original ignores this Open error and continues parsing the buffer.
	handshakeCipher.Open(serverPublic[:0], nonceMaximumAt509A1F0, serverPublic, nil)
	kemShared, err := kemPrivate.Decapsulate(serverPublic[:1088])
	if err != nil {
		return nil, err
	}
	xPublic, err := ecdh.X25519().NewPublicKey(serverPublic[1088:1120])
	if err != nil {
		return nil, err
	}
	xShared, err := xPrivate.ECDH(xPublic)
	if err != nil {
		return nil, err
	}
	sessionKey := make([]byte, 64)
	copy(sessionKey, kemShared)
	copy(sessionKey[32:], xShared)
	result.ENEXKXNftX = append(sessionKey, shared...)
	result.AVDV5UPQ = WrnUSnC6(publicPair, result.ENEXKXNftX, result.HiUDG1Fl)
	result.EufjD7UrtI = WrnUSnC6(serverPublic[:1120], result.ENEXKXNftX, result.HiUDG1Fl)
	token := make([]byte, 32)
	if _, err := io.ReadFull(conn, token); err != nil {
		return nil, err
	}
	if _, err := result.EufjD7UrtI.Open(token[:0], nil, token, nil); err != nil {
		return nil, err
	}
	if s.BsCV7uFJFW != 0 {
		seconds := binary.BigEndian.Uint16(token[:2])
		if seconds != 0 {
			s.QgCGo4Q6lLF.Lock()
			s.EQa3Gaz = time.Now().Add(time.Duration(seconds) * time.Second)
			s.Wliqugg3J89 = sessionKey
			s.HMnsiUUhd = token[:16]
			s.QgCGo4Q6lLF.Unlock()
		}
	}
	paddingLength := make([]byte, 18)
	if _, err := io.ReadFull(conn, paddingLength); err != nil {
		return nil, err
	}
	if _, err := result.EufjD7UrtI.Open(paddingLength[:0], nil, paddingLength, nil); err != nil {
		return nil, err
	}
	result.WWjo2UGyM = make([]byte, binary.BigEndian.Uint16(paddingLength[:2]))
	if s.S799IXI == 2 {
		result.DrpXao = &LUKUT5W0R6{DrpXao: conn, D1_SMs0WWKz: GdvNyLM(result.ENEXKXNftX, request[:16]), EO5NXI6ex6: GdvNyLM(result.ENEXKXNftX, token[:16]), DhRIe1pBzWdS: make([]byte, 0, 5), LMy55MYg_l: len(result.WWjo2UGyM), TfVMo2pTBnt: make([]byte, 0, 5)}
	}
	return result, nil
}
