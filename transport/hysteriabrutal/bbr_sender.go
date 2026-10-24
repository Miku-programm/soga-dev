// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package hysteriabrutal

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/apernet/quic-go/congestion"
	common "recovered.local/soga/transport/hysteriaerr"
)

// BbrSender implements BBR congestion control algorithm.  BBR aims to estimate
// the current available Bottleneck Bandwidth and RTT (hence the name), and
// regulates the pacing rate and the size of the congestion window based on
// those signals.
//
// BBR relies on pacing in order to function properly.  Do not use BBR when
// pacing is disabled.
//

const (
	minBps = 65536 // 64 kbps

	invalidPacketNumber            = -1
	initialCongestionWindowPackets = 32

	// Constants based on TCP defaults.
	// The minimum CWND to ensure delayed acks don't reduce bandwidth measurements.
	// Does not inflate the pacing rate.
	defaultMinimumCongestionWindow = 4 * congestion.ByteCount(congestion.InitialPacketSizeIPv4)

	// The gain used for the STARTUP, equal to 2/ln(2).
	defaultHighGain = 2.885
	// The newly derived gain for STARTUP, equal to 4 * ln(2)
	derivedHighGain = 2.773
	// The newly derived CWND gain for STARTUP, 2.
	derivedHighCWNDGain = 2.0

	debugEnv = "HYSTERIA_BBR_DEBUG"
)

// The cycle of gains used during the PROBE_BW stage.
var pacingGain = [...]float64{1.25, 0.75, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}

const (
	// The length of the gain cycle.
	gainCycleLength = len(pacingGain)
	// The size of the bandwidth filter window, in round-trips.
	bandwidthWindowSize = gainCycleLength + 2

	// The time after which the current min_rtt value expires.
	minRttExpiry = 10 * time.Second
	// The minimum time the connection can spend in PROBE_RTT mode.
	probeRttTime = 200 * time.Millisecond
	// If the bandwidth does not increase by the factor of |kStartupGrowthTarget|
	// within |kRoundTripsWithoutGrowthBeforeExitingStartup| rounds, the connection
	// will exit the STARTUP mode.
	startupGrowthTarget                         = 1.25
	roundTripsWithoutGrowthBeforeExitingStartup = int64(3)

	// Flag.
	defaultStartupFullLossCount  = 8
	quicBbr2DefaultLossThreshold = 0.02
	maxBbrBurstPackets           = 10
)

type lajiJA int

const (
	// Startup phase of the connection.
	bbrModeStartup = iota
	// After achieving the highest possible bandwidth during the startup, lower
	// the pacing rate in order to drain the queue.
	bbrModeDrain
	// Cruising mode.
	bbrModeProbeBw
	// Temporarily slow down sending in order to empty the buffer and measure
	// the real minimum RTT.
	bbrModeProbeRtt
)

// Indicates how the congestion control limits the amount of bytes in flight.
type hKvOmYf_pbF int

const (
	// Do not limit.
	bbrRecoveryStateNotInRecovery = iota
	// Allow an extra outstanding byte for each byte acknowledged.
	bbrRecoveryStateConservation
	// Allow two extra outstanding bytes for each byte acknowledged (slow
	// start).
	bbrRecoveryStateGrowth
)

type tGYu23kout struct {
	jzNuIRaj5L4f congestion.RTTStatsProvider
	ajdznEwg3HP  KYybyfkN
	gMj1ag2rOX9T *common.OQpLDX

	jOTShB lajiJA

	// Bandwidth sampler provides BBR with the bandwidth measurements at
	// individual points.
	eiff4Dk6J *oahYXcYd4

	// The number of the round trips that have occurred during the connection.
	g4B8Dg awuEkJg9p

	// The packet number of the most recently sent packet.
	wUnF0r congestion.PacketNumber
	// Acknowledgement of any packet after |current_round_trip_end_| will cause
	// the round trip counter to advance.
	bkPDow congestion.PacketNumber

	// Number of congestion events with some losses, in the current round.
	pC_xa8 uint64

	// Number of total bytes lost in the current round.
	sSPtxam congestion.ByteCount

	// The filter that tracks the maximum bandwidth over the multiple recent
	// round-trips.
	clm34LfHOfH *Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]

	// Minimum RTT estimate.  Automatically expires within 10 seconds (and
	// triggers PROBE_RTT mode) if no new value is sampled during that period.
	es2Rg_vDy time.Duration
	// The time at which the current value of |min_rtt_| was assigned.
	p4pnrwiNi time.Time

	// The maximum allowed number of bytes in flight.
	ndBPgZ congestion.ByteCount

	// The initial value of the |congestion_window_|.
	kPUkZLy congestion.ByteCount

	// The largest value the |congestion_window_| can achieve.
	qkdjbFU0l congestion.ByteCount

	// The smallest value the |congestion_window_| can achieve.
	irwaYXRFoja congestion.ByteCount

	// The pacing gain applied during the STARTUP phase.
	ytPxW1_vL float64

	// The CWND gain applied during the STARTUP phase.
	fI9EUh float64

	// The pacing gain applied during the DRAIN phase.
	s_Zx6iZ float64

	// The current pacing rate of the connection.
	rbHB4J1ksIV JaadYOR_Vaf6

	// The gain currently applied to the pacing rate.
	rDtgaAw float64
	// The gain currently applied to the congestion window.
	qAZ13bY float64

	// The gain used for the congestion window during PROBE_BW.  Latched from
	// quic_bbr_cwnd_gain flag.
	sZxgbLDX float64
	// The number of RTTs to stay in STARTUP mode.  Defaults to 3.
	xIZ4Qm6J81 int64

	// Number of round-trips in PROBE_BW mode, used for determining the current
	// pacing gain cycle.
	sewKJeM4 int
	// The time at which the last pacing gain cycle was started.
	_62XJdnnNP time.Time

	// Indicates whether the connection has reached the full bandwidth mode.
	dv_p1GSp bool
	// Number of rounds during which there was no significant bandwidth increase.
	ccZF6DrRpP int64
	// The bandwidth compared to which the increase is measured.
	afwE420 JaadYOR_Vaf6

	// Set to true upon exiting quiescence.
	bmSW0VMoFYsi bool

	// Time at which PROBE_RTT has to be exited.  Setting it to zero indicates
	// that the time is yet unknown as the number of packets in flight has not
	// reached the required value.
	hq2wUhQa time.Time
	// Indicates whether a round-trip has passed since PROBE_RTT became active.
	zFN_SDSij bool

	// Indicates whether the most recent bandwidth sample was marked as
	// app-limited.
	gZ285RGO1K bool
	// Indicates whether any non app-limited samples have been recorded.
	hcFC7avZ bool

	// Current state of recovery.
	daiTa6d hKvOmYf_pbF
	// Receiving acknowledgement of a packet after |end_recovery_at_| will cause
	// BBR to exit the recovery mode.  A value above zero indicates at least one
	// loss has been detected, so it must not be set back to zero.
	zdCu53B_Jn congestion.PacketNumber
	// A window used to limit the number of bytes in flight during loss recovery.
	voFtPR2C congestion.ByteCount
	// If true, consider all samples in recovery app-limited.
	xTTndm bool // not used

	// When true, pace at 1.5x and disable packet conservation in STARTUP.
	dxYh9Z bool // not used
	// When true, disables packet conservation in STARTUP.
	i5pDu7Gmkf bool // not used

	// When true, add the most recent ack aggregation measurement during STARTUP.
	zQhgaatIp bool
	// When true, expire the windowed ack aggregation values in STARTUP when
	// bandwidth increases more than 25%.
	f4pwsMIQjjC bool

	// If true, will not exit low gain mode until bytes_in_flight drops below BDP
	// or it's time for high gain mode.
	btA4YPFSwnmT bool

	// If true, slow down pacing rate in STARTUP when overshooting is detected.
	yL5qU8a bool
	// Bytes lost while detect_overshooting_ is true.
	wnghShWZg congestion.ByteCount
	// Slow down pacing rate if
	// bytes_lost_while_detecting_overshooting_ *
	// bytes_lost_multiplier_while_detecting_overshooting_ > IW.
	gX1RE3uW uint8
	// When overshooting is detected, do not drop pacing_rate_ below this value /
	// min_rtt.
	zaXoVmSmYwgs congestion.ByteCount

	// Max congestion window when adjusting network parameters.
	dzxASp congestion.ByteCount // not used

	// Params.
	hIPu5Tz9 congestion.ByteCount
	// Recorded on packet sent. equivalent |unacked_packets_->bytes_in_flight()|
	ehQ4ZmGJ4u congestion.ByteCount

	d2Q76d8Uzv8K bool
}

var _ congestion.CongestionControl = &tGYu23kout{}

func NewSenderFromInline(
	clock KYybyfkN, initialMaxDatagramSize congestion.ByteCount,
) *tGYu23kout {
	return bmvaaWF(
		clock,
		initialMaxDatagramSize,
		initialCongestionWindowPackets*initialMaxDatagramSize,
		congestion.MaxCongestionWindowPackets*initialMaxDatagramSize,
	)
}

func bmvaaWF(
	clock KYybyfkN, initialMaxDatagramSize,
	initialCongestionWindow,
	initialMaxCongestionWindow congestion.ByteCount,
) *tGYu23kout {
	debug, _ := strconv.ParseBool(os.Getenv(debugEnv))
	b := &tGYu23kout{
		ajdznEwg3HP:  clock,
		jOTShB:       bbrModeStartup,
		eiff4Dk6J:    iTmwEHfnE0Fe(awuEkJg9p(bandwidthWindowSize)),
		wUnF0r:       invalidPacketNumber,
		bkPDow:       invalidPacketNumber,
		clm34LfHOfH:  NewWindowedFilter(awuEkJg9p(bandwidthWindowSize), Ez2Vlj[JaadYOR_Vaf6]),
		ndBPgZ:       initialCongestionWindow,
		kPUkZLy:      initialCongestionWindow,
		qkdjbFU0l:    initialMaxCongestionWindow,
		irwaYXRFoja:  defaultMinimumCongestionWindow,
		ytPxW1_vL:    defaultHighGain,
		fI9EUh:       defaultHighGain,
		s_Zx6iZ:      1.0 / defaultHighGain,
		rDtgaAw:      1.0,
		qAZ13bY:      1.0,
		sZxgbLDX:     2.0,
		xIZ4Qm6J81:   roundTripsWithoutGrowthBeforeExitingStartup,
		daiTa6d:      bbrRecoveryStateNotInRecovery,
		zdCu53B_Jn:   invalidPacketNumber,
		voFtPR2C:     initialMaxCongestionWindow,
		gX1RE3uW:     2,
		zaXoVmSmYwgs: initialCongestionWindow,
		dzxASp:       initialMaxCongestionWindow,
		hIPu5Tz9:     initialMaxDatagramSize,
		d2Q76d8Uzv8K: debug,
	}
	b.gMj1ag2rOX9T = common.NewFromInline(b.nDThIaL)

	/*
		if b.tracer != nil {
			b.lastState = logging.CongestionStateStartup
			b.tracer.UpdatedCongestionState(logging.CongestionStateStartup)
		}
	*/

	b.rUoHx19(b.ajdznEwg3HP.Now())
	b.setHighCwndGain(derivedHighCWNDGain)

	return b
}

func (b *tGYu23kout) SetRTTStatsProvider(provider congestion.RTTStatsProvider) {
	b.jzNuIRaj5L4f = provider
}

// TimeUntilSend implements the SendAlgorithm interface.
func (b *tGYu23kout) TimeUntilSend(bytesInFlight congestion.ByteCount) time.Time {
	return b.gMj1ag2rOX9T.TimeUntilSend()
}

// HasPacingBudget implements the SendAlgorithm interface.
func (b *tGYu23kout) HasPacingBudget(now time.Time) bool {
	return b.gMj1ag2rOX9T.Budget(now) >= b.hIPu5Tz9
}

// OnPacketSent implements the SendAlgorithm interface.
func (b *tGYu23kout) OnPacketSent(
	sentTime time.Time,
	bytesInFlight congestion.ByteCount,
	packetNumber congestion.PacketNumber,
	bytes congestion.ByteCount,
	isRetransmittable bool,
) {
	b.gMj1ag2rOX9T.SentPacket(sentTime, bytes)

	b.wUnF0r = packetNumber
	b.ehQ4ZmGJ4u = bytesInFlight

	if bytesInFlight == 0 {
		b.bmSW0VMoFYsi = true
	}

	b.eiff4Dk6J.OnPacketSent(sentTime, packetNumber, bytes, bytesInFlight, isRetransmittable)
}

// CanSend implements the SendAlgorithm interface.
func (b *tGYu23kout) CanSend(bytesInFlight congestion.ByteCount) bool {
	return bytesInFlight < b.GetCongestionWindow()
}

// MaybeExitSlowStart implements the SendAlgorithm interface.
func (b *tGYu23kout) MaybeExitSlowStart() {
	// Do nothing
}

// OnPacketAcked implements the SendAlgorithm interface.
func (b *tGYu23kout) OnPacketAcked(number congestion.PacketNumber, ackedBytes, priorInFlight congestion.ByteCount, eventTime time.Time) {
	// Do nothing.
}

// OnPacketLost implements the SendAlgorithm interface.
func (b *tGYu23kout) OnPacketLost(number congestion.PacketNumber, lostBytes, priorInFlight congestion.ByteCount) {
	// Do nothing.
}

// OnRetransmissionTimeout implements the SendAlgorithm interface.
func (b *tGYu23kout) OnRetransmissionTimeout(packetsRetransmitted bool) {
	// Do nothing.
}

// SetMaxDatagramSize implements the SendAlgorithm interface.
func (b *tGYu23kout) SetMaxDatagramSize(s congestion.ByteCount) {
	if s < b.hIPu5Tz9 {
		panic(fmt.Sprintf("congestion BUG: decreased max datagram size from %d to %d", b.hIPu5Tz9, s))
	}
	cwndIsMinCwnd := b.ndBPgZ == b.irwaYXRFoja
	b.hIPu5Tz9 = s
	if cwndIsMinCwnd {
		b.ndBPgZ = b.irwaYXRFoja
	}
	b.gMj1ag2rOX9T.SetMaxDatagramSize(s)
}

// InSlowStart implements the SendAlgorithmWithDebugInfos interface.
func (b *tGYu23kout) InSlowStart() bool {
	return b.jOTShB == bbrModeStartup
}

// InRecovery implements the SendAlgorithmWithDebugInfos interface.
func (b *tGYu23kout) InRecovery() bool {
	return b.daiTa6d != bbrRecoveryStateNotInRecovery
}

// GetCongestionWindow implements the SendAlgorithmWithDebugInfos interface.
func (b *tGYu23kout) GetCongestionWindow() congestion.ByteCount {
	if b.jOTShB == bbrModeProbeRtt {
		return b.probeRttCongestionWindow()
	}

	if b.InRecovery() {
		return min(b.ndBPgZ, b.voFtPR2C)
	}

	return b.ndBPgZ
}

func (b *tGYu23kout) OnCongestionEvent(number congestion.PacketNumber, lostBytes, priorInFlight congestion.ByteCount) {
	// Do nothing.
}

func (b *tGYu23kout) OnCongestionEventEx(priorInFlight congestion.ByteCount, eventTime time.Time, ackedPackets []congestion.AckedPacketInfo, lostPackets []congestion.LostPacketInfo) {
	totalBytesAckedBefore := b.eiff4Dk6J.TotalBytesAcked()
	totalBytesLostBefore := b.eiff4Dk6J.TotalBytesLost()

	var isRoundStart, minRttExpired bool
	var excessAcked, bytesLost congestion.ByteCount

	// The send state of the largest packet in acked_packets, unless it is
	// empty. If acked_packets is empty, it's the send state of the largest
	// packet in lost_packets.
	var lastPacketSendState iC6_vHy61ee

	b.maybeAppLimited(priorInFlight)

	// Update bytesInFlight
	b.ehQ4ZmGJ4u = priorInFlight
	for _, p := range ackedPackets {
		b.ehQ4ZmGJ4u -= p.BytesAcked
	}
	for _, p := range lostPackets {
		b.ehQ4ZmGJ4u -= p.BytesLost
	}

	if len(ackedPackets) != 0 {
		lastAckedPacket := ackedPackets[len(ackedPackets)-1].PacketNumber
		isRoundStart = b.updateRoundTripCounter(lastAckedPacket)
		b.updateRecoveryState(lastAckedPacket, len(lostPackets) != 0, isRoundStart)
	}

	sample := b.eiff4Dk6J.OnCongestionEvent(eventTime,
		ackedPackets, lostPackets, b.clm34LfHOfH.GetBest(), infBandwidth, b.g4B8Dg)
	if sample.xDGjNq0R.o00cUt {
		b.gZ285RGO1K = sample.xDGjNq0R.rcyLtXR67hl
		b.hcFC7avZ = b.hcFC7avZ || !b.gZ285RGO1K
	}
	// Avoid updating |max_bandwidth_| if a) this is a loss-only event, or b) all
	// packets in |acked_packets| did not generate valid samples. (e.g. ack of
	// ack-only packets). In both cases, sampler_.total_bytes_acked() will not
	// change.
	if totalBytesAckedBefore != b.eiff4Dk6J.TotalBytesAcked() {
		if !sample.zkhqa7g1Vp3 || sample.w87e3ZbfBKJR > b.clm34LfHOfH.GetBest() {
			b.clm34LfHOfH.Update(sample.w87e3ZbfBKJR, b.g4B8Dg)
		}
	}

	if sample.e1LhJmPmT != infRTT {
		minRttExpired = b.a4jZwIN(eventTime, sample.e1LhJmPmT)
	}
	bytesLost = b.eiff4Dk6J.TotalBytesLost() - totalBytesLostBefore

	excessAcked = sample.kPsWl9
	lastPacketSendState = sample.xDGjNq0R

	if len(lostPackets) != 0 {
		b.pC_xa8++
		b.sSPtxam += bytesLost
	}

	// Handle logic specific to PROBE_BW mode.
	if b.jOTShB == bbrModeProbeBw {
		b.naN0aU3Oq(eventTime, priorInFlight, len(lostPackets) != 0)
	}

	// Handle logic specific to STARTUP and DRAIN modes.
	if isRoundStart && !b.dv_p1GSp {
		b.i7czOLhu3mnJ(&lastPacketSendState)
	}

	b.iMbUKxNV3Xp(eventTime)

	// Handle logic specific to PROBE_RTT.
	b.te9aUC2_(eventTime, isRoundStart, minRttExpired)

	// Calculate number of packets acked and lost.
	bytesAcked := b.eiff4Dk6J.TotalBytesAcked() - totalBytesAckedBefore

	// After the model is updated, recalculate the pacing rate and congestion
	// window.
	b.n0_hVN2CE0Lj(bytesLost)
	b.cr9DME(bytesAcked, excessAcked)
	b.calculateRecoveryWindow(bytesAcked, bytesLost)

	// Cleanup internal state.
	// This is where we clean up obsolete (acked or lost) packets from the bandwidth sampler.
	// The "least unacked" should actually be FirstOutstanding, but since we are not passing
	// that through OnCongestionEventEx, we will only do an estimate using acked/lost packets
	// for now. Because of fast retransmission, they should differ by no more than 2 packets.
	// (this is controlled by packetThreshold in quic-go's sentPacketHandler)
	var leastUnacked congestion.PacketNumber
	if len(ackedPackets) != 0 {
		leastUnacked = ackedPackets[len(ackedPackets)-1].PacketNumber - 2
	} else {
		leastUnacked = lostPackets[len(lostPackets)-1].PacketNumber + 1
	}
	b.eiff4Dk6J.RemoveObsoletePackets(leastUnacked)

	if isRoundStart {
		b.pC_xa8 = 0
		b.sSPtxam = 0
	}
}

func (b *tGYu23kout) PacingRate() JaadYOR_Vaf6 {
	if b.rbHB4J1ksIV == 0 {
		return JaadYOR_Vaf6(b.ytPxW1_vL * float64(
			BandwidthFromDelta(b.kPUkZLy, b.yZgIw31())))
	}

	return b.rbHB4J1ksIV
}

// Sets the pacing gain used in STARTUP.  Must be greater than 1.

// Sets the CWND gain used in STARTUP.  Must be greater than 1.
func (b *tGYu23kout) setHighCwndGain(highCwndGain float64) {
	b.fI9EUh = highCwndGain
	if b.jOTShB == bbrModeStartup {
		b.qAZ13bY = highCwndGain
	}
}

// Sets the gain used in DRAIN.  Must be less than 1.

// Get the current bandwidth estimate. Note that Bandwidth is in bits per second.
func (b *tGYu23kout) bandwidthEstimate() JaadYOR_Vaf6 {
	return b.clm34LfHOfH.GetBest()
}

func (b *tGYu23kout) nDThIaL() congestion.ByteCount {
	bps := congestion.ByteCount(float64(b.bandwidthEstimate()) * b.qAZ13bY / float64(BytesPerSecond))
	if bps < minBps {
		// We need to make sure that the bandwidth value for pacer is never zero,
		// otherwise it will go into an edge case where HasPacingBudget = false
		// but TimeUntilSend is before, causing the quic-go send loop to go crazy and get stuck.
		return minBps
	}
	return bps
}

// Returns the current estimate of the RTT of the connection.  Outside of the
// edge cases, this is minimum RTT.
func (b *tGYu23kout) yZgIw31() time.Duration {
	if b.es2Rg_vDy != 0 {
		return b.es2Rg_vDy
	}
	// min_rtt could be available if the handshake packet gets neutered then
	// gets acknowledged. This could only happen for QUIC crypto where we do not
	// drop keys.
	minRtt := b.jzNuIRaj5L4f.MinRTT()
	if minRtt == 0 {
		return 100 * time.Millisecond
	} else {
		return minRtt
	}
}

// Computes the target congestion window using the specified gain.
func (b *tGYu23kout) iIG_jAPC3U(gain float64) congestion.ByteCount {
	bdp := bdpFromRttAndBandwidth(b.yZgIw31(), b.bandwidthEstimate())
	congestionWindow := congestion.ByteCount(gain * float64(bdp))

	// BDP estimate will be zero if no bandwidth samples are available yet.
	if congestionWindow == 0 {
		congestionWindow = congestion.ByteCount(gain * float64(b.kPUkZLy))
	}

	return max(congestionWindow, b.irwaYXRFoja)
}

// The target congestion window during PROBE_RTT.
func (b *tGYu23kout) probeRttCongestionWindow() congestion.ByteCount {
	return b.irwaYXRFoja
}

func (b *tGYu23kout) a4jZwIN(now time.Time, sampleMinRtt time.Duration) bool {
	// Do not expire min_rtt if none was ever available.
	minRttExpired := b.es2Rg_vDy != 0 && now.After(b.p4pnrwiNi.Add(minRttExpiry))
	if minRttExpired || sampleMinRtt < b.es2Rg_vDy || b.es2Rg_vDy == 0 {
		b.es2Rg_vDy = sampleMinRtt
		b.p4pnrwiNi = now
	}

	return minRttExpired
}

// Enters the STARTUP mode.
func (b *tGYu23kout) rUoHx19(now time.Time) {
	b.jOTShB = bbrModeStartup
	// b.maybeTraceStateChange(logging.CongestionStateStartup)
	b.rDtgaAw = b.ytPxW1_vL
	b.qAZ13bY = b.fI9EUh

	if b.d2Q76d8Uzv8K {
		b.zS48CWRdj("Phase: STARTUP")
	}
}

// Enters the PROBE_BW mode.
func (b *tGYu23kout) zMter8W0(now time.Time) {
	b.jOTShB = bbrModeProbeBw
	// b.maybeTraceStateChange(logging.CongestionStateProbeBw)
	b.qAZ13bY = b.sZxgbLDX

	// Pick a random offset for the gain cycle out of {0, 2..7} range. 1 is
	// excluded because in that case increased gain and decreased gain would not
	// follow each other.
	b.sewKJeM4 = int(rand.Int31n(congestion.PacketsPerConnectionID)) % (gainCycleLength - 1)
	if b.sewKJeM4 >= 1 {
		b.sewKJeM4 += 1
	}

	b._62XJdnnNP = now
	b.rDtgaAw = pacingGain[b.sewKJeM4]

	if b.d2Q76d8Uzv8K {
		b.zS48CWRdj("Phase: PROBE_BW")
	}
}

// Updates the round-trip counter if a round-trip has passed.  Returns true if
// the counter has been advanced.
func (b *tGYu23kout) updateRoundTripCounter(lastAckedPacket congestion.PacketNumber) bool {
	if b.bkPDow == invalidPacketNumber || lastAckedPacket > b.bkPDow {
		b.g4B8Dg++
		b.bkPDow = b.wUnF0r
		return true
	}
	return false
}

// Updates the current gain used in PROBE_BW mode.
func (b *tGYu23kout) naN0aU3Oq(now time.Time, priorInFlight congestion.ByteCount, hasLosses bool) {
	// In most cases, the cycle is advanced after an RTT passes.
	shouldAdvanceGainCycling := now.After(b._62XJdnnNP.Add(b.yZgIw31()))
	// If the pacing gain is above 1.0, the connection is trying to probe the
	// bandwidth by increasing the number of bytes in flight to at least
	// pacing_gain * BDP.  Make sure that it actually reaches the target, as long
	// as there are no losses suggesting that the buffers are not able to hold
	// that much.
	if b.rDtgaAw > 1.0 && !hasLosses && priorInFlight < b.iIG_jAPC3U(b.rDtgaAw) {
		shouldAdvanceGainCycling = false
	}

	// If pacing gain is below 1.0, the connection is trying to drain the extra
	// queue which could have been incurred by probing prior to it.  If the number
	// of bytes in flight falls down to the estimated BDP value earlier, conclude
	// that the queue has been successfully drained and exit this cycle early.
	if b.rDtgaAw < 1.0 && b.ehQ4ZmGJ4u <= b.iIG_jAPC3U(1) {
		shouldAdvanceGainCycling = true
	}

	if shouldAdvanceGainCycling {
		b.sewKJeM4 = (b.sewKJeM4 + 1) % gainCycleLength
		b._62XJdnnNP = now
		// Stay in low gain mode until the target BDP is hit.
		// Low gain mode will be exited immediately when the target BDP is achieved.
		if b.btA4YPFSwnmT && b.rDtgaAw < 1 &&
			pacingGain[b.sewKJeM4] == 1 &&
			b.ehQ4ZmGJ4u > b.iIG_jAPC3U(1) {
			return
		}
		b.rDtgaAw = pacingGain[b.sewKJeM4]
	}
}

// Tracks for how many round-trips the bandwidth has not increased
// significantly.
func (b *tGYu23kout) i7czOLhu3mnJ(lastPacketSendState *iC6_vHy61ee) {
	if b.gZ285RGO1K {
		return
	}

	target := JaadYOR_Vaf6(float64(b.afwE420) * startupGrowthTarget)
	if b.bandwidthEstimate() >= target {
		b.afwE420 = b.bandwidthEstimate()
		b.ccZF6DrRpP = 0
		if b.f4pwsMIQjjC {
			// Expire old excess delivery measurements now that bandwidth increased.
			b.eiff4Dk6J.ResetMaxAckHeightTracker(0, b.g4B8Dg)
		}
		return
	}

	b.ccZF6DrRpP++
	if b.ccZF6DrRpP >= b.xIZ4Qm6J81 ||
		b.shouldExitStartupDueToLoss(lastPacketSendState) {
		b.dv_p1GSp = true
	}
}

func (b *tGYu23kout) maybeAppLimited(bytesInFlight congestion.ByteCount) {
	if bytesInFlight < b.iIG_jAPC3U(1) {
		b.eiff4Dk6J.OnAppLimited()
	}
}

// Transitions from STARTUP to DRAIN and from DRAIN to PROBE_BW if
// appropriate.
func (b *tGYu23kout) iMbUKxNV3Xp(now time.Time) {
	if b.jOTShB == bbrModeStartup && b.dv_p1GSp {
		b.jOTShB = bbrModeDrain
		// b.maybeTraceStateChange(logging.CongestionStateDrain)
		b.rDtgaAw = b.s_Zx6iZ
		b.qAZ13bY = b.fI9EUh

		if b.d2Q76d8Uzv8K {
			b.zS48CWRdj("Phase: DRAIN")
		}
	}
	if b.jOTShB == bbrModeDrain && b.ehQ4ZmGJ4u <= b.iIG_jAPC3U(1) {
		b.zMter8W0(now)
	}
}

// Decides whether to enter or exit PROBE_RTT.
func (b *tGYu23kout) te9aUC2_(now time.Time, isRoundStart, minRttExpired bool) {
	if minRttExpired && !b.bmSW0VMoFYsi && b.jOTShB != bbrModeProbeRtt {
		b.jOTShB = bbrModeProbeRtt
		// b.maybeTraceStateChange(logging.CongestionStateProbRtt)
		b.rDtgaAw = 1.0
		// Do not decide on the time to exit PROBE_RTT until the |bytes_in_flight|
		// is at the target small value.
		b.hq2wUhQa = time.Time{}

		if b.d2Q76d8Uzv8K {
			b.zS48CWRdj("BandwidthEstimate: %s, CongestionWindowGain: %.2f, PacingGain: %.2f, PacingRate: %s",
				bijN9oB(b.bandwidthEstimate()), b.qAZ13bY, b.rDtgaAw, bijN9oB(b.PacingRate()))
			b.zS48CWRdj("Phase: PROBE_RTT")
		}
	}

	if b.jOTShB == bbrModeProbeRtt {
		b.eiff4Dk6J.OnAppLimited()
		// b.maybeTraceStateChange(logging.CongestionStateApplicationLimited)

		if b.hq2wUhQa.IsZero() {
			// If the window has reached the appropriate size, schedule exiting
			// PROBE_RTT.  The CWND during PROBE_RTT is kMinimumCongestionWindow, but
			// we allow an extra packet since QUIC checks CWND before sending a
			// packet.
			if b.ehQ4ZmGJ4u < b.probeRttCongestionWindow()+congestion.MaxPacketBufferSize {
				b.hq2wUhQa = now.Add(probeRttTime)
				b.zFN_SDSij = false
			}
		} else {
			if isRoundStart {
				b.zFN_SDSij = true
			}
			if now.Sub(b.hq2wUhQa) >= 0 && b.zFN_SDSij {
				b.p4pnrwiNi = now
				if b.d2Q76d8Uzv8K {
					b.zS48CWRdj("MinRTT: %s", b.yZgIw31())
				}
				if !b.dv_p1GSp {
					b.rUoHx19(now)
				} else {
					b.zMter8W0(now)
				}
			}
		}
	}

	b.bmSW0VMoFYsi = false
}

// Determines whether BBR needs to enter, exit or advance state of the
// recovery.
func (b *tGYu23kout) updateRecoveryState(lastAckedPacket congestion.PacketNumber, hasLosses, isRoundStart bool) {
	// Disable recovery in startup, if loss-based exit is enabled.
	if !b.dv_p1GSp {
		return
	}

	// Exit recovery when there are no losses for a round.
	if hasLosses {
		b.zdCu53B_Jn = b.wUnF0r
	}

	switch b.daiTa6d {
	case bbrRecoveryStateNotInRecovery:
		if hasLosses {
			b.daiTa6d = bbrRecoveryStateConservation
			// This will cause the |recovery_window_| to be set to the correct
			// value in CalculateRecoveryWindow().
			b.voFtPR2C = 0
			// Since the conservation phase is meant to be lasting for a whole
			// round, extend the current round as if it were started right now.
			b.bkPDow = b.wUnF0r
		}
	case bbrRecoveryStateConservation:
		if isRoundStart {
			b.daiTa6d = bbrRecoveryStateGrowth
		}
		fallthrough
	case bbrRecoveryStateGrowth:
		// Exit recovery if appropriate.
		if !hasLosses && lastAckedPacket > b.zdCu53B_Jn {
			b.daiTa6d = bbrRecoveryStateNotInRecovery
		}
	}
}

// Determines the appropriate pacing rate for the connection.
func (b *tGYu23kout) n0_hVN2CE0Lj(bytesLost congestion.ByteCount) {
	if b.bandwidthEstimate() == 0 {
		return
	}

	targetRate := JaadYOR_Vaf6(b.rDtgaAw * float64(b.bandwidthEstimate()))
	if b.dv_p1GSp {
		b.rbHB4J1ksIV = targetRate
		return
	}

	// Pace at the rate of initial_window / RTT as soon as RTT measurements are
	// available.
	if b.rbHB4J1ksIV == 0 && b.jzNuIRaj5L4f.MinRTT() != 0 {
		b.rbHB4J1ksIV = BandwidthFromDelta(b.kPUkZLy, b.jzNuIRaj5L4f.MinRTT())
		return
	}

	if b.yL5qU8a {
		b.wnghShWZg += bytesLost
		// Check for overshooting with network parameters adjusted when pacing rate
		// > target_rate and loss has been detected.
		if b.rbHB4J1ksIV > targetRate && b.wnghShWZg > 0 {
			if b.hcFC7avZ ||
				b.wnghShWZg*congestion.ByteCount(b.gX1RE3uW) > b.kPUkZLy {
				// We are fairly sure overshoot happens if 1) there is at least one
				// non app-limited bw sample or 2) half of IW gets lost. Slow pacing
				// rate.
				b.rbHB4J1ksIV = max(targetRate, BandwidthFromDelta(b.zaXoVmSmYwgs, b.jzNuIRaj5L4f.MinRTT()))
				b.wnghShWZg = 0
				b.yL5qU8a = false
			}
		}
	}

	// Do not decrease the pacing rate during startup.
	b.rbHB4J1ksIV = max(b.rbHB4J1ksIV, targetRate)
}

// Determines the appropriate congestion window for the connection.
func (b *tGYu23kout) cr9DME(bytesAcked, excessAcked congestion.ByteCount) {
	if b.jOTShB == bbrModeProbeRtt {
		return
	}

	targetWindow := b.iIG_jAPC3U(b.qAZ13bY)
	if b.dv_p1GSp {
		// Add the max recently measured ack aggregation to CWND.
		targetWindow += b.eiff4Dk6J.MaxAckHeight()
	} else if b.zQhgaatIp {
		// Add the most recent excess acked.  Because CWND never decreases in
		// STARTUP, this will automatically create a very localized max filter.
		targetWindow += excessAcked
	}

	// Instead of immediately setting the target CWND as the new one, BBR grows
	// the CWND towards |target_window| by only increasing it |bytes_acked| at a
	// time.
	if b.dv_p1GSp {
		b.ndBPgZ = min(targetWindow, b.ndBPgZ+bytesAcked)
	} else if b.ndBPgZ < targetWindow ||
		b.eiff4Dk6J.TotalBytesAcked() < b.kPUkZLy {
		// If the connection is not yet out of startup phase, do not decrease the
		// window.
		b.ndBPgZ += bytesAcked
	}

	// Enforce the limits on the congestion window.
	b.ndBPgZ = max(b.ndBPgZ, b.irwaYXRFoja)
	b.ndBPgZ = min(b.ndBPgZ, b.qkdjbFU0l)
}

// Determines the appropriate window that constrains the in-flight during recovery.
func (b *tGYu23kout) calculateRecoveryWindow(bytesAcked, bytesLost congestion.ByteCount) {
	if b.daiTa6d == bbrRecoveryStateNotInRecovery {
		return
	}

	// Set up the initial recovery window.
	if b.voFtPR2C == 0 {
		b.voFtPR2C = b.ehQ4ZmGJ4u + bytesAcked
		b.voFtPR2C = max(b.irwaYXRFoja, b.voFtPR2C)
		return
	}

	// Remove losses from the recovery window, while accounting for a potential
	// integer underflow.
	if b.voFtPR2C >= bytesLost {
		b.voFtPR2C = b.voFtPR2C - bytesLost
	} else {
		b.voFtPR2C = b.hIPu5Tz9
	}

	// In CONSERVATION mode, just subtracting losses is sufficient.  In GROWTH,
	// release additional |bytes_acked| to achieve a slow-start-like behavior.
	if b.daiTa6d == bbrRecoveryStateGrowth {
		b.voFtPR2C += bytesAcked
	}

	// Always allow sending at least |bytes_acked| in response.
	b.voFtPR2C = max(b.voFtPR2C, b.ehQ4ZmGJ4u+bytesAcked)
	b.voFtPR2C = max(b.irwaYXRFoja, b.voFtPR2C)
}

// Return whether we should exit STARTUP due to excessive loss.
func (b *tGYu23kout) shouldExitStartupDueToLoss(lastPacketSendState *iC6_vHy61ee) bool {
	if b.pC_xa8 < defaultStartupFullLossCount || !lastPacketSendState.o00cUt {
		return false
	}

	inflightAtSend := lastPacketSendState.tcFhoDpK1YNk

	if inflightAtSend > 0 && b.sSPtxam > 0 {
		if b.sSPtxam > congestion.ByteCount(float64(inflightAtSend)*quicBbr2DefaultLossThreshold) {
			return true
		}
		return false
	}
	return false
}

func (b *tGYu23kout) zS48CWRdj(format string, a ...any) {
	fmt.Printf("[BBRSender] [%s] %s\n",
		time.Now().Format("15:04:05"),
		fmt.Sprintf(format, a...))
}

func bdpFromRttAndBandwidth(rtt time.Duration, bandwidth JaadYOR_Vaf6) congestion.ByteCount {
	return congestion.ByteCount(rtt) * congestion.ByteCount(bandwidth) / congestion.ByteCount(BytesPerSecond) / congestion.ByteCount(time.Second)
}

func CMryyxB(addr net.Addr) congestion.ByteCount {
	// If this is not a UDP address, we don't know anything about the MTU.
	// Use the minimum size of an Initial packet as the max packet size.
	if udpAddr, ok := addr.(*net.UDPAddr); ok {
		if udpAddr.IP.To4() != nil {
			return congestion.InitialPacketSizeIPv4
		} else {
			return congestion.InitialPacketSizeIPv6
		}
	} else {
		return congestion.MinInitialPacketSize
	}
}

func bijN9oB(bw JaadYOR_Vaf6) string {
	bwf := float64(bw)
	units := []string{"bps", "Kbps", "Mbps", "Gbps"}
	unitIndex := 0
	for bwf > 1000 && unitIndex < len(units)-1 {
		bwf /= 1000
		unitIndex++
	}
	return fmt.Sprintf("%.2f %s", bwf, units[unitIndex])
}
