// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	var x IntSet
	assert.Equal(t, 0, x.Len())
	x.Add(4)
	x.Add(8)
	x.Add(15)
	x.Add(16)
	x.Add(23)
	x.Add(42)
	assert.Equal(t, 6, x.Len())

	x.Add(4)
	assert.Equal(t, 6, x.Len())
	x.Add(42)
	assert.Equal(t, 6, x.Len())

	x.Remove(15)
	assert.Equal(t, 5, x.Len())
	x.Remove(16)
	assert.Equal(t, 4, x.Len())

	y := x.Copy()

	x.Clear()
	assert.Equal(t, 0, x.Len())
	assert.Equal(t, 4, y.Len())
}
