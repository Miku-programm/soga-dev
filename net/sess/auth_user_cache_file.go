package sess

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"os"
	"time"

	"github.com/jellydator/ttlcache/v3"
	B "recovered.local/soga/buf/pool"
	C "recovered.local/soga/config"
	U "recovered.local/soga/user"
)

// 原 RTTI 中的缓存文件记录：16 字节用户项和 48 字节 IP 项。
type YHksCjPM struct {
	W28KAG2qD7aT int64
	CoHdPBk_a    int64
}
type CXTczIZ5hN7o []*YHksCjPM
type B07Fb7qv struct {
	SRR4_uk   net.IP
	ZFFCVgG8d CXTczIZ5hN7o
}

// LAQrKU @0x187c820；记录为地址类型、IP、32 位数量和成对 64 位整数，均为大端。
func LAQrKU(b *bytes.Buffer) (*B07Fb7qv, error) {
	kind, err := b.ReadByte()
	if err != nil {
		return nil, err
	}
	var size int
	switch kind {
	case 1:
		size = 4
	case 2:
		size = 16
	default:
		return nil, errors.New("unknown ip type")
	}
	ip := b.Next(size)
	if len(ip) != size {
		return nil, io.ErrUnexpectedEOF
	}
	n := b.Next(4)
	if len(n) != 4 {
		return nil, io.ErrUnexpectedEOF
	}
	count := binary.BigEndian.Uint32(n)
	users := make(CXTczIZ5hN7o, 0, int(count))
	for i := uint32(0); i < count; i++ {
		id := b.Next(8)
		// 原汇编 0x187cc76/0x187cc5e 在这两条短记录分支返回保存的 nil error。
		// 保留原行为；Load 随后的解引用因此仍可能 panic，不能擅自改成兼容实现。
		if len(id) != 8 {
			return nil, nil
		}
		expires := b.Next(8)
		if len(expires) != 8 {
			return nil, nil
		}
		users = append(users, &YHksCjPM{int64(binary.BigEndian.Uint64(id)), int64(binary.BigEndian.Uint64(expires))})
	}
	return &B07Fb7qv{ip, users}, nil
}

func (r *B07Fb7qv) Write(b *B.X_aXX5) { // 0x187cd80
	if ip := r.SRR4_uk.To4(); ip != nil {
		b.WriteByte(1)
		b.Write(ip)
	} else {
		b.WriteByte(2)
		b.Write(r.SRR4_uk)
	}
	var n [8]byte
	binary.BigEndian.PutUint32(n[:4], uint32(len(r.ZFFCVgG8d)))
	b.Write(n[:4])
	for _, u := range r.ZFFCVgG8d {
		binary.BigEndian.PutUint64(n[:], uint64(u.W28KAG2qD7aT))
		b.Write(n[:])
		binary.BigEndian.PutUint64(n[:], uint64(u.CoHdPBk_a))
		b.Write(n[:])
	}
}

func (c *EafxyUZD) Load(path string, users *U.ZpNFqUVZp) (bool, error) { // 0x187d6c0
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	b := bytes.NewBuffer(data)
	now := time.Now().UnixMilli()
	for {
		r, err := LAQrKU(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return true, nil
			}
			return false, err
		}
		ip := r.SRR4_uk.String()
		for _, entry := range r.ZFFCVgG8d {
			if u, ok := users.Get(entry.W28KAG2qD7aT); ok && entry.CoHdPBk_a > now {
				c.AddUser(ip, u, entry.CoHdPBk_a)
			}
		}
	}
}

func (c *EafxyUZD) Save(path string) error { // 0x187d960
	c.aRCMnhMrKi.RLock()
	records := make([]*B07Fb7qv, 0, c.eGhvKTJIaDi.Len())
	c.eGhvKTJIaDi.Range(func(item *ttlcache.Item[string, *Elkpyuoj]) bool {
		ip := net.ParseIP(item.Key())
		if ip == nil {
			return true
		}
		l := item.Value()
		r := &B07Fb7qv{ip, make(CXTczIZ5hN7o, 0, l.Count())}
		l.Range(func(v *HRRGg0nld) bool {
			r.ZFFCVgG8d = append(r.ZFFCVgG8d, &YHksCjPM{v.mHY8JyIJ6.GetId(), v.jMdUpB})
			return true
		}, false)
		records = append(records, r)
		return true
	})
	c.aRCMnhMrKi.RUnlock()
	// 原代码使用 0644 创建目录且忽略该错误；不在重建时修改这一行为。
	_ = os.MkdirAll(C.IPUserCacheDirectory(), 0644)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	b := B.S4OQXWIZ(16384)
	defer b.Free()
	for _, r := range records {
		r.Write(b)
		if b.Len() >= 15360 {
			if _, err = f.Write(b.Bytes()); err != nil {
				return err
			}
			b.Reset()
		}
	}
	if b.Len() > 0 {
		_, err = f.Write(b.Bytes())
	}
	return err
}
