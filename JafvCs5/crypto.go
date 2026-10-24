package JafvCs5

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

// E_IV2PKO 对应 0x181b7a0。五个切片的顺序经 Go 寄存器/栈 ABI 核实。
// 保留原 Open 的 dst 追加语义及丢弃返回切片的行为。
func E_IV2PKO(key, nonce, dst, ciphertext, additional []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	_, err = aead.Open(dst, nonce, ciphertext, additional)
	return err
}

type FIo_wn2vRGa struct {
	_QLr1_kE  cipher.Stream
	m7MKueoGA io.Reader
}

// NewStreamReader 是用于表达原调用者内联结构赋值的恢复桥接，并非原函数名。
func NewStreamReader(stream cipher.Stream, reader io.Reader) *FIo_wn2vRGa {
	return &FIo_wn2vRGa{_QLr1_kE: stream, m7MKueoGA: reader}
}
func (r *FIo_wn2vRGa) Read(p []byte) (int, error) { // 0x181c0c0
	n, err := r.m7MKueoGA.Read(p)
	if n > 0 {
		r._QLr1_kE.XORKeyStream(p[:n], p[:n])
	}
	return n, err
}

// Fc9Ososz7 对应 0x181b620，保留 Seal 追加写入而丢弃返回切片的语义。
func Fc9Ososz7(key, nonce, dst, plaintext, additional []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	aead.Seal(dst, nonce, plaintext, additional)
	return nil
}
