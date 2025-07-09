package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)


type Unpadded struct {
	counter1 int64
	counter2 int64
}

type Padded struct {
	counter1 int64
	_ [8]int64
	counter2 int64
}

func runTest(usePadded bool, iterations int) time.Duration {
	var start time.Time
	var elapsed time.Duration

	if usePadded {
		p := &Padded{}
		start = time.Now()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&p.counter1, 1)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&p.counter2, 1)
			}
			wg.Done()
		}()
		wg.Wait()
		elapsed = time.Since(start)
	} else {
		u := &Unpadded{}
		start = time.Now()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&u.counter1, 1)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&u.counter2, 1)
			}
			wg.Done()
		}()
		wg.Wait()
		elapsed = time.Since(start)
	}
	return elapsed
}


func main() {
	const iterations = 1_000_000_000

	fmt.Println("Cache line size (bytes):", unsafe.Sizeof(Padded{}) - unsafe.Sizeof(int64(0))*2)

	unpaddedTime := runTest(false, iterations)
	fmt.Printf("Unpadded (false sharing): %v\n", unpaddedTime)

	paddedTime := runTest(true, iterations)
	fmt.Printf("Padded (no false sharing): %v\n", paddedTime)

	speedup := float64(unpaddedTime) / float64(paddedTime)
	fmt.Printf("Speedup: %.2fx\n", speedup)
}
