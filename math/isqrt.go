package math

import (
	omath "math"
)

// ISqrt returns floor(sqrt(n)). Typical run time is few hundreds of ns.
func ISqrt(n uint32) (x uint32) {
	if n == 0 {
		return
	}

	if n >= omath.MaxUint16*omath.MaxUint16 {
		return omath.MaxUint16
	}

	var px, nx uint32
	for x = n; ; px, x = x, nx {
		nx = (x + n/x) / 2
		if nx == x || nx == px {
			break
		}
	}
	return
}
