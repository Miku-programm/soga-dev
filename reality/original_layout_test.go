// 根据原样本 RTTI 验证 61 个结构的大小、字段名、偏移、类型、嵌入标志和标签。
package reality

import (
	"reflect"
	"testing"
)

func TestOriginalTypeLayouts(t *testing.T) {
	{
		typ := reflect.TypeOf(serverHelloDoneMsg{})
		if typ.Size() != 0 || typ.NumField() != 0 {
			t.Fatalf("serverHelloDoneMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
	}
	{
		typ := reflect.TypeOf(helloRequestMsg{})
		if typ.Size() != 0 || typ.NumField() != 0 {
			t.Fatalf("helloRequestMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
	}
	{
		typ := reflect.TypeOf(endOfEarlyDataMsg{})
		if typ.Size() != 0 || typ.NumField() != 0 {
			t.Fatalf("endOfEarlyDataMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
	}
	{
		typ := reflect.TypeOf(ClientSessionState{})
		if typ.Size() != 8 || typ.NumField() != 1 {
			t.Fatalf("ClientSessionState layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "session" || f.Offset != 0 || f.Type.String() != "*reality.SessionState" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientSessionState field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(keyUpdateMsg{})
		if typ.Size() != 1 || typ.NumField() != 1 {
			t.Fatalf("keyUpdateMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "updateRequested" || f.Offset != 0 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("keyUpdateMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(finishedMsg{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("finishedMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "verifyData" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(newSessionTicketMsg{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("newSessionTicketMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "ticket" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("newSessionTicketMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(certificateMsg{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("certificateMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "certificates" || f.Offset != 0 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(certificateStatusMsg{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("certificateStatusMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "response" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateStatusMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(serverKeyExchangeMsg{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("serverKeyExchangeMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "key" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverKeyExchangeMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(ECHRejectionError{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("ECHRejectionError layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "RetryConfigList" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ECHRejectionError field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(clientKeyExchangeMsg{})
		if typ.Size() != 24 || typ.NumField() != 1 {
			t.Fatalf("clientKeyExchangeMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "ciphertext" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientKeyExchangeMsg field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(echConfigErr{})
		if typ.Size() != 16 || typ.NumField() != 1 {
			t.Fatalf("echConfigErr layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "field" || f.Offset != 0 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echConfigErr field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(permanentError{})
		if typ.Size() != 16 || typ.NumField() != 1 {
			t.Fatalf("permanentError layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "err" || f.Offset != 0 || f.Type.String() != "net.Error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("permanentError field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(weakCertCache{})
		if typ.Size() != 48 || typ.NumField() != 1 {
			t.Fatalf("weakCertCache layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Map" || f.Offset != 0 || f.Type.String() != "sync.Map" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("weakCertCache field 0: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(keyShare{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatalf("keyShare layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "group" || f.Offset != 0 || f.Type.String() != "reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("keyShare field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "data" || f.Offset != 8 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("keyShare field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(pskIdentity{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatalf("pskIdentity layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "label" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("pskIdentity field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "obfuscatedTicketAge" || f.Offset != 24 || f.Type.String() != "uint32" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("pskIdentity field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(EchCipher{})
		if typ.Size() != 4 || typ.NumField() != 2 {
			t.Fatalf("EchCipher layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "KDFID" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchCipher field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "AEADID" || f.Offset != 2 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchCipher field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(echExtension{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatalf("echExtension layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Type" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echExtension field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "Data" || f.Offset != 8 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echExtension field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(CertificateVerificationError{})
		if typ.Size() != 40 || typ.NumField() != 2 {
			t.Fatalf("CertificateVerificationError layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "UnverifiedCertificates" || f.Offset != 0 || f.Type.String() != "[]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("CertificateVerificationError field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "Err" || f.Offset != 24 || f.Type.String() != "error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("CertificateVerificationError field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(atLeastReader{})
		if typ.Size() != 24 || typ.NumField() != 2 {
			t.Fatalf("atLeastReader layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "R" || f.Offset != 0 || f.Type.String() != "io.Reader" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("atLeastReader field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "N" || f.Offset != 16 || f.Type.String() != "int64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("atLeastReader field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(xorNonceAEAD{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatalf("xorNonceAEAD layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "nonceMask" || f.Offset != 0 || f.Type.String() != "[12]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("xorNonceAEAD field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "aead" || f.Offset != 16 || f.Type.String() != "cipher.AEAD" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("xorNonceAEAD field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(prefixNonceAEAD{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatalf("prefixNonceAEAD layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "nonce" || f.Offset != 0 || f.Type.String() != "[12]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("prefixNonceAEAD field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "aead" || f.Offset != 16 || f.Type.String() != "cipher.AEAD" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("prefixNonceAEAD field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(rawExtension{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatalf("rawExtension layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "extType" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("rawExtension field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "data" || f.Offset != 8 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("rawExtension field 1: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(rsaKeyAgreement{})
		if typ.Size() != 0 || typ.NumField() != 0 {
			t.Fatalf("rsaKeyAgreement layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
	}
	{
		typ := reflect.TypeOf(LimitFallback{})
		if typ.Size() != 24 || typ.NumField() != 3 {
			t.Fatalf("LimitFallback layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "AfterBytes" || f.Offset != 0 || f.Type.String() != "uint64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("LimitFallback field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "BytesPerSec" || f.Offset != 8 || f.Type.String() != "uint64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("LimitFallback field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "BurstBytesPerSec" || f.Offset != 16 || f.Type.String() != "uint64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("LimitFallback field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(EncryptedClientHelloKey{})
		if typ.Size() != 56 || typ.NumField() != 3 {
			t.Fatalf("EncryptedClientHelloKey layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Config" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EncryptedClientHelloKey field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "PrivateKey" || f.Offset != 24 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EncryptedClientHelloKey field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "SendAsRetry" || f.Offset != 48 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EncryptedClientHelloKey field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(ticketKey{})
		if typ.Size() != 56 || typ.NumField() != 3 {
			t.Fatalf("ticketKey layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "aesKey" || f.Offset != 0 || f.Type.String() != "[16]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ticketKey field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "hmacKey" || f.Offset != 16 || f.Type.String() != "[16]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ticketKey field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "created" || f.Offset != 32 || f.Type.String() != "time.Time" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ticketKey field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(keySharePrivateKeys{})
		if typ.Size() != 24 || typ.NumField() != 3 {
			t.Fatalf("keySharePrivateKeys layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "curveID" || f.Offset != 0 || f.Type.String() != "reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("keySharePrivateKeys field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ecdhe" || f.Offset != 8 || f.Type.String() != "*ecdh.PrivateKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("keySharePrivateKeys field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "mlkem" || f.Offset != 16 || f.Type.String() != "*mlkem.DecapsulationKey768" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("keySharePrivateKeys field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(certificateVerifyMsg{})
		if typ.Size() != 32 || typ.NumField() != 3 {
			t.Fatalf("certificateVerifyMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "hasSignatureAlgorithm" || f.Offset != 0 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateVerifyMsg field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "signatureAlgorithm" || f.Offset != 2 || f.Type.String() != "reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateVerifyMsg field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "signature" || f.Offset != 8 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateVerifyMsg field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(certificateMsgTLS13{})
		if typ.Size() != 128 || typ.NumField() != 3 {
			t.Fatalf("certificateMsgTLS13 layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "certificate" || f.Offset != 0 || f.Type.String() != "reality.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateMsgTLS13 field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ocspStapling" || f.Offset != 120 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateMsgTLS13 field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "scts" || f.Offset != 121 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateMsgTLS13 field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(RecordHeaderError{})
		if typ.Size() != 40 || typ.NumField() != 3 {
			t.Fatalf("RecordHeaderError layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Msg" || f.Offset != 0 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("RecordHeaderError field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "RecordHeader" || f.Offset != 16 || f.Type.String() != "[5]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("RecordHeaderError field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Conn" || f.Offset != 24 || f.Type.String() != "net.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("RecordHeaderError field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(CertificateRequestInfo{})
		if typ.Size() != 72 || typ.NumField() != 4 {
			t.Fatalf("CertificateRequestInfo layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "AcceptableCAs" || f.Offset != 0 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("CertificateRequestInfo field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "SignatureSchemes" || f.Offset != 24 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("CertificateRequestInfo field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Version" || f.Offset != 48 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("CertificateRequestInfo field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "ctx" || f.Offset != 56 || f.Type.String() != "context.Context" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("CertificateRequestInfo field 3: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(cipherSuiteTLS13{})
		if typ.Size() != 32 || typ.NumField() != 4 {
			t.Fatalf("cipherSuiteTLS13 layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "id" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuiteTLS13 field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "keyLen" || f.Offset != 8 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuiteTLS13 field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "aead" || f.Offset != 16 || f.Type.String() != "func([]uint8, []uint8) reality.aead" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuiteTLS13 field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "hash" || f.Offset != 24 || f.Type.String() != "crypto.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuiteTLS13 field 3: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(encryptedExtensionsMsg{})
		if typ.Size() != 72 || typ.NumField() != 4 {
			t.Fatalf("encryptedExtensionsMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "alpnProtocol" || f.Offset != 0 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("encryptedExtensionsMsg field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "quicTransportParameters" || f.Offset != 16 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("encryptedExtensionsMsg field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "earlyData" || f.Offset != 40 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("encryptedExtensionsMsg field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "echRetryConfigs" || f.Offset != 48 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("encryptedExtensionsMsg field 3: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(certificateRequestMsg{})
		if typ.Size() != 80 || typ.NumField() != 4 {
			t.Fatalf("certificateRequestMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "hasSignatureAlgorithm" || f.Offset != 0 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsg field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "certificateTypes" || f.Offset != 8 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsg field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "supportedSignatureAlgorithms" || f.Offset != 32 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsg field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "certificateAuthorities" || f.Offset != 56 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsg field 3: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(QUICEvent{})
		if typ.Size() != 56 || typ.NumField() != 5 {
			t.Fatalf("QUICEvent layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Kind" || f.Offset != 0 || f.Type.String() != "reality.QUICEventKind" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("QUICEvent field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "Level" || f.Offset != 8 || f.Type.String() != "reality.QUICEncryptionLevel" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("QUICEvent field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Data" || f.Offset != 16 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("QUICEvent field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "Suite" || f.Offset != 40 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("QUICEvent field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "SessionState" || f.Offset != 48 || f.Type.String() != "*reality.SessionState" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("QUICEvent field 4: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(newSessionTicketMsgTLS13{})
		if typ.Size() != 64 || typ.NumField() != 5 {
			t.Fatalf("newSessionTicketMsgTLS13 layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "lifetime" || f.Offset != 0 || f.Type.String() != "uint32" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("newSessionTicketMsgTLS13 field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ageAdd" || f.Offset != 4 || f.Type.String() != "uint32" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("newSessionTicketMsgTLS13 field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "nonce" || f.Offset != 8 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("newSessionTicketMsgTLS13 field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "label" || f.Offset != 32 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("newSessionTicketMsgTLS13 field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "maxEarlyData" || f.Offset != 56 || f.Type.String() != "uint32" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("newSessionTicketMsgTLS13 field 4: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(echServerContext{})
		if typ.Size() != 40 || typ.NumField() != 5 {
			t.Fatalf("echServerContext layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "hpkeContext" || f.Offset != 0 || f.Type.String() != "*hpke.Recipient" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echServerContext field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "configID" || f.Offset != 8 || f.Type.String() != "uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echServerContext field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "ciphersuite" || f.Offset != 10 || f.Type.String() != "reality.EchCipher" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echServerContext field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "transcript" || f.Offset != 16 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echServerContext field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "inner" || f.Offset != 32 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echServerContext field 4: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(certificateRequestMsgTLS13{})
		if typ.Size() != 80 || typ.NumField() != 5 {
			t.Fatalf("certificateRequestMsgTLS13 layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "ocspStapling" || f.Offset != 0 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsgTLS13 field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "scts" || f.Offset != 1 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsgTLS13 field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "supportedSignatureAlgorithms" || f.Offset != 8 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsgTLS13 field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "supportedSignatureAlgorithmsCert" || f.Offset != 32 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsgTLS13 field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "certificateAuthorities" || f.Offset != 56 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("certificateRequestMsgTLS13 field 4: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(ecdheKeyAgreement{})
		if typ.Size() != 48 || typ.NumField() != 5 {
			t.Fatalf("ecdheKeyAgreement layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "version" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ecdheKeyAgreement field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "isRSA" || f.Offset != 2 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ecdheKeyAgreement field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "key" || f.Offset != 8 || f.Type.String() != "*ecdh.PrivateKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ecdheKeyAgreement field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "ckx" || f.Offset != 16 || f.Type.String() != "*reality.clientKeyExchangeMsg" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ecdheKeyAgreement field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "preMasterSecret" || f.Offset != 24 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ecdheKeyAgreement field 4: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(listener{})
		if typ.Size() != 48 || typ.NumField() != 4 {
			t.Fatalf("listener layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "O6B2qZ0p0" || f.Offset != 0 || f.Type.String() != "net.Listener" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("listener field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "config" || f.Offset != 16 || f.Type.String() != "*reality.Config" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("listener field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "conns" || f.Offset != 24 || f.Type.String() != "chan net.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("listener field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "err" || f.Offset != 32 || f.Type.String() != "error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("listener field 3: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(Certificate{})
		if typ.Size() != 120 || typ.NumField() != 6 {
			t.Fatalf("Certificate layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Certificate" || f.Offset != 0 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Certificate field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "PrivateKey" || f.Offset != 24 || f.Type.String() != "crypto.PrivateKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Certificate field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "SupportedSignatureAlgorithms" || f.Offset != 40 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Certificate field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "OCSPStaple" || f.Offset != 64 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Certificate field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "SignedCertificateTimestamps" || f.Offset != 88 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Certificate field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "Leaf" || f.Offset != 112 || f.Type.String() != "*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Certificate field 5: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(MirrorConn{})
		if typ.Size() != 40 || typ.NumField() != 3 {
			t.Fatalf("MirrorConn layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Mutex" || f.Offset != 0 || f.Type.String() != "*sync.Mutex" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("MirrorConn field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "DrpXao" || f.Offset != 8 || f.Type.String() != "net.Conn" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("MirrorConn field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Target" || f.Offset != 24 || f.Type.String() != "net.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("MirrorConn field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(DetectConn{})
		if typ.Size() != 40 || typ.NumField() != 3 {
			t.Fatalf("DetectConn layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "DrpXao" || f.Offset != 0 || f.Type.String() != "net.Conn" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("DetectConn field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "Key" || f.Offset != 16 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("DetectConn field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "CcsSent" || f.Offset != 32 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("DetectConn field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(RatelimitedConn{})
		if typ.Size() != 32 || typ.NumField() != 3 {
			t.Fatalf("RatelimitedConn layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "DrpXao" || f.Offset != 0 || f.Type.String() != "net.Conn" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("RatelimitedConn field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "After" || f.Offset != 16 || f.Type.String() != "int64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("RatelimitedConn field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Bucket" || f.Offset != 24 || f.Type.String() != "*ratelimit.Bucket" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("RatelimitedConn field 2: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(echClientContext{})
		if typ.Size() != 96 || typ.NumField() != 9 {
			t.Fatalf("echClientContext layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "config" || f.Offset != 0 || f.Type.String() != "*reality.EchConfig" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "hpkeContext" || f.Offset != 8 || f.Type.String() != "*hpke.Sender" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "encapsulatedKey" || f.Offset != 16 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "innerHello" || f.Offset != 40 || f.Type.String() != "*reality.clientHelloMsg" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "innerTranscript" || f.Offset != 48 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "kdfID" || f.Offset != 64 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "aeadID" || f.Offset != 66 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "echRejected" || f.Offset != 68 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "retryConfigs" || f.Offset != 72 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("echClientContext field 8: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(cipherSuite{})
		if typ.Size() != 72 || typ.NumField() != 9 {
			t.Fatalf("cipherSuite layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "id" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "keyLen" || f.Offset != 8 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "macLen" || f.Offset != 16 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "ivLen" || f.Offset != 24 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "ka" || f.Offset != 32 || f.Type.String() != "func(uint16) reality.keyAgreement" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "flags" || f.Offset != 40 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "cipher" || f.Offset != 48 || f.Type.String() != "func([]uint8, []uint8, bool) interface {}" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "mac" || f.Offset != 56 || f.Type.String() != "func([]uint8) hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "aead" || f.Offset != 64 || f.Type.String() != "func([]uint8, []uint8) reality.aead" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("cipherSuite field 8: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(clientHandshakeState{})
		if typ.Size() != 208 || typ.NumField() != 9 {
			t.Fatalf("clientHandshakeState layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "c" || f.Offset != 0 || f.Type.String() != "*reality.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ctx" || f.Offset != 8 || f.Type.String() != "context.Context" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "serverHello" || f.Offset != 24 || f.Type.String() != "*reality.serverHelloMsg" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "hello" || f.Offset != 32 || f.Type.String() != "*reality.clientHelloMsg" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "suite" || f.Offset != 40 || f.Type.String() != "*reality.cipherSuite" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "finishedHash" || f.Offset != 48 || f.Type.String() != "reality.finishedHash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "masterSecret" || f.Offset != 152 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "session" || f.Offset != 176 || f.Type.String() != "*reality.SessionState" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "ticket" || f.Offset != 184 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHandshakeState field 8: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(finishedHash{})
		if typ.Size() != 104 || typ.NumField() != 7 {
			t.Fatalf("finishedHash layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "client" || f.Offset != 0 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "server" || f.Offset != 16 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "clientMD5" || f.Offset != 32 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "serverMD5" || f.Offset != 48 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "buffer" || f.Offset != 64 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "version" || f.Offset != 88 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "prf" || f.Offset != 96 || f.Type.String() != "reality.prfFunc" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("finishedHash field 6: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(EchConfig{})
		if typ.Size() != 136 || typ.NumField() != 10 {
			t.Fatalf("EchConfig layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "raw" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "Version" || f.Offset != 24 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Length" || f.Offset != 26 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "ConfigID" || f.Offset != 28 || f.Type.String() != "uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "KemID" || f.Offset != 30 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "PublicKey" || f.Offset != 32 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "SymmetricCipherSuite" || f.Offset != 56 || f.Type.String() != "[]reality.EchCipher" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "MaxNameLength" || f.Offset != 80 || f.Type.String() != "uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "PublicName" || f.Offset != 88 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "Extensions" || f.Offset != 112 || f.Type.String() != "[]reality.echExtension" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("EchConfig field 9: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(ClientHelloInfo{})
		if typ.Size() != 224 || typ.NumField() != 11 {
			t.Fatalf("ClientHelloInfo layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "CipherSuites" || f.Offset != 0 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ServerName" || f.Offset != 24 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "SupportedCurves" || f.Offset != 40 || f.Type.String() != "[]reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "SupportedPoints" || f.Offset != 64 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "SignatureSchemes" || f.Offset != 88 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "SupportedProtos" || f.Offset != 112 || f.Type.String() != "[]string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "SupportedVersions" || f.Offset != 136 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "Extensions" || f.Offset != 160 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "Conn" || f.Offset != 184 || f.Type.String() != "net.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "config" || f.Offset != 200 || f.Type.String() != "*reality.Config" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "ctx" || f.Offset != 208 || f.Type.String() != "context.Context" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ClientHelloInfo field 10: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(quicState{})
		if typ.Size() != 584 || typ.NumField() != 12 {
			t.Fatalf("quicState layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "events" || f.Offset != 0 || f.Type.String() != "[]reality.QUICEvent" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "nextEvent" || f.Offset != 24 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "eventArr" || f.Offset != 32 || f.Type.String() != "[8]reality.QUICEvent" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "started" || f.Offset != 480 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "signalc" || f.Offset != 488 || f.Type.String() != "chan struct {}" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "blockedc" || f.Offset != 496 || f.Type.String() != "chan struct {}" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "cancelc" || f.Offset != 504 || f.Type.String() != "<-chan struct {}" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "cancel" || f.Offset != 512 || f.Type.String() != "context.CancelFunc" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "waitingForDrain" || f.Offset != 520 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "readbuf" || f.Offset != 528 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "transportParams" || f.Offset != 552 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "enableSessionEvents" || f.Offset != 576 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("quicState field 11: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(halfConn{})
		if typ.Size() != 232 || typ.NumField() != 13 {
			t.Fatalf("halfConn layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "handshakeLen" || f.Offset != 0 || f.Type.String() != "[7]int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "handshakeBuf" || f.Offset != 56 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Mutex" || f.Offset != 80 || f.Type.String() != "sync.Mutex" || f.Anonymous != true || string(f.Tag) != "" {
			t.Errorf("halfConn field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "err" || f.Offset != 88 || f.Type.String() != "error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "version" || f.Offset != 104 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "cipher" || f.Offset != 112 || f.Type.String() != "interface {}" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "mac" || f.Offset != 128 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "seq" || f.Offset != 144 || f.Type.String() != "[8]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "scratchBuf" || f.Offset != 152 || f.Type.String() != "[13]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "nextCipher" || f.Offset != 168 || f.Type.String() != "interface {}" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "nextMac" || f.Offset != 184 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "level" || f.Offset != 200 || f.Type.String() != "reality.QUICEncryptionLevel" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "trafficSecret" || f.Offset != 208 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("halfConn field 12: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(ConnectionState{})
		if typ.Size() != 192 || typ.NumField() != 16 {
			t.Fatalf("ConnectionState layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Version" || f.Offset != 0 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "HandshakeComplete" || f.Offset != 2 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "DidResume" || f.Offset != 3 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "CipherSuite" || f.Offset != 4 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "CurveID" || f.Offset != 6 || f.Type.String() != "reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "NegotiatedProtocol" || f.Offset != 8 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "NegotiatedProtocolIsMutual" || f.Offset != 24 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "ServerName" || f.Offset != 32 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "PeerCertificates" || f.Offset != 48 || f.Type.String() != "[]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "VerifiedChains" || f.Offset != 72 || f.Type.String() != "[][]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "SignedCertificateTimestamps" || f.Offset != 96 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "OCSPResponse" || f.Offset != 120 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "TLSUnique" || f.Offset != 144 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "ECHAccepted" || f.Offset != 168 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "ekm" || f.Offset != 176 || f.Type.String() != "func(string, []uint8, int) ([]uint8, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "testingOnlyDidHRR" || f.Offset != 184 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("ConnectionState field 15: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(SessionState{})
		if typ.Size() != 232 || typ.NumField() != 17 {
			t.Fatalf("SessionState layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "Extra" || f.Offset != 0 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "EarlyData" || f.Offset != 24 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "version" || f.Offset != 26 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "isClient" || f.Offset != 28 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "cipherSuite" || f.Offset != 30 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "createdAt" || f.Offset != 32 || f.Type.String() != "uint64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "secret" || f.Offset != 40 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "extMasterSecret" || f.Offset != 64 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "peerCertificates" || f.Offset != 72 || f.Type.String() != "[]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "ocspResponse" || f.Offset != 96 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "scts" || f.Offset != 120 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "verifiedChains" || f.Offset != 144 || f.Type.String() != "[][]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "alpnProtocol" || f.Offset != 168 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "useBy" || f.Offset != 184 || f.Type.String() != "uint64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "ageAdd" || f.Offset != 192 || f.Type.String() != "uint32" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "ticket" || f.Offset != 200 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 15: %+v", f)
		}
		if f := typ.Field(16); f.Name != "curveID" || f.Offset != 224 || f.Type.String() != "reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("SessionState field 16: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(serverHandshakeStateTLS13{})
		if typ.Size() != 192 || typ.NumField() != 18 {
			t.Fatalf("serverHandshakeStateTLS13 layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "c" || f.Offset != 0 || f.Type.String() != "*reality.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ctx" || f.Offset != 8 || f.Type.String() != "context.Context" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "clientHello" || f.Offset != 24 || f.Type.String() != "*reality.clientHelloMsg" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "hello" || f.Offset != 32 || f.Type.String() != "*reality.serverHelloMsg" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "sentDummyCCS" || f.Offset != 40 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "usingPSK" || f.Offset != 41 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "earlyData" || f.Offset != 42 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "suite" || f.Offset != 48 || f.Type.String() != "*reality.cipherSuiteTLS13" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "cert" || f.Offset != 56 || f.Type.String() != "*reality.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "sigAlg" || f.Offset != 64 || f.Type.String() != "reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "earlySecret" || f.Offset != 72 || f.Type.String() != "*tls13.EarlySecret" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "sharedKey" || f.Offset != 80 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "handshakeSecret" || f.Offset != 104 || f.Type.String() != "*tls13.HandshakeSecret" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "masterSecret" || f.Offset != 112 || f.Type.String() != "*tls13.MasterSecret" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "trafficSecret" || f.Offset != 120 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "transcript" || f.Offset != 144 || f.Type.String() != "hash.Hash" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 15: %+v", f)
		}
		if f := typ.Field(16); f.Name != "clientFinished" || f.Offset != 160 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 16: %+v", f)
		}
		if f := typ.Field(17); f.Name != "echContext" || f.Offset != 184 || f.Type.String() != "*reality.echServerContext" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHandshakeStateTLS13 field 17: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(serverHelloMsg{})
		if typ.Size() != 296 || typ.NumField() != 22 {
			t.Fatalf("serverHelloMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "original" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "vers" || f.Offset != 24 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "random" || f.Offset != 32 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "sessionId" || f.Offset != 56 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "cipherSuite" || f.Offset != 80 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "compressionMethod" || f.Offset != 82 || f.Type.String() != "uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "ocspStapling" || f.Offset != 83 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "ticketSupported" || f.Offset != 84 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "secureRenegotiationSupported" || f.Offset != 85 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "secureRenegotiation" || f.Offset != 88 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "extendedMasterSecret" || f.Offset != 112 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "alpnProtocol" || f.Offset != 120 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "scts" || f.Offset != 136 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "supportedVersion" || f.Offset != 160 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "serverShare" || f.Offset != 168 || f.Type.String() != "reality.keyShare" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "selectedIdentityPresent" || f.Offset != 200 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 15: %+v", f)
		}
		if f := typ.Field(16); f.Name != "selectedIdentity" || f.Offset != 202 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 16: %+v", f)
		}
		if f := typ.Field(17); f.Name != "supportedPoints" || f.Offset != 208 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 17: %+v", f)
		}
		if f := typ.Field(18); f.Name != "encryptedClientHello" || f.Offset != 232 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 18: %+v", f)
		}
		if f := typ.Field(19); f.Name != "serverNameAck" || f.Offset != 256 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 19: %+v", f)
		}
		if f := typ.Field(20); f.Name != "cookie" || f.Offset != 264 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 20: %+v", f)
		}
		if f := typ.Field(21); f.Name != "selectedGroup" || f.Offset != 288 || f.Type.String() != "reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("serverHelloMsg field 21: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(clientHelloMsg{})
		if typ.Size() != 576 || typ.NumField() != 29 {
			t.Fatalf("clientHelloMsg layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "original" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "vers" || f.Offset != 24 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "random" || f.Offset != 32 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "sessionId" || f.Offset != 56 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "cipherSuites" || f.Offset != 80 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "compressionMethods" || f.Offset != 104 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "serverName" || f.Offset != 128 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "ocspStapling" || f.Offset != 144 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "supportedCurves" || f.Offset != 152 || f.Type.String() != "[]reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "supportedPoints" || f.Offset != 176 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "ticketSupported" || f.Offset != 200 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "sessionTicket" || f.Offset != 208 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "supportedSignatureAlgorithms" || f.Offset != 232 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "supportedSignatureAlgorithmsCert" || f.Offset != 256 || f.Type.String() != "[]reality.SignatureScheme" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "secureRenegotiationSupported" || f.Offset != 280 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "secureRenegotiation" || f.Offset != 288 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 15: %+v", f)
		}
		if f := typ.Field(16); f.Name != "extendedMasterSecret" || f.Offset != 312 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 16: %+v", f)
		}
		if f := typ.Field(17); f.Name != "alpnProtocols" || f.Offset != 320 || f.Type.String() != "[]string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 17: %+v", f)
		}
		if f := typ.Field(18); f.Name != "scts" || f.Offset != 344 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 18: %+v", f)
		}
		if f := typ.Field(19); f.Name != "supportedVersions" || f.Offset != 352 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 19: %+v", f)
		}
		if f := typ.Field(20); f.Name != "cookie" || f.Offset != 376 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 20: %+v", f)
		}
		if f := typ.Field(21); f.Name != "keyShares" || f.Offset != 400 || f.Type.String() != "[]reality.keyShare" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 21: %+v", f)
		}
		if f := typ.Field(22); f.Name != "earlyData" || f.Offset != 424 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 22: %+v", f)
		}
		if f := typ.Field(23); f.Name != "pskModes" || f.Offset != 432 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 23: %+v", f)
		}
		if f := typ.Field(24); f.Name != "pskIdentities" || f.Offset != 456 || f.Type.String() != "[]reality.pskIdentity" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 24: %+v", f)
		}
		if f := typ.Field(25); f.Name != "pskBinders" || f.Offset != 480 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 25: %+v", f)
		}
		if f := typ.Field(26); f.Name != "quicTransportParameters" || f.Offset != 504 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 26: %+v", f)
		}
		if f := typ.Field(27); f.Name != "encryptedClientHello" || f.Offset != 528 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 27: %+v", f)
		}
		if f := typ.Field(28); f.Name != "extensions" || f.Offset != 552 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("clientHelloMsg field 28: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(Conn{})
		if typ.Size() != 1088 || typ.NumField() != 48 {
			t.Fatalf("Conn layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "AuthKey" || f.Offset != 0 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "ClientVer" || f.Offset != 24 || f.Type.String() != "[3]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "ClientTime" || f.Offset != 32 || f.Type.String() != "time.Time" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "ClientShortId" || f.Offset != 56 || f.Type.String() != "[8]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "conn" || f.Offset != 64 || f.Type.String() != "net.Conn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "isClient" || f.Offset != 80 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "handshakeFn" || f.Offset != 88 || f.Type.String() != "func(context.Context) error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "quic" || f.Offset != 96 || f.Type.String() != "*reality.quicState" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "isHandshakeComplete" || f.Offset != 104 || f.Type.String() != "atomic.Bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "handshakeMutex" || f.Offset != 108 || f.Type.String() != "sync.Mutex" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "handshakeErr" || f.Offset != 120 || f.Type.String() != "error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "vers" || f.Offset != 136 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "haveVers" || f.Offset != 138 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "config" || f.Offset != 144 || f.Type.String() != "*reality.Config" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "handshakes" || f.Offset != 152 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "extMasterSecret" || f.Offset != 160 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 15: %+v", f)
		}
		if f := typ.Field(16); f.Name != "didResume" || f.Offset != 161 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 16: %+v", f)
		}
		if f := typ.Field(17); f.Name != "didHRR" || f.Offset != 162 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 17: %+v", f)
		}
		if f := typ.Field(18); f.Name != "cipherSuite" || f.Offset != 164 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 18: %+v", f)
		}
		if f := typ.Field(19); f.Name != "curveID" || f.Offset != 166 || f.Type.String() != "reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 19: %+v", f)
		}
		if f := typ.Field(20); f.Name != "ocspResponse" || f.Offset != 168 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 20: %+v", f)
		}
		if f := typ.Field(21); f.Name != "scts" || f.Offset != 192 || f.Type.String() != "[][]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 21: %+v", f)
		}
		if f := typ.Field(22); f.Name != "peerCertificates" || f.Offset != 216 || f.Type.String() != "[]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 22: %+v", f)
		}
		if f := typ.Field(23); f.Name != "verifiedChains" || f.Offset != 240 || f.Type.String() != "[][]*x509.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 23: %+v", f)
		}
		if f := typ.Field(24); f.Name != "serverName" || f.Offset != 264 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 24: %+v", f)
		}
		if f := typ.Field(25); f.Name != "secureRenegotiation" || f.Offset != 280 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 25: %+v", f)
		}
		if f := typ.Field(26); f.Name != "ekm" || f.Offset != 288 || f.Type.String() != "func(string, []uint8, int) ([]uint8, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 26: %+v", f)
		}
		if f := typ.Field(27); f.Name != "resumptionSecret" || f.Offset != 296 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 27: %+v", f)
		}
		if f := typ.Field(28); f.Name != "echAccepted" || f.Offset != 320 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 28: %+v", f)
		}
		if f := typ.Field(29); f.Name != "ticketKeys" || f.Offset != 328 || f.Type.String() != "[]reality.ticketKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 29: %+v", f)
		}
		if f := typ.Field(30); f.Name != "clientFinishedIsFirst" || f.Offset != 352 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 30: %+v", f)
		}
		if f := typ.Field(31); f.Name != "closeNotifyErr" || f.Offset != 360 || f.Type.String() != "error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 31: %+v", f)
		}
		if f := typ.Field(32); f.Name != "closeNotifySent" || f.Offset != 376 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 32: %+v", f)
		}
		if f := typ.Field(33); f.Name != "clientFinished" || f.Offset != 377 || f.Type.String() != "[12]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 33: %+v", f)
		}
		if f := typ.Field(34); f.Name != "serverFinished" || f.Offset != 389 || f.Type.String() != "[12]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 34: %+v", f)
		}
		if f := typ.Field(35); f.Name != "clientProtocol" || f.Offset != 408 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 35: %+v", f)
		}
		if f := typ.Field(36); f.Name != "in" || f.Offset != 424 || f.Type.String() != "reality.halfConn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 36: %+v", f)
		}
		if f := typ.Field(37); f.Name != "out" || f.Offset != 656 || f.Type.String() != "reality.halfConn" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 37: %+v", f)
		}
		if f := typ.Field(38); f.Name != "rawInput" || f.Offset != 888 || f.Type.String() != "bytes.Buffer" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 38: %+v", f)
		}
		if f := typ.Field(39); f.Name != "input" || f.Offset != 928 || f.Type.String() != "bytes.Reader" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 39: %+v", f)
		}
		if f := typ.Field(40); f.Name != "hand" || f.Offset != 968 || f.Type.String() != "bytes.Buffer" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 40: %+v", f)
		}
		if f := typ.Field(41); f.Name != "buffering" || f.Offset != 1008 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 41: %+v", f)
		}
		if f := typ.Field(42); f.Name != "sendBuf" || f.Offset != 1016 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 42: %+v", f)
		}
		if f := typ.Field(43); f.Name != "bytesSent" || f.Offset != 1040 || f.Type.String() != "int64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 43: %+v", f)
		}
		if f := typ.Field(44); f.Name != "packetsSent" || f.Offset != 1048 || f.Type.String() != "int64" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 44: %+v", f)
		}
		if f := typ.Field(45); f.Name != "retryCount" || f.Offset != 1056 || f.Type.String() != "int" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 45: %+v", f)
		}
		if f := typ.Field(46); f.Name != "activeCall" || f.Offset != 1064 || f.Type.String() != "atomic.Int32" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 46: %+v", f)
		}
		if f := typ.Field(47); f.Name != "tmp" || f.Offset != 1068 || f.Type.String() != "[16]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Conn field 47: %+v", f)
		}
	}
	{
		typ := reflect.TypeOf(Config{})
		if typ.Size() != 688 || typ.NumField() != 49 {
			t.Fatalf("Config layout: size=%d fields=%d", typ.Size(), typ.NumField())
		}
		if f := typ.Field(0); f.Name != "DialContext" || f.Offset != 0 || f.Type.String() != "func(context.Context, string, string) (net.Conn, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 0: %+v", f)
		}
		if f := typ.Field(1); f.Name != "Show" || f.Offset != 8 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 1: %+v", f)
		}
		if f := typ.Field(2); f.Name != "Type" || f.Offset != 16 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 2: %+v", f)
		}
		if f := typ.Field(3); f.Name != "Dest" || f.Offset != 32 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 3: %+v", f)
		}
		if f := typ.Field(4); f.Name != "Xver" || f.Offset != 48 || f.Type.String() != "uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 4: %+v", f)
		}
		if f := typ.Field(5); f.Name != "ServerNames" || f.Offset != 56 || f.Type.String() != "map[string]bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 5: %+v", f)
		}
		if f := typ.Field(6); f.Name != "PrivateKey" || f.Offset != 64 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 6: %+v", f)
		}
		if f := typ.Field(7); f.Name != "MinClientVer" || f.Offset != 88 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 7: %+v", f)
		}
		if f := typ.Field(8); f.Name != "MaxClientVer" || f.Offset != 112 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 8: %+v", f)
		}
		if f := typ.Field(9); f.Name != "MaxTimeDiff" || f.Offset != 136 || f.Type.String() != "time.Duration" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 9: %+v", f)
		}
		if f := typ.Field(10); f.Name != "ShortIds" || f.Offset != 144 || f.Type.String() != "map[[8]uint8]bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 10: %+v", f)
		}
		if f := typ.Field(11); f.Name != "Mldsa65Key" || f.Offset != 152 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 11: %+v", f)
		}
		if f := typ.Field(12); f.Name != "LimitFallbackUpload" || f.Offset != 176 || f.Type.String() != "reality.LimitFallback" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 12: %+v", f)
		}
		if f := typ.Field(13); f.Name != "LimitFallbackDownload" || f.Offset != 200 || f.Type.String() != "reality.LimitFallback" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 13: %+v", f)
		}
		if f := typ.Field(14); f.Name != "Rand" || f.Offset != 224 || f.Type.String() != "io.Reader" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 14: %+v", f)
		}
		if f := typ.Field(15); f.Name != "Time" || f.Offset != 240 || f.Type.String() != "func() time.Time" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 15: %+v", f)
		}
		if f := typ.Field(16); f.Name != "Certificates" || f.Offset != 248 || f.Type.String() != "[]reality.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 16: %+v", f)
		}
		if f := typ.Field(17); f.Name != "NameToCertificate" || f.Offset != 272 || f.Type.String() != "map[string]*reality.Certificate" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 17: %+v", f)
		}
		if f := typ.Field(18); f.Name != "GetCertificate" || f.Offset != 280 || f.Type.String() != "func(*reality.ClientHelloInfo) (*reality.Certificate, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 18: %+v", f)
		}
		if f := typ.Field(19); f.Name != "GetClientCertificate" || f.Offset != 288 || f.Type.String() != "func(*reality.CertificateRequestInfo) (*reality.Certificate, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 19: %+v", f)
		}
		if f := typ.Field(20); f.Name != "GetConfigForClient" || f.Offset != 296 || f.Type.String() != "func(*reality.ClientHelloInfo) (*reality.Config, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 20: %+v", f)
		}
		if f := typ.Field(21); f.Name != "VerifyPeerCertificate" || f.Offset != 304 || f.Type.String() != "func([][]uint8, [][]*x509.Certificate) error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 21: %+v", f)
		}
		if f := typ.Field(22); f.Name != "VerifyConnection" || f.Offset != 312 || f.Type.String() != "func(reality.ConnectionState) error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 22: %+v", f)
		}
		if f := typ.Field(23); f.Name != "RootCAs" || f.Offset != 320 || f.Type.String() != "*x509.CertPool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 23: %+v", f)
		}
		if f := typ.Field(24); f.Name != "NextProtos" || f.Offset != 328 || f.Type.String() != "[]string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 24: %+v", f)
		}
		if f := typ.Field(25); f.Name != "ServerName" || f.Offset != 352 || f.Type.String() != "string" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 25: %+v", f)
		}
		if f := typ.Field(26); f.Name != "ClientAuth" || f.Offset != 368 || f.Type.String() != "reality.ClientAuthType" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 26: %+v", f)
		}
		if f := typ.Field(27); f.Name != "ClientCAs" || f.Offset != 376 || f.Type.String() != "*x509.CertPool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 27: %+v", f)
		}
		if f := typ.Field(28); f.Name != "InsecureSkipVerify" || f.Offset != 384 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 28: %+v", f)
		}
		if f := typ.Field(29); f.Name != "CipherSuites" || f.Offset != 392 || f.Type.String() != "[]uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 29: %+v", f)
		}
		if f := typ.Field(30); f.Name != "PreferServerCipherSuites" || f.Offset != 416 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 30: %+v", f)
		}
		if f := typ.Field(31); f.Name != "SessionTicketsDisabled" || f.Offset != 417 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 31: %+v", f)
		}
		if f := typ.Field(32); f.Name != "SessionTicketKey" || f.Offset != 418 || f.Type.String() != "[32]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 32: %+v", f)
		}
		if f := typ.Field(33); f.Name != "ClientSessionCache" || f.Offset != 456 || f.Type.String() != "reality.ClientSessionCache" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 33: %+v", f)
		}
		if f := typ.Field(34); f.Name != "UnwrapSession" || f.Offset != 472 || f.Type.String() != "func([]uint8, reality.ConnectionState) (*reality.SessionState, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 34: %+v", f)
		}
		if f := typ.Field(35); f.Name != "WrapSession" || f.Offset != 480 || f.Type.String() != "func(reality.ConnectionState, *reality.SessionState) ([]uint8, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 35: %+v", f)
		}
		if f := typ.Field(36); f.Name != "MinVersion" || f.Offset != 488 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 36: %+v", f)
		}
		if f := typ.Field(37); f.Name != "MaxVersion" || f.Offset != 490 || f.Type.String() != "uint16" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 37: %+v", f)
		}
		if f := typ.Field(38); f.Name != "CurvePreferences" || f.Offset != 496 || f.Type.String() != "[]reality.CurveID" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 38: %+v", f)
		}
		if f := typ.Field(39); f.Name != "DynamicRecordSizingDisabled" || f.Offset != 520 || f.Type.String() != "bool" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 39: %+v", f)
		}
		if f := typ.Field(40); f.Name != "Renegotiation" || f.Offset != 528 || f.Type.String() != "reality.RenegotiationSupport" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 40: %+v", f)
		}
		if f := typ.Field(41); f.Name != "KeyLogWriter" || f.Offset != 536 || f.Type.String() != "io.Writer" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 41: %+v", f)
		}
		if f := typ.Field(42); f.Name != "EncryptedClientHelloConfigList" || f.Offset != 552 || f.Type.String() != "[]uint8" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 42: %+v", f)
		}
		if f := typ.Field(43); f.Name != "EncryptedClientHelloRejectionVerify" || f.Offset != 576 || f.Type.String() != "func(reality.ConnectionState) error" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 43: %+v", f)
		}
		if f := typ.Field(44); f.Name != "GetEncryptedClientHelloKeys" || f.Offset != 584 || f.Type.String() != "func(*reality.ClientHelloInfo) ([]reality.EncryptedClientHelloKey, error)" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 44: %+v", f)
		}
		if f := typ.Field(45); f.Name != "EncryptedClientHelloKeys" || f.Offset != 592 || f.Type.String() != "[]reality.EncryptedClientHelloKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 45: %+v", f)
		}
		if f := typ.Field(46); f.Name != "mutex" || f.Offset != 616 || f.Type.String() != "sync.RWMutex" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 46: %+v", f)
		}
		if f := typ.Field(47); f.Name != "sessionTicketKeys" || f.Offset != 640 || f.Type.String() != "[]reality.ticketKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 47: %+v", f)
		}
		if f := typ.Field(48); f.Name != "autoSessionTicketKeys" || f.Offset != 664 || f.Type.String() != "[]reality.ticketKey" || f.Anonymous != false || string(f.Tag) != "" {
			t.Errorf("Config field 48: %+v", f)
		}
	}
}
