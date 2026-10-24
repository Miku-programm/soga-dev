// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package hysteriabrutal

import (
	"golang.org/x/exp/constraints"
)

// Implements Kathleen Nichols' algorithm for tracking the minimum (or maximum)
// estimate of a stream of samples over some fixed time interval. (E.g.,
// the minimum RTT over the past five minutes.) The algorithm keeps track of
// the best, second best, and third best min (or max) estimates, maintaining an
// invariant that the measurement time of the n'th best >= n-1'th best.

// The algorithm works as follows. On a reset, all three estimates are set to
// the same sample. The second best estimate is then recorded in the second
// quarter of the window, and a third best estimate is recorded in the second
// half of the window, bounding the worst case error when the true min is
// monotonically increasing (or true max is monotonically decreasing) over the
// window.
//
// A new best sample replaces all three estimates, since the new best is lower
// (or higher) than everything else in the window and it is the most recent.
// The window thus effectively gets reset on every new min. The same property
// holds true for second best and third best estimates. Specifically, when a
// sample arrives that is better than the second best but not better than the
// best, it replaces the second and third best estimates but not the best
// estimate. Similarly, a sample that is better than the third best estimate
// but not the other estimates replaces only the third best estimate.
//
// Finally, when the best expires, it is replaced by the second best, which in
// turn is replaced by the third best. The newest sample replaces the third
// best.

type WindowedFilterValue interface {
	any
}

type WindowedFilterTime interface {
	constraints.Integer | constraints.Float
}

type Nqk7YFvzgcZg[V WindowedFilterValue, T WindowedFilterTime] struct {
	// Time length of window.
	aaF6ngDQX  T
	fPLAVE6X   []f2ypPZu[V, T]
	enH1PpcAyo func(V, V) int
}

type f2ypPZu[V WindowedFilterValue, T WindowedFilterTime] struct {
	nvK8flg3g V
	jgaVJt    T
}

// Compares two values and returns true if the first is greater than or equal
// to the second.
func Ez2Vlj[O constraints.Ordered](a, b O) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// Compares two values and returns true if the first is less than or equal
// to the second.

func NewWindowedFilter[V WindowedFilterValue, T WindowedFilterTime](windowLength T, comparator func(V, V) int) *Nqk7YFvzgcZg[V, T] {
	return &Nqk7YFvzgcZg[V, T]{
		aaF6ngDQX:  windowLength,
		fPLAVE6X:   make([]f2ypPZu[V, T], 3, 3),
		enH1PpcAyo: comparator,
	}
}

// Changes the window length.  Does not update any current samples.
func (f *Nqk7YFvzgcZg[V, T]) SetWindowLength(windowLength T) {
	f.aaF6ngDQX = windowLength
}

func (f *Nqk7YFvzgcZg[V, T]) GetBest() V {
	return f.fPLAVE6X[0].nvK8flg3g
}

func (f *Nqk7YFvzgcZg[V, T]) GetSecondBest() V {
	return f.fPLAVE6X[1].nvK8flg3g
}

func (f *Nqk7YFvzgcZg[V, T]) GetThirdBest() V {
	return f.fPLAVE6X[2].nvK8flg3g
}

// Updates best estimates with |sample|, and expires and updates best
// estimates as necessary.
func (f *Nqk7YFvzgcZg[V, T]) Update(newSample V, newTime T) {
	// Reset all estimates if they have not yet been initialized, if new sample
	// is a new best, or if the newest recorded estimate is too old.
	if f.enH1PpcAyo(f.fPLAVE6X[0].nvK8flg3g, *new(V)) == 0 ||
		f.enH1PpcAyo(newSample, f.fPLAVE6X[0].nvK8flg3g) >= 0 ||
		newTime-f.fPLAVE6X[2].jgaVJt > f.aaF6ngDQX {
		f.Reset(newSample, newTime)
		return
	}

	if f.enH1PpcAyo(newSample, f.fPLAVE6X[1].nvK8flg3g) >= 0 {
		f.fPLAVE6X[1] = f2ypPZu[V, T]{newSample, newTime}
		f.fPLAVE6X[2] = f.fPLAVE6X[1]
	} else if f.enH1PpcAyo(newSample, f.fPLAVE6X[2].nvK8flg3g) >= 0 {
		f.fPLAVE6X[2] = f2ypPZu[V, T]{newSample, newTime}
	}

	// Expire and update estimates as necessary.
	if newTime-f.fPLAVE6X[0].jgaVJt > f.aaF6ngDQX {
		// The best estimate hasn't been updated for an entire window, so promote
		// second and third best estimates.
		f.fPLAVE6X[0] = f.fPLAVE6X[1]
		f.fPLAVE6X[1] = f.fPLAVE6X[2]
		f.fPLAVE6X[2] = f2ypPZu[V, T]{newSample, newTime}
		// Need to iterate one more time. Check if the new best estimate is
		// outside the window as well, since it may also have been recorded a
		// long time ago. Don't need to iterate once more since we cover that
		// case at the beginning of the method.
		if newTime-f.fPLAVE6X[0].jgaVJt > f.aaF6ngDQX {
			f.fPLAVE6X[0] = f.fPLAVE6X[1]
			f.fPLAVE6X[1] = f.fPLAVE6X[2]
		}
		return
	}
	if f.enH1PpcAyo(f.fPLAVE6X[1].nvK8flg3g, f.fPLAVE6X[0].nvK8flg3g) == 0 &&
		newTime-f.fPLAVE6X[1].jgaVJt > f.aaF6ngDQX/4 {
		// A quarter of the window has passed without a better sample, so the
		// second-best estimate is taken from the second quarter of the window.
		f.fPLAVE6X[1] = f2ypPZu[V, T]{newSample, newTime}
		f.fPLAVE6X[2] = f.fPLAVE6X[1]
		return
	}

	if f.enH1PpcAyo(f.fPLAVE6X[2].nvK8flg3g, f.fPLAVE6X[1].nvK8flg3g) == 0 &&
		newTime-f.fPLAVE6X[2].jgaVJt > f.aaF6ngDQX/2 {
		// We've passed a half of the window without a better estimate, so take
		// a third-best estimate from the second half of the window.
		f.fPLAVE6X[2] = f2ypPZu[V, T]{newSample, newTime}
	}
}

// Resets all estimates to new sample.
func (f *Nqk7YFvzgcZg[V, T]) Reset(newSample V, newTime T) {
	f.fPLAVE6X[2] = f2ypPZu[V, T]{newSample, newTime}
	f.fPLAVE6X[1] = f.fPLAVE6X[2]
	f.fPLAVE6X[0] = f.fPLAVE6X[1]
}

func (f *Nqk7YFvzgcZg[V, T]) Clear() {
	f.fPLAVE6X = make([]f2ypPZu[V, T], 3, 3)
}
