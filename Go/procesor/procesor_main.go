package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func cpuIntensiveTask() float64 {
	result := 0.0
	for i := 0; i < 100000; i++ {
		result += math.Sin(float64(i)) * math.Tan(float64(i))
	}
	return result
}

func main() {
	loadBefore := runtime.NumCPU()
	fmt.Printf("Użycie CPU przed skryptem: %d\n", loadBefore)

	runtime.GC() // Wyłącz GC

	start := time.Now()
	for i := 0; i < 1000; i++ {
		cpuIntensiveTask()
	}
	duration := time.Since(start)
	fmt.Printf("Czas wykonania bez GC: %v\n", duration)

	runtime.GC() // Włącz GC

	start = time.Now()
	for i := 0; i < 1000; i++ {
		cpuIntensiveTask()
	}
	duration = time.Since(start)
	fmt.Printf("Czas wykonania z GC: %v\n", duration)
}
