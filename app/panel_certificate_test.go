package app

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"testing"
)

type startupCertificatePanel struct {
	certificateNodeFixture
	id int
}

func (p *startupCertificatePanel) GetAPIConfig() *C.FCEAfQJwDA7e {
	return &C.FCEAfQJwDA7e{Yf1KqSqJpNI7: p.id}
}

func TestPanelCertificateSkipsMissingContentAndPreservesWriteFailure(t *testing.T) {
	oldCert, oldKey := C.ManualCertificateFiles()
	t.Cleanup(func() { C.SetCertificateFiles(oldCert, oldKey) })
	C.SetCertificateFiles("existing.crt", "existing.key")
	plain := &certificatePanelFixture{}
	missing := &startupCertificatePanel{certificateNodeFixture: certificateNodeFixture{certificatePanelFixture: certificatePanelFixture{node: &L.Wr5SVP{I3k4JnNp7nB: "one half"}}}, id: 127}
	invalid := &startupCertificatePanel{certificateNodeFixture: certificateNodeFixture{certificatePanelFixture: certificatePanelFixture{node: &L.Wr5SVP{EMlXEc: 128, I3k4JnNp7nB: "invalid certificate", XATiVun: "invalid key"}}}, id: 128}
	after := &startupCertificatePanel{certificateNodeFixture: certificateNodeFixture{certificatePanelFixture: certificatePanelFixture{node: &L.Wr5SVP{}}}, id: 129}
	if !gJwSzy([]Q.VyFk5uT{plain, missing, invalid, after}) {
		t.Fatal("original returns true even when saving panel certificate fails")
	}
	if plain.calls != 0 || missing.nodeCalls != 1 || invalid.nodeCalls != 1 || after.nodeCalls != 0 {
		t.Fatal("optional interface/iteration/early return mismatch")
	}
	if cert, key := C.ManualCertificateFiles(); cert != "existing.crt" || key != "existing.key" {
		t.Fatal("invalid certificate changed configured paths")
	}
	if gJwSzy([]Q.VyFk5uT{plain, missing}) || gJwSzy(nil) {
		t.Fatal("empty certificate content must return false")
	}
}
