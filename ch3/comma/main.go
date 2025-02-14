// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package comma

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

const (
	segmentLen = 3
	commaChar  = ','
)

func CommaBuffer(s string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		buf.WriteString(s[:1])
		s = s[1:]
	}

	i := len(s) % segmentLen
	if i == 0 {
		i = segmentLen
	}
	buf.WriteString(s[:i])
	for ; i < len(s); i += segmentLen {
		buf.WriteRune(commaChar)
		buf.WriteString(s[i : i+segmentLen])
	}
	return buf.String()
}
