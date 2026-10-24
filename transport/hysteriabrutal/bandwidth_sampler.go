// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package hysteriabrutal

import (
	"math"
	"time"

	"github.com/apernet/quic-go/congestion"
)

const (
	infRTT                             = time.Duration(math.MaxInt64)
	defaultConnectionStateMapQueueSize = 256
	defaultCandidatesBufferSize        = 256
)

type awuEkJg9p uint64

// SendTimeState is a subset of ConnectionStateOnSentPacket which is returned
// to the caller when the packet is acked or lost.
type iC6_vHy61ee struct {
	// Whether other states in this object is valid.
	o00cUt bool
	// Whether the sender is app limited at the time the packet was sent.
	// App limited bandwidth sample might be artificially low because the sender
	// did not have enough data to send in order to saturate the link.
	rcyLtXR67hl bool
	// Total number of sent bytes at the time the packet was sent.
	// Includes the packet itself.
	fyuyH6zvX congestion.ByteCount
	// Total number of acked bytes at the time the packet was sent.
	gvnxB1rdE congestion.ByteCount
	// Total number of lost bytes at the time the packet was sent.
	pWaFRsxfpvC congestion.ByteCount
	// Total number of inflight bytes at the time the packet was sent.
	// Includes the packet itself.
	// It should be equal to |total_bytes_sent| minus the sum of
	// |total_bytes_acked|, |total_bytes_lost| and total neutered bytes.
	tcFhoDpK1YNk congestion.ByteCount
}

func newSendTimeState(
	isAppLimited bool,
	totalBytesSent congestion.ByteCount,
	totalBytesAcked congestion.ByteCount,
	totalBytesLost congestion.ByteCount,
	bytesInFlight congestion.ByteCount,
) *iC6_vHy61ee {
	return &iC6_vHy61ee{
		o00cUt:       true,
		rcyLtXR67hl:  isAppLimited,
		fyuyH6zvX:    totalBytesSent,
		gvnxB1rdE:    totalBytesAcked,
		pWaFRsxfpvC:  totalBytesLost,
		tcFhoDpK1YNk: bytesInFlight,
	}
}

type k8ofg5VH struct {
	// The excess bytes acknowlwedged in the time delta for this event.
	sCcx3Ca6E congestion.ByteCount

	// The bytes acknowledged and time delta from the event.
	yhpZvxZ congestion.ByteCount
	tSOB2v  time.Duration
	// The round trip of the event.
	qr1Cd4llNbRm awuEkJg9p
}

func mUaFqjDBO(a, b k8ofg5VH) int {
	if a.sCcx3Ca6E > b.sCcx3Ca6E {
		return 1
	} else if a.sCcx3Ca6E < b.sCcx3Ca6E {
		return -1
	}
	return 0
}

// BandwidthSample
type bandwidthSample struct {
	// The bandwidth at that particular sample. Zero if no valid bandwidth sample
	// is available.
	bandwidth JaadYOR_Vaf6
	// The RTT measurement at this particular sample.  Zero if no RTT sample is
	// available.  Does not correct for delayed ack time.
	rtt time.Duration
	// |send_rate| is computed from the current packet being acked('P') and an
	// earlier packet that is acked before P was sent.
	sendRate JaadYOR_Vaf6
	// States captured when the packet was sent.
	stateAtSend iC6_vHy61ee
}

func newBandwidthSample() *bandwidthSample {
	return &bandwidthSample{
		sendRate: infBandwidth,
	}
}

// MaxAckHeightTracker is part of the BandwidthSampler. It is called after every
// ack event to keep track the degree of ack aggregation(a.k.a "ack height").
type a4a_ewpUFjg struct {
	// Tracks the maximum number of bytes acked faster than the estimated
	// bandwidth.
	jm_R6V *Nqk7YFvzgcZg[k8ofg5VH, awuEkJg9p]
	// The time this aggregation started and the number of bytes acked during it.
	r10Waeh time.Time
	cM5vf8  congestion.ByteCount
	// The last sent packet number before the current aggregation epoch started.
	pLwLGJtFf congestion.PacketNumber
	// The number of ack aggregation epochs ever started, including the ongoing
	// one. Stats only.
	uxti2K1    uint64
	v9KunT     float64
	tpXKRSNXF  bool
	sI7jVWaahw bool
}

func newMaxAckHeightTracker(windowLength awuEkJg9p) *a4a_ewpUFjg {
	return &a4a_ewpUFjg{
		jm_R6V:    NewWindowedFilter(windowLength, mUaFqjDBO),
		pLwLGJtFf: invalidPacketNumber,
		v9KunT:    1.0,
	}
}

func (m *a4a_ewpUFjg) Get() congestion.ByteCount {
	return m.jm_R6V.GetBest().sCcx3Ca6E
}

func (m *a4a_ewpUFjg) Update(
	bandwidthEstimate JaadYOR_Vaf6, isNewMaxBandwidth bool,
	roundTripCount awuEkJg9p,
	lastSentPacketNumber congestion.PacketNumber,
	lastAckedPacketNumber congestion.PacketNumber,
	ackTime time.Time,
	bytesAcked congestion.ByteCount,
) congestion.ByteCount {
	forceNewEpoch := false

	if m.sI7jVWaahw && isNewMaxBandwidth {
		// Save and clear existing entries.
		best := m.jm_R6V.GetBest()
		secondBest := m.jm_R6V.GetSecondBest()
		thirdBest := m.jm_R6V.GetThirdBest()
		m.jm_R6V.Clear()

		// Reinsert the heights into the filter after recalculating.
		expectedBytesAcked := bytesFromBandwidthAndTimeDelta(bandwidthEstimate, best.tSOB2v)
		if expectedBytesAcked < best.yhpZvxZ {
			best.sCcx3Ca6E = best.yhpZvxZ - expectedBytesAcked
			m.jm_R6V.Update(best, best.qr1Cd4llNbRm)
		}
		expectedBytesAcked = bytesFromBandwidthAndTimeDelta(bandwidthEstimate, secondBest.tSOB2v)
		if expectedBytesAcked < secondBest.yhpZvxZ {
			secondBest.sCcx3Ca6E = secondBest.yhpZvxZ - expectedBytesAcked
			m.jm_R6V.Update(secondBest, secondBest.qr1Cd4llNbRm)
		}
		expectedBytesAcked = bytesFromBandwidthAndTimeDelta(bandwidthEstimate, thirdBest.tSOB2v)
		if expectedBytesAcked < thirdBest.yhpZvxZ {
			thirdBest.sCcx3Ca6E = thirdBest.yhpZvxZ - expectedBytesAcked
			m.jm_R6V.Update(thirdBest, thirdBest.qr1Cd4llNbRm)
		}
	}

	// If any packet sent after the start of the epoch has been acked, start a new
	// epoch.
	if m.tpXKRSNXF &&
		m.pLwLGJtFf != invalidPacketNumber &&
		lastAckedPacketNumber != invalidPacketNumber &&
		lastAckedPacketNumber > m.pLwLGJtFf {
		forceNewEpoch = true
	}
	if m.r10Waeh.IsZero() || forceNewEpoch {
		m.cM5vf8 = bytesAcked
		m.r10Waeh = ackTime
		m.pLwLGJtFf = lastSentPacketNumber
		m.uxti2K1++
		return 0
	}

	// Compute how many bytes are expected to be delivered, assuming max bandwidth
	// is correct.
	aggregationDelta := ackTime.Sub(m.r10Waeh)
	expectedBytesAcked := bytesFromBandwidthAndTimeDelta(bandwidthEstimate, aggregationDelta)
	// Reset the current aggregation epoch as soon as the ack arrival rate is less
	// than or equal to the max bandwidth.
	if m.cM5vf8 <= congestion.ByteCount(m.v9KunT*float64(expectedBytesAcked)) {
		// Reset to start measuring a new aggregation epoch.
		m.cM5vf8 = bytesAcked
		m.r10Waeh = ackTime
		m.pLwLGJtFf = lastSentPacketNumber
		m.uxti2K1++
		return 0
	}

	m.cM5vf8 += bytesAcked

	// Compute how many extra bytes were delivered vs max bandwidth.
	extraBytesAcked := m.cM5vf8 - expectedBytesAcked
	newEvent := k8ofg5VH{
		sCcx3Ca6E: expectedBytesAcked,
		yhpZvxZ:   m.cM5vf8,
		tSOB2v:    aggregationDelta,
	}
	m.jm_R6V.Update(newEvent, roundTripCount)
	return extraBytesAcked
}

func (m *a4a_ewpUFjg) SetFilterWindowLength(length awuEkJg9p) {
	m.jm_R6V.SetWindowLength(length)
}

func (m *a4a_ewpUFjg) Reset(newHeight congestion.ByteCount, newTime awuEkJg9p) {
	newEvent := k8ofg5VH{
		sCcx3Ca6E:    newHeight,
		qr1Cd4llNbRm: newTime,
	}
	m.jm_R6V.Reset(newEvent, newTime)
}

func (m *a4a_ewpUFjg) SetAckAggregationBandwidthThreshold(threshold float64) {
	m.v9KunT = threshold
}

func (m *a4a_ewpUFjg) SetStartNewAggregationEpochAfterFullRound(value bool) {
	m.tpXKRSNXF = value
}

func (m *a4a_ewpUFjg) SetReduceExtraAckedOnBandwidthIncrease(value bool) {
	m.sI7jVWaahw = value
}

func (m *a4a_ewpUFjg) AckAggregationBandwidthThreshold() float64 {
	return m.v9KunT
}

func (m *a4a_ewpUFjg) NumAckAggregationEpochs() uint64 {
	return m.uxti2K1
}

// AckPoint represents a point on the ack line.
type kO8aHEaU9 struct {
	m7dTK3NJaO time.Time
	joNWs9     congestion.ByteCount
}

// RecentAckPoints maintains the most recent 2 ack points at distinct times.
type h4dvYPRUDTrt struct {
	yCiDnZ [2]kO8aHEaU9
}

func (r *h4dvYPRUDTrt) Update(ackTime time.Time, totalBytesAcked congestion.ByteCount) {
	if ackTime.Before(r.yCiDnZ[1].m7dTK3NJaO) {
		r.yCiDnZ[1].m7dTK3NJaO = ackTime
	} else if ackTime.After(r.yCiDnZ[1].m7dTK3NJaO) {
		r.yCiDnZ[0] = r.yCiDnZ[1]
		r.yCiDnZ[1].m7dTK3NJaO = ackTime
	}

	r.yCiDnZ[1].joNWs9 = totalBytesAcked
}

func (r *h4dvYPRUDTrt) Clear() {
	r.yCiDnZ[0] = kO8aHEaU9{}
	r.yCiDnZ[1] = kO8aHEaU9{}
}

func (r *h4dvYPRUDTrt) MostRecentPoint() *kO8aHEaU9 {
	return &r.yCiDnZ[1]
}

func (r *h4dvYPRUDTrt) LessRecentPoint() *kO8aHEaU9 {
	if r.yCiDnZ[0].joNWs9 != 0 {
		return &r.yCiDnZ[0]
	}

	return &r.yCiDnZ[1]
}

// ConnectionStateOnSentPacket represents the information about a sent packet
// and the state of the connection at the moment the packet was sent,
// specifically the information about the most recently acknowledged packet at
// that moment.
type wzBRgOvr8BC struct {
	// Time at which the packet is sent.
	p0vSfsOlU time.Time
	// Size of the packet.
	jNhPqIq congestion.ByteCount
	// The value of |totalBytesSentAtLastAckedPacket| at the time the
	// packet was sent.
	pLa7WgMaenyW congestion.ByteCount
	// The value of |lastAckedPacketSentTime| at the time the packet was
	// sent.
	v_nXPcs0HTe time.Time
	// The value of |lastAckedPacketAckTime| at the time the packet was
	// sent.
	yXecWTdyLG time.Time
	// Send time states that are returned to the congestion controller when the
	// packet is acked or lost.
	ujKE8LV1OP iC6_vHy61ee
}

// Snapshot constructor. Records the current state of the bandwidth
// sampler.
// |bytes_in_flight| is the bytes in flight right after the packet is sent.
func newConnectionStateOnSentPacket(
	sentTime time.Time,
	size congestion.ByteCount,
	bytesInFlight congestion.ByteCount,
	sampler *oahYXcYd4,
) *wzBRgOvr8BC {
	return &wzBRgOvr8BC{
		p0vSfsOlU:    sentTime,
		jNhPqIq:      size,
		pLa7WgMaenyW: sampler.g4GkKaB7,
		v_nXPcs0HTe:  sampler.fhJyzuFBO,
		yXecWTdyLG:   sampler.ciarX4f,
		ujKE8LV1OP: *newSendTimeState(
			sampler.q7YCyTJjNN,
			sampler.eoyKL70a,
			sampler.cyAQ1b,
			sampler.oZZJNF,
			bytesInFlight,
		),
	}
}

// BandwidthSampler keeps track of sent and acknowledged packets and outputs a
// bandwidth sample for every packet acknowledged. The samples are taken for
// individual packets, and are not filtered; the consumer has to filter the
// bandwidth samples itself. In certain cases, the sampler will locally severely
// underestimate the bandwidth, hence a maximum filter with a size of at least
// one RTT is recommended.
//
// This class bases its samples on the slope of two curves: the number of bytes
// sent over time, and the number of bytes acknowledged as received over time.
// It produces a sample of both slopes for every packet that gets acknowledged,
// based on a slope between two points on each of the corresponding curves. Note
// that due to the packet loss, the number of bytes on each curve might get
// further and further away from each other, meaning that it is not feasible to
// compare byte values coming from different curves with each other.
//
// The obvious points for measuring slope sample are the ones corresponding to
// the packet that was just acknowledged. Let us denote them as S_1 (point at
// which the current packet was sent) and A_1 (point at which the current packet
// was acknowledged). However, taking a slope requires two points on each line,
// so estimating bandwidth requires picking a packet in the past with respect to
// which the slope is measured.
//
// For that purpose, BandwidthSampler always keeps track of the most recently
// acknowledged packet, and records it together with every outgoing packet.
// When a packet gets acknowledged (A_1), it has not only information about when
// it itself was sent (S_1), but also the information about the latest
// acknowledged packet right before it was sent (S_0 and A_0).
//
// Based on that data, send and ack rate are estimated as:
//
//	send_rate = (bytes(S_1) - bytes(S_0)) / (time(S_1) - time(S_0))
//	ack_rate = (bytes(A_1) - bytes(A_0)) / (time(A_1) - time(A_0))
//
// Here, the ack rate is intuitively the rate we want to treat as bandwidth.
// However, in certain cases (e.g. ack compression) the ack rate at a point may
// end up higher than the rate at which the data was originally sent, which is
// not indicative of the real bandwidth. Hence, we use the send rate as an upper
// bound, and the sample value is
//
//	rate_sample = min(send_rate, ack_rate)
//
// An important edge case handled by the sampler is tracking the app-limited
// samples. There are multiple meaning of "app-limited" used interchangeably,
// hence it is important to understand and to be able to distinguish between
// them.
//
// Meaning 1: connection state. The connection is said to be app-limited when
// there is no outstanding data to send. This means that certain bandwidth
// samples in the future would not be an accurate indication of the link
// capacity, and it is important to inform consumer about that. Whenever
// connection becomes app-limited, the sampler is notified via OnAppLimited()
// method.
//
// Meaning 2: a phase in the bandwidth sampler. As soon as the bandwidth
// sampler becomes notified about the connection being app-limited, it enters
// app-limited phase. In that phase, all *sent* packets are marked as
// app-limited. Note that the connection itself does not have to be
// app-limited during the app-limited phase, and in fact it will not be
// (otherwise how would it send packets?). The boolean flag below indicates
// whether the sampler is in that phase.
//
// Meaning 3: a flag on the sent packet and on the sample. If a sent packet is
// sent during the app-limited phase, the resulting sample related to the
// packet will be marked as app-limited.
//
// With the terminology issue out of the way, let us consider the question of
// what kind of situation it addresses.
//
// Consider a scenario where we first send packets 1 to 20 at a regular
// bandwidth, and then immediately run out of data. After a few seconds, we send
// packets 21 to 60, and only receive ack for 21 between sending packets 40 and
// 41. In this case, when we sample bandwidth for packets 21 to 40, the S_0/A_0
// we use to compute the slope is going to be packet 20, a few seconds apart
// from the current packet, hence the resulting estimate would be extremely low
// and not indicative of anything. Only at packet 41 the S_0/A_0 will become 21,
// meaning that the bandwidth sample would exclude the quiescence.
//
// Based on the analysis of that scenario, we implement the following rule: once
// OnAppLimited() is called, all sent packets will produce app-limited samples
// up until an ack for a packet that was sent after OnAppLimited() was called.
// Note that while the scenario above is not the only scenario when the
// connection is app-limited, the approach works in other cases too.

type ri4dMyMP struct {
	// The maximum bandwidth sample from all acked packets.
	// QuicBandwidth::Zero() if no samples are available.
	w87e3ZbfBKJR JaadYOR_Vaf6
	// Whether |sample_max_bandwidth| is from a app-limited sample.
	zkhqa7g1Vp3 bool
	// The minimum rtt sample from all acked packets.
	// QuicTime::Delta::Infinite() if no samples are available.
	e1LhJmPmT time.Duration
	// For each packet p in acked packets, this is the max value of INFLIGHT(p),
	// where INFLIGHT(p) is the number of bytes acked while p is inflight.
	jSNh_B5Te congestion.ByteCount
	// The send state of the largest packet in acked_packets, unless it is
	// empty. If acked_packets is empty, it's the send state of the largest
	// packet in lost_packets.
	xDGjNq0R iC6_vHy61ee
	// The number of extra bytes acked from this ack event, compared to what is
	// expected from the flow's bandwidth. Larger value means more ack
	// aggregation.
	kPsWl9 congestion.ByteCount
}

func newCongestionEventSample() *ri4dMyMP {
	return &ri4dMyMP{
		e1LhJmPmT: infRTT,
	}
}

type oahYXcYd4 struct {
	// The total number of congestion controlled bytes sent during the connection.
	eoyKL70a congestion.ByteCount

	// The total number of congestion controlled bytes which were acknowledged.
	cyAQ1b congestion.ByteCount

	// The total number of congestion controlled bytes which were lost.
	oZZJNF congestion.ByteCount

	// The total number of congestion controlled bytes which have been neutered.
	o8yuvw congestion.ByteCount

	// The value of |total_bytes_sent_| at the time the last acknowledged packet
	// was sent. Valid only when |last_acked_packet_sent_time_| is valid.
	g4GkKaB7 congestion.ByteCount

	// The time at which the last acknowledged packet was sent. Set to
	// QuicTime::Zero() if no valid timestamp is available.
	fhJyzuFBO time.Time

	// The time at which the most recent packet was acknowledged.
	ciarX4f time.Time

	// The most recently sent packet.
	n7S7Dan congestion.PacketNumber

	// The most recently acked packet.
	fhILy8i congestion.PacketNumber

	// Indicates whether the bandwidth sampler is currently in an app-limited
	// phase.
	q7YCyTJjNN bool

	// The packet that will be acknowledged after this one will cause the sampler
	// to exit the app-limited phase.
	zISoP8ufhtI congestion.PacketNumber

	// Record of the connection state at the point where each packet in flight was
	// sent, indexed by the packet number.
	e7KrTFhqp *h8SYlleQ[wzBRgOvr8BC]

	y_Z4NrytDBe4 h4dvYPRUDTrt
	vFlrYPDp4    YXV6CyZ5Ux[kO8aHEaU9]

	// Maximum number of tracked packets.
	rDjgJc congestion.ByteCount

	qgtQO9f7j *a4a_ewpUFjg
	lWRr8Cm0w congestion.ByteCount

	// True if connection option 'BSAO' is set.
	i4NwjwX59Gl bool

	// True if connection option 'BBRB' is set.
	exqrreSd1pd bool
}

func iTmwEHfnE0Fe(maxAckHeightTrackerWindowLength awuEkJg9p) *oahYXcYd4 {
	b := &oahYXcYd4{
		qgtQO9f7j:   newMaxAckHeightTracker(maxAckHeightTrackerWindowLength),
		e7KrTFhqp:   newPacketNumberIndexedQueue[wzBRgOvr8BC](defaultConnectionStateMapQueueSize),
		n7S7Dan:     invalidPacketNumber,
		fhILy8i:     invalidPacketNumber,
		zISoP8ufhtI: invalidPacketNumber,
	}

	b.vFlrYPDp4.Init(defaultCandidatesBufferSize)

	return b
}

func (b *oahYXcYd4) MaxAckHeight() congestion.ByteCount {
	return b.qgtQO9f7j.Get()
}

func (b *oahYXcYd4) NumAckAggregationEpochs() uint64 {
	return b.qgtQO9f7j.NumAckAggregationEpochs()
}

func (b *oahYXcYd4) SetMaxAckHeightTrackerWindowLength(length awuEkJg9p) {
	b.qgtQO9f7j.SetFilterWindowLength(length)
}

func (b *oahYXcYd4) ResetMaxAckHeightTracker(newHeight congestion.ByteCount, newTime awuEkJg9p) {
	b.qgtQO9f7j.Reset(newHeight, newTime)
}

func (b *oahYXcYd4) SetStartNewAggregationEpochAfterFullRound(value bool) {
	b.qgtQO9f7j.SetStartNewAggregationEpochAfterFullRound(value)
}

func (b *oahYXcYd4) SetLimitMaxAckHeightTrackerBySendRate(value bool) {
	b.exqrreSd1pd = value
}

func (b *oahYXcYd4) SetReduceExtraAckedOnBandwidthIncrease(value bool) {
	b.qgtQO9f7j.SetReduceExtraAckedOnBandwidthIncrease(value)
}

func (b *oahYXcYd4) EnableOverestimateAvoidance() {
	if b.i4NwjwX59Gl {
		return
	}

	b.i4NwjwX59Gl = true
	b.qgtQO9f7j.SetAckAggregationBandwidthThreshold(2.0)
}

func (b *oahYXcYd4) IsOverestimateAvoidanceEnabled() bool {
	return b.i4NwjwX59Gl
}

func (b *oahYXcYd4) OnPacketSent(
	sentTime time.Time,
	packetNumber congestion.PacketNumber,
	bytes congestion.ByteCount,
	bytesInFlight congestion.ByteCount,
	isRetransmittable bool,
) {
	b.n7S7Dan = packetNumber

	if !isRetransmittable {
		return
	}

	b.eoyKL70a += bytes

	// If there are no packets in flight, the time at which the new transmission
	// opens can be treated as the A_0 point for the purpose of bandwidth
	// sampling. This underestimates bandwidth to some extent, and produces some
	// artificially low samples for most packets in flight, but it provides with
	// samples at important points where we would not have them otherwise, most
	// importantly at the beginning of the connection.
	if bytesInFlight == 0 {
		b.ciarX4f = sentTime
		if b.i4NwjwX59Gl {
			b.y_Z4NrytDBe4.Clear()
			b.y_Z4NrytDBe4.Update(sentTime, b.cyAQ1b)
			b.vFlrYPDp4.Clear()
			b.vFlrYPDp4.PushBack(*b.y_Z4NrytDBe4.MostRecentPoint())
		}
		b.g4GkKaB7 = b.eoyKL70a

		// In this situation ack compression is not a concern, set send rate to
		// effectively infinite.
		b.fhJyzuFBO = sentTime
	}

	b.e7KrTFhqp.Emplace(packetNumber, newConnectionStateOnSentPacket(
		sentTime,
		bytes,
		bytesInFlight+bytes,
		b,
	))
}

func (b *oahYXcYd4) OnCongestionEvent(
	ackTime time.Time,
	ackedPackets []congestion.AckedPacketInfo,
	lostPackets []congestion.LostPacketInfo,
	maxBandwidth JaadYOR_Vaf6, estBandwidthUpperBound JaadYOR_Vaf6, roundTripCount awuEkJg9p,
) ri4dMyMP {
	eventSample := newCongestionEventSample()

	var lastLostPacketSendState iC6_vHy61ee

	for _, p := range lostPackets {
		sendState := b.OnPacketLost(p.PacketNumber, p.BytesLost)
		if sendState.o00cUt {
			lastLostPacketSendState = sendState
		}
	}

	if len(ackedPackets) == 0 {
		// Only populate send state for a loss-only event.
		eventSample.xDGjNq0R = lastLostPacketSendState
		return *eventSample
	}

	var lastAckedPacketSendState iC6_vHy61ee
	var maxSendRate JaadYOR_Vaf6

	for _, p := range ackedPackets {
		sample := b.nNZK6dj0MNti(ackTime, p.PacketNumber)
		if !sample.stateAtSend.o00cUt {
			continue
		}

		lastAckedPacketSendState = sample.stateAtSend

		if sample.rtt != 0 {
			eventSample.e1LhJmPmT = min(eventSample.e1LhJmPmT, sample.rtt)
		}
		if sample.bandwidth > eventSample.w87e3ZbfBKJR {
			eventSample.w87e3ZbfBKJR = sample.bandwidth
			eventSample.zkhqa7g1Vp3 = sample.stateAtSend.rcyLtXR67hl
		}
		if sample.sendRate != infBandwidth {
			maxSendRate = max(maxSendRate, sample.sendRate)
		}
		inflightSample := b.cyAQ1b - lastAckedPacketSendState.gvnxB1rdE
		if inflightSample > eventSample.jSNh_B5Te {
			eventSample.jSNh_B5Te = inflightSample
		}
	}

	if !lastLostPacketSendState.o00cUt {
		eventSample.xDGjNq0R = lastAckedPacketSendState
	} else if !lastAckedPacketSendState.o00cUt {
		eventSample.xDGjNq0R = lastLostPacketSendState
	} else {
		// If two packets are inflight and an alarm is armed to lose a packet and it
		// wakes up late, then the first of two in flight packets could have been
		// acknowledged before the wakeup, which re-evaluates loss detection, and
		// could declare the later of the two lost.
		if lostPackets[len(lostPackets)-1].PacketNumber > ackedPackets[len(ackedPackets)-1].PacketNumber {
			eventSample.xDGjNq0R = lastLostPacketSendState
		} else {
			eventSample.xDGjNq0R = lastAckedPacketSendState
		}
	}

	isNewMaxBandwidth := eventSample.w87e3ZbfBKJR > maxBandwidth
	maxBandwidth = max(maxBandwidth, eventSample.w87e3ZbfBKJR)
	if b.exqrreSd1pd {
		maxBandwidth = max(maxBandwidth, maxSendRate)
	}

	eventSample.kPsWl9 = b.av4xVCcxem5(min(estBandwidthUpperBound, maxBandwidth), isNewMaxBandwidth, roundTripCount)

	return *eventSample
}

func (b *oahYXcYd4) OnPacketLost(packetNumber congestion.PacketNumber, bytesLost congestion.ByteCount) (s iC6_vHy61ee) {
	b.oZZJNF += bytesLost
	if sentPacketPointer := b.e7KrTFhqp.GetEntry(packetNumber); sentPacketPointer != nil {
		sentPacketToSendTimeState(sentPacketPointer, &s)
	}
	return s
}

func (b *oahYXcYd4) OnPacketNeutered(packetNumber congestion.PacketNumber) {
	b.e7KrTFhqp.Remove(packetNumber, func(sentPacket wzBRgOvr8BC) {
		b.o8yuvw += sentPacket.jNhPqIq
	})
}

func (b *oahYXcYd4) OnAppLimited() {
	b.q7YCyTJjNN = true
	b.zISoP8ufhtI = b.n7S7Dan
}

func (b *oahYXcYd4) RemoveObsoletePackets(leastUnacked congestion.PacketNumber) {
	// A packet can become obsolete when it is removed from QuicUnackedPacketMap's
	// view of inflight before it is acked or marked as lost. For example, when
	// QuicSentPacketManager::RetransmitCryptoPackets retransmits a crypto packet,
	// the packet is removed from QuicUnackedPacketMap's inflight, but is not
	// marked as acked or lost in the BandwidthSampler.
	b.e7KrTFhqp.RemoveUpTo(leastUnacked)
}

func (b *oahYXcYd4) TotalBytesSent() congestion.ByteCount {
	return b.eoyKL70a
}

func (b *oahYXcYd4) TotalBytesLost() congestion.ByteCount {
	return b.oZZJNF
}

func (b *oahYXcYd4) TotalBytesAcked() congestion.ByteCount {
	return b.cyAQ1b
}

func (b *oahYXcYd4) TotalBytesNeutered() congestion.ByteCount {
	return b.o8yuvw
}

func (b *oahYXcYd4) IsAppLimited() bool {
	return b.q7YCyTJjNN
}

func (b *oahYXcYd4) EndOfAppLimitedPhase() congestion.PacketNumber {
	return b.zISoP8ufhtI
}

func (b *oahYXcYd4) bxaXSIOnnib(totalBytesAcked congestion.ByteCount, a0 *kO8aHEaU9) bool {
	if b.vFlrYPDp4.Empty() {
		return false
	}

	if b.vFlrYPDp4.Len() == 1 {
		*a0 = *b.vFlrYPDp4.Front()
		return true
	}

	for i := 1; i < b.vFlrYPDp4.Len(); i++ {
		if b.vFlrYPDp4.Offset(i).joNWs9 > totalBytesAcked {
			*a0 = *b.vFlrYPDp4.Offset(i - 1)
			if i > 1 {
				for j := 0; j < i-1; j++ {
					b.vFlrYPDp4.PopFront()
				}
			}
			return true
		}
	}

	*a0 = *b.vFlrYPDp4.Back()
	for k := 0; k < b.vFlrYPDp4.Len()-1; k++ {
		b.vFlrYPDp4.PopFront()
	}
	return true
}

func (b *oahYXcYd4) nNZK6dj0MNti(ackTime time.Time, packetNumber congestion.PacketNumber) bandwidthSample {
	sample := newBandwidthSample()
	b.fhILy8i = packetNumber
	sentPacketPointer := b.e7KrTFhqp.GetEntry(packetNumber)
	if sentPacketPointer == nil {
		return *sample
	}

	// OnPacketAcknowledgedInner
	b.cyAQ1b += sentPacketPointer.jNhPqIq
	b.g4GkKaB7 = sentPacketPointer.ujKE8LV1OP.fyuyH6zvX
	b.fhJyzuFBO = sentPacketPointer.p0vSfsOlU
	b.ciarX4f = ackTime
	if b.i4NwjwX59Gl {
		b.y_Z4NrytDBe4.Update(ackTime, b.cyAQ1b)
	}

	if b.q7YCyTJjNN {
		// Exit app-limited phase in two cases:
		// (1) end_of_app_limited_phase_ is not initialized, i.e., so far all
		// packets are sent while there are buffered packets or pending data.
		// (2) The current acked packet is after the sent packet marked as the end
		// of the app limit phase.
		if b.zISoP8ufhtI == invalidPacketNumber ||
			packetNumber > b.zISoP8ufhtI {
			b.q7YCyTJjNN = false
		}
	}

	// There might have been no packets acknowledged at the moment when the
	// current packet was sent. In that case, there is no bandwidth sample to
	// make.
	if sentPacketPointer.v_nXPcs0HTe.IsZero() {
		return *sample
	}

	// Infinite rate indicates that the sampler is supposed to discard the
	// current send rate sample and use only the ack rate.
	sendRate := infBandwidth
	if sentPacketPointer.p0vSfsOlU.After(sentPacketPointer.v_nXPcs0HTe) {
		sendRate = BandwidthFromDelta(
			sentPacketPointer.ujKE8LV1OP.fyuyH6zvX-sentPacketPointer.pLa7WgMaenyW,
			sentPacketPointer.p0vSfsOlU.Sub(sentPacketPointer.v_nXPcs0HTe))
	}

	var a0 kO8aHEaU9
	if b.i4NwjwX59Gl && b.bxaXSIOnnib(sentPacketPointer.ujKE8LV1OP.gvnxB1rdE, &a0) {
	} else {
		a0.m7dTK3NJaO = sentPacketPointer.yXecWTdyLG
		a0.joNWs9 = sentPacketPointer.ujKE8LV1OP.gvnxB1rdE
	}

	// During the slope calculation, ensure that ack time of the current packet is
	// always larger than the time of the previous packet, otherwise division by
	// zero or integer underflow can occur.
	if ackTime.Sub(a0.m7dTK3NJaO) <= 0 {
		return *sample
	}

	ackRate := BandwidthFromDelta(b.cyAQ1b-a0.joNWs9, ackTime.Sub(a0.m7dTK3NJaO))

	sample.bandwidth = min(sendRate, ackRate)
	// Note: this sample does not account for delayed acknowledgement time.  This
	// means that the RTT measurements here can be artificially high, especially
	// on low bandwidth connections.
	sample.rtt = ackTime.Sub(sentPacketPointer.p0vSfsOlU)
	sample.sendRate = sendRate
	sentPacketToSendTimeState(sentPacketPointer, &sample.stateAtSend)

	return *sample
}

func (b *oahYXcYd4) av4xVCcxem5(
	bandwidthEstimate JaadYOR_Vaf6, isNewMaxBandwidth bool,
	roundTripCount awuEkJg9p,
) congestion.ByteCount {
	newlyAckedBytes := b.cyAQ1b - b.lWRr8Cm0w
	if newlyAckedBytes == 0 {
		return 0
	}
	b.lWRr8Cm0w = b.cyAQ1b
	extraAcked := b.qgtQO9f7j.Update(
		bandwidthEstimate,
		isNewMaxBandwidth,
		roundTripCount,
		b.n7S7Dan,
		b.fhILy8i,
		b.ciarX4f,
		newlyAckedBytes)
	// If |extra_acked| is zero, i.e. this ack event marks the start of a new ack
	// aggregation epoch, save LessRecentPoint, which is the last ack point of the
	// previous epoch, as a A0 candidate.
	if b.i4NwjwX59Gl && extraAcked == 0 {
		b.vFlrYPDp4.PushBack(*b.y_Z4NrytDBe4.LessRecentPoint())
	}
	return extraAcked
}

func sentPacketToSendTimeState(sentPacket *wzBRgOvr8BC, sendTimeState *iC6_vHy61ee) {
	*sendTimeState = sentPacket.ujKE8LV1OP
	sendTimeState.o00cUt = true
}

// BytesFromBandwidthAndTimeDelta calculates the bytes
// from a bandwidth(bits per second) and a time delta
func bytesFromBandwidthAndTimeDelta(bandwidth JaadYOR_Vaf6, delta time.Duration) congestion.ByteCount {
	return (congestion.ByteCount(bandwidth) * congestion.ByteCount(delta)) /
		(congestion.ByteCount(time.Second) * 8)
}
