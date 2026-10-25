package relay

import "testing"

// 直连 IP 后被 sniff 改写为域名时，AccessHost 仍是客户端原始请求的 IP。
func TestAccessHostKeepsTargetBeforeSniff(t *testing.T) {
	direct, err := U80FDt("31.13.92.37:443")
	if err != nil {
		t.Fatal(err)
	}
	direct.SetHost("www.facebook.com")
	if direct.GetHost() != "www.facebook.com" || direct.AccessHost() != "31.13.92.37" {
		t.Fatalf("direct ip: host=%v access=%v", direct.GetHost(), direct.AccessHost())
	}
	direct.SetHost("other.example")
	if direct.AccessHost() != "31.13.92.37" {
		t.Fatal("second SetHost must keep the first original target")
	}

	byDomain := NewDomainAddress("facebook.com", 443)
	byDomain.SetHost("www.facebook.com")
	if byDomain.AccessHost() != "facebook.com" {
		t.Fatal(byDomain.AccessHost())
	}

	untouched := NewDomainAddress("www.facebook.com", 443)
	if untouched.AccessHost() != "www.facebook.com" {
		t.Fatal(untouched.AccessHost())
	}

	// sniff_redirect 新建地址后保留原始目标。
	redirected := NewDomainAddress("www.facebook.com", 443)
	redirected.SetAccessHost(direct.AccessHost())
	if redirected.AccessHost() != "31.13.92.37" || redirected.GetHost() != "www.facebook.com" {
		t.Fatal(redirected.AccessHost())
	}
}
