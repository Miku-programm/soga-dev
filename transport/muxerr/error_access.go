package muxerr

// ErrorAt5098410 exposes the original error interface to its cross-package
// callers. The decompiler's address label is retained; this is the same error
// instance used by native legacy SS authentication, not a recreated error.
func ErrorAt5098410() error { return errorAt5098410 }
