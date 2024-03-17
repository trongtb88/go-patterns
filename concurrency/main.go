package main

import (
	"fmt"
	"time"
)

// simple channel
func doWork(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("DONE")
			return
		default:
			fmt.Println("DO WORKING")
		}
	}
}

// for pipeline
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// communicate between channel and go routine
	done := make(chan bool)
	go doWork(done)

	time.Sleep(2 * time.Second)
	// done after 2 second
	done <- true

	// pineline
	numbers := []int{4, 3, 2, 8, 9}

	fmt.Println("START PINELINE")

	// stage 1
	dataChannel := sliceToChannel(numbers)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
	fmt.Println("DONE PINELINE")

}
