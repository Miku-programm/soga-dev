package vmess

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"net"
	Replay "recovered.local/soga/crypto/replaywindow"
	B "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	K "recovered.local/soga/util/ssrhelper"
	U "recovered.local/soga/user"
)

var replayErrorAt5098420 = errors.New("vmess salt repeat") // init 0x1d12f40

type KUJQ1W struct {
	*DVEkfear
	d2_4RdfaUPn  *WR4lMNfab
	aMP1QMakFbpY *U.QKaIWV2
	kpF6Hxsr8E   *Wpqfspy6zZ_B
}

func EAQVlR(conn net.Conn, auth *WR4lMNfab) (*KUJQ1W, error) { // 0x1d178a0
	framed, header, err := uaQJpenJx8c(conn, auth)
	if err != nil {
		H.Li3c4CJ1_D(conn)
		return nil, err
	}
	return &KUJQ1W{framed, auth, auth.Xwwpafq, header}, nil
}
func uaQJpenJx8c(conn net.Conn, auth *WR4lMNfab) (*DVEkfear, *Wpqfspy6zZ_B, error) { // 0x1d179c0
	var reader io.Reader
	if auth.Ih7r8HyMvOa {
		buf := B.LdHEmSs9(128)
		defer buf.Free()
		if _, err := buf.ReadFullFrom(conn, 26); err != nil {
			return nil, nil, err
		}
		lengthBlock := buf.BytesTo(18)
		connectionNonce := append([]byte(nil), buf.BytesFrom(18)...)
		commandKey := auth.GetCmdKey()
		key := K.CVqKigbE26(commandKey, []byte("VMess Header AEAD Key_Length"), auth.T_ul7Q_0j, connectionNonce)[:16]
		nonce := K.CVqKigbE26(commandKey, []byte("VMess Header AEAD Nonce_Length"), auth.T_ul7Q_0j, connectionNonce)[:12]
		if err := K.E_IV2PKO(key, nonce, lengthBlock[:0], lengthBlock, auth.T_ul7Q_0j); err != nil {
			return nil, nil, err
		}
		length := int(binary.BigEndian.Uint16(lengthBlock))
		buf.Reset()
		if _, err := buf.ReadFullFrom(conn, length+16); err != nil {
			return nil, nil, err
		}
		key = K.CVqKigbE26(commandKey, []byte("VMess Header AEAD Key"), auth.T_ul7Q_0j, connectionNonce)[:16]
		nonce = K.CVqKigbE26(commandKey, []byte("VMess Header AEAD Nonce"), auth.T_ul7Q_0j, connectionNonce)[:12]
		data := buf.Bytes()
		if err := K.E_IV2PKO(key, nonce, data[:0], data, auth.T_ul7Q_0j); err != nil {
			return nil, nil, err
		}
		reader = bytes.NewReader(data[:length])
	} else {
		block, err := aes.NewCipher(auth.GetCmdKey())
		if err != nil {
			return nil, nil, err
		}
		reader = K.NewStreamReader(cipher.NewCFBDecrypter(block, auth.GetIV()), conn)
	}
	header, err := Gk1af7wDMnD(reader)
	if err != nil {
		return nil, nil, err
	}
	replay := B.LdHEmSs9(48)
	defer replay.Free()
	replay.Write(auth.B2pr6gMaPO[:])
	replay.Write(header.cEcPFEa7)
	replay.Write(header.ezCHRvHCePBC)
	if !Replay.HFlnXWty(replay.Bytes()) {
		return nil, nil, replayErrorAt5098420
	}
	framed, err := W2tVTGQ(conn, auth, header)
	if err != nil {
		return nil, nil, err
	}
	return framed, header, nil
}
func (c *KUJQ1W) GetUser() U.PgaNd3014U       { return c.aMP1QMakFbpY }             // 0x1d18a60
func (c *KUJQ1W) GetDestAddr() *G.AHU2MRmjzQu { return c.kpF6Hxsr8E.GetDestAddr() } // 0x1d18a80
func (c *KUJQ1W) NetworkType() J.ZES6F5       { return c.kpF6Hxsr8E.NetworkType() } // 0x1d18aa0
func (c *KUJQ1W) Free()                       { H.Li3c4CJ1_D(c.DVEkfear) }          // 0x1d18ae0
