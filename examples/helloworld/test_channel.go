package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var l sync.Mutex

func createChannelData(in chan<- int, n int) {
	for i := 0; i < n; i++ {
		in <- i
		log.Println("i: ", i)
	}
}

func getChannelData(out <-chan int, n int) {
	// log.Println(<-out)
	// log.Println(<-out)
	// log.Println(len(out))
	// close(out)
	for i := 0; i < n; i++ {
		log.Println("out: ", <-out)
	}
}

func main() {
	n := 3
	in := make(chan int, n)
	out := make(chan int, n)
	// done := make(chan bool)

	createChannelData(in, n)

	// We now supply 2 channels to the `multiplyByTwo` function
	// One for sending data and one for receiving
	wg.Add(1)
	go multiplyByTwo(in, out)
	// go multiplyByTwo(in, out)
	// go multiplyByTwo(in, out)
	// We then send it data through the channel and wait for the result
	// log.Println(len(out))

	close(in)
	getChannelData(out, n)

	wg.Wait()

	// log.Println(<-out)
	// log.Println(<-out)
	// log.Println(<-out)
	// fmt.Println(<-out)
	// fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	// This line is just to illustrate that there is code that is
	// executed before we have to wait on the `in` channel

	log.Println("Initializing goroutine...")

	// The goroutine does not proceed until data is received on the `in` channel
	// num := <-in
	for num := range in {
		// wg.Add(1)

		result := num * 2
		// fmt.Println(result)
		// log.Println(result)
		out <- result

	}
	wg.Done()
	// l.Unlock()
	// done <- true

	// defer wg.Done()
	// The rest is unchanged
	// result := num * 2
	// result := n * 2
	// out <- result
}
