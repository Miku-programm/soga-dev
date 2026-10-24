package rand

// Tij6hx recovers 0xe8e520. It uses this package's crypto-random integer
// functions and the original Int63n/Int31n split; not math/rand's source.
func Tij6hx(n int, swap func(int, int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}
	i := n - 1
	for ; i > 0x7ffffffe; i-- {
		swap(i, int(TfHpoOcv(int64(i+1))))
	}
	for ; i > 0; i-- {
		swap(i, int(NMT2SOqf(int32(i+1))))
	}
}
