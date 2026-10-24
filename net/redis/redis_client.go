// The original yd0VtxczG client is an obfuscated go-redis/v8 dependency.
// RTTI Options224B/Client56B, NewClient @0xf7ba00, defaults @0xf61f80,
// pool construction @0xf62220 and TLS/plain dialer @0xf8c280 match v8.11.5.
// This is a verified dependency/API mapping, not proof of the original patch
// version. Public dependency field names are used for the erased identifiers.
package redis

import "github.com/go-redis/redis/v8"

type HXzAD33aVBDR = redis.Client
type Csx1an = redis.Options

func ZHNMa9Nz(options *Csx1an) *HXzAD33aVBDR { return redis.NewClient(options) }
