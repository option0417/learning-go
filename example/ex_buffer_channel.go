package main

import "fmt"
import "time"

func put(c chan int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Put\n")
		c <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func get(c chan int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Get\n")
		fmt.Printf("%v\n", <-c)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 2)
	go put(ch)
	go get(ch)
	time.Sleep(1000 * time.Millisecond)
}
