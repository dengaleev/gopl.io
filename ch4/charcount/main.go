// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func printCountMap(name string, counts map[rune]int) {
	fmt.Printf("%s counts:\n", name)
	fmt.Printf("\trune\tcount\n")
	for c, n := range counts {
		fmt.Printf("\t%q\t%d\n", c, n)
	}
}

func main() {
	letters := make(map[rune]int)   // counts of Unicode letters
	digits := make(map[rune]int)    // counts of Unicode digits
	other := make(map[rune]int)     // counts of Unicode digits
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsDigit(r):
			digits[r]++
		case unicode.IsLetter(r):
			letters[r]++
		default:
			other[r]++
		}
		utflen[n]++
	}
	printCountMap("Digits", digits)
	printCountMap("Letters", letters)
	printCountMap("Other", other)
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
