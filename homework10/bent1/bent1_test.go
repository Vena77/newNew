package main

import "testing"

var glob uint

func BenchmarkFibb(b *testing.B) {
	var res uint
	for i := 0; i < b.N; i++ {
		res = fibbonach1(100)
	}
	glob = res
}
