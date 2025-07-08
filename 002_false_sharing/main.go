package main

import (
	"sync"
	"time"
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
	}
}
