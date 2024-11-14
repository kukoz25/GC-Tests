package main

import (
	"runtime"
	"testing"
)

func BenchmarkCpuIntensiveTaskWithoutGC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cpuIntensiveTask()
	}
}

func BenchmarkCpuIntensiveTaskWithGC(b *testing.B) {
	b.ReportAllocs()
	runtime.GC()
	for i := 0; i < b.N; i++ {
		cpuIntensiveTask()
	}
}
