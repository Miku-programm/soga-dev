package vmessaead

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ELIQWabl31 is the original 96-byte WebSocket configuration structure.
type ELIQWabl31 struct {
	OhyeDUGTCGH string
	RTIB8Nkai   string
	WOT5UbYYto  string
	RlZC_oDxj   http.Header
	Xtp2zWRK    bool
	VbB82ra     *tls.Config
	A3jvUyhB    int
	I_3sggK     string
}

// ki0pcfT's original 56-byte layout. The hashed dependency's Dialer/Conn
// RTTI matches gorilla/websocket v1.5.3; the original exact version is unproven.
type ki0pcfT struct {
	bcfUf2fzTgX  *websocket.Conn
	jRVoqU9TmTn  io.Reader
	fuo5hf0ygZcx net.Addr
	aCoNgxuV     sync.Mutex
	aHBpwQ       sync.Mutex
}

// tlVyZZk: 0x19bc700. Message type is intentionally ignored.
func (c *ki0pcfT) tlVyZZk() (io.Reader, error) {
	if c.jRVoqU9TmTn != nil {
		return c.jRVoqU9TmTn, nil
	}
	_, r, err := c.bcfUf2fzTgX.NextReader()
	if err != nil {
		return nil, err
	}
	c.jRVoqU9TmTn = r
	return r, nil
}

// Read: 0x19bc040. EOF advances to the next WebSocket message.
func (c *ki0pcfT) Read(p []byte) (int, error) {
	c.aCoNgxuV.Lock()
	defer c.aCoNgxuV.Unlock()
	for {
		r, err := c.tlVyZZk()
		if err != nil {
			return 0, err
		}
		n, err := r.Read(p)
		if err != io.EOF {
			return n, err
		}
		c.jRVoqU9TmTn = nil
	}
}

// Write: 0x19bc260; sends one binary frame per call under its own mutex.
func (c *ki0pcfT) Write(p []byte) (int, error) {
	c.aHBpwQ.Lock()
	defer c.aHBpwQ.Unlock()
	if err := c.bcfUf2fzTgX.WriteMessage(websocket.BinaryMessage, p); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Close: 0x19bc400; send status 1000 with a five-second deadline, then close
// even on failure. Errors are joined as text, not with errors.Join.
func (c *ki0pcfT) Close() error {
	var failures []string
	if err := c.bcfUf2fzTgX.WriteControl(websocket.CloseMessage, []byte{3, 232}, time.Now().Add(5*time.Second)); err != nil {
		failures = append(failures, err.Error())
	}
	if err := c.bcfUf2fzTgX.Close(); err != nil {
		failures = append(failures, err.Error())
	}
	if len(failures) != 0 {
		return fmt.Errorf("failed to close connection: %s", strings.Join(failures, ","))
	}
	return nil
}

// LocalAddr: 0x19bc7a0; RemoteAddr: 0x19bc7e0.
func (c *ki0pcfT) LocalAddr() net.Addr  { return c.bcfUf2fzTgX.LocalAddr() }
func (c *ki0pcfT) RemoteAddr() net.Addr { return c.fuo5hf0ygZcx }

// Deadline methods: 0x19bc800/0x19bc8e0/0x19bc940.
func (c *ki0pcfT) SetDeadline(t time.Time) error {
	if err := c.SetReadDeadline(t); err != nil {
		return err
	}
	return c.SetWriteDeadline(t)
}
func (c *ki0pcfT) SetReadDeadline(t time.Time) error  { return c.bcfUf2fzTgX.SetReadDeadline(t) }
func (c *ki0pcfT) SetWriteDeadline(t time.Time) error { return c.bcfUf2fzTgX.SetWriteDeadline(t) }

// caPKDWZ1qaY9: 0x19bd500. NetDial returns the already-open connection;
// headers are rebuilt using only Get's first value for each original key.
func caPKDWZ1qaY9(conn net.Conn, config *ELIQWabl31, early *bytes.Buffer) (net.Conn, error) {
	dialer := &websocket.Dialer{NetDial: func(string, string) (net.Conn, error) { return conn, nil }, ReadBufferSize: 4096, WriteBufferSize: 4096, HandshakeTimeout: 8 * time.Second}
	scheme := "ws"
	if config.Xtp2zWRK {
		scheme = "wss"
		dialer.TLSClientConfig = config.VbB82ra
	}
	path, err := url.Parse(config.WOT5UbYYto)
	if err != nil {
		return nil, fmt.Errorf("parse url %s error: %w", config.WOT5UbYYto, err)
	}
	target := url.URL{Scheme: scheme, Host: net.JoinHostPort(config.OhyeDUGTCGH, config.RTIB8Nkai), Path: path.Path, RawQuery: path.RawQuery}
	headers := make(http.Header)
	for key := range config.RlZC_oDxj {
		headers.Add(key, config.RlZC_oDxj.Get(key))
	}
	if early != nil {
		if config.I_3sggK != "" {
			headers.Set(config.I_3sggK, early.String())
		} else {
			target.Path += early.String()
		}
	}
	ws, response, err := dialer.DialContext(context.Background(), target.String(), headers)
	if err != nil {
		message := err.Error()
		if response != nil {
			message = response.Status
		}
		return nil, fmt.Errorf("dial %s error: %s", target.Host, message)
	}
	return &ki0pcfT{bcfUf2fzTgX: ws, fuo5hf0ygZcx: conn.RemoteAddr()}, nil
}

// E6YB0lYJp: 0x19bdfc0. A valid ed query parameter overrides early-data
// settings and is removed from the retained configuration path.
func E6YB0lYJp(conn net.Conn, config *ELIQWabl31) (net.Conn, error) {
	if path, err := url.Parse(config.WOT5UbYYto); err == nil {
		query := path.Query()
		if value := query.Get("ed"); value != "" {
			if n, err := strconv.Atoi(value); err == nil {
				config.A3jvUyhB = n
				config.I_3sggK = "Sec-WebSocket-Protocol"
				query.Del("ed")
				path.RawQuery = query.Encode()
				config.WOT5UbYYto = path.String()
			}
		}
	}
	if config.A3jvUyhB > 0 {
		return uTCy4V(conn, config)
	}
	return caPKDWZ1qaY9(conn, config, nil)
}

var _ net.Conn = (*ki0pcfT)(nil)
