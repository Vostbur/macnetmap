package main

import "testing"

//     157           6987664 ns/op          524695 B/op       1364 allocs/op
func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run()
	}
}
