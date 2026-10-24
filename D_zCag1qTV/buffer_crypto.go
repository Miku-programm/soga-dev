package D_zCag1qTV

import "crypto/cipher"

// OpenWrite @ 0xe98340。原实现忽略 AEAD 返回切片，按预计算的长度推进。
func (b *X_aXX5) OpenWrite(aead cipher.AEAD, nonce, ciphertext, additional []byte) (int, error) {
	size := len(ciphertext) - aead.Overhead()
	b.GrowBy(size)
	_, err := aead.Open(b.kNItjuFiwnR[b.__XpzgdOCbm:b.__XpzgdOCbm], nonce, ciphertext, additional)
	if err != nil {
		return 0, err
	}
	b.__XpzgdOCbm += size
	return size, nil
}

// Open @ 0xe98520。只在认证成功后读取 Overhead 并调整写位置。
func (b *X_aXX5) Open(aead cipher.AEAD, nonce, additional []byte) (int, error) {
	_, err := aead.Open(b.kNItjuFiwnR[b.cnahaWq6:b.cnahaWq6], nonce, b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm], additional)
	if err != nil {
		return 0, err
	}
	b.__XpzgdOCbm -= aead.Overhead()
	return b.__XpzgdOCbm - b.cnahaWq6, nil
}

// Seal @ 0xe986c0。前后两次 Overhead 调用均保留。
func (b *X_aXX5) Seal(aead cipher.AEAD, nonce, additional []byte) {
	b.GrowBy(aead.Overhead())
	aead.Seal(b.kNItjuFiwnR[b.cnahaWq6:b.cnahaWq6], nonce, b.kNItjuFiwnR[b.cnahaWq6:b.__XpzgdOCbm], additional)
	b.__XpzgdOCbm += aead.Overhead()
}

// SealWrite @ 0xe98880。
func (b *X_aXX5) SealWrite(aead cipher.AEAD, nonce, plaintext, additional []byte) int {
	size := aead.Overhead() + len(plaintext)
	b.GrowBy(size)
	aead.Seal(b.kNItjuFiwnR[b.__XpzgdOCbm:b.__XpzgdOCbm], nonce, plaintext, additional)
	b.__XpzgdOCbm += size
	return size
}

// SealWriteWithStart @ 0xe98a40。start 是底层切片绝对位置，压紧后也不改写。
func (b *X_aXX5) SealWriteWithStart(aead cipher.AEAD, start int, nonce, plaintext, additional []byte) int {
	_, _ = b.Write(plaintext)
	b.GrowBy(aead.Overhead())
	aead.Seal(b.kNItjuFiwnR[start:start], nonce, b.kNItjuFiwnR[start:b.__XpzgdOCbm], additional)
	b.__XpzgdOCbm += aead.Overhead()
	return len(plaintext) + aead.Overhead()
}

func (b *X_aXX5) XORKeyStreamWrite(stream cipher.Stream, plaintext []byte) int { // 0xe98c40
	b.GrowBy(len(plaintext))
	stream.XORKeyStream(b.kNItjuFiwnR[b.__XpzgdOCbm:], plaintext)
	b.__XpzgdOCbm += len(plaintext)
	return len(plaintext)
}

func (b *X_aXX5) XORKeyStreamRead(stream cipher.Stream, data []byte) (int, error) { // 0xe98d40
	n, err := b.Read(data)
	if err != nil {
		return 0, err
	}
	stream.XORKeyStream(data, data[:n])
	return n, nil
}
