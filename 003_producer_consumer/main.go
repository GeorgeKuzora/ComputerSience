package main

import (
	"fmt"
	"time"
)

type Producer struct {
	id int
}

func (p Producer) Produce(value int) int {
	fmt.Printf("Producer with ID: %d produced VALUE: %d\n", p.id, value)
	return value
}

type Consumer struct {
	id int
}

func (c Consumer) Consume(value int) {
	fmt.Printf("Consumer with ID: %d consumed VALUE: %d\n", c.id, value)
}

func main() {

	out := make(chan int, 1000)

	producers := []Producer {
		Producer{id: 0},
		Producer{id: 1},
		Producer{id: 2},
		Producer{id: 3},
		Producer{id: 4},
	}
	consumers := []Consumer {
		Consumer{id: 0},
		Consumer{id: 1},
		Consumer{id: 2},
		Consumer{id: 3},
		Consumer{id: 4},
	}

	input := []int {0,1,2,3,4,5,6,7,8,9}

	for _, p := range producers {
		go func(p Producer, in []int, ch chan <- int) {
			for _, v := range in {
				ch <- p.Produce(v)
			}
		}(p, input, out)
	}


	for _, c := range consumers {
		go func(c Consumer, ch <- chan int) {
			start := time.Now()
			for {
				v, ok := <- ch
				if v == 0 && ok == false {
				 	fmt.Printf("CONSUMER %d: channel OUT was closed\n", c.id)
					return
				}
				dur_in_ms := time.Now().Sub(start).Microseconds()
				c.Consume(v)
				fmt.Printf("TOTAL TIME for CONSUMER %d VALUE %d: %d mrs\n", c.id, v, dur_in_ms)
			}
		}(c, out)
	}
	time.Sleep(1 * time.Second)
	close(out)
	time.Sleep(1 * time.Second)
	fmt.Println("End of the program")
}
