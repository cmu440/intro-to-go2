package main

import "fmt"

const (
	N_INCREMENTS = 10000000
)

func main() {
	
	counter := 0
	donechan := make(chan bool)

	go func(done chan<- bool) {
		for i := 0; i < N_INCREMENTS; i++ {
			counter++
		}
		done <- true
	}(donechan)
	
	for i := 0; i < N_INCREMENTS; i++ {
		counter++
	}
	
	_ = <-donechan
	
	fmt.Println("Count: ", counter)

}