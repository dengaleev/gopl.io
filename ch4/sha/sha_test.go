package sha_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopl.io/ch4/sha"
)

func TestByteDiff(t *testing.T) {
	for _, data := range []struct {
		lhs, rhs byte
		diff     int
	}{
		{
			0b00000000,
			0b00000000,
			8,
		},
		{
			0b10101010,
			0b00000000,
			4,
		},
		{
			0b10101010,
			0b01010101,
			0,
		},
		{
			0b11111111,
			0b00111000,
			3,
		},
	} {
		assert.Equal(t, data.diff, sha.BitsDiff(data.lhs, data.rhs))
	}
}
