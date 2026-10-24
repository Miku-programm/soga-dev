package mieru

import "fmt"

type cKThZd6LHk4 struct{}

func (cKThZd6LHk4) Error() string   { return "i/o timeout" }
func (cKThZd6LHk4) Timeout() bool   { return true }
func (cKThZd6LHk4) Temporary() bool { return true }
func (k xofvu_nD97c) String() string {
	switch k {
	case 0:
		return "closeConnRequest"
	case 1:
		return "closeConnResponse"
	case 2:
		return "openSessionRequest"
	case 3:
		return "openSessionResponse"
	case 4:
		return "closeSessionRequest"
	case 5:
		return "closeSessionResponse"
	case 6:
		return "dataClientToServer"
	case 7:
		return "dataServerToClient"
	case 8:
		return "ackClientToServer"
	case 9:
		return "ackServerToClient"
	default:
		return fmt.Sprintf("UNKNOWN(%v)", byte(k))
	}
}

func (m *IqB1foHWg) String() string { //0x1c3a280
	switch m.Rf6N7kY6GVCs | 1 {
	case 3, 5:
		return fmt.Sprintf("{%v sid=%d seq=%d status=%d payloadLen=%d suffixLen=%d}", m.Rf6N7kY6GVCs, m.HNaaDh, m.Tt1OfazA, m.DGMUmJ7, m.HPUJVQ0S, m.FO7MEJ)
	default:
		return fmt.Sprintf("{%v sid=%d seq=%d unAck=%d win=%d frag=%d prefixLen=%d payloadLen=%d suffixLen=%d}", m.Rf6N7kY6GVCs, m.HNaaDh, m.Tt1OfazA, m.N01v7C, m.UlMT6MeaV2, m.VGN7gjQ4p, m.B8QsA0N2yo, m.HPUJVQ0S, m.FO7MEJ)
	}
}
