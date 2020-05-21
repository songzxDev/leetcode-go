package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 3
	ch1 <- 1
	ch1 <- 2
	elemt := <-ch1
	fmt.Printf("%v\n", elemt)
}
