package main

import "fmt"

func main() {
	strings := []string{"goole", "baidu"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x的值 = %d\n", i, x)
	}
	myStr := [10]int{1, 10, 20, 30, 40, 50}
	fmt.Println(myStr)
}
