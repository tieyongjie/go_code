package main

import "fmt"

// 通道是一个先进先出的队列
func main() {
	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	fmt.Printf("this is first value %v\n", <-channel)
	fmt.Printf("this is first value %v\n", <-channel)
	//fmt.Printf("this is sencode value %v\n", <-channel)
	fmt.Printf("this is first value %v\n", <-channel)
	fmt.Printf("this is first value %v\n", <-channel)
}
