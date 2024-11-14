package main

import (
	"fmt"
	"runtime"
)

type Node struct {
	ref *Node
}

func createCycle() {
	a := &Node{}
	b := &Node{}

	a.ref = b
	b.ref = a
}

func main() {
	runtime.GC() // Disable GC

	startMemory := runtime.MemStats{}
	runtime.ReadMemStats(&startMemory)
	for i := 0; i < 1000000; i++ {
		createCycle()
	}
	endMemory := runtime.MemStats{}
	runtime.ReadMemStats(&endMemory)
	fmt.Printf("Zużycie pamięci bez GC: %d bitów\n", endMemory.HeapAlloc-startMemory.HeapAlloc)

	runtime.GC() // Enable GC

	startMemory = runtime.MemStats{}
	runtime.ReadMemStats(&startMemory)
	for i := 0; i < 1000000; i++ {
		createCycle()
	}
	endMemory = runtime.MemStats{}
	runtime.ReadMemStats(&endMemory)
	fmt.Printf("Zużycie pamięci z GC: %d bitów\n", endMemory.HeapAlloc-startMemory.HeapAlloc)
}
