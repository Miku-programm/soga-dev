package ss2022

import (
	"bytes"
	"context"
	"encoding/binary"
	"net"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
	"testing"
	"time"
)

func clientPacket(t *testing.T, m *N07NaA, user *user.WtK9HQSy66, id, sequence uint64, payload []byte) []byte {
	t.Helper()
	header := make([]byte, 16)
	binary.BigEndian.PutUint64(header, id)
	binary.BigEndian.PutUint64(header[8:], sequence)
	hash, ok := m.GetHashById(user.GetId())
	if !ok {
		t.Fatal("no hash")
	}
	var identity [16]byte
	for i := range identity {
		identity[i] = header[i] ^ hash[i]
	}
	m.su5Rzwte.Encrypt(identity[:], identity[:])
	plaintext := make([]byte, 11)
	binary.BigEndian.PutUint64(plaintext[1:], uint64(rtx.C5Jff5vs0()))
	plaintext = append(plaintext, 1, 192, 0, 2, 1, 0, 53)
	plaintext = append(plaintext, payload...)
	a, err := user.GetCipher().NewAEAD(user.GetCipher().SubKey(user.PSK(), header[:8]))
	if err != nil {
		t.Fatal(err)
	}
	sealed := a.Seal(nil, header[4:16], plaintext, nil)
	m.su5Rzwte.Encrypt(header, header)
	return append(append(header, identity[:]...), sealed...)
}
func TestNativeUDPAuthenticationReplayAndResponse(t *testing.T) {
	for _, method := range []ssmethod.C6q3Upu{"2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		m := managerFixture(t, ctx, method)
		user := userFixture(t, method, 23, 17)
		m.UpdateUser(user)
		server, err := net.ListenPacket("udp4", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer server.Close()
		client, err := net.ListenPacket("udp4", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()
		server.SetDeadline(time.Now().Add(3 * time.Second))
		client.SetDeadline(time.Now().Add(3 * time.Second))
		wrapped := NewPacketConnForFactory(server, m)
		reader := wrapped.FcUA_X7yd
		writer := wrapped.JxxVAb
		packet := clientPacket(t, m, user, 234, 0, []byte("request"))
		if _, err = client.WriteTo(packet, server.LocalAddr()); err != nil {
			t.Fatal(err)
		}
		buf := make([]byte, 4096)
		n, gotUser, source, target, err := reader.ReadFromProxy(buf)
		if err != nil || gotUser != user || string(buf[:n]) != "request" || target.GetHost() != "192.0.2.1" || target.GetPort() != 53 {
			t.Fatalf("decode %v", err)
		}
		if n, err = writer.WriteToProxy([]byte("reply"), user, source, &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 53}); err != nil || n != 5 {
			t.Fatal(err)
		}
		n, _, err = client.ReadFrom(buf)
		if err != nil {
			t.Fatal(err)
		}
		user.GetBlock().Decrypt(buf[:16], buf[:16])
		sid := binary.BigEndian.Uint64(buf[:8])
		if binary.BigEndian.Uint64(buf[8:16]) != 0 {
			t.Fatal("initial packet id")
		}
		a, err := method.NewAEAD(method.SubKey(user.PSK(), buf[:8]))
		if err != nil {
			t.Fatal(err)
		}
		plaintext, err := a.Open(nil, buf[4:16], buf[16:n], nil)
		if err != nil {
			t.Fatal(err)
		}
		if plaintext[0] != 1 || binary.BigEndian.Uint64(plaintext[9:17]) != 234 {
			t.Fatal("response IDs")
		}
		pad := int(binary.BigEndian.Uint16(plaintext[17:19]))
		if pad < 1 || pad > 895 {
			t.Fatal("padding", pad)
		}
		if !bytes.Equal(plaintext[19+pad:], append([]byte{1, 192, 0, 2, 1, 0, 53}, []byte("reply")...)) {
			t.Fatal("response payload")
		}
		s, ok := m.GetSessionByClientAddr(source)
		if !ok || s.fl_h6QOtKast != sid {
			t.Fatal("session")
		}
		// 重放及认证失败不返回应用数据，也不阻止下一条有效报文。
		client.WriteTo(packet, server.LocalAddr())
		corrupt := clientPacket(t, m, user, 234, 1, []byte("bad"))
		corrupt[len(corrupt)-1] ^= 1
		client.WriteTo(corrupt, server.LocalAddr())
		next := clientPacket(t, m, user, 234, 1, []byte("next"))
		client.WriteTo(next, server.LocalAddr())
		n, gotUser, _, _, err = reader.ReadFromProxy(buf)
		if err != nil || gotUser != user || string(buf[:n]) != "next" {
			t.Fatal("skip damaged/replayed", err)
		}
		if !s.paeyZKdn.Check(2) || s.paeyZKdn.Check(1) {
			t.Fatal("replay window commit")
		}
	}
}

var _ sess.KsKg3yls = (*KNAU41ESnn)(nil)
