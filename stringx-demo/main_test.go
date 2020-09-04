package main

import (
	"testing"
)

/**
Benchmark_StringxReplace-8   	  229603	      4875 ns/op	    3688 B/op	      61 allocs/op
PASS
ok  	stringx-demo	1.591s
*/
func Benchmark_StringxReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringxReplace()
	}
}

/**
Benchmark_SysReplace-8   	 2138535	       558 ns/op	     288 B/op	       6 allocs/op
PASS
ok  	stringx-demo	2.243s
*/
func Benchmark_SysReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SysReplace()
	}
}
