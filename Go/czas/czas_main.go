package main

import (
	"fmt"
	"runtime"
	"time"
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

	start := time.Now()
	for i := 0; i < 1500000; i++ {
		createCycle()
	}
	duration := time.Since(start)
	fmt.Printf("Czas bez GC: %v\n", duration)

	runtime.GC() // Enable GC

	start = time.Now()
	for i := 0; i < 1500000; i++ {
		createCycle()
	}
	duration = time.Since(start)
	fmt.Printf("Czas z GC: %v\n", duration)
}
