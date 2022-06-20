package main

import "fmt"

func max(num1, num2 int) int {
	var result int
	if num1 < num2 {
		result = num2
	} else {
		result = num1
	}
	return result
}

func main() {
	a := 100
	b := 200
	var ret int
	ret = max(a, b)
	fmt.Printf("最大值是: %d", ret)
}
