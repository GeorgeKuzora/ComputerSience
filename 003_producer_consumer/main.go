package main

type Producer struct {}

func (p Producer) Produce(value int) int {
	return value
}

type Consumer struct {}

func (c Consumer) Consume() {}

func main() {

	var in <- chan int
	var out chan <- int

	in = make(chan int)
	out = make(chan int)

	producer := Producer{}

	input := []int {1,2,3,4,5,6,7,8,9}

	for _, v := range input {
		go func(value int) {

		}(v)
	}
}
