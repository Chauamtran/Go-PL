package main

import (
	"fmt"
	"sync"
)

// var c = make(chan int, 10)
// var a string

// func f() {
// 	a = "hello, world"
// 	// c <- 0
// 	close(c)
// }
// func main()  {
// 	fmt.Println("Hello World")
// 	go f()
// 	<- c
// 	print(a)
// }

var a string
var done bool
var once sync.Once
var l sync.Mutex

func setup() {
	a = "hello, world"
	done = true
	l.Unlock()
}

func doprint() {
	fmt.Println(a)
	once.Do(setup)
	fmt.Println(a, done)
}

func twoprint() {
	fmt.Println("HELLO")
	go doprint()
	go doprint()
}

func main() {
	l.Lock()
	twoprint()
	fmt.Println("TESTING")
}
