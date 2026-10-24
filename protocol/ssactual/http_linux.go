//go:build linux

package ssactual

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	E "recovered.local/soga/net/outbound"
	O "recovered.local/soga/net/socket"
	"strconv"
	"time"
)

// DlqIvv retains all original fields and the 56-byte RTTI layout.
type DlqIvv struct {
	*B8EHsVY7a
	l_ViT2M   string
	bqekXVE   string
	savFEH    *tls.Config
	DvfjZBhXc http.Header
}

// HTTPOptions is an analysis name: the original by-value parameter's name
// is not retained. Its 136-byte layout follows 0x1b58340's argument offsets.
type HTTPOptions struct {
	FFrN8gdu
	Name           string
	Server         string
	Port           int
	Username       string
	Password       string
	TLS            bool
	SNI            string
	SkipCertVerify bool
	Headers        map[string]string
}

// CJW16StU3I: 0x1b58340. Header values are added, including values whose
// differently cased input keys canonicalize to the same HTTP header.
func CJW16StU3I(o HTTPOptions) *DlqIvv {
	var config *tls.Config
	if o.TLS {
		name := o.Server
		if o.SNI != "" {
			name = o.SNI
		}
		config = &tls.Config{ServerName: name, InsecureSkipVerify: o.SkipCertVerify}
	}
	headers := make(http.Header)
	for key, value := range o.Headers {
		headers.Add(key, value)
	}
	return &DlqIvv{
		B8EHsVY7a: &B8EHsVY7a{m50I26ZDL: o.Name, x3HUaU8Ln1S1: net.JoinHostPort(o.Server, strconv.Itoa(o.Port)), xDan8hA5LV: o.PiTAGr6U5KDG, iOKxDl: 6, bEbqkM8: o.VsBRTR},
		l_ViT2M:   o.Username, bqekXVE: o.Password, savFEH: config, DvfjZBhXc: headers,
	}
}

// g9TEl8UN: 0x1b57680. The temporary response reader is discarded and the
// response body is not closed here, exactly as in the original function.
func (h *DlqIvv) g9TEl8UN(m *E.D0mLgQF, conn io.ReadWriter) error {
	address := m.RemoteAddress()
	request := &http.Request{Method: "CONNECT", URL: &url.URL{Host: address}, Host: address, Header: h.DvfjZBhXc.Clone()}
	request.Header.Add("Proxy-Connection", "Keep-Alive")
	if h.l_ViT2M != "" && h.bqekXVE != "" {
		request.Header.Add("Proxy-Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(h.l_ViT2M+":"+h.bqekXVE)))
	}
	if err := request.Write(conn); err != nil {
		return err
	}
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		return err
	}
	switch response.StatusCode {
	case 200:
		return nil
	case 407:
		return errors.New("HTTP need auth")
	case 405:
		return errors.New("CONNECT method not allowed by proxy")
	default:
		if response.StatusCode >= 500 {
			return errors.New(response.Status)
		}
		return fmt.Errorf("can not connect remote err code: %d", response.StatusCode)
	}
}

// StreamConn: 0x1b56ae0. Only the TLS handshake receives a five-second
// background timeout; the CONNECT exchange has no additional deadline.
func (h *DlqIvv) StreamConn(conn net.Conn, m *E.D0mLgQF) (net.Conn, error) {
	if h.savFEH != nil {
		secured := tls.Client(conn, h.savFEH)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := secured.HandshakeContext(ctx); err != nil {
			return nil, fmt.Errorf("%s connect error: %w", h.Addr(), err)
		}
		conn = secured
	}
	if err := h.g9TEl8UN(m, conn); err != nil {
		return nil, err
	}
	return conn, nil
}

// DialContext: 0x1b56f80; failure cleanup closure: 0x1b72c40.
func (h *DlqIvv) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.PnyceLR, err error) {
	conn, err := O.I_y9nlK1q8(ctx, "tcp", h.Addr(), h.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", h.Addr(), err)
	}
	mS0F0Ti(conn)
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	secured, err := h.StreamConn(conn, m)
	if err != nil {
		return nil, err
	}
	return &lGta_1QkvQ{Conn: secured, bnOtGuS: E.GYe3KD{h.Name()}}, nil
}

var _ E.C74CUtQ9FqH = (*DlqIvv)(nil)
