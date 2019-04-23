package golisp

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

func BenchmarkList100(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	y := Pair{v, nil}
	for i := 0; i < 100; i++ {
		x = Pair{v, x}
		y = Pair{v, y}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x == y
	}
}

func BenchmarkList1k(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	y := Pair{v, nil}
	for i := 0; i < 1000; i++ {
		x = Pair{v, x}
		y = Pair{v, y}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x == y
	}
}

func BenchmarkList10k(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	y := Pair{v, nil}
	for i := 0; i < 10000; i++ {
		x = Pair{v, x}
		y = Pair{v, y}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x == y
	}
}

func BenchmarkList10kDiff(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	y := Pair{v, v}
	for i := 0; i < 10000; i++ {
		x = Pair{v, x}
		y = Pair{v, y}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x == y
	}
}

/* This shows that if the two interface chains differ early in
 * the chain, the comparison is very fast.  Vs. if the two interface
 * chains differ at the end, the comparison is slow.*/

func BenchmarkList10kDiffSoon(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	y := Pair{v, nil}
	for i := 0; i < 10000; i++ {
		if i == 9990 {
			x = Pair{v, x}
			y = Pair{Symbol("B"), y}
		} else {
			x = Pair{v, x}
			y = Pair{v, y}
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = x == y
	}
}

// These two tests show that, unlike ==, calling a function
// on these long chains of interfaces doesn't seem to incur
// a copy (even though with == they seem to have value semantics).
func BenchmarkCdr1k(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	for i := 0; i < 1000; i++ {
		x = Pair{v, x}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cdr(x)
	}
}

func BenchmarkCdr10k(b *testing.B) {
	v := Symbol("A")
	x := Pair{v, nil}
	for i := 0; i < 10000; i++ {
		x = Pair{v, x}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cdr(x)
	}
}
