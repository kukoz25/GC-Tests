
package main

import (
	"runtime"
	"testing"
)

func BenchmarkCreateCycleMemoryWithoutGC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		createCycle()
	}
}

func BenchmarkCreateCycleMemoryWithGC(b *testing.B) {
	b.ReportAllocs()
	runtime.GC()
	for i := 0; i < b.N; i++ {
		createCycle()
	}
}
