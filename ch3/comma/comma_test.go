package comma_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopl.io/ch3/comma"
)

func TestComma(t *testing.T) {
	for _, data := range []struct {
		in, out string
	}{
		{"1", "1"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"999999999", "999,999,999"},
	} {
		assert.Equal(t, data.out, comma.CommaBuffer(data.in))
	}
}
