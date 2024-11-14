package main

import (
	"runtime"
	"testing"
)

func BenchmarkCreateCycleWithoutGC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		createCycle()
	}
}

func BenchmarkCreateCycleWithGC(b *testing.B) {
	b.ReportAllocs()
	runtime.GC()
	for i := 0; i < b.N; i++ {
		createCycle()
	}
}
