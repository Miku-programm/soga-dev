//go:build linux

// 根据 IDA 0x149fc00、0x14a0960、0x14a1aa0、0x14a1d60、0x153ae00、0x153af20 恢复。
package a3PdV_bmpRUK

import (
	"context"
	"crypto"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
)

// s_CjqL_8 原 RTTI 大小 24 字节。
type s_CjqL_8 struct {
	zlC6WUcq    crypto.PrivateKey
	w0zKHy_Ex0W *registration.Resource
}

func (u *s_CjqL_8) GetEmail() string                        { return "" }            // 0x153ada0
func (u *s_CjqL_8) GetRegistration() *registration.Resource { return u.w0zKHy_Ex0W } // 0x153adc0
func (u *s_CjqL_8) GetPrivateKey() crypto.PrivateKey        { return u.zlC6WUcq }    // 0x153ade0

func rLtaOfBZh() (crypto.PrivateKey, error) {
	const path = "/etc/soga/acme_account.key"
	contents, err := os.ReadFile(path)
	if err == nil {
		return certcrypto.ParsePEMPrivateKey(contents)
	}
	key, err := certcrypto.GeneratePrivateKey(certcrypto.EC256)
	if err != nil {
		return nil, err
	}
	if err = os.WriteFile(path, certcrypto.PEMEncode(key), 0600); err != nil {
		return nil, err
	}
	return key, nil
}

func q0kcjsmh(domain, certPath, keyPath, keyLength string) error {
	key, err := rLtaOfBZh()
	if err != nil {
		return fmt.Errorf("load acme account key: %w", err)
	}
	user := &s_CjqL_8{zlC6WUcq: key}
	config := lego.NewConfig(user)
	config.CADirURL = "https://acme-v02.api.letsencrypt.org/directory"
	config.Certificate.KeyType = certcrypto.EC256
	if keyLength == "ec-384" {
		config.Certificate.KeyType = certcrypto.EC384
	}
	client, err := lego.NewClient(config)
	if err != nil {
		return err
	}
	if err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "80")); err != nil {
		return err
	}
	user.w0zKHy_Ex0W, err = client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return fmt.Errorf("acme register: %w", err)
	}
	cert, err := client.Certificate.Obtain(certificate.ObtainRequest{Domains: []string{domain}, Bundle: true})
	if err != nil {
		return err
	}
	if err = os.WriteFile(keyPath, cert.PrivateKey, 0600); err != nil {
		return err
	}
	return os.WriteFile(certPath, cert.Certificate, 0644)
}

func cKh3OZ4rAaWw(provider string) ([]string, []string, int) {
	args := append([]string{}, "--dns", provider)
	env := []string{}
	count := 0
	for _, field := range FHSbLFuj2() {
		if value := cNT24rG4F.CertificateDNSValue(strings.ToLower(field)); value != "" {
			env = append(env, field[4:]+"="+value)
			count++
		}
	}
	return args, env, count
}

func u9iC_y2SB(renew, force bool) error {
	mode, domain, keyLength, provider := cNT24rG4F.CertificateIssuanceSettings()
	certPath, keyPath := cNT24rG4F.ManualCertificateFiles()
	keyArgs := []string{}
	if keyLength != "" {
		if keyLength != "ec-256" && keyLength != "ec-384" {
			return errors.New("cert_key_length only support ec-256 or ec-384 for now")
		}
		keyArgs = []string{"--keylength", keyLength}
	}
	if mode != "dns" {
		muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "start apply certificate, please waiting...")
		err := q0kcjsmh(domain, certPath, keyPath, keyLength)
		if err != nil {
			muc279Mpa6S4.TRTbrquD(slog.LevelError, "apply certificate failed:", err)
		} else {
			muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "apply certificate success")
		}
		return err
	}
	env := append([]string{}, os.Environ()...)
	action := "--issue"
	if renew {
		action = "--renew"
	}
	args := []string{"/root/.acme.sh/acme.sh", action, "-d", domain, "--fullchain-file", certPath, "--key-file", keyPath}
	args = append(args, keyArgs...)
	args = append(args, "--server", "letsencrypt")
	dnsArgs, dnsEnv, count := cKh3OZ4rAaWw(provider)
	if count == 0 {
		return errors.New("can not find \"DNS_XXX\" in config, please configure it")
	}
	args = append(args, dnsArgs...)
	env = append(env, dnsEnv...)
	if force {
		args = append(args, "--force")
	}
	// 原函数直接拼接 shell 字符串；保留该行为，没有新增转义规则。
	cmd := exec.Command("sh", "-c", strings.Join(args, " "))
	cmd.Env = env
	muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "start apply certificate, please waiting...")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	_ = cmd.Start() // 原版忽略 Start 和 Pipe 错误，由 Wait 返回错误。
	go u9iC_y2SBfunc1(stdout)
	go u9iC_y2SBfunc2(stderr)
	err := cmd.Wait()
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(slog.LevelError, "exec \"%s\" failed\n", strings.Join(args, " "))
		return err
	}
	if cmd.ProcessState.ExitCode() != 0 {
		muc279Mpa6S4.V8GN5Stk(slog.LevelError, "apply certificate failed, exit code: %d\n", cmd.ProcessState.ExitCode())
		return err
	}
	muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "apply certificate success")
	return nil
}

// 原匿名闭包 0x15472c0/0x15471c0；Read 返回 n>0 和 err 时也直接退出，丢弃该批数据。
func u9iC_y2SBfunc1(reader io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			return
		}
		log.Print(string(buf[:n]))
	}
}
func u9iC_y2SBfunc2(reader io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			return
		}
		log.Print(string(buf[:n]))
	}
}

func TFwUsh(ctx context.Context) error {
	muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "start apply certificate, it make take up to 60 seconds.")
	mode, domain, _, provider := cNT24rG4F.CertificateIssuanceSettings()
	if domain == "" {
		return errors.New("please configure cert_domain")
	}
	muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "cert domain:", domain)
	if mode != "http" && mode != "dns" && mode != "self" {
		return fmt.Errorf("please configure cert_mode, must be one of: 'http', 'dns', 'self', not '%v'", mode)
	}
	if mode == "dns" && provider == "" {
		return errors.New("please configure dns_provider")
	}
	certPath, keyPath := fmt.Sprintf("/etc/soga/%s.crt", domain), fmt.Sprintf("/etc/soga/%s.key", domain)
	muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "cert file:", certPath)
	muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "key file:", keyPath)
	cNT24rG4F.SetCertificateFiles(certPath, keyPath)
	certExists, keyExists := CAodbhZf.VSCJLy2No(certPath), CAodbhZf.VSCJLy2No(keyPath)
	if mode == "self" {
		if certExists && keyExists {
			muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "self-signed cert already exists, use it directly")
			return nil
		}
		return lLTDzR(domain, certPath, keyPath, 1)
	}
	if certExists && !keyExists {
		return fmt.Errorf("key_file not exist: %v", keyPath)
	}
	if !certExists && keyExists {
		return fmt.Errorf("cert_file not exist: %v", certPath)
	}
	if !certExists {
		muc279Mpa6S4.TRTbrquD(slog.LevelInfo, "start force renew")
	}
	go hRYbCxp(ctx)
	if certExists {
		expires, err := sur_b39JM6l(certPath)
		if err != nil {
			return err
		}
		hours := expires.Sub(time.Now()).Hours()
		muc279Mpa6S4.V8GN5Stk(slog.LevelInfo, "cert expiry time is %v, remain %v days", expires, int(hours/24))
		if hours > 240 {
			return nil
		}
	}
	return u9iC_y2SB(false, true)
}

func hRYbCxp(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			certPath, _ := cNT24rG4F.ManualCertificateFiles()
			expires, err := sur_b39JM6l(certPath)
			if err != nil {
				muc279Mpa6S4.TRTbrquD(slog.LevelDebug, err)
				continue
			}
			if expires.Sub(time.Now()).Hours() <= 120 {
				if err = u9iC_y2SB(false, true); err != nil {
					muc279Mpa6S4.TRTbrquD(slog.LevelError, "apply cert failed:", err)
				} else if err = CAodbhZf.Jx2ev8z(); err != nil {
					muc279Mpa6S4.TRTbrquD(slog.LevelError, "reload failed:", err)
				}
			}
		}
	}
}
