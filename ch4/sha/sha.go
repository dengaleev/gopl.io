package sha

import "gopl.io/ch2/popcount"

func BitsDiff(a, b byte) int {
	d := ^(a ^ b)
	return popcount.PopCount(uint64(d))
}

func SHABitsDiff(a, b [32]byte) (cnt int) {
	for i := range a {
		cnt += BitsDiff(a[i], b[i])
	}
	return
}
