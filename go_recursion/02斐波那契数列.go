package main

import "fmt"

func fibonaqi(n int) (result int) {
	if n < 2 {
		return n
	}
	result = fibonaqi(n-2) + fibonaqi(n-1)
	return result
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", fibonaqi(i))
	}
}
