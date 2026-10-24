package FLhkKwJ8M

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"math/rand"
	"net"
)

type GC3qACg18wgq struct {
	roEClw      []*OMOYhoLsaabd
	gyJC3dk_8   *uuid.UUID
	dKNGuf      uint8
	kOlsNE5LaUi bool
}

// ClientConfiguration is an analyst-named representation of the 80-byte value
// passed to 0x19ba4c0; the original type name is absent from RTTI. Field offsets
// follow caller 0x1b6e7e0. Port and Host are passed but never read by this function.
type ClientConfiguration struct {
	UUID     string
	AlterID  uint16
	Security string
	Port     string
	Host     string
	AEAD     bool
}

// Iw0IEMM3ha: 0x19ba4c0.
func Iw0IEMM3ha(config ClientConfiguration) (*GC3qACg18wgq, error) {
	id := new(uuid.UUID)
	if err := id.Parse(config.UUID); err != nil {
		return nil, err
	}
	var security byte
	switch config.Security {
	case "aes-128-gcm", "auto":
		security = 3
	case "chacha20-poly1305":
		security = 4
	case "none":
		security = 5
	case "zero":
		security = 6
	default:
		return nil, fmt.Errorf("unknown security type: %s", config.Security)
	}
	return &GC3qACg18wgq{roEClw: o_nBA2(h7OyM4GsmHif(id), config.AlterID), gyJC3dk_8: id, dKNGuf: security, kOlsNE5LaUi: config.AEAD}, nil
}

// StreamConn: 0x19ba400. On failure this returns a typed nil *HRzHVG1yP5O
// inside net.Conn, matching the unconditional original interface construction.
func (c *GC3qACg18wgq) StreamConn(conn net.Conn, dst *IxhMZ66CqoH1) (net.Conn, error) {
	result, err := wv3qc3(conn, c.roEClw[rand.Intn(len(c.roEClw))], dst, c.dKNGuf, c.kOlsNE5LaUi)
	return result, err
}
