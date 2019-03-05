package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	in := make(chan int)

	// Create 3 `multiplyByTwo` goroutines.
	go no(in, out, "3")

	go no(in, out, "1")
	fmt.Println("000000000")
	go no(in, out, "2")
	go no(in, out, "0")
	//go multiplyByTwo(in, out, "1")
	//go multiplyByTwo(in, out, "2")
	//go multiplyByTwo(in, out, "4")

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data
	in <- 1
	in <- 2
	in <- 3
	in <- 4

	// Now we wait for each result to come in
	fmt.Println(".....")
	fmt.Println(<-out)
	fmt.Println(".....")
	fmt.Println(<-out)
	fmt.Println(".....")
	fmt.Println(<-out)
	fmt.Println(".....")
	fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int, i string) {
	fmt.Printf("Initializing goroutine... %s", i)
	fmt.Println(".")
	num := <-in
	result := num * 2
	out <- result
}

func no(in <-chan int, out chan<- int, i string) {
	if i == "0" {
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("potato %s", i)
	fmt.Println(".")
	num := <-in
	fmt.Println("!")
	fmt.Println(num)
	fmt.Println("?")

	out <- num
}
