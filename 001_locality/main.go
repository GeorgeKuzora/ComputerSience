package main

import (
	"fmt"
	"time"
)


func sumarrayrows(a [][]int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			sum += a[i][j];
		}
	}
	return sum
}

func sumarraycols(a [][]int) int {
	sum := 0
	for j := 0; j < len(a[0]); j++ {
		for i := 0; i < len(a); i++ {
			sum += a[i][j];
		}
	}
	return sum
}

func main() {
	a := make([][]int,0,10000)
	for range 10000 {
		b := make([]int,0,10000)
		for range 10000 {
			b = append(b, 1)
		}
		a = append(a, b)
	}
	start := time.Now()
	sum := sumarrayrows(a)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Rows sum %d, took %s\n", sum, delta)

	start = time.Now()
	sum = sumarraycols(a)
	end = time.Now()
	delta = end.Sub(start)
	fmt.Printf("Cols sum %d, took %s\n", sum, delta)
}
