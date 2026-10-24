package tkbUwvW3ex

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// KUD0DUCag is the original 72-byte request, RTTI 0x36b2180.
// The derived session key is private and therefore never serialized.
type KUD0DUCag struct {
	YQ02ZfnPFDqE string `json:"data"`
	Nc_QRM4fWhLJ string `json:"key"`
	EK9X3prNY    string `json:"nonce"`
	tbNcCO9yF    []byte
}

// D4bEdSC7On is the original 32-byte encrypted response, RTTI 0x362a350.
type D4bEdSC7On struct {
	LGzvNG string `json:"data"`
	C4YQM8 string `json:"nonce"`
}

// These exact strings were decoded from the immutable original instructions;
// see session-public-key.json and session-kdf-strings.json in the evidence tree.
const licensePublicKey = "BErNqt680F4x/BfxAEeaQeMD1prMFbkzftSE9NPdWgPxU4hChGUz0UqIFj48EQgpzbJSJAcIBUTklOOQKLDp7Do="
const licenseKeyContext = "7720ecbefa703d76dbdddf5bf9743611"

// kVljP2 recovers 0xec1260. The function value at 0x376a390 points
// to crypto/sha256.New (0x72a4e0); Extract and Expand are separate calls.
func kVljP2(secret, salt []byte, info string, length int) ([]byte, error) {
	key, err := hkdf.Extract(sha256.New, secret, salt)
	if err != nil {
		return nil, err
	}
	return hkdf.Expand(sha256.New, key, info, length)
}

// aStQje recovers 0xecb7a0, including both identical decoded salt and info.
func aStQje(secret []byte) ([]byte, error) {
	return kVljP2(secret, []byte(licenseKeyContext), licenseKeyContext, 32)
}

// LBab24XA recovers 0xec2d00. off_5029B50 points to the P-256 curve.
// Requests authenticate the fixed context string; responses use no AAD.
func LBab24XA(data []byte) (*KUD0DUCag, error) {
	encodedKey, err := base64.StdEncoding.DecodeString(licensePublicKey)
	if err != nil {
		return nil, fmt.Errorf("decode public key: %w", err)
	}
	publicKey, err := ecdh.P256().NewPublicKey(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("parse public key: %w", err)
	}
	privateKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate ephemeral key: %w", err)
	}
	secret, err := privateKey.ECDH(publicKey)
	if err != nil {
		return nil, fmt.Errorf("ecdh: %w", err)
	}
	key, err := aStQje(secret)
	if err != nil {
		return nil, fmt.Errorf("derive key: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create gcm: %w", err)
	}
	nonce := make([]byte, 12)
	if _, err = rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("generate nonce: %w", err)
	}
	encrypted := gcm.Seal(nil, nonce, data, []byte(licenseKeyContext))
	return &KUD0DUCag{
		YQ02ZfnPFDqE: base64.StdEncoding.EncodeToString(encrypted),
		Nc_QRM4fWhLJ: base64.StdEncoding.EncodeToString(privateKey.PublicKey().Bytes()),
		EK9X3prNY:    base64.StdEncoding.EncodeToString(nonce),
		tbNcCO9yF:    key,
	}, nil
}

// Decode recovers 0xec8060. Its original Open call has a nil AAD slice;
// nonce length validation is delegated to GCM, including its panic behavior.
func (response *D4bEdSC7On) Decode(request *KUD0DUCag) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(response.LGzvNG)
	if err != nil {
		return nil, fmt.Errorf("decode response data: %w", err)
	}
	nonce, err := base64.StdEncoding.DecodeString(response.C4YQM8)
	if err != nil {
		return nil, fmt.Errorf("decode nonce: %w", err)
	}
	block, err := aes.NewCipher(request.tbNcCO9yF)
	if err != nil {
		return nil, fmt.Errorf("create cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create gcm: %w", err)
	}
	data, err = gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypt response: %w", err)
	}
	return data, nil
}

// UnopiKFbiaEJ is the decrypted 24-byte response, RTTI 0x3652978.
// The dispatcher reads code; expiry and max_user_num remain decoded fields.
type UnopiKFbiaEJ struct {
	IJ9cKUVV2nxt int   `json:"code"`
	HiBS9AChX    int64 `json:"expiry_time"`
	BwXCgH6ywzx  int   `json:"max_user_num"`
}
