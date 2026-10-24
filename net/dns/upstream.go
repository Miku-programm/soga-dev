package dns

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"recovered.local/soga/util/logx"
	"strings"
	"time"
)

// cDpHJRFxZ: 0x1593360。
func cDpHJRFxZ(server string, localIP net.IP) Fr3ugVA_Lh4_ {
	scheme, address, found := strings.Cut(server, "://")
	if found {
		switch scheme {
		case "udp":
			return sLGJmZdZhG(address, localIP)
		case "tcp":
			return dNpGCa1ta_M(address, localIP)
		case "tls", "tcp-tls":
			return gXgZlI_Zv(address, localIP)
		case "https":
			return eXmNUd9(server, localIP)
		}
		logx.V8GN5Stk(slog.LevelError, "dns server <%v> 协议不支持，仅支持 udp:// tcp:// tcp-tls:// https://，已按 udp 处理", server)
		return sLGJmZdZhG(address, localIP)
	}
	if strings.HasPrefix(server, "local") {
		return &BkMuy3LHBq{}
	}
	return sLGJmZdZhG(server, localIP)
}

// 原构造函数重复的 SplitHostPort / JoinHostPort 分支；即使有错误也按 port 是否为空判断。
func serverAddress(server, defaultPort string) (string, string) {
	host, port, _ := net.SplitHostPort(server)
	if port != "" {
		return server, host
	}
	return net.JoinHostPort(server, defaultPort), server
}

// dNpGCa1ta_M: 0x1594aa0。
func dNpGCa1ta_M(server string, localIP net.IP) *KZFqTq8jd {
	address, host := serverAddress(server, "53")
	return &KZFqTq8jd{&JcYJA9vKYWNF{&dns.Client{Net: "tcp", UDPSize: 1024, Dialer: dsVT1Buuy(localIP, false, host), Timeout: 5 * time.Second}, address}}
}

// sLGJmZdZhG: 0x1594d20。Net 保持空字符串，由 DNS Client 使用 UDP。
func sLGJmZdZhG(server string, localIP net.IP) *Zq5YXn {
	address, host := serverAddress(server, "53")
	return &Zq5YXn{&JcYJA9vKYWNF{&dns.Client{UDPSize: 1024, Dialer: dsVT1Buuy(localIP, true, host), Timeout: 5 * time.Second}, address}}
}

// gXgZlI_Zv: 0x1591dc0。
func gXgZlI_Zv(server string, localIP net.IP) *OPLAviioh9T {
	address, host := serverAddress(server, "853")
	return &OPLAviioh9T{&JcYJA9vKYWNF{&dns.Client{Net: "tcp-tls", UDPSize: 1024, Dialer: dsVT1Buuy(localIP, false, host), Timeout: 5 * time.Second}, address}}
}

// dsVT1Buuy: 0x1594fa0。不匹配的地址族使用共享 dialer。
func dsVT1Buuy(localIP net.IP, udp bool, host string) *net.Dialer {
	if localIP == nil {
		return dialerAt5029CB0
	}
	if ip := net.ParseIP(host); ip != nil && (ip.To4() != nil) != (localIP.To4() != nil) {
		return dialerAt5029CB0
	}
	dialer := *dialerAt5029CB0
	if udp {
		dialer.LocalAddr = &net.UDPAddr{IP: localIP}
	} else {
		dialer.LocalAddr = &net.TCPAddr{IP: localIP}
	}
	return &dialer
}

// eXmNUd9: 0x1586fa0。
func eXmNUd9(server string, localIP net.IP) *NJKG0MHuXGIP {
	host := ""
	if u, err := url.Parse(server); err == nil {
		host = u.Hostname()
	}
	dialer := dsVT1Buuy(localIP, false, host)
	transport := &http.Transport{MaxIdleConns: 30, IdleConnTimeout: 90 * time.Second, TLSHandshakeTimeout: 30 * time.Second, ForceAttemptHTTP2: true, DialContext: dialer.DialContext}
	return &NJKG0MHuXGIP{&http.Client{Transport: transport, Timeout: 10 * time.Second}, server}
}

// zKnn59rsA1: 0x15856e0。
func (r *JcYJA9vKYWNF) zKnn59rsA1(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(host), typ)
	response, _, err := r.gRxcpc77sP.ExchangeContext(ctx, msg, r.dqrvhAOMdU6)
	if err != nil {
		return nil, err
	}
	if response.Rcode != 0 {
		return nil, errors.New(dns.RcodeToString[response.Rcode])
	}
	return answerIPs(response, host, typ)
}

// 两个原 lookup 共有的 Answer 类型分支（0x15858be、0x15878b4）。
func answerIPs(msg *dns.Msg, host string, typ uint16) ([]net.IP, error) {
	ips := make([]net.IP, 0)
	for _, record := range msg.Answer {
		switch record := record.(type) {
		case *dns.A:
			if typ == 1 {
				ips = append(ips, record.A)
			}
		case *dns.AAAA:
			if typ == 28 {
				ips = append(ips, record.AAAA)
			}
		}
	}
	if len(ips) == 0 {
		return nil, fmt.Errorf("can not resolve host: %v", host)
	}
	return ips, nil
}

// LookupDNSType: 0x1585c60。
func (r *JcYJA9vKYWNF) LookupDNSType(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	return r.zKnn59rsA1(ctx, host, typ)
}

// zKnn59rsA1: 0x15871a0。原实现未验证 HTTP status 或 DNS Rcode。
func (r *NJKG0MHuXGIP) zKnn59rsA1(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(host), typ)
	payload, err := msg.PackBuffer(nil)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", r.jZWWEh, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/dns-message")
	req.Header.Add("Content-Type", "application/dns-message")
	resp, err := r.s9Guhvp0K.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	payload, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = msg.Unpack(payload); err != nil {
		return nil, err
	}
	return answerIPs(msg, host, typ)
}

// LookupDNSType: 0x1587d60。
func (r *NJKG0MHuXGIP) LookupDNSType(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	return r.zKnn59rsA1(ctx, host, typ)
}
