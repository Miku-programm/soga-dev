package mXe7cPU8Mzi

import (
	"crypto/cipher"
	"fmt"
	"io"
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	G "recovered.local/soga/G6bHD5JA"
)

// Original RTTI 0x3659278 (56 bytes) and 0x3631bb0 (32 bytes).
type rZDq_kinV4G struct {
	aMKXGpHBfg  net.Conn
	n49jm3JWYo  cipher.AEAD
	xUtUMCF3grV []byte
}
type jjyc26 struct {
	jaPY7RU   *IqB1foHWg
	mzrzt1uBj []byte
}

func (r *rZDq_kinV4G) Read() (*jjyc26, error) { //0x1c4cec0
	encrypted := make([]byte, 48)
	if _, err := io.ReadFull(r.aMKXGpHBfg, encrypted); err != nil {
		return nil, fmt.Errorf("read metadata: %w", err)
	}
	incrementNonce(r.xUtUMCF3grV)
	plain, err := r.n49jm3JWYo.Open(nil, r.xUtUMCF3grV, encrypted, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypt metadata: %w", err)
	}
	m, err := Ox0UfI5kIfcR(plain)
	if err != nil {
		return nil, fmt.Errorf("parse metadata: %w", err)
	}
	return r.jQ8yjlkt(m)
}

func (r *rZDq_kinV4G) jQ8yjlkt(m *IqB1foHWg) (*jjyc26, error) { //0x1c4d360
	kind := m.Rf6N7kY6GVCs | 1
	if (kind == 7 || kind == 9) && m.B8QsA0N2yo != 0 {
		if err := G.KlTAqrF42mR(r.aMKXGpHBfg, int(m.B8QsA0N2yo)); err != nil {
			return nil, fmt.Errorf("discard prefix padding: %w", err)
		}
	}
	var payload []byte
	if m.HPUJVQ0S != 0 {
		encrypted := D.PoolAt50961D8.Get(int(m.HPUJVQ0S) + 16)
		if _, err := io.ReadFull(r.aMKXGpHBfg, encrypted); err != nil {
			D.PoolAt50961D8.Put(encrypted)
			return nil, fmt.Errorf("read payload: %w", err)
		}
		incrementNonce(r.xUtUMCF3grV)
		dest := D.PoolAt50961D8.Get(int(m.HPUJVQ0S))
		var err error
		payload, err = r.n49jm3JWYo.Open(dest[:0], r.xUtUMCF3grV, encrypted, nil)
		D.PoolAt50961D8.Put(encrypted)
		if err != nil {
			D.PoolAt50961D8.Put(dest)
			return nil, fmt.Errorf("decrypt payload: %w", err)
		}
	}
	if m.FO7MEJ != 0 {
		// The original does not return an already decrypted payload to the pool on
		// this failure branch. Do not introduce a different ownership side effect.
		if err := G.KlTAqrF42mR(r.aMKXGpHBfg, int(m.FO7MEJ)); err != nil {
			return nil, fmt.Errorf("discard suffix padding: %w", err)
		}
	}
	return &jjyc26{jaPY7RU: m, mzrzt1uBj: payload}, nil
}
