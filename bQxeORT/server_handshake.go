package bQxeORT

import (
	"bytes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

// Handshake @0x1cfbce0。保留原非前向密钥链、混合密钥交换、会话复用、
// 重放检测、AES/ChaCha尝试顺序、流量伪装和错误返回，不使用协议占位实现。
func (s *LgF3OB1QX) Handshake(raw net.Conn, captured *[]byte) (*RfBfjVN4XE, error) {
	if s.EiKprMhvVvc == nil {
		return nil, errors.New("uninitialized")
	}
	conn := &RfBfjVN4XE{DrpXao: raw, HiUDG1Fl: true}
	initial := make([]byte, s.DptoPRgY+16)
	if _, err := io.ReadFull(raw, initial); err != nil {
		return nil, err
	}
	if captured != nil {
		*captured = append(*captured, initial...)
	}
	remaining := initial[16:]
	var shared []byte
	var previous cipher.Stream
	for i, key := range s.EiKprMhvVvc {
		if previous != nil {
			previous.XORKeyStream(remaining, remaining[:32])
		}
		keySize := 32
		if _, ok := key.(*mlkem.DecapsulationKey768); ok {
			keySize = 1088
		}
		if s.F9GbshEXgaiY != 0 {
			GdvNyLM(s.SpB4gP6aCP[i], initial[:16]).XORKeyStream(remaining, remaining[:keySize])
		}
		switch private := key.(type) {
		case *ecdh.PrivateKey:
			public, err := ecdh.X25519().NewPublicKey(remaining[:keySize])
			if err != nil {
				return nil, err
			}
			if public.Bytes()[31] > 127 {
				return nil, errors.New("the highest bit of the last byte of the peer-sent X25519 public key is not 0")
			}
			shared, err = private.ECDH(public)
			if err != nil {
				return nil, err
			}
		case *mlkem.DecapsulationKey768:
			var err error
			shared, err = private.Decapsulate(remaining[:keySize])
			if err != nil {
				return nil, err
			}
		}
		if i == len(s.EiKprMhvVvc)-1 {
			break
		}
		previous = GdvNyLM(shared, initial[:16])
		remaining = remaining[keySize:]
		previous.XORKeyStream(remaining, remaining[:32])
		if !bytes.Equal(remaining[:32], s.CPktYs15Wt[i+1][:]) {
			return nil, fmt.Errorf("unexpected hash32: %v", remaining[:32])
		}
		remaining = remaining[32:]
	}
	handshakeCipher := WrnUSnC6(initial[:16], shared, conn.HiUDG1Fl)
	header := make([]byte, 18)
	if _, err := io.ReadFull(raw, header); err != nil {
		return nil, err
	}
	if captured != nil {
		*captured = append(*captured, header...)
	}
	lengthBytes := make([]byte, 2)
	if _, err := handshakeCipher.Open(lengthBytes[:0], nil, header, nil); err != nil {
		conn.HiUDG1Fl = !conn.HiUDG1Fl
		handshakeCipher = WrnUSnC6(initial[:16], shared, conn.HiUDG1Fl)
		if _, err = handshakeCipher.Open(lengthBytes[:0], nil, header, nil); err != nil {
			return nil, err
		}
	}
	if captured != nil {
		*captured = nil
	}
	size := int(binary.BigEndian.Uint16(lengthBytes))
	if size == 32 {
		if s.MOy_iE1vq == 0 && s.OIHUERfuz == 0 {
			return nil, errors.New("0-RTT is not allowed")
		}
		ticketCiphertext := make([]byte, 32)
		if _, err := io.ReadFull(raw, ticketCiphertext); err != nil {
			return nil, err
		}
		ticket, err := handshakeCipher.Open(nil, nil, ticketCiphertext, nil)
		if err != nil {
			return nil, err
		}
		s.VgRY0_5S9F37.RLock()
		session := s.GIHdzaFGhi8[[16]byte(ticket[:16])]
		s.VgRY0_5S9F37.RUnlock()
		if session == nil {
			noise := make([]byte, IaLb4jaWtI(1279, 2279))
			for {
				_, _ = rand.Read(noise)
				if _, err := QkDUgv(noise); err != nil {
					break
				}
			}
			_, _ = raw.Write(noise)
			return nil, errors.New("expired ticket")
		}
		if _, loaded := session.VaNnCL.LoadOrStore([32]byte(shared[:32]), true); loaded {
			return nil, errors.New("replay detected")
		}
		conn.ENEXKXNftX = append(session.MwUsryYRwlpa, shared...)
		conn.A9M03eHa = make([]byte, 16)
		_, _ = rand.Read(conn.A9M03eHa)
		conn.AVDV5UPQ = WrnUSnC6(conn.A9M03eHa, conn.ENEXKXNftX, conn.HiUDG1Fl)
		conn.EufjD7UrtI = WrnUSnC6(ticketCiphertext, conn.ENEXKXNftX, conn.HiUDG1Fl)
		if s.F9GbshEXgaiY == 2 {
			conn.DrpXao = &LUKUT5W0R6{DrpXao: raw,
				D1_SMs0WWKz: GdvNyLM(conn.ENEXKXNftX, conn.A9M03eHa),
				EO5NXI6ex6:  GdvNyLM(conn.ENEXKXNftX, initial[:16]), IdEXNQK5b1k: 16,
				DhRIe1pBzWdS: make([]byte, 0, 5), TfVMo2pTBnt: make([]byte, 0, 5)}
		}
		return conn, nil
	}
	if size < 1232 {
		return nil, errors.New("too short length")
	}
	clientHello := make([]byte, size)
	if _, err := io.ReadFull(raw, clientHello); err != nil {
		return nil, err
	}
	if _, err := handshakeCipher.Open(clientHello[:0], nil, clientHello, nil); err != nil {
		return nil, err
	}
	publicKEM, err := mlkem.NewEncapsulationKey768(clientHello[:1184])
	if err != nil {
		return nil, err
	}
	kemShared, kemCiphertext := publicKEM.Encapsulate()
	publicDH, err := ecdh.X25519().NewPublicKey(clientHello[1184:1216])
	if err != nil {
		return nil, err
	}
	privateDH, _ := ecdh.X25519().GenerateKey(rand.Reader)
	dhShared, err := privateDH.ECDH(publicDH)
	if err != nil {
		return nil, err
	}
	forwardShared := make([]byte, 64)
	copy(forwardShared, kemShared)
	copy(forwardShared[32:], dhShared)
	serverHello := append(kemCiphertext, privateDH.PublicKey().Bytes()...)
	conn.ENEXKXNftX = append(forwardShared, shared...)
	conn.AVDV5UPQ = WrnUSnC6(serverHello, conn.ENEXKXNftX, conn.HiUDG1Fl)
	conn.EufjD7UrtI = WrnUSnC6(clientHello[:1216], conn.ENEXKXNftX, conn.HiUDG1Fl)
	var ticket [16]byte
	_, _ = rand.Read(ticket[:])
	var lifetime int64
	if s.OIHUERfuz != 0 {
		lifetime = int64(IaLb4jaWtI(int(s.MOy_iE1vq), int(s.OIHUERfuz)))
	} else {
		lifetime = int64(IaLb4jaWtI(50, 100)) * s.MOy_iE1vq / 100
	}
	binary.BigEndian.PutUint16(ticket[:2], uint16(lifetime))
	if lifetime > 0 {
		s.VgRY0_5S9F37.Lock()
		s.GPXD9SMy[(time.Now().Unix()+max(s.MOy_iE1vq, s.OIHUERfuz))/60+2] = ticket
		s.CpNeohe = append(s.CpNeohe, ticket)
		s.GIHdzaFGhi8[ticket] = &MZKKvEA81HV{MwUsryYRwlpa: forwardShared}
		s.VgRY0_5S9F37.Unlock()
	}
	padding, sizes, delays := OODk_yfADN(s.EA3t4sK, s.PBHCqS)
	response := make([]byte, 1168+padding)
	handshakeCipher.Seal(response[:0], nonceMaximumAt509A1F0, serverHello, nil)
	conn.AVDV5UPQ.Seal(response[:1136], nil, ticket[:], nil)
	pad := response[1168:]
	var padLength [2]byte
	binary.BigEndian.PutUint16(padLength[:], uint16(padding-18))
	conn.AVDV5UPQ.Seal(pad[:0], nil, padLength[:], nil)
	conn.AVDV5UPQ.Seal(pad[:18], nil, pad[18:padding-16], nil)
	sizes[0] += 1168
	unsent := response
	for i, n := range sizes {
		if n > 0 {
			if _, err := raw.Write(unsent[:n]); err != nil {
				return nil, err
			}
			unsent = unsent[n:]
		}
		if i < len(delays) {
			time.Sleep(delays[i])
		}
	}
	if _, err := io.ReadFull(raw, header); err != nil {
		return nil, err
	}
	if _, err := handshakeCipher.Open(header[:0], nil, header, nil); err != nil {
		return nil, err
	}
	clientPadding := make([]byte, int(binary.BigEndian.Uint16(header[:2])))
	if _, err := io.ReadFull(raw, clientPadding); err != nil {
		return nil, err
	}
	if _, err := handshakeCipher.Open(clientPadding[:0], nil, clientPadding, nil); err != nil {
		return nil, err
	}
	if s.F9GbshEXgaiY == 2 {
		conn.DrpXao = &LUKUT5W0R6{DrpXao: raw,
			D1_SMs0WWKz:  GdvNyLM(conn.ENEXKXNftX, ticket[:]),
			EO5NXI6ex6:   GdvNyLM(conn.ENEXKXNftX, initial[:16]),
			DhRIe1pBzWdS: make([]byte, 0, 5), TfVMo2pTBnt: make([]byte, 0, 5)}
	}
	return conn, nil
}
